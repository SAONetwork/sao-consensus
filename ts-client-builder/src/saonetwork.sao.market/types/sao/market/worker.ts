/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { DecCoin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.market";

export interface Worker {
  workername: string;
  storage: number;
  debt: DecCoin | undefined;
  reward: DecCoin | undefined;
  incomePerSecond: DecCoin | undefined;
  totalStorage: number;
  lastRewardAt: number;
}

function createBaseWorker(): Worker {
  return {
    workername: "",
    storage: 0,
    debt: undefined,
    reward: undefined,
    incomePerSecond: undefined,
    totalStorage: 0,
    lastRewardAt: 0,
  };
}

export const Worker = {
  encode(message: Worker, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.workername !== "") {
      writer.uint32(10).string(message.workername);
    }
    if (message.storage !== 0) {
      writer.uint32(16).uint64(message.storage);
    }
    if (message.debt !== undefined) {
      DecCoin.encode(message.debt, writer.uint32(26).fork()).ldelim();
    }
    if (message.reward !== undefined) {
      DecCoin.encode(message.reward, writer.uint32(34).fork()).ldelim();
    }
    if (message.incomePerSecond !== undefined) {
      DecCoin.encode(message.incomePerSecond, writer.uint32(42).fork()).ldelim();
    }
    if (message.totalStorage !== 0) {
      writer.uint32(48).uint64(message.totalStorage);
    }
    if (message.lastRewardAt !== 0) {
      writer.uint32(56).int64(message.lastRewardAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Worker {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorker();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.workername = reader.string();
          break;
        case 2:
          message.storage = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.debt = DecCoin.decode(reader, reader.uint32());
          break;
        case 4:
          message.reward = DecCoin.decode(reader, reader.uint32());
          break;
        case 5:
          message.incomePerSecond = DecCoin.decode(reader, reader.uint32());
          break;
        case 6:
          message.totalStorage = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.lastRewardAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Worker {
    return {
      workername: isSet(object.workername) ? String(object.workername) : "",
      storage: isSet(object.storage) ? Number(object.storage) : 0,
      debt: isSet(object.debt) ? DecCoin.fromJSON(object.debt) : undefined,
      reward: isSet(object.reward) ? DecCoin.fromJSON(object.reward) : undefined,
      incomePerSecond: isSet(object.incomePerSecond) ? DecCoin.fromJSON(object.incomePerSecond) : undefined,
      totalStorage: isSet(object.totalStorage) ? Number(object.totalStorage) : 0,
      lastRewardAt: isSet(object.lastRewardAt) ? Number(object.lastRewardAt) : 0,
    };
  },

  toJSON(message: Worker): unknown {
    const obj: any = {};
    message.workername !== undefined && (obj.workername = message.workername);
    message.storage !== undefined && (obj.storage = Math.round(message.storage));
    message.debt !== undefined && (obj.debt = message.debt ? DecCoin.toJSON(message.debt) : undefined);
    message.reward !== undefined && (obj.reward = message.reward ? DecCoin.toJSON(message.reward) : undefined);
    message.incomePerSecond !== undefined
      && (obj.incomePerSecond = message.incomePerSecond ? DecCoin.toJSON(message.incomePerSecond) : undefined);
    message.totalStorage !== undefined && (obj.totalStorage = Math.round(message.totalStorage));
    message.lastRewardAt !== undefined && (obj.lastRewardAt = Math.round(message.lastRewardAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Worker>, I>>(object: I): Worker {
    const message = createBaseWorker();
    message.workername = object.workername ?? "";
    message.storage = object.storage ?? 0;
    message.debt = (object.debt !== undefined && object.debt !== null) ? DecCoin.fromPartial(object.debt) : undefined;
    message.reward = (object.reward !== undefined && object.reward !== null)
      ? DecCoin.fromPartial(object.reward)
      : undefined;
    message.incomePerSecond = (object.incomePerSecond !== undefined && object.incomePerSecond !== null)
      ? DecCoin.fromPartial(object.incomePerSecond)
      : undefined;
    message.totalStorage = object.totalStorage ?? 0;
    message.lastRewardAt = object.lastRewardAt ?? 0;
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
