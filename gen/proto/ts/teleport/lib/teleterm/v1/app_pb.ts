/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/lib/teleterm/v1/app.proto" (package "teleport.lib.teleterm.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
//
// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Label } from "./label_pb";
/**
 * App describes an app resource.
 *
 * @generated from protobuf message teleport.lib.teleterm.v1.App
 */
export interface App {
    /**
     * uri uniquely identifies an app within Teleport Connect.
     *
     * @generated from protobuf field: string uri = 1;
     */
    uri: string;
    /**
     * name is the name of the app.
     *
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * endpoint_uri is the URI to which the app service is going to proxy requests. It corresponds to
     * app_service.apps[].uri in the Teleport configuration.
     *
     * @generated from protobuf field: string endpoint_uri = 3;
     */
    endpointUri: string;
    /**
     * desc is the app description.
     *
     * @generated from protobuf field: string desc = 4;
     */
    desc: string;
    /**
     * aws_console is true if this app is AWS management console.
     *
     * @generated from protobuf field: bool aws_console = 5;
     */
    awsConsole: boolean;
    /**
     * public_addr is the public address the application is accessible at.
     *
     * If the app resource has its public_addr field set, this field returns the value of public_addr
     * from the app resource.
     *
     * If the app resource does not have public_addr field set, this field returns the name of the app
     * under the proxy hostname of the cluster to which the app belongs, e.g.,
     * dumper.root-cluster.com, example-app.leaf-cluster.org.
     *
     * In both cases public_addr does not include a port number. An app resource cannot include a port
     * number in its public_addr, as the backend will (wrongly) reject such resource with an error
     * saying "public_addr "example.com:1337" can not contain a port, applications will be available
     * on the same port as the web proxy". This is obviously not the case for custom public addresses.
     * Ultimately, it means that public_addr alone is not enough to access the app, unless either the
     * cluster or the custom domain use the default port of 443.
     *
     * Always empty for SAML applications.
     *
     * @generated from protobuf field: string public_addr = 6;
     */
    publicAddr: string;
    /**
     * friendly_name is a user readable name of the app.
     * Right now, it is set only for Okta applications.
     * It is constructed from a label value.
     * See more in api/types/resource.go.
     *
     * @generated from protobuf field: string friendly_name = 7;
     */
    friendlyName: string;
    /**
     * saml_app is true if the application is a SAML Application (Service Provider).
     *
     * @generated from protobuf field: bool saml_app = 8;
     */
    samlApp: boolean;
    /**
     * labels is a list of labels for the app.
     *
     * @generated from protobuf field: repeated teleport.lib.teleterm.v1.Label labels = 9;
     */
    labels: Label[];
    /**
     * fqdn is the hostname under which the app is accessible within the root cluster. It is used by
     * the Web UI to route the requests from the /web/launch URL to the correct app. fqdn by itself
     * does not include the port number, so fqdn alone cannot be used to launch an app, hence why it's
     * incorporated into the /web/launch URL.
     *
     * If the app belongs to a root cluster, fqdn is equal to public_addr or [name].[root cluster
     * proxy hostname] if public_addr is not present.
     * If the app belongs to a leaf cluster, fqdn is equal to [name].[root cluster proxy hostname].
     *
     * fqdn is not present for SAML applications.
     *
     * @generated from protobuf field: string fqdn = 10;
     */
    fqdn: string;
    /**
     * aws_roles is a list of AWS IAM roles for the application representing AWS console.
     *
     * @generated from protobuf field: repeated teleport.lib.teleterm.v1.AWSRole aws_roles = 11;
     */
    awsRoles: AWSRole[];
}
/**
 * AwsRole describes AWS IAM role.
 *
 * @generated from protobuf message teleport.lib.teleterm.v1.AWSRole
 */
