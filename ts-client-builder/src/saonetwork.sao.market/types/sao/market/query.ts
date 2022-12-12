/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Worker } from "./worker";

export const protobufPackage = "saonetwork.sao.market";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetWorkerRequest {
  workername: string;
}

export interface QueryGetWorkerResponse {
  worker: Worker | undefined;
}

export interface QueryAllWorkerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllWorkerResponse {
  worker: Worker[];
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

function createBaseQueryGetWorkerRequest(): QueryGetWorkerRequest {
  return { workername: "" };
}

export const QueryGetWorkerRequest = {
  encode(message: QueryGetWorkerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.workername !== "") {
      writer.uint32(10).string(message.workername);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetWorkerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetWorkerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.workername = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWorkerRequest {
    return { workername: isSet(object.workername) ? String(object.workername) : "" };
  },

  toJSON(message: QueryGetWorkerRequest): unknown {
    const obj: any = {};
    message.workername !== undefined && (obj.workername = message.workername);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetWorkerRequest>, I>>(object: I): QueryGetWorkerRequest {
    const message = createBaseQueryGetWorkerRequest();
    message.workername = object.workername ?? "";
    return message;
  },
};

function createBaseQueryGetWorkerResponse(): QueryGetWorkerResponse {
  return { worker: undefined };
}

export const QueryGetWorkerResponse = {
  encode(message: QueryGetWorkerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.worker !== undefined) {
      Worker.encode(message.worker, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetWorkerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetWorkerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.worker = Worker.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWorkerResponse {
    return { worker: isSet(object.worker) ? Worker.fromJSON(object.worker) : undefined };
  },

  toJSON(message: QueryGetWorkerResponse): unknown {
    const obj: any = {};
    message.worker !== undefined && (obj.worker = message.worker ? Worker.toJSON(message.worker) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetWorkerResponse>, I>>(object: I): QueryGetWorkerResponse {
    const message = createBaseQueryGetWorkerResponse();
    message.worker = (object.worker !== undefined && object.worker !== null)
      ? Worker.fromPartial(object.worker)
      : undefined;
    return message;
  },
};

function createBaseQueryAllWorkerRequest(): QueryAllWorkerRequest {
  return { pagination: undefined };
}

export const QueryAllWorkerRequest = {
  encode(message: QueryAllWorkerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllWorkerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllWorkerRequest();
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

  fromJSON(object: any): QueryAllWorkerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllWorkerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllWorkerRequest>, I>>(object: I): QueryAllWorkerRequest {
    const message = createBaseQueryAllWorkerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllWorkerResponse(): QueryAllWorkerResponse {
  return { worker: [], pagination: undefined };
}

export const QueryAllWorkerResponse = {
  encode(message: QueryAllWorkerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.worker) {
      Worker.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllWorkerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllWorkerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.worker.push(Worker.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllWorkerResponse {
    return {
      worker: Array.isArray(object?.worker) ? object.worker.map((e: any) => Worker.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllWorkerResponse): unknown {
    const obj: any = {};
    if (message.worker) {
      obj.worker = message.worker.map((e) => e ? Worker.toJSON(e) : undefined);
    } else {
      obj.worker = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllWorkerResponse>, I>>(object: I): QueryAllWorkerResponse {
    const message = createBaseQueryAllWorkerResponse();
    message.worker = object.worker?.map((e) => Worker.fromPartial(e)) || [];
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
  /** Queries a Worker by index. */
  Worker(request: QueryGetWorkerRequest): Promise<QueryGetWorkerResponse>;
  /** Queries a list of Worker items. */
  WorkerAll(request: QueryAllWorkerRequest): Promise<QueryAllWorkerResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Worker = this.Worker.bind(this);
    this.WorkerAll = this.WorkerAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.market.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Worker(request: QueryGetWorkerRequest): Promise<QueryGetWorkerResponse> {
    const data = QueryGetWorkerRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.market.Query", "Worker", data);
    return promise.then((data) => QueryGetWorkerResponse.decode(new _m0.Reader(data)));
  }

  WorkerAll(request: QueryAllWorkerRequest): Promise<QueryAllWorkerResponse> {
    const data = QueryAllWorkerRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.market.Query", "WorkerAll", data);
    return promise.then((data) => QueryAllWorkerResponse.decode(new _m0.Reader(data)));
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
