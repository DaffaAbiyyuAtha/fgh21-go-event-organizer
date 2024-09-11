create table "wishlist" (
	"id" serial primary key,
	"user_id" int REFERENCES "users"("id"),
    "event_id" int REFERENCES "events"("id")
);