import request from '@/utils/request'
import api from './api'


export function listDrafts(params) {
    return request({
        url: api.drafts,
        method: 'GET',
        params: {
            ...params
        }
    })
}

export function getDraft(id) {
    return request({
        url: api.drafts + `/${id}`,
        method: "GET"
    })
}


export function addDraft(data) {
    return request({
        url: api.drafts,
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function deleteDraft(id) {
    return request({
        url: api.drafts + `/${id}`,
        method:'DELETE'
    })
}

export function putDraft(id, data) {
    return request({
        url: api.drafts + `/${id}`,
        method:'PUT',
        data: {
            ...data
        }
    })
}