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

      <el-table-column prop="service_type" label="应用类型" width="180"></el-table-column>

      <!-- <el-table-column prop="is_open_check" label="检测开关" width="180"></el-table-column>
       -->

      <el-table-column prop="is_open_check" label="检测开关">
        <el-switch v-model="is_open_check" disabled> </el-switch>
      </el-table-column>

      <el-table-column prop="desc" label="描述"> </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            @click="handleEdit(scope.$index, scope.row)"
            >查看</el-button
          >

          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">
            编辑
          </el-button>

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
            name: "web1",
            service_type: "前后端项目",
            is_open_check: true,
            desc: "xxxxxx",
          },
          {
            id: 2,
            name: "web2",
            service_type: "仅后端服务",
            is_open_check: false,

            desc: "xxxxxx",
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
