<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px" :label-position="labelPosition">
      <el-form-item label="ip">
        <el-input size="medium" v-model="form.ip" />
      </el-form-item>
      <el-form-item label="ssh 账号">
        <el-input size="medium" v-model="form.ssh_user" />
      </el-form-item>
      <el-form-item label="ssh 密码">
        <el-input size="medium" v-model="form.ssh_passwd" />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即添加</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { myhostAdd } from "@/api/myhost";
export default {
  data() {
    return {
      labelPosition: "left",
      form: {
        ip: "",
        ssh_user: "",
        ssh_passwd: "",
        desc: "",
      },
    };
  },
  methods: {
    onSubmit() {
      console.log(this.form);
      myhostAdd(this.form).then((response) => {
        this.list = response.data;
      });
    },
  },
};
</script>
