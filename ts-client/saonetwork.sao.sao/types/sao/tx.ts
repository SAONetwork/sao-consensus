/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "saonetwork.sao.sao";

export interface MsgStore {
  creator: string;
  owner: string;
  cid: string;
  provider: string;
  duration: number;
  replica: number;
}

export interface MsgStoreResponse {
  orderId: number;
}

export interface MsgCancel {
  creator: string;
  orderId: number;
}

export interface MsgCancelResponse {}

export interface MsgComplete {
  creator: string;
  orderId: number;
  cid: string;
  size: number;
}

export interface MsgCompleteResponse {}

export interface MsgReject {
  creator: string;
  orderId: number;
}

export interface MsgRejectResponse {}

export interface MsgTerminate {
  creator: string;
  orderId: number;
}

export interface MsgTerminateResponse {}

export interface MsgReady {
  creator: string;
  orderId: number;
}

export interface MsgReadyResponse {}

const baseMsgStore: object = {
  creator: "",
  owner: "",
  cid: "",
  provider: "",
  duration: 0,
  replica: 0,
};

export const MsgStore = {
  encode(message: MsgStore, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    if (message.provider !== "") {
      writer.uint32(34).string(message.provider);
    }
    if (message.duration !== 0) {
      writer.uint32(40).int32(message.duration);
    }
    if (message.replica !== 0) {
      writer.uint32(48).int32(message.replica);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStore {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStore } as MsgStore;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.cid = reader.string();
          break;
        case 4:
          message.provider = reader.string();
          break;
        case 5:
          message.duration = reader.int32();
          break;
        case 6:
          message.replica = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStore {
    const message = { ...baseMsgStore } as MsgStore;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = String(object.provider);
    } else {
      message.provider = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = Number(object.replica);
    } else {
      message.replica = 0;
    }
    return message;
  },

  toJSON(message: MsgStore): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.owner !== undefined && (obj.owner = message.owner);
    message.cid !== undefined && (obj.cid = message.cid);
    message.provider !== undefined && (obj.provider = message.provider);
    message.duration !== undefined && (obj.duration = message.duration);
    message.replica !== undefined && (obj.replica = message.replica);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStore>): MsgStore {
    const message = { ...baseMsgStore } as MsgStore;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = object.provider;
    } else {
      message.provider = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.replica !== undefined && object.replica !== null) {
      message.replica = object.replica;
    } else {
      message.replica = 0;
    }
    return message;
  },
};

const baseMsgStoreResponse: object = { orderId: 0 };

