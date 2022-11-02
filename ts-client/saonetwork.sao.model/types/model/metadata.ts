/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.model";

export interface Metadata {
  dataId: string;
  owner: string;
  alias: string;
  familyId: string;
  tags: string[];
  cids: string[];
  extendInfo: string;
}

const baseMetadata: object = {
  dataId: "",
  owner: "",
  alias: "",
  familyId: "",
  tags: "",
  cids: "",
  extendInfo: "",
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
    if (message.familyId !== "") {
      writer.uint32(34).string(message.familyId);
    }
    for (const v of message.tags) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.cids) {
      writer.uint32(50).string(v!);
    }
    if (message.extendInfo !== "") {
      writer.uint32(58).string(message.extendInfo);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Metadata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMetadata } as Metadata;
    message.tags = [];
    message.cids = [];
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
          message.familyId = reader.string();
          break;
        case 5:
          message.tags.push(reader.string());
          break;
        case 6:
          message.cids.push(reader.string());
          break;
        case 7:
          message.extendInfo = reader.string();
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
    message.cids = [];
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
    if (object.familyId !== undefined && object.familyId !== null) {
      message.familyId = String(object.familyId);
    } else {
      message.familyId = "";
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(String(e));
      }
    }
    if (object.cids !== undefined && object.cids !== null) {
      for (const e of object.cids) {
        message.cids.push(String(e));
      }
    }
    if (object.extendInfo !== undefined && object.extendInfo !== null) {
      message.extendInfo = String(object.extendInfo);
    } else {
      message.extendInfo = "";
    }
    return message;
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.dataId !== undefined && (obj.dataId = message.dataId);
    message.owner !== undefined && (obj.owner = message.owner);
    message.alias !== undefined && (obj.alias = message.alias);
    message.familyId !== undefined && (obj.familyId = message.familyId);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    if (message.cids) {
      obj.cids = message.cids.map((e) => e);
    } else {
      obj.cids = [];
    }
    message.extendInfo !== undefined && (obj.extendInfo = message.extendInfo);
    return obj;
  },

  fromPartial(object: DeepPartial<Metadata>): Metadata {
    const message = { ...baseMetadata } as Metadata;
    message.tags = [];
    message.cids = [];
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
    if (object.familyId !== undefined && object.familyId !== null) {
      message.familyId = object.familyId;
    } else {
      message.familyId = "";
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(e);
      }
    }
    if (object.cids !== undefined && object.cids !== null) {
      for (const e of object.cids) {
        message.cids.push(e);
      }
    }
    if (object.extendInfo !== undefined && object.extendInfo !== null) {
      message.extendInfo = object.extendInfo;
    } else {
      message.extendInfo = "";
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
