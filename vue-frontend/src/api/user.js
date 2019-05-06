import request from '@/utils/request'
// 登录
export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data
  })
}

// 获取用户信息
export function getInfo() {
  return request({
    url: '/api/v1/getinfo',
    method: 'get'
  })
}

// 获取用户列表
export function getUserList(params) {
  return request({
    url: '/api/v1/users/',
    method: 'get',
    params
  })
}

// 创建用户
export function createUser(data) {
  return request({
    url: '/api/v1/register',
    method: 'post',
    data
  })
}

// 修改用户
export function updateUser(data) {
  return request({
    url: '/api/v1/users',
    method: 'put',
    data
  })
}

// 删除用户
export function deleteUser(id) {
  return request({
    url: '/api/v1/users/' + id,
    method: 'delete'
  })
}

// 将用户加入组中
export function updateUserGroup(id, data) {
  return request({
    url: '/users/' + id + '/',
    method: 'patch',
    data
  })
}

// 将用户设置为管理员
export function setUserSuper(id, data) {
  return request({
    url: '/setusersuper/' + id + '/',
    method: 'put',
    data
  })
}

