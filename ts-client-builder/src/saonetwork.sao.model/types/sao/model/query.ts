/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { ExpiredData } from "./expired_data";
import { Metadata } from "./metadata";
import { Model } from "./model";
import { Params } from "./params";
import { ShardMeta } from "./shard_meta";

export const protobufPackage = "saonetwork.sao.model";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetMetadataRequest {
  dataId: string;
}

export interface QueryGetMetadataResponse {
  metadata: Metadata | undefined;
  orderId: number;
  shards: { [key: string]: ShardMeta };
}

export interface QueryGetMetadataResponse_ShardsEntry {
  key: string;
  value: ShardMeta | undefined;
}

export interface QueryAllMetadataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMetadataResponse {
  metadata: Metadata[];
  pagination: PageResponse | undefined;
}

export interface QueryGetModelRequest {
  key: string;
}

export interface QueryGetModelResponse {
  model: Model | undefined;
}

export interface QueryAllModelRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllModelResponse {
  model: Model[];
  pagination: PageResponse | undefined;
}

export interface QueryGetExpiredDataRequest {
  height: number;
}

export interface QueryGetExpiredDataResponse {
  expiredData: ExpiredData | undefined;
}

export interface QueryAllExpiredDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllExpiredDataResponse {
  expiredData: ExpiredData[];
  pagination: PageResponse | undefined;
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

function createBaseQueryGetMetadataRequest(): QueryGetMetadataRequest {
  return { dataId: "" };
}

export const QueryGetMetadataRequest = {
  encode(message: QueryGetMetadataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dataId !== "") {
      writer.uint32(10).string(message.dataId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMetadataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMetadataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dataId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMetadataRequest {
    return { dataId: isSet(object.dataId) ? String(object.dataId) : "" };
  },

  toJSON(message: QueryGetMetadataRequest): unknown {
    const obj: any = {};
    message.dataId !== undefined && (obj.dataId = message.dataId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMetadataRequest>, I>>(object: I): QueryGetMetadataRequest {
    const message = createBaseQueryGetMetadataRequest();
    message.dataId = object.dataId ?? "";
    return message;
  },
};

function createBaseQueryGetMetadataResponse(): QueryGetMetadataResponse {
  return { metadata: undefined, orderId: 0, shards: {} };
}

export const QueryGetMetadataResponse = {
  encode(message: QueryGetMetadataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(10).fork()).ldelim();
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    Object.entries(message.shards).forEach(([key, value]) => {
      QueryGetMetadataResponse_ShardsEntry.encode({ key: key as any, value }, writer.uint32(26).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMetadataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMetadataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          const entry3 = QueryGetMetadataResponse_ShardsEntry.decode(reader, reader.uint32());
          if (entry3.value !== undefined) {
            message.shards[entry3.key] = entry3.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMetadataResponse {
    return {
      metadata: isSet(object.metadata) ? Metadata.fromJSON(object.metadata) : undefined,
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
      shards: isObject(object.shards)
        ? Object.entries(object.shards).reduce<{ [key: string]: ShardMeta }>((acc, [key, value]) => {
          acc[key] = ShardMeta.fromJSON(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: QueryGetMetadataResponse): unknown {
    const obj: any = {};
    message.metadata !== undefined && (obj.metadata = message.metadata ? Metadata.toJSON(message.metadata) : undefined);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    obj.shards = {};
    if (message.shards) {
      Object.entries(message.shards).forEach(([k, v]) => {
        obj.shards[k] = ShardMeta.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMetadataResponse>, I>>(object: I): QueryGetMetadataResponse {
    const message = createBaseQueryGetMetadataResponse();
    message.metadata = (object.metadata !== undefined && object.metadata !== null)
      ? Metadata.fromPartial(object.metadata)
      : undefined;
    message.orderId = object.orderId ?? 0;
    message.shards = Object.entries(object.shards ?? {}).reduce<{ [key: string]: ShardMeta }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ShardMeta.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseQueryGetMetadataResponse_ShardsEntry(): QueryGetMetadataResponse_ShardsEntry {
  return { key: "", value: undefined };
}

export const QueryGetMetadataResponse_ShardsEntry = {
  encode(message: QueryGetMetadataResponse_ShardsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ShardMeta.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMetadataResponse_ShardsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMetadataResponse_ShardsEntry();
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

  fromJSON(object: any): QueryGetMetadataResponse_ShardsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? ShardMeta.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: QueryGetMetadataResponse_ShardsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? ShardMeta.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMetadataResponse_ShardsEntry>, I>>(
    object: I,
  ): QueryGetMetadataResponse_ShardsEntry {
    const message = createBaseQueryGetMetadataResponse_ShardsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? ShardMeta.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMetadataRequest(): QueryAllMetadataRequest {
  return { pagination: undefined };
}

export const QueryAllMetadataRequest = {
  encode(message: QueryAllMetadataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMetadataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMetadataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMetadataRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllMetadataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMetadataRequest>, I>>(object: I): QueryAllMetadataRequest {
    const message = createBaseQueryAllMetadataRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMetadataResponse(): QueryAllMetadataResponse {
  return { metadata: [], pagination: undefined };
}

export const QueryAllMetadataResponse = {
  encode(message: QueryAllMetadataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.metadata) {
      Metadata.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMetadataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMetadataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadata.push(Metadata.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMetadataResponse {
    return {
      metadata: Array.isArray(object?.metadata) ? object.metadata.map((e: any) => Metadata.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllMetadataResponse): unknown {
    const obj: any = {};
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) => e ? Metadata.toJSON(e) : undefined);
    } else {
      obj.metadata = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMetadataResponse>, I>>(object: I): QueryAllMetadataResponse {
    const message = createBaseQueryAllMetadataResponse();
    message.metadata = object.metadata?.map((e) => Metadata.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetModelRequest(): QueryGetModelRequest {
  return { key: "" };
}

export const QueryGetModelRequest = {
  encode(message: QueryGetModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetModelRequest {
    return { key: isSet(object.key) ? String(object.key) : "" };
  },

  toJSON(message: QueryGetModelRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelRequest>, I>>(object: I): QueryGetModelRequest {
    const message = createBaseQueryGetModelRequest();
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseQueryGetModelResponse(): QueryGetModelResponse {
  return { model: undefined };
}

export const QueryGetModelResponse = {
  encode(message: QueryGetModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.model !== undefined) {
      Model.encode(message.model, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.model = Model.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetModelResponse {
    return { model: isSet(object.model) ? Model.fromJSON(object.model) : undefined };
  },

  toJSON(message: QueryGetModelResponse): unknown {
    const obj: any = {};
    message.model !== undefined && (obj.model = message.model ? Model.toJSON(message.model) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelResponse>, I>>(object: I): QueryGetModelResponse {
    const message = createBaseQueryGetModelResponse();
    message.model = (object.model !== undefined && object.model !== null) ? Model.fromPartial(object.model) : undefined;
    return message;
  },
};

function createBaseQueryAllModelRequest(): QueryAllModelRequest {
  return { pagination: undefined };
}

export const QueryAllModelRequest = {
  encode(message: QueryAllModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllModelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllModelRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllModelRequest>, I>>(object: I): QueryAllModelRequest {
    const message = createBaseQueryAllModelRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllModelResponse(): QueryAllModelResponse {
  return { model: [], pagination: undefined };
}

export const QueryAllModelResponse = {
  encode(message: QueryAllModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.model) {
      Model.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.model.push(Model.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllModelResponse {
    return {
      model: Array.isArray(object?.model) ? object.model.map((e: any) => Model.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllModelResponse): unknown {
    const obj: any = {};
    if (message.model) {
      obj.model = message.model.map((e) => e ? Model.toJSON(e) : undefined);
    } else {
      obj.model = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllModelResponse>, I>>(object: I): QueryAllModelResponse {
    const message = createBaseQueryAllModelResponse();
    message.model = object.model?.map((e) => Model.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetExpiredDataRequest(): QueryGetExpiredDataRequest {
  return { height: 0 };
}

export const QueryGetExpiredDataRequest = {
  encode(message: QueryGetExpiredDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.height !== 0) {
      writer.uint32(8).uint64(message.height);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetExpiredDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetExpiredDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExpiredDataRequest {
    return { height: isSet(object.height) ? Number(object.height) : 0 };
  },

  toJSON(message: QueryGetExpiredDataRequest): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = Math.round(message.height));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetExpiredDataRequest>, I>>(object: I): QueryGetExpiredDataRequest {
    const message = createBaseQueryGetExpiredDataRequest();
    message.height = object.height ?? 0;
    return message;
  },
};

function createBaseQueryGetExpiredDataResponse(): QueryGetExpiredDataResponse {
  return { expiredData: undefined };
}

export const QueryGetExpiredDataResponse = {
  encode(message: QueryGetExpiredDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.expiredData !== undefined) {
      ExpiredData.encode(message.expiredData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetExpiredDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetExpiredDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.expiredData = ExpiredData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExpiredDataResponse {
    return { expiredData: isSet(object.expiredData) ? ExpiredData.fromJSON(object.expiredData) : undefined };
  },

  toJSON(message: QueryGetExpiredDataResponse): unknown {
    const obj: any = {};
    message.expiredData !== undefined
      && (obj.expiredData = message.expiredData ? ExpiredData.toJSON(message.expiredData) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetExpiredDataResponse>, I>>(object: I): QueryGetExpiredDataResponse {
    const message = createBaseQueryGetExpiredDataResponse();
    message.expiredData = (object.expiredData !== undefined && object.expiredData !== null)
      ? ExpiredData.fromPartial(object.expiredData)
      : undefined;
    return message;
  },
};

function createBaseQueryAllExpiredDataRequest(): QueryAllExpiredDataRequest {
  return { pagination: undefined };
}

export const QueryAllExpiredDataRequest = {
  encode(message: QueryAllExpiredDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllExpiredDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllExpiredDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExpiredDataRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllExpiredDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllExpiredDataRequest>, I>>(object: I): QueryAllExpiredDataRequest {
    const message = createBaseQueryAllExpiredDataRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllExpiredDataResponse(): QueryAllExpiredDataResponse {
  return { expiredData: [], pagination: undefined };
}

export const QueryAllExpiredDataResponse = {
  encode(message: QueryAllExpiredDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.expiredData) {
      ExpiredData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllExpiredDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllExpiredDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.expiredData.push(ExpiredData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExpiredDataResponse {
    return {
      expiredData: Array.isArray(object?.expiredData)
        ? object.expiredData.map((e: any) => ExpiredData.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllExpiredDataResponse): unknown {
    const obj: any = {};
    if (message.expiredData) {
      obj.expiredData = message.expiredData.map((e) => e ? ExpiredData.toJSON(e) : undefined);
    } else {
      obj.expiredData = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllExpiredDataResponse>, I>>(object: I): QueryAllExpiredDataResponse {
    const message = createBaseQueryAllExpiredDataResponse();
    message.expiredData = object.expiredData?.map((e) => ExpiredData.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Metadata by index. */
  Metadata(request: QueryGetMetadataRequest): Promise<QueryGetMetadataResponse>;
  /** Queries a list of Metadata items. */
  MetadataAll(request: QueryAllMetadataRequest): Promise<QueryAllMetadataResponse>;
  /** Queries a Model by index. */
  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
  /** Queries a list of Model items. */
  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
  /** Queries a ExpiredData by index. */
  ExpiredData(request: QueryGetExpiredDataRequest): Promise<QueryGetExpiredDataResponse>;
  /** Queries a list of ExpiredData items. */
  ExpiredDataAll(request: QueryAllExpiredDataRequest): Promise<QueryAllExpiredDataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Metadata = this.Metadata.bind(this);
    this.MetadataAll = this.MetadataAll.bind(this);
    this.Model = this.Model.bind(this);
    this.ModelAll = this.ModelAll.bind(this);
    this.ExpiredData = this.ExpiredData.bind(this);
    this.ExpiredDataAll = this.ExpiredDataAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Metadata(request: QueryGetMetadataRequest): Promise<QueryGetMetadataResponse> {
    const data = QueryGetMetadataRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "Metadata", data);
    return promise.then((data) => QueryGetMetadataResponse.decode(new _m0.Reader(data)));
  }

  MetadataAll(request: QueryAllMetadataRequest): Promise<QueryAllMetadataResponse> {
    const data = QueryAllMetadataRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "MetadataAll", data);
    return promise.then((data) => QueryAllMetadataResponse.decode(new _m0.Reader(data)));
  }

  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse> {
    const data = QueryGetModelRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "Model", data);
    return promise.then((data) => QueryGetModelResponse.decode(new _m0.Reader(data)));
  }

  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse> {
    const data = QueryAllModelRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "ModelAll", data);
    return promise.then((data) => QueryAllModelResponse.decode(new _m0.Reader(data)));
  }

  ExpiredData(request: QueryGetExpiredDataRequest): Promise<QueryGetExpiredDataResponse> {
    const data = QueryGetExpiredDataRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "ExpiredData", data);
    return promise.then((data) => QueryGetExpiredDataResponse.decode(new _m0.Reader(data)));
  }

  ExpiredDataAll(request: QueryAllExpiredDataRequest): Promise<QueryAllExpiredDataResponse> {
    const data = QueryAllExpiredDataRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.model.Query", "ExpiredDataAll", data);
    return promise.then((data) => QueryAllExpiredDataResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
