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
        url: api.post,
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function deletePost(postId) {
    return request({
        url: api.post + `/${postId}`,
        method:'DELETE'
    })
}

export function putPost(postId, data) {
    return request({
        url: api.post + `/${postId}`,
        method:'PUT',
        data: {
            ...data
        }
    })
}


export function listPostTags(postId) {
    return request({
        url: api.post + `/${postId}/tags`,
        method: "GET"
    })
}

export function addPostTags(postId, data) {
    return request({
        url: api.post + `/${postId}/tags`,
        method: 'POST',
        data: {
            ...data
        }
    })
}


export function putPostTags(postId, data) {
    return request({
        url: api.post +`/${postId}/tags`,
        method: 'PUT',
        data: {
            ...data
        }
    })
}


export function deletePostTags(postId, tagId) {
    return request({
        url: api.post + `/${postId}/${tagId}`,
        method: 'DELETE'
    })
}


export function listPostComments(postId, params) {
    return request({
        url: api.post + `/${postId}/comments`,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function addPostComment(postId, data) {
    return request({
        url: api.post + `/${postId}/comments`,
        method: 'POST',
        data: {
            ...data
        }
    })
}