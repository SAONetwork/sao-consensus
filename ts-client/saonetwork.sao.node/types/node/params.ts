/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

/** Params defines the parameters for the module. */
export interface Params {
  blockReward: number;
  earnDenom: string;
}

const baseParams: object = { blockReward: 0, earnDenom: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.blockReward !== 0) {
      writer.uint32(8).uint64(message.blockReward);
    }
    if (message.earnDenom !== "") {
      writer.uint32(18).string(message.earnDenom);
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
          message.blockReward = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.earnDenom = reader.string();
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
    if (object.earnDenom !== undefined && object.earnDenom !== null) {
      message.earnDenom = String(object.earnDenom);
    } else {
      message.earnDenom = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.blockReward !== undefined &&
      (obj.blockReward = message.blockReward);
    message.earnDenom !== undefined && (obj.earnDenom = message.earnDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.blockReward !== undefined && object.blockReward !== null) {
      message.blockReward = object.blockReward;
    } else {
      message.blockReward = 0;
    }
    if (object.earnDenom !== undefined && object.earnDenom !== null) {
      message.earnDenom = object.earnDenom;
    } else {
      message.earnDenom = "";
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
