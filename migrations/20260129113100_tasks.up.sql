CREATE TABLE tasks (

                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           task TEXT NOT NULL,
                           status VARCHAR(50) DEFAULT 'active'
);