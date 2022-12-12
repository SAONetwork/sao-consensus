/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface SaoJwsSignature {
  protected?: string;
  signature?: string;
}

export interface SaoMetadata {
  dataId?: string;
  owner?: string;
  alias?: string;
  groupId?: string;

  /** @format uint64 */
  orderId?: string;
  tags?: string[];
  cid?: string;
  commits?: string[];
  extendInfo?: string;
  update?: boolean;
  commit?: string;
  rule?: string;

  /** @format uint64 */
  duration?: string;

  /** @format uint64 */
  createdAt?: string;
  provider?: string;

  /** @format int32 */
  expire?: number;

  /** @format int32 */
  status?: number;

  /** @format int32 */
  replica?: number;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  amount?: V1Beta1Coin;

  /** @format uint64 */
  size?: string;

  /** @format int64 */
  operation?: number;
}

export type SaoMsgCancelResponse = object;

export type SaoMsgCompleteResponse = object;

export type SaoMsgReadyResponse = object;

export type SaoMsgRejectResponse = object;

export interface SaoMsgRenewResponse {
  result?: Record<string, string>;
}

export interface SaoMsgStoreResponse {
  /** @format uint64 */
  orderId?: string;
}

export type SaoMsgTerminateResponse = object;

export type SaoMsgUpdataPermissionResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type SaoParams = object;

export interface SaoPermissionProposal {
  owner?: string;
  dataId?: string;
  readonlyDids?: string[];
  readwriteDids?: string[];
}

export interface SaoProposal {
  owner?: string;
  provider?: string;
  groupId?: string;

  /** @format uint64 */
  duration?: string;

  /** @format int32 */
  replica?: number;

  /** @format int32 */
  timeout?: number;
  alias?: string;
  dataId?: string;
  commitId?: string;
  tags?: string[];
  cid?: string;
  rule?: string;
  extendInfo?: string;

  /** @format uint64 */
  size?: string;

  /**
   * 0: new|update, 1:force-push
   * @format int64
   */
  operation?: number;
  readonlyDids?: string[];
  readwriteDids?: string[];
}

export interface SaoQueryMetadataResponse {
  metadata?: SaoMetadata;
  shards?: Record<string, SaoShardMeta>;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface SaoQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: SaoParams;
}

export interface SaoQueryProposal {
  owner?: string;
  keyword?: string;
  groupId?: string;

  /**
   * 0,1 - query by dataId, 2 - query by alias
   * @format int64
   */
  type_?: number;

  /** @format uint64 */
  lastValidHeight?: string;
  gateway?: string;
  commitId?: string;
  version?: string;
}

export interface SaoRenewProposal {
  owner?: string;

  /** @format uint64 */
  duration?: string;

  /** @format int32 */
  timeout?: number;
  data?: string[];
}

export interface SaoShardMeta {
  /** @format uint64 */
  shardId?: string;
  peer?: string;
  cid?: string;
  provider?: string;
}

export interface SaoTerminateProposal {
  owner?: string;
  dataId?: string;
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title sao/sao/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryMetadata
   * @summary Queries a list of Metadata items.
   * @request GET:/SaoNetwork/sao/sao/metadata
   */
  queryMetadata = (
    query?: {
      "proposal.owner"?: string;
      "proposal.keyword"?: string;
      "proposal.groupId"?: string;
      "proposal.type_"?: number;
      "proposal.lastValidHeight"?: string;
      "proposal.gateway"?: string;
      "proposal.commitId"?: string;
      "proposal.version"?: string;
      "jws_signature.protected"?: string;
      "jws_signature.signature"?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<SaoQueryMetadataResponse, RpcStatus>({
      path: `/SaoNetwork/sao/sao/metadata`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/SaoNetwork/sao/sao/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<SaoQueryParamsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/sao/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
