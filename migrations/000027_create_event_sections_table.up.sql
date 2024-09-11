create table "event_sections" (
	"id" serial primary key,
    "name" varchar(50),
    "price" INT,
    "quantity" INT,
    "event_id" int references "events"("id")
)