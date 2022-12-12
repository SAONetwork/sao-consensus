/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.did";

export interface AccountList {
  did: string;
  accountDids: string[];
}

function createBaseAccountList(): AccountList {
  return { did: "", accountDids: [] };
}

export const AccountList = {
  encode(message: AccountList, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    for (const v of message.accountDids) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccountList {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountList();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.accountDids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountList {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      accountDids: Array.isArray(object?.accountDids) ? object.accountDids.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: AccountList): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    if (message.accountDids) {
      obj.accountDids = message.accountDids.map((e) => e);
    } else {
      obj.accountDids = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountList>, I>>(object: I): AccountList {
    const message = createBaseAccountList();
    message.did = object.did ?? "";
    message.accountDids = object.accountDids?.map((e) => e) || [];
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
