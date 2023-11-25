namespace go edit

struct BaseResponse {
    1: i64 status_code
    2: string status_msg
}

struct FileNode {
    1: string name,
    2: bool is_dir,
    3: list<FileNode> children
}

struct GetDirTreeResponse {
    1: BaseResponse base_resp
    2: list<FileNode> dir_tree
}

struct GetFileRequest {
    1: string path
}

struct GetFileResponse {
    1: BaseResponse base_resp
    2: string file_content
}

struct CreateFileRequest {
    1: string path
}

struct CreateFileResponse {
    1: BaseResponse base_resp
}

struct CreateDirRequest {
    1: string path
}

struct CreateDirResponse {
    1: BaseResponse base_resp
}

service EditService {
    GetDirTreeResponse GetDirTree()
    GetFileResponse GetFile(1: GetFileRequest req)
    CreateFileResponse CreateFile(1: CreateFileRequest req)
    CreateDirResponse CreateDir(1: CreateDirRequest req)
}