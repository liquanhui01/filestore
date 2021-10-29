<template>
  <div class="delete-file-total">
    <div class="delete-file-total-top">
      <div style="margin-left: 30px">全部文件</div>
      <div style="margin-right: 30px">
        已全部加载，共{{ this.$store.state.trashDataCount }}项
      </div>
    </div>
    <div class="delete-file-total-content">
      <el-table :data="this.$store.state.trashData" style="width: 100%">
        <el-table-column prop="name" label="文件名" width="400">
          <template #default="scope" class="cell">
            <span style="margin-left: 10px" class="folder">{{
              scope.row.name
            }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="type" label="文件类型" width="280">
          <template #default="scope" class="cell">
            <span style="margin-left: 10px" class="folder">{{
              scope.row.type == 0 ? "文件夹" : "文件"
            }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="date" label="删除日期" width="280">
          <template #default="scope" class="cell">
            <span style="margin-left: 10px" class="folder">{{
              scope.row.CreatedAt
            }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              size="mini"
              type="success"
              plain
              @click="handleRestore(scope.$index, scope.row)"
              >还原</el-button
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
  </div>
</template>

<script>
import { FindTrashFiles, DeleteTrashFile, RestoreTrashFile } from "@/axios/api";
export default {
  data() {
    return {};
  },
  methods: {
    handleRestore(index, val) {
      var data = {
        user_id: val.user_id,
        type: val.type,
        folder_id: val.folder_id,
        file_sha1: val.file_sha1,
        id: val.id,
      };
      RestoreTrashFile(data)
        .then((res) => {
          this.$store.dispatch("deleteTrashFilesElement", index);
        })
        .catch((err) => console.log(err));
    },
    handleDelete(index, val) {
      var data = {
        id: val.id,
        type: val.type,
        file_sha1: val.file_sha1,
        user_id: val.user_id,
        folder_id: val.folder_id,
      };
      DeleteTrashFile(data)
        .then((res) => {
          this.$store.dispatch("deleteTrashFilesElement", index);
        })
        .catch((err) => console.log(err));
    },
  },
  mounted() {
    var data = {
      user_id: localStorage.getItem("id"),
    };
    FindTrashFiles(data)
      .then((res) => {
        this.$store.dispatch("addTrashFiles", res.data.data);
      })
      .catch((err) => console.log(err));
  },
};
</script>

<style>
.delete-file-total {
  width: 100%;
}

.delete-file-total-top {
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

.delete-file-total-content {
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
</style>
