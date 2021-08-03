--
-- PostgreSQL database dump
--

-- Dumped from database version 10.17
-- Dumped by pg_dump version 10.17

-- Started on 2021-08-03 15:49:53

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
-- TOC entry 6 (class 2615 OID 16393)
-- Name: db_goskeleton; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA db_goskeleton;


ALTER SCHEMA db_goskeleton OWNER TO postgres;

--
-- TOC entry 2856 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA db_goskeleton; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA db_goskeleton IS '创建测试数据库';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 201 (class 1259 OID 16519)
-- Name: tb_auth_casbin_rule; Type: TABLE; Schema: db_goskeleton; Owner: postgres
--

CREATE TABLE db_goskeleton.tb_auth_casbin_rule (
                                                   id integer NOT NULL,
                                                   ptype character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p0 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p1 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p2 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p3 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p4 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   p5 character varying(100) DEFAULT ''::character varying NOT NULL,
                                                   created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                                   updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE db_goskeleton.tb_auth_casbin_rule OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 16517)
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE; Schema: db_goskeleton; Owner: postgres
--

CREATE SEQUENCE db_goskeleton.tb_auth_casbin_rule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE db_goskeleton.tb_auth_casbin_rule_id_seq OWNER TO postgres;

--
-- TOC entry 2857 (class 0 OID 0)
-- Dependencies: 200
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: db_goskeleton; Owner: postgres
--

ALTER SEQUENCE db_goskeleton.tb_auth_casbin_rule_id_seq OWNED BY db_goskeleton.tb_auth_casbin_rule.id;


--
-- TOC entry 203 (class 1259 OID 16539)
-- Name: tb_oauth_access_tokens; Type: TABLE; Schema: db_goskeleton; Owner: postgres
--

CREATE TABLE db_goskeleton.tb_oauth_access_tokens (
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


ALTER TABLE db_goskeleton.tb_oauth_access_tokens OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 16537)
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE; Schema: db_goskeleton; Owner: postgres
--

CREATE SEQUENCE db_goskeleton.tb_oauth_access_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE db_goskeleton.tb_oauth_access_tokens_id_seq OWNER TO postgres;

--
-- TOC entry 2858 (class 0 OID 0)
-- Dependencies: 202
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: db_goskeleton; Owner: postgres
--

ALTER SEQUENCE db_goskeleton.tb_oauth_access_tokens_id_seq OWNED BY db_goskeleton.tb_oauth_access_tokens.id;


--
-- TOC entry 199 (class 1259 OID 16446)
-- Name: tb_users; Type: TABLE; Schema: db_goskeleton; Owner: postgres
--

CREATE TABLE db_goskeleton.tb_users (
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


ALTER TABLE db_goskeleton.tb_users OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 16444)
-- Name: tb_users_id_seq; Type: SEQUENCE; Schema: db_goskeleton; Owner: postgres
--

CREATE SEQUENCE db_goskeleton.tb_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE db_goskeleton.tb_users_id_seq OWNER TO postgres;

--
-- TOC entry 2859 (class 0 OID 0)
-- Dependencies: 198
-- Name: tb_users_id_seq; Type: SEQUENCE OWNED BY; Schema: db_goskeleton; Owner: postgres
--

ALTER SEQUENCE db_goskeleton.tb_users_id_seq OWNED BY db_goskeleton.tb_users.id;


--
-- TOC entry 2697 (class 2604 OID 16522)
-- Name: tb_auth_casbin_rule id; Type: DEFAULT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_auth_casbin_rule ALTER COLUMN id SET DEFAULT nextval('db_goskeleton.tb_auth_casbin_rule_id_seq'::regclass);


--
-- TOC entry 2707 (class 2604 OID 16542)
-- Name: tb_oauth_access_tokens id; Type: DEFAULT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_oauth_access_tokens ALTER COLUMN id SET DEFAULT nextval('db_goskeleton.tb_oauth_access_tokens_id_seq'::regclass);


--
-- TOC entry 2686 (class 2604 OID 16449)
-- Name: tb_users id; Type: DEFAULT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_users ALTER COLUMN id SET DEFAULT nextval('db_goskeleton.tb_users_id_seq'::regclass);


--
-- TOC entry 2848 (class 0 OID 16519)
-- Dependencies: 201
-- Data for Name: tb_auth_casbin_rule; Type: TABLE DATA; Schema: db_goskeleton; Owner: postgres
--



--
-- TOC entry 2850 (class 0 OID 16539)
-- Dependencies: 203
-- Data for Name: tb_oauth_access_tokens; Type: TABLE DATA; Schema: db_goskeleton; Owner: postgres
--



--
-- TOC entry 2846 (class 0 OID 16446)
-- Dependencies: 199
-- Data for Name: tb_users; Type: TABLE DATA; Schema: db_goskeleton; Owner: postgres
--

INSERT INTO db_goskeleton.tb_users VALUES (1, 'admin', '123456', '张三丰', '1363991xxxx', 1, '备注信息', NULL, '127.0.0.1', 0, '2021-08-03 15:15:00.954634', '2021-08-03 15:15:00.954634');


--
-- TOC entry 2860 (class 0 OID 0)
-- Dependencies: 200
-- Name: tb_auth_casbin_rule_id_seq; Type: SEQUENCE SET; Schema: db_goskeleton; Owner: postgres
--

SELECT pg_catalog.setval('db_goskeleton.tb_auth_casbin_rule_id_seq', 1, false);


--
-- TOC entry 2861 (class 0 OID 0)
-- Dependencies: 202
-- Name: tb_oauth_access_tokens_id_seq; Type: SEQUENCE SET; Schema: db_goskeleton; Owner: postgres
--

SELECT pg_catalog.setval('db_goskeleton.tb_oauth_access_tokens_id_seq', 1, false);


--
-- TOC entry 2862 (class 0 OID 0)
-- Dependencies: 198
-- Name: tb_users_id_seq; Type: SEQUENCE SET; Schema: db_goskeleton; Owner: postgres
--

SELECT pg_catalog.setval('db_goskeleton.tb_users_id_seq', 1, true);


--
-- TOC entry 2721 (class 2606 OID 16536)
-- Name: tb_auth_casbin_rule tb_auth_casbin_rule_pkey; Type: CONSTRAINT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_auth_casbin_rule
    ADD CONSTRAINT tb_auth_casbin_rule_pkey PRIMARY KEY (id);


--
-- TOC entry 2723 (class 2606 OID 16557)
-- Name: tb_oauth_access_tokens tb_oauth_access_tokens_pkey; Type: CONSTRAINT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_oauth_access_tokens
    ADD CONSTRAINT tb_oauth_access_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 2719 (class 2606 OID 16461)
-- Name: tb_users tb_users_pkey; Type: CONSTRAINT; Schema: db_goskeleton; Owner: postgres
--

ALTER TABLE ONLY db_goskeleton.tb_users
    ADD CONSTRAINT tb_users_pkey PRIMARY KEY (id);


-- Completed on 2021-08-03 15:49:53

--
-- PostgreSQL database dump complete
--

