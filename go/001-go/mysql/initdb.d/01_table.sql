/* ユーザーマスタ */
create table if not exists user_master (
    user_id varchar(16) not null,
    password varchar(255),
    user_name varchar(255),
    mail_address varchar(255),
    authority decimal(3),
    last_login_date datetime,
    create_date datetime,
    update_date datetime
);
alter table user_master add primary key (user_id);