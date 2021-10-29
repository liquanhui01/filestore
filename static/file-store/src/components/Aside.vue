<template>
  <el-menu
    :default-active="$route.path"
    router
    class="el-menu-vertical-demo"
    style="min-height: calc(90vh - 50px); border-right: 1px solid #f5f5f5"
  >
    <el-menu-item index="/home/all" @click="changeIndex(1)">
      <i class="el-icon-s-platform"></i>
      <template #title>所有文件</template>
    </el-menu-item>
    <el-menu-item index="/home/delete" @click="changeIndex(2)">
      <i class="el-icon-delete-solid"></i>
      <template #title>垃圾箱</template>
    </el-menu-item>
    <el-submenu index="/home">
      <template #title>
        <i class="el-icon-s-tools"></i>
        <span>设置</span>
      </template>
      <el-menu-item index="/home/message" @click="changeIndex(3)"
        >个人信息</el-menu-item
      >
      <el-menu-item index="/home/pass" @click="changeIndex(3)"
        >修改密码</el-menu-item
      >
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    changeIndex(value) {
      this.$store.dispatch("editHeaderIndex", value);
    },
  },
};
</script>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
