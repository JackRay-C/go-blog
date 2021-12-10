import request from '../utils/request'
import api from './api'


export function getFileById(id) {
    return request({
        url:  api.file + `/${id}`,
        method: 'GET'
    })
}


export function listFiles(params) {
    return request ({
        url: api.file,
        method: 'GET',
        params: {
            ...params
        }
    })
}