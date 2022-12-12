/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Metadata } from "./metadata";
import { Shard } from "./shard";

export const protobufPackage = "saonetwork.sao.order";

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
  metadata: Metadata | undefined;
  shards: { [key: string]: Shard };
  amount: Coin | undefined;
  size: number;
  operation: number;
  createdAt: number;
}

export interface Order_ShardsEntry {
  key: string;
  value: Shard | undefined;
}

function createBaseOrder(): Order {
  return {
    creator: "",
    owner: "",
    id: 0,
    provider: "",
    cid: "",
    duration: 0,
    expire: 0,
    status: 0,
    replica: 0,
    metadata: undefined,
    shards: {},
    amount: undefined,
    size: 0,
    operation: 0,
    createdAt: 0,
  };
}

export const Order = {
  encode(message: Order, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      writer.uint32(48).uint64(message.duration);
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
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(82).fork()).ldelim();
    }
    Object.entries(message.shards).forEach(([key, value]) => {
      Order_ShardsEntry.encode({ key: key as any, value }, writer.uint32(90).fork()).ldelim();
    });
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(98).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(104).uint64(message.size);
    }
    if (message.operation !== 0) {
      writer.uint32(112).uint32(message.operation);
    }
    if (message.createdAt !== 0) {
      writer.uint32(120).int32(message.createdAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder();
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
          message.duration = longToNumber(reader.uint64() as Long);
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
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 11:
          const entry11 = Order_ShardsEntry.decode(reader, reader.uint32());
          if (entry11.value !== undefined) {
            message.shards[entry11.key] = entry11.value;
          }
          break;
        case 12:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 13:
          message.size = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.operation = reader.uint32();
          break;
        case 15:
          message.createdAt = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      provider: isSet(object.provider) ? String(object.provider) : "",
      cid: isSet(object.cid) ? String(object.cid) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      expire: isSet(object.expire) ? Number(object.expire) : 0,
      status: isSet(object.status) ? Number(object.status) : 0,
      replica: isSet(object.replica) ? Number(object.replica) : 0,
      metadata: isSet(object.metadata) ? Metadata.fromJSON(object.metadata) : undefined,
      shards: isObject(object.shards)
        ? Object.entries(object.shards).reduce<{ [key: string]: Shard }>((acc, [key, value]) => {
          acc[key] = Shard.fromJSON(value);
          return acc;
        }, {})
        : {},
      amount: isSet(object.amount) ? Coin.fromJSON(object.amount) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      operation: isSet(object.operation) ? Number(object.operation) : 0,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
    };
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.owner !== undefined && (obj.owner = message.owner);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.provider !== undefined && (obj.provider = message.provider);
    message.cid !== undefined && (obj.cid = message.cid);
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    message.expire !== undefined && (obj.expire = Math.round(message.expire));
    message.status !== undefined && (obj.status = Math.round(message.status));
    message.replica !== undefined && (obj.replica = Math.round(message.replica));
    message.metadata !== undefined && (obj.metadata = message.metadata ? Metadata.toJSON(message.metadata) : undefined);
    obj.shards = {};
    if (message.shards) {
      Object.entries(message.shards).forEach(([k, v]) => {
        obj.shards[k] = Shard.toJSON(v);
      });
    }
    message.amount !== undefined && (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.operation !== undefined && (obj.operation = Math.round(message.operation));
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Order>, I>>(object: I): Order {
    const message = createBaseOrder();
    message.creator = object.creator ?? "";
    message.owner = object.owner ?? "";
    message.id = object.id ?? 0;
    message.provider = object.provider ?? "";
    message.cid = object.cid ?? "";
    message.duration = object.duration ?? 0;
    message.expire = object.expire ?? 0;
    message.status = object.status ?? 0;
    message.replica = object.replica ?? 0;
    message.metadata = (object.metadata !== undefined && object.metadata !== null)
      ? Metadata.fromPartial(object.metadata)
      : undefined;
    message.shards = Object.entries(object.shards ?? {}).reduce<{ [key: string]: Shard }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Shard.fromPartial(value);
      }
      return acc;
    }, {});
    message.amount = (object.amount !== undefined && object.amount !== null)
      ? Coin.fromPartial(object.amount)
      : undefined;
    message.size = object.size ?? 0;
    message.operation = object.operation ?? 0;
    message.createdAt = object.createdAt ?? 0;
    return message;
  },
};

function createBaseOrder_ShardsEntry(): Order_ShardsEntry {
  return { key: "", value: undefined };
}

export const Order_ShardsEntry = {
  encode(message: Order_ShardsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Shard.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order_ShardsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder_ShardsEntry();
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
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Shard.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Order_ShardsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? Shard.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Order_ShardsEntry>, I>>(object: I): Order_ShardsEntry {
    const message = createBaseOrder_ShardsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null) ? Shard.fromPartial(object.value) : undefined;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
