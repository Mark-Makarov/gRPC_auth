INSERT INTO apps (id, name, secret)
VALUES (1, 'tests', 'tests-secret')
ON CONFLICT DO NOTHING;