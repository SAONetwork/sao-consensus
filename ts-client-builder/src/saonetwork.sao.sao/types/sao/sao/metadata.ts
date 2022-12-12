/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "saonetwork.sao.sao";

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
  duration: number;
  createdAt: number;
  provider: string;
  expire: number;
  status: number;
  replica: number;
  amount: Coin | undefined;
  size: number;
  operation: number;
}

function createBaseMetadata(): Metadata {
  return {
    dataId: "",
    owner: "",
    alias: "",
    groupId: "",
    orderId: 0,
    tags: [],
    cid: "",
    commits: [],
    extendInfo: "",
    update: false,
    commit: "",
    rule: "",
    duration: 0,
    createdAt: 0,
    provider: "",
    expire: 0,
    status: 0,
    replica: 0,
    amount: undefined,
    size: 0,
    operation: 0,
  };
}

export const Metadata = {
  encode(message: Metadata, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.duration !== 0) {
      writer.uint32(104).uint64(message.duration);
    }
    if (message.createdAt !== 0) {
      writer.uint32(112).uint64(message.createdAt);
    }
    if (message.provider !== "") {
      writer.uint32(122).string(message.provider);
    }
    if (message.expire !== 0) {
      writer.uint32(128).int32(message.expire);
    }
    if (message.status !== 0) {
      writer.uint32(136).int32(message.status);
    }
    if (message.replica !== 0) {
      writer.uint32(144).int32(message.replica);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(154).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(160).uint64(message.size);
    }
    if (message.operation !== 0) {
      writer.uint32(168).uint32(message.operation);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Metadata {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMetadata();
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
        case 13:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.createdAt = longToNumber(reader.uint64() as Long);
          break;
        case 15:
          message.provider = reader.string();
          break;
        case 16:
          message.expire = reader.int32();
          break;
        case 17:
          message.status = reader.int32();
          break;
        case 18:
          message.replica = reader.int32();
          break;
        case 19:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        case 20:
          message.size = longToNumber(reader.uint64() as Long);
          break;
        case 21:
          message.operation = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Metadata {
    return {
      dataId: isSet(object.dataId) ? String(object.dataId) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      alias: isSet(object.alias) ? String(object.alias) : "",
      groupId: isSet(object.groupId) ? String(object.groupId) : "",
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
      tags: Array.isArray(object?.tags) ? object.tags.map((e: any) => String(e)) : [],
      cid: isSet(object.cid) ? String(object.cid) : "",
      commits: Array.isArray(object?.commits) ? object.commits.map((e: any) => String(e)) : [],
      extendInfo: isSet(object.extendInfo) ? String(object.extendInfo) : "",
      update: isSet(object.update) ? Boolean(object.update) : false,
      commit: isSet(object.commit) ? String(object.commit) : "",
      rule: isSet(object.rule) ? String(object.rule) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      provider: isSet(object.provider) ? String(object.provider) : "",
      expire: isSet(object.expire) ? Number(object.expire) : 0,
      status: isSet(object.status) ? Number(object.status) : 0,
      replica: isSet(object.replica) ? Number(object.replica) : 0,
      amount: isSet(object.amount) ? Coin.fromJSON(object.amount) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      operation: isSet(object.operation) ? Number(object.operation) : 0,
    };
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.dataId !== undefined && (obj.dataId = message.dataId);
    message.owner !== undefined && (obj.owner = message.owner);
    message.alias !== undefined && (obj.alias = message.alias);
    message.groupId !== undefined && (obj.groupId = message.groupId);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
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
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.provider !== undefined && (obj.provider = message.provider);
    message.expire !== undefined && (obj.expire = Math.round(message.expire));
    message.status !== undefined && (obj.status = Math.round(message.status));
    message.replica !== undefined && (obj.replica = Math.round(message.replica));
    message.amount !== undefined && (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.operation !== undefined && (obj.operation = Math.round(message.operation));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Metadata>, I>>(object: I): Metadata {
    const message = createBaseMetadata();
    message.dataId = object.dataId ?? "";
    message.owner = object.owner ?? "";
    message.alias = object.alias ?? "";
    message.groupId = object.groupId ?? "";
    message.orderId = object.orderId ?? 0;
    message.tags = object.tags?.map((e) => e) || [];
    message.cid = object.cid ?? "";
    message.commits = object.commits?.map((e) => e) || [];
    message.extendInfo = object.extendInfo ?? "";
    message.update = object.update ?? false;
    message.commit = object.commit ?? "";
    message.rule = object.rule ?? "";
    message.duration = object.duration ?? 0;
    message.createdAt = object.createdAt ?? 0;
    message.provider = object.provider ?? "";
    message.expire = object.expire ?? 0;
    message.status = object.status ?? 0;
    message.replica = object.replica ?? 0;
    message.amount = (object.amount !== undefined && object.amount !== null)
      ? Coin.fromPartial(object.amount)
      : undefined;
    message.size = object.size ?? 0;
    message.operation = object.operation ?? 0;
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
