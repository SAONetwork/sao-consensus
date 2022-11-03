/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.model";

export interface ShardMeta {
  shardId: number;
  peer: string;
  cid: string;
}

const baseShardMeta: object = { shardId: 0, peer: "", cid: "" };

export const ShardMeta = {
  encode(message: ShardMeta, writer: Writer = Writer.create()): Writer {
    if (message.shardId !== 0) {
      writer.uint32(8).uint64(message.shardId);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ShardMeta {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShardMeta } as ShardMeta;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.shardId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.peer = reader.string();
          break;
        case 3:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShardMeta {
    const message = { ...baseShardMeta } as ShardMeta;
    if (object.shardId !== undefined && object.shardId !== null) {
      message.shardId = Number(object.shardId);
    } else {
      message.shardId = 0;
    }
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = String(object.peer);
    } else {
      message.peer = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: ShardMeta): unknown {
    const obj: any = {};
    message.shardId !== undefined && (obj.shardId = message.shardId);
    message.peer !== undefined && (obj.peer = message.peer);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<ShardMeta>): ShardMeta {
    const message = { ...baseShardMeta } as ShardMeta;
    if (object.shardId !== undefined && object.shardId !== null) {
      message.shardId = object.shardId;
    } else {
      message.shardId = 0;
    }
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = object.peer;
    } else {
      message.peer = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
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
