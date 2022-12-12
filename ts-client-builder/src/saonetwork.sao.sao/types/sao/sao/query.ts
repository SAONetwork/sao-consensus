/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { JwsSignature } from "./jws_signature";
import { Metadata } from "./metadata";
import { Params } from "./params";
import { QueryProposal } from "./query_proposal";
import { ShardMeta } from "./shard_meta";

export const protobufPackage = "saonetwork.sao.sao";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryMetadataRequest {
  proposal: QueryProposal | undefined;
  jwsSignature: JwsSignature | undefined;
}

export interface QueryMetadataResponse {
  metadata: Metadata | undefined;
  shards: { [key: string]: ShardMeta };
}

export interface QueryMetadataResponse_ShardsEntry {
  key: string;
  value: ShardMeta | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryMetadataRequest(): QueryMetadataRequest {
  return { proposal: undefined, jwsSignature: undefined };
}

export const QueryMetadataRequest = {
  encode(message: QueryMetadataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proposal !== undefined) {
      QueryProposal.encode(message.proposal, writer.uint32(10).fork()).ldelim();
    }
    if (message.jwsSignature !== undefined) {
      JwsSignature.encode(message.jwsSignature, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMetadataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMetadataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposal = QueryProposal.decode(reader, reader.uint32());
          break;
        case 2:
          message.jwsSignature = JwsSignature.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMetadataRequest {
    return {
      proposal: isSet(object.proposal) ? QueryProposal.fromJSON(object.proposal) : undefined,
      jwsSignature: isSet(object.jwsSignature) ? JwsSignature.fromJSON(object.jwsSignature) : undefined,
    };
  },

  toJSON(message: QueryMetadataRequest): unknown {
    const obj: any = {};
    message.proposal !== undefined
      && (obj.proposal = message.proposal ? QueryProposal.toJSON(message.proposal) : undefined);
    message.jwsSignature !== undefined
      && (obj.jwsSignature = message.jwsSignature ? JwsSignature.toJSON(message.jwsSignature) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMetadataRequest>, I>>(object: I): QueryMetadataRequest {
    const message = createBaseQueryMetadataRequest();
    message.proposal = (object.proposal !== undefined && object.proposal !== null)
      ? QueryProposal.fromPartial(object.proposal)
      : undefined;
    message.jwsSignature = (object.jwsSignature !== undefined && object.jwsSignature !== null)
      ? JwsSignature.fromPartial(object.jwsSignature)
      : undefined;
    return message;
  },
};

function createBaseQueryMetadataResponse(): QueryMetadataResponse {
  return { metadata: undefined, shards: {} };
}

export const QueryMetadataResponse = {
  encode(message: QueryMetadataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(10).fork()).ldelim();
    }
    Object.entries(message.shards).forEach(([key, value]) => {
      QueryMetadataResponse_ShardsEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMetadataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMetadataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 2:
          const entry2 = QueryMetadataResponse_ShardsEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.shards[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMetadataResponse {
    return {
      metadata: isSet(object.metadata) ? Metadata.fromJSON(object.metadata) : undefined,
      shards: isObject(object.shards)
        ? Object.entries(object.shards).reduce<{ [key: string]: ShardMeta }>((acc, [key, value]) => {
          acc[key] = ShardMeta.fromJSON(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: QueryMetadataResponse): unknown {
    const obj: any = {};
    message.metadata !== undefined && (obj.metadata = message.metadata ? Metadata.toJSON(message.metadata) : undefined);
    obj.shards = {};
    if (message.shards) {
      Object.entries(message.shards).forEach(([k, v]) => {
        obj.shards[k] = ShardMeta.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMetadataResponse>, I>>(object: I): QueryMetadataResponse {
    const message = createBaseQueryMetadataResponse();
    message.metadata = (object.metadata !== undefined && object.metadata !== null)
      ? Metadata.fromPartial(object.metadata)
      : undefined;
    message.shards = Object.entries(object.shards ?? {}).reduce<{ [key: string]: ShardMeta }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ShardMeta.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseQueryMetadataResponse_ShardsEntry(): QueryMetadataResponse_ShardsEntry {
  return { key: "", value: undefined };
}

export const QueryMetadataResponse_ShardsEntry = {
  encode(message: QueryMetadataResponse_ShardsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ShardMeta.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMetadataResponse_ShardsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMetadataResponse_ShardsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ShardMeta.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMetadataResponse_ShardsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? ShardMeta.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: QueryMetadataResponse_ShardsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? ShardMeta.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMetadataResponse_ShardsEntry>, I>>(
    object: I,
  ): QueryMetadataResponse_ShardsEntry {
    const message = createBaseQueryMetadataResponse_ShardsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? ShardMeta.fromPartial(object.value)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Metadata items. */
  Metadata(request: QueryMetadataRequest): Promise<QueryMetadataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Metadata = this.Metadata.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Metadata(request: QueryMetadataRequest): Promise<QueryMetadataResponse> {
    const data = QueryMetadataRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Query", "Metadata", data);
    return promise.then((data) => QueryMetadataResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
