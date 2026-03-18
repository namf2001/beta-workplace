-- ============================================================
-- Migration 003: Sessions & Verification Tokens
-- ============================================================

-- ─── SESSIONS ────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.sessions (
    id            BIGSERIAL   PRIMARY KEY,
    user_id       BIGINT      NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    expires       TIMESTAMPTZ NOT NULL,
    session_token TEXT        NOT NULL UNIQUE CHECK (session_token <> ''::text)
);

CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON public.sessions(user_id);

-- ─── VERIFICATION TOKENS ─────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.verification_token (
    identifier TEXT        NOT NULL CHECK (identifier <> ''::text),
    expires    TIMESTAMPTZ NOT NULL,
    token      TEXT        NOT NULL CHECK (token <> ''::text),
    PRIMARY KEY (identifier, token)
);
