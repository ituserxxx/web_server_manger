<template>
  <div class="app-container">
    <el-table
      :data="tableData"
      v-loading="listLoading"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column prop="id" label="ID" width="180"> </el-table-column>
      <el-table-column prop="ip" label="ip" width="180"> </el-table-column>
      <el-table-column prop="ssh_user" label="ssh_user" width="180"> </el-table-column>
      <el-table-column prop="ssh_passwd" label="ssh_passwd" width="180">
      </el-table-column>
      <el-table-column prop="desc" label="描述"> </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button
          >
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-drawer
      title=""
      :visible.sync="editDrawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <el-form ref="form" :model="editForm" label-width="120px">
        <el-input size="medium" v-model="editForm.id" style="hidden" />
        <el-form-item label="ip">
          <el-input size="medium" v-model="editForm.ip" />
        </el-form-item>
        <el-form-item label="ssh 账号">
          <el-input size="medium" v-model="editForm.ssh_user" />
        </el-form-item>
        <el-form-item label="ssh 密码">
          <el-input size="medium" v-model="editForm.ssh_passwd" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="editForm.desc" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">保存</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
import { getList } from "@/api/table";

export default {
  data() {
    return {
      tableData: null,
      listLoading: true,
      editDrawer: false,
      direction: "rtl",
      editForm: {
        id: "",
        ip: "",
        ssh_user: "",
        ssh_passwd: "",
        desc: "",
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      getList().then((response) => {
        this.tableData = [
          {
            id: 1,
            ip: "127.0.0.1",
            ssh_user: "root",
            ssh_passwd: "123456",
            desc: "上海市普陀区金沙江路 1518 弄",
          },
        ];
        this.listLoading = false;
      });
    },
    handleEdit(index, row) {
      console.log(index, row);
      this.editDrawer = true;
      this.editForm = this.tableData[index];
    },
    onSubmit() {
      console.log(this.editForm);
      this.$message("submit!");
    },
    handleDelete(index, row) {
      console.log(index, row);
    },

    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          this.drawerData = {};
          done();
        })
        .catch((_) => {});
    },
  },
};
</script>
