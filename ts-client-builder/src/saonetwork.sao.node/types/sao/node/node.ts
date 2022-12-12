/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

export interface Node {
  creator: string;
  peer: string;
  reputation: number;
  status: number;
  lastAliveHeigh: number;
}

function createBaseNode(): Node {
  return { creator: "", peer: "", reputation: 0, status: 0, lastAliveHeigh: 0 };
}

export const Node = {
  encode(message: Node, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    if (message.reputation !== 0) {
      writer.uint32(29).float(message.reputation);
    }
    if (message.status !== 0) {
      writer.uint32(32).uint32(message.status);
    }
    if (message.lastAliveHeigh !== 0) {
      writer.uint32(40).int64(message.lastAliveHeigh);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Node {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.peer = reader.string();
          break;
        case 3:
          message.reputation = reader.float();
          break;
        case 4:
          message.status = reader.uint32();
          break;
        case 5:
          message.lastAliveHeigh = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Node {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      peer: isSet(object.peer) ? String(object.peer) : "",
      reputation: isSet(object.reputation) ? Number(object.reputation) : 0,
      status: isSet(object.status) ? Number(object.status) : 0,
      lastAliveHeigh: isSet(object.lastAliveHeigh) ? Number(object.lastAliveHeigh) : 0,
    };
  },

  toJSON(message: Node): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.peer !== undefined && (obj.peer = message.peer);
    message.reputation !== undefined && (obj.reputation = message.reputation);
    message.status !== undefined && (obj.status = Math.round(message.status));
    message.lastAliveHeigh !== undefined && (obj.lastAliveHeigh = Math.round(message.lastAliveHeigh));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Node>, I>>(object: I): Node {
    const message = createBaseNode();
    message.creator = object.creator ?? "";
    message.peer = object.peer ?? "";
    message.reputation = object.reputation ?? 0;
    message.status = object.status ?? 0;
    message.lastAliveHeigh = object.lastAliveHeigh ?? 0;
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
