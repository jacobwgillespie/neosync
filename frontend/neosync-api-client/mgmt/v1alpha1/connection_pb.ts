// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file mgmt/v1alpha1/connection.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message mgmt.v1alpha1.GetConnectionsRequest
 */
export class GetConnectionsRequest extends Message<GetConnectionsRequest> {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId = "";

  constructor(data?: PartialMessage<GetConnectionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetConnectionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConnectionsRequest {
    return new GetConnectionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConnectionsRequest {
    return new GetConnectionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConnectionsRequest {
    return new GetConnectionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetConnectionsRequest | PlainMessage<GetConnectionsRequest> | undefined, b: GetConnectionsRequest | PlainMessage<GetConnectionsRequest> | undefined): boolean {
    return proto3.util.equals(GetConnectionsRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetConnectionsResponse
 */
export class GetConnectionsResponse extends Message<GetConnectionsResponse> {
  /**
   * @generated from field: repeated mgmt.v1alpha1.Connection connections = 1;
   */
  connections: Connection[] = [];

  constructor(data?: PartialMessage<GetConnectionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetConnectionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "connections", kind: "message", T: Connection, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConnectionsResponse {
    return new GetConnectionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConnectionsResponse {
    return new GetConnectionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConnectionsResponse {
    return new GetConnectionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetConnectionsResponse | PlainMessage<GetConnectionsResponse> | undefined, b: GetConnectionsResponse | PlainMessage<GetConnectionsResponse> | undefined): boolean {
    return proto3.util.equals(GetConnectionsResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetConnectionRequest
 */
export class GetConnectionRequest extends Message<GetConnectionRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetConnectionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetConnectionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConnectionRequest {
    return new GetConnectionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConnectionRequest {
    return new GetConnectionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConnectionRequest {
    return new GetConnectionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetConnectionRequest | PlainMessage<GetConnectionRequest> | undefined, b: GetConnectionRequest | PlainMessage<GetConnectionRequest> | undefined): boolean {
    return proto3.util.equals(GetConnectionRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.GetConnectionResponse
 */
export class GetConnectionResponse extends Message<GetConnectionResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.Connection connection = 1;
   */
  connection?: Connection;

  constructor(data?: PartialMessage<GetConnectionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.GetConnectionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "connection", kind: "message", T: Connection },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetConnectionResponse {
    return new GetConnectionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetConnectionResponse {
    return new GetConnectionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetConnectionResponse {
    return new GetConnectionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetConnectionResponse | PlainMessage<GetConnectionResponse> | undefined, b: GetConnectionResponse | PlainMessage<GetConnectionResponse> | undefined): boolean {
    return proto3.util.equals(GetConnectionResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.CreateConnectionRequest
 */
export class CreateConnectionRequest extends Message<CreateConnectionRequest> {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: mgmt.v1alpha1.ConnectionConfig connection_config = 3;
   */
  connectionConfig?: ConnectionConfig;

  constructor(data?: PartialMessage<CreateConnectionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CreateConnectionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "connection_config", kind: "message", T: ConnectionConfig },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConnectionRequest {
    return new CreateConnectionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConnectionRequest {
    return new CreateConnectionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConnectionRequest {
    return new CreateConnectionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateConnectionRequest | PlainMessage<CreateConnectionRequest> | undefined, b: CreateConnectionRequest | PlainMessage<CreateConnectionRequest> | undefined): boolean {
    return proto3.util.equals(CreateConnectionRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.CreateConnectionResponse
 */
export class CreateConnectionResponse extends Message<CreateConnectionResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.Connection connection = 1;
   */
  connection?: Connection;

  constructor(data?: PartialMessage<CreateConnectionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CreateConnectionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "connection", kind: "message", T: Connection },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConnectionResponse {
    return new CreateConnectionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConnectionResponse {
    return new CreateConnectionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConnectionResponse {
    return new CreateConnectionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateConnectionResponse | PlainMessage<CreateConnectionResponse> | undefined, b: CreateConnectionResponse | PlainMessage<CreateConnectionResponse> | undefined): boolean {
    return proto3.util.equals(CreateConnectionResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.UpdateConnectionRequest
 */
export class UpdateConnectionRequest extends Message<UpdateConnectionRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: mgmt.v1alpha1.ConnectionConfig connection_config = 2;
   */
  connectionConfig?: ConnectionConfig;

  constructor(data?: PartialMessage<UpdateConnectionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.UpdateConnectionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "connection_config", kind: "message", T: ConnectionConfig },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConnectionRequest {
    return new UpdateConnectionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConnectionRequest {
    return new UpdateConnectionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConnectionRequest {
    return new UpdateConnectionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConnectionRequest | PlainMessage<UpdateConnectionRequest> | undefined, b: UpdateConnectionRequest | PlainMessage<UpdateConnectionRequest> | undefined): boolean {
    return proto3.util.equals(UpdateConnectionRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.UpdateConnectionResponse
 */
export class UpdateConnectionResponse extends Message<UpdateConnectionResponse> {
  /**
   * @generated from field: mgmt.v1alpha1.Connection connection = 1;
   */
  connection?: Connection;

  constructor(data?: PartialMessage<UpdateConnectionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.UpdateConnectionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "connection", kind: "message", T: Connection },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConnectionResponse {
    return new UpdateConnectionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConnectionResponse {
    return new UpdateConnectionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConnectionResponse {
    return new UpdateConnectionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateConnectionResponse | PlainMessage<UpdateConnectionResponse> | undefined, b: UpdateConnectionResponse | PlainMessage<UpdateConnectionResponse> | undefined): boolean {
    return proto3.util.equals(UpdateConnectionResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.DeleteConnectionRequest
 */
export class DeleteConnectionRequest extends Message<DeleteConnectionRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteConnectionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.DeleteConnectionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConnectionRequest {
    return new DeleteConnectionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConnectionRequest {
    return new DeleteConnectionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConnectionRequest {
    return new DeleteConnectionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteConnectionRequest | PlainMessage<DeleteConnectionRequest> | undefined, b: DeleteConnectionRequest | PlainMessage<DeleteConnectionRequest> | undefined): boolean {
    return proto3.util.equals(DeleteConnectionRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.DeleteConnectionResponse
 */
export class DeleteConnectionResponse extends Message<DeleteConnectionResponse> {
  constructor(data?: PartialMessage<DeleteConnectionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.DeleteConnectionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConnectionResponse {
    return new DeleteConnectionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConnectionResponse {
    return new DeleteConnectionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConnectionResponse {
    return new DeleteConnectionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteConnectionResponse | PlainMessage<DeleteConnectionResponse> | undefined, b: DeleteConnectionResponse | PlainMessage<DeleteConnectionResponse> | undefined): boolean {
    return proto3.util.equals(DeleteConnectionResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.CheckConnectionConfigRequest
 */
export class CheckConnectionConfigRequest extends Message<CheckConnectionConfigRequest> {
  /**
   * @generated from field: mgmt.v1alpha1.ConnectionConfig connection_config = 1;
   */
  connectionConfig?: ConnectionConfig;

  constructor(data?: PartialMessage<CheckConnectionConfigRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CheckConnectionConfigRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "connection_config", kind: "message", T: ConnectionConfig },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckConnectionConfigRequest {
    return new CheckConnectionConfigRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckConnectionConfigRequest {
    return new CheckConnectionConfigRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckConnectionConfigRequest {
    return new CheckConnectionConfigRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CheckConnectionConfigRequest | PlainMessage<CheckConnectionConfigRequest> | undefined, b: CheckConnectionConfigRequest | PlainMessage<CheckConnectionConfigRequest> | undefined): boolean {
    return proto3.util.equals(CheckConnectionConfigRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.CheckConnectionConfigResponse
 */
export class CheckConnectionConfigResponse extends Message<CheckConnectionConfigResponse> {
  /**
   * @generated from field: bool is_connected = 1;
   */
  isConnected = false;

  /**
   * @generated from field: optional string connection_error = 2;
   */
  connectionError?: string;

  constructor(data?: PartialMessage<CheckConnectionConfigResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.CheckConnectionConfigResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "is_connected", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "connection_error", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckConnectionConfigResponse {
    return new CheckConnectionConfigResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckConnectionConfigResponse {
    return new CheckConnectionConfigResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckConnectionConfigResponse {
    return new CheckConnectionConfigResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CheckConnectionConfigResponse | PlainMessage<CheckConnectionConfigResponse> | undefined, b: CheckConnectionConfigResponse | PlainMessage<CheckConnectionConfigResponse> | undefined): boolean {
    return proto3.util.equals(CheckConnectionConfigResponse, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.Connection
 */
export class Connection extends Message<Connection> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: mgmt.v1alpha1.ConnectionConfig connection_config = 3;
   */
  connectionConfig?: ConnectionConfig;

  /**
   * @generated from field: string created_by_user_id = 4;
   */
  createdByUserId = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string updated_by_user_id = 6;
   */
  updatedByUserId = "";

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Connection>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.Connection";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "connection_config", kind: "message", T: ConnectionConfig },
    { no: 4, name: "created_by_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "updated_by_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Connection {
    return new Connection().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Connection {
    return new Connection().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Connection {
    return new Connection().fromJsonString(jsonString, options);
  }

  static equals(a: Connection | PlainMessage<Connection> | undefined, b: Connection | PlainMessage<Connection> | undefined): boolean {
    return proto3.util.equals(Connection, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.ConnectionConfig
 */
export class ConnectionConfig extends Message<ConnectionConfig> {
  /**
   * @generated from oneof mgmt.v1alpha1.ConnectionConfig.config
   */
  config: {
    /**
     * @generated from field: mgmt.v1alpha1.PostgresConnectionConfig pg_config = 1;
     */
    value: PostgresConnectionConfig;
    case: "pgConfig";
  } | {
    /**
     * @generated from field: mgmt.v1alpha1.AwsS3ConnectionConfig aws_s3_config = 2;
     */
    value: AwsS3ConnectionConfig;
    case: "awsS3Config";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<ConnectionConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.ConnectionConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pg_config", kind: "message", T: PostgresConnectionConfig, oneof: "config" },
    { no: 2, name: "aws_s3_config", kind: "message", T: AwsS3ConnectionConfig, oneof: "config" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConnectionConfig {
    return new ConnectionConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConnectionConfig {
    return new ConnectionConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConnectionConfig {
    return new ConnectionConfig().fromJsonString(jsonString, options);
  }

  static equals(a: ConnectionConfig | PlainMessage<ConnectionConfig> | undefined, b: ConnectionConfig | PlainMessage<ConnectionConfig> | undefined): boolean {
    return proto3.util.equals(ConnectionConfig, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.PostgresConnectionConfig
 */
export class PostgresConnectionConfig extends Message<PostgresConnectionConfig> {
  /**
   * @generated from oneof mgmt.v1alpha1.PostgresConnectionConfig.connection_config
   */
  connectionConfig: {
    /**
     * @generated from field: string url = 1;
     */
    value: string;
    case: "url";
  } | {
    /**
     * @generated from field: mgmt.v1alpha1.PostgresConnection connection = 2;
     */
    value: PostgresConnection;
    case: "connection";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<PostgresConnectionConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.PostgresConnectionConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "connection_config" },
    { no: 2, name: "connection", kind: "message", T: PostgresConnection, oneof: "connection_config" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostgresConnectionConfig {
    return new PostgresConnectionConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostgresConnectionConfig {
    return new PostgresConnectionConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostgresConnectionConfig {
    return new PostgresConnectionConfig().fromJsonString(jsonString, options);
  }

  static equals(a: PostgresConnectionConfig | PlainMessage<PostgresConnectionConfig> | undefined, b: PostgresConnectionConfig | PlainMessage<PostgresConnectionConfig> | undefined): boolean {
    return proto3.util.equals(PostgresConnectionConfig, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.PostgresConnection
 */
export class PostgresConnection extends Message<PostgresConnection> {
  /**
   * @generated from field: string host = 1;
   */
  host = "";

  /**
   * @generated from field: int32 port = 2;
   */
  port = 0;

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string user = 4;
   */
  user = "";

  /**
   * @generated from field: string pass = 5;
   */
  pass = "";

  /**
   * @generated from field: optional string ssl_mode = 6;
   */
  sslMode?: string;

  constructor(data?: PartialMessage<PostgresConnection>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.PostgresConnection";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "pass", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "ssl_mode", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostgresConnection {
    return new PostgresConnection().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostgresConnection {
    return new PostgresConnection().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostgresConnection {
    return new PostgresConnection().fromJsonString(jsonString, options);
  }

  static equals(a: PostgresConnection | PlainMessage<PostgresConnection> | undefined, b: PostgresConnection | PlainMessage<PostgresConnection> | undefined): boolean {
    return proto3.util.equals(PostgresConnection, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.AwsS3ConnectionConfig
 */
export class AwsS3ConnectionConfig extends Message<AwsS3ConnectionConfig> {
  /**
   * @generated from field: string bucket_arn = 1;
   */
  bucketArn = "";

  /**
   * @generated from field: optional string path_prefix = 2;
   */
  pathPrefix?: string;

  constructor(data?: PartialMessage<AwsS3ConnectionConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.AwsS3ConnectionConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bucket_arn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AwsS3ConnectionConfig {
    return new AwsS3ConnectionConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AwsS3ConnectionConfig {
    return new AwsS3ConnectionConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AwsS3ConnectionConfig {
    return new AwsS3ConnectionConfig().fromJsonString(jsonString, options);
  }

  static equals(a: AwsS3ConnectionConfig | PlainMessage<AwsS3ConnectionConfig> | undefined, b: AwsS3ConnectionConfig | PlainMessage<AwsS3ConnectionConfig> | undefined): boolean {
    return proto3.util.equals(AwsS3ConnectionConfig, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.IsConnectionNameAvailableRequest
 */
export class IsConnectionNameAvailableRequest extends Message<IsConnectionNameAvailableRequest> {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId = "";

  /**
   * @generated from field: string connection_name = 2;
   */
  connectionName = "";

  constructor(data?: PartialMessage<IsConnectionNameAvailableRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.IsConnectionNameAvailableRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "connection_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IsConnectionNameAvailableRequest {
    return new IsConnectionNameAvailableRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IsConnectionNameAvailableRequest {
    return new IsConnectionNameAvailableRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IsConnectionNameAvailableRequest {
    return new IsConnectionNameAvailableRequest().fromJsonString(jsonString, options);
  }

  static equals(a: IsConnectionNameAvailableRequest | PlainMessage<IsConnectionNameAvailableRequest> | undefined, b: IsConnectionNameAvailableRequest | PlainMessage<IsConnectionNameAvailableRequest> | undefined): boolean {
    return proto3.util.equals(IsConnectionNameAvailableRequest, a, b);
  }
}

/**
 * @generated from message mgmt.v1alpha1.IsConnectionNameAvailableResponse
 */
export class IsConnectionNameAvailableResponse extends Message<IsConnectionNameAvailableResponse> {
  /**
   * @generated from field: bool is_available = 1;
   */
  isAvailable = false;

  constructor(data?: PartialMessage<IsConnectionNameAvailableResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "mgmt.v1alpha1.IsConnectionNameAvailableResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "is_available", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IsConnectionNameAvailableResponse {
    return new IsConnectionNameAvailableResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IsConnectionNameAvailableResponse {
    return new IsConnectionNameAvailableResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IsConnectionNameAvailableResponse {
    return new IsConnectionNameAvailableResponse().fromJsonString(jsonString, options);
  }

  static equals(a: IsConnectionNameAvailableResponse | PlainMessage<IsConnectionNameAvailableResponse> | undefined, b: IsConnectionNameAvailableResponse | PlainMessage<IsConnectionNameAvailableResponse> | undefined): boolean {
    return proto3.util.equals(IsConnectionNameAvailableResponse, a, b);
  }
}

