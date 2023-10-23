import request from '@/utils/request'

export function notifyUpdate(data) {
  return request({
    url: '/notify/Update',
    method: 'post',
    data
  })
}
export function notifyGet(data) {
  return request({
    url: '/notify/Get',
    method: 'post',
    data
  })
}
 