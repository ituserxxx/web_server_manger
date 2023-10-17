import request from '@/utils/request'

export function myhostAdd(data) {
  return request({
    url: '/myhost/add',
    method: 'post',
    data
  })
}
export function myhostDel(data) {
  return request({
    url: '/myhost/del',
    method: 'post',
    data
  })
}
export function myhostUpdate(data) {
  return request({
    url: '/myhost/update',
    method: 'post',
    data
  })
}
export function myhostGetList(data) {
  return request({
    url: '/myhost/getList',
    method: 'post',
    data
  })
}