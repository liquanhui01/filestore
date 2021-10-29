<template>
  <div
    style="
      height: 50px;
      line-height: 50px;
      border-bottom: 1px solid #f5f5f5;
      display: flex;
      background-color: #ffffff;
    "
  >
    <div class="title">
      <img
        src="../../public/tubiao.jpeg"
        alt=""
        style="width: 50px; height: 50px; border-radius: 50%"
      />
      <span style="margin-left: 5px">赛昂云盘</span>
    </div>
    <div style="flex: 1; margin-left: 20px">
      <div class="middle">
        <div class="upAndDown" v-if="this.$store.state.changeIndex == 4">
          <el-tooltip
            class="item"
            effect="light"
            content="上传文件"
            placement="bottom"
          >
            <el-icon :size="23" class="icons" @click="navigateToUpload">
              <Upload />
            </el-icon>
          </el-tooltip>
          <el-tooltip
            class="item"
            effect="light"
            content="下载文件"
            placement="bottom"
          >
            <el-icon :size="23" class="icons" @click="downloadFile">
              <Download />
            </el-icon>
          </el-tooltip>
        </div>
        <div v-if="this.$store.state.changeIndex == 1">
          <el-tooltip
            class="item"
            effect="light"
            content="新建文件夹"
            placement="bottom"
          >
            <el-icon :size="23" class="icons" @click="createFolder">
              <FolderOpened />
            </el-icon>
          </el-tooltip>
        </div>
        <div v-else></div>
      </div>
    </div>
    <div class="right">
      <div
        class="demo-type"
        style="margin-top: 10px; margin-right: 10px; padding-top: 10px"
      >
        <div>
          <el-avatar :src="avatar"></el-avatar>
        </div>
      </div>
      <el-dropdown>
        <span class="el-dropdown-link">
          {{ user_name }}<i class="el-icon-arrow-down el-icon--right"></i>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="getMessage">我的账户</el-dropdown-item>
            <el-dropdown-item @click="changePassword"
              >修改密码</el-dropdown-item
            >
            <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <!-- 文件上传框 -->
    <el-dialog v-model="dialogTableVisible"> </el-dialog>
    <!-- 新建文件夹框 -->
    <el-dialog title="新建文件夹" v-model="dialogFormVisible">
      <el-form>
        <el-form-item label="文件夹名">
          <el-input v-model="folder_name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelCreateFolder">取 消</el-button>
          <el-button type="primary" @click="submitFolder">确 定</el-button>
        </span>
      </template>
    </el-dialog>
    <!--遮罩层-->
    <div class="loading" v-if="loading">
      <h4 class="tips">{{ tips }}</h4>
      <!--进度条-->
      <el-progress
        type="line"
        :percentage="percentage"
        class="progress"
        :show-text="true"
      ></el-progress>
    </div>
  </div>
</template>

