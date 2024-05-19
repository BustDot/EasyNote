namespace go note.note

include "base.thrift"

struct Note {
    1: i64 note_id
    2: i64 user_id
    3: string user_name
    4: string user_avatar
    5: string title
    6: string content
    7: i64 create_time
    8: i64 update_time
}

struct CreateNoteReq {
    1: i64 user_id
    2: string title
    3: string content
}

struct CreateNoteResp {
    1: base.BaseResp base_resp
}

struct DeleteNoteReq {
    1: i64 note_id
    2: i64 user_id
}

struct DeleteNoteResp {
    1: base.BaseResp base_resp
}

struct UpdateNoteReq {
    1: i64 note_id
    2: i64 user_id
    3: string title
    4: string content
}

struct UpdateNoteResp {
    1: base.BaseResp base_resp
}

struct MGetNoteReq {
    1: list<i64> note_ids
}

struct MGetNoteResp {
    1: list<Note> notes
    2: base.BaseResp base_resp
}

struct QueryNoteReq {
    1: i64 user_id
    2: optional string search_key
    3: i64 offset
    4: i64 limit
}

struct QueryNoteResp {
    1: list<Note> notes
    2: i64 total
    3: base.BaseResp base_resp
}

service NoteService {
    CreateNoteResp CreateNote(1: CreateNoteReq req)
    DeleteNoteResp DeleteNote(1: DeleteNoteReq req)
    UpdateNoteResp UpdateNote(1: UpdateNoteReq req)
    MGetNoteResp MGetNote(1: MGetNoteReq req)
    QueryNoteResp QueryNote(1: QueryNoteReq req)
}