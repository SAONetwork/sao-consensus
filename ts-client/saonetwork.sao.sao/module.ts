// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgTerminate } from "./types/sao/tx";
import { MsgComplete } from "./types/sao/tx";
import { MsgStore } from "./types/sao/tx";
import { MsgReject } from "./types/sao/tx";
import { MsgCancel } from "./types/sao/tx";


export { MsgTerminate, MsgComplete, MsgStore, MsgReject, MsgCancel };

type sendMsgTerminateParams = {
  value: MsgTerminate,
  fee?: StdFee,
  memo?: string
};

type sendMsgCompleteParams = {
  value: MsgComplete,
  fee?: StdFee,
  memo?: string
};

type sendMsgStoreParams = {
  value: MsgStore,
  fee?: StdFee,
  memo?: string
};

type sendMsgRejectParams = {
  value: MsgReject,
  fee?: StdFee,
  memo?: string
};

type sendMsgCancelParams = {
  value: MsgCancel,
  fee?: StdFee,
  memo?: string
};


type msgTerminateParams = {
  value: MsgTerminate,
};

type msgCompleteParams = {
  value: MsgComplete,
};

type msgStoreParams = {
  value: MsgStore,
};

type msgRejectParams = {
  value: MsgReject,
};

type msgCancelParams = {
  value: MsgCancel,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgTerminate({ value, fee, memo }: sendMsgTerminateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgTerminate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgTerminate({ value: MsgTerminate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgTerminate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgComplete({ value, fee, memo }: sendMsgCompleteParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgComplete: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgComplete({ value: MsgComplete.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgComplete: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgStore({ value, fee, memo }: sendMsgStoreParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgStore: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgStore({ value: MsgStore.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgStore: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgReject({ value, fee, memo }: sendMsgRejectParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgReject: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgReject({ value: MsgReject.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgReject: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCancel({ value, fee, memo }: sendMsgCancelParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCancel: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCancel({ value: MsgCancel.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCancel: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgTerminate({ value }: msgTerminateParams): EncodeObject {
			try {
				return { typeUrl: "/saonetwork.sao.sao.MsgTerminate", value: MsgTerminate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgTerminate: Could not create message: ' + e.message)
			}
		},
		
		msgComplete({ value }: msgCompleteParams): EncodeObject {
			try {
				return { typeUrl: "/saonetwork.sao.sao.MsgComplete", value: MsgComplete.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgComplete: Could not create message: ' + e.message)
			}
		},
		
		msgStore({ value }: msgStoreParams): EncodeObject {
			try {
				return { typeUrl: "/saonetwork.sao.sao.MsgStore", value: MsgStore.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgStore: Could not create message: ' + e.message)
			}
		},
		
		msgReject({ value }: msgRejectParams): EncodeObject {
			try {
				return { typeUrl: "/saonetwork.sao.sao.MsgReject", value: MsgReject.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgReject: Could not create message: ' + e.message)
			}
		},
		
		msgCancel({ value }: msgCancelParams): EncodeObject {
			try {
				return { typeUrl: "/saonetwork.sao.sao.MsgCancel", value: MsgCancel.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCancel: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			SaonetworkSaoSao: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;