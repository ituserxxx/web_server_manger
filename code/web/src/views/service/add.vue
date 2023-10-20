<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px" :label-position="labelPosition">
      <el-form-item label="*应用名称">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="项目描述">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item label="*部署服务器">
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

      <div>
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
      <el-form-item>
        <el-button type="primary" @click="onSubmit()">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { myhostGetList } from "@/api/myhost";
import { myserviceAdd } from "@/api/myservice";
export default {
  data() {
    return {
      labelPosition: "left",
      host_list: [],
      form: {
        name: "",
        desc: "",
        host_id: "",
        service_type: "", //qianhou   hou
        deploy_method: "", //qianhou   hou
        git_list: [
          // {
          //   url: "", // git resp url
          //   desc: "", // desc
          // },
        ],
        is_open_check: false,
        check_list: [
          // {
          //   type: "", // http tcp
          //   value: "", // url  ip:port
          //   time: "", // time interval
          //   desc: "", // desc
          // },
        ],
      },
    };
  },
  created() {
    this.fetchData();
    // this.addCheck();
    // this.addGit();
  },
  methods: {
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
    fetchData() {
      myhostGetList().then((res) => {
        if (res.code == 33) {
          this.host_list = res.data;
        } else {
          this.$message.error(res.data);
        }
      });
    },

    onSubmit() {
      console.log(this.form);
      if (this.form.git_list.length > 0) {
        for (let i = 0; i < this.form.git_list.length; i++) {
          const item = this.form.git_list[i];
          if (item.url.trim() === "") {
            // 数据为空，进行相应的处理逻辑
            this.$message.error("仓库git项不能为空");
            return;
          }
        }
      }
      if (this.form.check_list.length > 0) {
        for (let i = 0; i < this.form.check_list.length; i++) {
          const item = this.form.check_list[i];
          if (
            item.type.trim() === "" ||
            item.value.trim() === "" ||
            item.time.trim() === ""
          ) {
            // 数据为空，进行相应的处理逻辑
            this.$message.error("检测项不能为空");
            return;
          }
        }
      }
      if (this.form.name == "" || this.form.host_id == "") {
        this.$message.error("必填项不能为空");
        return;
      }

      myserviceAdd(this.form).then((res) => {
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
