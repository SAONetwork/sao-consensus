/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { AccountAuth } from "./account_auth";
import { BindingProof } from "./binding_proof";
import { PubKey } from "./pub_key";

export const protobufPackage = "saonetwork.sao.did";

export interface MsgAddBinding {
  creator: string;
  accountId: string;
  proof: BindingProof | undefined;
}

export interface MsgAddBindingResponse {
}

export interface MsgUnbinding {
  creator: string;
  accountId: string;
}

export interface MsgUnbindingResponse {
}

export interface MsgAddAccountAuth {
  creator: string;
  did: string;
  accountAuth: AccountAuth | undefined;
}

export interface MsgAddAccountAuthResponse {
}

export interface MsgUpdateAccountAuths {
  creator: string;
  did: string;
  update: AccountAuth[];
  remove: string[];
}

export interface MsgUpdateAccountAuthsResponse {
}

export interface MsgUpdateSidDocument {
  creator: string;
  keys: PubKey[];
  rootDocId: string;
}

export interface MsgUpdateSidDocumentResponse {
  docId: string;
}

export interface MsgAddPastSeed {
  creator: string;
  did: string;
  pastSeed: string;
}

export interface MsgAddPastSeedResponse {
}

export interface MsgResetStore {
  creator: string;
}

export interface MsgResetStoreResponse {
}

export interface MsgUpdatePaymentAddress {
  creator: string;
  accountId: string;
  did: string;
}

export interface MsgUpdatePaymentAddressResponse {
}

export interface MsgBinding {
  creator: string;
  accountId: string;
  rootDocId: string;
  keys: PubKey[];
  accountAuth: AccountAuth | undefined;
  proof: BindingProof | undefined;
}

export interface MsgBindingResponse {
}

function createBaseMsgAddBinding(): MsgAddBinding {
  return { creator: "", accountId: "", proof: undefined };
}

