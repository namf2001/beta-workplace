-- Test data for users repository tests
INSERT INTO users (id, email, name, password, image, email_verified, created_at, updated_at)
VALUES
    (1, 'test1@example.com', 'Test User 1', 'hashedpassword1', 'https://example.com/1.png', NULL, NOW(), NOW()),
    (2, 'test2@example.com', 'Test User 2', 'hashedpassword2', 'https://example.com/2.png', NULL, NOW(), NOW()),
    (3, 'test3@example.com', 'Test User 3', 'hashedpassword3', 'https://example.com/3.png', NULL, NOW(), NOW());
