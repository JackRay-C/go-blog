import request from '@/utils/request'
import api from './api'


export function listDicts(params) {
    return request({
        url:  api.dicts,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function postDict(data) {
    return request({
        url: api.dicts,
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function getDict(id) {
    return request ({
        url: api.dicts + `${id}`,
        method: 'GET'
    })
}

export function putDict(id, data) {
    return request({
        url: api.dicts + `${id}`,
        method: 'PUT',
        data: {
            ...data
        }
    })
}