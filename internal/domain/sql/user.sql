-- 'user_role' table
CREATE TABLE golang.user_role (
	id serial NOT NULL,
	role_name varchar(30) NOT NULL,
	description text NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL DEFAULT null,
	CONSTRAINT user_role_pk PRIMARY KEY (id),
	CONSTRAINT user_role_name_un UNIQUE (role_name)
);
COMMENT ON TABLE golang.user_role IS 'User role containing role for the user';

-- 'users' table
CREATE TABLE golang.users (
	id uuid NOT NULL,
	username varchar(30) NOT NULL,
	first_name varchar(30) NOT NULL,
	last_name varchar(30) NULL,
	email varchar(75) NOT NULL,
	"password" varchar(100) NOT NULL,
	is_active bool NOT NULL DEFAULT false,
	role_id int NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_username_un UNIQUE (username),
	CONSTRAINT user_email_un UNIQUE (email),
	CONSTRAINT user_userrole_fk FOREIGN KEY (role_id) REFERENCES golang.user_role(id) ON DELETE SET NULL ON UPDATE CASCADE
);
COMMENT ON TABLE golang."user" IS 'Table user will containing user account data';

