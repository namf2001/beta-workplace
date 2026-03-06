# Beta Workspace API Documentation

**Version:** 1.0.0
**Base URL:** `http://localhost:8080/api/1.0.0`
**Description:** API documentation for Beta - an integrated workspace platform that helps teams work more efficiently through a unified ecosystem: team chat, collaborative documents, email, task management, and workflow automation.

## Authentication

Most endpoints require authentication using Bearer token.

**Header:**
```
Authorization: Bearer <your_token>
```

---

## Table of Contents

1. [Authentication APIs](#authentication-apis)
2. [User Management APIs](#user-management-apis)
3. [Organization APIs](#organization-apis)
4. [Project Management APIs](#project-management-apis)
5. [Task Management APIs](#task-management-apis)
6. [Room/Chat APIs](#roomchat-apis)
7. [Message APIs](#message-apis)
8. [File Management APIs](#file-management-apis)
9. [Notification APIs](#notification-apis)
10. [System APIs](#system-apis)

---

## Authentication APIs

### 1. Register Account (Multi-step)

**Endpoint:** `POST /auth/register`

**Description:** Đăng ký tài khoản với 3 bước: (1) Gửi mã xác thực, (2) Xác thực email, (3) Hoàn thành đăng ký

**Request Body:**

```json
{
  "step": 1,
  "email": "user@example.com"
}
```

**Step 1 - Send Verification Code:**
```json
{
  "step": 1,
  "email": "user@example.com"
}
```

**Step 2 - Verify Email:**
```json
{
  "step": 2,
  "email": "user@example.com",
  "verify_code": "123456"
}
```

**Step 3 - Complete Registration:**
```json
{
  "step": 3,
  "email": "user@example.com",
  "verify_code": "123456",
  "full_name": "John Doe",
  "password": "password123"
}
```

**Response (Step 3):**
```json
{
  "code": "INF004",
  "message": "Register user success",
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "encrypted_email",
      "full_name": "encrypted_name",
      "profile_image": "",
      "created_at": "2024-01-01T00:00:00Z"
    },
    "token": "encrypted_session_token"
  }
}
```

### 2. Login

**Endpoint:** `POST /login`

**Description:** Đăng nhập vào workspace platform

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "code": "INF001",
  "message": "Login success",
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "encrypted_email",
      "full_name": "encrypted_name",
      "profile_image": "",
      "colour": "#FF5733"
    },
    "token": "encrypted_session_token"
  }
}
```

### 3. Logout

**Endpoint:** `POST /auth/logout`

**Description:** Đăng xuất khỏi workspace platform

**Headers:** Authorization required

**Response:**
```json
{
  "code": "INF002",
  "message": "Logout success"
}
```

### 4. Google OAuth - Get Auth URL

**Endpoint:** `GET /auth/google/url`

**Description:** Lấy URL để redirect user đến Google OAuth

**Response:**
```json
{
  "code": "INF001",
  "message": "Google OAuth URL generated successfully",
  "data": {
    "auth_url": "https://accounts.google.com/o/oauth2/auth...",
    "state": "random_state_string"
  }
}
```

### 5. Google OAuth - Callback

**Endpoint:** `GET /auth/google/callback`

**Description:** Xử lý callback từ Google OAuth

**Query Parameters:**
- `code` (required): OAuth authorization code
- `state` (optional): OAuth state parameter

**Response:**
```json
{
  "code": "INF001",
  "message": "Login success",
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "encrypted_email",
      "full_name": "encrypted_name",
      "profile_image": "https://..."
    },
    "token": "encrypted_session_token"
  }
}
```

---

## User Management APIs

### 1. Get User Profile

**Endpoint:** `GET /auth/user/profile`

**Description:** Lấy thông tin người dùng hiện tại

**Headers:** Authorization required

**Response:**
```json
{
  "code": "INF008",
  "message": "Get user info success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "full_name": "John Doe",
    "email": "user@example.com",
    "phone_number": "+84123456789",
    "profile_image": "https://...",
    "colour": "#FF5733",
    "created_at": "2024-01-01 00:00:00",
    "updated_at": "2024-01-01 00:00:00"
  }
}
```

### 2. Update Profile

**Endpoint:** `PUT /auth/user/profile`

**Description:** Cập nhật thông tin cá nhân

**Headers:** Authorization required

**Request Body:**
```json
{
  "full_name": "base64_encrypted_name",
  "email": "base64_encrypted_email",
  "phone_number": "base64_encrypted_phone",
  "address": "base64_encrypted_address",
  "date_of_birth": "1990-01-01",
  "gender": 1,
  "colour": "#FF5733",
  "profile_image": "https://..."
}
```

**Response:**
```json
{
  "code": "INF010",
  "message": "Update profile success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "full_name": "John Doe",
    "email": "user@example.com"
  }
}
```

### 3. Change Password

**Endpoint:** `PATCH /auth/user/password`

**Description:** Thay đổi mật khẩu

**Headers:** Authorization required

**Request Body:**
```json
{
  "old_password": "old_password123",
  "new_password": "new_password123"
}
```

**Response:**
```json
{
  "code": "INF006",
  "message": "Change password success"
}
```

### 4. Forgot Password (Multi-step)

**Endpoint:** `POST /forgot_password`

**Description:** Quên mật khẩu, step = 1 gửi mail, step = 2 nhập mật khẩu mới

**Step 1 Request:**
```json
{
  "step": 1,
  "email": "user@example.com"
}
```

**Step 2 Request:**
```json
{
  "step": 2,
  "email": "user@example.com",
  "verify_code": "123456",
  "new_password": "new_password123"
}
```

**Response (Step 1):**
```json
{
  "code": "INF005",
  "message": "Send mail forgot password success",
  "data": {
    "code": "123456"
  }
}
```

**Response (Step 2):**
```json
{
  "code": "INF006",
  "message": "Change password success"
}
```

### 5. Delete Account

**Endpoint:** `DELETE /auth/user/account`

**Description:** Xóa user và tất cả thông tin liên quan

**Headers:** Authorization required

**Response:**
```json
{
  "code": "INF055",
  "message": "Delete user success"
}
```

---

## Organization APIs

### 1. Create Organization

**Endpoint:** `POST /auth/organizations`

**Description:** Tạo tổ chức mới với user hiện tại làm admin

**Headers:** Authorization required

**Request Body:**
```json
{
  "name": "My Organization",
  "industry": "Technology",
  "region": "Asia",
  "avatar": "https://...",
  "description": "Organization description",
  "size": 50
}
```

**Response:**
```json
{
  "code": "ORG001",
  "message": "Create organization success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "My Organization",
    "industry": "Technology",
    "size": 50,
    "region": "Asia",
    "avatar": "https://...",
    "description": "Organization description",
    "admin_count": 1,
    "member_count": 1,
    "sub_admin_count": 0,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Get User Organizations

**Endpoint:** `GET /auth/organizations`

**Description:** Lấy danh sách tổ chức mà user tham gia

**Headers:** Authorization required

**Query Parameters:**
- `page` (required): Page number
- `page_size` (required): Number of items per page

**Response:**
```json
{
  "code": "ORG002",
  "message": "Get user organizations success",
  "data": {
    "organizations": [
      {
        "organization": {
          "id": "507f1f77bcf86cd799439011",
          "name": "My Organization",
          "industry": "Technology"
        },
        "role": "admin",
        "joined_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1
  }
}
```

### 3. Get Organization Details

**Endpoint:** `GET /auth/organizations/{id}`

**Description:** Lấy thông tin chi tiết của một tổ chức

**Headers:** Authorization required

**Response:**
```json
{
  "code": "ORG003",
  "message": "Get organization success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "My Organization",
    "industry": "Technology",
    "size": 50,
    "region": "Asia",
    "avatar": "https://...",
    "description": "Organization description",
    "admin_count": 1,
    "member_count": 5,
    "sub_admin_count": 2
  }
}
```

### 4. Update Organization

**Endpoint:** `PUT /auth/organizations/{id}`

**Description:** Cập nhật thông tin tổ chức (chỉ admin)

**Headers:** Authorization required

**Request Body:**
```json
{
  "name": "Updated Organization Name",
  "industry": "Technology",
  "size": 100,
  "region": "Global",
  "avatar": "https://...",
  "description": "Updated description"
}
```

**Response:**
```json
{
  "code": "ORG004",
  "message": "Update organization success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "Updated Organization Name"
  }
}
```

### 5. Invite Member

**Endpoint:** `POST /auth/organizations/{id}/invite`

**Description:** Mời thành viên vào tổ chức (admin hoặc sub-admin)

**Headers:** Authorization required

**Request Body:**
```json
{
  "email": "newmember@example.com",
  "role": "member"
}
```

**Response:**
```json
{
  "code": "ORG005",
  "message": "Create invitation success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "organization_id": "507f1f77bcf86cd799439012",
    "invite_code": "ABC123XYZ",
    "email": "newmember@example.com",
    "expires_at": "2024-01-08T00:00:00Z",
    "status": "pending",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 6. Join Organization

**Endpoint:** `POST /auth/organizations/join`

**Description:** Tham gia tổ chức bằng mã mời

**Headers:** Authorization required

**Request Body:**
```json
{
  "invite_code": "ABC123XYZ"
}
```

**Response:**
```json
{
  "code": "ORG006",
  "message": "Join organization success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "organization_id": "507f1f77bcf86cd799439012",
    "user_id": "507f1f77bcf86cd799439013",
    "user_name": "John Doe",
    "user_email": "user@example.com",
    "user_avatar": "https://...",
    "role": "member",
    "invited_by": "507f1f77bcf86cd799439014",
    "joined_at": "2024-01-01T00:00:00Z"
  }
}
```

### 7. Get Organization Members

**Endpoint:** `GET /auth/organizations/{id}/members`

**Description:** Lấy danh sách thành viên của tổ chức

**Headers:** Authorization required

**Response:**
```json
{
  "code": "MEM001",
  "message": "Get members success",
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "organization_id": "507f1f77bcf86cd799439012",
      "user_id": "507f1f77bcf86cd799439013",
      "role": "admin",
      "invited_by": "507f1f77bcf86cd799439014",
      "joined_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 8. Update Member Role

**Endpoint:** `PUT /auth/organizations/{id}/members/role`

**Description:** Cập nhật vai trò của thành viên (chỉ admin)

**Headers:** Authorization required

**Request Body:**
```json
{
  "user_id": "507f1f77bcf86cd799439013",
  "role": "sub_admin"
}
```

**Response:**
```json
{
  "code": "MEM002",
  "message": "Update member role success"
}
```

### 9. Remove Member

**Endpoint:** `DELETE /auth/organizations/{id}/members/{memberId}`

**Description:** Xóa thành viên khỏi tổ chức (chỉ admin)

**Headers:** Authorization required

**Response:**
```json
{
  "code": "MEM003",
  "message": "Remove member success"
}
```

---

## Project Management APIs

### 1. Create Project

**Endpoint:** `POST /auth/projects`

**Description:** Tạo dự án mới trong organization

**Headers:** Authorization required

**Request Body:**
```json
{
  "organization_id": "507f1f77bcf86cd799439011",
  "name": "Project Alpha",
  "key": "ALPHA",
  "description": "Project description",
  "avatar": "https://...",
  "project_type": "software",
  "lead": "507f1f77bcf86cd799439012",
  "settings": {
    "allow_guests": false,
    "enable_notifications": true,
    "enable_due_dates": true,
    "enable_time_tracking": false
  }
}
```

**Response:**
```json
{
  "code": "PRJ001",
  "message": "Create project success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "organization_id": "507f1f77bcf86cd799439012",
    "name": "Project Alpha",
    "key": "ALPHA",
    "description": "Project description",
    "avatar": "https://...",
    "project_type": "software",
    "lead": "507f1f77bcf86cd799439013",
    "status": "active",
    "settings": {
      "allow_guests": false,
      "enable_notifications": true,
      "enable_due_dates": true,
      "enable_time_tracking": false
    },
    "created_by": "507f1f77bcf86cd799439014",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Get Project Details

**Endpoint:** `GET /auth/projects/{project_id}`

**Description:** Lấy thông tin chi tiết dự án

**Headers:** Authorization required

**Response:**
```json
{
  "code": "PRJ002",
  "message": "Get project success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "Project Alpha",
    "key": "ALPHA",
    "description": "Project description",
    "status": "active"
  }
}
```

### 3. Get User Projects

**Endpoint:** `GET /auth/projects/my-projects`

**Description:** Lấy tất cả dự án mà user tham gia

**Headers:** Authorization required

**Response:**
```json
{
  "code": "PRJ003",
  "message": "Get projects success",
  "data": [
    {
      "project": {
        "id": "507f1f77bcf86cd799439011",
        "name": "Project Alpha",
        "key": "ALPHA"
      }
    }
  ]
}
```

### 4. Update Project

**Endpoint:** `PUT /auth/projects/{project_id}`

**Description:** Cập nhật thông tin dự án (chỉ admin)

**Headers:** Authorization required

**Request Body:**
```json
{
  "name": "Updated Project Name",
  "description": "Updated description",
  "avatar": "https://...",
  "project_type": "software",
  "lead": "507f1f77bcf86cd799439012",
  "status": "active"
}
```

**Response:**
```json
{
  "code": "PRJ004",
  "message": "Update project success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "Updated Project Name"
  }
}
```

### 5. Add Project Member

**Endpoint:** `POST /auth/projects/{project_id}/members`

**Description:** Thêm thành viên vào dự án (chỉ admin)

**Headers:** Authorization required

**Request Body:**
```json
{
  "user_id": "507f1f77bcf86cd799439011",
  "role": "member"
}
```

**Response:**
```json
{
  "code": "MEM001",
  "message": "Add member success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "project_id": "507f1f77bcf86cd799439012",
    "user": {
      "id": "507f1f77bcf86cd799439013"
    },
    "role": "member",
    "joined_at": "2024-01-01T00:00:00Z"
  }
}
```

### 6. Get Project Members

**Endpoint:** `GET /auth/projects/{project_id}/members`

**Description:** Lấy danh sách thành viên trong dự án

**Headers:** Authorization required

**Response:**
```json
{
  "code": "MEM001",
  "message": "Get members success",
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "project_id": "507f1f77bcf86cd799439012",
      "user": {
        "id": "507f1f77bcf86cd799439013"
      },
      "role": "admin",
      "joined_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

---

## Task Management APIs

### 1. Create Task List

**Endpoint:** `POST /auth/task-lists`

**Description:** Tạo danh sách task mới (Kanban column)

**Headers:** Authorization required

**Request Body:**
```json
{
  "project_id": "507f1f77bcf86cd799439011",
  "name": "To Do",
  "description": "Tasks to be done",
  "color": "#FF5733",
  "position": 1,
  "is_default": false
}
```

**Response:**
```json
{
  "code": "TSK001",
  "message": "Create task list success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "project_id": "507f1f77bcf86cd799439012",
    "name": "To Do",
    "description": "Tasks to be done",
    "color": "#FF5733",
    "position": 1,
    "is_default": false,
    "created_by": "507f1f77bcf86cd799439013",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Create Task

**Endpoint:** `POST /auth/tasks`

**Description:** Tạo task mới trong danh sách

**Headers:** Authorization required

**Request Body:**
```json
{
  "project_id": "507f1f77bcf86cd799439011",
  "list_id": "507f1f77bcf86cd799439012",
  "title": "Implement login feature",
  "description": "Create login page with email/password",
  "priority": "high",
  "assignee": "507f1f77bcf86cd799439013",
  "due_date": "2024-01-15T00:00:00Z",
  "due_reminder": {
    "enabled": true,
    "time_before": 86400
  },
  "labels": ["507f1f77bcf86cd799439014"],
  "parent_task_id": "507f1f77bcf86cd799439015"
}
```

**Response:**
```json
{
  "code": "TSK002",
  "message": "Create task success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "project_id": "507f1f77bcf86cd799439012",
    "list_id": "507f1f77bcf86cd799439013",
    "title": "Implement login feature",
    "description": "Create login page",
    "completed": false,
    "position": 1,
    "reporter": {
      "id": "507f1f77bcf86cd799439014",
      "full_name": "John Doe",
      "email": "john@example.com"
    },
    "assignee": {
      "id": "507f1f77bcf86cd799439015",
      "full_name": "Jane Doe"
    },
    "due_date": "2024-01-15T00:00:00Z",
    "priority": "high",
    "labels": ["507f1f77bcf86cd799439016"],
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 3. Get Task Details

**Endpoint:** `GET /auth/tasks/{task_id}`

**Description:** Lấy thông tin chi tiết task

**Headers:** Authorization required

**Response:**
```json
{
  "code": "TSK003",
  "message": "Get task success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "project_id": "507f1f77bcf86cd799439012",
    "list_id": "507f1f77bcf86cd799439013",
    "title": "Implement login feature",
    "description": "Create login page",
    "completed": false,
    "priority": "high"
  }
}
```

### 4. Update Task

**Endpoint:** `PUT /auth/tasks/{task_id}`

**Description:** Cập nhật thông tin task

**Headers:** Authorization required

**Request Body:**
```json
{
  "title": "Updated task title",
  "description": "Updated description",
  "priority": "medium",
  "assignee": "507f1f77bcf86cd799439011",
  "due_date": "2024-01-20T00:00:00Z",
  "labels": ["507f1f77bcf86cd799439012"],
  "completed": false
}
```

**Response:**
```json
{
  "code": "TSK004",
  "message": "Update task success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Updated task title",
    "priority": "medium"
  }
}
```

### 5. Move Task

**Endpoint:** `PUT /auth/tasks/{task_id}/move`

**Description:** Di chuyển task sang list khác hoặc thay đổi vị trí

**Headers:** Authorization required

**Request Body:**
```json
{
  "target_list_id": "507f1f77bcf86cd799439011",
  "target_position": 2
}
```

**Response:**
```json
{
  "code": "TSK005",
  "message": "Move task success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "list_id": "507f1f77bcf86cd799439012",
    "position": 2
  }
}
```

### 6. Get Tasks by Project

**Endpoint:** `GET /auth/projects/{project_id}/tasks`

**Description:** Lấy danh sách tasks trong dự án với filter và pagination

**Headers:** Authorization required

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `page_size` (optional): Page size (default: 20)
- `assignee` (optional): Filter by assignee ID
- `status` (optional): Filter by task status
- `priority` (optional): Filter by priority

**Response:**
```json
{
  "code": "TSK006",
  "message": "Get tasks success",
  "data": {
    "tasks": [
      {
        "id": "507f1f77bcf86cd799439011",
        "title": "Task 1",
        "priority": "high"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 20,
    "total_pages": 5
  }
}
```

### 7. Add Task Comment

**Endpoint:** `POST /auth/tasks/{task_id}/comments`

**Description:** Thêm comment vào task

**Headers:** Authorization required

**Request Body:**
```json
{
  "content": "This is a comment",
  "mentions": ["507f1f77bcf86cd799439011"],
  "parent_id": "507f1f77bcf86cd799439012"
}
```

**Response:**
```json
{
  "code": "TSK007",
  "message": "Add task comment success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "task_id": "507f1f77bcf86cd799439012",
    "author": {
      "id": "507f1f77bcf86cd799439013",
      "full_name": "John Doe"
    },
    "content": "This is a comment",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 8. Get Kanban Board

**Endpoint:** `GET /auth/projects/{project_id}/kanban`

**Description:** Lấy toàn bộ Kanban board với task lists và tasks

**Headers:** Authorization required

**Response:**
```json
{
  "code": "TSK006",
  "message": "Get tasks success",
  "data": {
    "project": {
      "id": "507f1f77bcf86cd799439011",
      "name": "Project Alpha"
    },
    "lists": [
      {
        "list": {
          "id": "507f1f77bcf86cd799439012",
          "name": "To Do",
          "tasks_count": 5
        },
        "tasks": [
          {
            "id": "507f1f77bcf86cd799439013",
            "title": "Task 1"
          }
        ]
      }
    ],
    "members": [],
    "labels": ["Bug", "Feature"],
    "priorities": ["low", "medium", "high"]
  }
}
```

### 9. Get Project Dashboard

**Endpoint:** `GET /auth/projects/{project_id}/dashboard`

**Description:** Lấy dashboard với thống kê và hoạt động gần đây

**Headers:** Authorization required

**Response:**
```json
{
  "code": "PRJ002",
  "message": "Get project success",
  "data": {
    "project": {
      "id": "507f1f77bcf86cd799439011",
      "name": "Project Alpha"
    },
    "tasks_summary": {
      "total": 50,
      "completed": 30,
      "in_progress": 15,
      "todo": 5
    },
    "recent_activity": [],
    "team_activity": [],
    "upcoming_deadlines": []
  }
}
```

### 10. Delete Task

**Endpoint:** `DELETE /auth/tasks/{task_id}`

**Description:** Xóa task khỏi dự án

**Headers:** Authorization required

**Response:**
```json
{
  "code": "TSK008",
  "message": "Delete task success"
}
```

### 11. Delete Task List

**Endpoint:** `DELETE /auth/task-lists/{list_id}`

**Description:** Xóa danh sách task (không thể xóa default list)

**Headers:** Authorization required

**Response:**
```json
{
  "code": "TSK009",
  "message": "Delete task list success"
}
```

### 12. Delete Task Comment

**Endpoint:** `DELETE /auth/task-comments/{comment_id}`

**Description:** Xóa comment khỏi task

**Headers:** Authorization required

**Response:**
```json
{
  "code": "TSK010",
  "message": "Delete task comment success"
}
```

### 13. Delete Multiple Tasks

**Endpoint:** `DELETE /auth/tasks/batch-delete`

**Description:** Xóa nhiều tasks cùng lúc

**Headers:** Authorization required

**Request Body:**
```json
{
  "task_ids": [
    "507f1f77bcf86cd799439011",
    "507f1f77bcf86cd799439012"
  ]
}
```

**Response:**
```json
{
  "code": "TSK008",
  "message": "Delete task success",
  "data": {
    "deleted_task_ids": ["507f1f77bcf86cd799439011"],
    "failed_task_ids": ["507f1f77bcf86cd799439012"],
    "total_requested": 2,
    "total_deleted": 1,
    "total_failed": 1
  }
}
```

---

## Room/Chat APIs

### 1. Get or Create Room

**Endpoint:** `POST /auth/room/`

**Description:** Lấy thông tin room nếu đã tồn tại, ngược lại tạo room mới

**Headers:** Authorization required

**Request Body:**
```json
{
  "organization_id": "507f1f77bcf86cd799439011",
  "name": "Project Discussion",
  "type_of_room": "direct",
  "members": [
    "507f1f77bcf86cd799439012",
    "507f1f77bcf86cd799439013"
  ]
}
```

**Type of Room:**
- `direct`: Direct message (1-1)
- `multi_member`: Group chat
- `circle`: Organization-wide room

**Response (Room Created - 201):**
```json
{
  "code": "INF027",
  "message": "Create room success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "Project Discussion",
    "members": [
      {
        "id": "507f1f77bcf86cd799439012",
        "full_name": "John Doe",
        "email": "john@example.com"
      }
    ],
    "latest_message": {},
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "has_notification": false
  }
}
```

**Response (Room Exists - 200):**
```json
{
  "code": "INF030",
  "message": "Get room success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "Project Discussion"
  }
}
```

### 2. Get List of Rooms

**Endpoint:** `GET /auth/room/{circle_id}`

**Description:** Lấy danh sách room

**Headers:** Authorization required

**Query Parameters:**
- `page` (required): Page number
- `page_size` (required): Number of items per page

**Response:**
```json
{
  "code": "INF030",
  "message": "Get room success",
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "name": "Room 1",
      "members": [],
      "latest_message": {
        "id": "507f1f77bcf86cd799439012",
        "content": "Hello"
      },
      "has_notification": true
    }
  ]
}
```

---

## Message APIs

### 1. Create Message

**Endpoint:** `POST /chat/{room_id}`

**Description:** Tạo và gửi tin nhắn thông qua websocket

**Headers:** Authorization required

**Request Body:**
```json
{
  "content": "Hello, how are you?",
  "type_content": "text"
}
```

**Type Content:**
- `text`: Text message
- `image`: Image message
- `file`: File attachment
- `video`: Video message

**Response:**
```json
{
  "code": "INF028",
  "message": "Create message success"
}
```

### 2. Get Messages

**Endpoint:** `GET /chat/{room_id}`

**Description:** Lấy danh sách tin nhắn trong phòng chat

**Headers:** Authorization required

**Query Parameters:**
- `page` (required): Page number
- `page_size` (required): Number of items per page

**Response:**
```json
{
  "code": "INF029",
  "message": "Get message success",
  "data": {
    "room": {
      "id": "507f1f77bcf86cd799439011",
      "name": "Room Name"
    },
    "members": [
      {
        "id": "507f1f77bcf86cd799439012",
        "full_name": "John Doe"
      }
    ],
    "messages": [
      {
        "id": "507f1f77bcf86cd799439013",
        "content": "Hello",
        "type_content": "text",
        "created_at": "2024-01-01T00:00:00Z",
        "user": {
          "id": "507f1f77bcf86cd799439014",
          "full_name": "John Doe",
          "email": "john@example.com",
          "profile_image": "https://...",
          "colour": "#FF5733"
        }
      }
    ]
  }
}
```

---

## File Management APIs

### 1. Upload File

**Endpoint:** `POST /auth/upload`

**Description:** Upload file lên server

**Headers:** Authorization required

**Request Body:** `multipart/form-data`
- `file`: File data

**Response:**
```json
{
  "code": "INF007",
  "message": "Upload file success",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "user_id": "507f1f77bcf86cd799439012",
    "key": "user_id/year=2024/month=01/day=01/uuid_timestamp_filename.jpg",
    "mime": "image/jpeg",
    "url": "https://s3.amazonaws.com/bucket/...",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Download File

**Endpoint:** `GET /auth/download`

**Description:** Tải file từ server

**Headers:** Authorization required

**Query Parameters (one required):**
- `object_id`: Object S3 ID
- `object_key`: Object S3 key

**Response:** File data (binary)

---

## Notification APIs

### 1. Register Device Token

**Endpoint:** `POST /auth/notification/register_device`

**Description:** Đăng ký token device để nhận thông báo FCM

**Headers:** Authorization required

**Request Body:**
```json
{
  "token_device": "firebase_device_token_here"
}
```

**Response:**
```json
{
  "code": "INF052",
  "message": "Device registration success"
}
```

### 2. Send Notification

**Endpoint:** `POST /auth/notification/send`

**Description:** Gửi thông báo cho workspace

**Headers:** Authorization required

**Request Body:**
```json
{
  "notify_type": "1",
  "data": {
    "room": {
      "id": "507f1f77bcf86cd799439011",
      "name": "Room Name"
    },
    "user": {
      "id": "507f1f77bcf86cd799439012",
      "full_name": "John Doe"
    },
    "message": {
      "id": "507f1f77bcf86cd799439013",
      "content": "New message"
    }
  }
}
```

**Notification Types:**
- `1`: New message
- `2`: Task assigned
- `3`: Task completed
- `4`: Mention in comment

**Response:**
```json
{
  "code": "INF031",
  "message": "Send notification success"
}
```

---

## System APIs

### 1. Get Public Key

**Endpoint:** `GET /public_key`

**Description:** Lấy public key RSA để client mã hóa dữ liệu

**Response:** Public Key PEM format (text/plain)

```
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...
-----END PUBLIC KEY-----
```

### 2. Get Android Asset Links

**Endpoint:** `GET /.well-known/assetlinks.json`

**Description:** Lấy thông tin file assetlinks.json cho Android App Links

**Response:**
```json
[
  {
    "relation": ["delegate_permission/common.handle_all_urls"],
    "target": {
      "namespace": "android_app",
      "package_name": "com.example.app",
      "sha256_cert_fingerprints": ["..."]
    }
  }
]
```

### 3. Get Apple App Site Association

**Endpoint:** `GET /.well-known/apple-app-site-association`

**Description:** Lấy thông tin file apple-app-site-association cho iOS Universal Links

**Response:**
```json
{
  "applinks": {
    "apps": [],
    "details": [
      {
        "appID": "TEAM_ID.com.example.app",
        "paths": ["*"]
      }
    ]
  }
}
```

---

## Common Response Codes

### Success Codes
- `INF001`: Login success
- `INF002`: Logout success
- `INF004`: Register user success
- `INF005`: Send mail forgot password success
- `INF006`: Change password success
- `INF007`: Upload file success
- `INF008`: Get user info success
- `INF010`: Update profile success
- `INF027`: Create room success
- `INF028`: Create message success
- `INF029`: Get message success
- `INF030`: Get room success
- `INF031`: Send notification success
- `INF052`: Device registration success
- `INF055`: Delete user success

### Error Codes
- `ERR001`: Invalid request params
- `ERR002`: Token invalid
- `ERR003`: User not found
- `ERR004`: Password incorrect
- `ERR005`: Email does not exist
- `ERR006`: Email already exists
- `ERR007`: Phone number already exists
- `ERR008`: Access denied
- `ERR009`: Room not found
- `ERR010`: Organization not found
- `ERR011`: Project not found
- `ERR012`: Task not found

---

## WebSocket Events

### Connect to WebSocket

**URL:** `ws://localhost:8080/ws`

**Authentication:** Send token in first message after connection

### Events

#### 1. Join Room
```json
{
  "event": "join_room",
  "room_id": "507f1f77bcf86cd799439011"
}
```

#### 2. Leave Room
```json
{
  "event": "leave_room",
  "room_id": "507f1f77bcf86cd799439011"
}
```

#### 3. New Message (Receive)
```json
{
  "event": "new_message",
  "room_id": "507f1f77bcf86cd799439011",
  "message": {
    "id": "507f1f77bcf86cd799439012",
    "content": "Hello",
    "type_content": "text",
    "user": {
      "id": "507f1f77bcf86cd799439013",
      "full_name": "John Doe"
    },
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 4. User Typing
```json
{
  "event": "typing",
  "room_id": "507f1f77bcf86cd799439011",
  "user_id": "507f1f77bcf86cd799439012",
  "is_typing": true
}
```

---

## Rate Limiting

API có giới hạn số lượng request:
- **Rate Limit:** 100 requests per minute per user
- **Response Header:** `X-RateLimit-Remaining`, `X-RateLimit-Reset`

---

## Error Handling

Tất cả API đều trả về format response như sau:

**Success Response:**
```json
{
  "code": "INF001",
  "message": "Success message",
  "data": {}
}
```

**Error Response:**
```json
{
  "code": "ERR001",
  "message": "Error message",
  "data": null
}
```

---

## Security

### RSA Encryption

Một số trường dữ liệu nhạy cảm (email, full_name, phone_number, address) cần được mã hóa RSA trước khi gửi lên server:

1. Lấy public key từ endpoint `/public_key`
2. Sử dụng public key để mã hóa dữ liệu
3. Gửi dữ liệu đã mã hóa dưới dạng base64 string

**Example fields requiring encryption:**
- `full_name` trong update profile
- `email` trong update profile
- `phone_number` trong update profile
- `address` trong update profile

### Password Requirements

- Minimum length: 6 characters
- Được hash bằng bcrypt trước khi lưu vào database

---

## Notes

1. Tất cả datetime đều sử dụng format **ISO 8601** (RFC3339): `2024-01-01T00:00:00Z`
2. Tất cả ID đều là **MongoDB ObjectID** dạng hex string (24 characters)
3. Timezone mặc định: **Asia/Ho_Chi_Minh** (UTC+7)
4. File upload size limit: **10MB**
5. Supported file types: images (jpg, png, gif), documents (pdf, doc, docx), archives (zip)

---

## Contact & Support

For API support or questions, please contact:
- **Email:** support@beta-workspace.com
- **Documentation:** https://docs.beta-workspace.com
- **GitHub Issues:** https://github.com/betasoft/beta-backend/issues
