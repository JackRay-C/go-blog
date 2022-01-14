import request from '@/utils/request'
import api from './api'

export function getTags(parmas) {
    return request({
        url: api.tags,
        method: 'GET',
        params: {
            ...parmas
        }
    })
}

export function getTagsById(id) {
    return request({
        url: `${api.tags}/${id}`,
        method: 'GET'
    })
}


export function getPostByTagId(tagId, pageNo, pageSize) {
    return request({
        url: `${api.tags}/${tagId}/posts`,
        method: 'GET',
        params: {
            pageNo,
            pageSize
        }
    })
}