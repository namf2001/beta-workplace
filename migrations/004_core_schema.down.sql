-- Rollback migration 004

DROP TABLE IF EXISTS friendships;
DROP TABLE IF EXISTS workplace_invitations;
DROP TABLE IF EXISTS workplace_members;
DROP TABLE IF EXISTS workplaces;
DROP TABLE IF EXISTS organization_members;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS verification_codes;
DROP TABLE IF EXISTS auth_providers;

ALTER TABLE users
    DROP COLUMN IF EXISTS is_active,
    DROP COLUMN IF EXISTS email_verified_at,
    DROP COLUMN IF EXISTS avatar_url,
    DROP COLUMN IF EXISTS display_name,
    DROP COLUMN IF EXISTS username;

DROP TYPE IF EXISTS friendship_status;
DROP TYPE IF EXISTS workplace_size;
DROP TYPE IF EXISTS workplace_role;
DROP TYPE IF EXISTS org_role;
DROP TYPE IF EXISTS verification_code_type;
