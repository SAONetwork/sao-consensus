/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.earn";

/** Params defines the parameters for the module. */
export interface Params {
  blockReward: number;
  rewardDenom: string;
}

const baseParams: object = { blockReward: 0, rewardDenom: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.blockReward !== 0) {
      writer.uint32(8).int32(message.blockReward);
    }
    if (message.rewardDenom !== "") {
      writer.uint32(18).string(message.rewardDenom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockReward = reader.int32();
          break;
        case 2:
          message.rewardDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.blockReward !== undefined && object.blockReward !== null) {
      message.blockReward = Number(object.blockReward);
    } else {
      message.blockReward = 0;
    }
    if (object.rewardDenom !== undefined && object.rewardDenom !== null) {
      message.rewardDenom = String(object.rewardDenom);
    } else {
      message.rewardDenom = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.blockReward !== undefined &&
      (obj.blockReward = message.blockReward);
    message.rewardDenom !== undefined &&
      (obj.rewardDenom = message.rewardDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.blockReward !== undefined && object.blockReward !== null) {
      message.blockReward = object.blockReward;
    } else {
      message.blockReward = 0;
    }
    if (object.rewardDenom !== undefined && object.rewardDenom !== null) {
      message.rewardDenom = object.rewardDenom;
    } else {
      message.rewardDenom = "";
    }
    return message;
  },
};

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
