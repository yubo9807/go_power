/*
 Navicat Premium Data Transfer

 Source Server         : blog
 Source Server Type    : MariaDB
 Source Server Version : 50568 (5.5.68-MariaDB)
 Source Host           : 47.92.216.215:3306
 Source Schema         : permissions

 Target Server Type    : MariaDB
 Target Server Version : 50568 (5.5.68-MariaDB)
 File Encoding         : 65001

 Date: 09/11/2023 09:31:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for s_correlation
-- ----------------------------
DROP TABLE IF EXISTS `s_correlation`;
CREATE TABLE `s_correlation` (
  `id` bigint(11) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `table_id` bigint(20) NOT NULL COMMENT '可以是 menu / interface / button 的 id',
  `table_type` varchar(255) NOT NULL COMMENT 'menu / interface / button ',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for s_element
-- ----------------------------
DROP TABLE IF EXISTS `s_element`;
CREATE TABLE `s_element` (
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
-- Table structure for s_interface
-- ----------------------------
DROP TABLE IF EXISTS `s_interface`;
CREATE TABLE `s_interface` (
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
-- Table structure for s_menu
-- ----------------------------
DROP TABLE IF EXISTS `s_menu`;
CREATE TABLE `s_menu` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `parent` bigint(20) DEFAULT NULL COMMENT '父级菜单ID',
  `count` int(11) NOT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `name` (`name`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for s_roles
-- ----------------------------
DROP TABLE IF EXISTS `s_roles`;
CREATE TABLE `s_roles` (
  `id` bigint(20) NOT NULL,
  `role` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `parent` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
