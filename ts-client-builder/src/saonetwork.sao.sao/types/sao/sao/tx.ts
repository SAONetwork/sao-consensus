/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { JwsSignature } from "./jws_signature";
import { PermissionProposal } from "./permission_proposal";
import { Proposal } from "./proposal";
import { RenewProposal } from "./renew_proposal";
import { TerminateProposal } from "./terminate_proposal";

export const protobufPackage = "saonetwork.sao.sao";

export interface MsgCancel {
  creator: string;
  orderId: number;
}

export interface MsgCancelResponse {
}

export interface MsgComplete {
  creator: string;
  orderId: number;
  cid: string;
  size: number;
}

export interface MsgCompleteResponse {
}

export interface MsgReject {
  creator: string;
  orderId: number;
}

export interface MsgRejectResponse {
}

export interface MsgTerminate {
  creator: string;
  proposal: TerminateProposal | undefined;
  jwsSignature: JwsSignature | undefined;
}

export interface MsgTerminateResponse {
}

export interface MsgReady {
  creator: string;
  orderId: number;
}

export interface MsgReadyResponse {
}

export interface MsgStore {
  creator: string;
  proposal: Proposal | undefined;
  jwsSignature: JwsSignature | undefined;
}

export interface MsgStoreResponse {
  orderId: number;
}

export interface MsgRenew {
  creator: string;
  proposal: RenewProposal | undefined;
  jwsSignature: JwsSignature | undefined;
}

export interface MsgRenewResponse {
  result: { [key: string]: string };
}

export interface MsgRenewResponse_ResultEntry {
  key: string;
  value: string;
}

export interface MsgUpdataPermission {
  creator: string;
  proposal: PermissionProposal | undefined;
  jwsSignature: JwsSignature | undefined;
}

export interface MsgUpdataPermissionResponse {
}

function createBaseMsgCancel(): MsgCancel {
  return { creator: "", orderId: 0 };
}

