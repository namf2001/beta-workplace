-- ============================================================
-- Migration 008: Organization Invitations + Extend Organizations
-- Covers: add industry/size/region to organizations,
--         create organization_invitations table
-- ============================================================

-- ─── ENUMS ────────────────────────────────────────────────────

CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'expired', 'cancelled');

-- ─── EXTEND ORGANIZATIONS ─────────────────────────────────────

ALTER TABLE public.organizations
    ADD COLUMN IF NOT EXISTS industry TEXT,
    ADD COLUMN IF NOT EXISTS size     INT,
    ADD COLUMN IF NOT EXISTS region   TEXT;

-- ─── ORGANIZATION INVITATIONS ─────────────────────────────────

CREATE TABLE IF NOT EXISTS public.organization_invitations (
    id              BIGSERIAL         PRIMARY KEY,
    organization_id BIGINT            NOT NULL REFERENCES public.organizations(id) ON DELETE CASCADE,
    invite_code     TEXT              NOT NULL UNIQUE CHECK (invite_code <> ''::text),
    email           TEXT              NOT NULL CHECK (email <> ''::text),
    role            org_role          NOT NULL DEFAULT 'member',
    invited_by      BIGINT            REFERENCES public.users(id) ON DELETE SET NULL,
    status          invitation_status NOT NULL DEFAULT 'pending',
    expires_at      TIMESTAMPTZ       NOT NULL,
    created_at      TIMESTAMPTZ       NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_org_invitations_code
    ON public.organization_invitations(invite_code);

CREATE INDEX IF NOT EXISTS idx_org_invitations_org
    ON public.organization_invitations(organization_id);
