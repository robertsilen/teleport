package azureoidc

import (
	"context"
	"net/url"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/gravitational/trace"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/applicationtemplates"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
)

// A special application template ID in Microsoft Graph, equivalent to the "create your own application" option in Azure portal.
// Only non-gallery apps ("Create your own application" option in the UI) are allowed to use SAML SSO,
// hence we use this template.
// Ref: https://learn.microsoft.com/en-us/graph/api/applicationtemplate-instantiate
const nonGalleryAppTemplateID = "8adf8e6e-67b2-4cf2-a259-e3dc5476c621"

// A list of Microsoft Graph permissions ("app roles") for directory sync as performed by the Entra ID plugin.
// Ref: https://learn.microsoft.com/en-us/graph/permissions-reference
var appRoles = []string{
	// Application.Read.All
	"9a5d68dd-52b0-4cc2-bd40-abcf44ac3a30",
	// Directory.Read.All
	"7ab1d382-f21e-4acd-a863-ba3e13f7da61",
	// Policy.Read.All
	"246dd0d5-5bd0-4def-940b-0421030a5b68",
}

// SetupEnterpriseApp sets up an Enterprise Application in the Entra ID directory.
// The enterprise application:
//   - Provides Teleport with OIDC authentication to Azure
//   - Is given the permissions to access certain Microsoft Graph API endpoints for this tenant.
//   - Provides SSO to the Teleport cluster via SAML.
func SetupEnterpriseApp(ctx context.Context, proxyPublicAddr string, authConnectorName string) (string, string, error) {
	var appID, tenantID string

	tenantID, err := getTenantID()
	if err != nil {
		return appID, tenantID, trace.Wrap(err)
	}

	graphClient, err := createGraphClient()
	if err != nil {
		return appID, tenantID, trace.Wrap(err)
	}

	displayName := "Teleport" + " " + strings.TrimPrefix(proxyPublicAddr, "https://")

	instantiateRequest := applicationtemplates.NewItemInstantiatePostRequestBody()
	instantiateRequest.SetDisplayName(&displayName)
	appAndSP, err := graphClient.ApplicationTemplates().
		ByApplicationTemplateId(nonGalleryAppTemplateID).
		Instantiate().
		Post(ctx, instantiateRequest, nil)

	if err != nil {
		return appID, tenantID, trace.Wrap(err, "failed to instantiate application template")
	}

	app := appAndSP.GetApplication()
	sp := appAndSP.GetServicePrincipal()
	appID = *app.GetAppId()
	spID := *sp.GetId()

	msGraphResourceID, err := getMSGraphResourceID(ctx, graphClient)
	if err != nil {
		return appID, tenantID, trace.Wrap(err, "failed to get MS Graph API resource ID")
	}

	msGraphResourceUUID := uuid.MustParse(msGraphResourceID)

	for _, appRoleID := range appRoles {
		assignment := models.NewAppRoleAssignment()
		spUUID := uuid.MustParse(spID)
		assignment.SetPrincipalId(&spUUID)

		assignment.SetResourceId(&msGraphResourceUUID)

		appRoleUUID := uuid.MustParse(appRoleID)
		assignment.SetAppRoleId(&appRoleUUID)
		_, err := graphClient.ServicePrincipals().
			ByServicePrincipalId(spID).
			AppRoleAssignments().
			Post(ctx, assignment, nil)
		if err != nil {
			return appID, tenantID, trace.Wrap(err, "failed to assign app role %s", appRoleID)
		}
	}

	if err := createFederatedAuthCredential(ctx, graphClient, *app.GetId(), proxyPublicAddr); err != nil {
		return appID, tenantID, trace.Wrap(err, "failed to create an OIDC federated auth credential")
	}

	acsURL, err := url.Parse(proxyPublicAddr)
	if err != nil {
		return appID, tenantID, trace.Wrap(err, "failed to parse proxy public address")
	}
	acsURL.Path = path.Join("/v1/webapi/saml/acs", authConnectorName)
	if err := setupSSO(ctx, graphClient, *app.GetId(), spID, acsURL.String()); err != nil {
		return appID, tenantID, trace.Wrap(err, "failed to set up SSO for the enterprise app")
	}

	return appID, tenantID, nil
}

// createFederatedAuthCredential creates a new federated (OIDC) auth credential for the given Entra application.
func createFederatedAuthCredential(ctx context.Context, graphClient *msgraphsdk.GraphServiceClient, appObjectID string, proxyPublicAddr string) error {
	credential := models.NewFederatedIdentityCredential()
	name := "teleport-oidc"
	audiences := []string{azureDefaultJWTAudience}
	subject := azureSubject
	credential.SetName(&name)
	credential.SetIssuer(&proxyPublicAddr)
	credential.SetAudiences(audiences)
	credential.SetSubject(&subject)

	// ByApplicationID here means the object ID,
	// i.e. app.GetId(), not app.GetAppId().
	_, err := graphClient.Applications().ByApplicationId(appObjectID).
		FederatedIdentityCredentials().Post(ctx, credential, nil)

	return trace.Wrap(err)

}

// getMSGraphResourceID gets the resource ID for the Microsoft Graph app in the Entra directory.
func getMSGraphResourceID(ctx context.Context, graphClient *msgraphsdk.GraphServiceClient) (string, error) {
	requestFilter := "displayName eq 'Microsoft Graph'"

	requestParameters := &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
		Filter: &requestFilter,
	}
	configuration := &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	spResponse, err := graphClient.ServicePrincipals().Get(ctx, configuration)
	if err != nil {
		return "", trace.Wrap(err)
	}

	spList := spResponse.GetValue()
	if len(spList) < 1 {
		return "", trace.NotFound("Microsoft Graph app not found in the tenant")
	}

	return *spList[0].GetId(), nil
}
