import request from '@/utils/request'
import api from '@/api/admin/admin'

export function listUsers(params){
    return request({
        url: api.user,
        method: 'GET',
        params: {
            ...params
        }
    })
}