// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SaonetworkSaoNode from './saonetwork.sao.node'
import SaonetworkSaoSao from './saonetwork.sao.sao'


export default { 
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