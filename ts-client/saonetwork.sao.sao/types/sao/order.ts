/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Shard } from "../sao/shard";

export const protobufPackage = "saonetwork.sao.sao";

export interface Order {
  creator: string;
  owner: string;
  id: number;
  provider: string;
  cid: string;
  duration: number;
  expire: number;
  status: number;
  replica: number;
  metadata: string;
  shards: { [key: string]: Shard };
}

export interface Order_ShardsEntry {
  key: string;
  value: Shard | undefined;
}

const baseOrder: object = {
  creator: "",
  owner: "",
  id: 0,
  provider: "",
  cid: "",
  duration: 0,
  expire: 0,
  status: 0,
  replica: 0,
  metadata: "",
};

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.id !== 0) {
      writer.uint32(24).uint64(message.id);
    }
    if (message.provider !== "") {
      writer.uint32(34).string(message.provider);
    }
    if (message.cid !== "") {
      writer.uint32(42).string(message.cid);
    }
    if (message.duration !== 0) {
      writer.uint32(48).int32(message.duration);
    }
    if (message.expire !== 0) {
      writer.uint32(56).int32(message.expire);
    }
    if (message.status !== 0) {
      writer.uint32(64).int32(message.status);
    }
    if (message.replica !== 0) {
      writer.uint32(72).int32(message.replica);
    }
    if (message.metadata !== "") {
      writer.uint32(82).string(message.metadata);
    }
    Object.entries(message.shards).forEach(([key, value]) => {
      Order_ShardsEntry.encode(
        { key: key as any, value },
        writer.uint32(90).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    message.shards = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.provider = reader.string();
          break;
        case 5:
          message.cid = reader.string();
          break;
        case 6:
          message.duration = reader.int32();
          break;
        case 7:
          message.expire = reader.int32();
          break;
        case 8:
          message.status = reader.int32();
          break;
        case 9:
          message.replica = reader.int32();
          break;
        case 10:
          message.metadata = reader.string();
          break;
        case 11:
          const entry11 = Order_ShardsEntry.decode(reader, reader.uint32());
          if (entry11.value !== undefined) {
            message.shards[entry11.key] = entry11.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    message.shards = {};
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = String(object.provider);
    } else {
      message.provider = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.expire !== undefined && object.expire !== null) {
      message.expire = Number(object.expire);
    } else {
      message.expire = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = Number(object.replica);
    } else {
      message.replica = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = String(object.metadata);
    } else {
      message.metadata = "";
    }
    if (object.shards !== undefined && object.shards !== null) {
      Object.entries(object.shards).forEach(([key, value]) => {
        message.shards[key] = Shard.fromJSON(value);
      });
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.owner !== undefined && (obj.owner = message.owner);
    message.id !== undefined && (obj.id = message.id);
    message.provider !== undefined && (obj.provider = message.provider);
    message.cid !== undefined && (obj.cid = message.cid);
    message.duration !== undefined && (obj.duration = message.duration);
    message.expire !== undefined && (obj.expire = message.expire);
    message.status !== undefined && (obj.status = message.status);
    message.replica !== undefined && (obj.replica = message.replica);
    message.metadata !== undefined && (obj.metadata = message.metadata);
    obj.shards = {};
    if (message.shards) {
      Object.entries(message.shards).forEach(([k, v]) => {
        obj.shards[k] = Shard.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    message.shards = {};
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = object.provider;
    } else {
      message.provider = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.expire !== undefined && object.expire !== null) {
      message.expire = object.expire;
    } else {
      message.expire = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = object.replica;
    } else {
      message.replica = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = object.metadata;
    } else {
      message.metadata = "";
    }
    if (object.shards !== undefined && object.shards !== null) {
      Object.entries(object.shards).forEach(([key, value]) => {
        if (value !== undefined) {
          message.shards[key] = Shard.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseOrder_ShardsEntry: object = { key: "" };

export const Order_ShardsEntry = {
  encode(message: Order_ShardsEntry, writer: Writer = Writer.create()): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Shard.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order_ShardsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder_ShardsEntry } as Order_ShardsEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = Shard.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order_ShardsEntry {
    const message = { ...baseOrder_ShardsEntry } as Order_ShardsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Shard.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Order_ShardsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? Shard.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Order_ShardsEntry>): Order_ShardsEntry {
    const message = { ...baseOrder_ShardsEntry } as Order_ShardsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Shard.fromPartial(object.value);
    } else {
      message.value = undefined;
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
