/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

export interface Node {
  creator: string;
  peer: string;
}

const baseNode: object = { creator: "", peer: "" };

export const Node = {
  encode(message: Node, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Node {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNode } as Node;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.peer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Node {
    const message = { ...baseNode } as Node;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = String(object.peer);
    } else {
      message.peer = "";
    }
    return message;
  },

  toJSON(message: Node): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.peer !== undefined && (obj.peer = message.peer);
    return obj;
  },

  fromPartial(object: DeepPartial<Node>): Node {
    const message = { ...baseNode } as Node;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = object.peer;
    } else {
      message.peer = "";
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
