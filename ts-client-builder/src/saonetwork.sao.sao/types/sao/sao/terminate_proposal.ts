/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface TerminateProposal {
  owner: string;
  dataId: string;
}

function createBaseTerminateProposal(): TerminateProposal {
  return { owner: "", dataId: "" };
}

export const TerminateProposal = {
  encode(message: TerminateProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.dataId !== "") {
      writer.uint32(18).string(message.dataId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TerminateProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTerminateProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.dataId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TerminateProposal {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      dataId: isSet(object.dataId) ? String(object.dataId) : "",
    };
  },

  toJSON(message: TerminateProposal): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.dataId !== undefined && (obj.dataId = message.dataId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TerminateProposal>, I>>(object: I): TerminateProposal {
    const message = createBaseTerminateProposal();
    message.owner = object.owner ?? "";
    message.dataId = object.dataId ?? "";
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
