
import request from '@/utils/request'
// import api from './api'

export function login(data) {
    return request({
        url:'/api/v1/auth/login',
        method: 'POST',
        data: {
            ...data
        }
    })
}

export function getRolesByAccessToken(access_token){
    return request({
        url: '/api/v1/auth/roles',
        method: 'GET',
        params: {
            access_token
        }
    })
}

export function getInfoByAccessToken(access_token) {
    return request({
        url: '/api/v1/auth/info',
        method: 'GET',
        params: {
            access_token
        }
    })
}

export function getPermissionByAccessToken(access_token){
    return request({
        url: '/api/v1/auth/permissions',
        method: 'GET',
        params: {
            access_token
        }
    })
}

export function refreshToken(refresh_token){
    return request({
        url: '/api/v1/auth/refresh',
        method: 'GET',
        params: {
            refresh_token
        }
    })
}

