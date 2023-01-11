import request from '@/utils/request'
import api from './api'

export function listTags(params){
    return request({
        url: api.tags,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getTagById(id) {
    return request({
        url: api.tags + `/${id}`,
        menthod: 'GET',
    })
}

export function addTag(data) {
    return request({
        url: api.tags,
        method: 'POST',
        data: {
            ...data
        }
    })
}