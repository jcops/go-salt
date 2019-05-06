import request from '@/utils/request'

export function saltKeyping() {
  return request({
    url: '/api/v1/keyping',
    method: 'get'
  })
}
export function saltlist() {
  return request({
    url: '/api/v1/keylist',
    method: 'get'
  })
}
export function saltdelete(data) {
  return request({
    url: '/api/v1/keydelete',
    method: 'post',
    data
  })
}
export function saltaccept(data) {
  return request({
    url: '/api/v1/keyaccept',
    method: 'post',
    data
  })
}
export function saltrun(data) {
  return request({
    url: '/api/v1/keyrun',
    method: 'post',
    data
  })
}
export function saltdeploy(data) {
  return request({
    url: '/api/v1/keydeploy',
    method: 'post',
    timeout: 300000,
    data
  })
}
export function saltcopy(data) {
  return request({
    url: '/api/v1/keycopy',
    method: 'post',
    timeout: 300000,
    data
  })
}
