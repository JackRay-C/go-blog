import request from '@/utils/request'
import api from '@/api/admin/admin'

export function listPosts(params) {
    return request({
        url: api.post,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getPost(id) {
    return request({
        url: api.post + `/${id}`,
        method: "GET"
    })
}


export function addPost(data) {
    return request({
        url: api.posts,
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function deletePost(postId) {
    return request({
        url: api.posts + `/${postId}`,
        method:'DELETE'
    })
}

export function putPost(postId, data) {
    return request({
        url: api.posts + `/${postId}`,
        method:'PUT',
        data: {
            ...data
        }
    })
}