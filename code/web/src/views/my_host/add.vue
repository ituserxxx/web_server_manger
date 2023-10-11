<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px" >
      <el-form-item label="主机名称">
        <el-input   size="medium" v-model="form.name" />
      </el-form-item>
      <el-form-item label="ip">
        <el-input   size="medium" v-model="form.ip" />
      </el-form-item>
      <el-form-item label="ssh 账号">
        <el-input   size="medium" v-model="form.ssh_user" />
      </el-form-item>
      <el-form-item label="ssh 密码">
        <el-input   size="medium" v-model="form.ssh_passd" />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即添加</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: '',
        ip: '',
        ssh_user: '',
        ssh_passd: '',
        desc: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>
 
