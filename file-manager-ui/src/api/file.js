// 查询项目列表
import request from "@/utils/request";
import {download} from "@/utils/request";

export function listFile(data) {
    return request({
        url: '/file/list',
        method: 'post',
        data: data
    })
}

export function DeleteFile(data) {
    return request({
        url: '/file/delete',
        method: 'post',
        data: data
    })
}

export function downFile(data) {
    return download('/file/down', data, data.name)
}
