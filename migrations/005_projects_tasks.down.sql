-- Rollback migration 005

DROP TABLE IF EXISTS task_activity_logs;
DROP TABLE IF EXISTS task_watchers;
DROP TABLE IF EXISTS task_links;
DROP TABLE IF EXISTS task_attachments;
DROP TABLE IF EXISTS task_comments;
DROP TABLE IF EXISTS task_labels;
DROP TABLE IF EXISTS labels;
DROP TABLE IF EXISTS task_assignees;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS project_statuses;
DROP TABLE IF EXISTS project_members;
DROP TABLE IF EXISTS projects;

DROP TYPE IF EXISTS task_link_type;
DROP TYPE IF EXISTS task_priority;
DROP TYPE IF EXISTS project_role;
DROP TYPE IF EXISTS project_access;
