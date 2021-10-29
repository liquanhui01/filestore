<template>
  <div class="all-file-total">
    <div class="all-file-total-top">
      <div style="margin-left: 30px">全部文件</div>
      <div style="margin-right: 30px">
        已全部加载，共{{ this.$store.state.folderFilesCount }}项
      </div>
    </div>
    <div class="all-file-total-content">
      <el-table
        :data="JSON.parse(JSON.stringify(this.$store.state.folderFiles))"
        style="width: 100%"
      >
        <el-table-column width="50" label="单选">
          <template #default="scope">
            <el-radio
              :label="scope.$index + 1"
              v-model="this.$store.state.radioId"
              @change="change(scope)"
            ></el-radio>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="文件名" width="400">
          <template #default="scope" class="cell">
            <span style="margin-left: 10px" class="folder">{{
              scope.row.file_name
            }}</span>
          </template>
        </el-table-column>
        <el-table-column label="类型" width="170">
          <template #default="scope" class="cell">
            {{ scope.row.type }}
          </template>
        </el-table-column>
        <el-table-column prop="size" label="大小" width="170">
          <template #default="scope" class="cell">
            {{ scope.row.file_size }}
          </template>
        </el-table-column>
        <el-table-column label="修改时间" width="180">
          <template #default="scope" class="cell">
            <span>{{ scope.row.upload_at }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              size="mini"
              type="primary"
              plain
              @click="handleEdit(scope.$index, scope.row)"
              >重命名</el-button
            >
            <el-button
              size="mini"
              type="success"
              plain
              @click="handleMove(scope.$index, scope.row.file_sha1)"
              >移动</el-button
            >
            <el-button
              size="mini"
              type="danger"
              plain
              @click="handleDelete(scope.$index, scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 文件重命名 -->
    <el-dialog title="文件重命名" v-model="dialogFormVisible">
      <el-form>
        <el-form-item label="文件名">
          <el-input
            v-model="file_name"
            autocomplete="off"
            placeholder="文件名不需要带点号和后缀名"
          ></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelEditFolder">取 消</el-button>
          <el-button type="primary" @click="submitEdit">确 定</el-button>
        </span>
      </template>
    </el-dialog>
    <el-dialog title="文件夹" v-model="dialogFolderVisible">
      <el-form>
        <el-form-item label="文件夹名">
          <el-select
            v-model="value"
            placeholder="请选择要移动的文件夹"
            @change="currentSelect"
          >
            <el-option
              v-for="item in folders"
              :key="item.id"
              :label="item.folder_name"
              :value="item.id"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelMoveFolder">取 消</el-button>
          <el-button type="primary" @click="submitMove">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { FolderOpened } from "@element-plus/icons";
import {
  FindUserFolderFiles,
  UpdateUserFileName,
  MoveFileToFolder,
  FindFolderByUserId,
  UpdateFileStatus,
} from "@/axios/api.js";
import { toRaw } from "@vue/reactivity";
export default {
  components: {
    FolderOpened,
  },
  data() {
    return {
      file_name: "",
      origin_file_name: "",
      folder_id: 0,
      index: 0,
      file_sha1: 0,
      file_id: 0,
      el: {},
      dialogFormVisible: false,
      dialogFolderVisible: false,
      folders: [],
      value: "",
    };
  },
  methods: {
    beforeunloadFu(e) {
      var data = toRaw(this.$store.state.fileData);
      FindUserFolderFiles(data)
        .then((res) => {
          this.$store.dispatch("addFolderFiles", res.data.data);
        })
        .catch((err) => console.log(err));
    },
    submitEdit() {
      var data = {
        new_file_name: this.file_name,
        origin_file_name: this.origin_file_name,
        file_sha1: this.file_sha1,
      };
      UpdateUserFileName(data)
        .then((res) => {
          var newData = toRaw(this.el);
          newData["file_name"] = res.data.data.FileName;
          newData["last_update"] = res.data.data.LastUpdate;
          var payload = {
            index: this.index,
            el: newData,
          };
          this.$store.dispatch("editFolderFilesElement", payload);
          this.$notify({
            message: "文件重命名成功",
            type: "success",
            duration: 800,
          });
        })
        .catch((err) => {
          this.$notify({
            message: "文件重命名失败",
            type: "error",
            duration: 800,
          });
        });
      this.dialogFormVisible = false;
    },
    handleDelete(index, val) {
      var data = {
        file_sha1: val.file_sha1,
        user_id: localStorage.getItem("id"),
        status: 1,
        file_name: val.file_name,
        folder_id: val.folder_id,
      };
      this.$confirm("请确认是否删除?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          UpdateFileStatus(data)
            .then(() => {
              this.$store.dispatch("deleteFolderFilesElement", index);
              this.$notify({
                message: "文件删除成功",
                type: "success",
                duration: 800,
              });
            })
            .catch((err) => {
              console.log(err);
              this.$notify({
                message: "文件删除失败",
                type: "error",
                duration: 800,
              });
            });
        })
        .catch(() => {
          this.$notify({
            message: "已取消删除",
            type: "error",
            duration: 800,
          });
        });
    },
    handleEdit(index, val) {
      this.index = index;
      this.file_sha1 = val.file_sha1;
      this.origin_file_name = val.file_name;
      this.el = val;
      this.dialogFormVisible = true;
    },
    cancelEditFolder() {
      this.dialogFormVisible = false;
    },
    handleMove(index, val) {
      this.index = index;
      this.file_sha1 = val;
      FindFolderByUserId(this.$store.state.fileData)
        .then((res) => {
          this.folders = res.data.data;
        })
        .catch((err) => console.log(err));
      this.dialogFolderVisible = true;
    },
    currentSelect(val) {
      this.folder_id = val;
    },
    submitMove() {
      var data = {
        folder_id: this.folder_id,
        file_sha1: this.file_sha1,
      };
      MoveFileToFolder(data)
        .then(() => {
          this.$store.dispatch("deleteFolderFilesElement", this.index);
          this.$notify({
            message: "文件移动成功",
            type: "success",
            duration: 800,
          });
        })
        .catch(() => {
          this.$notify({
            message: "文件移动失败",
            type: "error",
            duration: 800,
          });
        });
      this.dialogFolderVisible = false;
      this.$store.dispatch("editRadioId", 0);
    },
    cancelMoveFolder() {
      this.folders = [];
      this.dialogFolderVisible = false;
    },
    // 单选框绑定值变化时触发的事件（选中的 Radio label 值）
    change(scope) {
      this.$store.dispatch("editRadioId", scope.$index + 1);
      this.$store.dispatch("editDownloadFileSha1", scope.row.file_sha1);
    },
  },
  mounted() {
    var data = {
      user_id: localStorage.getItem("id"),
      folder_id: localStorage.getItem("folder_id"),
    };
    FindUserFolderFiles(data)
      .then((res) => {
        console.log(res.data.data);
        this.$store.dispatch("addFolderFiles", res.data.data);
      })
      .catch((err) => console.log(err));
  },
};
</script>

<style>
.all-file-total {
  width: 100%;
}

.all-file-total-top {
  width: 100%;
  height: 35px;
  background-color: #fafafa;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  color: #969696;
  font-size: 13px;
}

.all-file-total-content {
  width: 100%;
  min-height: calc(97vh - 80px);
}
.pagination {
  width: 100%;
  height: 80px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.cell {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.folder:hover {
  color: #409eff;
}
</style>
