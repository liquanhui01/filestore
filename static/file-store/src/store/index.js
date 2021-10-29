import {
  createStore
} from "vuex";

export default createStore({
  // 变量
  state: {
    tableData: [],
    folderCount: 0,
    changeIndex: 1,
    fileData: {},
    folderFiles: [],
    folderFilesCount: 0,
    trashData: [],
    trashDataCount: 0,
    downloadFileSha1: "",
    radioId: 0,
  },
  mutations: {
    // 添加folder数据到数组中
    insertElement(state, el) {
      state.tableData.push(el);
      state.folderCount++;
    },
    // 根据index删除folder数组红的指定元素
    deleteElement(state, index) {
      state.tableData.splice(index, 1)
      state.folderCount--;
    },
    // 获取用户所有的文件夹数据并赋值给数组
    getAllFolders(state, folders) {
      state.tableData = folders;
      state.folderCount = folders.length;
    },
    // 修改指定下标元素的内容
    EditElement(state, param) {
      state.tableData.splice(param["index"], 1, param["el"]);
    },
    // 改变index修改header上的展示内容
    EditIndex(state, index) {
      state.changeIndex = index;
    },
    // 修改fileData的值
    EditFileData(state, param) {
      state.fileData = param;
    },
    // 获取文件夹下所有的文件
    GetAllFolderFiles(state, param) {
      state.folderFiles = param;
      state.folderFilesCount = param.length;
    },
    // 修改指定下标元素的内容
    EditFolderFileElement(state, param) {
      state.folderFiles.splice(param["index"], 1, param["el"]);
    },
    // 根据index删除folderFiles数组红的指定元素
    DeleteFolderFilesElement(state, index) {
      state.folderFiles.splice(index, 1)
      state.folderFilesCount--;
    },
    // 获取垃圾箱下所有的文件
    GetAllTrashFiles(state, param) {
      state.trashData = param;
      state.trashDataCount = param.length;
    },
    // 修改指定下标元素的垃圾箱内容
    EditTrashFileElement(state, param) {
      state.trashData.splice(param["index"], 1, param["el"]);
    },
    // 根据index删除trashData数组红的指定元素
    DeleteTrashFilesElement(state, index) {
      state.trashData.splice(index, 1)
      state.trashDataCount--;
    },
    // 修改下载的文件的id
    EditDownloadFileSha1(state, param) {
      state.downloadFileSha1 = param
    },
    // 修改radioId的值
    EditRadioId(state, value) {
      state.radioId = value;
    },
  },
  actions: {
    // 添加folder数据到数组中
    addTableData(context, payload) {
      context.commit("getAllFolders", payload);
    },
    // 根据index删除folder数组红的指定元素
    removeTableElement(context, payload) {
      context.commit("deleteElement", payload);
    },
    // 获取用户所有的文件夹数据并赋值给数组
    addTableElement(context, payload) {
      context.commit("insertElement", payload);
    },
    // 修改数组指定下标元素数据
    editTableElement(context, payload) {
      context.commit("EditElement", payload);
    },
    // 改变index修改header上的展示内容
    editHeaderIndex(context, payload) {
      context.commit("EditIndex", payload);
    },
    // 修改fileData
    editFileDataValue(context, payload) {
      context.commit("EditFileData", payload);
    },
    // 获取folder下所有的文件
    addFolderFiles(context, payload) {
      context.commit("GetAllFolderFiles", payload);
    },
    // 修改folderFiles中指定的元素
    editFolderFilesElement(context, payload) {
      context.commit("EditFolderFileElement", payload);
    },
    // 删除folderFiles中指定的元素
    deleteFolderFilesElement(context, payload) {
      context.commit("DeleteFolderFilesElement", payload);
    },
    // 获取folder下所有的文件
    addTrashFiles(context, payload) {
      context.commit("GetAllTrashFiles", payload);
    },
    // 修改folderFiles中指定的元素
    editTrashFilesElement(context, payload) {
      context.commit("EditTrashFileElement", payload);
    },
    // 删除folderFiles中指定的元素
    deleteTrashFilesElement(context, payload) {
      context.commit("DeleteTrashFilesElement", payload);
    },
    // 修改下载的文件的id
    editDownloadFileSha1(context, payload) {
      context.commit("EditDownloadFileSha1", payload);
    },
    // 修改radioId的值
    editRadioId(context, payload) {
      context.commit("EditRadioId", payload);
    },
  },
  modules: {},
})