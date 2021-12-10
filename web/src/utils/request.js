import axios from 'axios'
import router from '../router'

const instance = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000,
    headers: { 'Content-Type': 'application/json;charset=utf8' },
    withCredentials: true
})

instance.interceptors.request.use(config => {
    if(localStorage.getItem("token")) {
        config.headers.token = localStorage.getItem("token")
    }
    return config
}, error => {
    return Promise.reject(error)
})


instance.interceptors.response.use(res => {
    // 通用response 处理
    if (res.status !== 200) {
        return Promise.reject(res.statusText)
    } else {
        if(res.data.code === 401 || res.data.code === 1003 || res.data.code === 1004) {
            localStorage.removeItem("token")
            router.push("/login")
            return Promise.reject(res)
        }
        return res.data
    }
}, error => {
    return Promise.reject(error)
})

export default instance