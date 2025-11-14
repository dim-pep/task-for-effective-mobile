CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255) NOT NULL,
    price_kopecks INT NOT NULL CHECK (price_kopecks > 0),
    user_id UUID NOT NULL,
    start_date DATE NOT NULL CHECK (EXTRACT(DAY FROM start_date) = 1),
    end_date DATE CHECK (end_date IS NULL OR EXTRACT(DAY FROM end_date) = 1)
);