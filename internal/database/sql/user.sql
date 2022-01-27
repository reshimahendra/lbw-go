-- DROP TABLE public.user_role;
CREATE TABLE public.user_role (
	id int2 NOT NULL,
	role_name varchar(30) NOT NULL,
	description text NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT user_role_name_un UNIQUE (role_name),
	CONSTRAINT user_role_pk PRIMARY KEY (id)
);
COMMENT ON TABLE public.user_role IS 'user role containing role hold by the user';

-- Permissions
ALTER TABLE public.user_role OWNER TO lotus;
GRANT ALL ON TABLE public.user_role TO lotus;


-- DROP TABLE public.users;
CREATE TABLE public.users (
	id uuid NOT NULL,
	username varchar(30) NOT NULL,
	firstname varchar(30) NOT NULL,
	lastname varchar(30) NULL,
	email varchar(100) NOT NULL,
	passkey varchar(50) NOT NULL,
	user_status_id int2 NOT NULL DEFAULT 0, -- user status
	user_role_id int2 NOT NULL DEFAULT 0, -- user role on system
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	activated_at timestamp NULL, -- account activation datetime
	deleted_at timestamp NULL, -- account (soft) delete datetime
	CONSTRAINT users_email_un UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_username_un UNIQUE (username),
	CONSTRAINT users_user_role_fk FOREIGN KEY (user_role_id) REFERENCES public.user_role(id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT users_user_status_fk FOREIGN KEY (user_status_id) REFERENCES public.user_status(id) ON DELETE SET NULL ON UPDATE CASCADE
);
COMMENT ON TABLE public.users IS 'User table';

-- Column comments
COMMENT ON COLUMN public.users.user_status_id IS 'user status';
COMMENT ON COLUMN public.users.user_role_id IS 'user role on system';
COMMENT ON COLUMN public.users.activated_at IS 'account activation datetime';
COMMENT ON COLUMN public.users.deleted_at IS 'account (soft) delete datetime';

-- Permissions
ALTER TABLE public.users OWNER TO lotus;
GRANT ALL ON TABLE public.users TO lotus;

