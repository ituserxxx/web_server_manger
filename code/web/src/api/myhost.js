import request from '@/utils/request'

export function myhostAdd(data) {
  return request({
    url: '/myhost/Add',
    method: 'post',
    data
  })
}
export function myhostDel(data) {
  return request({
    url: '/myhost/Del',
    method: 'post',
    data
  })
}
export function myhostUpdate(data) {
  return request({
    url: '/myhost/Update',
    method: 'post',
    data
  })
}
export function myhostGetList(data) {
  return request({
    url: '/myhost/GetList',
    method: 'post',
    data
  })
}