--
-- PostgreSQL database dump
--

-- Dumped from database version 10.17
-- Dumped by pg_dump version 10.17

-- Started on 2021-08-03 18:27:18

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 5 (class 2615 OID 16570)
-- Name: web; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA web;


ALTER SCHEMA web OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 204 (class 1259 OID 16652)
-- Name: _tb_auth_casbin_rule; Type: TABLE; Schema: web; Owner: postgres
--

CREATE TABLE web._tb_auth_casbin_rule (
                                          id bigint NOT NULL,
                                          ptype character varying(100),
                                          v0 character varying(100),
                                          v1 character varying(100),
                                          v2 character varying(100),
                                          v3 character varying(100),
                                          v4 character varying(100),
                                          v5 character varying(100)
);


ALTER TABLE web._tb_auth_casbin_rule OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16650)
-- Name: _tb_auth_casbin_rule_id_seq; Type: SEQUENCE; Schema: web; Owner: postgres
--

CREATE SEQUENCE web._tb_auth_casbin_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE web._tb_auth_casbin_rule_id_seq OWNER TO postgres;

--
-- TOC entry 2867 (class 0 OID 0)
-- Dependencies: 203
-- Name: _tb_auth_casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: web; Owner: postgres
--

ALTER SEQUENCE web._tb_auth_casbin_rule_id_seq OWNED BY web._tb_auth_casbin_rule.id;


--
-- TOC entry 200 (class 1259 OID 16609)
-- Name: tb_auth_casbin_rule; Type: TABLE; Schema: web; Owner: postgres
--

CREATE TABLE web.tb_auth_casbin_rule (
                                         id integer NOT NULL,
                                         ptype character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p0 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p1 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p2 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p3 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p4 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         p5 character varying(100) DEFAULT ''::character varying NOT NULL,
                                         created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                         updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                         v0 character varying(100),
                                         v1 character varying(100),
                                         v2 character varying(100),
                                         v3 character varying(100),
                                         v4 character varying(100),
                                         v5 character varying(100)
);


ALTER TABLE web.tb_auth_casbin_rule OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 16607)
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE; Schema: web; Owner: postgres
--

CREATE SEQUENCE web.tb_auth_casbin_rule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE web.tb_auth_casbin_rule_id_seq OWNER TO postgres;

--
-- TOC entry 2868 (class 0 OID 0)
-- Dependencies: 199
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: web; Owner: postgres
--

ALTER SEQUENCE web.tb_auth_casbin_rule_id_seq OWNED BY web.tb_auth_casbin_rule.id;


--
-- TOC entry 202 (class 1259 OID 16629)
-- Name: tb_oauth_access_tokens; Type: TABLE; Schema: web; Owner: postgres
--

CREATE TABLE web.tb_oauth_access_tokens (
                                            id integer NOT NULL,
                                            fr_user_id integer DEFAULT 0,
                                            client_id integer DEFAULT 1,
                                            token character varying(520) DEFAULT ''::character varying NOT NULL,
                                            action_name character varying(100) DEFAULT ''::character varying NOT NULL,
                                            scopes character varying(100) DEFAULT '*'::character varying NOT NULL,
                                            revoked smallint DEFAULT 0 NOT NULL,
                                            client_ip character varying(20) DEFAULT ''::character varying,
                                            expires_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                            created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                            updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE web.tb_oauth_access_tokens OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16627)
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE; Schema: web; Owner: postgres
--

CREATE SEQUENCE web.tb_oauth_access_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE web.tb_oauth_access_tokens_id_seq OWNER TO postgres;

--
-- TOC entry 2869 (class 0 OID 0)
-- Dependencies: 201
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: web; Owner: postgres
--

ALTER SEQUENCE web.tb_oauth_access_tokens_id_seq OWNED BY web.tb_oauth_access_tokens.id;


--
-- TOC entry 198 (class 1259 OID 16591)
-- Name: tb_users; Type: TABLE; Schema: web; Owner: postgres
--

CREATE TABLE web.tb_users (
                              id integer NOT NULL,
                              user_name character varying(30) DEFAULT ''::character varying NOT NULL,
                              pass character varying(128) DEFAULT ''::character varying NOT NULL,
                              real_name character varying(30) DEFAULT ''::character varying,
                              phone character(11) DEFAULT ''::bpchar,
                              status smallint DEFAULT 1,
                              remark character varying(120) DEFAULT ''::character varying,
                              last_login_time timestamp without time zone,
                              last_login_ip character varying(20) DEFAULT ''::character varying,
                              login_times integer DEFAULT 0,
                              created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE web.tb_users OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 16589)
-- Name: tb_users_id_seq; Type: SEQUENCE; Schema: web; Owner: postgres
--

CREATE SEQUENCE web.tb_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE web.tb_users_id_seq OWNER TO postgres;

--
-- TOC entry 2870 (class 0 OID 0)
-- Dependencies: 197
-- Name: tb_users_id_seq; Type: SEQUENCE OWNED BY; Schema: web; Owner: postgres
--

