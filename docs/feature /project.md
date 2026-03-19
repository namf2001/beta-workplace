# Project Management Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Project là đơn vị quản lý công việc bên trong Workplace. Mỗi project có:
- Kanban board (project statuses)
- Danh sách thành viên với role `owner` | `member`
- Channel chat riêng (tự động tạo)
- Task management (xem [task.md](./task.md))

---

## Endpoints

### POST `/workplaces/{workplace_id}/projects`

Tạo project mới trong workplace.

**Request:**
```json
{
  "name": "Backend API",
  "description": "REST API for the platform",
  "color": "#6366F1",
  "access": "private"
}
```

`access` = `public` | `private`

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "workplace_id": 1,
    "name": "Backend API",
    "description": "REST API for the platform",
    "color": "#6366F1",
    "access": "private",
    "created_by": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

> Tự động tạo channel loại `project` và thêm creator vào với role `owner`.

---

### GET `/workplaces/{workplace_id}/projects`

Lấy danh sách project trong workplace.

**Query params:**
- `access` — `public` | `private`

---

### GET `/projects/{id}`

Lấy chi tiết project.

---

### PUT `/projects/{id}`

Cập nhật project (chỉ `owner`).

**Request:**
```json
{ "name": "Updated Name", "description": "...", "color": "#10B981", "access": "public" }
```

---

### DELETE `/projects/{id}`

Xóa project và tất cả tasks bên trong (chỉ `owner`).

---

## Member Endpoints

### GET `/projects/{id}/members`

Lấy danh sách thành viên.

**Response (200):**
```json
{
  "data": [
    { "user_id": 1, "name": "John Doe", "role": "owner",  "joined_at": "..." },
    { "user_id": 2, "name": "Jane Doe", "role": "member", "joined_at": "..." }
  ]
}
```

---

### POST `/projects/{id}/members`

Thêm thành viên (chỉ từ danh sách workplace members).

**Request:** `{ "user_id": 3, "role": "member" }`

---

### PUT `/projects/{id}/members/{user_id}/role`

Đổi role thành viên.

**Request:** `{ "role": "owner" }`

---

### DELETE `/projects/{id}/members/{user_id}`

Xóa thành viên khỏi project.

---

### DELETE `/projects/{id}/leave`

Tự rời project.

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `projects` | Thông tin project |
| `project_members` | Thành viên + role (`owner`/`member`) |
| `project_statuses` | Các cột Kanban |
| `channels` | Channel `type=project` tự tạo |
