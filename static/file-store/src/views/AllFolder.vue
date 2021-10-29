<template>
  <div class="all-file-total">
    <div class="all-file-total-top">
      <div style="margin-left: 30px">全部文件</div>
      <div style="margin-right: 30px">
        已全部加载，共{{ this.$store.state.folderCount }}项
      </div>
    </div>
    <div class="all-file-total-content">
      <el-table
        :data="JSON.parse(JSON.stringify(this.$store.state.tableData))"
        style="width: 100%"
      >
        <el-table-column prop="name" label="文件名" width="380">
          <template #default="scope" class="cell">
            <el-icon :size="28">
              <FolderOpened />
            </el-icon>
            <span
              style="margin-left: 10px"
              class="folder"
              @click="navigateToFiles(scope.row.id)"
              >{{ scope.row.folder_name }}</span
            >
          </template>
        </el-table-column>
        <el-table-column prop="size" label="大小" width="380">
          -
        </el-table-column>
        <el-table-column prop="size" label="修改时间">
          <template #default="scope" class="cell">
            <span>{{ scope.row.UpdatedAt }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新建文件夹框 -->
    <el-dialog title="修改文件夹" v-model="dialogFormVisible">
      <el-form>
        <el-form-item label="文件夹名">
          <el-input v-model="folder_name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelEditFolder">取 消</el-button>
          <el-button type="primary" @click="submitEdit">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { FolderOpened } from "@element-plus/icons";
import { FindFolderByUserId, UpdateFolder, DeleteFolder } from "@/axios/api.js";
import { toRaw } from "@vue/reactivity";
export default {
  components: {
    FolderOpened,
  },
  data() {
    return {
      folder_name: "",
      dialogFormVisible: false,
      folder_id: 0,
      index: 0,
      el: {},
    };
  },
  methods: {
    submitEdit() {
      var data = {
        user_id: localStorage.getItem("id"),
        folder_id: this.folder_id,
        folder_name: this.folder_name,
      };
      UpdateFolder(data)
        .then((res) => {
          var newData = toRaw(this.el);
          newData["folder_name"] = res.data.data;
          var payload = {
            index: this.index,
            el: newData,
          };
          this.$store.dispatch("editTableElement", payload);
        })
        .catch((err) => {
          console.log(err);
        });
      this.dialogFormVisible = false;
      this.folder_name = "";
    },
    handleDelete(index, val) {
      var data = {
        user_id: localStorage.getItem("id"),
        folder_id: val.id,
        folder_name: val.folder_name,
      };
      DeleteFolder(data)
        .then((res) => {
          this.$store.dispatch("removeTableElement", index);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    handleEdit(index, val) {
      this.dialogFormVisible = true;
      this.folder_id = val.id;
      this.index = index;
      this.el = val;
    },
    cancelEditFolder() {
      this.dialogFormVisible = false;
    },
    navigateToFiles(val) {
      var data = {
        user_id: localStorage.getItem("id"),
        folder_id: val,
      };
      this.$store.dispatch("editHeaderIndex", 4);
      this.$store.dispatch("editFileDataValue", data);
      localStorage.setItem("folder_id", val);
      this.$router.push({
        path: "/home/all/folder/files",
      });
    },
  },
  created() {
    var data = {
      user_id: localStorage.getItem("id"),
    };
    FindFolderByUserId(data)
      .then((res) => {
        this.$store.dispatch("addTableData", res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
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
