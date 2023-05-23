CREATE TABLE "refresh_sessions" (
    "id" SERIAL PRIMARY KEY,
    "id_client" uuid REFERENCES users(id) ON DELETE CASCADE,
    "id_refresh_token" uuid NOT NULL,
    "issued_at" timestamp with time zone NOT NULL DEFAULT now(),
    "expires_in" timestamp with time zone NOT NULL
);