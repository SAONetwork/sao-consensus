import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLogin } from "./types/node/tx";
import { MsgReset } from "./types/node/tx";
import { MsgLogout } from "./types/node/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.node.MsgLogin", MsgLogin],
    ["/saonetwork.sao.node.MsgReset", MsgReset],
    ["/saonetwork.sao.node.MsgLogout", MsgLogout],
    
];

export { msgTypes }