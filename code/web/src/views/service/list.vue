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

      <el-table-column label="状态检测">
        <template slot-scope="scope">
          <el-switch v-model="scope.row.is_open_check"></el-switch>
        </template>
      </el-table-column>

      <el-table-column prop="desc" label="描述"> </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleView(scope.$index, scope.row)"
            type="primary"
            style="margin-left: 16px"
          >
            查看
          </el-button>

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

    <el-drawer
      title=""
      :visible.sync="drawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <p>应用名称：web1</p>
      <p>部署服务器：192.168.1.1</p>
      <p>应用类型：前后端项目</p>
    </el-drawer>
  </div>
</template>

<script>
import { getList } from "@/api/table";

export default {
  data() {
    return {
      tableData: [],
      listLoading: true,
      drawer: false,
      drawerData: {},
      direction: "rtl",
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
    handleView(index, row) {
      this.drawer = true;
      this.drawerData = this.tableData[index];
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
    handleDelete(index, row) {
      console.log(index, row);
    },
  },
};
</script>
