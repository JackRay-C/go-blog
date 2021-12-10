import request from '@/utils/request'
import api from '@/api/admin/admin'

export function listSubjects(params) {
    return request({
        url: api.subject,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getSubjectById(id) {
    return request({
        url: api.subject + `/${id}`,
        method: "GET"
    })
}

export function getSubjectPostCount(id) {
    return request({
        url: api.subject + `/${id}/posts`,
        method: "GET"
    })
}

export function getPostBySubjectId(id, pageNo, pageSize) {
    return request({
        url: api.subject + `/${id}/posts`,
        method: "GET",
        params: {
            pageNo,
            pageSize
        }
    })
}
