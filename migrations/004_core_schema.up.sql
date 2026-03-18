-- ============================================================
-- Migration 004: Core Schema
-- Covers: users (extended), auth_providers, verification_codes,
--         sessions, organizations, workplaces, workplace_invitations,
--         friendships
-- ============================================================

-- ─── ENUMS ───────────────────────────────────────────────────

CREATE TYPE verification_code_type AS ENUM (
    'email_verification',
    'organization_join'
);

CREATE TYPE org_role        AS ENUM ('admin', 'sub_admin', 'member');
CREATE TYPE workplace_role  AS ENUM ('admin', 'member');
CREATE TYPE workplace_size  AS ENUM ('1-10', '11-50', '51-200', '201-500', '500+');
CREATE TYPE friendship_status AS ENUM ('pending', 'accepted', 'blocked');

-- ─── EXTEND USERS ────────────────────────────────────────────

ALTER TABLE public.users
    ADD COLUMN IF NOT EXISTS username          VARCHAR(100) UNIQUE,
    ADD COLUMN IF NOT EXISTS display_name      VARCHAR(255),
    ADD COLUMN IF NOT EXISTS avatar_url        TEXT,
    ADD COLUMN IF NOT EXISTS email_verified_at TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS is_active         BOOLEAN NOT NULL DEFAULT true;

CREATE INDEX IF NOT EXISTS idx_users_username ON public.users(username);

-- ─── AUTH PROVIDERS ──────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.auth_providers (
    id                  BIGSERIAL PRIMARY KEY,
    user_id             BIGINT      NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    provider            TEXT        NOT NULL CHECK (provider <> ''::text),
    provider_account_id TEXT        NOT NULL CHECK (provider_account_id <> ''::text),
    access_token        TEXT,
    refresh_token       TEXT,
    token_expires_at    TIMESTAMPTZ,
    id_token            TEXT,
    scope               TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (provider, provider_account_id)
);

CREATE INDEX IF NOT EXISTS idx_auth_providers_user
    ON public.auth_providers(user_id);

-- ─── VERIFICATION CODES ──────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.verification_codes (
    id          BIGSERIAL              PRIMARY KEY,
    user_id     BIGINT                 REFERENCES public.users(id) ON DELETE CASCADE,
    identifier  TEXT        NOT NULL   CHECK (identifier <> ''::text),
    code        TEXT        NOT NULL   CHECK (code <> ''::text),
    type        verification_code_type NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    used_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL   DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_verification_codes_identifier
    ON public.verification_codes(identifier, type);

-- ─── SESSIONS ────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.sessions (
    id            BIGSERIAL PRIMARY KEY,
    user_id       BIGINT      NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    session_token TEXT        NOT NULL UNIQUE CHECK (session_token <> ''::text),
    expires_at    TIMESTAMPTZ NOT NULL,
    user_agent    TEXT,
    ip_address    INET,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_sessions_user
    ON public.sessions(user_id);

-- ─── ORGANIZATIONS ───────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.organizations (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT        NOT NULL   CHECK (name <> ''::text),
    slug        TEXT        NOT NULL   UNIQUE CHECK (slug <> ''::text),
    logo_url    TEXT,
    description TEXT,
    created_by  BIGINT                 REFERENCES public.users(id) ON DELETE SET NULL,
    created_at  TIMESTAMPTZ NOT NULL   DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL   DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.organization_members (
    id              BIGSERIAL PRIMARY KEY,
    organization_id BIGINT   NOT NULL REFERENCES public.organizations(id) ON DELETE CASCADE,
    user_id         BIGINT   NOT NULL REFERENCES public.users(id)         ON DELETE CASCADE,
    role            org_role NOT NULL DEFAULT 'member',
    joined_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (organization_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_org_members_user
    ON public.organization_members(user_id);

-- ─── WORKPLACES ──────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.workplaces (
    id         BIGSERIAL      PRIMARY KEY,
    name       TEXT           NOT NULL CHECK (name <> ''::text),
    icon_url   TEXT,
    size       workplace_size,
    created_by BIGINT         REFERENCES public.users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ    NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.workplace_members (
    id           BIGSERIAL      PRIMARY KEY,
    workplace_id BIGINT         NOT NULL REFERENCES public.workplaces(id) ON DELETE CASCADE,
    user_id      BIGINT         NOT NULL REFERENCES public.users(id)      ON DELETE CASCADE,
    role         workplace_role NOT NULL DEFAULT 'member',
    joined_at    TIMESTAMPTZ    NOT NULL DEFAULT now(),
    UNIQUE (workplace_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_workplace_members_user
    ON public.workplace_members(user_id);

CREATE TABLE IF NOT EXISTS public.workplace_invitations (
    id           BIGSERIAL   PRIMARY KEY,
    workplace_id BIGINT      NOT NULL REFERENCES public.workplaces(id) ON DELETE CASCADE,
    invite_token TEXT        NOT NULL UNIQUE CHECK (invite_token <> ''::text),
    created_by   BIGINT      REFERENCES public.users(id) ON DELETE SET NULL,
    max_uses     INT,
    use_count    INT         NOT NULL DEFAULT 0,
    expires_at   TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_workplace_invitations_token
    ON public.workplace_invitations(invite_token);

-- ─── FRIENDSHIPS ─────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS public.friendships (
    id           BIGSERIAL        PRIMARY KEY,
    requester_id BIGINT           NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    receiver_id  BIGINT           NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    status       friendship_status NOT NULL DEFAULT 'pending',
    created_at   TIMESTAMPTZ      NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ      NOT NULL DEFAULT now(),
    UNIQUE (requester_id, receiver_id),
    CHECK (requester_id <> receiver_id)
);

CREATE INDEX IF NOT EXISTS idx_friendships_receiver
    ON public.friendships(receiver_id);
