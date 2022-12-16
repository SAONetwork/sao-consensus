import {queryClient as didQueryClient} from "./saonetwork.sao.did"
import {queryClient as nodeQueryClient} from "./saonetwork.sao.node"
import {Api as DidApi} from "./saonetwork.sao.did/rest"
import {Api as NodeApi} from "./saonetwork.sao.node/rest"

export {Client} from "./index"
export {MsgUpdateSidDocumentResponse} from "./saonetwork.sao.did/types/sao/did/tx"
export {
    // MsgStoreResponse,
    Proposal,
    QueryProposal,
    RenewProposal,
    PermissionProposal
} from "./saonetwork.sao.sao"
export {MsgStoreResponse} from "./saonetwork.sao.sao/types/sao/sao/tx"
export {TxMsgData} from "./cosmos.tx.v1beta1/types/cosmos/base/abci/v1beta1/abci";
export {nodeQueryClient, didQueryClient, NodeApi, DidApi}