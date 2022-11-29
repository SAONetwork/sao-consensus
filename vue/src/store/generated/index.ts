// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SaonetworkSaoDid from './saonetwork.sao.did'
import SaonetworkSaoModel from './saonetwork.sao.model'
import SaonetworkSaoNode from './saonetwork.sao.node'
import SaonetworkSaoOrder from './saonetwork.sao.order'
import SaonetworkSaoSao from './saonetwork.sao.sao'


export default { 
  SaonetworkSaoDid: load(SaonetworkSaoDid, 'saonetwork.sao.did'),
  SaonetworkSaoModel: load(SaonetworkSaoModel, 'saonetwork.sao.model'),
  SaonetworkSaoNode: load(SaonetworkSaoNode, 'saonetwork.sao.node'),
  SaonetworkSaoOrder: load(SaonetworkSaoOrder, 'saonetwork.sao.order'),
  SaonetworkSaoSao: load(SaonetworkSaoSao, 'saonetwork.sao.sao'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}