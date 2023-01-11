import request from '@/utils/request'
import api from './api'

export function getSubjects(params) {
    return request({
        url: api.subjects,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getSubjectById(id) {
    return request({
        url: api.subjects + `/${id}`,
        method: "GET"
    })
}

export function getSubjectPostCount(id) {
    return request({
        url: api.posts,
        method: "GET",
        params: {
            page_no: 1,
            page_size: 10,
            subject_id: id
        }
    })
}

export function getPostBySubjectId(id, pageNo, pageSize) {
    return request({
        url: api.posts,
        method: "GET",
        params: {
            pageNo,
            pageSize,
            subject_id: id
        }
    })
}
