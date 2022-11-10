import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgReady } from "./types/sao/tx";
import { MsgStore } from "./types/sao/tx";
import { MsgReject } from "./types/sao/tx";
import { MsgCancel } from "./types/sao/tx";
import { MsgComplete } from "./types/sao/tx";
import { MsgTerminate } from "./types/sao/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.sao.MsgReady", MsgReady],
    ["/saonetwork.sao.sao.MsgStore", MsgStore],
    ["/saonetwork.sao.sao.MsgReject", MsgReject],
    ["/saonetwork.sao.sao.MsgCancel", MsgCancel],
    ["/saonetwork.sao.sao.MsgComplete", MsgComplete],
    ["/saonetwork.sao.sao.MsgTerminate", MsgTerminate],
    
];

export { msgTypes }