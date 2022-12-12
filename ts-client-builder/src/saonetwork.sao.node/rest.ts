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

export type NodeMsgClaimRewardResponse = object;

export type NodeMsgLoginResponse = object;

export type NodeMsgLogoutResponse = object;

export type NodeMsgResetResponse = object;

export interface NodeNode {
  creator?: string;
  peer?: string;

  /** @format float */
  reputation?: number;

  /** @format int64 */
  status?: number;

  /** @format int64 */
  lastAliveHeigh?: string;
}

/**
 * Params defines the parameters for the module.
 */
export interface NodeParams {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  block_reward?: V1Beta1Coin;
}

export interface NodePledge {
  creator?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  total_order_pledged?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  total_storage_pledged?: V1Beta1Coin;

  /**
   * DecCoin defines a token with a denomination and a decimal amount.
   *
   * NOTE: The amount field is an Dec which implements the custom method
   * signatures required by gogoproto.
   */
  reward?: V1Beta1DecCoin;

  /**
   * DecCoin defines a token with a denomination and a decimal amount.
   *
   * NOTE: The amount field is an Dec which implements the custom method
   * signatures required by gogoproto.
   */
  reward_debt?: V1Beta1DecCoin;

  /** @format int64 */
  total_storage?: string;

  /** @format int64 */
  last_reward_at?: string;
}

export interface NodePool {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  total_pledged?: V1Beta1Coin;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  total_reward?: V1Beta1Coin;

  /**
   * DecCoin defines a token with a denomination and a decimal amount.
   *
   * NOTE: The amount field is an Dec which implements the custom method
   * signatures required by gogoproto.
   */
  acc_pledge_per_byte?: V1Beta1DecCoin;

  /**
   * DecCoin defines a token with a denomination and a decimal amount.
   *
   * NOTE: The amount field is an Dec which implements the custom method
   * signatures required by gogoproto.
   */
  acc_reward_per_byte?: V1Beta1DecCoin;

  /**
   * DecCoin defines a token with a denomination and a decimal amount.
   *
   * NOTE: The amount field is an Dec which implements the custom method
   * signatures required by gogoproto.
   */
  reward_per_block?: V1Beta1DecCoin;

  /** @format int64 */
  total_storage?: string;

  /** @format int64 */
  last_reward_block?: string;
}

export interface NodeQueryAllNodeResponse {
  node?: NodeNode[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface NodeQueryAllPledgeResponse {
  pledge?: NodePledge[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface NodeQueryGetNodeResponse {
  node?: NodeNode;
}

export interface NodeQueryGetPledgeResponse {
  pledge?: NodePledge;
}

export interface NodeQueryGetPoolResponse {
  Pool?: NodePool;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface NodeQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: NodeParams;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
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

/**
* DecCoin defines a token with a denomination and a decimal amount.

NOTE: The amount field is an Dec which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1DecCoin {
  denom?: string;
  amount?: string;
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
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
 * @title sao/node/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryNodeAll
   * @summary Queries a list of Node items.
   * @request GET:/SaoNetwork/sao/node/node
   */
  queryNodeAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
      status?: number;
    },
    params: RequestParams = {},
  ) =>
    this.request<NodeQueryAllNodeResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/node`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNode
   * @summary Queries a Node by index.
   * @request GET:/SaoNetwork/sao/node/node/{creator}
   */
  queryNode = (creator: string, params: RequestParams = {}) =>
    this.request<NodeQueryGetNodeResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/node/${creator}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/SaoNetwork/sao/node/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<NodeQueryParamsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPledgeAll
   * @summary Queries a list of Pledge items.
   * @request GET:/SaoNetwork/sao/node/pledge
   */
  queryPledgeAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<NodeQueryAllPledgeResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/pledge`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPledge
   * @summary Queries a Pledge by index.
   * @request GET:/SaoNetwork/sao/node/pledge/{creator}
   */
  queryPledge = (creator: string, params: RequestParams = {}) =>
    this.request<NodeQueryGetPledgeResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/pledge/${creator}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPool
   * @summary Queries a Pool by index.
   * @request GET:/SaoNetwork/sao/node/pool
   */
  queryPool = (params: RequestParams = {}) =>
    this.request<NodeQueryGetPoolResponse, RpcStatus>({
      path: `/SaoNetwork/sao/node/pool`,
      method: "GET",
      format: "json",
      ...params,
    });
}
