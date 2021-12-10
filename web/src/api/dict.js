import request from '../utils/request'
import api from './api'


export function listDicts(params) {
    return request({
        url:  api.dict,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function postDict(data) {
    return request({
        url: api.dict,
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function getDict(id) {
    return request ({
        url: api.dict + `${id}`,
        method: 'GET'
    })
}

export function putDict(id, data) {
    return request({
        url: api.dict + `${id}`,
        method: 'PUT',
        data: {
            ...data
        }
    })
}