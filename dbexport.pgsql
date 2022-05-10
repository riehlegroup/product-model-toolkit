--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: components; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.components (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    product_refer integer,
    uid text NOT NULL,
    name text,
    package text,
    version text,
    license_refer integer,
    copyright text,
    artifact_refer integer
);


ALTER TABLE public.components OWNER TO postgres;

--
-- Name: components_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.components_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.components_id_seq OWNER TO postgres;

--
-- Name: components_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.components_id_seq OWNED BY public.components.id;


--
-- Name: dep_graphs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dep_graphs (
    dependencies jsonb DEFAULT '{}'::jsonb NOT NULL
);


ALTER TABLE public.dep_graphs OWNER TO postgres;

--
-- Name: licenses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.licenses (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    spdx_id text NOT NULL,
    declared_licesne text,
    concluded_license text
);


ALTER TABLE public.licenses OWNER TO postgres;

--
-- Name: licenses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.licenses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.licenses_id_seq OWNER TO postgres;

--
-- Name: licenses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.licenses_id_seq OWNED BY public.licenses.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    version text,
    license text,
    vcs text,
    description text,
    comment text,
    home_page_url text,
    external_reference text,
    clearing_state text,
    dep_graph_refer integer,
    infrastructure text
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: usage_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usage_types (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    product_refer integer,
    name text
);


ALTER TABLE public.usage_types OWNER TO postgres;

--
-- Name: usage_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usage_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.usage_types_id_seq OWNER TO postgres;

--
-- Name: usage_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usage_types_id_seq OWNED BY public.usage_types.id;


--
-- Name: components id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.components ALTER COLUMN id SET DEFAULT nextval('public.components_id_seq'::regclass);


