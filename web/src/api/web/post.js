import request from '@/utils/request'
import api from './api'

export function listPosts(params) {
    return request({
        url: api.posts,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getPost(id) {
    return request({
        url: api.posts + `/${id}`,
        method: "GET"
    })
}






