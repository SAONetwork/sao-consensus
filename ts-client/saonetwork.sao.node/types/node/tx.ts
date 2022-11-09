/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

export interface MsgLogin {
  creator: string;
  peer: string;
}

export interface MsgLoginResponse {}

export interface MsgLogout {
  creator: string;
}

export interface MsgLogoutResponse {}

export interface MsgReset {
  creator: string;
  peer: string;
}

export interface MsgResetResponse {}

export interface MsgClaimReward {
  creator: string;
}

export interface MsgClaimRewardResponse {}

const baseMsgLogin: object = { creator: "", peer: "" };

export const MsgLogin = {
  encode(message: MsgLogin, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLogin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLogin } as MsgLogin;
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

  fromJSON(object: any): MsgLogin {
    const message = { ...baseMsgLogin } as MsgLogin;
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

  toJSON(message: MsgLogin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.peer !== undefined && (obj.peer = message.peer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLogin>): MsgLogin {
    const message = { ...baseMsgLogin } as MsgLogin;
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

const baseMsgLoginResponse: object = {};

export const MsgLoginResponse = {
  encode(_: MsgLoginResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLoginResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLoginResponse } as MsgLoginResponse;
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

  fromJSON(_: any): MsgLoginResponse {
    const message = { ...baseMsgLoginResponse } as MsgLoginResponse;
    return message;
  },

  toJSON(_: MsgLoginResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgLoginResponse>): MsgLoginResponse {
    const message = { ...baseMsgLoginResponse } as MsgLoginResponse;
    return message;
  },
};

const baseMsgLogout: object = { creator: "" };

export const MsgLogout = {
  encode(message: MsgLogout, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLogout {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLogout } as MsgLogout;
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

  fromJSON(object: any): MsgLogout {
    const message = { ...baseMsgLogout } as MsgLogout;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgLogout): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLogout>): MsgLogout {
    const message = { ...baseMsgLogout } as MsgLogout;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgLogoutResponse: object = {};

export const MsgLogoutResponse = {
  encode(_: MsgLogoutResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLogoutResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLogoutResponse } as MsgLogoutResponse;
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

  fromJSON(_: any): MsgLogoutResponse {
    const message = { ...baseMsgLogoutResponse } as MsgLogoutResponse;
    return message;
  },

  toJSON(_: MsgLogoutResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgLogoutResponse>): MsgLogoutResponse {
    const message = { ...baseMsgLogoutResponse } as MsgLogoutResponse;
    return message;
  },
};

const baseMsgReset: object = { creator: "", peer: "" };

export const MsgReset = {
  encode(message: MsgReset, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReset {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReset } as MsgReset;
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

  fromJSON(object: any): MsgReset {
    const message = { ...baseMsgReset } as MsgReset;
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

  toJSON(message: MsgReset): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.peer !== undefined && (obj.peer = message.peer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReset>): MsgReset {
    const message = { ...baseMsgReset } as MsgReset;
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

const baseMsgResetResponse: object = {};

export const MsgResetResponse = {
  encode(_: MsgResetResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgResetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgResetResponse } as MsgResetResponse;
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

  fromJSON(_: any): MsgResetResponse {
    const message = { ...baseMsgResetResponse } as MsgResetResponse;
    return message;
  },

  toJSON(_: MsgResetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgResetResponse>): MsgResetResponse {
    const message = { ...baseMsgResetResponse } as MsgResetResponse;
    return message;
  },
};

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
  Login(request: MsgLogin): Promise<MsgLoginResponse>;
  Logout(request: MsgLogout): Promise<MsgLogoutResponse>;
  Reset(request: MsgReset): Promise<MsgResetResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClaimReward(request: MsgClaimReward): Promise<MsgClaimRewardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Login(request: MsgLogin): Promise<MsgLoginResponse> {
    const data = MsgLogin.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Login", data);
    return promise.then((data) => MsgLoginResponse.decode(new Reader(data)));
  }

  Logout(request: MsgLogout): Promise<MsgLogoutResponse> {
    const data = MsgLogout.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Logout", data);
    return promise.then((data) => MsgLogoutResponse.decode(new Reader(data)));
  }

  Reset(request: MsgReset): Promise<MsgResetResponse> {
    const data = MsgReset.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Reset", data);
    return promise.then((data) => MsgResetResponse.decode(new Reader(data)));
  }

  ClaimReward(request: MsgClaimReward): Promise<MsgClaimRewardResponse> {
    const data = MsgClaimReward.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.node.Msg",
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
