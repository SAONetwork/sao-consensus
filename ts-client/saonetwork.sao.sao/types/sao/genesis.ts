/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../sao/params";
import { Order } from "../sao/order";
import { Shard } from "../sao/shard";

export const protobufPackage = "saonetwork.sao.sao";

/** GenesisState defines the sao module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  orderList: Order[];
  orderCount: number;
  shardList: Shard[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  shardCount: number;
}

const baseGenesisState: object = { orderCount: 0, shardCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.shardList = [];
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
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.shardList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromJSON(e));
      }
    }
    if (object.orderCount !== undefined && object.orderCount !== null) {
      message.orderCount = Number(object.orderCount);
    } else {
      message.orderCount = 0;
    }
    if (object.shardList !== undefined && object.shardList !== null) {
      for (const e of object.shardList) {
        message.shardList.push(Shard.fromJSON(e));
      }
    }
    if (object.shardCount !== undefined && object.shardCount !== null) {
      message.shardCount = Number(object.shardCount);
    } else {
      message.shardCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.orderList) {
      obj.orderList = message.orderList.map((e) =>
        e ? Order.toJSON(e) : undefined
      );
    } else {
      obj.orderList = [];
    }
    message.orderCount !== undefined && (obj.orderCount = message.orderCount);
    if (message.shardList) {
      obj.shardList = message.shardList.map((e) =>
        e ? Shard.toJSON(e) : undefined
      );
    } else {
      obj.shardList = [];
    }
    message.shardCount !== undefined && (obj.shardCount = message.shardCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.shardList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromPartial(e));
      }
    }
    if (object.orderCount !== undefined && object.orderCount !== null) {
      message.orderCount = object.orderCount;
    } else {
      message.orderCount = 0;
    }
    if (object.shardList !== undefined && object.shardList !== null) {
      for (const e of object.shardList) {
        message.shardList.push(Shard.fromPartial(e));
      }
    }
    if (object.shardCount !== undefined && object.shardCount !== null) {
      message.shardCount = object.shardCount;
    } else {
      message.shardCount = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
