create table "users"(
	"id" serial primary key,
	"username" varchar(80),
	"email" varchar(80),
	"password" varchar(255)
);