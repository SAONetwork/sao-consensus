// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SaonetworkSaoEarn from './saonetwork.sao.earn'
import SaonetworkSaoModel from './saonetwork.sao.model'
import SaonetworkSaoNode from './saonetwork.sao.node'
import SaonetworkSaoSao from './saonetwork.sao.sao'


export default { 
  SaonetworkSaoEarn: load(SaonetworkSaoEarn, 'saonetwork.sao.earn'),
  SaonetworkSaoModel: load(SaonetworkSaoModel, 'saonetwork.sao.model'),
  SaonetworkSaoNode: load(SaonetworkSaoNode, 'saonetwork.sao.node'),
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