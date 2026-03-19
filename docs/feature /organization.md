# Organization Feature

**Base URL:** `http://localhost:8080/api/v1`  
**Headers:** `Authorization: Bearer <token>`

## Overview

Organization là cấp cao nhất — một công ty, nhóm lớn. User tham gia bằng mã OTP được gửi qua email.

```
Role hierarchy:
  admin → sub_admin → member
```

| Role | Quyền hạn |
|------|-----------|
| `admin` | Toàn quyền: cấu hình org, quản lý member, phân quyền |
| `sub_admin` | Mời member, quản lý member thường |
| `member` | Xem và tham gia project |

---

## Endpoints

### POST `/organizations`

Tạo tổ chức mới. User tạo sẽ tự động là `admin`.

**Request:**
```json
{
  "name": "Acme Corp",
  "slug": "acme-corp",
  "logo_url": "https://...",
  "description": "We build things"
}
```

**Response (201):**
```json
{
  "data": {
    "id": 1,
    "name": "Acme Corp",
    "slug": "acme-corp",
    "logo_url": "https://...",
    "created_by": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### GET `/organizations`

Lấy danh sách organization của user.

---

### GET `/organizations/{id}`

Lấy chi tiết organization.

---

### PUT `/organizations/{id}`

Cập nhật tổ chức (chỉ admin).

**Request:**
```json
{ "name": "New Name", "logo_url": "https://...", "description": "Updated" }
```

---

### DELETE `/organizations/{id}`

Xóa tổ chức (chỉ admin tạo).

---

## Join Organization Flow

User tham gia bằng mã OTP:

```
Admin gửi invite → email có mã OTP
User nhập OTP tại app → POST /organizations/join-otp
```

### POST `/organizations/{id}/send-otp`

Admin/sub_admin gửi OTP mời qua email.

**Request:** `{ "email": "newmember@example.com" }`

---

### POST `/organizations/join-otp`

User nhập OTP để tham gia.

**Request:**
```json
{
  "organization_id": 1,
  "otp": "123456"
}
```

**Response (200):**
```json
{
  "data": {
    "organization_id": 1,
    "user_id": 5,
    "role": "member",
    "joined_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## Member Endpoints

### GET `/organizations/{id}/members`

Danh sách thành viên.

**Query params:**
- `role` — lọc theo role: `admin` | `sub_admin` | `member`
- `page`, `page_size`

---

### PUT `/organizations/{id}/members/{user_id}/role`

Đổi role (chỉ admin).

**Request:** `{ "role": "sub_admin" }`

---

### DELETE `/organizations/{id}/members/{user_id}`

Xóa thành viên (admin/sub_admin).

---

### DELETE `/organizations/{id}/leave`

Tự rời tổ chức. Admin không thể rời nếu là admin duy nhất.

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `organizations` | Thông tin tổ chức |
| `organization_members` | Thành viên + role |
| `verification_codes` | OTP mã mời (type: `organization_join`) |
