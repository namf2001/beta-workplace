-- ============================================================
-- Migration 001: Users Table
-- ============================================================

CREATE TABLE IF NOT EXISTS public.users (
    id             BIGSERIAL   PRIMARY KEY,
    name           TEXT        CHECK (name <> ''::text),
    email          TEXT        UNIQUE CHECK (email <> ''::text),
    email_verified TIMESTAMPTZ,
    password       TEXT,
    image          TEXT,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_users_email ON public.users(email);
