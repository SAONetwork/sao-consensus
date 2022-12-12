/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.order";

export interface Shard {
  id: number;
  orderId: number;
  status: number;
  size: number;
  cid: string;
  pledge: Coin | undefined;
}

function createBaseShard(): Shard {
  return { id: 0, orderId: 0, status: 0, size: 0, cid: "", pledge: undefined };
}

export const Shard = {
  encode(message: Shard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    if (message.status !== 0) {
      writer.uint32(24).int32(message.status);
    }
    if (message.size !== 0) {
      writer.uint32(32).int32(message.size);
    }
    if (message.cid !== "") {
      writer.uint32(42).string(message.cid);
    }
    if (message.pledge !== undefined) {
      Coin.encode(message.pledge, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Shard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShard();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.status = reader.int32();
          break;
        case 4:
          message.size = reader.int32();
          break;
        case 5:
          message.cid = reader.string();
          break;
        case 6:
          message.pledge = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Shard {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
      status: isSet(object.status) ? Number(object.status) : 0,
      size: isSet(object.size) ? Number(object.size) : 0,
      cid: isSet(object.cid) ? String(object.cid) : "",
      pledge: isSet(object.pledge) ? Coin.fromJSON(object.pledge) : undefined,
    };
  },

  toJSON(message: Shard): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    message.status !== undefined && (obj.status = Math.round(message.status));
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.cid !== undefined && (obj.cid = message.cid);
    message.pledge !== undefined && (obj.pledge = message.pledge ? Coin.toJSON(message.pledge) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Shard>, I>>(object: I): Shard {
    const message = createBaseShard();
    message.id = object.id ?? 0;
    message.orderId = object.orderId ?? 0;
    message.status = object.status ?? 0;
    message.size = object.size ?? 0;
    message.cid = object.cid ?? "";
    message.pledge = (object.pledge !== undefined && object.pledge !== null)
      ? Coin.fromPartial(object.pledge)
      : undefined;
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
