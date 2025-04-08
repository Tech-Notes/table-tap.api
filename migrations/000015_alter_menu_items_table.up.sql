-- update photo_url column to photo_id int nullable
ALTER TABLE menu_items
DROP COLUMN photo_url,
ADD COLUMN photo_id INT NULL;