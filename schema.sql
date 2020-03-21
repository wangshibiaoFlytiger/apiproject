create database apiproject default character set utf8mb4 collate utf8mb4_general_ci;
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

# 王世彪 2019-05-16 14:51:45
# 将主键类型改为整数
alter table video change id id bigint comment '主键';

# 王世彪 2020-03-20 10:40:52
# 增加定时任务表
create table cron_task(
    id bigint primary key comment '主键',
    type varchar(20) comment '任务类型: 1种任务对应后台的一个任务处理函数',
    entry_id int comment '该ID由任务管理器生成, 作为任务的唯一标识',
    title varchar(100) comment '标题',
    spec varchar(20) comment '定时描述表达式',
    remark varchar(255) comment '备注',
    status int default 1 comment '状态:1启用,2禁用',
    created_at datetime null comment '创建时间',
    updated_at datetime null comment '更新时间',
    deleted_at datetime null comment '删除时间'
) comment '定时任务表' collate=utf8mb4_unicode_ci;