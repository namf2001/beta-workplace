# Task Management Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Task management theo mô hình Kanban (tương tự Jira/Linear).

- **Project** → có nhiều **Statuses** (cột Kanban)
- **Task** → thuộc 1 status, có thể có subtask (`parent_id`)
- **Assignee** — người thực hiện, **Reporter** — người giao việc, **Created by** — người tạo task
- Task có thể gắn label, comment, attachment, watcher, link với task khác

---

## Project Status Endpoints

### POST `/projects/{id}/statuses`

Tạo cột Kanban mới.

**Request:**
```json
{ "name": "In Progress", "color": "#3B82F6", "position": 1 }
```

**Response (201):**
```json
{
  "data": { "id": 1, "project_id": 1, "name": "In Progress", "color": "#3B82F6", "position": 1 }
}
```

---

### GET `/projects/{id}/statuses`

Lấy tất cả cột Kanban của project.

---

### PUT `/projects/{id}/statuses/{status_id}`

Cập nhật tên, màu, vị trí cột.

---

### DELETE `/projects/{id}/statuses/{status_id}`

Xóa cột (chỉ được xóa khi không còn task trong cột).

---

## Task Endpoints

### POST `/projects/{id}/tasks`

Tạo task mới.

**Request:**
```json
{
  "status_id": 1,
  "parent_id": 0,
  "title": "Implement login API",
  "description": "Add POST /auth/login endpoint",
  "priority": "high",
  "due_date": "2024-02-01T00:00:00Z",
  "estimate": 4,
  "reporter_id": 2,
  "assignee_ids": [3, 4]
}
```

> `parent_id = 0` → root task. `parent_id > 0` → subtask.

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "project_id": 1,
    "status_id": 1,
    "parent_id": 0,
    "title": "Implement login API",
    "priority": "high",
    "position": 65536,
    "due_date": "2024-02-01T00:00:00Z",
    "estimate": 4,
    "created_by": 1,
    "reporter_id": 2,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/projects/{id}/tasks`

Lấy tất cả task trong project (phân nhóm theo status).

**Query params:**
- `status_id` — lọc theo cột
- `assignee_id` — lọc theo người làm
- `priority` — `highest` | `high` | `medium` | `low` | `lowest`
- `parent_id` = `0` để chỉ lấy root tasks

---

### GET `/tasks/{id}`

Lấy chi tiết task (bao gồm assignees, labels, comments).

---

### PUT `/tasks/{id}`

Cập nhật task.

**Request:** bất kỳ field nào của task.

---

### DELETE `/tasks/{id}`

Xóa task (và tất cả subtask).

---

### PATCH `/tasks/{id}/status`

Chuyển task sang cột khác (drag-and-drop Kanban).

**Request:**
```json
{ "status_id": 2, "position": 98304 }
```

> `position` dùng thuật toán midpoint: luôn là số float nằm giữa 2 task trong cột.

---

## Assignee Endpoints

### POST `/tasks/{id}/assignees`

Giao task cho user.

**Request:** `{ "user_id": 3 }`

---

### DELETE `/tasks/{id}/assignees/{user_id}`

Bỏ giao task.

---

## Comment Endpoints

### POST `/tasks/{id}/comments`

Thêm comment (hoặc reply vào thread).

**Request:**
```json
{ "content": "This is done.", "parent_id": 0 }
```

---

### PUT `/comments/{id}`

Sửa comment.

---

### DELETE `/comments/{id}`

Xóa comment.

---

## Label Endpoints

### POST `/projects/{id}/labels`

Tạo label cho project.

**Request:** `{ "name": "bug", "color": "#EF4444" }`

---

### POST `/tasks/{id}/labels`

Gắn label vào task.

**Request:** `{ "label_id": 1 }`

---

### DELETE `/tasks/{id}/labels/{label_id}`

Gỡ label.

---

## Task Link Endpoints

### POST `/tasks/{id}/links`

Liên kết 2 task.

**Request:**
```json
{ "target_id": 5, "link_type": "blocks" }
```

`link_type` = `blocks` | `is_blocked_by` | `duplicates` | `relates_to`

---

### DELETE `/tasks/{id}/links/{link_id}`

Xóa liên kết.

---

## Watcher Endpoints

### POST `/tasks/{id}/watch`

Theo dõi task (nhận notification khi có thay đổi).

---

### DELETE `/tasks/{id}/watch`

Bỏ theo dõi.

---

## Activity Log

### GET `/tasks/{id}/activity`

Lịch sử thay đổi của task.

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "actor_id": 1,
      "action": "status_changed",
      "old_value": { "status_id": 1 },
      "new_value": { "status_id": 2 },
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

`action` phổ biến: `created`, `status_changed`, `assignee_added`, `assignee_removed`, `priority_changed`, `due_date_changed`, `commented`

---

## Task Priority

| Value | Mô tả |
|-------|-------|
| `highest` | Khẩn cấp nhất |
| `high` | Cao |
| `medium` | Trung bình (mặc định) |
| `low` | Thấp |
| `lowest` | Thấp nhất |

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `project_statuses` | Cột Kanban |
| `tasks` | Task chính |
| `task_assignees` | Người thực hiện (có `assigned_by`) |
| `labels` + `task_labels` | Label và gán label |
| `task_comments` | Comment với thread |
| `task_attachments` | File gắn vào task (ref → files) |
| `task_links` | Liên kết giữa các task |
| `task_watchers` | Người theo dõi task |
| `task_activity_logs` | Audit trail |
