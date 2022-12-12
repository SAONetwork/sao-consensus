/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.did";

export interface BindingProof {
  version: number;
  message: string;
  signature: string;
  account: string;
  did: string;
  timestamp: number;
}

function createBaseBindingProof(): BindingProof {
  return { version: 0, message: "", signature: "", account: "", did: "", timestamp: 0 };
}

export const BindingProof = {
  encode(message: BindingProof, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.version !== 0) {
      writer.uint32(8).int32(message.version);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    if (message.account !== "") {
      writer.uint32(34).string(message.account);
    }
    if (message.did !== "") {
      writer.uint32(42).string(message.did);
    }
    if (message.timestamp !== 0) {
      writer.uint32(48).uint64(message.timestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BindingProof {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBindingProof();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.version = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        case 4:
          message.account = reader.string();
          break;
        case 5:
          message.did = reader.string();
          break;
        case 6:
          message.timestamp = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BindingProof {
    return {
      version: isSet(object.version) ? Number(object.version) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      signature: isSet(object.signature) ? String(object.signature) : "",
      account: isSet(object.account) ? String(object.account) : "",
      did: isSet(object.did) ? String(object.did) : "",
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
    };
  },

  toJSON(message: BindingProof): unknown {
    const obj: any = {};
    message.version !== undefined && (obj.version = Math.round(message.version));
    message.message !== undefined && (obj.message = message.message);
    message.signature !== undefined && (obj.signature = message.signature);
    message.account !== undefined && (obj.account = message.account);
    message.did !== undefined && (obj.did = message.did);
    message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BindingProof>, I>>(object: I): BindingProof {
    const message = createBaseBindingProof();
    message.version = object.version ?? 0;
    message.message = object.message ?? "";
    message.signature = object.signature ?? "";
    message.account = object.account ?? "";
    message.did = object.did ?? "";
    message.timestamp = object.timestamp ?? 0;
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
