# Chat Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Hệ thống chat hỗ trợ 4 loại channel:

| Type | Mô tả |
|------|-------|
| `global` | Channel chung của toàn workplace |
| `dm` | Direct message giữa 2 người (workplace_id = NULL) |
| `group` | Group chat trong workplace |
| `project` | Channel gắn với một project cụ thể |

## Flow

```
[DM]
  Tìm kiếm user → POST /friendships (gửi kết bạn)
  Chấp nhận → PUT /friendships/{id}/accept
  Tạo DM → POST /channels (type: dm)
  Gửi tin → POST /channels/{id}/messages

[Group / Global]
  POST /channels (type: group hoặc global)
  Thêm member → POST /channels/{id}/members
  Gửi tin → POST /channels/{id}/messages

[Project channel]
  Tự động tạo khi tạo project (type: project)
  Gửi tin → POST /channels/{id}/messages
```

---

## Channel Endpoints

### POST `/channels`

Tạo channel mới.

**Request:**
```json
{
  "workplace_id": 1,
  "name": "general",
  "type": "group"
}
```
> DM: bỏ `workplace_id` và `name`, truyền `type: "dm"` và `member_ids: [user_id]`.

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "workplace_id": 1,
    "name": "general",
    "type": "group",
    "created_by": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/channels`

Lấy danh sách channels của user trong workplace.

**Query params:**
- `workplace_id` (required)
- `type` (optional) — `global` | `dm` | `group` | `project`

**Response (200):**
```json
{
  "data": [
    { "id": 1, "name": "general", "type": "group", "unread_count": 3 }
  ]
}
```

---

### GET `/channels/{id}`

Lấy chi tiết channel.

---

### POST `/channels/{id}/members`

Thêm member vào channel.

**Request:**
```json
{ "user_ids": [2, 3] }
```

---

### DELETE `/channels/{id}/members/{user_id}`

Xóa member khỏi channel.

---

## Message Endpoints

### POST `/channels/{id}/messages`

Gửi tin nhắn.

**Request:**
```json
{
  "content": "Hello everyone!",
  "parent_id": 0
}
```
> `parent_id` > 0 = thread reply.

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "channel_id": 1,
    "sender_id": 1,
    "parent_id": 0,
    "content": "Hello everyone!",
    "is_edited": false,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/channels/{id}/messages`

Lấy tin nhắn (phân trang, mới nhất trước).

**Query params:**
- `limit` (default: 50)
- `before_id` — cursor pagination

---

### PUT `/messages/{id}`

Sửa tin nhắn.

**Request:** `{ "content": "Updated content" }`

---

### DELETE `/messages/{id}`

Xóa tin nhắn (soft delete — `is_deleted: true`).

---

### POST `/messages/{id}/reactions`

Thêm reaction.

**Request:** `{ "emoji": "👍" }`

---

### DELETE `/messages/{id}/reactions/{emoji}`

Xóa reaction.

---

## Friendship Endpoints

Cần là bạn bè trước khi DM.

### POST `/friendships`

Gửi lời mời kết bạn.

**Request:** `{ "receiver_id": 2 }`

**Response (201):** `{ "data": { "id": 1, "status": "pending" } }`

---

### PUT `/friendships/{id}/accept`

Chấp nhận lời mời.

---

### PUT `/friendships/{id}/reject`

Từ chối lời mời.

---

### GET `/friendships`

Danh sách bạn bè.

**Query params:** `status` = `pending` | `accepted` | `blocked`

---

## Real-time (WebSocket)

```
ws://localhost:8080/ws?token=<jwt>
```

**Nhận events:**
```json
{ "event": "new_message",   "data": { ...message } }
{ "event": "message_edited", "data": { ...message } }
{ "event": "message_deleted","data": { "id": 1 } }
{ "event": "new_reaction",  "data": { ...reaction } }
{ "event": "user_online",   "data": { "user_id": 1 } }
{ "event": "user_offline",  "data": { "user_id": 1 } }
```

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `channels` | Global/DM/Group/Project channels |
| `channel_members` | Thành viên + last_read_at for unread count |
| `messages` | Tin nhắn + thread support |
| `message_attachments` | File đính kèm (ref → files) |
| `message_reactions` | Emoji reactions |
| `friendships` | Quản lý friend request/accept |
