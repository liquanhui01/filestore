<template>
  <div class="message-total">
    <div class="avatar-total">
      <img :src="avatar" alt="" class="avatar" />
    </div>
    <div class="message-form">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="账户名" class="message-form-item">
          <el-input v-model="form.user_name"></el-input>
        </el-form-item>
        <el-form-item label="手机号" class="form-item">
          <el-input v-model="form.phone"></el-input>
        </el-form-item>
        <el-form-item label="个人简介">
          <el-input type="textarea" v-model="form.profile"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <span style="color: gray">{{ status }}</span>
        </el-form-item>
        <el-form-item label="登陆时间">
          <span style="color: gray">{{ last_active }}</span>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">修改</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { ArrowRight } from "@element-plus/icons";
import { GetUserInfo, UpdateUser } from "@/axios/api";
export default {
  components: {
    ArrowRight,
  },
  data() {
    return {
      form: {
        user_name: "",
        phone: "",
        profile: "",
      },
      avatar:
        "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      last_active: "",
      status: 0,
    };
  },
  mounted() {
    var data = {
      id: localStorage.getItem("id"),
    };
    GetUserInfo(data)
      .then((res) => {
        var data = res.data.data;
        if (data.avatar != "") {
          this.avatar = data.avatar;
        }
        this.last_active = data.last_active;
        var st = data.status;
        this.status =
          st == 0 ? "启用" : st == 1 ? "禁用" : st == 2 ? "锁定" : "删除";
        this.form = {
          user_name: data.user_name,
          phone: data.phone,
          profile: data.profile,
        };
      })
      .catch((err) => {
        console.log("获取失败:", err);
      });
  },
  methods: {
    onSubmit() {
      var data = {
        user_name: this.form.user_name,
        phone: this.form.phone,
        profile: this.form.profile,
        id: localStorage.getItem("id"),
      };
      UpdateUser(data)
        .then((res) => {
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style>
.message-total {
  width: 50%;
  margin-left: 20%;
  margin-top: 7%;
}

.avatar-total {
  height: 160px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding-left: 10px;
  padding-right: 10px;
  color: #303030;
  font-size: 17px;
  font-weight: 500;
}

.avatar {
  width: 85px;
  height: 85px;
  border-radius: 50%;
}

.message-form-item {
  height: 50px;
  line-height: 50px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}
.icons {
  color: #969696;
}
</style>
