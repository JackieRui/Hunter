
-- desc: 职位信息表
-- date: 2021-11-23 20:45:00
CREATE TABLE `job` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(20) NOT NULL COMMENT '职位名称',
  `company` varchar(20) NOT NULL COMMENT '公司名称',
  `content` text COMMENT '页面内容',
  `url` varchar(50) NOT NULL '页面链接',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '职位状态 0:初始',
  `task` varchar(50) NOT NULL COMMENT 'Task信息',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建记录的时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '职位信息表';

