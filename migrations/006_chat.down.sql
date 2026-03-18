-- Rollback migration 006

DROP TABLE IF EXISTS message_reactions;
DROP TABLE IF EXISTS message_attachments;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS channel_members;
DROP TABLE IF EXISTS channels;

DROP TYPE IF EXISTS channel_type;
