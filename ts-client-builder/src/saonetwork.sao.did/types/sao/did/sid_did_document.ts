/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.did";

export interface SidDidDocument {
  versionId: string;
  keys: string;
}

function createBaseSidDidDocument(): SidDidDocument {
  return { versionId: "", keys: "" };
}

export const SidDidDocument = {
  encode(message: SidDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.versionId !== "") {
      writer.uint32(10).string(message.versionId);
    }
    if (message.keys !== "") {
      writer.uint32(18).string(message.keys);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SidDidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSidDidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.versionId = reader.string();
          break;
        case 2:
          message.keys = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SidDidDocument {
    return {
      versionId: isSet(object.versionId) ? String(object.versionId) : "",
      keys: isSet(object.keys) ? String(object.keys) : "",
    };
  },

  toJSON(message: SidDidDocument): unknown {
    const obj: any = {};
    message.versionId !== undefined && (obj.versionId = message.versionId);
    message.keys !== undefined && (obj.keys = message.keys);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SidDidDocument>, I>>(object: I): SidDidDocument {
    const message = createBaseSidDidDocument();
    message.versionId = object.versionId ?? "";
    message.keys = object.keys ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
