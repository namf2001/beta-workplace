# Feature Documentation

Tài liệu cho từng tính năng của Beta Workplace platform.

## Danh sách features

| Feature | File | Mô tả |
|---------|------|-------|
| 🔐 Authentication | [auth.md](./auth.md) | Đăng ký, đăng nhập, OAuth, OTP |
| 👤 User Management | [user.md](./user.md) | Profile, đổi password, tìm kiếm user |
| 🏢 Organization | [organization.md](./organization.md) | Tổ chức, phân quyền admin/sub_admin/member |
| 🏠 Workplace | [workplace.md](./workplace.md) | Workspace, invite link, member |
| 📁 Project | [project.md](./project.md) | Project trong workplace, Kanban board |
| ✅ Task | [task.md](./task.md) | Task management (Jira-style), subtask, comment |
| 💬 Chat | [chat.md](./chat.md) | Channel, DM, tin nhắn, reaction, WebSocket |
| 📎 Upload | [upload.md](./upload.md) | File upload, gắn vào task/message |

## Architecture tổng quan

```
Organization
  └── (users join by OTP)

Workplace
  ├── Members (admin | member)
  ├── Channels (global, group)
  └── Projects
        ├── Members (owner | member)
        ├── Project Channel (auto-created)
        ├── Statuses (Kanban columns)
        └── Tasks
              ├── Assignees
              ├── Labels
              ├── Comments
              ├── Attachments → Files
              ├── Links
              └── Watchers

DM
  └── Friendship → Channel (type: dm)

Files
  └── Referenced by task_attachments + message_attachments
```

## Common Response Format

```json
{
  "data": { ... },
  "message": "...",
  "meta": {
    "page": 1,
    "page_size": 20,
    "total": 100
  }
}
```

## Error Format

```json
{
  "error": {
    "code": "ERR_NOT_FOUND",
    "message": "Resource not found"
  }
}
```

## HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK |
| 201 | Created |
| 400 | Bad request / validation error |
| 401 | Unauthorized (no/invalid token) |
| 403 | Forbidden (no permission) |
| 404 | Not found |
| 409 | Conflict (duplicate) |
| 410 | Gone (expired) |
| 422 | Unprocessable entity |
| 500 | Internal server error |
