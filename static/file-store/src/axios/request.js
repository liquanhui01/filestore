// 一、配置axios
import axios from 'axios'
import qs from 'qs'
import {
  ElMessage
} from 'element-plus'

// import store from '@/store/index' 如果使用vuex，那么token，userinfo都可以在登录以后存储到store中，不需要使用storage
// 获取浏览器的接口地址。
let baseUrl = "http://127.0.0.1:8000"
// axios配置
axios.defaults.baseURL = baseUrl
// 设置请求最大时长
axios.defaults.timeout = 5000

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8'

// 请求拦截器，设置token
axios.interceptors.request.use(config => {
  if (localStorage && localStorage.getItem('token')) {
    const token = localStorage.getItem('token')
    token && (config.headers.Authorization = token)
  }
  return config
}, error => {
  // 可以安装elementui等ui组件，将错误信息输出到界面。
  return Promise.error(error)
})
// 响应拦截器
axios.interceptors.response.use(response => {
  return Promise.resolve(response)
}, error => {
  if (error.response.status == 401) {
    // this.$message = "登陆过期，请重新登陆"
    ElMessage.error({
      message: "登陆过期，请重新登陆",
      type: "error",
    })
    window.location.href = "/"
  } else {
    return Promise.reject(error)
  }
})

// 2、封装请求方式
// @param url 接口地址
// @param data 携带参数
// @param file 上传文件对象
// @param auth 是否携带token
// get请求
export function get(url, params) {
  return axios({
    method: "get",
    url: `${baseUrl}${url}`,
    params: params,
  })
}
// post请求
export function post(url, data) {
  return axios.post(url, qs.stringify(data))
}
// put请求
export function put(url, data) {
  return axios({
    method: "put",
    url: `${baseUrl}${url}`,
    params: data,
  })
}
// patch请求
export function patch(url, data) {
  return axios({
    method: "PATCH",
    url: `${baseUrl}${url}`,
    param: data,
  })
}
// delete 请求
export function del(url, data) {
  return axios({
    method: "delete",
    url: `${baseUrl}${url}`,
    params: data,
  })
}
// upload 请求
export function uploader(url, params, formData, config) {
  return axios({
    method: "post",
    url: `${baseUrl}${url}`,
    params: params,
    headers: {
      "Content-Type": "multipart/form-data",
    },
    config,
    formData
  })
}

// download请求
export function download(url, data) {
  return axios({
    method: "get",
    url: `${baseUrl}${url}`,
    params: data,
    responseType: "blob", // 服务器返回的数据类型
  })
}