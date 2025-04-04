ALTER TABLE businesses_users
ADD FOREIGN KEY (role_id) REFERENCES businesses_roles(id)
ON DELETE SET NULL; 