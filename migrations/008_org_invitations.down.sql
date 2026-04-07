-- ============================================================
-- Migration 008 Down: Rollback Organization Invitations
-- ============================================================

DROP TABLE IF EXISTS public.organization_invitations;

ALTER TABLE public.organizations
    DROP COLUMN IF EXISTS industry,
    DROP COLUMN IF EXISTS size,
    DROP COLUMN IF EXISTS region;

DROP TYPE IF EXISTS invitation_status;
