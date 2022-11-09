/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../node/params";
import { Node } from "../node/node";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Pool } from "../node/pool";
import { Pledge } from "../node/pledge";

export const protobufPackage = "saonetwork.sao.node";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNodeRequest {
  creator: string;
}

export interface QueryGetNodeResponse {
  node: Node | undefined;
}

export interface QueryAllNodeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNodeResponse {
  node: Node[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPoolRequest {}

export interface QueryGetPoolResponse {
  Pool: Pool | undefined;
}

export interface QueryGetPledgeRequest {
  creator: string;
}

export interface QueryGetPledgeResponse {
  pledge: Pledge | undefined;
}

export interface QueryAllPledgeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPledgeResponse {
  pledge: Pledge[];
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

const baseQueryGetNodeRequest: object = { creator: "" };

export const QueryGetNodeRequest = {
  encode(
    message: QueryGetNodeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNodeRequest {
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: QueryGetNodeRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNodeRequest>): QueryGetNodeRequest {
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseQueryGetNodeResponse: object = {};

export const QueryGetNodeResponse = {
  encode(
    message: QueryGetNodeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.node !== undefined) {
      Node.encode(message.node, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.node = Node.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNodeResponse {
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse;
    if (object.node !== undefined && object.node !== null) {
      message.node = Node.fromJSON(object.node);
    } else {
      message.node = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNodeResponse): unknown {
    const obj: any = {};
    message.node !== undefined &&
      (obj.node = message.node ? Node.toJSON(message.node) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNodeResponse>): QueryGetNodeResponse {
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse;
    if (object.node !== undefined && object.node !== null) {
      message.node = Node.fromPartial(object.node);
    } else {
      message.node = undefined;
    }
    return message;
  },
};

const baseQueryAllNodeRequest: object = {};

export const QueryAllNodeRequest = {
  encode(
    message: QueryAllNodeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest;
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

  fromJSON(object: any): QueryAllNodeRequest {
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNodeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllNodeRequest>): QueryAllNodeRequest {
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNodeResponse: object = {};

export const QueryAllNodeResponse = {
  encode(
    message: QueryAllNodeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.node) {
      Node.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse;
    message.node = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.node.push(Node.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllNodeResponse {
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse;
    message.node = [];
    if (object.node !== undefined && object.node !== null) {
      for (const e of object.node) {
        message.node.push(Node.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNodeResponse): unknown {
    const obj: any = {};
    if (message.node) {
      obj.node = message.node.map((e) => (e ? Node.toJSON(e) : undefined));
    } else {
      obj.node = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllNodeResponse>): QueryAllNodeResponse {
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse;
    message.node = [];
    if (object.node !== undefined && object.node !== null) {
      for (const e of object.node) {
        message.node.push(Node.fromPartial(e));
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

const baseQueryGetPoolRequest: object = {};

export const QueryGetPoolRequest = {
  encode(_: QueryGetPoolRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
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

  fromJSON(_: any): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    return message;
  },

  toJSON(_: QueryGetPoolRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetPoolRequest>): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    return message;
  },
};

const baseQueryGetPoolResponse: object = {};

export const QueryGetPoolResponse = {
  encode(
    message: QueryGetPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Pool !== undefined) {
      Pool.encode(message.Pool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Pool = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.Pool !== undefined && object.Pool !== null) {
      message.Pool = Pool.fromJSON(object.Pool);
    } else {
      message.Pool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPoolResponse): unknown {
    const obj: any = {};
    message.Pool !== undefined &&
      (obj.Pool = message.Pool ? Pool.toJSON(message.Pool) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPoolResponse>): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.Pool !== undefined && object.Pool !== null) {
      message.Pool = Pool.fromPartial(object.Pool);
    } else {
      message.Pool = undefined;
    }
    return message;
  },
};

const baseQueryGetPledgeRequest: object = { creator: "" };

export const QueryGetPledgeRequest = {
  encode(
    message: QueryGetPledgeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPledgeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPledgeRequest } as QueryGetPledgeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPledgeRequest {
    const message = { ...baseQueryGetPledgeRequest } as QueryGetPledgeRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: QueryGetPledgeRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPledgeRequest>
  ): QueryGetPledgeRequest {
    const message = { ...baseQueryGetPledgeRequest } as QueryGetPledgeRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseQueryGetPledgeResponse: object = {};

export const QueryGetPledgeResponse = {
  encode(
    message: QueryGetPledgeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pledge !== undefined) {
      Pledge.encode(message.pledge, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPledgeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPledgeResponse } as QueryGetPledgeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pledge = Pledge.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPledgeResponse {
    const message = { ...baseQueryGetPledgeResponse } as QueryGetPledgeResponse;
    if (object.pledge !== undefined && object.pledge !== null) {
      message.pledge = Pledge.fromJSON(object.pledge);
    } else {
      message.pledge = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPledgeResponse): unknown {
    const obj: any = {};
    message.pledge !== undefined &&
      (obj.pledge = message.pledge ? Pledge.toJSON(message.pledge) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPledgeResponse>
  ): QueryGetPledgeResponse {
    const message = { ...baseQueryGetPledgeResponse } as QueryGetPledgeResponse;
    if (object.pledge !== undefined && object.pledge !== null) {
      message.pledge = Pledge.fromPartial(object.pledge);
    } else {
      message.pledge = undefined;
    }
    return message;
  },
};

const baseQueryAllPledgeRequest: object = {};

export const QueryAllPledgeRequest = {
  encode(
    message: QueryAllPledgeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPledgeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPledgeRequest } as QueryAllPledgeRequest;
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

  fromJSON(object: any): QueryAllPledgeRequest {
    const message = { ...baseQueryAllPledgeRequest } as QueryAllPledgeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPledgeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPledgeRequest>
  ): QueryAllPledgeRequest {
    const message = { ...baseQueryAllPledgeRequest } as QueryAllPledgeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPledgeResponse: object = {};

export const QueryAllPledgeResponse = {
  encode(
    message: QueryAllPledgeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pledge) {
      Pledge.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPledgeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPledgeResponse } as QueryAllPledgeResponse;
    message.pledge = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pledge.push(Pledge.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllPledgeResponse {
    const message = { ...baseQueryAllPledgeResponse } as QueryAllPledgeResponse;
    message.pledge = [];
    if (object.pledge !== undefined && object.pledge !== null) {
      for (const e of object.pledge) {
        message.pledge.push(Pledge.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPledgeResponse): unknown {
    const obj: any = {};
    if (message.pledge) {
      obj.pledge = message.pledge.map((e) =>
        e ? Pledge.toJSON(e) : undefined
      );
    } else {
      obj.pledge = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPledgeResponse>
  ): QueryAllPledgeResponse {
    const message = { ...baseQueryAllPledgeResponse } as QueryAllPledgeResponse;
    message.pledge = [];
    if (object.pledge !== undefined && object.pledge !== null) {
      for (const e of object.pledge) {
        message.pledge.push(Pledge.fromPartial(e));
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
  /** Queries a Pool by index. */
  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse>;
  /** Queries a Node by index. */
  Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse>;
  /** Queries a list of Node items. */
  NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse>;
  /** Queries a Pledge by index. */
  Pledge(request: QueryGetPledgeRequest): Promise<QueryGetPledgeResponse>;
  /** Queries a list of Pledge items. */
  PledgeAll(request: QueryAllPledgeRequest): Promise<QueryAllPledgeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.node.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse> {
    const data = QueryGetPoolRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Query", "Pool", data);
    return promise.then((data) =>
      QueryGetPoolResponse.decode(new Reader(data))
    );
  }

  Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse> {
    const data = QueryGetNodeRequest.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Query", "Node", data);
    return promise.then((data) =>
      QueryGetNodeResponse.decode(new Reader(data))
    );
  }

  NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse> {
    const data = QueryAllNodeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.node.Query",
      "NodeAll",
      data
    );
    return promise.then((data) =>
      QueryAllNodeResponse.decode(new Reader(data))
    );
  }

  Pledge(request: QueryGetPledgeRequest): Promise<QueryGetPledgeResponse> {
    const data = QueryGetPledgeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.node.Query",
      "Pledge",
      data
    );
    return promise.then((data) =>
      QueryGetPledgeResponse.decode(new Reader(data))
    );
  }

  PledgeAll(request: QueryAllPledgeRequest): Promise<QueryAllPledgeResponse> {
    const data = QueryAllPledgeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.node.Query",
      "PledgeAll",
      data
    );
    return promise.then((data) =>
      QueryAllPledgeResponse.decode(new Reader(data))
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
