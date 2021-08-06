#数据库
/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : blog-service

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 06/08/2021 10:33:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_ahth
-- ----------------------------
DROP TABLE IF EXISTS `blog_ahth`;
CREATE TABLE `blog_ahth`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Secrect',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否被删除 0 未被删除 1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '认证管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标题',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '文章简介',
  `cover_image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '封面地址',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文章内容',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `is_del` tinyint(3) NOT NULL DEFAULT 0 COMMENT '0 未删除 1已删除',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用 1可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章id',
  `tag_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签id',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) NOT NULL COMMENT '0 未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章标签关联' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标签',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除 1已删除',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用 1可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签管理' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
