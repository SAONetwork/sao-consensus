/* eslint-disable */
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

export interface Pledge {
  creator: string;
  pledged: Coin | undefined;
  reward: Coin | undefined;
  rewardDebt: Coin | undefined;
}

const basePledge: object = { creator: "" };

export const Pledge = {
  encode(message: Pledge, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.pledged !== undefined) {
      Coin.encode(message.pledged, writer.uint32(18).fork()).ldelim();
    }
    if (message.reward !== undefined) {
      Coin.encode(message.reward, writer.uint32(26).fork()).ldelim();
    }
    if (message.rewardDebt !== undefined) {
      Coin.encode(message.rewardDebt, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pledge {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePledge } as Pledge;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.pledged = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.reward = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.rewardDebt = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pledge {
    const message = { ...basePledge } as Pledge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.pledged !== undefined && object.pledged !== null) {
      message.pledged = Coin.fromJSON(object.pledged);
    } else {
      message.pledged = undefined;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = Coin.fromJSON(object.reward);
    } else {
      message.reward = undefined;
    }
    if (object.rewardDebt !== undefined && object.rewardDebt !== null) {
      message.rewardDebt = Coin.fromJSON(object.rewardDebt);
    } else {
      message.rewardDebt = undefined;
    }
    return message;
  },

  toJSON(message: Pledge): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pledged !== undefined &&
      (obj.pledged = message.pledged
        ? Coin.toJSON(message.pledged)
        : undefined);
    message.reward !== undefined &&
      (obj.reward = message.reward ? Coin.toJSON(message.reward) : undefined);
    message.rewardDebt !== undefined &&
      (obj.rewardDebt = message.rewardDebt
        ? Coin.toJSON(message.rewardDebt)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Pledge>): Pledge {
    const message = { ...basePledge } as Pledge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.pledged !== undefined && object.pledged !== null) {
      message.pledged = Coin.fromPartial(object.pledged);
    } else {
      message.pledged = undefined;
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = Coin.fromPartial(object.reward);
    } else {
      message.reward = undefined;
    }
    if (object.rewardDebt !== undefined && object.rewardDebt !== null) {
      message.rewardDebt = Coin.fromPartial(object.rewardDebt);
    } else {
      message.rewardDebt = undefined;
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
