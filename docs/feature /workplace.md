# Workplace Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Workplace là không gian làm việc của một nhóm. Mỗi user có thể tạo hoặc tham gia nhiều workplace.

```
User
 ├── Tạo workplace → trở thành admin
 ├── Join workplace bằng invite link/token
 └── Tham gia với role: admin | member
```

Mỗi workplace có:
- Các **project** với Kanban board
- Các **channel** (global, group, project)
- **Members** với role: `admin` | `member`

---

## Endpoints

### POST `/workplaces`

Tạo workspace mới.

**Request:**
```json
{
  "name": "My Team",
  "icon_url": "https://...",
  "size": "11-50"
}
```

`size` = `1-10` | `11-50` | `51-200` | `201-500` | `500+`

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "name": "My Team",
    "icon_url": "https://...",
    "size": "11-50",
    "created_by": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/workplaces`

Lấy danh sách workplaces của user hiện tại.

**Response (200):**
```json
{
  "data": [
    { "id": 1, "name": "My Team", "role": "admin", "joined_at": "2024-01-01T00:00:00Z" }
  ]
}
```

---

### GET `/workplaces/{id}`

Lấy chi tiết workplace.

---

### PUT `/workplaces/{id}`

Cập nhật workplace (chỉ admin).

**Request:**
```json
{ "name": "Updated Name", "icon_url": "https://...", "size": "51-200" }
```

---

### DELETE `/workplaces/{id}`

Xóa workplace (chỉ admin tạo).

---

## Invitation Endpoints

### POST `/workplaces/{id}/invitations`

Tạo invite link.

**Request:**
```json
{ "max_uses": 10, "expires_at": "2024-02-01T00:00:00Z" }
```

> Bỏ `max_uses` hoặc `expires_at` nếu không giới hạn.

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "invite_token": "abc123xyz",
    "invite_url": "https://app.example.com/join/abc123xyz",
    "max_uses": 10,
    "use_count": 0,
    "expires_at": "2024-02-01T00:00:00Z"
  }
}
```

---

### POST `/workplaces/join`

Tham gia workplace bằng invite token.

**Request:**
```json
{ "invite_token": "abc123xyz" }
```

**Response (200):**
```json
{
  "data": {
    "workplace_id": 1,
    "user_id": 5,
    "role": "member",
    "joined_at": "2024-01-01T00:00:00Z"
  }
}
```

**Error cases:**
| Status | Reason |
|--------|--------|
| 404 | Token không tồn tại |
| 410 | Token đã hết hạn hoặc hết lượt dùng |
| 409 | User đã là member |

---

### DELETE `/workplaces/{id}/invitations/{invitation_id}`

Revoke invite link.

---

## Member Endpoints

### GET `/workplaces/{id}/members`

Lấy danh sách thành viên.

**Response (200):**
```json
{
  "data": [
    { "user_id": 1, "name": "John Doe", "email": "john@example.com", "role": "admin", "joined_at": "..." }
  ]
}
```

---

### PUT `/workplaces/{id}/members/{user_id}/role`

Đổi role thành viên (chỉ admin).

**Request:** `{ "role": "admin" }` hoặc `{ "role": "member" }`

---

### DELETE `/workplaces/{id}/members/{user_id}`

Xóa thành viên khỏi workplace (chỉ admin).

---

### DELETE `/workplaces/{id}/leave`

Tự rời khỏi workplace.

> Admin không thể rời nếu là admin duy nhất — phải chuyển quyền trước.

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `workplaces` | Thông tin workspace |
| `workplace_members` | Thành viên + role |
| `workplace_invitations` | Invite links |
