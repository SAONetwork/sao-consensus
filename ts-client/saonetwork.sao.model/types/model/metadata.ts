/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.model";

export interface Metadata {
  dataId: string;
  owner: string;
  alias: string;
  groupId: string;
  orderId: number;
  tags: string[];
  cid: string;
  commits: string[];
  extendInfo: string;
  update: boolean;
  commit: string;
  rule: string;
}

const baseMetadata: object = {
  dataId: "",
  owner: "",
  alias: "",
  groupId: "",
  orderId: 0,
  tags: "",
  cid: "",
  commits: "",
  extendInfo: "",
  update: false,
  commit: "",
  rule: "",
};

export const Metadata = {
  encode(message: Metadata, writer: Writer = Writer.create()): Writer {
    if (message.dataId !== "") {
      writer.uint32(10).string(message.dataId);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    if (message.groupId !== "") {
      writer.uint32(34).string(message.groupId);
    }
    if (message.orderId !== 0) {
      writer.uint32(40).uint64(message.orderId);
    }
    for (const v of message.tags) {
      writer.uint32(50).string(v!);
    }
    if (message.cid !== "") {
      writer.uint32(58).string(message.cid);
    }
    for (const v of message.commits) {
      writer.uint32(66).string(v!);
    }
    if (message.extendInfo !== "") {
      writer.uint32(74).string(message.extendInfo);
    }
    if (message.update === true) {
      writer.uint32(80).bool(message.update);
    }
    if (message.commit !== "") {
      writer.uint32(90).string(message.commit);
    }
    if (message.rule !== "") {
      writer.uint32(98).string(message.rule);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Metadata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMetadata } as Metadata;
    message.tags = [];
    message.commits = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dataId = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.alias = reader.string();
          break;
        case 4:
          message.groupId = reader.string();
          break;
        case 5:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.tags.push(reader.string());
          break;
        case 7:
          message.cid = reader.string();
          break;
        case 8:
          message.commits.push(reader.string());
          break;
        case 9:
          message.extendInfo = reader.string();
          break;
        case 10:
          message.update = reader.bool();
          break;
        case 11:
          message.commit = reader.string();
          break;
        case 12:
          message.rule = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Metadata {
    const message = { ...baseMetadata } as Metadata;
    message.tags = [];
    message.commits = [];
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = String(object.dataId);
    } else {
      message.dataId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (object.groupId !== undefined && object.groupId !== null) {
      message.groupId = String(object.groupId);
    } else {
      message.groupId = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(String(e));
      }
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.commits !== undefined && object.commits !== null) {
      for (const e of object.commits) {
        message.commits.push(String(e));
      }
    }
    if (object.extendInfo !== undefined && object.extendInfo !== null) {
      message.extendInfo = String(object.extendInfo);
    } else {
      message.extendInfo = "";
    }
    if (object.update !== undefined && object.update !== null) {
      message.update = Boolean(object.update);
    } else {
      message.update = false;
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = String(object.commit);
    } else {
      message.commit = "";
    }
    if (object.rule !== undefined && object.rule !== null) {
      message.rule = String(object.rule);
    } else {
      message.rule = "";
    }
    return message;
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.dataId !== undefined && (obj.dataId = message.dataId);
    message.owner !== undefined && (obj.owner = message.owner);
    message.alias !== undefined && (obj.alias = message.alias);
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    message.cid !== undefined && (obj.cid = message.cid);
    if (message.commits) {
      obj.commits = message.commits.map((e) => e);
    } else {
      obj.commits = [];
    }
    message.extendInfo !== undefined && (obj.extendInfo = message.extendInfo);
    message.update !== undefined && (obj.update = message.update);
    message.commit !== undefined && (obj.commit = message.commit);
    message.rule !== undefined && (obj.rule = message.rule);
    return obj;
  },

  fromPartial(object: DeepPartial<Metadata>): Metadata {
    const message = { ...baseMetadata } as Metadata;
    message.tags = [];
    message.commits = [];
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = object.dataId;
    } else {
      message.dataId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (object.groupId !== undefined && object.groupId !== null) {
      message.groupId = object.groupId;
    } else {
      message.groupId = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(e);
      }
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.commits !== undefined && object.commits !== null) {
      for (const e of object.commits) {
        message.commits.push(e);
      }
    }
    if (object.extendInfo !== undefined && object.extendInfo !== null) {
      message.extendInfo = object.extendInfo;
    } else {
      message.extendInfo = "";
    }
    if (object.update !== undefined && object.update !== null) {
      message.update = object.update;
    } else {
      message.update = false;
    }
    if (object.commit !== undefined && object.commit !== null) {
      message.commit = object.commit;
    } else {
      message.commit = "";
    }
    if (object.rule !== undefined && object.rule !== null) {
      message.rule = object.rule;
    } else {
      message.rule = "";
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
