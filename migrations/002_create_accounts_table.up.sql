-- ============================================================
-- Migration 002: Accounts Table (OAuth accounts)
-- ============================================================

CREATE TABLE IF NOT EXISTS public.accounts (
    id                  BIGSERIAL PRIMARY KEY,
    user_id             BIGINT      NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    type                TEXT        NOT NULL CHECK (type <> ''::text),
    provider            TEXT        NOT NULL CHECK (provider <> ''::text),
    provider_account_id TEXT        NOT NULL CHECK (provider_account_id <> ''::text),
    refresh_token       TEXT,
    access_token        TEXT,
    expires_at          BIGINT,
    id_token            TEXT,
    scope               TEXT,
    session_state       TEXT,
    token_type          TEXT,
    UNIQUE (provider, provider_account_id)
);

CREATE INDEX IF NOT EXISTS idx_accounts_user_id ON public.accounts(user_id);