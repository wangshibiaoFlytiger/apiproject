create database apiproject default character set utf8mb4 collate utf8mb4_unicode_ci;
use apiproject;

create table if not exists video
(
	id varchar(255) not null
		primary key,
	site_id varchar(255) null comment '站点ID',
	title varchar(255) null comment '标题',
	created_at datetime null comment '创建时间',
	updated_at datetime null comment '更新时间',
	deleted_at datetime null comment '删除时间'
)
comment '视频表' collate=utf8mb4_unicode_ci;