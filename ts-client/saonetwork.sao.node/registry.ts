import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLogout } from "./types/node/tx";
import { MsgLogin } from "./types/node/tx";
import { MsgReset } from "./types/node/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.node.MsgLogout", MsgLogout],
    ["/saonetwork.sao.node.MsgLogin", MsgLogin],
    ["/saonetwork.sao.node.MsgReset", MsgReset],
    
];

export { msgTypes }