/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface JwsSignature {
  protected: string;
  signature: string;
}

const baseJwsSignature: object = { protected: "", signature: "" };

export const JwsSignature = {
  encode(message: JwsSignature, writer: Writer = Writer.create()): Writer {
    if (message.protected !== "") {
      writer.uint32(10).string(message.protected);
    }
    if (message.signature !== "") {
      writer.uint32(18).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): JwsSignature {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseJwsSignature } as JwsSignature;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.protected = reader.string();
          break;
        case 2:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): JwsSignature {
    const message = { ...baseJwsSignature } as JwsSignature;
    if (object.protected !== undefined && object.protected !== null) {
      message.protected = String(object.protected);
    } else {
      message.protected = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: JwsSignature): unknown {
    const obj: any = {};
    message.protected !== undefined && (obj.protected = message.protected);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<JwsSignature>): JwsSignature {
    const message = { ...baseJwsSignature } as JwsSignature;
    if (object.protected !== undefined && object.protected !== null) {
      message.protected = object.protected;
    } else {
      message.protected = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
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
