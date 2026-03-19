# File Upload Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Hệ thống upload file tập trung — tất cả file (task attachment, message attachment, avatar...) đều được lưu qua bảng `files` trước, sau đó gắn vào entity tương ứng bằng FK.

```
Upload file → nhận file_id
→ Gắn vào task:    POST /tasks/{id}/attachments
→ Gắn vào message: POST /messages/{id}/attachments
```

---

## Endpoints

### POST `/files`

Upload một file.

**Content-Type:** `multipart/form-data`

**Form fields:**
| Field | Type | Required | Mô tả |
|-------|------|----------|-------|
| `file` | file | ✅ | File cần upload |
| `bucket` | string | ✅ | Tên bucket lưu trữ |

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "uploaded_by": 1,
    "bucket": "beta-workplace-uploads",
    "file_key": "uploads/2024/01/abc123.png",
    "file_url": "https://cdn.example.com/uploads/2024/01/abc123.png",
    "file_name": "screenshot.png",
    "file_size": 204800,
    "mime_type": "image/png",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/files/{id}`

Lấy thông tin file.

**Response (200):** Giống response của POST `/files`.

---

### DELETE `/files/{id}`

Xóa file (chỉ `uploaded_by` mới được xóa).

> ⚠️ Khi xóa file, các attachment records dùng `file_id` này cũng bị xóa theo (CASCADE).

---

### POST `/tasks/{id}/attachments`

Gắn file đã upload vào task.

**Request:**
```json
{ "file_id": 1 }
```

**Response (201):**
```json
{
  "data": {
    "task_id": 10,
    "file_id": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### DELETE `/tasks/{task_id}/attachments/{file_id}`

Gỡ file khỏi task (không xóa file gốc).

---

### POST `/messages/{id}/attachments`

Gắn file vào message.

**Request:**
```json
{ "file_id": 1 }
```

---

### DELETE `/messages/{message_id}/attachments/{file_id}`

Gỡ file khỏi message.

---

## Upload Limits

| Loại file | Max size |
|-----------|----------|
| Image (png, jpg, gif, webp) | 10 MB |
| Document (pdf, docx, xlsx) | 50 MB |
| Video | 200 MB |
| Khác | 25 MB |

---

## Storage Flow

```
Client
  │
  ▼
POST /files  (multipart upload)
  │
  ▼
Server validate mime_type + file_size
  │
  ▼
Upload to Object Storage (S3 / GCS / R2)
  │  ← trả về bucket + file_key + file_url
  ▼
INSERT INTO public.files (...)
  │  ← trả về file_id
  ▼
Client dùng file_id để gắn vào task/message
```

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `files` | Metadata của tất cả file đã upload |
| `task_attachments` | Junction: task_id + file_id |
| `message_attachments` | Junction: message_id + file_id |
