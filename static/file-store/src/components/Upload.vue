<template>
  <div>
    <!--文件上传入口-->
    <div class="uploadfile">
      <el-upload
        ref="upload"
        class="upload-demo"
        :before-upload="beforeUpload"
        drag
        :auto-upload="false"
        :on-exceed="handleExceed"
      >
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">
          将文件拖到此处，或<em>点击选择文件</em>
        </div>
      </el-upload>
      <el-button
        style="margin-left: 10px"
        size="small"
        type="success"
        @click="submitUpload"
        >上传</el-button
      >
    </div>
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
import { GetUploadPath } from "@/axios/api";
import axios from "axios";
export default {
  name: "UploadFile",
  data() {
    return {
      loading: false,
      percentage: 0,
      tips: "",
      dialogVisible: false,
    };
  },
  methods: {
    beforeUpload(file) {
      let fd = new FormData();
      fd.append("file", file);
      fd.append("user_id", this.$store.state.fileData.user_id);
      fd.append("folder_id", this.$store.state.fileData.folder_id);
      let config = {
        onUploadProgress: (progressEvent) => {
          //progressEvent.loaded:已上传文件大小
          //progressEvent.total:被上传文件的总大小
          let complete =
            (progressEvent.loaded / progressEvent.total).toFixed(2) * 100;
          this.percentage = complete;
          if (this.percentage >= 100) {
            this.dialogVisible = true;
          }
        },
      };
      axios
        .post(GetUploadPath(), fd, config)
        .then(() => {
          this.loading = false;
          this.$notify({
            message: "上传成功",
            duration: 900,
            type: "success",
          });
        })
        .catch(() => {
          this.loading = false;
          this.$notify({
            message: "上传失败，请重新上传",
            duration: 900,
            type: "error",
          });
        });
    },
    handleExceed() {},
    submitUpload() {
      this.loading = true;
      this.tips = "正在上传中。。。";
      this.$refs.upload.submit();
    },
    ensure() {
      this.dialogVisible = false;
      this.loading = false;
    },
  },
};
</script>

<style scoped>
.uploadfile {
  width: 200px;
  height: 200px;
  position: absolute;
  top: 50%;
  left: 50%;
  margin-left: -100px;
  margin-top: -100px;
}
.loading {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  background: black;
  opacity: 0.8;
}
.progress {
  width: 200px;
  height: 200px;
  position: absolute;
  top: 50%;
  left: 50%;
  margin-left: -100px;
  margin-top: -100px;
}
.tips {
  color: #409eff;
  position: absolute;
  top: 50%;
  left: 50%;
  margin-left: -100px;
  margin-top: -150px;
}
</style>