ALTER SEQUENCE web.tb_users_id_seq OWNED BY web.tb_users.id;


--
-- TOC entry 2724 (class 2604 OID 16655)
-- Name: _tb_auth_casbin_rule id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web._tb_auth_casbin_rule ALTER COLUMN id SET DEFAULT nextval('web._tb_auth_casbin_rule_id_seq'::regclass);


--
-- TOC entry 2703 (class 2604 OID 16612)
-- Name: tb_auth_casbin_rule id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_auth_casbin_rule ALTER COLUMN id SET DEFAULT nextval('web.tb_auth_casbin_rule_id_seq'::regclass);


--
-- TOC entry 2715 (class 2604 OID 16632)
-- Name: tb_oauth_access_tokens id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_oauth_access_tokens ALTER COLUMN id SET DEFAULT nextval('web.tb_oauth_access_tokens_id_seq'::regclass);


--
-- TOC entry 2692 (class 2604 OID 16594)
-- Name: tb_users id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_users ALTER COLUMN id SET DEFAULT nextval('web.tb_users_id_seq'::regclass);


--
-- TOC entry 2861 (class 0 OID 16652)
-- Dependencies: 204
-- Data for Name: _tb_auth_casbin_rule; Type: TABLE DATA; Schema: web; Owner: postgres
--



--
-- TOC entry 2857 (class 0 OID 16609)
-- Dependencies: 200
-- Data for Name: tb_auth_casbin_rule; Type: TABLE DATA; Schema: web; Owner: postgres
--



--
-- TOC entry 2859 (class 0 OID 16629)
-- Dependencies: 202
-- Data for Name: tb_oauth_access_tokens; Type: TABLE DATA; Schema: web; Owner: postgres
--

INSERT INTO web.tb_oauth_access_tokens VALUES (1, 3, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VyX25hbWUiOiJhZG1pbiIsInBob25lIjoiICAgICAgICAgICAiLCJleHAiOjE2MjgwMjE3NDYsIm5iZiI6MTYyNzk4NTI4OX0.5YtCQyWrMh-yH6cgihOJS0i3AFjEjXNxxc87bhQhXQs', 'refresh', '*', 0, '127.0.0.1', '2021-08-04 04:15:46', '2021-08-03 18:08:19.103374', '2021-08-03 18:15:46.174945');


--
-- TOC entry 2855 (class 0 OID 16591)
-- Dependencies: 198
-- Data for Name: tb_users; Type: TABLE DATA; Schema: web; Owner: postgres
--

INSERT INTO web.tb_users VALUES (1, 'admin2', '123456', '张三丰', '1363991xxxx', 1, '备注信息', NULL, '127.0.0.1', 0, '2021-08-03 18:04:08.595046', '2021-08-03 18:04:08.595046');
INSERT INTO web.tb_users VALUES (3, 'admin', '188bda0c10088d7c2e6d7c00592679e7', '', '           ', 1, '', NULL, '127.0.0.1', 0, '2021-08-03 18:06:55.424588', '2021-08-03 18:06:55.424588');
INSERT INTO web.tb_users VALUES (4, 'admin007', '188bda0c10088d7c2e6d7c00592679e7', 'postgres', '13639917240', 1, '测试添加-通过postgres', NULL, '', 0, '2021-08-03 18:12:01.5921', '2021-08-03 18:12:01.5921');


--
-- TOC entry 2871 (class 0 OID 0)
-- Dependencies: 203
-- Name: _tb_auth_casbin_rule_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web._tb_auth_casbin_rule_id_seq', 1, false);


--
-- TOC entry 2872 (class 0 OID 0)
-- Dependencies: 199
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_auth_casbin_rule_id_seq', 1, false);


--
-- TOC entry 2873 (class 0 OID 0)
-- Dependencies: 201
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_oauth_access_tokens_id_seq', 1, true);


--
-- TOC entry 2874 (class 0 OID 0)
-- Dependencies: 197
-- Name: tb_users_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_users_id_seq', 4, true);


--
-- TOC entry 2732 (class 2606 OID 16660)
-- Name: _tb_auth_casbin_rule _tb_auth_casbin_rule_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web._tb_auth_casbin_rule
    ADD CONSTRAINT _tb_auth_casbin_rule_pkey PRIMARY KEY (id);


--
-- TOC entry 2728 (class 2606 OID 16626)
-- Name: tb_auth_casbin_rule tb_auth_casbin_rule_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_auth_casbin_rule
    ADD CONSTRAINT tb_auth_casbin_rule_pkey PRIMARY KEY (id);


--
-- TOC entry 2730 (class 2606 OID 16647)
-- Name: tb_oauth_access_tokens tb_oauth_access_tokens_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_oauth_access_tokens
    ADD CONSTRAINT tb_oauth_access_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 2726 (class 2606 OID 16606)
-- Name: tb_users tb_users_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_users
    ADD CONSTRAINT tb_users_pkey PRIMARY KEY (id);


-- Completed on 2021-08-03 18:27:19

--
-- PostgreSQL database dump complete
--

