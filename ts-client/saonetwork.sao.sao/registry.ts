import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgTerminate } from "./types/sao/tx";
import { MsgCancel } from "./types/sao/tx";
import { MsgStore } from "./types/sao/tx";
import { MsgComplete } from "./types/sao/tx";
import { MsgReject } from "./types/sao/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.sao.MsgTerminate", MsgTerminate],
    ["/saonetwork.sao.sao.MsgCancel", MsgCancel],
    ["/saonetwork.sao.sao.MsgStore", MsgStore],
    ["/saonetwork.sao.sao.MsgComplete", MsgComplete],
    ["/saonetwork.sao.sao.MsgReject", MsgReject],
    
];

export { msgTypes }