/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface PermissionProposal {
  owner: string;
  dataId: string;
  readonlyDids: string[];
  readwriteDids: string[];
}

function createBasePermissionProposal(): PermissionProposal {
  return { owner: "", dataId: "", readonlyDids: [], readwriteDids: [] };
}

export const PermissionProposal = {
  encode(message: PermissionProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.dataId !== "") {
      writer.uint32(18).string(message.dataId);
    }
    for (const v of message.readonlyDids) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.readwriteDids) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PermissionProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePermissionProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.dataId = reader.string();
          break;
        case 3:
          message.readonlyDids.push(reader.string());
          break;
        case 4:
          message.readwriteDids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PermissionProposal {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      dataId: isSet(object.dataId) ? String(object.dataId) : "",
      readonlyDids: Array.isArray(object?.readonlyDids) ? object.readonlyDids.map((e: any) => String(e)) : [],
      readwriteDids: Array.isArray(object?.readwriteDids) ? object.readwriteDids.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: PermissionProposal): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.dataId !== undefined && (obj.dataId = message.dataId);
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

  fromPartial<I extends Exact<DeepPartial<PermissionProposal>, I>>(object: I): PermissionProposal {
    const message = createBasePermissionProposal();
    message.owner = object.owner ?? "";
    message.dataId = object.dataId ?? "";
    message.readonlyDids = object.readonlyDids?.map((e) => e) || [];
    message.readwriteDids = object.readwriteDids?.map((e) => e) || [];
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
