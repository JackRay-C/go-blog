
import request from '@/utils/request'
import api from './api'

export function login(data) {
    return request({
        url:'/api/v1/auth/login',
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function getUserInfo() {
    return request({
        url: `${api.users}/info`,
        method: 'GET',
    })
}