export const MsgStoreResponse = {
  encode(message: MsgStoreResponse, writer: Writer = Writer.create()): Writer {
    if (message.orderId !== 0) {
      writer.uint32(8).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStoreResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStoreResponse } as MsgStoreResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreResponse {
    const message = { ...baseMsgStoreResponse } as MsgStoreResponse;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgStoreResponse): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStoreResponse>): MsgStoreResponse {
    const message = { ...baseMsgStoreResponse } as MsgStoreResponse;
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

const baseMsgCancel: object = { creator: "", orderId: 0 };

export const MsgCancel = {
  encode(message: MsgCancel, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancel } as MsgCancel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancel {
    const message = { ...baseMsgCancel } as MsgCancel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgCancel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancel>): MsgCancel {
    const message = { ...baseMsgCancel } as MsgCancel;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

const baseMsgCancelResponse: object = {};

export const MsgCancelResponse = {
  encode(_: MsgCancelResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelResponse } as MsgCancelResponse;
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

  fromJSON(_: any): MsgCancelResponse {
    const message = { ...baseMsgCancelResponse } as MsgCancelResponse;
    return message;
  },

  toJSON(_: MsgCancelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCancelResponse>): MsgCancelResponse {
    const message = { ...baseMsgCancelResponse } as MsgCancelResponse;
    return message;
  },
};

const baseMsgComplete: object = { creator: "", orderId: 0, cid: "", size: 0 };

export const MsgComplete = {
  encode(message: MsgComplete, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    if (message.size !== 0) {
      writer.uint32(32).int32(message.size);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgComplete {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgComplete } as MsgComplete;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.cid = reader.string();
          break;
        case 4:
          message.size = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgComplete {
    const message = { ...baseMsgComplete } as MsgComplete;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = Number(object.size);
    } else {
      message.size = 0;
    }
    return message;
  },

  toJSON(message: MsgComplete): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    message.cid !== undefined && (obj.cid = message.cid);
    message.size !== undefined && (obj.size = message.size);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgComplete>): MsgComplete {
    const message = { ...baseMsgComplete } as MsgComplete;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    } else {
      message.size = 0;
    }
    return message;
  },
};

const baseMsgCompleteResponse: object = {};

export const MsgCompleteResponse = {
  encode(_: MsgCompleteResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCompleteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCompleteResponse } as MsgCompleteResponse;
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

  fromJSON(_: any): MsgCompleteResponse {
    const message = { ...baseMsgCompleteResponse } as MsgCompleteResponse;
    return message;
  },

  toJSON(_: MsgCompleteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCompleteResponse>): MsgCompleteResponse {
    const message = { ...baseMsgCompleteResponse } as MsgCompleteResponse;
    return message;
  },
};

const baseMsgReject: object = { creator: "", orderId: 0 };

export const MsgReject = {
  encode(message: MsgReject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReject } as MsgReject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReject {
    const message = { ...baseMsgReject } as MsgReject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgReject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReject>): MsgReject {
    const message = { ...baseMsgReject } as MsgReject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

const baseMsgRejectResponse: object = {};

export const MsgRejectResponse = {
  encode(_: MsgRejectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRejectResponse } as MsgRejectResponse;
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

  fromJSON(_: any): MsgRejectResponse {
    const message = { ...baseMsgRejectResponse } as MsgRejectResponse;
    return message;
  },

  toJSON(_: MsgRejectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRejectResponse>): MsgRejectResponse {
    const message = { ...baseMsgRejectResponse } as MsgRejectResponse;
    return message;
  },
};

const baseMsgTerminate: object = { creator: "", orderId: 0 };

export const MsgTerminate = {
  encode(message: MsgTerminate, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTerminate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTerminate } as MsgTerminate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTerminate {
    const message = { ...baseMsgTerminate } as MsgTerminate;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgTerminate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTerminate>): MsgTerminate {
    const message = { ...baseMsgTerminate } as MsgTerminate;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

const baseMsgTerminateResponse: object = {};

export const MsgTerminateResponse = {
  encode(_: MsgTerminateResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTerminateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTerminateResponse } as MsgTerminateResponse;
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

  fromJSON(_: any): MsgTerminateResponse {
    const message = { ...baseMsgTerminateResponse } as MsgTerminateResponse;
    return message;
  },

  toJSON(_: MsgTerminateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgTerminateResponse>): MsgTerminateResponse {
    const message = { ...baseMsgTerminateResponse } as MsgTerminateResponse;
    return message;
  },
};

const baseMsgReady: object = { creator: "", orderId: 0 };

export const MsgReady = {
  encode(message: MsgReady, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReady {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReady } as MsgReady;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReady {
    const message = { ...baseMsgReady } as MsgReady;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    return message;
  },

  toJSON(message: MsgReady): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReady>): MsgReady {
    const message = { ...baseMsgReady } as MsgReady;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    return message;
  },
};

const baseMsgReadyResponse: object = {};

export const MsgReadyResponse = {
  encode(_: MsgReadyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReadyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReadyResponse } as MsgReadyResponse;
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

  fromJSON(_: any): MsgReadyResponse {
    const message = { ...baseMsgReadyResponse } as MsgReadyResponse;
    return message;
  },

  toJSON(_: MsgReadyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgReadyResponse>): MsgReadyResponse {
    const message = { ...baseMsgReadyResponse } as MsgReadyResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Store(request: MsgStore): Promise<MsgStoreResponse>;
  Cancel(request: MsgCancel): Promise<MsgCancelResponse>;
  Complete(request: MsgComplete): Promise<MsgCompleteResponse>;
  Reject(request: MsgReject): Promise<MsgRejectResponse>;
  Terminate(request: MsgTerminate): Promise<MsgTerminateResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Ready(request: MsgReady): Promise<MsgReadyResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Store(request: MsgStore): Promise<MsgStoreResponse> {
    const data = MsgStore.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Store", data);
    return promise.then((data) => MsgStoreResponse.decode(new Reader(data)));
  }

  Cancel(request: MsgCancel): Promise<MsgCancelResponse> {
    const data = MsgCancel.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Cancel", data);
    return promise.then((data) => MsgCancelResponse.decode(new Reader(data)));
  }

  Complete(request: MsgComplete): Promise<MsgCompleteResponse> {
    const data = MsgComplete.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.sao.Msg",
      "Complete",
      data
    );
    return promise.then((data) => MsgCompleteResponse.decode(new Reader(data)));
  }

  Reject(request: MsgReject): Promise<MsgRejectResponse> {
    const data = MsgReject.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Reject", data);
    return promise.then((data) => MsgRejectResponse.decode(new Reader(data)));
  }

  Terminate(request: MsgTerminate): Promise<MsgTerminateResponse> {
    const data = MsgTerminate.encode(request).finish();
    const promise = this.rpc.request(
      "saonetwork.sao.sao.Msg",
      "Terminate",
      data
    );
    return promise.then((data) =>
      MsgTerminateResponse.decode(new Reader(data))
    );
  }

  Ready(request: MsgReady): Promise<MsgReadyResponse> {
    const data = MsgReady.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Ready", data);
    return promise.then((data) => MsgReadyResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
