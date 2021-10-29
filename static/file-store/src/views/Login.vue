<template>
  <div class="login-container">
    <div class="form">
      <el-tabs
        v-model="activeName"
        @tab-click="handleClick"
        tab-position="left"
      >
        <el-tab-pane label="登陆页面" name="first">
          <div style="margin-top: 110px">
            <el-form
              :model="ruleForm"
              status-icon
              :rules="rules"
              ref="ruleForm"
              label-width="100px"
              class="demo-ruleForm"
            >
              <el-form-item label="账号" prop="loginAccount">
                <el-input
                  v-model="ruleForm.loginAccount"
                  placeholder="请输入手机号/账户名"
                  @change="getName"
                >
                </el-input>
              </el-form-item>
              <el-form-item label="密码" prop="loginPassword">
                <el-input
                  type="password"
                  placeholder="请输入密码"
                  v-model="ruleForm.loginPassword"
                  @change="getPassword"
                  class="input-with-select"
                >
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="login">登陆</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
        <el-tab-pane label="注册页面" name="second">
          <div style="margin-top: 60px">
            <el-form
              :model="ruleForm"
              status-icon
              :rules="rules"
              ref="ruleForm"
              label-width="100px"
              class="demo-ruleForm"
            >
              <el-form-item label="账户名" prop="name">
                <el-input
                  v-model.number="ruleForm.name"
                  placeholder="请输入账户名"
                  @change="getName"
                >
                </el-input>
              </el-form-item>
              <el-form-item label="密码" prop="pass">
                <el-input
                  type="password"
                  v-model="ruleForm.pass"
                  placeholder="请输入密码"
                  @change="getPassword"
                ></el-input>
              </el-form-item>
              <el-form-item label="确认密码" prop="checkPass">
                <el-input
                  type="password"
                  v-model="ruleForm.checkPass"
                  placeholder="请输入密码"
                  @change="getSecondPassword"
                ></el-input>
              </el-form-item>
              <el-form-item label="手机号" prop="phone">
                <el-input
                  v-model.number="ruleForm.phone"
                  placeholder="请输入手机号"
                  @change="getPhone"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="register">注册</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import { ElMessage } from "element-plus";
import { UserRegiste, UserLogin } from "@/axios/api";
export default {
  data() {
    var checkName = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入用户名"));
      } else if (value.length < 2 && value.length > 20) {
        callback(new Error("用户名格式不正确"));
      } else {
        callback();
      }
    };
    var checkPhone = (rule, value, callback) => {
      if (value == "") {
        callback(new Error("请输入手机号"));
      }
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
      account: "",
      password: "",
      secondPassword: "",
      phone: "",
      activeName: "first",
      ruleForm: {
        pass: "",
        checkPass: "",
        name: "",
        phone: "",
        loginAccount: "",
        loginPassword: "",
      },
      rules: {
        pass: [
          {
            validator: validatePass,
            trigger: "blur",
          },
        ],
        loginPassword: [
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
        name: [
          {
            validator: checkName,
            trigger: "blur",
          },
        ],
        loginAccount: [
          {
            validator: checkName,
            trigger: "blur",
          },
        ],
        phone: [
          {
            validator: checkPhone,
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    // 登陆用户
    login(formName) {
      if (this.account != "" && this.password != "") {
        var data = {
          user_name: this.account,
          password: this.password,
        };
        UserLogin(data)
          .then((res) => {
            localStorage.setItem("id", res.data.data.UserId);
            localStorage.setItem("user_name", res.data.data.UserName);
            localStorage.setItem("avatar", res.data.data.Avatar);
            localStorage.setItem("token", res.data.data.Token);
            ElMessage.success({
              message: "登陆成功",
              duration: 800,
            })
            this.$router.push({
              path: "/home/all",
            });
          })
          .catch((err) => {
            this.$message.error("登陆失败，请重试", 200);
          });
      } else {
        this.$message.error("请输入账户信息", 200);
      }
    },
    // 注册用户
    register() {
      if (this.account != "" && this.password != "" && this.phone != "") {
        var data = {
          user_name: this.account,
          password: this.password,
          phone: this.phone,
        };
        UserRegiste(data)
          .then((res) => {
            this.$message.success("注册成功，请登陆账户");
            this.activeName = "first";
          })
          .catch((err) => {
            this.$message.error("注册失败，请重试");
          });
      } else {
        this.$message.error("内容不能为空");
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    handleClick(tab, event) {
      // console.log(tab, event);
    },
    getName(val) {
      this.account = val;
    },
    getPassword(val) {
      this.password = val;
    },
    getSecondPassword(val) {
      this.secondPassword = val;
    },
    getPhone(val) {
      this.phone = val;
    },
  },
};
</script>

<style>
.login-container {
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.form {
  margin-top: 15%;
  width: 470px;
  height: 400px;
  display: flex;
  background-color: #fafafa;
  border-radius: 10px;
}

.el-select .el-input {
  width: 200px;
}

.demo-input-label {
  display: inline-block;
  width: 130px;
}
</style>
