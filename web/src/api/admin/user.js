import request from '@/utils/request'
import api from './api'

export function listUsers(params){
    return request({
        url: api.users,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getUserById(id) {
    return request({
        url: api.users + `/${id}`,
        menthod: 'GET',
    })
}

export function getUserInfo() {
    return request({
        url: `${api.auth}/info`,
        method: 'GET',
    })
}