<script>
import {
  CopyDocument,
  FolderOpened,
  RefreshLeft,
  Delete,
  Download,
  Upload,
} from "@element-plus/icons";
import sha1 from "js-sha1";
import { CreateFolder, DownloadFile } from "@/axios/api";
export default {
  components: {
    CopyDocument,
    FolderOpened,
    RefreshLeft,
    Delete,
    Download,
    Upload,
  },
  props: {
    changeIndex: {
      type: Number,
      required: true,
    },
  },
  data() {
    var checkAge = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("年龄不能为空"));
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error("请输入数字值"));
        } else {
          if (value < 18) {
            callback(new Error("必须年满18岁"));
          } else {
            callback();
          }
        }
      }, 1000);
    };
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      dialogTableVisible: false,
      dialogFormVisible: false,
      dialogUploadVisible: false,
      folder_name: "",
      user_name: "",
      avatar:
        "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      user_id: "",
      ruleForm: {
        pass: "",
        checkPass: "",
      },
      index: 1,
      rules: {
        pass: [
          {
            validator: validatePass,
            trigger: "blur",
          },
        ],
        checkPass: [
          {
            validator: validatePass2,
            trigger: "blur",
          },
        ],
      },
      fileList: [],
      percentage: 0,
    };
  },
  mounted() {
    this.user_name = localStorage.getItem("user_name");
    if (localStorage.getItem("avatar") != "") {
      this.avatar = localStorage.getItem("avatar");
    }
    this.user_id = localStorage.getItem("user_id");
  },
  updated() {
    this.user_name = localStorage.getItem("user_name");
    if (localStorage.getItem("avatar") != "") {
      this.avatar = localStorage.getItem("avatar");
    }
    this.user_id = localStorage.getItem("user_id");
  },
  methods: {
    navigateToUpload() {
      this.$router.push("/home/upload");
    },
    beforeFileUpload(file) {
      let fd = new FormData();
      fd.append("file", file);
      let config = {
        onUploadProgress: (progressEvent) => {
          let complete =
            (progressEvent.loaded / progressEvent.total).toFixed(2) * 100;
          this.percentage = complete;
          // if (this.percentage >= 100) {
          // }
        },
        headers: {
          "Content-Type": "multipart/form-data",
        },
      };
      this.$axios
        .post(this.url, fd, config)
        .then((res) => {
          console.log(res);
        })
        .catch((err) => console.log(err));
    },
    openDialog() {
      this.dialogTableVisible = true;
    },
    createFolder() {
      this.dialogFormVisible = true;
    },
    cancelCreateFolder() {
      this.folder_name = "";
      this.dialogFormVisible = false;
    },
    submitFolder() {
      var data = {
        folder_name: this.folder_name,
        user_id: localStorage.getItem("id"),
      };
      CreateFolder(data)
        .then((res) => {
          this.$store.commit("insertElement", res.data.data);
        })
        .catch((err) => {
          console.log(err);
        });
      this.dialogFormVisible = false;
      this.folder_name = "";
    },
    logout() {
      localStorage.removeItem("token");
      this.$router.push("/");
    },
    getMessage() {
      this.$router.push("/home/message");
    },
    changePassword() {
      this.$router.push("/home/pass");
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert("submit!");
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    getResult(a) {
      console.log(a);
    },
    downloadFile() {
      var data = {
        file_sha1: this.$store.state.downloadFileSha1,
        user_id: localStorage.getItem("id"),
      };
      this.$confirm("是否确定下载该文件?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then((res) => {
          DownloadFile(data)
            .then((res) => {
              if (!res) {
                return;
              }
              // 根据content-disposition获取文件名
              const disposition = res.headers["content-disposition"];
              let file_name = disposition.substring(
                disposition.indexOf("filename=") + 9,
                disposition.length
              );
              // iso8859-1的字符转换成中文
              file_name = decodeURI(escape(file_name));
              // 去掉双引号
              file_name = file_name.replace(/\"/g, "");
              let url = window.URL.createObjectURL(res.data);
              let link = document.createElement("a");
              link.style.display = "none";
              link.href = url;
              link.setAttribute("download", file_name);
              document.body.appendChild(link);
              link.click();
              window.URL.revokeObjectURL(link.href);
              document.body.removeChild(link);
              this.$message({
                type: "success",
                message: "正在下载中...",
                duration: 600,
              });
            })
            .catch((err) => console.log(err));
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消下载",
          });
        });
      this.$store.dispatch("editDownloadFileSha1", "");
      this.$store.dispatch("editRadioId", 0);
    },
  },
};
</script>

<style>
.title {
  width: 200px;
  text-align: center;
  font-weight: 500;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.middle {
  height: 50px;
  color: lightgray;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.icons {
  margin-left: 20px;
}

.icons :hover {
  color: #409eff;
}

.right {
  width: 200px;
  height: 50px;
  line-height: 50px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.uploadTop {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}
</style>
