import request from '@/utils/request'


export function myserviceAdd(data) {
  return request({
    url: '/myservice/add',
    method: 'post',
    data
  })
}
export function myserviceDel(data) {
  return request({
    url: '/myservice/del',
    method: 'post',
    data
  })
}
export function myserviceUpdate(data) {
  return request({
    url: '/myservice/update',
    method: 'post',
    data
  })
}
export function myserviceGetList(data) {
  return request({
    url: '/myservice/getList',
    method: 'post',
    data
  })
}
