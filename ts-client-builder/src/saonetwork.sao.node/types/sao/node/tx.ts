/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.node";

export interface MsgLogin {
  creator: string;
}

export interface MsgLoginResponse {
}

export interface MsgLogout {
  creator: string;
}

export interface MsgLogoutResponse {
}

export interface MsgReset {
  creator: string;
  peer: string;
  status: number;
}

export interface MsgResetResponse {
}

export interface MsgClaimReward {
  creator: string;
}

export interface MsgClaimRewardResponse {
}

function createBaseMsgLogin(): MsgLogin {
  return { creator: "" };
}

export const MsgLogin = {
  encode(message: MsgLogin, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLogin {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLogin();
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

  fromJSON(object: any): MsgLogin {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgLogin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLogin>, I>>(object: I): MsgLogin {
    const message = createBaseMsgLogin();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgLoginResponse(): MsgLoginResponse {
  return {};
}

export const MsgLoginResponse = {
  encode(_: MsgLoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLoginResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLoginResponse();
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
    return {};
  },

  toJSON(_: MsgLoginResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLoginResponse>, I>>(_: I): MsgLoginResponse {
    const message = createBaseMsgLoginResponse();
    return message;
  },
};

function createBaseMsgLogout(): MsgLogout {
  return { creator: "" };
}

export const MsgLogout = {
  encode(message: MsgLogout, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLogout {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLogout();
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
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgLogout): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLogout>, I>>(object: I): MsgLogout {
    const message = createBaseMsgLogout();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgLogoutResponse(): MsgLogoutResponse {
  return {};
}

export const MsgLogoutResponse = {
  encode(_: MsgLogoutResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLogoutResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLogoutResponse();
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
    return {};
  },

  toJSON(_: MsgLogoutResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLogoutResponse>, I>>(_: I): MsgLogoutResponse {
    const message = createBaseMsgLogoutResponse();
    return message;
  },
};

function createBaseMsgReset(): MsgReset {
  return { creator: "", peer: "", status: 0 };
}

export const MsgReset = {
  encode(message: MsgReset, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.peer !== "") {
      writer.uint32(18).string(message.peer);
    }
    if (message.status !== 0) {
      writer.uint32(24).uint32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReset {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReset();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.peer = reader.string();
          break;
        case 3:
          message.status = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReset {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      peer: isSet(object.peer) ? String(object.peer) : "",
      status: isSet(object.status) ? Number(object.status) : 0,
    };
  },

  toJSON(message: MsgReset): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.peer !== undefined && (obj.peer = message.peer);
    message.status !== undefined && (obj.status = Math.round(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReset>, I>>(object: I): MsgReset {
    const message = createBaseMsgReset();
    message.creator = object.creator ?? "";
    message.peer = object.peer ?? "";
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseMsgResetResponse(): MsgResetResponse {
  return {};
}

export const MsgResetResponse = {
  encode(_: MsgResetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgResetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgResetResponse();
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
    return {};
  },

  toJSON(_: MsgResetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgResetResponse>, I>>(_: I): MsgResetResponse {
    const message = createBaseMsgResetResponse();
    return message;
  },
};

function createBaseMsgClaimReward(): MsgClaimReward {
  return { creator: "" };
}

export const MsgClaimReward = {
  encode(message: MsgClaimReward, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimReward {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimReward();
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
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgClaimReward): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimReward>, I>>(object: I): MsgClaimReward {
    const message = createBaseMsgClaimReward();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgClaimRewardResponse(): MsgClaimRewardResponse {
  return {};
}

export const MsgClaimRewardResponse = {
  encode(_: MsgClaimRewardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimRewardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimRewardResponse();
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
    return {};
  },

  toJSON(_: MsgClaimRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimRewardResponse>, I>>(_: I): MsgClaimRewardResponse {
    const message = createBaseMsgClaimRewardResponse();
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
    this.Login = this.Login.bind(this);
    this.Logout = this.Logout.bind(this);
    this.Reset = this.Reset.bind(this);
    this.ClaimReward = this.ClaimReward.bind(this);
  }
  Login(request: MsgLogin): Promise<MsgLoginResponse> {
    const data = MsgLogin.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Login", data);
    return promise.then((data) => MsgLoginResponse.decode(new _m0.Reader(data)));
  }

  Logout(request: MsgLogout): Promise<MsgLogoutResponse> {
    const data = MsgLogout.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Logout", data);
    return promise.then((data) => MsgLogoutResponse.decode(new _m0.Reader(data)));
  }

  Reset(request: MsgReset): Promise<MsgResetResponse> {
    const data = MsgReset.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "Reset", data);
    return promise.then((data) => MsgResetResponse.decode(new _m0.Reader(data)));
  }

  ClaimReward(request: MsgClaimReward): Promise<MsgClaimRewardResponse> {
    const data = MsgClaimReward.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.node.Msg", "ClaimReward", data);
    return promise.then((data) => MsgClaimRewardResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
