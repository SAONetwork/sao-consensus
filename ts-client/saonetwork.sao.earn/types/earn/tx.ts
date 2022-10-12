/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.earn";

export interface MsgClaimReward {
  creator: string;
}

export interface MsgClaimRewardResponse {}

const baseMsgClaimReward: object = { creator: "" };

export const MsgClaimReward = {
  encode(message: MsgClaimReward, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimReward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimReward } as MsgClaimReward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimReward {
    const message = { ...baseMsgClaimReward } as MsgClaimReward;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgClaimReward): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgClaimReward>): MsgClaimReward {
    const message = { ...baseMsgClaimReward } as MsgClaimReward;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgClaimRewardResponse: object = {};

export const MsgClaimRewardResponse = {
  encode(_: MsgClaimRewardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimRewardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimRewardResponse } as MsgClaimRewardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimRewardResponse {
    const message = { ...baseMsgClaimRewardResponse } as MsgClaimRewardResponse;
    return message;
  },

  toJSON(_: MsgClaimRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgClaimRewardResponse>): MsgClaimRewardResponse {
    const message = { ...baseMsgClaimRewardResponse } as MsgClaimRewardResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClaimReward(request: MsgClaimReward): Promise<MsgClaimRewardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ClaimReward(request: MsgClaimReward): Promise<MsgClaimRewardResponse> {
    const data = MsgClaimReward.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.earn.Msg",
      "ClaimReward",
      data
    );
    return promise.then((data) =>
      MsgClaimRewardResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
