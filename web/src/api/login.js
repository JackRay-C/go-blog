
import request from '../utils/request'

export function login(data) {
    return request({
        url:'/api/v1/auth/login',
        method: 'POST',
        data: {
            ...data
        }
    })
}