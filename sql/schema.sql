SET FOREIGN_KEY_CHECKS = 0;
SET NAMES utf8mb4;
-- sys_dept DDL
CREATE TABLE `sys_dept` (`id` BIGINT NOT NULL AUTO_INCREMENT Comment "id",
`parent_id` BIGINT NULL,
`dept_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
`create_time` DATETIME NULL,
`update_time` DATETIME NULL,
`delete_time` DATETIME NULL,
`del_flag` TINYINT NULL,
`remark` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
PRIMARY KEY (`id`)) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;
-- sys_role DDL
CREATE TABLE `sys_role` (`id` BIGINT NOT NULL AUTO_INCREMENT,
`role_code` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL Comment "角色编码",
`role_name` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL Comment "角色名称",
`del_flag` CHAR(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' Comment "0可用1已删除",
`status` TINYINT NOT NULL DEFAULT 0 Comment "0可用1停用",
`create_time` DATETIME NULL,
`update_time` DATETIME NULL,
`delete_time` DATETIME NULL,
`remark` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
PRIMARY KEY (`id`)) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci AUTO_INCREMENT = 2 ROW_FORMAT = Dynamic;
-- sys_user DDL
CREATE TABLE `sys_user` (`id` BIGINT NOT NULL AUTO_INCREMENT,
`dept_id` BIGINT NULL,
`user_name` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL Comment "用户名",
`real_name` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL Comment "姓名",
`password` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL Comment "密码",
`del_flag` CHAR(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' Comment "0可用1已删除",
`status` TINYINT NOT NULL DEFAULT 0 Comment "0正常1停用",
`create_time` DATETIME NULL,
`update_time` DATETIME NULL,
`delete_time` DATETIME NULL,
`remark` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
PRIMARY KEY (`id`)) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci AUTO_INCREMENT = 18 ROW_FORMAT = Dynamic;
-- sys_user_role DDL
CREATE TABLE `sys_user_role` (`user_id` BIGINT NOT NULL,
`role_id` BIGINT NOT NULL,
PRIMARY KEY (`user_id`,
`role_id`)) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;
