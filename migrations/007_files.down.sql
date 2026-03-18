-- Rollback migration 007

-- Restore message_attachments
DROP TABLE IF EXISTS public.message_attachments;

CREATE TABLE IF NOT EXISTS public.message_attachments (
    id         BIGSERIAL   PRIMARY KEY,
    message_id BIGINT      NOT NULL REFERENCES public.messages(id) ON DELETE CASCADE,
    file_url   TEXT        NOT NULL CHECK (file_url <> ''::text),
    file_name  TEXT        CHECK (file_name <> ''::text),
    file_size  BIGINT,
    mime_type  TEXT        CHECK (mime_type <> ''::text),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Restore task_attachments
DROP TABLE IF EXISTS public.task_attachments;

CREATE TABLE IF NOT EXISTS public.task_attachments (
    id          BIGSERIAL   PRIMARY KEY,
    task_id     BIGINT      NOT NULL REFERENCES public.tasks(id) ON DELETE CASCADE,
    uploaded_by BIGINT               REFERENCES public.users(id) ON DELETE SET NULL,
    file_url    TEXT        NOT NULL CHECK (file_url <> ''::text),
    file_name   TEXT        CHECK (file_name <> ''::text),
    file_size   BIGINT,
    mime_type   TEXT        CHECK (mime_type <> ''::text),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

DROP TABLE IF EXISTS public.files;
