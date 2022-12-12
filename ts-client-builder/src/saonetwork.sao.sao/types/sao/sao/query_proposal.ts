/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface QueryProposal {
  owner: string;
  keyword: string;
  groupId: string;
  /** 0,1 - query by dataId, 2 - query by alias */
  type: number;
  lastValidHeight: number;
  gateway: string;
  commitId: string;
  version: string;
}

function createBaseQueryProposal(): QueryProposal {
  return { owner: "", keyword: "", groupId: "", type: 0, lastValidHeight: 0, gateway: "", commitId: "", version: "" };
}

export const QueryProposal = {
  encode(message: QueryProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.keyword !== "") {
      writer.uint32(18).string(message.keyword);
    }
    if (message.groupId !== "") {
      writer.uint32(26).string(message.groupId);
    }
    if (message.type !== 0) {
      writer.uint32(32).uint32(message.type);
    }
    if (message.lastValidHeight !== 0) {
      writer.uint32(40).uint64(message.lastValidHeight);
    }
    if (message.gateway !== "") {
      writer.uint32(50).string(message.gateway);
    }
    if (message.commitId !== "") {
      writer.uint32(58).string(message.commitId);
    }
    if (message.version !== "") {
      writer.uint32(66).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.keyword = reader.string();
          break;
        case 3:
          message.groupId = reader.string();
          break;
        case 4:
          message.type = reader.uint32();
          break;
        case 5:
          message.lastValidHeight = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.gateway = reader.string();
          break;
        case 7:
          message.commitId = reader.string();
          break;
        case 8:
          message.version = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryProposal {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      keyword: isSet(object.keyword) ? String(object.keyword) : "",
      groupId: isSet(object.groupId) ? String(object.groupId) : "",
      type: isSet(object.type) ? Number(object.type) : 0,
      lastValidHeight: isSet(object.lastValidHeight) ? Number(object.lastValidHeight) : 0,
      gateway: isSet(object.gateway) ? String(object.gateway) : "",
      commitId: isSet(object.commitId) ? String(object.commitId) : "",
      version: isSet(object.version) ? String(object.version) : "",
    };
  },

  toJSON(message: QueryProposal): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.keyword !== undefined && (obj.keyword = message.keyword);
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.type !== undefined && (obj.type = Math.round(message.type));
    message.lastValidHeight !== undefined && (obj.lastValidHeight = Math.round(message.lastValidHeight));
    message.gateway !== undefined && (obj.gateway = message.gateway);
    message.commitId !== undefined && (obj.commitId = message.commitId);
    message.version !== undefined && (obj.version = message.version);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryProposal>, I>>(object: I): QueryProposal {
    const message = createBaseQueryProposal();
    message.owner = object.owner ?? "";
    message.keyword = object.keyword ?? "";
    message.groupId = object.groupId ?? "";
    message.type = object.type ?? 0;
    message.lastValidHeight = object.lastValidHeight ?? 0;
    message.gateway = object.gateway ?? "";
    message.commitId = object.commitId ?? "";
    message.version = object.version ?? "";
    return message;
  },
};

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
