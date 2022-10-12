/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.earn";

export interface Pool {
  denom: Coin | undefined;
  coinPerShare: number;
  lastRewardBlock: number;
  totalReward: Coin | undefined;
}

const basePool: object = { coinPerShare: 0, lastRewardBlock: 0 };

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.denom !== undefined) {
      Coin.encode(message.denom, writer.uint32(10).fork()).ldelim();
    }
    if (message.coinPerShare !== 0) {
      writer.uint32(16).uint64(message.coinPerShare);
    }
    if (message.lastRewardBlock !== 0) {
      writer.uint32(24).int64(message.lastRewardBlock);
    }
    if (message.totalReward !== undefined) {
      Coin.encode(message.totalReward, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = Coin.decode(reader, reader.uint32());
          break;
        case 2:
          message.coinPerShare = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.lastRewardBlock = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.totalReward = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = Coin.fromJSON(object.denom);
    } else {
      message.denom = undefined;
    }
    if (object.coinPerShare !== undefined && object.coinPerShare !== null) {
      message.coinPerShare = Number(object.coinPerShare);
    } else {
      message.coinPerShare = 0;
    }
    if (
      object.lastRewardBlock !== undefined &&
      object.lastRewardBlock !== null
    ) {
      message.lastRewardBlock = Number(object.lastRewardBlock);
    } else {
      message.lastRewardBlock = 0;
    }
    if (object.totalReward !== undefined && object.totalReward !== null) {
      message.totalReward = Coin.fromJSON(object.totalReward);
    } else {
      message.totalReward = undefined;
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.denom !== undefined &&
      (obj.denom = message.denom ? Coin.toJSON(message.denom) : undefined);
    message.coinPerShare !== undefined &&
      (obj.coinPerShare = message.coinPerShare);
    message.lastRewardBlock !== undefined &&
      (obj.lastRewardBlock = message.lastRewardBlock);
    message.totalReward !== undefined &&
      (obj.totalReward = message.totalReward
        ? Coin.toJSON(message.totalReward)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = Coin.fromPartial(object.denom);
    } else {
      message.denom = undefined;
    }
    if (object.coinPerShare !== undefined && object.coinPerShare !== null) {
      message.coinPerShare = object.coinPerShare;
    } else {
      message.coinPerShare = 0;
    }
    if (
      object.lastRewardBlock !== undefined &&
      object.lastRewardBlock !== null
    ) {
      message.lastRewardBlock = object.lastRewardBlock;
    } else {
      message.lastRewardBlock = 0;
    }
    if (object.totalReward !== undefined && object.totalReward !== null) {
      message.totalReward = Coin.fromPartial(object.totalReward);
    } else {
      message.totalReward = undefined;
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
