/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

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
  extendInfo: string;
  size: number;
  /** 0: new|update, 1:force-push */
  operation: number;
  readonlyDids: string[];
  readwriteDids: string[];
}

function createBaseProposal(): Proposal {
  return {
    owner: "",
    provider: "",
    groupId: "",
    duration: 0,
    replica: 0,
    timeout: 0,
    alias: "",
    dataId: "",
    commitId: "",
    tags: [],
    cid: "",
    rule: "",
    extendInfo: "",
    size: 0,
    operation: 0,
    readonlyDids: [],
    readwriteDids: [],
  };
}

export const Proposal = {
  encode(message: Proposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      writer.uint32(32).uint64(message.duration);
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
    if (message.extendInfo !== "") {
      writer.uint32(106).string(message.extendInfo);
    }
    if (message.size !== 0) {
      writer.uint32(112).uint64(message.size);
    }
    if (message.operation !== 0) {
      writer.uint32(120).uint32(message.operation);
    }
    for (const v of message.readonlyDids) {
      writer.uint32(130).string(v!);
    }
    for (const v of message.readwriteDids) {
      writer.uint32(138).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Proposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposal();
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
          message.duration = longToNumber(reader.uint64() as Long);
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
          message.extendInfo = reader.string();
          break;
        case 14:
          message.size = longToNumber(reader.uint64() as Long);
          break;
        case 15:
          message.operation = reader.uint32();
          break;
        case 16:
          message.readonlyDids.push(reader.string());
          break;
        case 17:
          message.readwriteDids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proposal {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      provider: isSet(object.provider) ? String(object.provider) : "",
      groupId: isSet(object.groupId) ? String(object.groupId) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      replica: isSet(object.replica) ? Number(object.replica) : 0,
      timeout: isSet(object.timeout) ? Number(object.timeout) : 0,
      alias: isSet(object.alias) ? String(object.alias) : "",
      dataId: isSet(object.dataId) ? String(object.dataId) : "",
      commitId: isSet(object.commitId) ? String(object.commitId) : "",
      tags: Array.isArray(object?.tags) ? object.tags.map((e: any) => String(e)) : [],
      cid: isSet(object.cid) ? String(object.cid) : "",
      rule: isSet(object.rule) ? String(object.rule) : "",
      extendInfo: isSet(object.extendInfo) ? String(object.extendInfo) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      operation: isSet(object.operation) ? Number(object.operation) : 0,
      readonlyDids: Array.isArray(object?.readonlyDids) ? object.readonlyDids.map((e: any) => String(e)) : [],
      readwriteDids: Array.isArray(object?.readwriteDids) ? object.readwriteDids.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Proposal): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.provider !== undefined && (obj.provider = message.provider);
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    message.replica !== undefined && (obj.replica = Math.round(message.replica));
    message.timeout !== undefined && (obj.timeout = Math.round(message.timeout));
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
    message.extendInfo !== undefined && (obj.extendInfo = message.extendInfo);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.operation !== undefined && (obj.operation = Math.round(message.operation));
    if (message.readonlyDids) {
      obj.readonlyDids = message.readonlyDids.map((e) => e);
    } else {
      obj.readonlyDids = [];
    }
    if (message.readwriteDids) {
      obj.readwriteDids = message.readwriteDids.map((e) => e);
    } else {
      obj.readwriteDids = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Proposal>, I>>(object: I): Proposal {
    const message = createBaseProposal();
    message.owner = object.owner ?? "";
    message.provider = object.provider ?? "";
    message.groupId = object.groupId ?? "";
    message.duration = object.duration ?? 0;
    message.replica = object.replica ?? 0;
    message.timeout = object.timeout ?? 0;
    message.alias = object.alias ?? "";
    message.dataId = object.dataId ?? "";
    message.commitId = object.commitId ?? "";
    message.tags = object.tags?.map((e) => e) || [];
    message.cid = object.cid ?? "";
    message.rule = object.rule ?? "";
    message.extendInfo = object.extendInfo ?? "";
    message.size = object.size ?? 0;
    message.operation = object.operation ?? 0;
    message.readonlyDids = object.readonlyDids?.map((e) => e) || [];
    message.readwriteDids = object.readwriteDids?.map((e) => e) || [];
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
