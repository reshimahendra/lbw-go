-- DROP TABLE public.membership_status;
CREATE TABLE public.membership_status (
	id int2 NOT NULL, -- membership id will refer to user id. using 1:1 relation
	status_name varchar(30) NOT NULL,
	description text NULL,
	CONSTRAINT membership_status_name_un UNIQUE (status_name),
	CONSTRAINT membership_status_pk PRIMARY KEY (id)
);
COMMENT ON TABLE public.membership_status IS 'membership status';

-- Column comments
COMMENT ON COLUMN public.membership_status.id IS 'membership id will refer to user id. using 1:1 relation';

-- Permissions
ALTER TABLE public.membership_status OWNER TO lotus;
GRANT ALL ON TABLE public.membership_status TO lotus;
-- ----------------------------------------------



-- DROP TABLE public.membership_mail_app_type;
CREATE TABLE public.membership_mail_app_type (
	id int2 NOT NULL,
	type_name varchar(50) NOT NULL, -- Membership type name
	description text NULL, -- description for the membership service
	price money NOT NULL DEFAULT 0, -- price of the service
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- date time the membership created
	updated_at timestamp NULL, -- datetime the membershipservice was updated
	deleted_at timestamp NULL, -- soft delete for the membership status
	CONSTRAINT membership_mail_app_type_name_un UNIQUE (type_name),
	CONSTRAINT membership_mail_app_type_pk PRIMARY KEY (id)
);
COMMENT ON TABLE public.membership_mail_app_type IS 'Membership of mail app service';

-- Column comments
COMMENT ON COLUMN public.membership_mail_app_type.type_name IS 'Membership type name';
COMMENT ON COLUMN public.membership_mail_app_type.description IS 'description for the membership service';
COMMENT ON COLUMN public.membership_mail_app_type.price IS 'price of the service';
COMMENT ON COLUMN public.membership_mail_app_type.created_at IS 'date time the membership created';
COMMENT ON COLUMN public.membership_mail_app_type.updated_at IS 'datetime the membershipservice was updated';
COMMENT ON COLUMN public.membership_mail_app_type.deleted_at IS 'soft delete for the membership status';

-- Permissions
ALTER TABLE public.membership_mail_app_type OWNER TO lotus;
GRANT ALL ON TABLE public.membership_mail_app_type TO lotus;
-- ----------------------------------------------



-- DROP TABLE public.membership_mail_app;
CREATE TABLE public.membership_mail_app (
	id uuid NOT NULL,
	type_id int2 NOT NULL DEFAULT 0, -- refer to membership mail app type
	status_id int2 NOT NULL DEFAULT 0, -- refer to membership status
	price money NOT NULL DEFAULT 0, -- membership subscription price
	last_paid_amount money NOT NULL DEFAULT 0,
	last_paid_at timestamp NULL,
	subscribed_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT membership_mail_app_pk PRIMARY KEY (id),
	CONSTRAINT membership_mail_app_status_fk FOREIGN KEY (status_id) REFERENCES public.membership_status(id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT membership_mail_app_type_fk FOREIGN KEY (type_id) REFERENCES public.membership_mail_app_type(id) ON UPDATE CASCADE,
	CONSTRAINT membership_mail_app_user_fk FOREIGN KEY (id) REFERENCES public.users(id) ON UPDATE CASCADE
);
COMMENT ON TABLE public.membership_mail_app IS 'user membership to mail app service for mail service handler on static site';

-- Column comments
COMMENT ON COLUMN public.membership_mail_app.type_id IS 'refer to membership mail app type';
COMMENT ON COLUMN public.membership_mail_app.status_id IS 'refer to membership status';
COMMENT ON COLUMN public.membership_mail_app.price IS 'membership subscription price';

-- Permissions
ALTER TABLE public.membership_mail_app OWNER TO lotus;
GRANT ALL ON TABLE public.membership_mail_app TO lotus;
-- ----------------------------------------------



-- DROP TABLE public.membership_mail_app_config;
CREATE TABLE public.membership_mail_app_config (
	cfg_id int2 NOT NULL,
	id uuid NOT NULL,
	config_name varchar(30) NOT NULL,
	default_config bool NOT NULL DEFAULT false, -- whether this config is default config or not
	smtp_server varchar(100) NOT NULL,
	smtp_port int4 NOT NULL,
	smtp_username varchar(75) NOT NULL,
	smtp_password varchar(50) NOT NULL,
	smtp_sender_email varchar(75) NOT NULL,
	smtp_sender_identity varchar(50) NOT NULL,
	active_status bool NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	deleted_at timestamp NULL, -- soft delete features
	CONSTRAINT membership_mail_app_config_name_un UNIQUE (config_name),
	CONSTRAINT membership_mail_app_config_pk PRIMARY KEY (cfg_id, id),
	CONSTRAINT membership_mail_app_config_users_fk FOREIGN KEY (id) REFERENCES public.users(id) ON DELETE SET NULL ON UPDATE CASCADE
);
COMMENT ON TABLE public.membership_mail_app_config IS 'user configuration for their mail server configuration';

-- Column comments
COMMENT ON COLUMN public.membership_mail_app_config.default_config IS 'whether this config is default config or not';
COMMENT ON COLUMN public.membership_mail_app_config.deleted_at IS 'soft delete features';

-- Permissions
ALTER TABLE public.membership_mail_app_config OWNER TO lotus;
GRANT ALL ON TABLE public.membership_mail_app_config TO lotus;
-- ----------------------------------------------

