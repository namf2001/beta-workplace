# User Management Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Quản lý thông tin cá nhân, đổi mật khẩu, tìm kiếm user khác.

---

## Endpoints

### GET `/users/me`

Lấy thông tin user hiện tại.

**Response (200):**
```json
{
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "image": "https://...",
    "email_verified": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### PUT `/users/me`

Cập nhật thông tin cá nhân.

**Request:**
```json
{
  "name": "John Updated",
  "image": "https://new-avatar.png"
}
```

**Response (200):** Giống GET `/users/me`.

---

### PATCH `/users/me/password`

Đổi mật khẩu (chỉ cho tài khoản credentials, không phải OAuth).

**Request:**
```json
{
  "old_password": "current_password",
  "new_password": "new_secure_password"
}
```

**Response (200):** `{ "message": "Password changed successfully" }`

**Error cases:**
| Status | Reason |
|--------|--------|
| 400 | Mật khẩu cũ sai |
| 422 | New password quá yếu |
| 403 | Tài khoản OAuth không có password |

---

### DELETE `/users/me`

Xóa tài khoản và tất cả dữ liệu liên quan.

**Response (200):** `{ "message": "Account deleted" }`

---

### GET `/users/search`

Tìm kiếm user theo username hoặc email (để gửi friend request, thêm vào workspace...).

**Query params:**
- `q` (required) — từ khóa tìm kiếm
- `limit` (default: 10)

**Response (200):**
```json
{
  "data": [
    { "id": 2, "name": "Jane Doe", "email": "jane@example.com", "image": "https://..." }
  ]
}
```

---

### GET `/users/{id}`

Lấy thông tin public của một user khác.

**Response (200):**
```json
{
  "data": {
    "id": 2,
    "name": "Jane Doe",
    "image": "https://..."
  }
}
```

> Email và thông tin nhạy cảm không được trả về khi xem profile người khác.

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `users` | Thông tin tài khoản |
| `accounts` | OAuth connections |
| `sessions` | Phiên đăng nhập |
