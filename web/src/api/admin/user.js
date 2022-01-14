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