CREATE TABLE IF NOT EXISTS tables (
    id SERIAL PRIMARY KEY,
    business_id INT NOT NULL,
    qr_code_url VARCHAR(255) NOT NULL,
    status VARCHAR(25) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (business_id) REFERENCES businesses(id) ON DELETE CASCADE
);