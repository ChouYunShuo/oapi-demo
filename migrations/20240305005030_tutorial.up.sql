CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS idm_users (
  uuid uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  username varchar(255) NOT NULL UNIQUE,
  password bytea,
  first_name varchar(255) NOT NULL,
  last_name varchar(255),
  created_at timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
  last_modified_at timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
  deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS idm_user_emails (
  uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) NOT NULL,
  user_uuid uuid NOT NULL references idm_users(uuid),
  created_at timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
  deleted_at timestamp
);
