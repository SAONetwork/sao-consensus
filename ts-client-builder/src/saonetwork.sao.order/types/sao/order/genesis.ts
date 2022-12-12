/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Order } from "./order";
import { Params } from "./params";
import { Shard } from "./shard";

export const protobufPackage = "saonetwork.sao.order";

/** GenesisState defines the order module's genesis state. */
export interface GenesisState {
  params:
    | Params
    | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  orderList: Order[];
  orderCount: number;
  shardList: Shard[];
  shardCount: number;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, orderList: [], orderCount: 0, shardList: [], shardCount: 0 };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.orderList) {
      Order.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.orderCount !== 0) {
      writer.uint32(24).uint64(message.orderCount);
    }
    for (const v of message.shardList) {
      Shard.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.shardCount !== 0) {
      writer.uint32(40).uint64(message.shardCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.orderList.push(Order.decode(reader, reader.uint32()));
          break;
        case 3:
          message.orderCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.shardList.push(Shard.decode(reader, reader.uint32()));
          break;
        case 5:
          message.shardCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      orderList: Array.isArray(object?.orderList) ? object.orderList.map((e: any) => Order.fromJSON(e)) : [],
      orderCount: isSet(object.orderCount) ? Number(object.orderCount) : 0,
      shardList: Array.isArray(object?.shardList) ? object.shardList.map((e: any) => Shard.fromJSON(e)) : [],
      shardCount: isSet(object.shardCount) ? Number(object.shardCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.orderList) {
      obj.orderList = message.orderList.map((e) => e ? Order.toJSON(e) : undefined);
    } else {
      obj.orderList = [];
    }
    message.orderCount !== undefined && (obj.orderCount = Math.round(message.orderCount));
    if (message.shardList) {
      obj.shardList = message.shardList.map((e) => e ? Shard.toJSON(e) : undefined);
    } else {
      obj.shardList = [];
    }
    message.shardCount !== undefined && (obj.shardCount = Math.round(message.shardCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.orderList = object.orderList?.map((e) => Order.fromPartial(e)) || [];
    message.orderCount = object.orderCount ?? 0;
    message.shardList = object.shardList?.map((e) => Shard.fromPartial(e)) || [];
    message.shardCount = object.shardCount ?? 0;
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