export const MsgAddBinding = {
  encode(message: MsgAddBinding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountId !== "") {
      writer.uint32(18).string(message.accountId);
    }
    if (message.proof !== undefined) {
      BindingProof.encode(message.proof, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBinding {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBinding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.accountId = reader.string();
          break;
        case 3:
          message.proof = BindingProof.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddBinding {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      proof: isSet(object.proof) ? BindingProof.fromJSON(object.proof) : undefined,
    };
  },

  toJSON(message: MsgAddBinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.proof !== undefined && (obj.proof = message.proof ? BindingProof.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBinding>, I>>(object: I): MsgAddBinding {
    const message = createBaseMsgAddBinding();
    message.creator = object.creator ?? "";
    message.accountId = object.accountId ?? "";
    message.proof = (object.proof !== undefined && object.proof !== null)
      ? BindingProof.fromPartial(object.proof)
      : undefined;
    return message;
  },
};

function createBaseMsgAddBindingResponse(): MsgAddBindingResponse {
  return {};
}

export const MsgAddBindingResponse = {
  encode(_: MsgAddBindingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBindingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBindingResponse();
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

  fromJSON(_: any): MsgAddBindingResponse {
    return {};
  },

  toJSON(_: MsgAddBindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBindingResponse>, I>>(_: I): MsgAddBindingResponse {
    const message = createBaseMsgAddBindingResponse();
    return message;
  },
};

function createBaseMsgUnbinding(): MsgUnbinding {
  return { creator: "", accountId: "" };
}

export const MsgUnbinding = {
  encode(message: MsgUnbinding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountId !== "") {
      writer.uint32(18).string(message.accountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnbinding {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnbinding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.accountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnbinding {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
    };
  },

  toJSON(message: MsgUnbinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accountId !== undefined && (obj.accountId = message.accountId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnbinding>, I>>(object: I): MsgUnbinding {
    const message = createBaseMsgUnbinding();
    message.creator = object.creator ?? "";
    message.accountId = object.accountId ?? "";
    return message;
  },
};

function createBaseMsgUnbindingResponse(): MsgUnbindingResponse {
  return {};
}

export const MsgUnbindingResponse = {
  encode(_: MsgUnbindingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnbindingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnbindingResponse();
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

  fromJSON(_: any): MsgUnbindingResponse {
    return {};
  },

  toJSON(_: MsgUnbindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnbindingResponse>, I>>(_: I): MsgUnbindingResponse {
    const message = createBaseMsgUnbindingResponse();
    return message;
  },
};

function createBaseMsgAddAccountAuth(): MsgAddAccountAuth {
  return { creator: "", did: "", accountAuth: undefined };
}

export const MsgAddAccountAuth = {
  encode(message: MsgAddAccountAuth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.accountAuth !== undefined) {
      AccountAuth.encode(message.accountAuth, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAccountAuth {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAccountAuth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.accountAuth = AccountAuth.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAccountAuth {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      accountAuth: isSet(object.accountAuth) ? AccountAuth.fromJSON(object.accountAuth) : undefined,
    };
  },

  toJSON(message: MsgAddAccountAuth): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.accountAuth !== undefined
      && (obj.accountAuth = message.accountAuth ? AccountAuth.toJSON(message.accountAuth) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAccountAuth>, I>>(object: I): MsgAddAccountAuth {
    const message = createBaseMsgAddAccountAuth();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.accountAuth = (object.accountAuth !== undefined && object.accountAuth !== null)
      ? AccountAuth.fromPartial(object.accountAuth)
      : undefined;
    return message;
  },
};

function createBaseMsgAddAccountAuthResponse(): MsgAddAccountAuthResponse {
  return {};
}

export const MsgAddAccountAuthResponse = {
  encode(_: MsgAddAccountAuthResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAccountAuthResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAccountAuthResponse();
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

  fromJSON(_: any): MsgAddAccountAuthResponse {
    return {};
  },

  toJSON(_: MsgAddAccountAuthResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAccountAuthResponse>, I>>(_: I): MsgAddAccountAuthResponse {
    const message = createBaseMsgAddAccountAuthResponse();
    return message;
  },
};

function createBaseMsgUpdateAccountAuths(): MsgUpdateAccountAuths {
  return { creator: "", did: "", update: [], remove: [] };
}

export const MsgUpdateAccountAuths = {
  encode(message: MsgUpdateAccountAuths, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    for (const v of message.update) {
      AccountAuth.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.remove) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAccountAuths {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAccountAuths();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.update.push(AccountAuth.decode(reader, reader.uint32()));
          break;
        case 4:
          message.remove.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateAccountAuths {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      update: Array.isArray(object?.update) ? object.update.map((e: any) => AccountAuth.fromJSON(e)) : [],
      remove: Array.isArray(object?.remove) ? object.remove.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgUpdateAccountAuths): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    if (message.update) {
      obj.update = message.update.map((e) => e ? AccountAuth.toJSON(e) : undefined);
    } else {
      obj.update = [];
    }
    if (message.remove) {
      obj.remove = message.remove.map((e) => e);
    } else {
      obj.remove = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAccountAuths>, I>>(object: I): MsgUpdateAccountAuths {
    const message = createBaseMsgUpdateAccountAuths();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.update = object.update?.map((e) => AccountAuth.fromPartial(e)) || [];
    message.remove = object.remove?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgUpdateAccountAuthsResponse(): MsgUpdateAccountAuthsResponse {
  return {};
}

export const MsgUpdateAccountAuthsResponse = {
  encode(_: MsgUpdateAccountAuthsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateAccountAuthsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateAccountAuthsResponse();
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

  fromJSON(_: any): MsgUpdateAccountAuthsResponse {
    return {};
  },

  toJSON(_: MsgUpdateAccountAuthsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateAccountAuthsResponse>, I>>(_: I): MsgUpdateAccountAuthsResponse {
    const message = createBaseMsgUpdateAccountAuthsResponse();
    return message;
  },
};

function createBaseMsgUpdateSidDocument(): MsgUpdateSidDocument {
  return { creator: "", keys: [], rootDocId: "" };
}

export const MsgUpdateSidDocument = {
  encode(message: MsgUpdateSidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.keys) {
      PubKey.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.rootDocId !== "") {
      writer.uint32(26).string(message.rootDocId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.keys.push(PubKey.decode(reader, reader.uint32()));
          break;
        case 3:
          message.rootDocId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSidDocument {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      keys: Array.isArray(object?.keys) ? object.keys.map((e: any) => PubKey.fromJSON(e)) : [],
      rootDocId: isSet(object.rootDocId) ? String(object.rootDocId) : "",
    };
  },

  toJSON(message: MsgUpdateSidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.keys) {
      obj.keys = message.keys.map((e) => e ? PubKey.toJSON(e) : undefined);
    } else {
      obj.keys = [];
    }
    message.rootDocId !== undefined && (obj.rootDocId = message.rootDocId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSidDocument>, I>>(object: I): MsgUpdateSidDocument {
    const message = createBaseMsgUpdateSidDocument();
    message.creator = object.creator ?? "";
    message.keys = object.keys?.map((e) => PubKey.fromPartial(e)) || [];
    message.rootDocId = object.rootDocId ?? "";
    return message;
  },
};

function createBaseMsgUpdateSidDocumentResponse(): MsgUpdateSidDocumentResponse {
  return { docId: "" };
}

export const MsgUpdateSidDocumentResponse = {
  encode(message: MsgUpdateSidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.docId !== "") {
      writer.uint32(10).string(message.docId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.docId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSidDocumentResponse {
    return { docId: isSet(object.docId) ? String(object.docId) : "" };
  },

  toJSON(message: MsgUpdateSidDocumentResponse): unknown {
    const obj: any = {};
    message.docId !== undefined && (obj.docId = message.docId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSidDocumentResponse>, I>>(object: I): MsgUpdateSidDocumentResponse {
    const message = createBaseMsgUpdateSidDocumentResponse();
    message.docId = object.docId ?? "";
    return message;
  },
};

function createBaseMsgAddPastSeed(): MsgAddPastSeed {
  return { creator: "", did: "", pastSeed: "" };
}

export const MsgAddPastSeed = {
  encode(message: MsgAddPastSeed, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.pastSeed !== "") {
      writer.uint32(26).string(message.pastSeed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddPastSeed {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddPastSeed();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.pastSeed = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddPastSeed {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      pastSeed: isSet(object.pastSeed) ? String(object.pastSeed) : "",
    };
  },

  toJSON(message: MsgAddPastSeed): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.pastSeed !== undefined && (obj.pastSeed = message.pastSeed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddPastSeed>, I>>(object: I): MsgAddPastSeed {
    const message = createBaseMsgAddPastSeed();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.pastSeed = object.pastSeed ?? "";
    return message;
  },
};

function createBaseMsgAddPastSeedResponse(): MsgAddPastSeedResponse {
  return {};
}

export const MsgAddPastSeedResponse = {
  encode(_: MsgAddPastSeedResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddPastSeedResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddPastSeedResponse();
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

  fromJSON(_: any): MsgAddPastSeedResponse {
    return {};
  },

  toJSON(_: MsgAddPastSeedResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddPastSeedResponse>, I>>(_: I): MsgAddPastSeedResponse {
    const message = createBaseMsgAddPastSeedResponse();
    return message;
  },
};

function createBaseMsgResetStore(): MsgResetStore {
  return { creator: "" };
}

export const MsgResetStore = {
  encode(message: MsgResetStore, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgResetStore {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgResetStore();
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

  fromJSON(object: any): MsgResetStore {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgResetStore): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgResetStore>, I>>(object: I): MsgResetStore {
    const message = createBaseMsgResetStore();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgResetStoreResponse(): MsgResetStoreResponse {
  return {};
}

export const MsgResetStoreResponse = {
  encode(_: MsgResetStoreResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgResetStoreResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgResetStoreResponse();
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

  fromJSON(_: any): MsgResetStoreResponse {
    return {};
  },

  toJSON(_: MsgResetStoreResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgResetStoreResponse>, I>>(_: I): MsgResetStoreResponse {
    const message = createBaseMsgResetStoreResponse();
    return message;
  },
};

function createBaseMsgUpdatePaymentAddress(): MsgUpdatePaymentAddress {
  return { creator: "", accountId: "", did: "" };
}

export const MsgUpdatePaymentAddress = {
  encode(message: MsgUpdatePaymentAddress, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountId !== "") {
      writer.uint32(18).string(message.accountId);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePaymentAddress {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePaymentAddress();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.accountId = reader.string();
          break;
        case 3:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdatePaymentAddress {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      did: isSet(object.did) ? String(object.did) : "",
    };
  },

  toJSON(message: MsgUpdatePaymentAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePaymentAddress>, I>>(object: I): MsgUpdatePaymentAddress {
    const message = createBaseMsgUpdatePaymentAddress();
    message.creator = object.creator ?? "";
    message.accountId = object.accountId ?? "";
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseMsgUpdatePaymentAddressResponse(): MsgUpdatePaymentAddressResponse {
  return {};
}

export const MsgUpdatePaymentAddressResponse = {
  encode(_: MsgUpdatePaymentAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePaymentAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePaymentAddressResponse();
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

  fromJSON(_: any): MsgUpdatePaymentAddressResponse {
    return {};
  },

  toJSON(_: MsgUpdatePaymentAddressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdatePaymentAddressResponse>, I>>(_: I): MsgUpdatePaymentAddressResponse {
    const message = createBaseMsgUpdatePaymentAddressResponse();
    return message;
  },
};

function createBaseMsgBinding(): MsgBinding {
  return { creator: "", accountId: "", rootDocId: "", keys: [], accountAuth: undefined, proof: undefined };
}

export const MsgBinding = {
  encode(message: MsgBinding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accountId !== "") {
      writer.uint32(18).string(message.accountId);
    }
    if (message.rootDocId !== "") {
      writer.uint32(26).string(message.rootDocId);
    }
    for (const v of message.keys) {
      PubKey.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.accountAuth !== undefined) {
      AccountAuth.encode(message.accountAuth, writer.uint32(42).fork()).ldelim();
    }
    if (message.proof !== undefined) {
      BindingProof.encode(message.proof, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBinding {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBinding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.accountId = reader.string();
          break;
        case 3:
          message.rootDocId = reader.string();
          break;
        case 4:
          message.keys.push(PubKey.decode(reader, reader.uint32()));
          break;
        case 5:
          message.accountAuth = AccountAuth.decode(reader, reader.uint32());
          break;
        case 6:
          message.proof = BindingProof.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBinding {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      rootDocId: isSet(object.rootDocId) ? String(object.rootDocId) : "",
      keys: Array.isArray(object?.keys) ? object.keys.map((e: any) => PubKey.fromJSON(e)) : [],
      accountAuth: isSet(object.accountAuth) ? AccountAuth.fromJSON(object.accountAuth) : undefined,
      proof: isSet(object.proof) ? BindingProof.fromJSON(object.proof) : undefined,
    };
  },

  toJSON(message: MsgBinding): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.rootDocId !== undefined && (obj.rootDocId = message.rootDocId);
    if (message.keys) {
      obj.keys = message.keys.map((e) => e ? PubKey.toJSON(e) : undefined);
    } else {
      obj.keys = [];
    }
    message.accountAuth !== undefined
      && (obj.accountAuth = message.accountAuth ? AccountAuth.toJSON(message.accountAuth) : undefined);
    message.proof !== undefined && (obj.proof = message.proof ? BindingProof.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBinding>, I>>(object: I): MsgBinding {
    const message = createBaseMsgBinding();
    message.creator = object.creator ?? "";
    message.accountId = object.accountId ?? "";
    message.rootDocId = object.rootDocId ?? "";
    message.keys = object.keys?.map((e) => PubKey.fromPartial(e)) || [];
    message.accountAuth = (object.accountAuth !== undefined && object.accountAuth !== null)
      ? AccountAuth.fromPartial(object.accountAuth)
      : undefined;
    message.proof = (object.proof !== undefined && object.proof !== null)
      ? BindingProof.fromPartial(object.proof)
      : undefined;
    return message;
  },
};

function createBaseMsgBindingResponse(): MsgBindingResponse {
  return {};
}

export const MsgBindingResponse = {
  encode(_: MsgBindingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBindingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBindingResponse();
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

  fromJSON(_: any): MsgBindingResponse {
    return {};
  },

  toJSON(_: MsgBindingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBindingResponse>, I>>(_: I): MsgBindingResponse {
    const message = createBaseMsgBindingResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddBinding(request: MsgAddBinding): Promise<MsgAddBindingResponse>;
  Unbinding(request: MsgUnbinding): Promise<MsgUnbindingResponse>;
  AddAccountAuth(request: MsgAddAccountAuth): Promise<MsgAddAccountAuthResponse>;
  UpdateAccountAuths(request: MsgUpdateAccountAuths): Promise<MsgUpdateAccountAuthsResponse>;
  UpdateSidDocument(request: MsgUpdateSidDocument): Promise<MsgUpdateSidDocumentResponse>;
  AddPastSeed(request: MsgAddPastSeed): Promise<MsgAddPastSeedResponse>;
  ResetStore(request: MsgResetStore): Promise<MsgResetStoreResponse>;
  UpdatePaymentAddress(request: MsgUpdatePaymentAddress): Promise<MsgUpdatePaymentAddressResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Binding(request: MsgBinding): Promise<MsgBindingResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddBinding = this.AddBinding.bind(this);
    this.Unbinding = this.Unbinding.bind(this);
    this.AddAccountAuth = this.AddAccountAuth.bind(this);
    this.UpdateAccountAuths = this.UpdateAccountAuths.bind(this);
    this.UpdateSidDocument = this.UpdateSidDocument.bind(this);
    this.AddPastSeed = this.AddPastSeed.bind(this);
    this.ResetStore = this.ResetStore.bind(this);
    this.UpdatePaymentAddress = this.UpdatePaymentAddress.bind(this);
    this.Binding = this.Binding.bind(this);
  }
  AddBinding(request: MsgAddBinding): Promise<MsgAddBindingResponse> {
    const data = MsgAddBinding.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "AddBinding", data);
    return promise.then((data) => MsgAddBindingResponse.decode(new _m0.Reader(data)));
  }

  Unbinding(request: MsgUnbinding): Promise<MsgUnbindingResponse> {
    const data = MsgUnbinding.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "Unbinding", data);
    return promise.then((data) => MsgUnbindingResponse.decode(new _m0.Reader(data)));
  }

  AddAccountAuth(request: MsgAddAccountAuth): Promise<MsgAddAccountAuthResponse> {
    const data = MsgAddAccountAuth.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "AddAccountAuth", data);
    return promise.then((data) => MsgAddAccountAuthResponse.decode(new _m0.Reader(data)));
  }

  UpdateAccountAuths(request: MsgUpdateAccountAuths): Promise<MsgUpdateAccountAuthsResponse> {
    const data = MsgUpdateAccountAuths.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "UpdateAccountAuths", data);
    return promise.then((data) => MsgUpdateAccountAuthsResponse.decode(new _m0.Reader(data)));
  }

  UpdateSidDocument(request: MsgUpdateSidDocument): Promise<MsgUpdateSidDocumentResponse> {
    const data = MsgUpdateSidDocument.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "UpdateSidDocument", data);
    return promise.then((data) => MsgUpdateSidDocumentResponse.decode(new _m0.Reader(data)));
  }

  AddPastSeed(request: MsgAddPastSeed): Promise<MsgAddPastSeedResponse> {
    const data = MsgAddPastSeed.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "AddPastSeed", data);
    return promise.then((data) => MsgAddPastSeedResponse.decode(new _m0.Reader(data)));
  }

  ResetStore(request: MsgResetStore): Promise<MsgResetStoreResponse> {
    const data = MsgResetStore.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "ResetStore", data);
    return promise.then((data) => MsgResetStoreResponse.decode(new _m0.Reader(data)));
  }

  UpdatePaymentAddress(request: MsgUpdatePaymentAddress): Promise<MsgUpdatePaymentAddressResponse> {
    const data = MsgUpdatePaymentAddress.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "UpdatePaymentAddress", data);
    return promise.then((data) => MsgUpdatePaymentAddressResponse.decode(new _m0.Reader(data)));
  }

  Binding(request: MsgBinding): Promise<MsgBindingResponse> {
    const data = MsgBinding.encode(request).finish();
    const promise = this.rpc.request("saonetwork.sao.did.Msg", "Binding", data);
    return promise.then((data) => MsgBindingResponse.decode(new _m0.Reader(data)));
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
