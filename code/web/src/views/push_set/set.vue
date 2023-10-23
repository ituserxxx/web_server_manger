<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :label-position="labelPosition">
      <el-form-item label="企业微信通知开关">
        <el-switch v-model="form.qywx_is_open" />
      </el-form-item>

      <el-form-item label="企业微信通知 webhook">
        <el-input size="medium" v-model="form.qywx_webhook" />
      </el-form-item>
      <el-form-item label="钉钉通知开关">
        <el-switch v-model="form.dd_is_open" />
      </el-form-item>
      <el-form-item label="钉钉通知 webhook">
        <el-input size="medium" v-model="form.dd_webhook" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { notifyUpdate, notifyGet } from "@/api/notify";
export default {
  data() {
    return {
      labelPosition: "left",
      form: {
        id: "",
        qywx_is_open: false,
        qywx_webhook: "",
        dd_is_open: false,
        dd_webhook: "",
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      notifyGet().then((res) => {
        if (res.code == 33) {
          this.form = res.data;
        } else {
          this.$message.error(res.data);
        }
      });
    },
    onSubmit() {
      console.log(this.form);
      notifyUpdate(this.form).then((res) => {
        if (res.code == 33) {
          this.$message("submit!");
        } else {
          this.$message.error(res.data);
        }
      });
    },
  },
};
</script>
