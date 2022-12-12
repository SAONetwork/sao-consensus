/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.did";

export interface AccountAuth {
  accountDid: string;
  accountEncryptedSeed: string;
  sidEncryptedAccount: string;
}

function createBaseAccountAuth(): AccountAuth {
  return { accountDid: "", accountEncryptedSeed: "", sidEncryptedAccount: "" };
}

export const AccountAuth = {
  encode(message: AccountAuth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountDid !== "") {
      writer.uint32(10).string(message.accountDid);
    }
    if (message.accountEncryptedSeed !== "") {
      writer.uint32(18).string(message.accountEncryptedSeed);
    }
    if (message.sidEncryptedAccount !== "") {
      writer.uint32(26).string(message.sidEncryptedAccount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccountAuth {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountAuth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountDid = reader.string();
          break;
        case 2:
          message.accountEncryptedSeed = reader.string();
          break;
        case 3:
          message.sidEncryptedAccount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountAuth {
    return {
      accountDid: isSet(object.accountDid) ? String(object.accountDid) : "",
      accountEncryptedSeed: isSet(object.accountEncryptedSeed) ? String(object.accountEncryptedSeed) : "",
      sidEncryptedAccount: isSet(object.sidEncryptedAccount) ? String(object.sidEncryptedAccount) : "",
    };
  },

  toJSON(message: AccountAuth): unknown {
    const obj: any = {};
    message.accountDid !== undefined && (obj.accountDid = message.accountDid);
    message.accountEncryptedSeed !== undefined && (obj.accountEncryptedSeed = message.accountEncryptedSeed);
    message.sidEncryptedAccount !== undefined && (obj.sidEncryptedAccount = message.sidEncryptedAccount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountAuth>, I>>(object: I): AccountAuth {
    const message = createBaseAccountAuth();
    message.accountDid = object.accountDid ?? "";
    message.accountEncryptedSeed = object.accountEncryptedSeed ?? "";
    message.sidEncryptedAccount = object.sidEncryptedAccount ?? "";
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
