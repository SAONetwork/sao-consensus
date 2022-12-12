import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgReset } from "./types/sao/node/tx";
import { MsgClaimReward } from "./types/sao/node/tx";
import { MsgLogin } from "./types/sao/node/tx";
import { MsgLogout } from "./types/sao/node/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.node.MsgReset", MsgReset],
    ["/saonetwork.sao.node.MsgClaimReward", MsgClaimReward],
    ["/saonetwork.sao.node.MsgLogin", MsgLogin],
    ["/saonetwork.sao.node.MsgLogout", MsgLogout],
    
];

export { msgTypes }