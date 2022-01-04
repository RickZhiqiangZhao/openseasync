create database if not exists `openseasync` charset = 'utf8';
use `openseasync`;

create table if not exists assets
(
	id bigint auto_increment
		primary key,
	user_address char(42) not null,
	title varchar(255) null,
	image_url text null,
	image_preview_url text null,
	image_thumbnail_url text null,
	description text null,
	contract_address char(42) not null,
	token_id varchar(128) not null,
	num_sales int null,
	owner char(42) not null,
	owner_img_url text null,
	creator char(42) not null,
	creator_img_url text null,
	token_metadata text null,
	slug varchar(255) not null,
	is_delete int null,
    refresh_time int null
)
charset=utf8mb4;

create table if not exists assets_top_ownerships
(
	id bigint auto_increment
		primary key,
    user_address char(42) not null,
	contract_address char(42) not null,
	token_id varchar(128) not null,
	owner char(42) not null,
	profile_img_url text null,
	quantity varchar(128) null,
	is_delete int null,
    refresh_time int null
)
charset=utf8mb4;

create table if not exists collections
(
	id bigint auto_increment
		primary key,
	slug varchar(255) not null,
    user_address char(42) not null,
	name varchar(255) null,
	banner_image_url text null,
	description text null,
	image_url text null,
	large_image_url text null,
	is_delete int null,
	create_date varchar(32) null,
    refresh_time int null,
    num_owners int null,
    total_supply int null,
    total_volume float null,
    owned_asset_count varchar(64) null

)
charset=utf8mb4;

create table if not exists contracts
(
	id bigint auto_increment
		primary key,
	address char(42) not null,
	contract_name varchar(255) null,
	contract_type varchar(64) null,
	symbol varchar(64) null,
	schema_name varchar(32) null,
	total_supply varchar(128) null,
	Description text null
)
charset=utf8mb4;

create table if not exists traits
(
	id bigint auto_increment
		primary key,
    user_address char(42) not null,
	contract_address char(42) not null,
	token_id varchar(128) not null,
	trait_type varchar(255) null,
	value varchar(255) null,
	display_type varchar(255) null,
	max_value int null,
	trait_count int null,
	order_by varchar(255) null,
	is_delete int null,
    refresh_time int null
)
charset=utf8mb4;