export const MsgCancel = {
  encode(message: MsgCancel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancel {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancel();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
    };
  },

  toJSON(message: MsgCancel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancel>, I>>(object: I): MsgCancel {
    const message = createBaseMsgCancel();
    message.creator = object.creator ?? "";
    message.orderId = object.orderId ?? 0;
    return message;
  },
};

function createBaseMsgCancelResponse(): MsgCancelResponse {
  return {};
}

export const MsgCancelResponse = {
  encode(_: MsgCancelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelResponse();
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
    return {};
  },

  toJSON(_: MsgCancelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelResponse>, I>>(_: I): MsgCancelResponse {
    const message = createBaseMsgCancelResponse();
    return message;
  },
};

function createBaseMsgComplete(): MsgComplete {
  return { creator: "", orderId: 0, cid: "", size: 0 };
}

export const MsgComplete = {
  encode(message: MsgComplete, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgComplete {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgComplete();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
      cid: isSet(object.cid) ? String(object.cid) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
    };
  },

  toJSON(message: MsgComplete): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    message.cid !== undefined && (obj.cid = message.cid);
    message.size !== undefined && (obj.size = Math.round(message.size));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgComplete>, I>>(object: I): MsgComplete {
    const message = createBaseMsgComplete();
    message.creator = object.creator ?? "";
    message.orderId = object.orderId ?? 0;
    message.cid = object.cid ?? "";
    message.size = object.size ?? 0;
    return message;
  },
};

function createBaseMsgCompleteResponse(): MsgCompleteResponse {
  return {};
}

export const MsgCompleteResponse = {
  encode(_: MsgCompleteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCompleteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCompleteResponse();
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
    return {};
  },

  toJSON(_: MsgCompleteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCompleteResponse>, I>>(_: I): MsgCompleteResponse {
    const message = createBaseMsgCompleteResponse();
    return message;
  },
};

function createBaseMsgReject(): MsgReject {
  return { creator: "", orderId: 0 };
}

export const MsgReject = {
  encode(message: MsgReject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReject();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
    };
  },

  toJSON(message: MsgReject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReject>, I>>(object: I): MsgReject {
    const message = createBaseMsgReject();
    message.creator = object.creator ?? "";
    message.orderId = object.orderId ?? 0;
    return message;
  },
};

function createBaseMsgRejectResponse(): MsgRejectResponse {
  return {};
}

export const MsgRejectResponse = {
  encode(_: MsgRejectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectResponse();
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
    return {};
  },

  toJSON(_: MsgRejectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectResponse>, I>>(_: I): MsgRejectResponse {
    const message = createBaseMsgRejectResponse();
    return message;
  },
};

function createBaseMsgTerminate(): MsgTerminate {
  return { creator: "", proposal: undefined, jwsSignature: undefined };
}

export const MsgTerminate = {
  encode(message: MsgTerminate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.proposal !== undefined) {
      TerminateProposal.encode(message.proposal, writer.uint32(18).fork()).ldelim();
    }
    if (message.jwsSignature !== undefined) {
      JwsSignature.encode(message.jwsSignature, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTerminate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTerminate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.proposal = TerminateProposal.decode(reader, reader.uint32());
          break;
        case 3:
          message.jwsSignature = JwsSignature.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTerminate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      proposal: isSet(object.proposal) ? TerminateProposal.fromJSON(object.proposal) : undefined,
      jwsSignature: isSet(object.jwsSignature) ? JwsSignature.fromJSON(object.jwsSignature) : undefined,
    };
  },

  toJSON(message: MsgTerminate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.proposal !== undefined
      && (obj.proposal = message.proposal ? TerminateProposal.toJSON(message.proposal) : undefined);
    message.jwsSignature !== undefined
      && (obj.jwsSignature = message.jwsSignature ? JwsSignature.toJSON(message.jwsSignature) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTerminate>, I>>(object: I): MsgTerminate {
    const message = createBaseMsgTerminate();
    message.creator = object.creator ?? "";
    message.proposal = (object.proposal !== undefined && object.proposal !== null)
      ? TerminateProposal.fromPartial(object.proposal)
      : undefined;
    message.jwsSignature = (object.jwsSignature !== undefined && object.jwsSignature !== null)
      ? JwsSignature.fromPartial(object.jwsSignature)
      : undefined;
    return message;
  },
};

function createBaseMsgTerminateResponse(): MsgTerminateResponse {
  return {};
}

export const MsgTerminateResponse = {
  encode(_: MsgTerminateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTerminateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTerminateResponse();
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
    return {};
  },

  toJSON(_: MsgTerminateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTerminateResponse>, I>>(_: I): MsgTerminateResponse {
    const message = createBaseMsgTerminateResponse();
    return message;
  },
};

function createBaseMsgReady(): MsgReady {
  return { creator: "", orderId: 0 };
}

export const MsgReady = {
  encode(message: MsgReady, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReady {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReady();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      orderId: isSet(object.orderId) ? Number(object.orderId) : 0,
    };
  },

  toJSON(message: MsgReady): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReady>, I>>(object: I): MsgReady {
    const message = createBaseMsgReady();
    message.creator = object.creator ?? "";
    message.orderId = object.orderId ?? 0;
    return message;
  },
};

function createBaseMsgReadyResponse(): MsgReadyResponse {
  return {};
}

export const MsgReadyResponse = {
  encode(_: MsgReadyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReadyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReadyResponse();
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
    return {};
  },

  toJSON(_: MsgReadyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReadyResponse>, I>>(_: I): MsgReadyResponse {
    const message = createBaseMsgReadyResponse();
    return message;
  },
};

function createBaseMsgStore(): MsgStore {
  return { creator: "", proposal: undefined, jwsSignature: undefined };
}

export const MsgStore = {
  encode(message: MsgStore, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.proposal !== undefined) {
      Proposal.encode(message.proposal, writer.uint32(18).fork()).ldelim();
    }
    if (message.jwsSignature !== undefined) {
      JwsSignature.encode(message.jwsSignature, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStore {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStore();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.proposal = Proposal.decode(reader, reader.uint32());
          break;
        case 3:
          message.jwsSignature = JwsSignature.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStore {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      proposal: isSet(object.proposal) ? Proposal.fromJSON(object.proposal) : undefined,
      jwsSignature: isSet(object.jwsSignature) ? JwsSignature.fromJSON(object.jwsSignature) : undefined,
    };
  },

  toJSON(message: MsgStore): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.proposal !== undefined && (obj.proposal = message.proposal ? Proposal.toJSON(message.proposal) : undefined);
    message.jwsSignature !== undefined
      && (obj.jwsSignature = message.jwsSignature ? JwsSignature.toJSON(message.jwsSignature) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStore>, I>>(object: I): MsgStore {
    const message = createBaseMsgStore();
    message.creator = object.creator ?? "";
    message.proposal = (object.proposal !== undefined && object.proposal !== null)
      ? Proposal.fromPartial(object.proposal)
      : undefined;
    message.jwsSignature = (object.jwsSignature !== undefined && object.jwsSignature !== null)
      ? JwsSignature.fromPartial(object.jwsSignature)
      : undefined;
    return message;
  },
};

function createBaseMsgStoreResponse(): MsgStoreResponse {
  return { orderId: 0 };
}

export const MsgStoreResponse = {
  encode(message: MsgStoreResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.orderId !== 0) {
      writer.uint32(8).uint64(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStoreResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStoreResponse();
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
    return { orderId: isSet(object.orderId) ? Number(object.orderId) : 0 };
  },

  toJSON(message: MsgStoreResponse): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = Math.round(message.orderId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStoreResponse>, I>>(object: I): MsgStoreResponse {
    const message = createBaseMsgStoreResponse();
    message.orderId = object.orderId ?? 0;
    return message;
  },
};

function createBaseMsgRenew(): MsgRenew {
  return { creator: "", proposal: undefined, jwsSignature: undefined };
}

export const MsgRenew = {
  encode(message: MsgRenew, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.proposal !== undefined) {
      RenewProposal.encode(message.proposal, writer.uint32(18).fork()).ldelim();
    }
    if (message.jwsSignature !== undefined) {
      JwsSignature.encode(message.jwsSignature, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRenew {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRenew();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.proposal = RenewProposal.decode(reader, reader.uint32());
          break;
        case 3:
          message.jwsSignature = JwsSignature.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRenew {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      proposal: isSet(object.proposal) ? RenewProposal.fromJSON(object.proposal) : undefined,
      jwsSignature: isSet(object.jwsSignature) ? JwsSignature.fromJSON(object.jwsSignature) : undefined,
    };
  },

  toJSON(message: MsgRenew): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.proposal !== undefined
      && (obj.proposal = message.proposal ? RenewProposal.toJSON(message.proposal) : undefined);
    message.jwsSignature !== undefined
      && (obj.jwsSignature = message.jwsSignature ? JwsSignature.toJSON(message.jwsSignature) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRenew>, I>>(object: I): MsgRenew {
    const message = createBaseMsgRenew();
    message.creator = object.creator ?? "";
    message.proposal = (object.proposal !== undefined && object.proposal !== null)
      ? RenewProposal.fromPartial(object.proposal)
      : undefined;
    message.jwsSignature = (object.jwsSignature !== undefined && object.jwsSignature !== null)
      ? JwsSignature.fromPartial(object.jwsSignature)
      : undefined;
    return message;
  },
};

function createBaseMsgRenewResponse(): MsgRenewResponse {
  return { result: {} };
}

export const MsgRenewResponse = {
  encode(message: MsgRenewResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.result).forEach(([key, value]) => {
      MsgRenewResponse_ResultEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRenewResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRenewResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = MsgRenewResponse_ResultEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.result[entry1.key] = entry1.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRenewResponse {
    return {
      result: isObject(object.result)
        ? Object.entries(object.result).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: MsgRenewResponse): unknown {
    const obj: any = {};
    obj.result = {};
    if (message.result) {
      Object.entries(message.result).forEach(([k, v]) => {
        obj.result[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRenewResponse>, I>>(object: I): MsgRenewResponse {
    const message = createBaseMsgRenewResponse();
    message.result = Object.entries(object.result ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseMsgRenewResponse_ResultEntry(): MsgRenewResponse_ResultEntry {
  return { key: "", value: "" };
}

export const MsgRenewResponse_ResultEntry = {
  encode(message: MsgRenewResponse_ResultEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRenewResponse_ResultEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRenewResponse_ResultEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRenewResponse_ResultEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: MsgRenewResponse_ResultEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRenewResponse_ResultEntry>, I>>(object: I): MsgRenewResponse_ResultEntry {
    const message = createBaseMsgRenewResponse_ResultEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgUpdataPermission(): MsgUpdataPermission {
  return { creator: "", proposal: undefined, jwsSignature: undefined };
}

export const MsgUpdataPermission = {
  encode(message: MsgUpdataPermission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.proposal !== undefined) {
      PermissionProposal.encode(message.proposal, writer.uint32(18).fork()).ldelim();
    }
    if (message.jwsSignature !== undefined) {
      JwsSignature.encode(message.jwsSignature, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdataPermission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdataPermission();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.proposal = PermissionProposal.decode(reader, reader.uint32());
          break;
        case 3:
          message.jwsSignature = JwsSignature.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdataPermission {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      proposal: isSet(object.proposal) ? PermissionProposal.fromJSON(object.proposal) : undefined,
      jwsSignature: isSet(object.jwsSignature) ? JwsSignature.fromJSON(object.jwsSignature) : undefined,
    };
  },

  toJSON(message: MsgUpdataPermission): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.proposal !== undefined
      && (obj.proposal = message.proposal ? PermissionProposal.toJSON(message.proposal) : undefined);
    message.jwsSignature !== undefined
      && (obj.jwsSignature = message.jwsSignature ? JwsSignature.toJSON(message.jwsSignature) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdataPermission>, I>>(object: I): MsgUpdataPermission {
    const message = createBaseMsgUpdataPermission();
    message.creator = object.creator ?? "";
    message.proposal = (object.proposal !== undefined && object.proposal !== null)
      ? PermissionProposal.fromPartial(object.proposal)
      : undefined;
    message.jwsSignature = (object.jwsSignature !== undefined && object.jwsSignature !== null)
      ? JwsSignature.fromPartial(object.jwsSignature)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdataPermissionResponse(): MsgUpdataPermissionResponse {
  return {};
}

export const MsgUpdataPermissionResponse = {
  encode(_: MsgUpdataPermissionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdataPermissionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdataPermissionResponse();
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

  fromJSON(_: any): MsgUpdataPermissionResponse {
    return {};
  },

  toJSON(_: MsgUpdataPermissionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdataPermissionResponse>, I>>(_: I): MsgUpdataPermissionResponse {
    const message = createBaseMsgUpdataPermissionResponse();
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
  Ready(request: MsgReady): Promise<MsgReadyResponse>;
  Renew(request: MsgRenew): Promise<MsgRenewResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdataPermission(request: MsgUpdataPermission): Promise<MsgUpdataPermissionResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Store = this.Store.bind(this);
    this.Cancel = this.Cancel.bind(this);
    this.Complete = this.Complete.bind(this);
    this.Reject = this.Reject.bind(this);
    this.Terminate = this.Terminate.bind(this);
    this.Ready = this.Ready.bind(this);
    this.Renew = this.Renew.bind(this);
    this.UpdataPermission = this.UpdataPermission.bind(this);
  }
  Store(request: MsgStore): Promise<MsgStoreResponse> {
    const data = MsgStore.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Store", data);
    return promise.then((data) => MsgStoreResponse.decode(new _m0.Reader(data)));
  }

  Cancel(request: MsgCancel): Promise<MsgCancelResponse> {
    const data = MsgCancel.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Cancel", data);
    return promise.then((data) => MsgCancelResponse.decode(new _m0.Reader(data)));
  }

  Complete(request: MsgComplete): Promise<MsgCompleteResponse> {
    const data = MsgComplete.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Complete", data);
    return promise.then((data) => MsgCompleteResponse.decode(new _m0.Reader(data)));
  }

  Reject(request: MsgReject): Promise<MsgRejectResponse> {
    const data = MsgReject.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Reject", data);
    return promise.then((data) => MsgRejectResponse.decode(new _m0.Reader(data)));
  }

  Terminate(request: MsgTerminate): Promise<MsgTerminateResponse> {
    const data = MsgTerminate.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Terminate", data);
    return promise.then((data) => MsgTerminateResponse.decode(new _m0.Reader(data)));
  }

  Ready(request: MsgReady): Promise<MsgReadyResponse> {
    const data = MsgReady.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Ready", data);
    return promise.then((data) => MsgReadyResponse.decode(new _m0.Reader(data)));
  }

  Renew(request: MsgRenew): Promise<MsgRenewResponse> {
    const data = MsgRenew.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "Renew", data);
    return promise.then((data) => MsgRenewResponse.decode(new _m0.Reader(data)));
  }

  UpdataPermission(request: MsgUpdataPermission): Promise<MsgUpdataPermissionResponse> {
    const data = MsgUpdataPermission.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.sao.Msg", "UpdataPermission", data);
    return promise.then((data) => MsgUpdataPermissionResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
