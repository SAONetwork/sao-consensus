/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin, DecCoin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.node";

export interface Pool {
  totalPledged: Coin | undefined;
  totalReward: Coin | undefined;
  accPledgePerByte: DecCoin | undefined;
  accRewardPerByte: DecCoin | undefined;
  rewardPerBlock: DecCoin | undefined;
  totalStorage: number;
  lastRewardBlock: number;
}

function createBasePool(): Pool {
  return {
    totalPledged: undefined,
    totalReward: undefined,
    accPledgePerByte: undefined,
    accRewardPerByte: undefined,
    rewardPerBlock: undefined,
    totalStorage: 0,
    lastRewardBlock: 0,
  };
}

export const Pool = {
  encode(message: Pool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.totalPledged !== undefined) {
      Coin.encode(message.totalPledged, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalReward !== undefined) {
      Coin.encode(message.totalReward, writer.uint32(18).fork()).ldelim();
    }
    if (message.accPledgePerByte !== undefined) {
      DecCoin.encode(message.accPledgePerByte, writer.uint32(26).fork()).ldelim();
    }
    if (message.accRewardPerByte !== undefined) {
      DecCoin.encode(message.accRewardPerByte, writer.uint32(34).fork()).ldelim();
    }
    if (message.rewardPerBlock !== undefined) {
      DecCoin.encode(message.rewardPerBlock, writer.uint32(42).fork()).ldelim();
    }
    if (message.totalStorage !== 0) {
      writer.uint32(48).int64(message.totalStorage);
    }
    if (message.lastRewardBlock !== 0) {
      writer.uint32(56).int64(message.lastRewardBlock);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalPledged = Coin.decode(reader, reader.uint32());
          break;
        case 2:
          message.totalReward = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.accPledgePerByte = DecCoin.decode(reader, reader.uint32());
          break;
        case 4:
          message.accRewardPerByte = DecCoin.decode(reader, reader.uint32());
          break;
        case 5:
          message.rewardPerBlock = DecCoin.decode(reader, reader.uint32());
          break;
        case 6:
          message.totalStorage = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.lastRewardBlock = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    return {
      totalPledged: isSet(object.totalPledged) ? Coin.fromJSON(object.totalPledged) : undefined,
      totalReward: isSet(object.totalReward) ? Coin.fromJSON(object.totalReward) : undefined,
      accPledgePerByte: isSet(object.accPledgePerByte) ? DecCoin.fromJSON(object.accPledgePerByte) : undefined,
      accRewardPerByte: isSet(object.accRewardPerByte) ? DecCoin.fromJSON(object.accRewardPerByte) : undefined,
      rewardPerBlock: isSet(object.rewardPerBlock) ? DecCoin.fromJSON(object.rewardPerBlock) : undefined,
      totalStorage: isSet(object.totalStorage) ? Number(object.totalStorage) : 0,
      lastRewardBlock: isSet(object.lastRewardBlock) ? Number(object.lastRewardBlock) : 0,
    };
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.totalPledged !== undefined
      && (obj.totalPledged = message.totalPledged ? Coin.toJSON(message.totalPledged) : undefined);
    message.totalReward !== undefined
      && (obj.totalReward = message.totalReward ? Coin.toJSON(message.totalReward) : undefined);
    message.accPledgePerByte !== undefined
      && (obj.accPledgePerByte = message.accPledgePerByte ? DecCoin.toJSON(message.accPledgePerByte) : undefined);
    message.accRewardPerByte !== undefined
      && (obj.accRewardPerByte = message.accRewardPerByte ? DecCoin.toJSON(message.accRewardPerByte) : undefined);
    message.rewardPerBlock !== undefined
      && (obj.rewardPerBlock = message.rewardPerBlock ? DecCoin.toJSON(message.rewardPerBlock) : undefined);
    message.totalStorage !== undefined && (obj.totalStorage = Math.round(message.totalStorage));
    message.lastRewardBlock !== undefined && (obj.lastRewardBlock = Math.round(message.lastRewardBlock));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Pool>, I>>(object: I): Pool {
    const message = createBasePool();
    message.totalPledged = (object.totalPledged !== undefined && object.totalPledged !== null)
      ? Coin.fromPartial(object.totalPledged)
      : undefined;
    message.totalReward = (object.totalReward !== undefined && object.totalReward !== null)
      ? Coin.fromPartial(object.totalReward)
      : undefined;
    message.accPledgePerByte = (object.accPledgePerByte !== undefined && object.accPledgePerByte !== null)
      ? DecCoin.fromPartial(object.accPledgePerByte)
      : undefined;
    message.accRewardPerByte = (object.accRewardPerByte !== undefined && object.accRewardPerByte !== null)
      ? DecCoin.fromPartial(object.accRewardPerByte)
      : undefined;
    message.rewardPerBlock = (object.rewardPerBlock !== undefined && object.rewardPerBlock !== null)
      ? DecCoin.fromPartial(object.rewardPerBlock)
      : undefined;
    message.totalStorage = object.totalStorage ?? 0;
    message.lastRewardBlock = object.lastRewardBlock ?? 0;
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
