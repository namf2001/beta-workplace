# Authentication Feature

**Base URL:** `http://localhost:8080/api/v1`

## Overview

Hệ thống xác thực hỗ trợ 2 phương thức:
1. **Credentials** — email + password với OTP verification
2. **OAuth** — Google, GitHub, Discord, Microsoft

## Flow

```
[Credentials Register]
  POST /auth/register (step 1 — send OTP)
  POST /auth/register (step 2 — verify OTP)
  POST /auth/register (step 3 — set name + password)
  → trả về JWT token

[Credentials Login]
  POST /auth/login
  → trả về JWT token

[OAuth Login]
  GET  /auth/{provider}/url     → redirect URL
  GET  /auth/{provider}/callback → JWT token

[Logout]
  POST /auth/logout
```

---

## Endpoints

### POST `/auth/register`

Đăng ký tài khoản — 3 bước.

**Step 1 — Gửi OTP:**
```json
{ "step": 1, "email": "user@example.com" }
```
> Server gửi mã OTP 6 chữ số đến email.

**Step 2 — Xác thực OTP:**
```json
{ "step": 2, "email": "user@example.com", "otp": "123456" }
```

**Step 3 — Hoàn tất đăng ký:**
```json
{
  "step": 3,
  "email": "user@example.com",
  "otp": "123456",
  "name": "John Doe",
  "password": "password123"
}
```

**Response (Step 3 — 201):**
```json
{
  "data": {
    "token": "<jwt>",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "user@example.com"
    }
  }
}
```

---

### POST `/auth/login`

Đăng nhập bằng email + password.

**Request:**
```json
{ "email": "user@example.com", "password": "password123" }
```

**Response (200):**
```json
{
  "data": {
    "token": "<jwt>",
    "user": { "id": 1, "name": "John Doe", "email": "user@example.com" }
  }
}
```

**Error cases:**
| Code | Reason |
|------|--------|
| 401 | Sai email/password |
| 403 | Email chưa được xác thực |

---

### GET `/auth/{provider}/url`

Lấy OAuth redirect URL.

`provider` = `google` | `github` | `discord` | `microsoft`

**Response (200):**
```json
{
  "data": {
    "auth_url": "https://accounts.google.com/o/oauth2/auth?...",
    "state": "random_state_string"
  }
}
```

---

### GET `/auth/{provider}/callback`

Callback từ OAuth provider.

**Query params:**
- `code` (required) — authorization code
- `state` (optional) — state string

**Response (200):** giống `/auth/login`

---

### POST `/auth/logout`

**Headers:** `Authorization: Bearer <token>`

**Response (200):**
```json
{ "message": "Logout success" }
```

---

### POST `/auth/forgot-password`

**Step 1 — Gửi OTP:**
```json
{ "step": 1, "email": "user@example.com" }
```

**Step 2 — Đặt lại mật khẩu:**
```json
{
  "step": 2,
  "email": "user@example.com",
  "otp": "123456",
  "new_password": "new_password123"
}
```

---

## JWT Token

Token được đính kèm trong header:
```
Authorization: Bearer <token>
```

Token payload:
```json
{ "user_id": 1, "email": "user@example.com", "exp": 1234567890 }
```

Token hết hạn sau **24 giờ**.

---

## Database Tables

| Table | Vai trò |
|-------|---------|
| `users` | Lưu thông tin người dùng |
| `accounts` | OAuth account connections |
| `sessions` | Active sessions |
| `verification_token` | OTP codes cho email verify & forgot password |
