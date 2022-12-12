/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface ShardMeta {
  shardId: number;
  peer: string;
  cid: string;
  provider: string;
}

function createBaseShardMeta(): ShardMeta {
  return { shardId: 0, peer: "", cid: "", provider: "" };
}

export const ShardMeta = {
  encode(message: ShardMeta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.shardId !== 0) {
      writer.uint32(8).uint64(message.shardId);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    if (message.provider !== "") {
      writer.uint32(34).string(message.provider);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ShardMeta {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShardMeta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.shardId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.peer = reader.string();
          break;
        case 3:
          message.cid = reader.string();
          break;
        case 4:
          message.provider = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShardMeta {
    return {
      shardId: isSet(object.shardId) ? Number(object.shardId) : 0,
      peer: isSet(object.peer) ? String(object.peer) : "",
      cid: isSet(object.cid) ? String(object.cid) : "",
      provider: isSet(object.provider) ? String(object.provider) : "",
    };
  },

  toJSON(message: ShardMeta): unknown {
    const obj: any = {};
    message.shardId !== undefined && (obj.shardId = Math.round(message.shardId));
    message.peer !== undefined && (obj.peer = message.peer);
    message.cid !== undefined && (obj.cid = message.cid);
    message.provider !== undefined && (obj.provider = message.provider);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ShardMeta>, I>>(object: I): ShardMeta {
    const message = createBaseShardMeta();
    message.shardId = object.shardId ?? 0;
    message.peer = object.peer ?? "";
    message.cid = object.cid ?? "";
    message.provider = object.provider ?? "";
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
