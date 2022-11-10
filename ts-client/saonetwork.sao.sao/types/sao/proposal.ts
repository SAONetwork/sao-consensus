/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface Proposal {
  owner: string;
  provider: string;
  groupId: string;
  duration: number;
  replica: number;
  timeout: number;
  alias: string;
  dataId: string;
  commitId: string;
  tags: string[];
  cid: string;
  rule: string;
  isUpdate: boolean;
  extendInfo: string;
}

const baseProposal: object = {
  owner: "",
  provider: "",
  groupId: "",
  duration: 0,
  replica: 0,
  timeout: 0,
  alias: "",
  dataId: "",
  commitId: "",
  tags: "",
  cid: "",
  rule: "",
  isUpdate: false,
  extendInfo: "",
};

export const Proposal = {
  encode(message: Proposal, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.provider !== "") {
      writer.uint32(18).string(message.provider);
    }
    if (message.groupId !== "") {
      writer.uint32(26).string(message.groupId);
    }
    if (message.duration !== 0) {
      writer.uint32(32).int32(message.duration);
    }
    if (message.replica !== 0) {
      writer.uint32(40).int32(message.replica);
    }
    if (message.timeout !== 0) {
      writer.uint32(48).int32(message.timeout);
    }
    if (message.alias !== "") {
      writer.uint32(58).string(message.alias);
    }
    if (message.dataId !== "") {
      writer.uint32(66).string(message.dataId);
    }
    if (message.commitId !== "") {
      writer.uint32(74).string(message.commitId);
    }
    for (const v of message.tags) {
      writer.uint32(82).string(v!);
    }
    if (message.cid !== "") {
      writer.uint32(90).string(message.cid);
    }
    if (message.rule !== "") {
      writer.uint32(98).string(message.rule);
    }
    if (message.isUpdate === true) {
      writer.uint32(104).bool(message.isUpdate);
    }
    if (message.extendInfo !== "") {
      writer.uint32(114).string(message.extendInfo);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Proposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProposal } as Proposal;
    message.tags = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.provider = reader.string();
          break;
        case 3:
          message.groupId = reader.string();
          break;
        case 4:
          message.duration = reader.int32();
          break;
        case 5:
          message.replica = reader.int32();
          break;
        case 6:
          message.timeout = reader.int32();
          break;
        case 7:
          message.alias = reader.string();
          break;
        case 8:
          message.dataId = reader.string();
          break;
        case 9:
          message.commitId = reader.string();
          break;
        case 10:
          message.tags.push(reader.string());
          break;
        case 11:
          message.cid = reader.string();
          break;
        case 12:
          message.rule = reader.string();
          break;
        case 13:
          message.isUpdate = reader.bool();
          break;
        case 14:
          message.extendInfo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proposal {
    const message = { ...baseProposal } as Proposal;
    message.tags = [];
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = String(object.provider);
    } else {
      message.provider = "";
    }
    if (object.groupId !== undefined && object.groupId !== null) {
      message.groupId = String(object.groupId);
    } else {
      message.groupId = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = Number(object.replica);
    } else {
      message.replica = 0;
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = Number(object.timeout);
    } else {
      message.timeout = 0;
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = String(object.dataId);
    } else {
      message.dataId = "";
    }
    if (object.commitId !== undefined && object.commitId !== null) {
      message.commitId = String(object.commitId);
    } else {
      message.commitId = "";
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
    if (object.rule !== undefined && object.rule !== null) {
      message.rule = String(object.rule);
    } else {
      message.rule = "";
    }
    if (object.isUpdate !== undefined && object.isUpdate !== null) {
      message.isUpdate = Boolean(object.isUpdate);
    } else {
      message.isUpdate = false;
    }
    if (object.extendInfo !== undefined && object.extendInfo !== null) {
      message.extendInfo = String(object.extendInfo);
    } else {
      message.extendInfo = "";
    }
    return message;
  },

  toJSON(message: Proposal): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.provider !== undefined && (obj.provider = message.provider);
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.duration !== undefined && (obj.duration = message.duration);
    message.replica !== undefined && (obj.replica = message.replica);
    message.timeout !== undefined && (obj.timeout = message.timeout);
    message.alias !== undefined && (obj.alias = message.alias);
    message.dataId !== undefined && (obj.dataId = message.dataId);
    message.commitId !== undefined && (obj.commitId = message.commitId);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    message.cid !== undefined && (obj.cid = message.cid);
    message.rule !== undefined && (obj.rule = message.rule);
    message.isUpdate !== undefined && (obj.isUpdate = message.isUpdate);
    message.extendInfo !== undefined && (obj.extendInfo = message.extendInfo);
    return obj;
  },

  fromPartial(object: DeepPartial<Proposal>): Proposal {
    const message = { ...baseProposal } as Proposal;
    message.tags = [];
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = object.provider;
    } else {
      message.provider = "";
    }
    if (object.groupId !== undefined && object.groupId !== null) {
      message.groupId = object.groupId;
    } else {
      message.groupId = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = object.replica;
    } else {
      message.replica = 0;
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = object.timeout;
    } else {
      message.timeout = 0;
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (object.dataId !== undefined && object.dataId !== null) {
      message.dataId = object.dataId;
    } else {
      message.dataId = "";
    }
    if (object.commitId !== undefined && object.commitId !== null) {
      message.commitId = object.commitId;
    } else {
      message.commitId = "";
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
    if (object.rule !== undefined && object.rule !== null) {
      message.rule = object.rule;
    } else {
      message.rule = "";
    }
    if (object.isUpdate !== undefined && object.isUpdate !== null) {
      message.isUpdate = object.isUpdate;
    } else {
      message.isUpdate = false;
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
