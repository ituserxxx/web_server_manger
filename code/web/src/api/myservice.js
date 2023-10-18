import request from '@/utils/request'


export function myserviceAdd(data) {
  return request({
    url: '/myservice/Add',
    method: 'post',
    data
  })
}
export function myserviceDel(data) {
  return request({
    url: '/myservice/Del',
    method: 'post',
    data
  })
}
export function myserviceUpdate(data) {
  return request({
    url: '/myservice/Update',
    method: 'post',
    data
  })
}
export function myserviceGetList(data) {
  return request({
    url: '/myservice/GetList',
    method: 'post',
    data
  })
  /**
   this.tableData = [
          {
            id: 1,
            name: "xxx1",
            desc: "desc",
            host_id: "22",
            service_type: "qianhou", //qianhou   hou
            deploy_method: "qianhou", //qianhou   hou
            git_hou: {
              url: "xxxx",
              local: "/dd",
              branch: "master",
            },
            git_qian: {
              url: "xxxx",
              local: "/dd",
              branch: "master",
            },
            is_open_check: true,
            is_open_check_qian: true,
            is_open_check_hou: true,
            check_time_interval: "111",
            open_check_qian: {
              outer_net: "123.0.0.1",
              internal_net: "127.0.0.1",
              outer_net_state: false,
              internal_net_state: true,
            },
            open_check_hou: {
              ip_port: "127.1.1.1:999",
              outer_net: "123.0.0.1",
              internal_net: "127.0.0.1",
              ip_port_state: false,
              outer_net_state: true,
              internal_net_state: true,
            },
          },
          
        ];
   */
}
