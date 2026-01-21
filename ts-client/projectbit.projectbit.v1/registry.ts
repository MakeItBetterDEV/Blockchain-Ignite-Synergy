import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/projectbit/projectbit/v1/tx";
import { MsgCreatePost } from "./types/projectbit/projectbit/v1/tx";
import { MsgUpdatePost } from "./types/projectbit/projectbit/v1/tx";
import { MsgDeletePost } from "./types/projectbit/projectbit/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/projectbit.projectbit.v1.MsgUpdateParams", MsgUpdateParams],
    ["/projectbit.projectbit.v1.MsgCreatePost", MsgCreatePost],
    ["/projectbit.projectbit.v1.MsgUpdatePost", MsgUpdatePost],
    ["/projectbit.projectbit.v1.MsgDeletePost", MsgDeletePost],
    
];

export { msgTypes }