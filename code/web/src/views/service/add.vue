<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="应用名称">
        <el-input v-model="form.name" />
      </el-form-item>

      <el-form-item label="部署服务器">
        <el-select v-model="value" placeholder="请选择">
          <el-option
            v-for="item in form.host_list"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="应用类型">
        <el-select v-model="form.service_type" placeholder="请选择">
          <el-option label="前后端项目" value="qianhou" />
          <el-option label="仅后端服务" value="hou" />
        </el-select>
      </el-form-item>

      <div v-if="form.service_type === 'qianhou'">
        <el-form-item label="部署方式">
          <el-select v-model="form.deploy_method" placeholder="请选择">
            <el-option label="前后分离" value="qianhou" />
            <el-option label="后端静态" value="hou" />
          </el-select>
        </el-form-item>
      </div>

      <div v-if="form.service_type === 'hou'">
        <p>后端部署信息</p>
        <el-form-item label="git 仓库地址">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="代码位置">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="线上分支">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
      </div>

      <div v-if="form.deploy_method === 'qianhou'">
        <div>
          <p>前端部署信息</p>
          <el-form-item label="git 仓库地址">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="代码位置">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="线上分支">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
        </div>
        <div>
          <p>后端部署信息</p>
          <el-form-item label="git 仓库地址">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="代码位置">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="线上分支">
            <!-- <el-switch v-model="form.delivery" /> -->
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
        </div>
      </div>

      <el-form-item label="状态检测">
        <el-switch v-model="form.is_open_check" />
      </el-form-item>

      <div v-if="form.is_open_check">
        <el-form-item label="前端" v-if="form.service_type !== 'hou'">
          <el-switch v-model="form.is_open_check_qian" />
        </el-form-item>

        <el-form-item label="后端">
          <el-switch v-model="form.is_open_check_hou" />
        </el-form-item>

        <el-form-item label="检测时间间隔(秒)">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
      </div>

      <div v-if="form.is_open_check_qian">
        <p>检测-配置前端</p>
        <el-form-item label="检测外网 Url">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="检测内网 Url">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
      </div>

      <div v-if="form.is_open_check_hou">
        <p>检测-配置后端</p>
        <el-form-item label="检测 ip:port">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="检测外网 Url">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
        <el-form-item label="检测内网 Url">
          <!-- <el-switch v-model="form.delivery" /> -->
          <el-input v-model="form.desc" type="textarea" />
        </el-form-item>
      </div>

      <el-form-item label="项目描述">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        service_type: "",
        is_open_check: false,
        is_open_check_qian: false,
        is_open_check_hou: false,
        deploy_method: "",
        host_list: [
          {
            id: 1,
            name: "王小虎",
            ip: "127.0.0.1",
            ssh_passwd: "123456",
            desc: "上海市普陀区金沙江路 1518 弄",
          },
          {
            id: 2,
            name: "王小虎",
            ip: "127.0.0.1",
            ssh_passwd: "123456",
            desc: "上海市普陀区金沙江路 1517 弄",
          },
        ],
        name: "",
        region: "",
        date1: "",
        date2: "",
        delivery: false,
        type: [],
        resource: "",
        desc: "",
      },
    };
  },
  methods: {
    onSubmit() {
      this.$message("submit!");
    },
    onCancel() {
      this.$message({
        message: "cancel!",
        type: "warning",
      });
    },
  },
};
</script>

<style scoped>
.line {
  text-align: center;
}
</style>
