-- ============================================================
-- Migration 007: Centralized Files Table
-- Refactors task_attachments and message_attachments to
-- reference a shared files table instead of storing file
-- metadata directly.
-- ============================================================

-- ─── FILES ───────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.files (
    id          BIGSERIAL   PRIMARY KEY,
    uploaded_by BIGINT      REFERENCES public.users(id) ON DELETE SET NULL,
    bucket      TEXT        NOT NULL CHECK (bucket <> ''::text),        -- storage bucket name (e.g. S3 bucket)
    file_key    TEXT        NOT NULL UNIQUE CHECK (file_key <> ''::text), -- object key / path inside bucket
    file_url    TEXT        NOT NULL CHECK (file_url <> ''::text),      -- public access URL
    file_name   TEXT        NOT NULL CHECK (file_name <> ''::text),     -- original file name
    file_size   BIGINT      NOT NULL CHECK (file_size > 0),             -- size in bytes
    mime_type   TEXT        NOT NULL CHECK (mime_type <> ''::text),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_files_uploaded_by ON public.files(uploaded_by);

-- ─── REFACTOR TASK ATTACHMENTS ───────────────────────────────
-- Drop old columns, add file_id FK instead

DROP TABLE IF EXISTS public.task_attachments;

CREATE TABLE IF NOT EXISTS public.task_attachments (
    task_id    BIGINT      NOT NULL REFERENCES public.tasks(id)  ON DELETE CASCADE,
    file_id    BIGINT      NOT NULL REFERENCES public.files(id)  ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (task_id, file_id)
);

-- ─── REFACTOR MESSAGE ATTACHMENTS ────────────────────────────

DROP TABLE IF EXISTS public.message_attachments;

CREATE TABLE IF NOT EXISTS public.message_attachments (
    message_id BIGINT      NOT NULL REFERENCES public.messages(id) ON DELETE CASCADE,
    file_id    BIGINT      NOT NULL REFERENCES public.files(id)    ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (message_id, file_id)
);
