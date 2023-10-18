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

      <el-table-column prop="service_type" label="应用类型" width="180">
        <template slot-scope="scope">
          <span v-if="scope.row.service_type === 'qianhou'">前后端</span>
          <span v-if="scope.row.service_type === 'hou'">后端</span>
        </template>
      </el-table-column>
      <el-table-column label="部署方式" width="180">
        <template slot-scope="scope">
          <span v-if="scope.row.deploy_method === 'qianhou'">前后分离</span>
          <span v-if="scope.row.deploy_method === 'hou'">后端静态</span>
        </template>
      </el-table-column>

      <!-- <el-table-column prop="is_open_check" label="检测开关" width="180"></el-table-column>
       -->

      <el-table-column label="状态检测">
        <template slot-scope="scope">
          <span v-if="scope.row.is_open_check">
            <span v-if="scope.row.is_open_check_hou">
              <el-button
                :type="
                  scope.row.open_check_hou.outer_net &&
                  scope.row.open_check_hou.outer_net_state
                    ? 'success'
                    : 'warning'
                "
                size="mini"
                plain
                disabled
                >后端(外)</el-button
              >
              <el-button
                :type="
                  scope.row.open_check_hou.internal_net &&
                  scope.row.open_check_hou.internal_net_state
                    ? 'success'
                    : 'warning'
                "
                size="mini"
                plain
                disabled
                >后端(内)</el-button
              >
              <el-button
                v-if="scope.row.open_check_hou.ip_port"
                :type="scope.row.open_check_hou.ip_port_state ? 'success' : 'warning'"
                size="mini"
                plain
                disabled
                >IP:PORT</el-button
              >
            </span>
            <br />
            <span v-if="scope.row.is_open_check_qian">
              <el-button
                :type="
                  scope.row.open_check_qian.outer_net &&
                  scope.row.open_check_qian.outer_net_state
                    ? 'success'
                    : 'warning'
                "
                size="mini"
                plain
                disabled
                >前端(外)</el-button
              >
              <el-button
                :type="
                  scope.row.open_check_qian.internal_net &&
                  scope.row.open_check_qian.internal_net_state
                    ? 'success'
                    : 'warning'
                "
                size="mini"
                plain
                disabled
                >前端(内)</el-button
              >
            </span>
          </span>
          <span v-else>
            <el-button type="info" size="mini" plain disabled>未开启检测</el-button>
          </span>
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
      <div style="height: 600px; width: 90%; margin: 0 auto; overflow: auto">
        <el-form ref="form" :model="form" label-width="120px">
          <el-form-item label="应用名称">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item label="项目描述">
            <el-input v-model="form.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="部署服务器">
            <el-select v-model="form.host_id" placeholder="请选择">
              <el-option
                v-for="item in host_list"
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

          <div v-if="form.service_type === 'hou' || form.deploy_method === 'hou'">
            <p>后端部署信息</p>
            <el-form-item label="git 仓库地址">
              <el-input v-model="form.git_hou.url" type="textarea" />
            </el-form-item>
            <el-form-item label="代码位置">
              <el-input v-model="form.git_hou.local" type="textarea" />
            </el-form-item>
            <el-form-item label="线上分支">
              <el-input v-model="form.git_hou.branch" type="textarea" />
            </el-form-item>
          </div>

          <div v-if="form.deploy_method === 'qianhou'">
            <div>
              <p>前端部署信息</p>
              <el-form-item label="git 仓库地址">
                <el-input v-model="form.git_qian.url" type="textarea" />
              </el-form-item>
              <el-form-item label="代码位置">
                <el-input v-model="form.git_qian.local" type="textarea" />
              </el-form-item>
              <el-form-item label="线上分支">
                <el-input v-model="form.git_qian.branch" type="textarea" />
              </el-form-item>
            </div>
            <div>
              <p>后端部署信息</p>
              <el-form-item label="git 仓库地址">
                <el-input v-model="form.git_hou.url" type="textarea" />
              </el-form-item>
              <el-form-item label="代码位置">
                <el-input v-model="form.git_hou.local" type="textarea" />
              </el-form-item>
              <el-form-item label="线上分支">
                <el-input v-model="form.git_hou.branch" type="textarea" />
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
              <el-input v-model="form.check_time_interval" type="textarea" />
            </el-form-item>
          </div>

          <div v-if="form.is_open_check_qian">
            <p>检测-配置前端</p>
            <el-form-item label="检测外网 Url">
              <el-input v-model="form.open_check_qian.outer_net" type="textarea" />
            </el-form-item>
            <el-form-item label="检测内网 Url">
              <el-input v-model="form.open_check_qian.internal_net" type="textarea" />
            </el-form-item>
          </div>

          <div v-if="form.is_open_check_hou">
            <p>检测-配置后端</p>
            <el-form-item label="检测 ip:port">
              <el-input v-model="form.open_check_hou.ip_port" type="textarea" />
            </el-form-item>
            <el-form-item label="检测外网 Url">
              <el-input v-model="form.open_check_hou.outer_net" type="textarea" />
            </el-form-item>
            <el-form-item label="检测内网 Url">
              <el-input v-model="form.open_check_hou.internal_net" type="textarea" />
            </el-form-item>
          </div>
        </el-form>
      </div>
      <el-button type="primary" @click="onUpdate">Update</el-button>
      <el-button @click="handleClose">Cancel</el-button>
    </el-drawer>
  </div>
</template>

<script>
import { myserviceGetList, myserviceDel, myserviceUpdate } from "@/api/myservice";

export default {
  data() {
    return {
      listLoading: true,
      direction: "rtl",
      drawer: false,
      drawerData: {},
      tableData: [],
      form: {},
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      myserviceGetList().then((res) => {
        this.listLoading = false;
        if (res.code === 33) {
          this.tableData = res.data;
        } else {
          this.$message.error(res.data);
        }
      });
    },

    handleView(index, row) {
      console.log(index, row);
      this.drawer = true;
      this.form = this.tableData[index];
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
      myserviceDel({ id: row.id }).then((res) => {
        if (res.code == 33) {
          this.tableData.splice(index, 1);
          this.$message("handleDelete ok");
        } else {
          this.$message.error(res.data);
        }
      });
    },
    onUpdate() {
      if (this.form.name == "" || this.form.host_id == "") {
        this.$message.error("必填项不能为空");
        return;
      }
      myserviceUpdate(this.form).then((res) => {
        if (res.code == 33) {
          this.$message("onUpdate ok");
        } else {
          this.$message.error(res.data);
        }
      });
    },
  },
};
</script>
