-- v25 (compatible with v19+): Store emoticon status for portals
ALTER TABLE portal ADD COLUMN has_emoticons BOOLEAN NOT NULL DEFAULT false;
