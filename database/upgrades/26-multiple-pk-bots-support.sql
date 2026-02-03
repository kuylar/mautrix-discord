-- v26 (compatible with v19+): Add support for multiple Plural proxy bots (currently only PK nd /plu/ral)
ALTER TABLE puppet ADD COLUMN plural_state TEXT NOT NULL DEFAULT 'single';
UPDATE puppet SET plural_state = CASE
    WHEN is_plural_kit_user THEN 'pk:user'
    WHEN is_plural_kit_proxy THEN 'pk:proxy'
    ELSE 'single'
END
