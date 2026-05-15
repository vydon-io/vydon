ALTER TABLE vydon_api.users
DROP COLUMN IF EXISTS user_type;

ALTER TABLE vydon_api.account_api_keys
DROP COLUMN IF EXISTS user_id;
