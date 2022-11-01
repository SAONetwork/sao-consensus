/* eslint-disable */
import { Params } from "../model/params";
import { Metadata } from "../model/metadata";
import { Model } from "../model/model";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.model";

/** GenesisState defines the model module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  metadataList: Metadata[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  modelList: Model[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.metadataList) {
      Metadata.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.modelList) {
      Model.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.metadataList = [];
    message.modelList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.metadataList.push(Metadata.decode(reader, reader.uint32()));
          break;
        case 3:
          message.modelList.push(Model.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.metadataList = [];
    message.modelList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.metadataList !== undefined && object.metadataList !== null) {
      for (const e of object.metadataList) {
        message.metadataList.push(Metadata.fromJSON(e));
      }
    }
    if (object.modelList !== undefined && object.modelList !== null) {
      for (const e of object.modelList) {
        message.modelList.push(Model.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.metadataList) {
      obj.metadataList = message.metadataList.map((e) =>
        e ? Metadata.toJSON(e) : undefined
      );
    } else {
      obj.metadataList = [];
    }
    if (message.modelList) {
      obj.modelList = message.modelList.map((e) =>
        e ? Model.toJSON(e) : undefined
      );
    } else {
      obj.modelList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.metadataList = [];
    message.modelList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.metadataList !== undefined && object.metadataList !== null) {
      for (const e of object.metadataList) {
        message.metadataList.push(Metadata.fromPartial(e));
      }
    }
    if (object.modelList !== undefined && object.modelList !== null) {
      for (const e of object.modelList) {
        message.modelList.push(Model.fromPartial(e));
      }
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
