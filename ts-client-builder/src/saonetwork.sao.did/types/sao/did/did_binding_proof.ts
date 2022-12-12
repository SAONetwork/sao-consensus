/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { BindingProof } from "./binding_proof";

export const protobufPackage = "saonetwork.sao.did";

export interface DidBindingProof {
  accountId: string;
  proof: BindingProof | undefined;
}

function createBaseDidBindingProof(): DidBindingProof {
  return { accountId: "", proof: undefined };
}

export const DidBindingProof = {
  encode(message: DidBindingProof, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.proof !== undefined) {
      BindingProof.encode(message.proof, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DidBindingProof {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDidBindingProof();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        case 2:
          message.proof = BindingProof.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidBindingProof {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      proof: isSet(object.proof) ? BindingProof.fromJSON(object.proof) : undefined,
    };
  },

  toJSON(message: DidBindingProof): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.proof !== undefined && (obj.proof = message.proof ? BindingProof.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DidBindingProof>, I>>(object: I): DidBindingProof {
    const message = createBaseDidBindingProof();
    message.accountId = object.accountId ?? "";
    message.proof = (object.proof !== undefined && object.proof !== null)
      ? BindingProof.fromPartial(object.proof)
      : undefined;
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
