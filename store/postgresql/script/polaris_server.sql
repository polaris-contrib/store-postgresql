/*
 Navicat Premium Data Transfer

 Source Server         : dev-postgresql
 Source Server Type    : PostgreSQL
 Source Server Version : 90224 (90224)
 Source Host           : 192.168.31.19:5432
 Source Catalog        : polaris_server
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90224 (90224)
 File Encoding         : 65001

 Date: 12/06/2023 23:23:58
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

-- ----------------------------
-- Records of auth_principal
-- ----------------------------
BEGIN;
COMMIT;

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
  "default" int2 NOT NULL DEFAULT 0::smallint,
  "revision" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."auth_strategy" OWNER TO "postgres";

-- ----------------------------
-- Records of auth_strategy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for auth_strategy_resource
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_strategy_resource";
CREATE TABLE "public"."auth_strategy_resource" (
  "strategy_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "res_type" int4 NOT NULL,
  "res_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."auth_strategy_resource" OWNER TO "postgres";

-- ----------------------------
-- Records of auth_strategy_resource
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for business
-- ----------------------------
DROP TABLE IF EXISTS "public"."business";
CREATE TABLE "public"."business" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "token" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."business" OWNER TO "postgres";

-- ----------------------------
-- Records of business
-- ----------------------------
BEGIN;
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."circuitbreaker_rule" OWNER TO "postgres";

-- ----------------------------
-- Records of circuitbreaker_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for circuitbreaker_rule_relation
-- ----------------------------
DROP TABLE IF EXISTS "public"."circuitbreaker_rule_relation";
CREATE TABLE "public"."circuitbreaker_rule_relation" (
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "rule_id" varchar(97) COLLATE "pg_catalog"."default" NOT NULL,
  "rule_version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."circuitbreaker_rule_relation" OWNER TO "postgres";

-- ----------------------------
-- Records of circuitbreaker_rule_relation
-- ----------------------------
BEGIN;
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now(),
  "etime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."circuitbreaker_rule_v2" OWNER TO "postgres";

-- ----------------------------
-- Records of circuitbreaker_rule_v2
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for cl5_module
-- ----------------------------
DROP TABLE IF EXISTS "public"."cl5_module";
CREATE TABLE "public"."cl5_module" (
  "module_id" int4 NOT NULL,
  "interface_id" int4 NOT NULL,
  "range_num" int4 NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."cl5_module" OWNER TO "postgres";

-- ----------------------------
-- Records of cl5_module
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for client
-- ----------------------------
DROP TABLE IF EXISTS "public"."client";
CREATE TABLE "public"."client" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "host" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "region" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "zone" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "campus" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."client" OWNER TO "postgres";

-- ----------------------------
-- Records of client
-- ----------------------------
BEGIN;
INSERT INTO "public"."client" ("id", "host", "type", "version", "region", "zone", "campus", "flag", "ctime", "mtime") VALUES ('client-0', 'client-0', 'UNKNOWN', 'client-0', 'client-0-region', 'client-0-zone', 'client-0-campus', 0, '2023-06-12 13:56:21', '2023-06-12 13:56:21');
INSERT INTO "public"."client" ("id", "host", "type", "version", "region", "zone", "campus", "flag", "ctime", "mtime") VALUES ('client-1', 'client-1', 'UNKNOWN', 'client-1', 'client-1-region', 'client-1-zone', 'client-1-campus', 0, '2023-06-12 13:56:21', '2023-06-12 13:56:21');
INSERT INTO "public"."client" ("id", "host", "type", "version", "region", "zone", "campus", "flag", "ctime", "mtime") VALUES ('client-2', 'client-2', 'UNKNOWN', 'client-2', 'client-2-region', 'client-2-zone', 'client-2-campus', 0, '2023-06-12 13:56:21', '2023-06-12 13:56:21');
INSERT INTO "public"."client" ("id", "host", "type", "version", "region", "zone", "campus", "flag", "ctime", "mtime") VALUES ('client-3', 'client-3', 'UNKNOWN', 'client-3', 'client-3-region', 'client-3-zone', 'client-3-campus', 0, '2023-06-12 13:56:21', '2023-06-12 13:56:21');
INSERT INTO "public"."client" ("id", "host", "type", "version", "region", "zone", "campus", "flag", "ctime", "mtime") VALUES ('client-4', 'client-4', 'UNKNOWN', 'client-4', 'client-4-region', 'client-4-zone', 'client-4-campus', 0, '2023-06-12 13:56:21', '2023-06-12 13:56:21');
COMMIT;

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

-- ----------------------------
-- Records of client_stat
-- ----------------------------
BEGIN;
INSERT INTO "public"."client_stat" ("client_id", "target", "port", "protocol", "path") VALUES ('client-0', 'prometheus', 8080, 'http', '/metrics');
INSERT INTO "public"."client_stat" ("client_id", "target", "port", "protocol", "path") VALUES ('client-1', 'prometheus', 8080, 'http', '/metrics');
INSERT INTO "public"."client_stat" ("client_id", "target", "port", "protocol", "path") VALUES ('client-2', 'prometheus', 8080, 'http', '/metrics');
INSERT INTO "public"."client_stat" ("client_id", "target", "port", "protocol", "path") VALUES ('client-3', 'prometheus', 8080, 'http', '/metrics');
INSERT INTO "public"."client_stat" ("client_id", "target", "port", "protocol", "path") VALUES ('client-4', 'prometheus', 8080, 'http', '/metrics');
COMMIT;

-- ----------------------------
-- Table structure for config_file
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file";
CREATE TABLE "public"."config_file" (
  "id" int8 NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for config_file_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_group";
CREATE TABLE "public"."config_file_group" (
  "id" int8 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_group" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for config_file_release
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_release";
CREATE TABLE "public"."config_file_release" (
  "id" int8 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "version" int4 NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_release" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file_release
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for config_file_release_history
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_release_history";
CREATE TABLE "public"."config_file_release_history" (
  "id" int8 NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "tags" varchar(2048) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "status" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'success'::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_release_history" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file_release_history
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for config_file_tag
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_tag";
CREATE TABLE "public"."config_file_tag" (
  "id" int8 NOT NULL,
  "key" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "value" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "file_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_tag" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file_tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for config_file_template
-- ----------------------------
DROP TABLE IF EXISTS "public"."config_file_template";
CREATE TABLE "public"."config_file_template" (
  "id" int8 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "format" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'text'::character varying,
  "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "modify_time" timestamp(6) NOT NULL DEFAULT now(),
  "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "public"."config_file_template" OWNER TO "postgres";

-- ----------------------------
-- Records of config_file_template
-- ----------------------------
BEGIN;
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."fault_detect_rule" OWNER TO "postgres";

-- ----------------------------
-- Records of fault_detect_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for health_check
-- ----------------------------
DROP TABLE IF EXISTS "public"."health_check";
CREATE TABLE "public"."health_check" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL DEFAULT 0::smallint,
  "ttl" int4 NOT NULL
)
;
ALTER TABLE "public"."health_check" OWNER TO "postgres";

-- ----------------------------
-- Records of health_check
-- ----------------------------
BEGIN;
INSERT INTO "public"."health_check" ("id", "type", "ttl") VALUES ('1112', 1, 2);
INSERT INTO "public"."health_check" ("id", "type", "ttl") VALUES ('1111e', 1, 1);
INSERT INTO "public"."health_check" ("id", "type", "ttl") VALUES ('1111u', 1, 1);
INSERT INTO "public"."health_check" ("id", "type", "ttl") VALUES ('1111i', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for instance
-- ----------------------------
DROP TABLE IF EXISTS "public"."instance";
CREATE TABLE "public"."instance" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "vpc_id" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "host" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "port" int4 NOT NULL,
  "protocol" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "version" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "health_status" int2 NOT NULL DEFAULT 1::smallint,
  "isolate" int2 NOT NULL DEFAULT 0::smallint,
  "weight" int2 NOT NULL DEFAULT 100::smallint,
  "enable_health_check" int2 NOT NULL DEFAULT 0::smallint,
  "logic_set" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_region" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_zone" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "cmdb_idc" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "priority" int2 NOT NULL DEFAULT 0::smallint,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."instance" OWNER TO "postgres";

-- ----------------------------
-- Records of instance
-- ----------------------------
BEGIN;
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111c', '111c', '4444c', '5555c', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 17:15:03', '2023-06-10 17:15:03');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111q', '111q', '4444q', '5555q', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 17:17:22', '2023-06-10 17:17:22');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111e', '111e', '4444e', '5555e', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 18:12:08', '2023-06-10 18:12:08');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111y', '111y', '4444y', '5555y', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 18:48:17', '2023-06-10 18:48:17');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111u', '111u', '4444u', '5555u', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 18:50:30', '2023-06-10 18:50:30');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111i', '111i', '4444i', '5555i', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 0, '2023-06-10 18:56:28', '2023-06-10 18:56:28');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111t', '111t', '4444t', '5555t', 1001, '', '', 1, 0, 0, 0, '', '', '', '', 0, '', 1, '2023-06-10 18:16:21', '2023-06-10 19:15:20');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111', '111', '4444', '55555', 1, '', '', 1, 0, 0, 0, '', '', '', '', 0, 'reversion', 0, '2023-06-10 17:05:31', '2023-06-11 01:09:55');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111a', '111a', '4444a', '55555a', 1, '', '', 1, 1, 0, 0, '', '', '', '', 0, 'reversion', 0, '2023-06-10 17:10:31', '2023-06-11 01:24:17');
INSERT INTO "public"."instance" ("id", "service_id", "vpc_id", "host", "port", "protocol", "version", "health_status", "isolate", "weight", "enable_health_check", "logic_set", "cmdb_region", "cmdb_zone", "cmdb_idc", "priority", "revision", "flag", "ctime", "mtime") VALUES ('1111b', '111a', '4444b', '55555b', 1001, '', '', 1, 1, 0, 0, '', '', '', '', 0, 'reversion', 0, '2023-06-10 17:12:04', '2023-06-11 01:24:17');
COMMIT;

-- ----------------------------
-- Table structure for instance_metadata
-- ----------------------------
DROP TABLE IF EXISTS "public"."instance_metadata";
CREATE TABLE "public"."instance_metadata" (
  "id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mkey" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mvalue" varchar(4096) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."instance_metadata" OWNER TO "postgres";

-- ----------------------------
-- Records of instance_metadata
-- ----------------------------
BEGIN;
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111u', 'mkey', '6666u', '2023-06-10 18:50:30', '2023-06-10 18:50:30');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111u', 'mvalue', '7777u', '2023-06-10 18:50:30', '2023-06-10 18:50:30');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111i', 'mkey', '6666i', '2023-06-10 18:56:28', '2023-06-10 18:56:28');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111i', 'mvalue', '7777i', '2023-06-10 18:56:28', '2023-06-10 18:56:28');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111t', 'mvalue', '7777t', '2023-06-10 19:11:05', '2023-06-10 19:11:05');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111t', 'mkey', '6666t', '2023-06-10 19:11:05', '2023-06-10 19:11:05');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('111', 'aaa', '1111', '2023-06-11 01:35:36', '2023-06-11 01:35:36');
INSERT INTO "public"."instance_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('111', 'bbb', '2222', '2023-06-11 01:35:36', '2023-06-11 01:35:36');
COMMIT;

-- ----------------------------
-- Table structure for leader_election
-- ----------------------------
DROP TABLE IF EXISTS "public"."leader_election";
CREATE TABLE "public"."leader_election" (
  "elect_key" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "version" int8 NOT NULL DEFAULT 0,
  "leader" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."leader_election" OWNER TO "postgres";

-- ----------------------------
-- Records of leader_election
-- ----------------------------
BEGIN;
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('aaa', 0, 'bbb', '2023-06-08 21:02:58.295652', '2023-06-08 21:02:58.295652');
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('val1', 0, 'val2', '2023-06-08 21:27:28.860113', '2023-06-08 21:27:28.860113');
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('val3', 0, 'val4', '2023-06-08 21:28:04.848416', '2023-06-08 21:28:04.848416');
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('val5', 0, 'val6', '2023-06-08 21:28:22.117857', '2023-06-08 21:28:22.117857');
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('val7', 0, 'val8', '2023-06-08 21:30:35.287969', '2023-06-08 21:30:35.287969');
INSERT INTO "public"."leader_election" ("elect_key", "version", "leader", "ctime", "mtime") VALUES ('test0', 0, '', '2023-06-08 22:04:50.696531', '2023-06-08 22:04:50.696531');
COMMIT;

-- ----------------------------
-- Table structure for mesh
-- ----------------------------
DROP TABLE IF EXISTS "public"."mesh";
CREATE TABLE "public"."mesh" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "department" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "business" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "managed" int2 NOT NULL,
  "istio_version" varchar(64) COLLATE "pg_catalog"."default",
  "data_cluster" varchar(1024) COLLATE "pg_catalog"."default",
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "token" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."mesh" OWNER TO "postgres";

-- ----------------------------
-- Records of mesh
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mesh_resource
-- ----------------------------
DROP TABLE IF EXISTS "public"."mesh_resource";
CREATE TABLE "public"."mesh_resource" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mesh_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "mesh_namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "type_url" varchar(96) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "body" text COLLATE "pg_catalog"."default",
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."mesh_resource" OWNER TO "postgres";

-- ----------------------------
-- Records of mesh_resource
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mesh_resource_revision
-- ----------------------------
DROP TABLE IF EXISTS "public"."mesh_resource_revision";
CREATE TABLE "public"."mesh_resource_revision" (
  "mesh_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "type_url" varchar(96) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."mesh_resource_revision" OWNER TO "postgres";

-- ----------------------------
-- Records of mesh_resource_revision
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mesh_service
-- ----------------------------
DROP TABLE IF EXISTS "public"."mesh_service";
CREATE TABLE "public"."mesh_service" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mesh_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mesh_namespace" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "mesh_service" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "location" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "export_to" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."mesh_service" OWNER TO "postgres";

-- ----------------------------
-- Records of mesh_service
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mesh_service_revision
-- ----------------------------
DROP TABLE IF EXISTS "public"."mesh_service_revision";
CREATE TABLE "public"."mesh_service_revision" (
  "mesh_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."mesh_service_revision" OWNER TO "postgres";

-- ----------------------------
-- Records of mesh_service_revision
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for namespace
-- ----------------------------
DROP TABLE IF EXISTS "public"."namespace";
CREATE TABLE "public"."namespace" (
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "comment" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "token" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(1024) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."namespace" OWNER TO "postgres";

-- ----------------------------
-- Records of namespace
-- ----------------------------
BEGIN;
INSERT INTO "public"."namespace" ("name", "comment", "token", "owner", "flag", "ctime", "mtime") VALUES ('Polaris', 'Polaris-server', '2d1bfe5d12e04d54b8ee69e62494c7fd', 'polaris', 0, '2019-09-06 07:55:07', '2019-09-06 07:55:07');
INSERT INTO "public"."namespace" ("name", "comment", "token", "owner", "flag", "ctime", "mtime") VALUES ('default', 'Default Environment', 'e2e473081d3d4306b52264e49f7ce227', 'polaris', 0, '2021-07-27 19:37:37', '2021-07-27 19:37:37');
INSERT INTO "public"."namespace" ("name", "comment", "token", "owner", "flag", "ctime", "mtime") VALUES ('Test1', 'Polaris-test1', '2d1bfe5d12e04d54b8ee69e62494c7fe', 'polaris', 0, '2023-06-07 21:14:27', '2023-06-07 21:48:01');
INSERT INTO "public"."namespace" ("name", "comment", "token", "owner", "flag", "ctime", "mtime") VALUES ('Test', 'Polaris-test1', '2d1bfe5d12e04d54b8ee69e62494c7fr', 'polaris', 0, '2023-06-03 16:02:03', '2023-06-07 21:49:45');
COMMIT;

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

-- ----------------------------
-- Records of owner_service_map
-- ----------------------------
BEGIN;
INSERT INTO "public"."owner_service_map" ("id", "owner", "service", "namespace") VALUES ('6f0b97ece75c4affbbccf958651c8430', 'polaris', 'Test1', '2222');
COMMIT;

-- ----------------------------
-- Table structure for ratelimit_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."ratelimit_config";
CREATE TABLE "public"."ratelimit_config" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "disable" int2 NOT NULL DEFAULT 0::smallint,
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "method" varchar(512) COLLATE "pg_catalog"."default" NOT NULL,
  "labels" text COLLATE "pg_catalog"."default" NOT NULL,
  "priority" int2 NOT NULL DEFAULT 0::smallint,
  "rule" text COLLATE "pg_catalog"."default" NOT NULL,
  "revision" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now(),
  "etime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."ratelimit_config" OWNER TO "postgres";

-- ----------------------------
-- Records of ratelimit_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ratelimit_revision
-- ----------------------------
DROP TABLE IF EXISTS "public"."ratelimit_revision";
CREATE TABLE "public"."ratelimit_revision" (
  "service_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "last_revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."ratelimit_revision" OWNER TO "postgres";

-- ----------------------------
-- Records of ratelimit_revision
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for routing_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."routing_config";
CREATE TABLE "public"."routing_config" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "in_bounds" text COLLATE "pg_catalog"."default",
  "out_bounds" text COLLATE "pg_catalog"."default",
  "revision" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."routing_config" OWNER TO "postgres";

-- ----------------------------
-- Records of routing_config
-- ----------------------------
BEGIN;
INSERT INTO "public"."routing_config" ("id", "in_bounds", "out_bounds", "revision", "flag", "ctime", "mtime") VALUES ('1111', '2223', '3333', '4444', 0, '2023-06-11 14:34:39', '2023-06-11 14:39:27');
COMMIT;

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
  "priority" int2 NOT NULL DEFAULT 0::smallint,
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now(),
  "etime" timestamp(6) NOT NULL DEFAULT now(),
  "extend_info" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT ''::character varying
)
;
ALTER TABLE "public"."routing_config_v2" OWNER TO "postgres";

-- ----------------------------
-- Records of routing_config_v2
-- ----------------------------
BEGIN;
INSERT INTO "public"."routing_config_v2" ("id", "name", "namespace", "policy", "config", "enable", "revision", "description", "priority", "flag", "ctime", "mtime", "etime", "extend_info") VALUES ('1111', '3333', '2222', '4444', '5555', 1, '6666', '7777', 1, 0, '2023-06-11 17:21:15', '2023-06-11 17:30:37', '2023-06-11 17:30:37', '');
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "reference" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "refer_filter" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "platform_id" varchar(32) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."service" OWNER TO "postgres";

-- ----------------------------
-- Records of service
-- ----------------------------
BEGIN;
INSERT INTO "public"."service" ("id", "name", "namespace", "ports", "business", "department", "cmdb_mod1", "cmdb_mod2", "cmdb_mod3", "comment", "token", "revision", "owner", "flag", "reference", "refer_filter", "platform_id", "ctime", "mtime") VALUES ('1111', 'Test1', '2222', '4444', '33333', '555', '666', '777', '888', 'Polaris-test', '2d1bfe5d12e04d54b8ee69e62494c7fe', '999', 'polaris', 0, '', NULL, '121212', '2023-06-09 15:50:31', '2023-06-09 22:02:22');
COMMIT;

-- ----------------------------
-- Table structure for service_metadata
-- ----------------------------
DROP TABLE IF EXISTS "public"."service_metadata";
CREATE TABLE "public"."service_metadata" (
  "id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mkey" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "mvalue" varchar(4096) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."service_metadata" OWNER TO "postgres";

-- ----------------------------
-- Records of service_metadata
-- ----------------------------
BEGIN;
INSERT INTO "public"."service_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1112', 'a', 'b', '2023-06-09 15:44:56', '2023-06-09 15:44:56');
INSERT INTO "public"."service_metadata" ("id", "mkey", "mvalue", "ctime", "mtime") VALUES ('1111', 'a', 'b', '2023-06-09 21:52:21', '2023-06-09 21:52:21');
COMMIT;

-- ----------------------------
-- Table structure for start_lock
-- ----------------------------
DROP TABLE IF EXISTS "public"."start_lock";
CREATE TABLE "public"."start_lock" (
  "lock_id" int4 NOT NULL,
  "lock_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "server" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."start_lock" OWNER TO "postgres";

-- ----------------------------
-- Records of start_lock
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_ip_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_ip_config";
CREATE TABLE "public"."t_ip_config" (
  "fip" int4 NOT NULL,
  "fareaid" int4 NOT NULL,
  "fcityid" int4 NOT NULL,
  "fidcid" int4 NOT NULL,
  "fflag" int2 DEFAULT 0::smallint,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_ip_config" OWNER TO "postgres";

-- ----------------------------
-- Records of t_ip_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_policy
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_policy";
CREATE TABLE "public"."t_policy" (
  "fmodid" int4 NOT NULL,
  "fdiv" int4 NOT NULL,
  "fmod" int4 NOT NULL,
  "fflag" int2 DEFAULT 0::smallint,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_policy" OWNER TO "postgres";

-- ----------------------------
-- Records of t_policy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_route
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_route";
CREATE TABLE "public"."t_route" (
  "fip" int4 NOT NULL,
  "fmodid" int4 NOT NULL,
  "fcmdid" int4 NOT NULL,
  "fsetid" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "fflag" int2 DEFAULT 0::smallint,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_route" OWNER TO "postgres";

-- ----------------------------
-- Records of t_route
-- ----------------------------
BEGIN;
INSERT INTO "public"."t_route" ("fip", "fmodid", "fcmdid", "fsetid", "fflag", "fstamp", "fflow") VALUES (1, 1, 1, '1', 4, '2023-06-11 17:30:37', 4);
COMMIT;

-- ----------------------------
-- Table structure for t_section
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_section";
CREATE TABLE "public"."t_section" (
  "fmodid" int4 NOT NULL,
  "ffrom" int4 NOT NULL,
  "fto" int4 NOT NULL,
  "fxid" int4 NOT NULL,
  "fflag" int2 DEFAULT 0::smallint,
  "fstamp" timestamp(6) NOT NULL,
  "fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_section" OWNER TO "postgres";

-- ----------------------------
-- Records of t_section
-- ----------------------------
BEGIN;
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

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
  "flag" int2 NOT NULL DEFAULT 0::smallint,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."user_group" OWNER TO "postgres";

-- ----------------------------
-- Records of user_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_group_relation
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group_relation";
CREATE TABLE "public"."user_group_relation" (
  "user_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "group_id" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "ctime" timestamp(6) NOT NULL DEFAULT now(),
  "mtime" timestamp(6) NOT NULL DEFAULT now()
)
;
ALTER TABLE "public"."user_group_relation" OWNER TO "postgres";

-- ----------------------------
-- Records of user_group_relation
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table auth_principal
-- ----------------------------
ALTER TABLE "public"."auth_principal" ADD CONSTRAINT "auth_principal_pkey" PRIMARY KEY ("strategy_id", "principal_id", "principal_role");

-- ----------------------------
-- Primary Key structure for table auth_strategy
-- ----------------------------
ALTER TABLE "public"."auth_strategy" ADD CONSTRAINT "auth_strategy_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table auth_strategy_resource
-- ----------------------------
ALTER TABLE "public"."auth_strategy_resource" ADD CONSTRAINT "auth_strategy_resource_pkey" PRIMARY KEY ("strategy_id", "res_type", "res_id");

-- ----------------------------
-- Primary Key structure for table business
-- ----------------------------
ALTER TABLE "public"."business" ADD CONSTRAINT "business_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule" ADD CONSTRAINT "circuitbreaker_rule_pkey" PRIMARY KEY ("id", "version");

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule_relation
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule_relation" ADD CONSTRAINT "circuitbreaker_rule_relation_pkey" PRIMARY KEY ("service_id");

-- ----------------------------
-- Primary Key structure for table circuitbreaker_rule_v2
-- ----------------------------
ALTER TABLE "public"."circuitbreaker_rule_v2" ADD CONSTRAINT "circuitbreaker_rule_v2_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table cl5_module
-- ----------------------------
ALTER TABLE "public"."cl5_module" ADD CONSTRAINT "cl5_module_pkey" PRIMARY KEY ("module_id");

-- ----------------------------
-- Primary Key structure for table client
-- ----------------------------
ALTER TABLE "public"."client" ADD CONSTRAINT "client_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table client_stat
-- ----------------------------
ALTER TABLE "public"."client_stat" ADD CONSTRAINT "client_stat_pkey" PRIMARY KEY ("client_id", "target", "port");

-- ----------------------------
-- Primary Key structure for table config_file
-- ----------------------------
ALTER TABLE "public"."config_file" ADD CONSTRAINT "config_file_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_group
-- ----------------------------
ALTER TABLE "public"."config_file_group" ADD CONSTRAINT "config_file_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_release
-- ----------------------------
ALTER TABLE "public"."config_file_release" ADD CONSTRAINT "config_file_release_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_release_history
-- ----------------------------
ALTER TABLE "public"."config_file_release_history" ADD CONSTRAINT "config_file_release_history_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_tag
-- ----------------------------
ALTER TABLE "public"."config_file_tag" ADD CONSTRAINT "config_file_tag_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table config_file_template
-- ----------------------------
ALTER TABLE "public"."config_file_template" ADD CONSTRAINT "config_file_template_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table fault_detect_rule
-- ----------------------------
ALTER TABLE "public"."fault_detect_rule" ADD CONSTRAINT "fault_detect_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table health_check
-- ----------------------------
ALTER TABLE "public"."health_check" ADD CONSTRAINT "health_check_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table instance
-- ----------------------------
CREATE INDEX "host" ON "public"."instance" USING btree (
  "host" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "mtime" ON "public"."instance" USING btree (
  "mtime" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);
CREATE INDEX "service_id" ON "public"."instance" USING btree (
  "service_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table instance
-- ----------------------------
ALTER TABLE "public"."instance" ADD CONSTRAINT "instance_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table instance_metadata
-- ----------------------------
CREATE INDEX "mkey" ON "public"."instance_metadata" USING btree (
  "mkey" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table instance_metadata
-- ----------------------------
ALTER TABLE "public"."instance_metadata" ADD CONSTRAINT "instance_metadata_pkey" PRIMARY KEY ("id", "mkey");

-- ----------------------------
-- Indexes structure for table leader_election
-- ----------------------------
CREATE INDEX "version" ON "public"."leader_election" USING btree (
  "version" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table leader_election
-- ----------------------------
ALTER TABLE "public"."leader_election" ADD CONSTRAINT "leader_election_pkey" PRIMARY KEY ("elect_key");

-- ----------------------------
-- Primary Key structure for table mesh
-- ----------------------------
ALTER TABLE "public"."mesh" ADD CONSTRAINT "mesh_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mesh_resource
-- ----------------------------
ALTER TABLE "public"."mesh_resource" ADD CONSTRAINT "mesh_resource_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mesh_resource_revision
-- ----------------------------
ALTER TABLE "public"."mesh_resource_revision" ADD CONSTRAINT "mesh_resource_revision_pkey" PRIMARY KEY ("mesh_id", "type_url");

-- ----------------------------
-- Primary Key structure for table mesh_service
-- ----------------------------
ALTER TABLE "public"."mesh_service" ADD CONSTRAINT "mesh_service_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mesh_service_revision
-- ----------------------------
ALTER TABLE "public"."mesh_service_revision" ADD CONSTRAINT "mesh_service_revision_pkey" PRIMARY KEY ("mesh_id");

-- ----------------------------
-- Uniques structure for table namespace
-- ----------------------------
ALTER TABLE "public"."namespace" ADD CONSTRAINT "namespace_name_key" UNIQUE ("name");

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
-- Primary Key structure for table routing_config_v2
-- ----------------------------
ALTER TABLE "public"."routing_config_v2" ADD CONSTRAINT "routing_config_v2_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table service
-- ----------------------------
ALTER TABLE "public"."service" ADD CONSTRAINT "service_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table service_metadata
-- ----------------------------
ALTER TABLE "public"."service_metadata" ADD CONSTRAINT "service_metadata_pkey" PRIMARY KEY ("id", "mkey");

-- ----------------------------
-- Primary Key structure for table start_lock
-- ----------------------------
ALTER TABLE "public"."start_lock" ADD CONSTRAINT "start_lock_pkey" PRIMARY KEY ("lock_id", "lock_key");

-- ----------------------------
-- Primary Key structure for table t_ip_config
-- ----------------------------
ALTER TABLE "public"."t_ip_config" ADD CONSTRAINT "t_ip_config_pkey" PRIMARY KEY ("fip");

-- ----------------------------
-- Primary Key structure for table t_policy
-- ----------------------------
ALTER TABLE "public"."t_policy" ADD CONSTRAINT "t_policy_pkey" PRIMARY KEY ("fmodid");

-- ----------------------------
-- Primary Key structure for table t_route
-- ----------------------------
ALTER TABLE "public"."t_route" ADD CONSTRAINT "t_route_pkey" PRIMARY KEY ("fip", "fmodid", "fcmdid");

-- ----------------------------
-- Primary Key structure for table t_section
-- ----------------------------
ALTER TABLE "public"."t_section" ADD CONSTRAINT "t_section_pkey" PRIMARY KEY ("fmodid", "ffrom", "fto");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_group
-- ----------------------------
ALTER TABLE "public"."user_group" ADD CONSTRAINT "user_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_group_relation
-- ----------------------------
ALTER TABLE "public"."user_group_relation" ADD CONSTRAINT "user_group_relation_pkey" PRIMARY KEY ("user_id", "group_id");
