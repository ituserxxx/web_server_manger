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
      <el-table-column prop="name" label="名称" width="180"> </el-table-column>
      <el-table-column prop="ip" label="ip" width="180"> </el-table-column>
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
  </div>
</template>

<script>
import { getList } from "@/api/table";

export default {
  data() {
    return {
      tableData: null,
      listLoading: true,
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
          {
            id: 3,
            name: "王小虎",
            ip: "127.0.0.1",
            ssh_passwd: "123456",
            desc: "上海市普陀区金沙江路 1519 弄",
          },
          {
            id: 4,
            name: "王小虎",
            ip: "127.0.0.1",
            ssh_passwd: "123456",
            desc: "上海市普陀区金沙江路 1516 弄",
          },
        ];
        this.listLoading = false;
      });
    },
    handleEdit(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      console.log(index, row);
    },
  },
};
</script>
