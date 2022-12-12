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

export interface DidAccountAuth {
  accountDid?: string;
  accountEncryptedSeed?: string;
  sidEncryptedAccount?: string;
}

export interface DidAccountList {
  did?: string;
  accountDids?: string[];
}

export interface DidBindingProof {
  /** @format int32 */
  version?: number;
  message?: string;
  signature?: string;
  account?: string;
  did?: string;

  /** @format uint64 */
  timestamp?: string;
}

export interface DidDidBindingProof {
  accountId?: string;
  proof?: DidBindingProof;
}

export type DidMsgAddAccountAuthResponse = object;

export type DidMsgAddBindingResponse = object;

export type DidMsgAddPastSeedResponse = object;

export type DidMsgBindingResponse = object;

export type DidMsgResetStoreResponse = object;

export type DidMsgUnbindingResponse = object;

export type DidMsgUpdateAccountAuthsResponse = object;

export type DidMsgUpdatePaymentAddressResponse = object;

export interface DidMsgUpdateSidDocumentResponse {
  docId?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type DidParams = object;

export interface DidPastSeeds {
  did?: string;
  seeds?: string[];
}

export interface DidPaymentAddress {
  did?: string;
  address?: string;
}

export interface DidPubKey {
  name?: string;
  value?: string;
}

export interface DidQueryAllAccountAuthResponse {
  accountAuth?: DidAccountAuth[];

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

export interface DidQueryAllAccountListResponse {
  accountList?: DidAccountList[];

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

export interface DidQueryAllDidBindingProofResponse {
  DidBindingProof?: DidDidBindingProof[];

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

export interface DidQueryAllPastSeedsResponse {
  pastSeeds?: DidPastSeeds[];

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

export interface DidQueryAllPaymentAddressResponse {
  paymentAddress?: DidPaymentAddress[];

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

export interface DidQueryAllSidDocumentResponse {
  sidDocument?: DidSidDocument[];

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

export interface DidQueryAllSidDocumentVersionResponse {
  sidDocumentVersion?: DidSidDocumentVersion[];

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

export interface DidQueryGetAccountAuthResponse {
  accountAuth?: DidAccountAuth;
}

export interface DidQueryGetAccountListResponse {
  accountList?: DidAccountList;
}

export interface DidQueryGetAllAccountAuthsResponse {
  accountAuths?: DidAccountAuth[];
}

export interface DidQueryGetDidBindingProofResponse {
  DidBindingProof?: DidDidBindingProof;
}

export interface DidQueryGetPastSeedsResponse {
  pastSeeds?: DidPastSeeds;
}

export interface DidQueryGetPaymentAddressResponse {
  paymentAddress?: DidPaymentAddress;
}

export interface DidQueryGetSidDocumentResponse {
  sidDocument?: DidSidDocument;
}

export interface DidQueryGetSidDocumentVersionResponse {
  sidDocumentVersion?: DidSidDocumentVersion;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface DidQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: DidParams;
}

export interface DidSidDocument {
  versionId?: string;
  keys?: DidPubKey[];
}

export interface DidSidDocumentVersion {
  docId?: string;
  versionList?: string[];
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
 * @title sao/did/account_auth.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountAuthAll
   * @summary Queries a list of AccountAuth items.
   * @request GET:/SaoNetwork/sao/did/account_auth
   */
  queryAccountAuthAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllAccountAuthResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/account_auth`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountAuth
   * @summary Queries a AccountAuth by index.
   * @request GET:/SaoNetwork/sao/did/account_auth/{accountDid}
   */
  queryAccountAuth = (accountDid: string, params: RequestParams = {}) =>
    this.request<DidQueryGetAccountAuthResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/account_auth/${accountDid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountListAll
   * @summary Queries a list of AccountList items.
   * @request GET:/SaoNetwork/sao/did/account_list
   */
  queryAccountListAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllAccountListResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/account_list`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountList
   * @summary Queries a AccountList by index.
   * @request GET:/SaoNetwork/sao/did/account_list/{did}
   */
  queryAccountList = (did: string, params: RequestParams = {}) =>
    this.request<DidQueryGetAccountListResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/account_list/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidBindingProofAll
   * @summary Queries a list of DidBindingProof items.
   * @request GET:/SaoNetwork/sao/did/did_binding_proof
   */
  queryDidBindingProofAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllDidBindingProofResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/did_binding_proof`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDidBindingProof
   * @summary Queries a DidBindingProof by index.
   * @request GET:/SaoNetwork/sao/did/did_binding_proof/{accountId}
   */
  queryDidBindingProof = (accountId: string, params: RequestParams = {}) =>
    this.request<DidQueryGetDidBindingProofResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/did_binding_proof/${accountId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetAllAccountAuths
   * @summary Queries a list of GetAllAccountAuth items.
   * @request GET:/SaoNetwork/sao/did/get_all_account_auths/{did}
   */
  queryGetAllAccountAuths = (did: string, params: RequestParams = {}) =>
    this.request<DidQueryGetAllAccountAuthsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/get_all_account_auths/${did}`,
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
   * @request GET:/SaoNetwork/sao/did/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<DidQueryParamsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPastSeedsAll
   * @summary Queries a list of PastSeeds items.
   * @request GET:/SaoNetwork/sao/did/past_seeds
   */
  queryPastSeedsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllPastSeedsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/past_seeds`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPastSeeds
   * @summary Queries a PastSeeds by index.
   * @request GET:/SaoNetwork/sao/did/past_seeds/{did}
   */
  queryPastSeeds = (did: string, params: RequestParams = {}) =>
    this.request<DidQueryGetPastSeedsResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/past_seeds/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPaymentAddressAll
   * @summary Queries a list of PaymentAddress items.
   * @request GET:/SaoNetwork/sao/did/payment_address
   */
  queryPaymentAddressAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllPaymentAddressResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/payment_address`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPaymentAddress
   * @summary Queries a PaymentAddress by index.
   * @request GET:/SaoNetwork/sao/did/payment_address/{did}
   */
  queryPaymentAddress = (did: string, params: RequestParams = {}) =>
    this.request<DidQueryGetPaymentAddressResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/payment_address/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySidDocumentAll
   * @summary Queries a list of SidDocument items.
   * @request GET:/SaoNetwork/sao/did/sid_document
   */
  querySidDocumentAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllSidDocumentResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/sid_document`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySidDocument
   * @summary Queries a SidDocument by index.
   * @request GET:/SaoNetwork/sao/did/sid_document/{versionId}
   */
  querySidDocument = (versionId: string, params: RequestParams = {}) =>
    this.request<DidQueryGetSidDocumentResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/sid_document/${versionId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySidDocumentVersionAll
   * @summary Queries a list of SidDocumentVersion items.
   * @request GET:/SaoNetwork/sao/did/sid_document_version
   */
  querySidDocumentVersionAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DidQueryAllSidDocumentVersionResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/sid_document_version`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySidDocumentVersion
   * @summary Queries a SidDocumentVersion by index.
   * @request GET:/SaoNetwork/sao/did/sid_document_version/{docId}
   */
  querySidDocumentVersion = (docId: string, params: RequestParams = {}) =>
    this.request<DidQueryGetSidDocumentVersionResponse, RpcStatus>({
      path: `/SaoNetwork/sao/did/sid_document_version/${docId}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
