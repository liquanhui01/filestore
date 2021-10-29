import {
    post,
    get,
    put,
    del,
    download,
} from "./request";

const UserRegisterPath = '/api/v1/user/register'
const UserLoginPath = '/api/v1/user/login'
const GetUserInfoPath = "/api/v1/token/user/find"
const UpdateUserPasswordPath = "/api/v1/token/user/password"
const UpdateUserPath = "/api/v1/token/user/update"
const UserFolder = "/api/v1/token/folder/"
const UserFolderFiles = "/api/v1/token/folder/files"
const MoveFileFolder = "/api/v1/token/folder/position"
const UserFileStatus = "/api/v1/token/folder/files/status"
const UserTrash = "/api/v1/token/trash/"
const UserTrashFile = "/api/v1/token/trash/file"
const UserTrashAll = "/api/v1/token/trash/all"
const UserTrashRestore = "/api/v1/token/trash/restore"
const UserFileDownload = "/api/v1/token/file/download"
const UserFileUpload = "/api/v1/token/file/upload"
const UserFileDownloadOSS = "/api/v1/token/file/downloadurl"


// 用户注册
export function UserRegiste(data) {
    return post(UserRegisterPath, data)
}

// 用户登陆
export function UserLogin(data) {
    return post(UserLoginPath, data)
}

// 获取用户信息
export function GetUserInfo(data) {
    return get(GetUserInfoPath, data)
}

// 更新用户密码
export function UpdateUserPassword(data) {
    return put(UpdateUserPasswordPath, data)
}

// 更新所有数据
export function UpdateUser(data) {
    return put(UpdateUserPath, data)
}

// 创建文件夹
export function CreateFolder(data) {
    return post(UserFolder, data)
}

// 查询用户所有的文件夹
export function FindFolderByUserId(data) {
    return get(UserFolder, data);
}

// 修改文件夹
export function UpdateFolder(data) {
    return put(UserFolder, data);
}

// 删除文件夹
export function DeleteFolder(data) {
    return del(UserFolder, data);
}

// 查询用户指定文件夹下所有的文件信息
export function FindUserFolderFiles(data) {
    return get(UserFolderFiles, data);
}

// 文件重命名
export function UpdateUserFileName(data) {
    return put(UserFolderFiles, data);
}

// 移动文件到指定的文件夹
export function MoveFileToFolder(data) {
    return put(MoveFileFolder, data);
}

// 修改文件的状态
export function UpdateFileStatus(data) {
    return put(UserFileStatus, data);
}

// 创建垃圾箱文件或文件夹
export function CreateTrashFile(data) {
    return post(UserTrash, data);
}

// 查找所有的垃圾箱内容
export function FindTrashFiles(data) {
    return get(UserTrash, data);
}

// 删除垃圾箱中指定的文件
export function DeleteTrashFile(data) {
    return del(UserTrashFile, data);
}

// 清空垃圾箱文件或文件夹
export function DeleteAllTrashFile(data) {
    return del(UserTrashAll, data);
}

// 垃圾箱文件或文件夹还原
export function RestoreTrashFile(data) {
    return del(UserTrashRestore, data);
}

//文件下载
export function DownloadFile(data) {
    return download(UserFileDownload, data);
}

// 返回文件上传路径
export function GetUploadPath() {
    return UserFileUpload
}