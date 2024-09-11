create table "transactions" (
	"id" serial primary key,
	"event_id" int references "events"("id"),
    "payment_method_id" int references "payment_methods"("id"),
    "user_id" int references "users"("id")
);