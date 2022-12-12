/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin, DecCoin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.node";

export interface Pledge {
  creator: string;
  totalOrderPledged: Coin | undefined;
  totalStoragePledged: Coin | undefined;
  reward: DecCoin | undefined;
  rewardDebt: DecCoin | undefined;
  totalStorage: number;
  lastRewardAt: number;
}

function createBasePledge(): Pledge {
  return {
    creator: "",
    totalOrderPledged: undefined,
    totalStoragePledged: undefined,
    reward: undefined,
    rewardDebt: undefined,
    totalStorage: 0,
    lastRewardAt: 0,
  };
}

export const Pledge = {
  encode(message: Pledge, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.totalOrderPledged !== undefined) {
      Coin.encode(message.totalOrderPledged, writer.uint32(18).fork()).ldelim();
    }
    if (message.totalStoragePledged !== undefined) {
      Coin.encode(message.totalStoragePledged, writer.uint32(26).fork()).ldelim();
    }
    if (message.reward !== undefined) {
      DecCoin.encode(message.reward, writer.uint32(34).fork()).ldelim();
    }
    if (message.rewardDebt !== undefined) {
      DecCoin.encode(message.rewardDebt, writer.uint32(42).fork()).ldelim();
    }
    if (message.totalStorage !== 0) {
      writer.uint32(48).int64(message.totalStorage);
    }
    if (message.lastRewardAt !== 0) {
      writer.uint32(56).int64(message.lastRewardAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Pledge {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePledge();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.totalOrderPledged = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.totalStoragePledged = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.reward = DecCoin.decode(reader, reader.uint32());
          break;
        case 5:
          message.rewardDebt = DecCoin.decode(reader, reader.uint32());
          break;
        case 6:
          message.totalStorage = longToNumber(reader.int64() as Long);
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

  fromJSON(object: any): Pledge {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      totalOrderPledged: isSet(object.totalOrderPledged) ? Coin.fromJSON(object.totalOrderPledged) : undefined,
      totalStoragePledged: isSet(object.totalStoragePledged) ? Coin.fromJSON(object.totalStoragePledged) : undefined,
      reward: isSet(object.reward) ? DecCoin.fromJSON(object.reward) : undefined,
      rewardDebt: isSet(object.rewardDebt) ? DecCoin.fromJSON(object.rewardDebt) : undefined,
      totalStorage: isSet(object.totalStorage) ? Number(object.totalStorage) : 0,
      lastRewardAt: isSet(object.lastRewardAt) ? Number(object.lastRewardAt) : 0,
    };
  },

  toJSON(message: Pledge): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.totalOrderPledged !== undefined
      && (obj.totalOrderPledged = message.totalOrderPledged ? Coin.toJSON(message.totalOrderPledged) : undefined);
    message.totalStoragePledged !== undefined
      && (obj.totalStoragePledged = message.totalStoragePledged ? Coin.toJSON(message.totalStoragePledged) : undefined);
    message.reward !== undefined && (obj.reward = message.reward ? DecCoin.toJSON(message.reward) : undefined);
    message.rewardDebt !== undefined
      && (obj.rewardDebt = message.rewardDebt ? DecCoin.toJSON(message.rewardDebt) : undefined);
    message.totalStorage !== undefined && (obj.totalStorage = Math.round(message.totalStorage));
    message.lastRewardAt !== undefined && (obj.lastRewardAt = Math.round(message.lastRewardAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Pledge>, I>>(object: I): Pledge {
    const message = createBasePledge();
    message.creator = object.creator ?? "";
    message.totalOrderPledged = (object.totalOrderPledged !== undefined && object.totalOrderPledged !== null)
      ? Coin.fromPartial(object.totalOrderPledged)
      : undefined;
    message.totalStoragePledged = (object.totalStoragePledged !== undefined && object.totalStoragePledged !== null)
      ? Coin.fromPartial(object.totalStoragePledged)
      : undefined;
    message.reward = (object.reward !== undefined && object.reward !== null)
      ? DecCoin.fromPartial(object.reward)
      : undefined;
    message.rewardDebt = (object.rewardDebt !== undefined && object.rewardDebt !== null)
      ? DecCoin.fromPartial(object.rewardDebt)
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
