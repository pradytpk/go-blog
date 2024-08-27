CREATE TABLE IF NOT EXISTS posts(
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    user_id bigint NOT NULL,
   content TEXT NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);
