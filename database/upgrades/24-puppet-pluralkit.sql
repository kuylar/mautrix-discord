-- v24 (compatible with v19+): Store PluralKit info for puppets
ALTER TABLE puppet ADD COLUMN is_plural_kit_user BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE puppet ADD COLUMN is_plural_kit_proxy BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE puppet ADD COLUMN plural_kit_id VARCHAR(255) NOT NULL DEFAULT '';
