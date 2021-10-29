<template>
  <div class="total">
    <!-- 头部 -->
    <el-affix :offset="0">
      <Header v-bind:changeIndex="index"></Header>
    </el-affix>
    <!-- 主体 -->
    <div style="display: flex">
      <!-- 侧边栏 -->
      <el-affix :offset="50">
        <Aside v-on:index="reactive"></Aside>
      </el-affix>
      <!-- 内容区域 -->
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/Header.vue";
import Aside from "@/components/Aside.vue";
export default {
  name: "Home",
  components: {
    Header,
    Aside,
  },
  data() {
	  return {
		  index: 1,
	  }
  },
  methods: {
	reactive(val) {
		console.log(val)
		this.index = val;
	}
  }
};
</script>

<style scoped>
body,
html,
.total {
  width: 100%;
  min-height: 90vh;
}

.content {
  width: calc(100% - 200px);
  min-height: calc(90vh - 50px);
  /* padding: 20px; */
  margin-left: 1;
}
</style>
