import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgComplete } from "./types/sao/tx";
<<<<<<< HEAD
<<<<<<< HEAD
import { MsgReady } from "./types/sao/tx";
import { MsgReject } from "./types/sao/tx";
import { MsgCancel } from "./types/sao/tx";
=======
>>>>>>> dcf6850 (add claim)
=======
import { MsgReject } from "./types/sao/tx";
>>>>>>> 1938811 (move pledge to node module)
import { MsgTerminate } from "./types/sao/tx";
import { MsgStore } from "./types/sao/tx";
import { MsgCancel } from "./types/sao/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/saonetwork.sao.sao.MsgComplete", MsgComplete],
<<<<<<< HEAD
<<<<<<< HEAD
    ["/saonetwork.sao.sao.MsgReady", MsgReady],
    ["/saonetwork.sao.sao.MsgReject", MsgReject],
    ["/saonetwork.sao.sao.MsgCancel", MsgCancel],
=======
>>>>>>> dcf6850 (add claim)
=======
    ["/saonetwork.sao.sao.MsgReject", MsgReject],
>>>>>>> 1938811 (move pledge to node module)
    ["/saonetwork.sao.sao.MsgTerminate", MsgTerminate],
    ["/saonetwork.sao.sao.MsgStore", MsgStore],
    ["/saonetwork.sao.sao.MsgCancel", MsgCancel],
    
];

export { msgTypes }