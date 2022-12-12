/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.did";

export interface SidDocumentVersion {
  docId: string;
  versionList: string[];
}

function createBaseSidDocumentVersion(): SidDocumentVersion {
  return { docId: "", versionList: [] };
}

export const SidDocumentVersion = {
  encode(message: SidDocumentVersion, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.docId !== "") {
      writer.uint32(10).string(message.docId);
    }
    for (const v of message.versionList) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SidDocumentVersion {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSidDocumentVersion();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.docId = reader.string();
          break;
        case 2:
          message.versionList.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SidDocumentVersion {
    return {
      docId: isSet(object.docId) ? String(object.docId) : "",
      versionList: Array.isArray(object?.versionList) ? object.versionList.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: SidDocumentVersion): unknown {
    const obj: any = {};
    message.docId !== undefined && (obj.docId = message.docId);
    if (message.versionList) {
      obj.versionList = message.versionList.map((e) => e);
    } else {
      obj.versionList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SidDocumentVersion>, I>>(object: I): SidDocumentVersion {
    const message = createBaseSidDocumentVersion();
    message.docId = object.docId ?? "";
    message.versionList = object.versionList?.map((e) => e) || [];
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