--
-- Name: licenses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.licenses ALTER COLUMN id SET DEFAULT nextval('public.licenses_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: usage_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usage_types ALTER COLUMN id SET DEFAULT nextval('public.usage_types_id_seq'::regclass);


--
-- Data for Name: components; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.components (id, created_at, updated_at, deleted_at, product_refer, uid, name, package, version, license_refer, copyright, artifact_refer) FROM stdin;
1	2022-01-09 21:15:23.573028+00	2022-01-09 21:15:23.573028+00	\N	1	Package-hello-go-src	hello-go-src			1		0
2	2022-01-18 20:51:24.338726+00	2022-01-18 20:51:24.338726+00	\N	2			doctrine/inflector	2.0.4	2		0
3	2022-01-18 20:51:24.340896+00	2022-01-18 20:51:24.340896+00	\N	2			doctrine/instantiator	1.4.0	3		0
4	2022-01-18 20:51:24.341858+00	2022-01-18 20:51:24.341858+00	\N	2			hamcrest/hamcrest-php	v2.0.1	4		0
5	2022-01-18 20:51:24.342842+00	2022-01-18 20:51:24.342842+00	\N	2			illuminate/collections	v8.78.1	5		0
6	2022-01-18 20:51:24.343815+00	2022-01-18 20:51:24.343815+00	\N	2			illuminate/console	v8.78.1	6		0
7	2022-01-18 20:51:24.344797+00	2022-01-18 20:51:24.344797+00	\N	2			illuminate/contracts	v8.78.1	7		0
8	2022-01-18 20:51:24.345691+00	2022-01-18 20:51:24.345691+00	\N	2			illuminate/macroable	v8.78.1	8		0
9	2022-01-18 20:51:24.346668+00	2022-01-18 20:51:24.346668+00	\N	2			illuminate/support	v8.78.1	9		0
10	2022-01-18 20:51:24.347658+00	2022-01-18 20:51:24.347658+00	\N	2			mockery/mockery	1.4.4	10		0
11	2022-01-18 20:51:24.348552+00	2022-01-18 20:51:24.348552+00	\N	2			myclabs/deep-copy	1.10.2	11		0
12	2022-01-18 20:51:24.350474+00	2022-01-18 20:51:24.350474+00	\N	2			nesbot/carbon	2.55.2	12		0
13	2022-01-18 20:51:24.353384+00	2022-01-18 20:51:24.353384+00	\N	2			nikic/php-parser	v4.13.2	13		0
14	2022-01-18 20:51:24.354873+00	2022-01-18 20:51:24.354873+00	\N	2			phar-io/manifest	2.0.3	14		0
15	2022-01-18 20:51:24.356185+00	2022-01-18 20:51:24.356185+00	\N	2			phar-io/version	3.1.0	15		0
16	2022-01-18 20:51:24.357544+00	2022-01-18 20:51:24.357544+00	\N	2			phpdocumentor/reflection-common	2.2.0	16		0
17	2022-01-18 20:51:24.358849+00	2022-01-18 20:51:24.358849+00	\N	2			phpdocumentor/reflection-docblock	5.3.0	17		0
18	2022-01-18 20:51:24.360385+00	2022-01-18 20:51:24.360385+00	\N	2			phpdocumentor/type-resolver	1.6.0	18		0
19	2022-01-18 20:51:24.361647+00	2022-01-18 20:51:24.361647+00	\N	2			phpspec/prophecy	v1.15.0	19		0
20	2022-01-18 20:51:24.362604+00	2022-01-18 20:51:24.362604+00	\N	2			phpunit/php-code-coverage	9.2.10	20		0
21	2022-01-18 20:51:24.363667+00	2022-01-18 20:51:24.363667+00	\N	2			phpunit/php-file-iterator	3.0.6	21		0
22	2022-01-18 20:51:24.364835+00	2022-01-18 20:51:24.364835+00	\N	2			phpunit/php-invoker	3.1.1	22		0
23	2022-01-18 20:51:24.365952+00	2022-01-18 20:51:24.365952+00	\N	2			phpunit/php-text-template	2.0.4	23		0
24	2022-01-18 20:51:24.367878+00	2022-01-18 20:51:24.367878+00	\N	2			phpunit/php-timer	5.0.3	24		0
25	2022-01-18 20:51:24.369051+00	2022-01-18 20:51:24.369051+00	\N	2			phpunit/phpunit	9.5.11	25		0
26	2022-01-18 20:51:24.371128+00	2022-01-18 20:51:24.371128+00	\N	2			psr/container	1.1.2	26		0
27	2022-01-18 20:51:24.373155+00	2022-01-18 20:51:24.373155+00	\N	2			psr/simple-cache	1.0.1	27		0
28	2022-01-18 20:51:24.375099+00	2022-01-18 20:51:24.375099+00	\N	2			psy/psysh	v0.11.1	28		0
29	2022-01-18 20:51:24.376523+00	2022-01-18 20:51:24.376523+00	\N	2			sebastian/cli-parser	1.0.1	29		0
30	2022-01-18 20:51:24.378231+00	2022-01-18 20:51:24.378231+00	\N	2			sebastian/code-unit	1.0.8	30		0
31	2022-01-18 20:51:24.380101+00	2022-01-18 20:51:24.380101+00	\N	2			sebastian/code-unit-reverse-lookup	2.0.3	31		0
32	2022-01-18 20:51:24.381361+00	2022-01-18 20:51:24.381361+00	\N	2			sebastian/comparator	4.0.6	32		0
33	2022-01-18 20:51:24.382729+00	2022-01-18 20:51:24.382729+00	\N	2			sebastian/complexity	2.0.2	33		0
34	2022-01-18 20:51:24.384648+00	2022-01-18 20:51:24.384648+00	\N	2			sebastian/diff	4.0.4	34		0
35	2022-01-18 20:51:24.386025+00	2022-01-18 20:51:24.386025+00	\N	2			sebastian/environment	5.1.3	35		0
36	2022-01-18 20:51:24.387308+00	2022-01-18 20:51:24.387308+00	\N	2			sebastian/exporter	4.0.4	36		0
37	2022-01-18 20:51:24.388693+00	2022-01-18 20:51:24.388693+00	\N	2			sebastian/global-state	5.0.3	37		0
38	2022-01-18 20:51:24.391169+00	2022-01-18 20:51:24.391169+00	\N	2			sebastian/lines-of-code	1.0.3	38		0
39	2022-01-18 20:51:24.393242+00	2022-01-18 20:51:24.393242+00	\N	2			sebastian/object-enumerator	4.0.4	39		0
40	2022-01-18 20:51:24.394922+00	2022-01-18 20:51:24.394922+00	\N	2			sebastian/object-reflector	2.0.4	40		0
41	2022-01-18 20:51:24.396658+00	2022-01-18 20:51:24.396658+00	\N	2			sebastian/recursion-context	4.0.4	41		0
42	2022-01-18 20:51:24.39835+00	2022-01-18 20:51:24.39835+00	\N	2			sebastian/resource-operations	3.0.3	42		0
43	2022-01-18 20:51:24.400455+00	2022-01-18 20:51:24.400455+00	\N	2			sebastian/type	2.3.4	43		0
44	2022-01-18 20:51:24.402512+00	2022-01-18 20:51:24.402512+00	\N	2			sebastian/version	3.0.2	44		0
45	2022-01-18 20:51:24.404059+00	2022-01-18 20:51:24.404059+00	\N	2			symfony/console	v5.4.2	45		0
46	2022-01-18 20:51:24.405451+00	2022-01-18 20:51:24.405451+00	\N	2			symfony/deprecation-contracts	v3.0.0	46		0
47	2022-01-18 20:51:24.407287+00	2022-01-18 20:51:24.407287+00	\N	2			symfony/polyfill-ctype	v1.24.0	47		0
48	2022-01-18 20:51:24.408991+00	2022-01-18 20:51:24.408991+00	\N	2			symfony/polyfill-intl-grapheme	v1.24.0	48		0
49	2022-01-18 20:51:24.410294+00	2022-01-18 20:51:24.410294+00	\N	2			symfony/polyfill-intl-normalizer	v1.24.0	49		0
50	2022-01-18 20:51:24.411523+00	2022-01-18 20:51:24.411523+00	\N	2			symfony/polyfill-mbstring	v1.24.0	50		0
51	2022-01-18 20:51:24.412672+00	2022-01-18 20:51:24.412672+00	\N	2			symfony/polyfill-php73	v1.24.0	51		0
52	2022-01-18 20:51:24.413756+00	2022-01-18 20:51:24.413756+00	\N	2			symfony/polyfill-php80	v1.24.0	52		0
53	2022-01-18 20:51:24.414808+00	2022-01-18 20:51:24.414808+00	\N	2			symfony/process	v5.4.2	53		0
54	2022-01-18 20:51:24.415848+00	2022-01-18 20:51:24.415848+00	\N	2			symfony/service-contracts	v2.4.1	54		0
55	2022-01-18 20:51:24.417885+00	2022-01-18 20:51:24.417885+00	\N	2			symfony/string	v6.0.2	55		0
56	2022-01-18 20:51:24.419794+00	2022-01-18 20:51:24.419794+00	\N	2			symfony/translation	v6.0.2	56		0
57	2022-01-18 20:51:24.421607+00	2022-01-18 20:51:24.421607+00	\N	2			symfony/translation-contracts	v3.0.0	57		0
58	2022-01-18 20:51:24.422912+00	2022-01-18 20:51:24.422912+00	\N	2			symfony/var-dumper	v6.0.2	58		0
59	2022-01-18 20:51:24.424736+00	2022-01-18 20:51:24.424736+00	\N	2			theseer/tokenizer	1.2.1	59		0
60	2022-01-18 20:51:24.425994+00	2022-01-18 20:51:24.425994+00	\N	2			voku/portable-ascii	1.5.6	60		0
61	2022-01-18 20:51:24.42742+00	2022-01-18 20:51:24.42742+00	\N	2			webmozart/assert	1.10.0	61		0
62	2022-01-18 20:55:07.462046+00	2022-01-18 20:55:07.462046+00	\N	3			doctrine/inflector	2.0.4	62		0
63	2022-01-18 20:55:07.462935+00	2022-01-18 20:55:07.462935+00	\N	3			doctrine/instantiator	1.4.0	63		0
64	2022-01-18 20:55:07.463643+00	2022-01-18 20:55:07.463643+00	\N	3			hamcrest/hamcrest-php	v2.0.1	64		0
65	2022-01-18 20:55:07.464289+00	2022-01-18 20:55:07.464289+00	\N	3			illuminate/collections	v8.78.1	65		0
66	2022-01-18 20:55:07.465023+00	2022-01-18 20:55:07.465023+00	\N	3			illuminate/console	v8.78.1	66		0
67	2022-01-18 20:55:07.465904+00	2022-01-18 20:55:07.465904+00	\N	3			illuminate/contracts	v8.78.1	67		0
68	2022-01-18 20:55:07.466716+00	2022-01-18 20:55:07.466716+00	\N	3			illuminate/macroable	v8.78.1	68		0
69	2022-01-18 20:55:07.467828+00	2022-01-18 20:55:07.467828+00	\N	3			illuminate/support	v8.78.1	69		0
70	2022-01-18 20:55:07.469324+00	2022-01-18 20:55:07.469324+00	\N	3			mockery/mockery	1.4.4	70		0
71	2022-01-18 20:55:07.470302+00	2022-01-18 20:55:07.470302+00	\N	3			myclabs/deep-copy	1.10.2	71		0
72	2022-01-18 20:55:07.47114+00	2022-01-18 20:55:07.47114+00	\N	3			nesbot/carbon	2.55.2	72		0
73	2022-01-18 20:55:07.472051+00	2022-01-18 20:55:07.472051+00	\N	3			nikic/php-parser	v4.13.2	73		0
74	2022-01-18 20:55:07.473301+00	2022-01-18 20:55:07.473301+00	\N	3			phar-io/manifest	2.0.3	74		0
75	2022-01-18 20:55:07.474432+00	2022-01-18 20:55:07.474432+00	\N	3			phar-io/version	3.1.0	75		0
76	2022-01-18 20:55:07.475745+00	2022-01-18 20:55:07.475745+00	\N	3			phpdocumentor/reflection-common	2.2.0	76		0
77	2022-01-18 20:55:07.476643+00	2022-01-18 20:55:07.476643+00	\N	3			phpdocumentor/reflection-docblock	5.3.0	77		0
78	2022-01-18 20:55:07.477661+00	2022-01-18 20:55:07.477661+00	\N	3			phpdocumentor/type-resolver	1.6.0	78		0
79	2022-01-18 20:55:07.478564+00	2022-01-18 20:55:07.478564+00	\N	3			phpspec/prophecy	v1.15.0	79		0
80	2022-01-18 20:55:07.479464+00	2022-01-18 20:55:07.479464+00	\N	3			phpunit/php-code-coverage	9.2.10	80		0
81	2022-01-18 20:55:07.480313+00	2022-01-18 20:55:07.480313+00	\N	3			phpunit/php-file-iterator	3.0.6	81		0
82	2022-01-18 20:55:07.481141+00	2022-01-18 20:55:07.481141+00	\N	3			phpunit/php-invoker	3.1.1	82		0
83	2022-01-18 20:55:07.48224+00	2022-01-18 20:55:07.48224+00	\N	3			phpunit/php-text-template	2.0.4	83		0
84	2022-01-18 20:55:07.48315+00	2022-01-18 20:55:07.48315+00	\N	3			phpunit/php-timer	5.0.3	84		0
85	2022-01-18 20:55:07.484753+00	2022-01-18 20:55:07.484753+00	\N	3			phpunit/phpunit	9.5.11	85		0
86	2022-01-18 20:55:07.486304+00	2022-01-18 20:55:07.486304+00	\N	3			psr/container	1.1.2	86		0
87	2022-01-18 20:55:07.488008+00	2022-01-18 20:55:07.488008+00	\N	3			psr/simple-cache	1.0.1	87		0
88	2022-01-18 20:55:07.489831+00	2022-01-18 20:55:07.489831+00	\N	3			psy/psysh	v0.11.1	88		0
89	2022-01-18 20:55:07.491535+00	2022-01-18 20:55:07.491535+00	\N	3			sebastian/cli-parser	1.0.1	89		0
90	2022-01-18 20:55:07.493309+00	2022-01-18 20:55:07.493309+00	\N	3			sebastian/code-unit	1.0.8	90		0
91	2022-01-18 20:55:07.49484+00	2022-01-18 20:55:07.49484+00	\N	3			sebastian/code-unit-reverse-lookup	2.0.3	91		0
92	2022-01-18 20:55:07.495938+00	2022-01-18 20:55:07.495938+00	\N	3			sebastian/comparator	4.0.6	92		0
93	2022-01-18 20:55:07.496845+00	2022-01-18 20:55:07.496845+00	\N	3			sebastian/complexity	2.0.2	93		0
94	2022-01-18 20:55:07.497752+00	2022-01-18 20:55:07.497752+00	\N	3			sebastian/diff	4.0.4	94		0
95	2022-01-18 20:55:07.498781+00	2022-01-18 20:55:07.498781+00	\N	3			sebastian/environment	5.1.3	95		0
96	2022-01-18 20:55:07.499876+00	2022-01-18 20:55:07.499876+00	\N	3			sebastian/exporter	4.0.4	96		0
97	2022-01-18 20:55:07.502057+00	2022-01-18 20:55:07.502057+00	\N	3			sebastian/global-state	5.0.3	97		0
98	2022-01-18 20:55:07.50372+00	2022-01-18 20:55:07.50372+00	\N	3			sebastian/lines-of-code	1.0.3	98		0
99	2022-01-18 20:55:07.505216+00	2022-01-18 20:55:07.505216+00	\N	3			sebastian/object-enumerator	4.0.4	99		0
100	2022-01-18 20:55:07.507588+00	2022-01-18 20:55:07.507588+00	\N	3			sebastian/object-reflector	2.0.4	100		0
101	2022-01-18 20:55:07.509428+00	2022-01-18 20:55:07.509428+00	\N	3			sebastian/recursion-context	4.0.4	101		0
102	2022-01-18 20:55:07.510979+00	2022-01-18 20:55:07.510979+00	\N	3			sebastian/resource-operations	3.0.3	102		0
103	2022-01-18 20:55:07.512157+00	2022-01-18 20:55:07.512157+00	\N	3			sebastian/type	2.3.4	103		0
104	2022-01-18 20:55:07.513221+00	2022-01-18 20:55:07.513221+00	\N	3			sebastian/version	3.0.2	104		0
105	2022-01-18 20:55:07.514201+00	2022-01-18 20:55:07.514201+00	\N	3			symfony/console	v5.4.2	105		0
106	2022-01-18 20:55:07.515239+00	2022-01-18 20:55:07.515239+00	\N	3			symfony/deprecation-contracts	v3.0.0	106		0
107	2022-01-18 20:55:07.516857+00	2022-01-18 20:55:07.516857+00	\N	3			symfony/polyfill-ctype	v1.24.0	107		0
108	2022-01-18 20:55:07.518954+00	2022-01-18 20:55:07.518954+00	\N	3			symfony/polyfill-intl-grapheme	v1.24.0	108		0
109	2022-01-18 20:55:07.520507+00	2022-01-18 20:55:07.520507+00	\N	3			symfony/polyfill-intl-normalizer	v1.24.0	109		0
110	2022-01-18 20:55:07.523282+00	2022-01-18 20:55:07.523282+00	\N	3			symfony/polyfill-mbstring	v1.24.0	110		0
111	2022-01-18 20:55:07.52498+00	2022-01-18 20:55:07.52498+00	\N	3			symfony/polyfill-php73	v1.24.0	111		0
112	2022-01-18 20:55:07.526511+00	2022-01-18 20:55:07.526511+00	\N	3			symfony/polyfill-php80	v1.24.0	112		0
113	2022-01-18 20:55:07.528328+00	2022-01-18 20:55:07.528328+00	\N	3			symfony/process	v5.4.2	113		0
114	2022-01-18 20:55:07.52974+00	2022-01-18 20:55:07.52974+00	\N	3			symfony/service-contracts	v2.4.1	114		0
115	2022-01-18 20:55:07.530904+00	2022-01-18 20:55:07.530904+00	\N	3			symfony/string	v6.0.2	115		0
116	2022-01-18 20:55:07.531983+00	2022-01-18 20:55:07.531983+00	\N	3			symfony/translation	v6.0.2	116		0
117	2022-01-18 20:55:07.533066+00	2022-01-18 20:55:07.533066+00	\N	3			symfony/translation-contracts	v3.0.0	117		0
118	2022-01-18 20:55:07.535003+00	2022-01-18 20:55:07.535003+00	\N	3			symfony/var-dumper	v6.0.2	118		0
119	2022-01-18 20:55:07.536567+00	2022-01-18 20:55:07.536567+00	\N	3			theseer/tokenizer	1.2.1	119		0
120	2022-01-18 20:55:07.538347+00	2022-01-18 20:55:07.538347+00	\N	3			voku/portable-ascii	1.5.6	120		0
121	2022-01-18 20:55:07.540384+00	2022-01-18 20:55:07.540384+00	\N	3			webmozart/assert	1.10.0	121		0
122	2022-03-29 13:39:51.870539+00	2022-03-29 13:39:51.870539+00	\N	4	Package-hello-go-bin	hello-go-bin			122		0
123	2022-04-11 21:45:59.427109+00	2022-04-11 21:45:59.427109+00	\N	5	Package-hello-go-bin	hello-go-bin			123		0
124	2022-04-11 21:46:32.584755+00	2022-04-11 21:46:32.584755+00	\N	6	Package-hello-go-bin	hello-go-bin			124		0
125	2022-05-10 00:49:01.568498+00	2022-05-10 00:49:01.568498+00	\N	7	Package-go.fmt	go.fmt		1.15.4	125		0
126	2022-05-10 00:49:01.571924+00	2022-05-10 00:49:01.571924+00	\N	7	Package-go.reflect	go.reflect		1.15.4	126		0
127	2022-05-10 00:49:01.574803+00	2022-05-10 00:49:01.574803+00	\N	7	Package-go.strconv	go.strconv		1.15.4	127		0
128	2022-05-10 00:49:01.578111+00	2022-05-10 00:49:01.578111+00	\N	7	Package-godist	go-1.15	Ubuntu snap distribution of Golang v1.15.4 linux/amd64	1.15.4	128		0
129	2022-05-10 00:49:01.580756+00	2022-05-10 00:49:01.580756+00	\N	7	Package-go-compiler	go		1.15.4	129		0
\.


--
-- Data for Name: dep_graphs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dep_graphs (dependencies) FROM stdin;
\.


--
-- Data for Name: licenses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.licenses (id, created_at, updated_at, deleted_at, spdx_id, declared_licesne, concluded_license) FROM stdin;
1	2022-01-09 21:15:23.571996+00	2022-01-09 21:15:23.571996+00	\N	Package-hello-go-src	GPL-3.0-or-later	NOASSERTION
2	2022-01-18 20:51:24.337405+00	2022-01-18 20:51:24.337405+00	\N	MIT		
3	2022-01-18 20:51:24.34037+00	2022-01-18 20:51:24.34037+00	\N	MIT		
4	2022-01-18 20:51:24.341458+00	2022-01-18 20:51:24.341458+00	\N	BSD-3-Clause		
5	2022-01-18 20:51:24.342367+00	2022-01-18 20:51:24.342367+00	\N	MIT		
6	2022-01-18 20:51:24.343356+00	2022-01-18 20:51:24.343356+00	\N	MIT		
7	2022-01-18 20:51:24.34438+00	2022-01-18 20:51:24.34438+00	\N	MIT		
8	2022-01-18 20:51:24.345246+00	2022-01-18 20:51:24.345246+00	\N	MIT		
9	2022-01-18 20:51:24.346224+00	2022-01-18 20:51:24.346224+00	\N	MIT		
10	2022-01-18 20:51:24.347197+00	2022-01-18 20:51:24.347197+00	\N	BSD-3-Clause		
11	2022-01-18 20:51:24.348143+00	2022-01-18 20:51:24.348143+00	\N	MIT		
12	2022-01-18 20:51:24.348987+00	2022-01-18 20:51:24.348987+00	\N	MIT		
13	2022-01-18 20:51:24.352601+00	2022-01-18 20:51:24.352601+00	\N	BSD-3-Clause		
14	2022-01-18 20:51:24.354183+00	2022-01-18 20:51:24.354183+00	\N	BSD-3-Clause		
15	2022-01-18 20:51:24.355554+00	2022-01-18 20:51:24.355554+00	\N	BSD-3-Clause		
16	2022-01-18 20:51:24.356804+00	2022-01-18 20:51:24.356804+00	\N	MIT		
17	2022-01-18 20:51:24.358307+00	2022-01-18 20:51:24.358307+00	\N	MIT		
18	2022-01-18 20:51:24.359688+00	2022-01-18 20:51:24.359688+00	\N	MIT		
19	2022-01-18 20:51:24.361119+00	2022-01-18 20:51:24.361119+00	\N	MIT		
20	2022-01-18 20:51:24.362191+00	2022-01-18 20:51:24.362191+00	\N	BSD-3-Clause		
21	2022-01-18 20:51:24.363198+00	2022-01-18 20:51:24.363198+00	\N	BSD-3-Clause		
22	2022-01-18 20:51:24.364268+00	2022-01-18 20:51:24.364268+00	\N	BSD-3-Clause		
23	2022-01-18 20:51:24.36541+00	2022-01-18 20:51:24.36541+00	\N	BSD-3-Clause		
24	2022-01-18 20:51:24.367053+00	2022-01-18 20:51:24.367053+00	\N	BSD-3-Clause		
25	2022-01-18 20:51:24.368514+00	2022-01-18 20:51:24.368514+00	\N	BSD-3-Clause		
26	2022-01-18 20:51:24.36979+00	2022-01-18 20:51:24.36979+00	\N	MIT		
27	2022-01-18 20:51:24.372188+00	2022-01-18 20:51:24.372188+00	\N	MIT		
28	2022-01-18 20:51:24.374118+00	2022-01-18 20:51:24.374118+00	\N	MIT		
29	2022-01-18 20:51:24.375895+00	2022-01-18 20:51:24.375895+00	\N	BSD-3-Clause		
30	2022-01-18 20:51:24.377433+00	2022-01-18 20:51:24.377433+00	\N	BSD-3-Clause		
31	2022-01-18 20:51:24.379339+00	2022-01-18 20:51:24.379339+00	\N	BSD-3-Clause		
32	2022-01-18 20:51:24.380767+00	2022-01-18 20:51:24.380767+00	\N	BSD-3-Clause		
33	2022-01-18 20:51:24.381988+00	2022-01-18 20:51:24.381988+00	\N	BSD-3-Clause		
34	2022-01-18 20:51:24.383775+00	2022-01-18 20:51:24.383775+00	\N	BSD-3-Clause		
35	2022-01-18 20:51:24.385404+00	2022-01-18 20:51:24.385404+00	\N	BSD-3-Clause		
36	2022-01-18 20:51:24.386661+00	2022-01-18 20:51:24.386661+00	\N	BSD-3-Clause		
37	2022-01-18 20:51:24.388002+00	2022-01-18 20:51:24.388002+00	\N	BSD-3-Clause		
38	2022-01-18 20:51:24.390025+00	2022-01-18 20:51:24.390025+00	\N	BSD-3-Clause		
39	2022-01-18 20:51:24.392421+00	2022-01-18 20:51:24.392421+00	\N	BSD-3-Clause		
40	2022-01-18 20:51:24.394156+00	2022-01-18 20:51:24.394156+00	\N	BSD-3-Clause		
41	2022-01-18 20:51:24.395824+00	2022-01-18 20:51:24.395824+00	\N	BSD-3-Clause		
42	2022-01-18 20:51:24.397565+00	2022-01-18 20:51:24.397565+00	\N	BSD-3-Clause		
43	2022-01-18 20:51:24.399052+00	2022-01-18 20:51:24.399052+00	\N	BSD-3-Clause		
44	2022-01-18 20:51:24.401461+00	2022-01-18 20:51:24.401461+00	\N	BSD-3-Clause		
45	2022-01-18 20:51:24.403409+00	2022-01-18 20:51:24.403409+00	\N	MIT		
46	2022-01-18 20:51:24.40478+00	2022-01-18 20:51:24.40478+00	\N	MIT		
47	2022-01-18 20:51:24.406469+00	2022-01-18 20:51:24.406469+00	\N	MIT		
48	2022-01-18 20:51:24.408177+00	2022-01-18 20:51:24.408177+00	\N	MIT		
49	2022-01-18 20:51:24.409764+00	2022-01-18 20:51:24.409764+00	\N	MIT		
50	2022-01-18 20:51:24.410956+00	2022-01-18 20:51:24.410956+00	\N	MIT		
51	2022-01-18 20:51:24.41212+00	2022-01-18 20:51:24.41212+00	\N	MIT		
52	2022-01-18 20:51:24.41326+00	2022-01-18 20:51:24.41326+00	\N	MIT		
53	2022-01-18 20:51:24.414319+00	2022-01-18 20:51:24.414319+00	\N	MIT		
54	2022-01-18 20:51:24.415362+00	2022-01-18 20:51:24.415362+00	\N	MIT		
55	2022-01-18 20:51:24.416829+00	2022-01-18 20:51:24.416829+00	\N	MIT		
56	2022-01-18 20:51:24.418796+00	2022-01-18 20:51:24.418796+00	\N	MIT		
57	2022-01-18 20:51:24.420859+00	2022-01-18 20:51:24.420859+00	\N	MIT		
58	2022-01-18 20:51:24.42225+00	2022-01-18 20:51:24.42225+00	\N	MIT		
59	2022-01-18 20:51:24.423892+00	2022-01-18 20:51:24.423892+00	\N	BSD-3-Clause		
60	2022-01-18 20:51:24.425399+00	2022-01-18 20:51:24.425399+00	\N	MIT		
61	2022-01-18 20:51:24.426687+00	2022-01-18 20:51:24.426687+00	\N	MIT		
62	2022-01-18 20:55:07.461435+00	2022-01-18 20:55:07.461435+00	\N	MIT		
63	2022-01-18 20:55:07.462589+00	2022-01-18 20:55:07.462589+00	\N	MIT		
64	2022-01-18 20:55:07.463301+00	2022-01-18 20:55:07.463301+00	\N	BSD-3-Clause		
65	2022-01-18 20:55:07.463989+00	2022-01-18 20:55:07.463989+00	\N	MIT		
66	2022-01-18 20:55:07.464669+00	2022-01-18 20:55:07.464669+00	\N	MIT		
67	2022-01-18 20:55:07.465427+00	2022-01-18 20:55:07.465427+00	\N	MIT		
68	2022-01-18 20:55:07.466341+00	2022-01-18 20:55:07.466341+00	\N	MIT		
69	2022-01-18 20:55:07.467245+00	2022-01-18 20:55:07.467245+00	\N	MIT		
70	2022-01-18 20:55:07.468463+00	2022-01-18 20:55:07.468463+00	\N	BSD-3-Clause		
71	2022-01-18 20:55:07.469931+00	2022-01-18 20:55:07.469931+00	\N	MIT		
72	2022-01-18 20:55:07.470775+00	2022-01-18 20:55:07.470775+00	\N	MIT		
73	2022-01-18 20:55:07.471567+00	2022-01-18 20:55:07.471567+00	\N	BSD-3-Clause		
74	2022-01-18 20:55:07.472782+00	2022-01-18 20:55:07.472782+00	\N	BSD-3-Clause		
75	2022-01-18 20:55:07.473858+00	2022-01-18 20:55:07.473858+00	\N	BSD-3-Clause		
76	2022-01-18 20:55:07.475232+00	2022-01-18 20:55:07.475232+00	\N	MIT		
77	2022-01-18 20:55:07.476202+00	2022-01-18 20:55:07.476202+00	\N	MIT		
78	2022-01-18 20:55:07.47721+00	2022-01-18 20:55:07.47721+00	\N	MIT		
79	2022-01-18 20:55:07.478113+00	2022-01-18 20:55:07.478113+00	\N	MIT		
80	2022-01-18 20:55:07.479016+00	2022-01-18 20:55:07.479016+00	\N	BSD-3-Clause		
81	2022-01-18 20:55:07.479907+00	2022-01-18 20:55:07.479907+00	\N	BSD-3-Clause		
82	2022-01-18 20:55:07.480754+00	2022-01-18 20:55:07.480754+00	\N	BSD-3-Clause		
83	2022-01-18 20:55:07.481744+00	2022-01-18 20:55:07.481744+00	\N	BSD-3-Clause		
84	2022-01-18 20:55:07.482743+00	2022-01-18 20:55:07.482743+00	\N	BSD-3-Clause		
85	2022-01-18 20:55:07.483746+00	2022-01-18 20:55:07.483746+00	\N	BSD-3-Clause		
86	2022-01-18 20:55:07.485564+00	2022-01-18 20:55:07.485564+00	\N	MIT		
87	2022-01-18 20:55:07.487261+00	2022-01-18 20:55:07.487261+00	\N	MIT		
88	2022-01-18 20:55:07.488921+00	2022-01-18 20:55:07.488921+00	\N	MIT		
89	2022-01-18 20:55:07.490699+00	2022-01-18 20:55:07.490699+00	\N	BSD-3-Clause		
90	2022-01-18 20:55:07.492359+00	2022-01-18 20:55:07.492359+00	\N	BSD-3-Clause		
91	2022-01-18 20:55:07.494239+00	2022-01-18 20:55:07.494239+00	\N	BSD-3-Clause		
92	2022-01-18 20:55:07.495494+00	2022-01-18 20:55:07.495494+00	\N	BSD-3-Clause		
93	2022-01-18 20:55:07.496464+00	2022-01-18 20:55:07.496464+00	\N	BSD-3-Clause		
94	2022-01-18 20:55:07.497329+00	2022-01-18 20:55:07.497329+00	\N	BSD-3-Clause		
95	2022-01-18 20:55:07.498252+00	2022-01-18 20:55:07.498252+00	\N	BSD-3-Clause		
96	2022-01-18 20:55:07.499376+00	2022-01-18 20:55:07.499376+00	\N	BSD-3-Clause		
97	2022-01-18 20:55:07.501131+00	2022-01-18 20:55:07.501131+00	\N	BSD-3-Clause		
98	2022-01-18 20:55:07.502875+00	2022-01-18 20:55:07.502875+00	\N	BSD-3-Clause		
99	2022-01-18 20:55:07.504589+00	2022-01-18 20:55:07.504589+00	\N	BSD-3-Clause		
100	2022-01-18 20:55:07.506692+00	2022-01-18 20:55:07.506692+00	\N	BSD-3-Clause		
101	2022-01-18 20:55:07.508695+00	2022-01-18 20:55:07.508695+00	\N	BSD-3-Clause		
102	2022-01-18 20:55:07.510273+00	2022-01-18 20:55:07.510273+00	\N	BSD-3-Clause		
103	2022-01-18 20:55:07.511652+00	2022-01-18 20:55:07.511652+00	\N	BSD-3-Clause		
104	2022-01-18 20:55:07.512735+00	2022-01-18 20:55:07.512735+00	\N	BSD-3-Clause		
105	2022-01-18 20:55:07.513776+00	2022-01-18 20:55:07.513776+00	\N	MIT		
106	2022-01-18 20:55:07.514775+00	2022-01-18 20:55:07.514775+00	\N	MIT		
107	2022-01-18 20:55:07.516198+00	2022-01-18 20:55:07.516198+00	\N	MIT		
108	2022-01-18 20:55:07.51808+00	2022-01-18 20:55:07.51808+00	\N	MIT		
109	2022-01-18 20:55:07.51984+00	2022-01-18 20:55:07.51984+00	\N	MIT		
110	2022-01-18 20:55:07.521312+00	2022-01-18 20:55:07.521312+00	\N	MIT		
111	2022-01-18 20:55:07.524063+00	2022-01-18 20:55:07.524063+00	\N	MIT		
112	2022-01-18 20:55:07.525825+00	2022-01-18 20:55:07.525825+00	\N	MIT		
113	2022-01-18 20:55:07.527549+00	2022-01-18 20:55:07.527549+00	\N	MIT		
114	2022-01-18 20:55:07.529108+00	2022-01-18 20:55:07.529108+00	\N	MIT		
115	2022-01-18 20:55:07.530354+00	2022-01-18 20:55:07.530354+00	\N	MIT		
116	2022-01-18 20:55:07.53148+00	2022-01-18 20:55:07.53148+00	\N	MIT		
117	2022-01-18 20:55:07.532552+00	2022-01-18 20:55:07.532552+00	\N	MIT		
118	2022-01-18 20:55:07.533936+00	2022-01-18 20:55:07.533936+00	\N	MIT		
119	2022-01-18 20:55:07.53589+00	2022-01-18 20:55:07.53589+00	\N	BSD-3-Clause		
120	2022-01-18 20:55:07.537573+00	2022-01-18 20:55:07.537573+00	\N	MIT		
121	2022-01-18 20:55:07.539407+00	2022-01-18 20:55:07.539407+00	\N	MIT		
122	2022-03-29 13:39:51.868318+00	2022-03-29 13:39:51.868318+00	\N	Package-hello-go-bin	NOASSERTION	GPL-3.0-or-later AND LicenseRef-Golang-BSD-plus-Patents
123	2022-04-11 21:45:59.425917+00	2022-04-11 21:45:59.425917+00	\N	Package-hello-go-bin	NOASSERTION	GPL-3.0-or-later AND LicenseRef-Golang-BSD-plus-Patents
124	2022-04-11 21:46:32.579717+00	2022-04-11 21:46:32.579717+00	\N	Package-hello-go-bin	NOASSERTION	GPL-3.0-or-later AND LicenseRef-Golang-BSD-plus-Patents
125	2022-05-10 00:49:01.566333+00	2022-05-10 00:49:01.566333+00	\N	Package-go.fmt	NOASSERTION	NOASSERTION
126	2022-05-10 00:49:01.570571+00	2022-05-10 00:49:01.570571+00	\N	Package-go.reflect	NOASSERTION	NOASSERTION
127	2022-05-10 00:49:01.573524+00	2022-05-10 00:49:01.573524+00	\N	Package-go.strconv	NOASSERTION	NOASSERTION
128	2022-05-10 00:49:01.576872+00	2022-05-10 00:49:01.576872+00	\N	Package-godist	LicenseRef-Golang-BSD-plus-Patents	NOASSERTION
129	2022-05-10 00:49:01.579725+00	2022-05-10 00:49:01.579725+00	\N	Package-go-compiler	NOASSERTION	NOASSERTION
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, created_at, updated_at, deleted_at, name, version, license, vcs, description, comment, home_page_url, external_reference, clearing_state, dep_graph_refer, infrastructure) FROM stdin;
1	2022-01-09 21:15:23.570467+00	2022-01-09 21:15:23.570467+00	\N	hello-go-src	SPDX-2.2					https://swinslow.net/spdx-examples/example6/hello-go-src-v2			0	
3	2022-01-18 20:55:07.460821+00	2022-01-18 20:55:07.460821+00	\N	product-dvULfhSSpN	VERSION HERE	Apache-2.0		DESCRIPTION HERE					0	
2	2022-01-18 20:51:24.327525+00	2022-01-18 20:51:24.327525+00	\N	product-dKuoXOduMS	VERSION HERE	LGPL-2.1-only		DESCRIPTION HERE					0	
4	2022-03-29 13:39:51.861319+00	2022-03-29 13:39:51.861319+00	\N	hello-go-bin	SPDX-2.2	LGPL-2.1-or-later				https://swinslow.net/spdx-examples/example6/hello-go-bin-v2	[]string		0	
5	2022-04-11 21:45:59.422241+00	2022-04-11 21:45:59.422241+00	\N	hello-go-bin	SPDX-2.2	LGPL-3.0-only				https://swinslow.net/spdx-examples/example6/hello-go-bin-v2	[]string		0	
6	2022-04-11 21:46:32.57499+00	2022-04-11 21:46:32.57499+00	\N	hello-go-bin	SPDX-2.2	GPL-3.0-or-later				https://swinslow.net/spdx-examples/example6/hello-go-bin-v2	[]string		0	
7	2022-05-10 00:49:01.56122+00	2022-05-10 00:49:01.56122+00	\N	go-lib	SPDX-2.2					https://swinslow.net/spdx-examples/example6/go-lib-v2			0	
\.


--
-- Data for Name: usage_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.usage_types (id, created_at, updated_at, deleted_at, product_refer, name) FROM stdin;
\.


--
-- Name: components_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.components_id_seq', 129, true);


--
-- Name: licenses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.licenses_id_seq', 129, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 7, true);


--
-- Name: usage_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.usage_types_id_seq', 1, false);


--
-- Name: components components_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.components
    ADD CONSTRAINT components_pkey PRIMARY KEY (id);


--
-- Name: licenses licenses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.licenses
    ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: usage_types usage_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usage_types
    ADD CONSTRAINT usage_types_pkey PRIMARY KEY (id);


--
-- Name: idx_components_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_components_deleted_at ON public.components USING btree (deleted_at);


--
-- Name: idx_licenses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_licenses_deleted_at ON public.licenses USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_usage_types_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_usage_types_deleted_at ON public.usage_types USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

