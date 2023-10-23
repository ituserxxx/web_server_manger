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
            <div v-if="scope.row.check_list.length > 0">
              <span v-for="item in scope.row.check_list">
                <el-button
                  :type="item.value && item.state ? 'success' : 'warning'"
                  size="mini"
                  plain
                  disabled
                  >{{ item.desc }}</el-button
                >
              </span>
            </div>
            <br />
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
      size="40%"
    >
      <div style="height: 600px; width: 80%; margin: 0 auto; overflow: auto">
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
                :label="item.ip"
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

          <!-- 仓库 -->
          <div>
            仓库 ：
            <el-button
              type="primary"
              size="small"
              icon="el-icon-plus"
              @click="addGit"
            ></el-button>
          </div>
          <div
            class="git-continer"
            v-for="(git, index) in form.git_list"
            :key="index"
            :prop="git.url"
          >
            <div class="git-inner-div">
              <div style="widht: 700px">
                <el-form-item label="url：">
                  <el-input v-model="git.url"></el-input>
                </el-form-item>
                <el-form-item label="备注：">
                  <el-input v-model="git.desc"></el-input>
                </el-form-item>
              </div>
              <el-button @click.prevent="removeGit(git)">删除</el-button>
            </div>
          </div>
          <br />
          <br />

          <el-form-item label="状态检测">
            <el-switch v-model="form.is_open_check" />
          </el-form-item>
          <!-- 检测项 -->

          <div v-if="form.is_open_check">
            <span style="width: 100px">检测项 ：</span>
            <el-button
              type="primary"
              size="small"
              icon="el-icon-plus"
              @click="addCheck"
            ></el-button>
          </div>
          <div
            class="check-continer"
            v-for="(check, index) in form.check_list"
            :key="index"
            :prop="check.url"
          >
            <div class="check-inner-div">
              <el-form-item label="检测方式">
                <el-select v-model="check.type" placeholder="请选择">
                  <el-option label="http" value="http" />
                  <el-option label="tcp" value="tcp" />
                </el-select>
              </el-form-item>
              <el-form-item label="内容">
                <el-input v-model="check.value"></el-input>
              </el-form-item>
              <el-form-item label="间隔时间（秒）">
                <el-input v-model="check.time"></el-input>
              </el-form-item>
              <el-form-item label="备注：">
                <el-input v-model="check.desc"></el-input>
              </el-form-item>
              <el-button @click.prevent="removeCheck(check)">删除</el-button>
            </div>
          </div>
          <br /><br />
        </el-form>
      </div>
      <el-button type="primary" @click="onUpdate">Update</el-button>
      <el-button @click="handleClose">Cancel</el-button>
    </el-drawer>
  </div>
</template>

<script>
import { myserviceGetList, myserviceDel, myserviceUpdate } from "@/api/myservice";
import { myhostGetList } from "@/api/myhost";
export default {
  data() {
    return {
      listLoading: true,
      direction: "rtl",
      drawer: false,
      drawerData: {},
      tableData: [
        // {
        //   id: "yeauq",
        //   name: "111",
        //   desc: "1111",
        //   host_id: "dofoy",
        //   service_type: "qianhou",
        //   deploy_method: "hou",
        //   git_list: [{ url: "http://localhost:9528/#/service/add", desc: "xxxx" }],
        //   is_open_check: true,
        //   check_list: [
        //     {
        //       type: "http",
        //       value: "http://localhost:9528/#/service/add",
        //       time: "22",
        //       desc: "ssss",
        //       state: false,
        //     },
        //   ],
        // },
      ],
      host_list: [],
      form: {
        //   id: "yeauq",
        //   name: "111",
        //   desc: "1111",
        //   host_id: "dofoy",
        //   service_type: "qianhou",
        //   deploy_method: "hou",
        //   git_list: [{ url: "http://localhost:9528/#/service/add", desc: "xxxx" }],
        //   is_open_check: true,
        //   check_list: [
        //     {
        //       type: "http",
        //       value: "http://localhost:9528/#/service/add",
        //       time: "22",
        //       desc: "ssss",
        //       state: false,
        //     },
        //   ],
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    // ---------- 列表
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
      myhostGetList().then((res) => {
        if (res.code == 33) {
          this.host_list = res.data;
        } else {
          this.$message.error(res.data);
        }
      });
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
    // ----------
    // ---------- 编辑
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          this.drawerData = {};
          done();
        })
        .catch((_) => {});
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
    addGit() {
      this.form.git_list.push({
        url: "", // git resp url
        desc: "", // desc
      });
    },
    removeGit(item) {
      var index = this.form.git_list.indexOf(item);
      if (index !== -1) {
        this.form.git_list.splice(index, 1);
      }
    },
    addCheck() {
      this.form.check_list.push({
        type: "", // http tcp
        value: "", // url  ip:port
        time: "", // time interval
        desc: "", // desc
      });
    },
    removeCheck(item) {
      var index = this.form.check_list.indexOf(item);
      if (index !== -1) {
        this.form.check_list.splice(index, 1);
      }
    },
  },
};
</script>
<style scoped>
.line {
  text-align: center;
}

.git-continer {
  width: 80%; /* 设置容器的宽度 */
  height: auto; /* 自适应高度 */
  display: flex; /* 将容器设置为弹性容器 */
  flex-wrap: wrap; /* 允许容器内的元素换行显示 */
}

.git-inner-div {
  width: 80%; /* 设置内层div的宽度 */
  height: auto; /* 自适应高度 */
  margin: 20px; /* 设置内层div之间的间距 */
  padding: 10px; /* 设置内层div的内边距 */
  border: 1px solid #ccc; /* 设置内层div的边框 */
  /*box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); 添加阴影效果 */
  border-radius: 4px; /* 使用圆角边框 */
  position: relative; /* 使子元素的绝对定位生效 */
}

.check-continer {
  width: 80%; /* 设置容器的宽度 */
  height: auto; /* 自适应高度 */
  display: flex; /* 将容器设置为弹性容器 */
  flex-wrap: wrap; /* 允许容器内的元素换行显示 */
}

.check-inner-div {
  width: 80%; /* 设置内层div的宽度 */
  height: auto; /* 自适应高度 */
  margin: 20px; /* 设置内层div之间的间距 */
  padding: 10px; /* 设置内层div的内边距 */
  border: 1px solid #ccc; /* 设置内层div的边框 */
  /*box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); 添加阴影效果 */
  border-radius: 4px; /* 使用圆角边框 */
  position: relative; /* 使子元素的绝对定位生效 */
}
</style>
