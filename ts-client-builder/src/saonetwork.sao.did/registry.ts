import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddPastSeed } from "./types/sao/did/tx";
import { MsgUpdateAccountAuths } from "./types/sao/did/tx";
import { MsgAddAccountAuth } from "./types/sao/did/tx";
import { MsgUnbinding } from "./types/sao/did/tx";
import { MsgResetStore } from "./types/sao/did/tx";
import { MsgAddBinding } from "./types/sao/did/tx";
import { MsgBinding } from "./types/sao/did/tx";
import { MsgUpdatePaymentAddress } from "./types/sao/did/tx";
import { MsgUpdateSidDocument } from "./types/sao/did/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.did.MsgAddPastSeed", MsgAddPastSeed],
    ["/saonetwork.sao.did.MsgUpdateAccountAuths", MsgUpdateAccountAuths],
    ["/saonetwork.sao.did.MsgAddAccountAuth", MsgAddAccountAuth],
    ["/saonetwork.sao.did.MsgUnbinding", MsgUnbinding],
    ["/saonetwork.sao.did.MsgResetStore", MsgResetStore],
    ["/saonetwork.sao.did.MsgAddBinding", MsgAddBinding],
    ["/saonetwork.sao.did.MsgBinding", MsgBinding],
    ["/saonetwork.sao.did.MsgUpdatePaymentAddress", MsgUpdatePaymentAddress],
    ["/saonetwork.sao.did.MsgUpdateSidDocument", MsgUpdateSidDocument],
    
];

export { msgTypes }