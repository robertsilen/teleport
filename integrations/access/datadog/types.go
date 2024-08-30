/*
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package datadog

// Datadog API types

type Metadata struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type PermissionsBody struct {
	Data []PermissionsData `json:"data,omitempty"`
}

type PermissionsData struct {
	Metadata
	Attributes PermissionsAttributes `json:"attributes,omitempty"`
}

type PermissionsAttributes struct {
	Name       string `json:"name,omitempty"`
	Restricted bool   `json:"restricted"`
}

type IncidentsBody struct {
	Data IncidentsData `json:"data,omitempty"`
}

type IncidentsData struct {
	Metadata
	Attributes IncidentsAttributes `json:"attributes,omitempty"`
}

type IncidentsAttributes struct {
	Title               string               `json:"title,omitempty"`
	Fields              IncidentsFields      `json:"fields,omitempty"`
	NotificationHandles []NotificationHandle `json:"notification_handles,omitempty"`
}

type IncidentsFields struct {
	Summary         *StringField      `json:"summary,omitempty"`
	Severity        *StringField      `json:"severity,omitempty"`
	State           *StringField      `json:"state,omitempty"`
	DetectionMethod *StringField      `json:"detection_method,omitempty"`
	RootCause       *StringField      `json:"root_cause,omitempty"`
	Teams           *StringSliceField `json:"teams,omitempty"`
	Services        *StringSliceField `json:"services,omitempty"`
}

type StringField struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type StringSliceField struct {
	Type  string   `json:"type,omitempty"`
	Value []string `json:"value,omitempty"`
}

type NotificationHandle struct {
	DisplayName string `json:"display_name,omitempty"`
	Handle      string `json:"handle,omitempty"`
}

type TimelineBody struct {
	Data TimelineData `json:"data,omitempty"`
}

type TimelineData struct {
	Metadata
	Attributes TimelineAttributes `json:"attributes,omitempty"`
}

type TimelineAttributes struct {
	CellType string          `json:"cell_type,omitempty"`
	Content  TimelineContent `json:"content,omitempty"`
}

type TimelineContent struct {
	Content string `json:"content,omitempty"`
}

type OncallTeamsBody struct {
	Data     []OncallTeamsData     `json:"data,omitempty"`
	Included []OncallTeamsIncluded `json:"included,omitempty"`
}

type OncallTeamsData struct {
	Metadata
	Attributes    OncallTeamsAttributes    `json:"attributes,omitempty"`
	Relationships OncallTeamsRelationships `json:"relationships,omitempty"`
}

type OncallTeamsAttributes struct {
	Name   string `json:"name,omitempty"`
	Handle string `json:"handle,omitempty"`
}

type OncallTeamsRelationships struct {
	OncallUsers OncallUsers `json:"oncall_users,omitempty"`
}

type OncallUsers struct {
	Data []OncallUsersData `json:"data,omitempty"`
}

type OncallUsersData struct {
	Metadata
}

type OncallTeamsIncluded struct {
	Metadata
	Attributes OncallTeamsIncludedAttributes `json:"attributes,omitempty"`
}

type OncallTeamsIncludedAttributes struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type UsersBody struct {
	Data []UsersData `json:"data,omitempty"`
}

type UsersData struct {
	Metadata
	Attributes UsersAttributes `json:"attributes,omitempty"`
}

type UsersAttributes struct {
	Name     string `json:"name,omitempty"`
	Handle   string `json:"handle,omitempty"`
	Email    string `json:"email,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type ErrorResult struct {
	Errors []string `json:"errors"`
}
