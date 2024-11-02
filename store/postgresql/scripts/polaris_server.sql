/*
 Navicat Premium Data Transfer

 Source Server         : postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 140001 (140001)
 Source Host           : 127.0.0.1:5432
 Source Catalog        : polaris_server
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140001 (140001)
 File Encoding         : 65001

 Date: 24/09/2024 22:47:55
*/

-- ----------------------------
-- Table structure for auth_principal
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_principal";
CREATE TABLE "public"."auth_principal" (
  "strategy_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "principal_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "principal_role" int4 NOT NULL
)
;
ALTER TABLE "public"."auth_principal" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_principal"."strategy_id" IS 'Strategy ID';
COMMENT ON COLUMN "public"."auth_principal"."principal_id" IS 'Principal ID';
COMMENT ON COLUMN "public"."auth_principal"."principal_role" IS 'PRINCIPAL type, 1 is User, 2 is Group, 3 is Role';
COMMENT ON TABLE "public"."auth_principal" IS 'Authentication principal table';

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role";
CREATE TABLE "public"."auth_role" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "source" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "role_type" int4 NOT NULL DEFAULT 20,
  "comment" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."auth_role" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_role"."id" IS 'Role ID';
COMMENT ON COLUMN "public"."auth_role"."name" IS 'Role name';
COMMENT ON COLUMN "public"."auth_role"."owner" IS 'Main account ID';
COMMENT ON COLUMN "public"."auth_role"."source" IS 'Role source';
COMMENT ON COLUMN "public"."auth_role"."role_type" IS 'Role type';
COMMENT ON COLUMN "public"."auth_role"."comment" IS 'Description';
COMMENT ON COLUMN "public"."auth_role"."flag" IS 'Whether the rules are valid, 0 is valid, 1 is invalid';
COMMENT ON COLUMN "public"."auth_role"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."auth_role"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."auth_role"."metadata" IS 'User metadata';
COMMENT ON TABLE "public"."auth_role" IS 'Authentication role table';

-- ----------------------------
-- Table structure for auth_role_principal
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_principal";
CREATE TABLE "public"."auth_role_principal" (
  "role_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "principal_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "principal_role" int4 NOT NULL
)
;
ALTER TABLE "public"."auth_role_principal" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_role_principal"."role_id" IS 'Role ID';
COMMENT ON COLUMN "public"."auth_role_principal"."principal_id" IS 'Principal ID';
COMMENT ON COLUMN "public"."auth_role_principal"."principal_role" IS 'Principal type, 1 is User, 2 is Group';
COMMENT ON TABLE "public"."auth_role_principal" IS 'Authentication role and principal relation table';

-- ----------------------------
-- Table structure for auth_strategy
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_strategy";
CREATE TABLE "public"."auth_strategy" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "action" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "default_status" int2 NOT NULL DEFAULT 0,
  "source" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."auth_strategy" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_strategy"."id" IS 'Strategy ID';
COMMENT ON COLUMN "public"."auth_strategy"."name" IS 'Policy name';
COMMENT ON COLUMN "public"."auth_strategy"."action" IS 'Read and write permission for this policy';
COMMENT ON COLUMN "public"."auth_strategy"."owner" IS 'The account ID to which this policy is';
COMMENT ON COLUMN "public"."auth_strategy"."comment" IS 'Description';
COMMENT ON COLUMN "public"."auth_strategy"."default_status" IS 'Default status flag';
COMMENT ON COLUMN "public"."auth_strategy"."source" IS 'Policy rule source';
COMMENT ON COLUMN "public"."auth_strategy"."revision" IS 'Authentication rule version';
COMMENT ON COLUMN "public"."auth_strategy"."flag" IS 'Validity flag';
COMMENT ON COLUMN "public"."auth_strategy"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."auth_strategy"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."auth_strategy"."metadata" IS 'Policy rule metadata';
COMMENT ON TABLE "public"."auth_strategy" IS 'Authentication strategy table';

