/*
 Navicat Premium Data Transfer

 Source Server         : blog
 Source Server Type    : MariaDB
 Source Server Version : 50568
 Source Host           : 47.92.216.215:3306
 Source Schema         : permissions

 Target Server Type    : MariaDB
 Target Server Version : 50568
 File Encoding         : 65001

 Date: 09/01/2023 20:17:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for correlation
-- ----------------------------
DROP TABLE IF EXISTS `correlation`;
CREATE TABLE `correlation` (
  `id` bigint(11) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `table_id` bigint(20) NOT NULL COMMENT '可以是 menu / interface / button 的 id',
  `table_type` varchar(255) NOT NULL COMMENT 'menu / interface / button ',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for element
-- ----------------------------
DROP TABLE IF EXISTS `element`;
CREATE TABLE `element` (
  `id` bigint(20) NOT NULL,
  `key` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `menu_id` bigint(20) DEFAULT NULL COMMENT '归属菜单（页面）',
  PRIMARY KEY (`id`),
  KEY `affiliation_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for interface
-- ----------------------------
DROP TABLE IF EXISTS `interface`;
CREATE TABLE `interface` (
  `id` bigint(20) NOT NULL,
  `method` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `menu_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `affiliation_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `hidden` tinyint(1) NOT NULL DEFAULT '0',
  `parent` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL,
  `role` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
