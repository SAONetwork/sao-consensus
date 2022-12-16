import { Client, registry, MissingWalletError } from 'SaoNetwork-sao-client-ts'

import { AccountAuth } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { AccountList } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { BindingProof } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { DidBindingProof } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { Params } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { PastSeeds } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { PaymentAddress } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { PubKey } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { SidDocument } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { SidDocumentVersion } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"


export { AccountAuth, AccountList, BindingProof, DidBindingProof, Params, PastSeeds, PaymentAddress, PubKey, SidDocument, SidDocumentVersion };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				DidBindingProof: {},
				DidBindingProofAll: {},
				AccountList: {},
				AccountListAll: {},
				AccountAuth: {},
				AccountAuthAll: {},
				GetAllAccountAuths: {},
				SidDocument: {},
				SidDocumentAll: {},
				SidDocumentVersion: {},
				SidDocumentVersionAll: {},
				PastSeeds: {},
				PastSeedsAll: {},
				PaymentAddress: {},
				PaymentAddressAll: {},
				ValidateDid: {},
				
				_Structure: {
						AccountAuth: getStructure(AccountAuth.fromPartial({})),
						AccountList: getStructure(AccountList.fromPartial({})),
						BindingProof: getStructure(BindingProof.fromPartial({})),
						DidBindingProof: getStructure(DidBindingProof.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PastSeeds: getStructure(PastSeeds.fromPartial({})),
						PaymentAddress: getStructure(PaymentAddress.fromPartial({})),
						PubKey: getStructure(PubKey.fromPartial({})),
						SidDocument: getStructure(SidDocument.fromPartial({})),
						SidDocumentVersion: getStructure(SidDocumentVersion.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getDidBindingProof: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidBindingProof[JSON.stringify(params)] ?? {}
		},
				getDidBindingProofAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidBindingProofAll[JSON.stringify(params)] ?? {}
		},
				getAccountList: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountList[JSON.stringify(params)] ?? {}
		},
				getAccountListAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountListAll[JSON.stringify(params)] ?? {}
		},
				getAccountAuth: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountAuth[JSON.stringify(params)] ?? {}
		},
				getAccountAuthAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountAuthAll[JSON.stringify(params)] ?? {}
		},
				getGetAllAccountAuths: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetAllAccountAuths[JSON.stringify(params)] ?? {}
		},
				getSidDocument: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SidDocument[JSON.stringify(params)] ?? {}
		},
				getSidDocumentAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SidDocumentAll[JSON.stringify(params)] ?? {}
		},
				getSidDocumentVersion: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SidDocumentVersion[JSON.stringify(params)] ?? {}
		},
				getSidDocumentVersionAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SidDocumentVersionAll[JSON.stringify(params)] ?? {}
		},
				getPastSeeds: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PastSeeds[JSON.stringify(params)] ?? {}
		},
				getPastSeedsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PastSeedsAll[JSON.stringify(params)] ?? {}
		},
				getPaymentAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PaymentAddress[JSON.stringify(params)] ?? {}
		},
				getPaymentAddressAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PaymentAddressAll[JSON.stringify(params)] ?? {}
		},
				getValidateDid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidateDid[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: saonetwork.sao.did initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidBindingProof({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryDidBindingProof( key.accountId)).data
				
					
				commit('QUERY', { query: 'DidBindingProof', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidBindingProof', payload: { options: { all }, params: {...key},query }})
				return getters['getDidBindingProof']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidBindingProof API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidBindingProofAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryDidBindingProofAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryDidBindingProofAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DidBindingProofAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidBindingProofAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDidBindingProofAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidBindingProofAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountList({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryAccountList( key.did)).data
				
					
				commit('QUERY', { query: 'AccountList', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountList', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountList']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountList API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountListAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryAccountListAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryAccountListAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AccountListAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountListAll', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountListAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountListAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountAuth({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryAccountAuth( key.accountDid)).data
				
					
				commit('QUERY', { query: 'AccountAuth', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountAuth', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountAuth']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountAuth API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountAuthAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryAccountAuthAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryAccountAuthAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AccountAuthAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountAuthAll', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountAuthAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountAuthAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetAllAccountAuths({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryGetAllAccountAuths( key.did)).data
				
					
				commit('QUERY', { query: 'GetAllAccountAuths', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetAllAccountAuths', payload: { options: { all }, params: {...key},query }})
				return getters['getGetAllAccountAuths']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetAllAccountAuths API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySidDocument({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.querySidDocument( key.versionId)).data
				
					
				commit('QUERY', { query: 'SidDocument', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySidDocument', payload: { options: { all }, params: {...key},query }})
				return getters['getSidDocument']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySidDocument API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySidDocumentAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.querySidDocumentAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.querySidDocumentAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SidDocumentAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySidDocumentAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSidDocumentAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySidDocumentAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySidDocumentVersion({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.querySidDocumentVersion( key.docId)).data
				
					
				commit('QUERY', { query: 'SidDocumentVersion', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySidDocumentVersion', payload: { options: { all }, params: {...key},query }})
				return getters['getSidDocumentVersion']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySidDocumentVersion API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySidDocumentVersionAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.querySidDocumentVersionAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.querySidDocumentVersionAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SidDocumentVersionAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySidDocumentVersionAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSidDocumentVersionAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySidDocumentVersionAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPastSeeds({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryPastSeeds( key.did)).data
				
					
				commit('QUERY', { query: 'PastSeeds', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPastSeeds', payload: { options: { all }, params: {...key},query }})
				return getters['getPastSeeds']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPastSeeds API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPastSeedsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryPastSeedsAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryPastSeedsAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PastSeedsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPastSeedsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPastSeedsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPastSeedsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPaymentAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryPaymentAddress( key.did)).data
				
					
				commit('QUERY', { query: 'PaymentAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPaymentAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getPaymentAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPaymentAddress API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPaymentAddressAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryPaymentAddressAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryPaymentAddressAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PaymentAddressAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPaymentAddressAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPaymentAddressAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPaymentAddressAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidateDid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryValidateDid( key.did)).data
				
					
				commit('QUERY', { query: 'ValidateDid', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidateDid', payload: { options: { all }, params: {...key},query }})
				return getters['getValidateDid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidateDid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgAddPastSeed({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgAddPastSeed({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddPastSeed:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddPastSeed:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUnbinding({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgUnbinding({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnbinding:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUnbinding:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdatePaymentAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgUpdatePaymentAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePaymentAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdatePaymentAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddBinding({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgAddBinding({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddBinding:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddBinding:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateAccountAuths({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgUpdateAccountAuths({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAccountAuths:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateAccountAuths:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBinding({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgBinding({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBinding:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBinding:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAccountAuth({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgAddAccountAuth({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAccountAuth:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAccountAuth:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateSidDocument({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgUpdateSidDocument({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSidDocument:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateSidDocument:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgResetStore({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgResetStore({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgResetStore:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgResetStore:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgAddPastSeed({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgAddPastSeed({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddPastSeed:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddPastSeed:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUnbinding({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgUnbinding({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUnbinding:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUnbinding:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdatePaymentAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgUpdatePaymentAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePaymentAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdatePaymentAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddBinding({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgAddBinding({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddBinding:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddBinding:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateAccountAuths({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgUpdateAccountAuths({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAccountAuths:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateAccountAuths:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBinding({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgBinding({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBinding:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBinding:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAccountAuth({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgAddAccountAuth({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAccountAuth:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAccountAuth:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateSidDocument({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgUpdateSidDocument({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSidDocument:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateSidDocument:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgResetStore({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgResetStore({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgResetStore:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgResetStore:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}