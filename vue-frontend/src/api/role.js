import request from '@/utils/request'

export function getRoutes() {
  return request({
    url: '/routes',
    method: 'get'
  })
}

export function getRoleList(params) {
  return request({
    url: '/api/v1/role',
    method: 'get',
    params
  })
}

export function createGroup(data) {
  return request({
    url: '/api/v1/role',
    method: 'post',
    data
  })
}

export function updateGroup(data) {
  return request({
    url: `/api/v1/role`,
    method: 'put',
    data
  })
}

// 获取权限列表
export function getPermissionList() {
  return request({
    url: '/api/v1/menu',
    method: 'get'
  })
}
// 修改角色权限
export function updateGroupPower(data) {
  return request({
    url: '/api/v1/role',
    method: 'put',
    data
  })
}
export function deleteGroup(id) {
  return request({
    url: `/api/v1/role/${id}`,
    method: 'delete'
  })
}

// 将指定用户从组里面删除
export function deleteGroupMember(gid, data) {
  return request({
    url: '/groupmembers/' + gid + '/',
    method: 'delete',
    data
  })
}
