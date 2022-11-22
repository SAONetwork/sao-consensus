import { Client, registry, MissingWalletError } from 'SaoNetwork-sao-client-ts'

import { Binding } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { BindingProof } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { DidBindingProofs } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"
import { Params } from "SaoNetwork-sao-client-ts/saonetwork.sao.did/types"


export { Binding, BindingProof, DidBindingProofs, Params };

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
				DidBindingProofs: {},
				DidBindingProofsAll: {},
				
				_Structure: {
						Binding: getStructure(Binding.fromPartial({})),
						BindingProof: getStructure(BindingProof.fromPartial({})),
						DidBindingProofs: getStructure(DidBindingProofs.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
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
				getDidBindingProofs: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidBindingProofs[JSON.stringify(params)] ?? {}
		},
				getDidBindingProofsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DidBindingProofsAll[JSON.stringify(params)] ?? {}
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
		
		
		
		
		 		
		
		
		async QueryDidBindingProofs({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryDidBindingProofs( key.accountId)).data
				
					
				commit('QUERY', { query: 'DidBindingProofs', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidBindingProofs', payload: { options: { all }, params: {...key},query }})
				return getters['getDidBindingProofs']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidBindingProofs API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDidBindingProofsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SaonetworkSaoDid.query.queryDidBindingProofsAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SaonetworkSaoDid.query.queryDidBindingProofsAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DidBindingProofsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDidBindingProofsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDidBindingProofsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDidBindingProofsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgGetBinding({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoDid.tx.sendMsgGetBinding({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgGetBinding:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgGetBinding:Send Could not broadcast Tx: '+ e.message)
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
		
		async MsgGetBinding({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoDid.tx.msgGetBinding({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgGetBinding:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgGetBinding:Create Could not create message: ' + e.message)
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
		
	}
}