-- ----------------------------
-- Table structure for auth_strategy_function
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_strategy_function";
CREATE TABLE "public"."auth_strategy_function" (
  "strategy_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "function" varchar(256) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."auth_strategy_function" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_strategy_function"."strategy_id" IS 'Strategy ID';
COMMENT ON COLUMN "public"."auth_strategy_function"."function" IS 'Server provider function name';
COMMENT ON TABLE "public"."auth_strategy_function" IS 'Authentication strategy functions';

-- ----------------------------
-- Table structure for auth_strategy_label
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_strategy_label";
CREATE TABLE "public"."auth_strategy_label" (
  "strategy_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "key" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "value" text COLLATE "pg_catalog"."default" NOT NULL,
  "compare_type" varchar(128) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."auth_strategy_label" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_strategy_label"."strategy_id" IS 'Strategy ID';
COMMENT ON COLUMN "public"."auth_strategy_label"."key" IS 'Tag key';
COMMENT ON COLUMN "public"."auth_strategy_label"."value" IS 'Tag value';
COMMENT ON COLUMN "public"."auth_strategy_label"."compare_type" IS 'Tag KV comparison function';
COMMENT ON TABLE "public"."auth_strategy_label" IS 'Authentication strategy labels';

-- ----------------------------
-- Table structure for auth_strategy_resource
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_strategy_resource";
CREATE TABLE "public"."auth_strategy_resource" (
  "strategy_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "res_type" int4 NOT NULL,
  "res_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."auth_strategy_resource" OWNER TO "postgres";
COMMENT ON COLUMN "public"."auth_strategy_resource"."strategy_id" IS 'Strategy ID';
COMMENT ON COLUMN "public"."auth_strategy_resource"."res_type" IS 'Resource Type, Namespaces = 0, Service = 1, configgroups = 2';
COMMENT ON COLUMN "public"."auth_strategy_resource"."res_id" IS 'Resource ID';
COMMENT ON COLUMN "public"."auth_strategy_resource"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."auth_strategy_resource"."mtime" IS 'Last updated time';
COMMENT ON TABLE "public"."auth_strategy_resource" IS 'Authentication strategy resource table';

-- ----------------------------
-- Table structure for circuitbreaker_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."circuitbreaker_rule";
CREATE TABLE "public"."circuitbreaker_rule" (
  "id" varchar(97) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'master'::character varying,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "business" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "department" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "comment" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "inbounds" text COLLATE "pg_catalog"."default" NOT NULL,
  "outbounds" text COLLATE "pg_catalog"."default" NOT NULL,
  "token" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."circuitbreaker_rule" OWNER TO "postgres";
COMMENT ON COLUMN "public"."circuitbreaker_rule"."id" IS 'Melting rule ID';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."version" IS 'Melting rule version, default is MASTER';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."name" IS 'Melting rule name';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."namespace" IS 'Melting rule belongs to namespace';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."business" IS 'Business information of fuse rule';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."department" IS 'Department information for the fuse rule';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."comment" IS 'Description of the fuse rule';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."inbounds" IS 'Service-tuned fuse rule';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."outbounds" IS 'Service motoring fuse rule';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."token" IS 'Token for writing operation check';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."owner" IS 'Melting rule owner information';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."revision" IS 'Melt rule version information';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."flag" IS 'Logic delete flag, 0 means visible, 1 means logically deleted';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."circuitbreaker_rule"."metadata" IS 'Circuit breaker rule metadata';

-- ----------------------------
-- Table structure for circuitbreaker_rule_relation
-- ----------------------------
DROP TABLE IF EXISTS "public"."circuitbreaker_rule_relation";
CREATE TABLE "public"."circuitbreaker_rule_relation" (
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "rule_id" varchar(97) COLLATE "pg_catalog"."default" NOT NULL,
  "rule_version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."circuitbreaker_rule_relation" OWNER TO "postgres";
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."service_id" IS 'Service ID';
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."rule_id" IS 'Melting rule ID';
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."rule_version" IS 'Melting rule version';
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."flag" IS 'Logic delete flag, 0 means visible, 1 means logically deleted';
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."circuitbreaker_rule_relation"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for circuitbreaker_rule_v2
-- ----------------------------
DROP TABLE IF EXISTS "public"."circuitbreaker_rule_v2";
CREATE TABLE "public"."circuitbreaker_rule_v2" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "enable" int4 NOT NULL DEFAULT 0,
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level" int4 NOT NULL,
  "src_service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "src_namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "dst_service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "dst_namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "dst_method" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "config" text COLLATE "pg_catalog"."default",
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "etime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."circuitbreaker_rule_v2" OWNER TO "postgres";
COMMENT ON COLUMN "public"."circuitbreaker_rule_v2"."metadata" IS 'circuit_breaker rule metadata';

-- ----------------------------
-- Table structure for cl5_module
-- ----------------------------
DROP TABLE IF EXISTS "public"."cl5_module";
CREATE TABLE "public"."cl5_module" (
  "module_id" int4 NOT NULL,
  "interface_id" int4 NOT NULL,
  "range_num" int4 NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."cl5_module" OWNER TO "postgres";
COMMENT ON COLUMN "public"."cl5_module"."module_id" IS 'Module ID';
COMMENT ON COLUMN "public"."cl5_module"."interface_id" IS 'Interface ID';
COMMENT ON COLUMN "public"."cl5_module"."range_num" IS 'Range number';
COMMENT ON COLUMN "public"."cl5_module"."mtime" IS 'Last updated time';
COMMENT ON TABLE "public"."cl5_module" IS 'To generate SID';

-- ----------------------------
-- Table structure for client
-- ----------------------------
DROP TABLE IF EXISTS "public"."client";
CREATE TABLE "public"."client" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "host" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "region" varchar(128) COLLATE "pg_catalog"."default",
  "zone" varchar(128) COLLATE "pg_catalog"."default",
  "campus" varchar(128) COLLATE "pg_catalog"."default",
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."client" OWNER TO "postgres";
COMMENT ON COLUMN "public"."client"."id" IS 'client id';
COMMENT ON COLUMN "public"."client"."host" IS 'client host IP';
COMMENT ON COLUMN "public"."client"."type" IS 'client type: polaris-java/polaris-go';
COMMENT ON COLUMN "public"."client"."version" IS 'client SDK version';
COMMENT ON COLUMN "public"."client"."region" IS 'region info for client';
COMMENT ON COLUMN "public"."client"."zone" IS 'zone info for client';
COMMENT ON COLUMN "public"."client"."campus" IS 'campus info for client';
COMMENT ON COLUMN "public"."client"."flag" IS '0 is valid, 1 is invalid(deleted)';
COMMENT ON COLUMN "public"."client"."ctime" IS 'create time';
COMMENT ON COLUMN "public"."client"."mtime" IS 'last updated time';

-- ----------------------------
-- Table structure for client_stat
-- ----------------------------
DROP TABLE IF EXISTS "public"."client_stat";
CREATE TABLE "public"."client_stat" (
  "client_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "target" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "port" int4 NOT NULL,
  "protocol" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "path" varchar(128) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."client_stat" OWNER TO "postgres";
COMMENT ON COLUMN "public"."client_stat"."client_id" IS 'client id';
COMMENT ON COLUMN "public"."client_stat"."target" IS 'target stat platform';
COMMENT ON COLUMN "public"."client_stat"."port" IS 'client port to get stat information';
COMMENT ON COLUMN "public"."client_stat"."protocol" IS 'stat info transport protocol';
COMMENT ON COLUMN "public"."client_stat"."path" IS 'stat metric path';

-- ----------------------------
-- Table structure for config_file
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file";
CREATE TABLE "public"."config_file" (
  "id" int8 NOT NULL DEFAULT nextval('config_file_id_seq'::regclass),
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "flag" int2 NOT NULL DEFAULT 0,
  "create_time" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file" OWNER TO "postgres";

-- ----------------------------
-- Table structure for config_file_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_group";
CREATE TABLE "public"."config_file_group" (
  "id" int8 NOT NULL DEFAULT nextval('config_file_group_id_seq'::regclass),
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "business" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "department" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "metadata" text COLLATE "pg_catalog"."default",
  "flag" int2 NOT NULL DEFAULT 0
)
;
ALTER TABLE "public"."config_file_group" OWNER TO "postgres";

-- ----------------------------
-- Table structure for config_file_release
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_release";
CREATE TABLE "public"."config_file_release" (
  "id" int8 NOT NULL DEFAULT nextval('config_file_release_id_seq'::regclass),
  "name" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "version" int8 NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "tags" text COLLATE "pg_catalog"."default",
  "active" int2 NOT NULL DEFAULT 0,
  "description" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "release_type" varchar(25) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
ALTER TABLE "public"."config_file_release" OWNER TO "postgres";

-- ----------------------------
-- Table structure for config_file_release_history
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_release_history";
CREATE TABLE "public"."config_file_release_history" (
  "id" int8 NOT NULL DEFAULT nextval('config_file_release_history_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "status" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'success'::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "tags" text COLLATE "pg_catalog"."default",
  "version" int8,
  "reason" varchar(3000) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "description" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_release_history" OWNER TO "postgres";

-- ----------------------------
-- Table structure for config_file_tag
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_tag";
CREATE TABLE "public"."config_file_tag" (
  "id" int8 NOT NULL DEFAULT nextval('config_file_tag_id_seq'::regclass),
  "key" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "value" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_tag" OWNER TO "postgres";

-- ----------------------------
-- Table structure for config_file_template
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_template";
CREATE TABLE "public"."config_file_template" (
  "id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY (
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1
),
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_template" OWNER TO "postgres";
COMMENT ON COLUMN "public"."config_file_template"."id" IS '主键';
COMMENT ON COLUMN "public"."config_file_template"."name" IS '配置文件模板名称';
COMMENT ON COLUMN "public"."config_file_template"."content" IS '配置文件模板内容';
COMMENT ON COLUMN "public"."config_file_template"."format" IS '模板文件格式';
COMMENT ON COLUMN "public"."config_file_template"."comment" IS '模板描述信息';
COMMENT ON COLUMN "public"."config_file_template"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."config_file_template"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."config_file_template"."modify_time" IS '最后更新时间';
COMMENT ON COLUMN "public"."config_file_template"."modify_by" IS '最后更新人';

-- ----------------------------
-- Table structure for fault_detect_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."fault_detect_rule";
CREATE TABLE "public"."fault_detect_rule" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'default'::character varying,
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "dst_service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "dst_namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "dst_method" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "config" text COLLATE "pg_catalog"."default",
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."fault_detect_rule" OWNER TO "postgres";
COMMENT ON COLUMN "public"."fault_detect_rule"."metadata" IS 'faultdetect rule metadata';

-- ----------------------------
-- Table structure for gray_resource
-- ----------------------------
DROP TABLE IF EXISTS "public"."gray_resource";
CREATE TABLE "public"."gray_resource" (
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "match_rule" text COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "flag" int2 DEFAULT 0
)
;
ALTER TABLE "public"."gray_resource" OWNER TO "postgres";
COMMENT ON COLUMN "public"."gray_resource"."name" IS '灰度资源';
COMMENT ON COLUMN "public"."gray_resource"."match_rule" IS '配置规则';
COMMENT ON COLUMN "public"."gray_resource"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."gray_resource"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."gray_resource"."modify_time" IS '最后更新时间';
COMMENT ON COLUMN "public"."gray_resource"."modify_by" IS '最后更新人';
COMMENT ON COLUMN "public"."gray_resource"."flag" IS '逻辑删除标志位, 0 为有效, 1 为逻辑删除';
COMMENT ON TABLE "public"."gray_resource" IS '灰度资源表';

-- ----------------------------
-- Table structure for health_check
-- ----------------------------
DROP TABLE IF EXISTS "public"."health_check";
CREATE TABLE "public"."health_check" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL DEFAULT 0,
  "ttl" int4 NOT NULL
)
;
ALTER TABLE "public"."health_check" OWNER TO "postgres";
COMMENT ON COLUMN "public"."health_check"."id" IS 'Instance ID';
COMMENT ON COLUMN "public"."health_check"."type" IS 'Instance health check type';
COMMENT ON COLUMN "public"."health_check"."ttl" IS 'TTL time jumping';

-- ----------------------------
-- Table structure for instance
-- ----------------------------
DROP TABLE IF EXISTS "public"."instance";
CREATE TABLE "public"."instance" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "vpc_id" varchar(64) COLLATE "pg_catalog"."default",
  "host" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "port" int4 NOT NULL,
  "protocol" varchar(32) COLLATE "pg_catalog"."default",
  "version" varchar(32) COLLATE "pg_catalog"."default",
  "health_status" int2 NOT NULL DEFAULT 1,
  "isolate" int2 NOT NULL DEFAULT 0,
  "weight" int2 NOT NULL DEFAULT 100,
  "enable_health_check" int2 NOT NULL DEFAULT 0,
  "logic_set" varchar(128) COLLATE "pg_catalog"."default",
  "cmdb_region" varchar(128) COLLATE "pg_catalog"."default",
  "cmdb_zone" varchar(128) COLLATE "pg_catalog"."default",
  "cmdb_idc" varchar(128) COLLATE "pg_catalog"."default",
  "priority" int2 NOT NULL DEFAULT 0,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."instance" OWNER TO "postgres";
COMMENT ON COLUMN "public"."instance"."id" IS 'Unique ID';
COMMENT ON COLUMN "public"."instance"."service_id" IS 'Service ID';
COMMENT ON COLUMN "public"."instance"."vpc_id" IS 'VPC ID';
COMMENT ON COLUMN "public"."instance"."host" IS 'Instance Host Information';
COMMENT ON COLUMN "public"."instance"."port" IS 'Instance port information';
COMMENT ON COLUMN "public"."instance"."protocol" IS 'Listening protocols for corresponding ports';
COMMENT ON COLUMN "public"."instance"."version" IS 'The version of the instance';
COMMENT ON COLUMN "public"."instance"."health_status" IS 'The health status of the instance, 1 is health, 0 is unhealthy';
COMMENT ON COLUMN "public"."instance"."isolate" IS 'Example isolation status flag, 0 is not isolated, 1 is isolated';
COMMENT ON COLUMN "public"."instance"."weight" IS 'The weight of the instance is mainly used for LoadBalance, default is 100';
COMMENT ON COLUMN "public"."instance"."enable_health_check" IS 'Whether to open a heartbeat on an instance, check the logic, 0 is not open, 1 is open';
COMMENT ON COLUMN "public"."instance"."logic_set" IS 'Example logic packet information';
COMMENT ON COLUMN "public"."instance"."cmdb_region" IS 'The region information of the instance is mainly used to close the route';
COMMENT ON COLUMN "public"."instance"."cmdb_zone" IS 'The ZONE information of the instance is mainly used to close the route.';
COMMENT ON COLUMN "public"."instance"."cmdb_idc" IS 'The IDC information of the instance is mainly used to close the route';
COMMENT ON COLUMN "public"."instance"."priority" IS 'Example priority, currently useless';
COMMENT ON COLUMN "public"."instance"."revision" IS 'Instance version information';
COMMENT ON COLUMN "public"."instance"."flag" IS 'Logic delete flag, 0 means visible, 1 means that it has been logically deleted';
COMMENT ON COLUMN "public"."instance"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."instance"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for instance_metadata
-- ----------------------------
DROP TABLE IF EXISTS "public"."instance_metadata";
CREATE TABLE "public"."instance_metadata" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mkey" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mvalue" varchar(4096) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."instance_metadata" OWNER TO "postgres";
COMMENT ON COLUMN "public"."instance_metadata"."id" IS 'Instance ID';
COMMENT ON COLUMN "public"."instance_metadata"."mkey" IS 'Instance label of Key';
COMMENT ON COLUMN "public"."instance_metadata"."mvalue" IS 'Instance label Value';
COMMENT ON COLUMN "public"."instance_metadata"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."instance_metadata"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for lane_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."lane_group";
CREATE TABLE "public"."lane_group" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "rule" text COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(3000) COLLATE "pg_catalog"."default",
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."lane_group" OWNER TO "postgres";
COMMENT ON COLUMN "public"."lane_group"."id" IS '泳道分组 ID';
COMMENT ON COLUMN "public"."lane_group"."name" IS '泳道分组名称';
COMMENT ON COLUMN "public"."lane_group"."rule" IS '规则的 json 字符串';
COMMENT ON COLUMN "public"."lane_group"."description" IS '规则描述';
COMMENT ON COLUMN "public"."lane_group"."revision" IS '规则摘要';
COMMENT ON COLUMN "public"."lane_group"."flag" IS '软删除标识位';
COMMENT ON COLUMN "public"."lane_group"."ctime" IS '创建时间';
COMMENT ON COLUMN "public"."lane_group"."mtime" IS '最后修改时间';
COMMENT ON COLUMN "public"."lane_group"."metadata" IS 'lane rule metadata';

-- ----------------------------
-- Table structure for lane_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."lane_rule";
CREATE TABLE "public"."lane_rule" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "rule" text COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(3000) COLLATE "pg_catalog"."default",
  "enable" int2,
  "flag" int2 DEFAULT 0,
  "priority" int8 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "etime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."lane_rule" OWNER TO "postgres";
COMMENT ON COLUMN "public"."lane_rule"."id" IS '规则 id';
COMMENT ON COLUMN "public"."lane_rule"."name" IS '规则名称';
COMMENT ON COLUMN "public"."lane_rule"."group_name" IS '泳道分组名称';
COMMENT ON COLUMN "public"."lane_rule"."rule" IS '规则的 json 字符串';
COMMENT ON COLUMN "public"."lane_rule"."revision" IS '规则摘要';
COMMENT ON COLUMN "public"."lane_rule"."description" IS '规则描述';
COMMENT ON COLUMN "public"."lane_rule"."enable" IS '是否启用';
COMMENT ON COLUMN "public"."lane_rule"."flag" IS '软删除标识位';
COMMENT ON COLUMN "public"."lane_rule"."priority" IS '泳道规则优先级';
COMMENT ON COLUMN "public"."lane_rule"."ctime" IS '创建时间';
COMMENT ON COLUMN "public"."lane_rule"."etime" IS '结束时间';
COMMENT ON COLUMN "public"."lane_rule"."mtime" IS '最后修改时间';

-- ----------------------------
-- Table structure for leader_election
-- ----------------------------
DROP TABLE IF EXISTS "public"."leader_election";
CREATE TABLE "public"."leader_election" (
  "elect_key" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "version" int8 NOT NULL DEFAULT 0,
  "leader" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."leader_election" OWNER TO "postgres";
COMMENT ON COLUMN "public"."leader_election"."elect_key" IS '选举键';
COMMENT ON COLUMN "public"."leader_election"."version" IS '版本';
COMMENT ON COLUMN "public"."leader_election"."leader" IS '领导者';
COMMENT ON COLUMN "public"."leader_election"."ctime" IS '创建时间';
COMMENT ON COLUMN "public"."leader_election"."mtime" IS '最后修改时间';

-- ----------------------------
-- Table structure for namespace
-- ----------------------------
DROP TABLE IF EXISTS "public"."namespace";
CREATE TABLE "public"."namespace" (
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(1024) COLLATE "pg_catalog"."default",
  "token" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "service_export_to" text COLLATE "pg_catalog"."default",
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."namespace" OWNER TO "postgres";
COMMENT ON COLUMN "public"."namespace"."name" IS 'Namespace name, unique';
COMMENT ON COLUMN "public"."namespace"."comment" IS 'Description of namespace';
COMMENT ON COLUMN "public"."namespace"."token" IS 'TOKEN named space for write operation check';
COMMENT ON COLUMN "public"."namespace"."owner" IS 'Responsible for named space Owner';
COMMENT ON COLUMN "public"."namespace"."flag" IS 'Logic delete flag, 0 means visible, 1 means that it has been logically deleted';
COMMENT ON COLUMN "public"."namespace"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."namespace"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."namespace"."service_export_to" IS 'Namespace metadata';
COMMENT ON COLUMN "public"."namespace"."metadata" IS 'Namespace metadata';

-- ----------------------------
-- Table structure for owner_service_map
-- ----------------------------
DROP TABLE IF EXISTS "public"."owner_service_map";
CREATE TABLE "public"."owner_service_map" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."owner_service_map" OWNER TO "postgres";
COMMENT ON COLUMN "public"."owner_service_map"."id" IS 'Primary key ID';
COMMENT ON COLUMN "public"."owner_service_map"."owner" IS 'Service Owner';
COMMENT ON COLUMN "public"."owner_service_map"."service" IS 'Service name';
COMMENT ON COLUMN "public"."owner_service_map"."namespace" IS 'Namespace name';

-- ----------------------------
-- Table structure for ratelimit_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."ratelimit_config";
CREATE TABLE "public"."ratelimit_config" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "disable" int2 NOT NULL DEFAULT 0,
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "method" varchar(512) COLLATE "pg_catalog"."default" NOT NULL,
  "labels" text COLLATE "pg_catalog"."default" NOT NULL,
  "priority" int2 NOT NULL DEFAULT 0,
  "rule" text COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "etime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."ratelimit_config" OWNER TO "postgres";
COMMENT ON COLUMN "public"."ratelimit_config"."id" IS 'ratelimit rule ID';
COMMENT ON COLUMN "public"."ratelimit_config"."name" IS 'ratelimt rule name';
COMMENT ON COLUMN "public"."ratelimit_config"."disable" IS 'ratelimit disable';
COMMENT ON COLUMN "public"."ratelimit_config"."service_id" IS 'Service ID';
COMMENT ON COLUMN "public"."ratelimit_config"."method" IS 'ratelimit method';
COMMENT ON COLUMN "public"."ratelimit_config"."labels" IS 'Conductive flow for a specific label';
COMMENT ON COLUMN "public"."ratelimit_config"."priority" IS 'ratelimit rule priority';
COMMENT ON COLUMN "public"."ratelimit_config"."rule" IS 'Current limiting rules';
COMMENT ON COLUMN "public"."ratelimit_config"."revision" IS 'Limiting version';
COMMENT ON COLUMN "public"."ratelimit_config"."flag" IS 'Logic delete flag, 0 means visible, 1 means that it has been logically deleted';
COMMENT ON COLUMN "public"."ratelimit_config"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."ratelimit_config"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."ratelimit_config"."etime" IS 'RateLimit rule enable time';
COMMENT ON COLUMN "public"."ratelimit_config"."metadata" IS 'ratelimit rule metadata';

-- ----------------------------
-- Table structure for ratelimit_revision
-- ----------------------------
DROP TABLE IF EXISTS "public"."ratelimit_revision";
CREATE TABLE "public"."ratelimit_revision" (
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "last_revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."ratelimit_revision" OWNER TO "postgres";
COMMENT ON COLUMN "public"."ratelimit_revision"."service_id" IS 'Service ID';
COMMENT ON COLUMN "public"."ratelimit_revision"."last_revision" IS 'The latest limited limiting rule version of the corresponding service';
COMMENT ON COLUMN "public"."ratelimit_revision"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for routing_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."routing_config";
CREATE TABLE "public"."routing_config" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "in_bounds" text COLLATE "pg_catalog"."default",
  "out_bounds" text COLLATE "pg_catalog"."default",
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."routing_config" OWNER TO "postgres";
COMMENT ON COLUMN "public"."routing_config"."id" IS 'Routing configuration ID';
COMMENT ON COLUMN "public"."routing_config"."in_bounds" IS 'Service is routing rules';
COMMENT ON COLUMN "public"."routing_config"."out_bounds" IS 'Service main routing rules';
COMMENT ON COLUMN "public"."routing_config"."revision" IS 'Routing rule version';
COMMENT ON COLUMN "public"."routing_config"."flag" IS 'Logic delete flag, 0 means visible, 1 means that it has been logically deleted';
COMMENT ON COLUMN "public"."routing_config"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."routing_config"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for routing_config_v2
-- ----------------------------
DROP TABLE IF EXISTS "public"."routing_config_v2";
CREATE TABLE "public"."routing_config_v2" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "policy" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "config" text COLLATE "pg_catalog"."default",
  "enable" int4 NOT NULL DEFAULT 0,
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "priority" int2 NOT NULL DEFAULT 0,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "etime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "extend_info" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."routing_config_v2" OWNER TO "postgres";
COMMENT ON COLUMN "public"."routing_config_v2"."priority" IS 'ratelimit rule priority';
COMMENT ON COLUMN "public"."routing_config_v2"."metadata" IS 'route rule metadata';

-- ----------------------------
-- Table structure for service
-- ----------------------------
DROP TABLE IF EXISTS "public"."service";
CREATE TABLE "public"."service" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "ports" text COLLATE "pg_catalog"."default",
  "business" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "department" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_mod1" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_mod2" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_mod3" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "comment" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "token" varchar(2048) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "reference" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "refer_filter" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "platform_id" varchar(32) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "export_to" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."service" OWNER TO "postgres";
COMMENT ON COLUMN "public"."service"."id" IS 'Service ID';
COMMENT ON COLUMN "public"."service"."name" IS 'Service name, only under the namespace';
COMMENT ON COLUMN "public"."service"."namespace" IS 'Namespace belongs to the service';
COMMENT ON COLUMN "public"."service"."ports" IS 'Service will have a list of all port information of the external exposure (single process exposing multiple protocols)';
COMMENT ON COLUMN "public"."service"."business" IS 'Service business information';
COMMENT ON COLUMN "public"."service"."department" IS 'Service department information';
COMMENT ON COLUMN "public"."service"."cmdb_mod1" IS 'Custom module field 1';
COMMENT ON COLUMN "public"."service"."cmdb_mod2" IS 'Custom module field 2';
COMMENT ON COLUMN "public"."service"."cmdb_mod3" IS 'Custom module field 3';
COMMENT ON COLUMN "public"."service"."comment" IS 'Description information';
COMMENT ON COLUMN "public"."service"."token" IS 'Service token, used to handle all the services involved in the service';
COMMENT ON COLUMN "public"."service"."revision" IS 'Service version information';
COMMENT ON COLUMN "public"."service"."owner" IS 'Owner information belonging to the service';
COMMENT ON COLUMN "public"."service"."flag" IS 'Logic delete flag, 0 means visible, 1 means that it has been logically deleted';
COMMENT ON COLUMN "public"."service"."reference" IS 'Service alias, what is the actual service name that the service is actually pointed out?';
COMMENT ON COLUMN "public"."service"."refer_filter" IS 'Custom reference filter';
COMMENT ON COLUMN "public"."service"."platform_id" IS 'The platform ID to which the service belongs';
COMMENT ON COLUMN "public"."service"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."service"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."service"."export_to" IS 'Service export to some namespace';

-- ----------------------------
-- Table structure for service_contract
-- ----------------------------
DROP TABLE IF EXISTS "public"."service_contract";
CREATE TABLE "public"."service_contract" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "protocol" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 DEFAULT 0,
  "content" text COLLATE "pg_catalog"."default",
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."service_contract" OWNER TO "postgres";
COMMENT ON COLUMN "public"."service_contract"."id" IS '服务契约主键';
COMMENT ON COLUMN "public"."service_contract"."type" IS '服务契约名称';
COMMENT ON COLUMN "public"."service_contract"."namespace" IS '命名空间';
COMMENT ON COLUMN "public"."service_contract"."service" IS '服务名称';
COMMENT ON COLUMN "public"."service_contract"."protocol" IS '当前契约对应的协议信息 e.g. http/dubbo/grpc/thrift';
COMMENT ON COLUMN "public"."service_contract"."version" IS '服务契约版本';
COMMENT ON COLUMN "public"."service_contract"."revision" IS '当前服务契约的全部内容版本摘要';
COMMENT ON COLUMN "public"."service_contract"."flag" IS '逻辑删除标志位，0 为有效，1 为逻辑删除';
COMMENT ON COLUMN "public"."service_contract"."content" IS '描述信息';
COMMENT ON COLUMN "public"."service_contract"."ctime" IS '创建时间';
COMMENT ON COLUMN "public"."service_contract"."mtime" IS '最后修改时间';

-- ----------------------------
-- Table structure for service_contract_detail
-- ----------------------------
DROP TABLE IF EXISTS "public"."service_contract_detail";
CREATE TABLE "public"."service_contract_detail" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "contract_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "protocol" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "method" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "path" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "source" int4,
  "content" text COLLATE "pg_catalog"."default",
  "revision" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."service_contract_detail" OWNER TO "postgres";
COMMENT ON COLUMN "public"."service_contract_detail"."id" IS '服务契约单个接口定义记录主键';
COMMENT ON COLUMN "public"."service_contract_detail"."contract_id" IS '服务契约 ID';
COMMENT ON COLUMN "public"."service_contract_detail"."type" IS '服务契约接口名称';
COMMENT ON COLUMN "public"."service_contract_detail"."namespace" IS '命名空间';
COMMENT ON COLUMN "public"."service_contract_detail"."service" IS '服务名称';
COMMENT ON COLUMN "public"."service_contract_detail"."protocol" IS '当前契约对应的协议信息 e.g. http/dubbo/grpc/thrift';
COMMENT ON COLUMN "public"."service_contract_detail"."version" IS '服务契约版本';
COMMENT ON COLUMN "public"."service_contract_detail"."method" IS 'http协议中的 method 字段, eg: POST/GET/PUT/DELETE';
COMMENT ON COLUMN "public"."service_contract_detail"."path" IS '接口具体全路径描述';
COMMENT ON COLUMN "public"."service_contract_detail"."source" IS '该条记录来源, 0: SDK/1: MANUAL';
COMMENT ON COLUMN "public"."service_contract_detail"."content" IS '描述信息';
COMMENT ON COLUMN "public"."service_contract_detail"."revision" IS '当前接口定义的全部内容版本摘要';
COMMENT ON COLUMN "public"."service_contract_detail"."flag" IS '逻辑删除标志位, 0 为有效, 1 为逻辑删除';
COMMENT ON COLUMN "public"."service_contract_detail"."ctime" IS '创建时间';
COMMENT ON COLUMN "public"."service_contract_detail"."mtime" IS '最后修改时间';

-- ----------------------------
-- Table structure for service_metadata
-- ----------------------------
DROP TABLE IF EXISTS "public"."service_metadata";
CREATE TABLE "public"."service_metadata" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mkey" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mvalue" varchar(4096) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."service_metadata" OWNER TO "postgres";
COMMENT ON COLUMN "public"."service_metadata"."id" IS 'Service ID';
COMMENT ON COLUMN "public"."service_metadata"."mkey" IS 'Service label key';
COMMENT ON COLUMN "public"."service_metadata"."mvalue" IS 'Service label Value';
COMMENT ON COLUMN "public"."service_metadata"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."service_metadata"."mtime" IS 'Last updated time';

-- ----------------------------
-- Table structure for start_lock
-- ----------------------------
DROP TABLE IF EXISTS "public"."start_lock";
CREATE TABLE "public"."start_lock" (
  "lock_id" int4 NOT NULL,
  "lock_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "server" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."start_lock" OWNER TO "postgres";
COMMENT ON COLUMN "public"."start_lock"."lock_id" IS 'Lock ID';
COMMENT ON COLUMN "public"."start_lock"."lock_key" IS 'Lock name';
COMMENT ON COLUMN "public"."start_lock"."server" IS 'Server holding launch lock';
COMMENT ON COLUMN "public"."start_lock"."mtime" IS 'Update time';

-- ----------------------------
-- Table structure for t_ip_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_ip_config";
CREATE TABLE "public"."t_ip_config" (
  "fip" int4 NOT NULL,
  "fareaid" int4 NOT NULL,
  "fcityid" int4 NOT NULL,
  "fidcid" int4 NOT NULL,
  "fflag" int2 DEFAULT 0,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_ip_config" OWNER TO "postgres";
COMMENT ON COLUMN "public"."t_ip_config"."fip" IS 'Machine IP';
COMMENT ON COLUMN "public"."t_ip_config"."fareaid" IS 'Area number';
COMMENT ON COLUMN "public"."t_ip_config"."fcityid" IS 'City number';
COMMENT ON COLUMN "public"."t_ip_config"."fidcid" IS 'IDC number';
COMMENT ON COLUMN "public"."t_ip_config"."fflag" IS 'Flag';
COMMENT ON COLUMN "public"."t_ip_config"."fstamp" IS 'Timestamp';
COMMENT ON COLUMN "public"."t_ip_config"."fflow" IS 'Flow';

-- ----------------------------
-- Table structure for t_policy
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_policy";
CREATE TABLE "public"."t_policy" (
  "fmodid" int4 NOT NULL,
  "fdiv" int4 NOT NULL,
  "fmod" int4 NOT NULL,
  "fflag" int2 DEFAULT 0,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_policy" OWNER TO "postgres";
COMMENT ON COLUMN "public"."t_policy"."fmodid" IS 'Module ID';
COMMENT ON COLUMN "public"."t_policy"."fdiv" IS 'Division';
COMMENT ON COLUMN "public"."t_policy"."fmod" IS 'Module';
COMMENT ON COLUMN "public"."t_policy"."fflag" IS 'Flag';
COMMENT ON COLUMN "public"."t_policy"."fstamp" IS 'Timestamp';
COMMENT ON COLUMN "public"."t_policy"."fflow" IS 'Flow';

-- ----------------------------
-- Table structure for t_route
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_route";
CREATE TABLE "public"."t_route" (
  "fip" int4 NOT NULL,
  "fmodid" int4 NOT NULL,
  "fcmdid" int4 NOT NULL,
  "fsetid" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "fflag" int2 DEFAULT 0,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_route" OWNER TO "postgres";
COMMENT ON COLUMN "public"."t_route"."fip" IS 'IP';
COMMENT ON COLUMN "public"."t_route"."fmodid" IS 'Module ID';
COMMENT ON COLUMN "public"."t_route"."fcmdid" IS 'Command ID';
COMMENT ON COLUMN "public"."t_route"."fsetid" IS 'Set ID';
COMMENT ON COLUMN "public"."t_route"."fflag" IS 'Flag';
COMMENT ON COLUMN "public"."t_route"."fstamp" IS 'Timestamp';
COMMENT ON COLUMN "public"."t_route"."fflow" IS 'Flow';

-- ----------------------------
-- Table structure for t_section
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_section";
CREATE TABLE "public"."t_section" (
  "fmodid" int4 NOT NULL,
  "ffrom" int4 NOT NULL,
  "fto" int4 NOT NULL,
  "fxid" int4 NOT NULL,
  "fflag" int2 DEFAULT 0,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_section" OWNER TO "postgres";
COMMENT ON COLUMN "public"."t_section"."fmodid" IS 'Module ID';
COMMENT ON COLUMN "public"."t_section"."ffrom" IS 'From';
COMMENT ON COLUMN "public"."t_section"."fto" IS 'To';
COMMENT ON COLUMN "public"."t_section"."fxid" IS 'XID';
COMMENT ON COLUMN "public"."t_section"."fflag" IS 'Flag';
COMMENT ON COLUMN "public"."t_section"."fstamp" IS 'Timestamp';
COMMENT ON COLUMN "public"."t_section"."fflow" IS 'Flow';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "source" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mobile" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "email" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "token" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "token_enable" int2 NOT NULL DEFAULT 1,
  "user_type" int4 NOT NULL DEFAULT 20,
  "comment" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group";
CREATE TABLE "public"."user_group" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "token" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "token_enable" int2 NOT NULL DEFAULT 1,
  "flag" int2 NOT NULL DEFAULT 0,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "metadata" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."user_group" OWNER TO "postgres";
COMMENT ON COLUMN "public"."user_group"."id" IS 'User group ID';
COMMENT ON COLUMN "public"."user_group"."name" IS 'User group name';
COMMENT ON COLUMN "public"."user_group"."owner" IS 'The main account ID of the user group';
COMMENT ON COLUMN "public"."user_group"."token" IS 'TOKEN information of this user group';
COMMENT ON COLUMN "public"."user_group"."comment" IS 'Description';
COMMENT ON COLUMN "public"."user_group"."token_enable" IS 'Token enable';
COMMENT ON COLUMN "public"."user_group"."flag" IS 'Whether the rules are valid';
COMMENT ON COLUMN "public"."user_group"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."user_group"."mtime" IS 'Last updated time';
COMMENT ON COLUMN "public"."user_group"."metadata" IS 'User group metadata';
COMMENT ON TABLE "public"."user_group" IS 'User group table';

-- ----------------------------
-- Table structure for user_group_relation
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group_relation";
CREATE TABLE "public"."user_group_relation" (
  "user_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "group_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "mtime" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."user_group_relation" OWNER TO "postgres";
COMMENT ON COLUMN "public"."user_group_relation"."user_id" IS 'User ID';
COMMENT ON COLUMN "public"."user_group_relation"."group_id" IS 'User group ID';
COMMENT ON COLUMN "public"."user_group_relation"."ctime" IS 'Create time';
COMMENT ON COLUMN "public"."user_group_relation"."mtime" IS 'Last updated time';
COMMENT ON TABLE "public"."user_group_relation" IS 'User group relation table';

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_group_id_seq"
OWNED BY "public"."config_file_group"."id";
SELECT setval('"public"."config_file_group_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_id_seq"
OWNED BY "public"."config_file"."id";
SELECT setval('"public"."config_file_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_release_history_id_seq"
OWNED BY "public"."config_file_release_history"."id";
SELECT setval('"public"."config_file_release_history_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_release_id_seq"
OWNED BY "public"."config_file_release"."id";
SELECT setval('"public"."config_file_release_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_tag_id_seq"
OWNED BY "public"."config_file_tag"."id";
SELECT setval('"public"."config_file_tag_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."config_file_template_id_seq"
OWNED BY "public"."config_file_template"."id";
SELECT setval('"public"."config_file_template_id_seq"', 1, true);

-- ----------------------------
-- Primary Key structure for table auth_principal
-- ----------------------------
ALTER TABLE "public"."auth_principal" ADD CONSTRAINT "auth_principal_pkey" PRIMARY KEY ("strategy_id", "principal_id", "principal_role");

-- ----------------------------
-- Indexes structure for table auth_role
-- ----------------------------
CREATE INDEX "idx_ar_mtime" ON "public"."auth_role" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "idx_ar_owner" ON "public"."auth_role" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table auth_role
-- ----------------------------
ALTER TABLE "public"."auth_role" ADD CONSTRAINT "auth_role_name_owner_key" UNIQUE ("name", "owner");

-- ----------------------------
-- Primary Key structure for table auth_role
-- ----------------------------
ALTER TABLE "public"."auth_role" ADD CONSTRAINT "auth_role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table auth_role_principal
-- ----------------------------
ALTER TABLE "public"."auth_role_principal" ADD CONSTRAINT "auth_role_principal_pkey" PRIMARY KEY ("role_id", "principal_id", "principal_role");

-- ----------------------------
-- Indexes structure for table auth_strategy
-- ----------------------------
CREATE INDEX "idx_a_owner" ON "public"."auth_strategy" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_mt" ON "public"."auth_strategy" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table auth_strategy
-- ----------------------------
ALTER TABLE "public"."auth_strategy" ADD CONSTRAINT "auth_strategy_name_owner_key" UNIQUE ("name", "owner");

-- ----------------------------
-- Primary Key structure for table auth_strategy
-- ----------------------------
ALTER TABLE "public"."auth_strategy" ADD CONSTRAINT "auth_strategy_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table auth_strategy_function
-- ----------------------------
ALTER TABLE "public"."auth_strategy_function" ADD CONSTRAINT "auth_strategy_function_pkey" PRIMARY KEY ("strategy_id", "function");

-- ----------------------------
-- Primary Key structure for table auth_strategy_label
-- ----------------------------
ALTER TABLE "public"."auth_strategy_label" ADD CONSTRAINT "auth_strategy_label_pkey" PRIMARY KEY ("strategy_id", "key");

-- ----------------------------
-- Indexes structure for table auth_strategy_resource
-- ----------------------------
CREATE INDEX "idx_asr_mtime" ON "public"."auth_strategy_resource" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_strategy_resource
-- ----------------------------
ALTER TABLE "public"."auth_strategy_resource" ADD CONSTRAINT "auth_strategy_resource_pkey" PRIMARY KEY ("strategy_id", "res_type", "res_id");

-- ----------------------------
-- Uniques structure for table circuitbreaker_rule
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule" ADD CONSTRAINT "circuitbreaker_rule_name_namespace_version_key" UNIQUE ("name", "namespace", "version");

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule" ADD CONSTRAINT "circuitbreaker_rule_pkey" PRIMARY KEY ("id", "version");

-- ----------------------------
-- Indexes structure for table circuitbreaker_rule_relation
-- ----------------------------
CREATE INDEX "idx_rule_id" ON "public"."circuitbreaker_rule_relation" USING btree (
  "rule_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule_relation
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule_relation" ADD CONSTRAINT "circuitbreaker_rule_relation_pkey" PRIMARY KEY ("service_id");

-- ----------------------------
-- Indexes structure for table circuitbreaker_rule_v2
-- ----------------------------
CREATE INDEX "idx_cr_mtime" ON "public"."circuitbreaker_rule_v2" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "idx_name" ON "public"."circuitbreaker_rule_v2" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule_v2
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule_v2" ADD CONSTRAINT "circuitbreaker_rule_v2_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table cl5_module
-- ----------------------------
ALTER TABLE "public"."cl5_module" ADD CONSTRAINT "cl5_module_pkey" PRIMARY KEY ("module_id");

-- ----------------------------
-- Indexes structure for table client
-- ----------------------------
CREATE INDEX "idx_c_mtime" ON "public"."client" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table client
-- ----------------------------
ALTER TABLE "public"."client" ADD CONSTRAINT "client_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table client_stat
-- ----------------------------
ALTER TABLE "public"."client_stat" ADD CONSTRAINT "client_stat_pkey" PRIMARY KEY ("client_id", "target", "port");

-- ----------------------------
-- Uniques structure for table config_file
-- ----------------------------
ALTER TABLE "public"."config_file" ADD CONSTRAINT "config_file_namespace_group_name_key" UNIQUE ("namespace", "group", "name");

-- ----------------------------
-- Primary Key structure for table config_file
-- ----------------------------
ALTER TABLE "public"."config_file" ADD CONSTRAINT "config_file_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table config_file_group
-- ----------------------------
ALTER TABLE "public"."config_file_group" ADD CONSTRAINT "config_file_group_namespace_name_key" UNIQUE ("namespace", "name");

-- ----------------------------
-- Primary Key structure for table config_file_group
-- ----------------------------
ALTER TABLE "public"."config_file_group" ADD CONSTRAINT "config_file_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table config_file_release
-- ----------------------------
ALTER TABLE "public"."config_file_release" ADD CONSTRAINT "config_file_release_namespace_group_file_name_name_key" UNIQUE ("namespace", "group", "file_name", "name");

-- ----------------------------
-- Primary Key structure for table config_file_release
-- ----------------------------
ALTER TABLE "public"."config_file_release" ADD CONSTRAINT "config_file_release_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_release_history
-- ----------------------------
ALTER TABLE "public"."config_file_release_history" ADD CONSTRAINT "config_file_release_history_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table config_file_tag
-- ----------------------------
ALTER TABLE "public"."config_file_tag" ADD CONSTRAINT "config_file_tag_key_value_namespace_group_file_name_key" UNIQUE ("key", "value", "namespace", "group", "file_name");

-- ----------------------------
-- Primary Key structure for table config_file_tag
-- ----------------------------
ALTER TABLE "public"."config_file_tag" ADD CONSTRAINT "config_file_tag_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table config_file_template
-- ----------------------------
ALTER TABLE "public"."config_file_template" ADD CONSTRAINT "uk_name" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table config_file_template
-- ----------------------------
ALTER TABLE "public"."config_file_template" ADD CONSTRAINT "config_file_template_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table fault_detect_rule
-- ----------------------------
CREATE INDEX "idx_fdr_mtime" ON "public"."fault_detect_rule" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "idx_fdr_name" ON "public"."fault_detect_rule" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table fault_detect_rule
-- ----------------------------
ALTER TABLE "public"."fault_detect_rule" ADD CONSTRAINT "fault_detect_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table gray_resource
-- ----------------------------
ALTER TABLE "public"."gray_resource" ADD CONSTRAINT "gray_resource_pkey" PRIMARY KEY ("name");

-- ----------------------------
-- Primary Key structure for table health_check
-- ----------------------------
ALTER TABLE "public"."health_check" ADD CONSTRAINT "health_check_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table instance
-- ----------------------------
CREATE INDEX "idx_host" ON "public"."instance" USING btree (
  "host" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_mtime" ON "public"."instance" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "idx_service_id" ON "public"."instance" USING btree (
  "service_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table instance
-- ----------------------------
ALTER TABLE "public"."instance" ADD CONSTRAINT "instance_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table instance_metadata
-- ----------------------------
CREATE INDEX "idx_mkey" ON "public"."instance_metadata" USING btree (
  "mkey" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table instance_metadata
-- ----------------------------
ALTER TABLE "public"."instance_metadata" ADD CONSTRAINT "instance_metadata_pkey" PRIMARY KEY ("id", "mkey");

-- ----------------------------
-- Indexes structure for table lane_group
-- ----------------------------
CREATE UNIQUE INDEX "idx_lane_group_name" ON "public"."lane_group" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table lane_group
-- ----------------------------
ALTER TABLE "public"."lane_group" ADD CONSTRAINT "lane_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table lane_rule
-- ----------------------------
CREATE UNIQUE INDEX "idx_lane_rule_unique" ON "public"."lane_rule" USING btree (
  "group_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table lane_rule
-- ----------------------------
ALTER TABLE "public"."lane_rule" ADD CONSTRAINT "lane_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table leader_election
-- ----------------------------
CREATE INDEX "idx_version" ON "public"."leader_election" USING btree (
  "version" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table leader_election
-- ----------------------------
ALTER TABLE "public"."leader_election" ADD CONSTRAINT "leader_election_pkey" PRIMARY KEY ("elect_key");

-- ----------------------------
-- Primary Key structure for table namespace
-- ----------------------------
ALTER TABLE "public"."namespace" ADD CONSTRAINT "namespace_pkey" PRIMARY KEY ("name");

-- ----------------------------
-- Indexes structure for table owner_service_map
-- ----------------------------
CREATE INDEX "idx_owner" ON "public"."owner_service_map" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_service_namespace" ON "public"."owner_service_map" USING btree (
  "service" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "namespace" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table owner_service_map
-- ----------------------------
ALTER TABLE "public"."owner_service_map" ADD CONSTRAINT "owner_service_map_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table ratelimit_config
-- ----------------------------
ALTER TABLE "public"."ratelimit_config" ADD CONSTRAINT "ratelimit_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table ratelimit_revision
-- ----------------------------
ALTER TABLE "public"."ratelimit_revision" ADD CONSTRAINT "ratelimit_revision_pkey" PRIMARY KEY ("service_id");

-- ----------------------------
-- Primary Key structure for table routing_config
-- ----------------------------
ALTER TABLE "public"."routing_config" ADD CONSTRAINT "routing_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table routing_config_v2
-- ----------------------------
CREATE INDEX "idx_rc_mtime" ON "public"."routing_config_v2" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table routing_config_v2
-- ----------------------------
ALTER TABLE "public"."routing_config_v2" ADD CONSTRAINT "routing_config_v2_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table service
-- ----------------------------
CREATE INDEX "idx_namespace" ON "public"."service" USING btree (
  "namespace" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_platform_id" ON "public"."service" USING btree (
  "platform_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_reference" ON "public"."service" USING btree (
  "reference" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table service
-- ----------------------------
ALTER TABLE "public"."service" ADD CONSTRAINT "service_name_namespace_key" UNIQUE ("name", "namespace");

-- ----------------------------
-- Primary Key structure for table service
-- ----------------------------
ALTER TABLE "public"."service" ADD CONSTRAINT "service_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table service_contract
-- ----------------------------
CREATE UNIQUE INDEX "idx_service_contract_unique" ON "public"."service_contract" USING btree (
  "namespace" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "service" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "version" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "protocol" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table service_contract
-- ----------------------------
ALTER TABLE "public"."service_contract" ADD CONSTRAINT "service_contract_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table service_contract_detail
-- ----------------------------
CREATE UNIQUE INDEX "idx_service_contract_detail_unique" ON "public"."service_contract_detail" USING btree (
  "contract_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "path" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "method" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "source" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table service_contract_detail
-- ----------------------------
ALTER TABLE "public"."service_contract_detail" ADD CONSTRAINT "service_contract_detail_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table service_metadata
-- ----------------------------
ALTER TABLE "public"."service_metadata" ADD CONSTRAINT "service_metadata_pkey" PRIMARY KEY ("id", "mkey");

-- ----------------------------
-- Primary Key structure for table start_lock
-- ----------------------------
ALTER TABLE "public"."start_lock" ADD CONSTRAINT "start_lock_pkey" PRIMARY KEY ("lock_id", "lock_key");

-- ----------------------------
-- Indexes structure for table t_ip_config
-- ----------------------------
CREATE INDEX "idx_fflow" ON "public"."t_ip_config" USING btree (
  "fflow" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table t_ip_config
-- ----------------------------
ALTER TABLE "public"."t_ip_config" ADD CONSTRAINT "t_ip_config_pkey" PRIMARY KEY ("fip");

-- ----------------------------
-- Primary Key structure for table t_policy
-- ----------------------------
ALTER TABLE "public"."t_policy" ADD CONSTRAINT "t_policy_pkey" PRIMARY KEY ("fmodid");

-- ----------------------------
-- Indexes structure for table t_route
-- ----------------------------
CREATE INDEX "idx1" ON "public"."t_route" USING btree (
  "fmodid" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "fcmdid" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "fsetid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table t_route
-- ----------------------------
ALTER TABLE "public"."t_route" ADD CONSTRAINT "t_route_pkey" PRIMARY KEY ("fip", "fmodid", "fcmdid");

-- ----------------------------
-- Primary Key structure for table t_section
-- ----------------------------
ALTER TABLE "public"."t_section" ADD CONSTRAINT "t_section_pkey" PRIMARY KEY ("fmodid", "ffrom", "fto");

-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_name_owner_key" UNIQUE ("name", "owner");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_group
-- ----------------------------
CREATE INDEX "mtime_idx" ON "public"."user_group" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "owner_idx" ON "public"."user_group" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table user_group
-- ----------------------------
ALTER TABLE "public"."user_group" ADD CONSTRAINT "user_group_name_owner_key" UNIQUE ("name", "owner");

-- ----------------------------
-- Primary Key structure for table user_group
-- ----------------------------
ALTER TABLE "public"."user_group" ADD CONSTRAINT "user_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_group_relation
-- ----------------------------
CREATE INDEX "idx_time" ON "public"."user_group_relation" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_group_relation
-- ----------------------------
ALTER TABLE "public"."user_group_relation" ADD CONSTRAINT "user_group_relation_pkey" PRIMARY KEY ("user_id", "group_id");
