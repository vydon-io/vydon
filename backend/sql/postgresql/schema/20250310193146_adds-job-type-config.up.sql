ALTER TABLE vydon_api.jobs
ADD COLUMN jobtype_config JSONB NOT NULL DEFAULT '{}'::JSONB;
