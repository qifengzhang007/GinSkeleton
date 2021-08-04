--
-- PostgreSQL database dump
--

-- Dumped from database version 10.17
-- Dumped by pg_dump version 10.17

-- Started on 2021-08-04 12:22:01

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
-- TOC entry 2856 (class 0 OID 0)
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
                                            token character varying(600) DEFAULT ''::character varying NOT NULL,
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
-- TOC entry 2857 (class 0 OID 0)
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
-- TOC entry 2858 (class 0 OID 0)
-- Dependencies: 197
-- Name: tb_users_id_seq; Type: SEQUENCE OWNED BY; Schema: web; Owner: postgres
--

ALTER SEQUENCE web.tb_users_id_seq OWNED BY web.tb_users.id;


--
-- TOC entry 2696 (class 2604 OID 16612)
-- Name: tb_auth_casbin_rule id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_auth_casbin_rule ALTER COLUMN id SET DEFAULT nextval('web.tb_auth_casbin_rule_id_seq'::regclass);


--
-- TOC entry 2708 (class 2604 OID 16632)
-- Name: tb_oauth_access_tokens id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_oauth_access_tokens ALTER COLUMN id SET DEFAULT nextval('web.tb_oauth_access_tokens_id_seq'::regclass);


--
-- TOC entry 2685 (class 2604 OID 16594)
-- Name: tb_users id; Type: DEFAULT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_users ALTER COLUMN id SET DEFAULT nextval('web.tb_users_id_seq'::regclass);


--
-- TOC entry 2848 (class 0 OID 16609)
-- Dependencies: 200
-- Data for Name: tb_auth_casbin_rule; Type: TABLE DATA; Schema: web; Owner: postgres
--



--
-- TOC entry 2850 (class 0 OID 16629)
-- Dependencies: 202
-- Data for Name: tb_oauth_access_tokens; Type: TABLE DATA; Schema: web; Owner: postgres
--



--
-- TOC entry 2846 (class 0 OID 16591)
-- Dependencies: 198
-- Data for Name: tb_users; Type: TABLE DATA; Schema: web; Owner: postgres
--



--
-- TOC entry 2859 (class 0 OID 0)
-- Dependencies: 199
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_auth_casbin_rule_id_seq', 1, false);


--
-- TOC entry 2860 (class 0 OID 0)
-- Dependencies: 201
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_oauth_access_tokens_id_seq', 2, true);


--
-- TOC entry 2861 (class 0 OID 0)
-- Dependencies: 197
-- Name: tb_users_id_seq; Type: SEQUENCE SET; Schema: web; Owner: postgres
--

SELECT pg_catalog.setval('web.tb_users_id_seq', 8, true);


--
-- TOC entry 2721 (class 2606 OID 16626)
-- Name: tb_auth_casbin_rule tb_auth_casbin_rule_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_auth_casbin_rule
    ADD CONSTRAINT tb_auth_casbin_rule_pkey PRIMARY KEY (id);


--
-- TOC entry 2723 (class 2606 OID 16647)
-- Name: tb_oauth_access_tokens tb_oauth_access_tokens_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_oauth_access_tokens
    ADD CONSTRAINT tb_oauth_access_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 2718 (class 2606 OID 16606)
-- Name: tb_users tb_users_pkey; Type: CONSTRAINT; Schema: web; Owner: postgres
--

ALTER TABLE ONLY web.tb_users
    ADD CONSTRAINT tb_users_pkey PRIMARY KEY (id);


--
-- TOC entry 2719 (class 1259 OID 16662)
-- Name: idx_web_tb_auth_casbin_rule; Type: INDEX; Schema: web; Owner: postgres
--

CREATE UNIQUE INDEX idx_web_tb_auth_casbin_rule ON web.tb_auth_casbin_rule USING btree (ptype, v0, v1, v2, v3, v4, v5);


-- Completed on 2021-08-04 12:22:02

--
-- PostgreSQL database dump complete
--