export interface AWSRole {
    /**
     * Name is the full role name with the entire path.
     *
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * Display is the role display name.
     *
     * @generated from protobuf field: string display = 2;
     */
    display: string;
    /**
     * ARN is the full role ARN.
     *
     * @generated from protobuf field: string arn = 3;
     */
    arn: string;
    /**
     * AccountID is the AWS Account ID this role refers to.
     *
     * @generated from protobuf field: string account_id = 4;
     */
    accountId: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class App$Type extends MessageType<App> {
    constructor() {
        super("teleport.lib.teleterm.v1.App", [
            { no: 1, name: "uri", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "endpoint_uri", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "desc", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "aws_console", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "public_addr", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "friendly_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "saml_app", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "labels", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Label },
            { no: 10, name: "fqdn", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "aws_roles", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => AWSRole }
        ]);
    }
    create(value?: PartialMessage<App>): App {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.uri = "";
        message.name = "";
        message.endpointUri = "";
        message.desc = "";
        message.awsConsole = false;
        message.publicAddr = "";
        message.friendlyName = "";
        message.samlApp = false;
        message.labels = [];
        message.fqdn = "";
        message.awsRoles = [];
        if (value !== undefined)
            reflectionMergePartial<App>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: App): App {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string uri */ 1:
                    message.uri = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* string endpoint_uri */ 3:
                    message.endpointUri = reader.string();
                    break;
                case /* string desc */ 4:
                    message.desc = reader.string();
                    break;
                case /* bool aws_console */ 5:
                    message.awsConsole = reader.bool();
                    break;
                case /* string public_addr */ 6:
                    message.publicAddr = reader.string();
                    break;
                case /* string friendly_name */ 7:
                    message.friendlyName = reader.string();
                    break;
                case /* bool saml_app */ 8:
                    message.samlApp = reader.bool();
                    break;
                case /* repeated teleport.lib.teleterm.v1.Label labels */ 9:
                    message.labels.push(Label.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string fqdn */ 10:
                    message.fqdn = reader.string();
                    break;
                case /* repeated teleport.lib.teleterm.v1.AWSRole aws_roles */ 11:
                    message.awsRoles.push(AWSRole.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: App, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string uri = 1; */
        if (message.uri !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.uri);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* string endpoint_uri = 3; */
        if (message.endpointUri !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.endpointUri);
        /* string desc = 4; */
        if (message.desc !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.desc);
        /* bool aws_console = 5; */
        if (message.awsConsole !== false)
            writer.tag(5, WireType.Varint).bool(message.awsConsole);
        /* string public_addr = 6; */
        if (message.publicAddr !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.publicAddr);
        /* string friendly_name = 7; */
        if (message.friendlyName !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.friendlyName);
        /* bool saml_app = 8; */
        if (message.samlApp !== false)
            writer.tag(8, WireType.Varint).bool(message.samlApp);
        /* repeated teleport.lib.teleterm.v1.Label labels = 9; */
        for (let i = 0; i < message.labels.length; i++)
            Label.internalBinaryWrite(message.labels[i], writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        /* string fqdn = 10; */
        if (message.fqdn !== "")
            writer.tag(10, WireType.LengthDelimited).string(message.fqdn);
        /* repeated teleport.lib.teleterm.v1.AWSRole aws_roles = 11; */
        for (let i = 0; i < message.awsRoles.length; i++)
            AWSRole.internalBinaryWrite(message.awsRoles[i], writer.tag(11, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.v1.App
 */
export const App = new App$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AWSRole$Type extends MessageType<AWSRole> {
    constructor() {
        super("teleport.lib.teleterm.v1.AWSRole", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "display", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "arn", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "account_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<AWSRole>): AWSRole {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.name = "";
        message.display = "";
        message.arn = "";
        message.accountId = "";
        if (value !== undefined)
            reflectionMergePartial<AWSRole>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AWSRole): AWSRole {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string display */ 2:
                    message.display = reader.string();
                    break;
                case /* string arn */ 3:
                    message.arn = reader.string();
                    break;
                case /* string account_id */ 4:
                    message.accountId = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: AWSRole, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string display = 2; */
        if (message.display !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.display);
        /* string arn = 3; */
        if (message.arn !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.arn);
        /* string account_id = 4; */
        if (message.accountId !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.accountId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.v1.AWSRole
 */
export const AWSRole = new AWSRole$Type();