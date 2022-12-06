import { Client, registry, MissingWalletError } from 'SaoNetwork-sao-client-ts'

import { JwsSignature } from "SaoNetwork-sao-client-ts/saonetwork.sao.sao/types"
import { Params } from "SaoNetwork-sao-client-ts/saonetwork.sao.sao/types"
import { PermissionProposal } from "SaoNetwork-sao-client-ts/saonetwork.sao.sao/types"
import { Proposal } from "SaoNetwork-sao-client-ts/saonetwork.sao.sao/types"
import { RenewProposal } from "SaoNetwork-sao-client-ts/saonetwork.sao.sao/types"


export { JwsSignature, Params, PermissionProposal, Proposal, RenewProposal };

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
				
				_Structure: {
						JwsSignature: getStructure(JwsSignature.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PermissionProposal: getStructure(PermissionProposal.fromPartial({})),
						Proposal: getStructure(Proposal.fromPartial({})),
						RenewProposal: getStructure(RenewProposal.fromPartial({})),
						
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
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: saonetwork.sao.sao initialized!')
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
				let value= (await client.SaonetworkSaoSao.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdataPermission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgUpdataPermission({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdataPermission:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdataPermission:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgReject({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgReject({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReject:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgReject:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgCancel({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgReady({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgReady({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReady:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgReady:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgTerminate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgTerminate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTerminate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgTerminate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRenew({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgRenew({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRenew:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRenew:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgStore({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgStore({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStore:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgStore:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgComplete({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SaonetworkSaoSao.tx.sendMsgComplete({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgComplete:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgComplete:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdataPermission({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgUpdataPermission({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdataPermission:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdataPermission:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgReject({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgReject({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReject:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgReject:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgCancel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancel:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgReady({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgReady({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReady:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgReady:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgTerminate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgTerminate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTerminate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgTerminate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRenew({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgRenew({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRenew:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRenew:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgStore({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgStore({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStore:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgStore:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgComplete({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SaonetworkSaoSao.tx.msgComplete({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgComplete:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgComplete:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
