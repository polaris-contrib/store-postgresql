/*
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

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
INSERT INTO "public"."namespace" ("name",
                                  "comment",
                                  "token",
                                  "owner",
                                  "flag",
                                  "ctime",
                                  "mtime")
VALUES ('Polaris',
        'Polaris-server',
        '2d1bfe5d12e04d54b8ee69e62494c7fd',
        'polaris',
        0,
        '2019-09-06 07:55:07',
        '2019-09-06 07:55:07'),
       ('default',
        'Default Environment',
        'e2e473081d3d4306b52264e49f7ce227',
        'polaris',
        0,
        '2021-07-27 19:37:37',
        '2021-07-27 19:37:37');
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
INSERT INTO "public"."service" ("id",
                                "name",
                                "namespace",
                                "ports",
                                "comment",
                                "business",
                                "department",
                                "cmdb_mod1",
                                "cmdb_mod2",
                                "cmdb_mod3",
                                "token",
                                "revision",
                                "owner",
                                "flag",
                                "reference",
                                "ctime",
                                "mtime")
VALUES ('fbca9bfa04ae4ead86e1ecf5811e32a9',
        'polaris.checker',
        'Polaris',
        '0',
        'polaris checker service',
        'polaris',
        'polaris department',
        'polaris cmdb_mod1',
        'polaris cmdb_mod2',
        'polaris cmdb_mod3',
        '7d19c46de327408d8709ee7392b7700b',
        '301b1e9f0bbd47a6b697e26e99dfe012',
        'polaris',
        0,
        '',
        '2021-09-06 07:55:07',
        '2021-09-06 07:55:09');
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
-- Table structure for t_ip_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_ip_config";
CREATE TABLE "public"."t_ip_config" (
    "Fip" int4 NOT NULL,
    "FareaId" int4 NOT NULL,
    "FcityId" int4 NOT NULL,
    "FidcId" int4 NOT NULL,
    "Fflag" int2 DEFAULT 0::smallint,
    "Fstamp" timestamp(6) NOT NULL,
    "Fflow" int4 NOT NULL
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
    "FmodId" int4 NOT NULL,
    "Fdiv" int4 NOT NULL,
    "Fmod" int4 NOT NULL,
    "Fflag" int2 DEFAULT 0::smallint,
    "Fstamp" timestamp(6) NOT NULL,
    "Fflow" int4 NOT NULL
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
    "Fip" int4 NOT NULL,
    "FmodId" int4 NOT NULL,
    "FcmdId" int4 NOT NULL,
    "Fsetid" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
    "Fflag" int2 DEFAULT 0::smallint,
    "Fstamp" timestamp(6) NOT NULL,
    "Fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_route" OWNER TO "postgres";

-- ----------------------------
-- Records of t_route
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_section
-- ----------------------------
DROP TABLE IF EXISTS "public"."t_section";
CREATE TABLE "public"."t_section" (
    "Fmodid" int4 NOT NULL,
    "Ffrom" int4 NOT NULL,
    "Fto" int4 NOT NULL,
    "Fxid" int4 NOT NULL,
    "Fflag" int2 DEFAULT 0::smallint,
    "Fstamp" timestamp(6) NOT NULL,
    "Fflow" int4 NOT NULL
)
;
ALTER TABLE "public"."t_section" OWNER TO "postgres";

-- ----------------------------
-- Records of t_section
-- ----------------------------
BEGIN;
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
INSERT INTO "start_lock" ("lock_id", "lock_key", "server", "mtime")
VALUES (1, 'sz', 'aaa', '2019-12-05 08:35:49');
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
insert into cl5_module(module_id, interface_id, range_num) values (3000001, 1, 0);
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
    "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "business" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "department" varchar(1024) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "metadata" text COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "flag" int4 NOT NULL
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
    "format" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
    "content" text COLLATE "pg_catalog"."default" NOT NULL,
    "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
    "version" int4 NOT NULL,
    "flag" int2 NOT NULL DEFAULT 0::smallint,
    "create_time" timestamp(6) NOT NULL DEFAULT now(),
    "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "modify_time" timestamp(6) NOT NULL DEFAULT now(),
    "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "tags" text COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "active" int2 NOT NULL DEFAULT 0::smallint,
    "description" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
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
    "comment" varchar(512) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "md5" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
    "type" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
    "status" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'success'::character varying,
    "create_time" timestamp(6) NOT NULL DEFAULT now(),
    "create_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "modify_time" timestamp(6) NOT NULL DEFAULT now(),
    "modify_by" varchar(32) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "tags" text COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "version" bigint NOT NULL,
    "reason" varchar(5000) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
    "description" varchar(1000) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
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
INSERT INTO "user" ("id",
                    "name",
                    "password",
                    "source",
                    "token",
                    "token_enable",
                    "user_type",
                    "comment",
                    "mobile",
                    "email",
                    "owner")
VALUES ('65e4789a6d5b49669adf1e9e8387549c',
        'polaris',
        '$2a$10$3izWuZtE5SBdAtSZci.gs.iZ2pAn9I8hEqYrC6gwJp1dyjqQnrrum',
        'Polaris',
        'nu/0WRA4EqSR1FagrjRj0fZwPXuGlMpX+zCuWu4uMqy8xr1vRjisSbA25aAC3mtU8MeeRsKhQiDAynUR09I=',
        1,
        20,
        'default polaris admin account',
        '12345678910',
        '12345678910',
        '');
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
-- Insert permission policies and association relationships for Polaris-Admin accounts
INSERT INTO auth_principal(`strategy_id`, `principal_id`, `principal_role`) VALUE (
    'fbca9bfa04ae4ead86e1ecf5811e32a9',
    '65e4789a6d5b49669adf1e9e8387549c',
    1
);
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
INSERT INTO auth_strategy("id",
                          "name",
                          "action",
                          "owner",
                          "comment",
                          "default",
                          "revision",
                          "flag",
                          "ctime",
                          "mtime")
VALUES ('fbca9bfa04ae4ead86e1ecf5811e32a9',
        '(用户) polaris的默认策略',
        'READ_WRITE',
        '65e4789a6d5b49669adf1e9e8387549c',
        'default admin',
        1,
        'fbca9bfa04ae4ead86e1ecf5811e32a9',
        0,
        current_date,
        current_date);
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
INSERT INTO auth_strategy_resource("strategy_id",
                                     "res_type",
                                     "res_id",
                                     "ctime",
                                     "mtime")
VALUES ('fbca9bfa04ae4ead86e1ecf5811e32a9',
        0,
        '*',
        current_date,
        current_date),
       ('fbca9bfa04ae4ead86e1ecf5811e32a9',
        1,
        '*',
        current_date,
        current_date),
       ('fbca9bfa04ae4ead86e1ecf5811e32a9',
        2,
        '*',
        current_date,
        current_date);
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
INSERT INTO "public"."config_file_template" ("id", "name", "content", "format", "comment", "create_time", "create_by", "modify_time", "modify_by") VALUES (1, 'spring-cloud-gateway-braining', '{
    "rules":[
        {
            "conditions":[
                {
                    "key":"${http.query.uid}",
                    "values":["10000"],
                    "operation":"EQUALS"
                }
            ],
            "labels":[
                {
                    "key":"env",
                    "value":"green"
                }
            ]
        }
    ]
}
', 'json', 'Spring Cloud Gateway  染色规则', '2023-06-13 01:40:24', 'polaris', '2023-06-13 01:40:34', 'polaris');
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
ALTER TABLE "public"."t_ip_config" ADD CONSTRAINT "t_ip_config_pkey" PRIMARY KEY ("Fip");

-- ----------------------------
-- Primary Key structure for table t_policy
-- ----------------------------
ALTER TABLE "public"."t_policy" ADD CONSTRAINT "t_policy_pkey" PRIMARY KEY ("FmodId");

-- ----------------------------
-- Primary Key structure for table t_route
-- ----------------------------
ALTER TABLE "public"."t_route" ADD CONSTRAINT "t_route_pkey" PRIMARY KEY ("Fip", "Fmodid", "Fcmdid");

-- ----------------------------
-- Primary Key structure for table t_section
-- ----------------------------
ALTER TABLE "public"."t_section" ADD CONSTRAINT "t_section_pkey" PRIMARY KEY ("Fmodid", "Ffrom", "Fto");

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
