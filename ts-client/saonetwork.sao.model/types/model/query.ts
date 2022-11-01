/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../model/params";
import { Metadata } from "../model/metadata";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Model } from "../model/model";

export const protobufPackage = "saonetwork.sao.model";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetMetadataRequest: object = { dataId: "" };

export const QueryGetMetadataRequest = {
  encode(
    message: QueryGetMetadataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.dataId !== "") {
      writer.uint32(10).string(message.dataId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMetadataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMetadataRequest,
    } as QueryGetMetadataRequest;
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
    const message = {
      ...baseQueryGetMetadataRequest,
    } as QueryGetMetadataRequest;
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = String(object.dataId);
    } else {
      message.dataId = "";
    }
    return message;
  },

  toJSON(message: QueryGetMetadataRequest): unknown {
    const obj: any = {};
    message.dataId !== undefined && (obj.dataId = message.dataId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMetadataRequest>
  ): QueryGetMetadataRequest {
    const message = {
      ...baseQueryGetMetadataRequest,
    } as QueryGetMetadataRequest;
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = object.dataId;
    } else {
      message.dataId = "";
    }
    return message;
  },
};

const baseQueryGetMetadataResponse: object = {};

export const QueryGetMetadataResponse = {
  encode(
    message: QueryGetMetadataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMetadataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMetadataResponse,
    } as QueryGetMetadataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMetadataResponse {
    const message = {
      ...baseQueryGetMetadataResponse,
    } as QueryGetMetadataResponse;
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMetadataResponse): unknown {
    const obj: any = {};
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMetadataResponse>
  ): QueryGetMetadataResponse {
    const message = {
      ...baseQueryGetMetadataResponse,
    } as QueryGetMetadataResponse;
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
    }
    return message;
  },
};

const baseQueryAllMetadataRequest: object = {};

export const QueryAllMetadataRequest = {
  encode(
    message: QueryAllMetadataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMetadataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMetadataRequest,
    } as QueryAllMetadataRequest;
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
    const message = {
      ...baseQueryAllMetadataRequest,
    } as QueryAllMetadataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMetadataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMetadataRequest>
  ): QueryAllMetadataRequest {
    const message = {
      ...baseQueryAllMetadataRequest,
    } as QueryAllMetadataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMetadataResponse: object = {};

export const QueryAllMetadataResponse = {
  encode(
    message: QueryAllMetadataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.metadata) {
      Metadata.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllMetadataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMetadataResponse,
    } as QueryAllMetadataResponse;
    message.metadata = [];
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
    const message = {
      ...baseQueryAllMetadataResponse,
    } as QueryAllMetadataResponse;
    message.metadata = [];
    if (object.metadata !== undefined && object.metadata !== null) {
      for (const e of object.metadata) {
        message.metadata.push(Metadata.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMetadataResponse): unknown {
    const obj: any = {};
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) =>
        e ? Metadata.toJSON(e) : undefined
      );
    } else {
      obj.metadata = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMetadataResponse>
  ): QueryAllMetadataResponse {
    const message = {
      ...baseQueryAllMetadataResponse,
    } as QueryAllMetadataResponse;
    message.metadata = [];
    if (object.metadata !== undefined && object.metadata !== null) {
      for (const e of object.metadata) {
        message.metadata.push(Metadata.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetModelRequest: object = { key: "" };

export const QueryGetModelRequest = {
  encode(
    message: QueryGetModelRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest;
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
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    return message;
  },

  toJSON(message: QueryGetModelRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetModelRequest>): QueryGetModelRequest {
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    return message;
  },
};

const baseQueryGetModelResponse: object = {};

export const QueryGetModelResponse = {
  encode(
    message: QueryGetModelResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.model !== undefined) {
      Model.encode(message.model, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse;
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
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse;
    if (object.model !== undefined && object.model !== null) {
      message.model = Model.fromJSON(object.model);
    } else {
      message.model = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetModelResponse): unknown {
    const obj: any = {};
    message.model !== undefined &&
      (obj.model = message.model ? Model.toJSON(message.model) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetModelResponse>
  ): QueryGetModelResponse {
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse;
    if (object.model !== undefined && object.model !== null) {
      message.model = Model.fromPartial(object.model);
    } else {
      message.model = undefined;
    }
    return message;
  },
};

const baseQueryAllModelRequest: object = {};

export const QueryAllModelRequest = {
  encode(
    message: QueryAllModelRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest;
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
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllModelRequest>): QueryAllModelRequest {
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllModelResponse: object = {};

export const QueryAllModelResponse = {
  encode(
    message: QueryAllModelResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.model) {
      Model.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse;
    message.model = [];
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
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse;
    message.model = [];
    if (object.model !== undefined && object.model !== null) {
      for (const e of object.model) {
        message.model.push(Model.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllModelResponse): unknown {
    const obj: any = {};
    if (message.model) {
      obj.model = message.model.map((e) => (e ? Model.toJSON(e) : undefined));
    } else {
      obj.model = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllModelResponse>
  ): QueryAllModelResponse {
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse;
    message.model = [];
    if (object.model !== undefined && object.model !== null) {
      for (const e of object.model) {
        message.model.push(Model.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
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
  MetadataAll(
    request: QueryAllMetadataRequest
  ): Promise<QueryAllMetadataResponse>;
  /** Queries a Model by index. */
  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
  /** Queries a list of Model items. */
  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.model.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Metadata(
    request: QueryGetMetadataRequest
  ): Promise<QueryGetMetadataResponse> {
    const data = QueryGetMetadataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.model.Query",
      "Metadata",
      data
    );
    return promise.then((data) =>
      QueryGetMetadataResponse.decode(new Reader(data))
    );
  }

  MetadataAll(
    request: QueryAllMetadataRequest
  ): Promise<QueryAllMetadataResponse> {
    const data = QueryAllMetadataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.model.Query",
      "MetadataAll",
      data
    );
    return promise.then((data) =>
      QueryAllMetadataResponse.decode(new Reader(data))
    );
  }

  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse> {
    const data = QueryGetModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.model.Query",
      "Model",
      data
    );
    return promise.then((data) =>
      QueryGetModelResponse.decode(new Reader(data))
    );
  }

  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse> {
    const data = QueryAllModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.model.Query",
      "ModelAll",
      data
    );
    return promise.then((data) =>
      QueryAllModelResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
