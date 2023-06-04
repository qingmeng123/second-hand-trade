create database if not exists second_hand_trade default character set utf8mb4 collate utf8mb4_general_ci;

use second_hand_trade;

#用户表
create table if not exists userinfo
(
    uid          int auto_increment
        primary key,
    username     varchar(20)                    null,
    password     varchar(80) default '666666'   null,
    gender       tinyint(1)  default 0          null,
    name         varchar(20) default '匿名用户' null,
    phone        varchar(20) default '0'        null,
    money        float       default 0          null,
    address_id   int         default 0          null,
    group_id     int         default 0          null,
    store_id     int         default 0          null,
    salt         varchar(80) default '0'        null,
    constraint user_username_uindex
        unique (username)

);

#地址表
create table if not exists address_info
(
    address_id int auto_increment
        primary key,
    uid        int         default 0   not null,
    name       varchar(20) default '0' null,
    phone      varchar(20) default '0' null,
    address    varchar(20) default '0' null,
    foreign key (uid) references userinfo(uid)
);

#二手店铺表
create table if not exists store
(
    store_id    int auto_increment
        primary key,
    store_name  varchar(20) default '0'        not null,
    notice      varchar(50) default '暂无公告' null,
    store_money float       default 0          null,
    constraint store_store_name_uindex
        unique (store_name)
);

#类别表（如二手家具，书等）
create table if not exists sort
(
    sort_id   int auto_increment
        primary key,
    sort_name varchar(20) default '0' null
);


#二手商品表
create table if not exists goods_info
(
    goods_id       int auto_increment
        primary key,
    sort_id        int          default 0   not null,
    store_id       int          default 0   null,
    goods_name     varchar(20)  default '0' null,
    picture        varchar(40)  default '0' null,
    price          float        default 0   null,
    goods_intro    varchar(100) default '0' null,
    turnover       int          default 0   null,
    style          varchar(20)  default '0' null,
    number         int          default 0   null,
    shelf_date     datetime                 null,
    index idx_goods_name(goods_name,price),
    foreign key (sort_id)references sort(sort_id),
    foreign key (store_id)references store(store_id)
);


#购物车
create table if not exists shopping_cart
(
    cart_id  int auto_increment
        primary key,
    uid      int          default 0   null,
    goods_id int          default 0   null,
    number   int          default 0   null,
    remark   varchar(100) default '0' null,
    state    tinyint(1)   default 0   null,
    order_id int          default 0   null,
    foreign key (uid)references userinfo(uid)
);

#订单详情
create table if not exists order_details
(
    order_id        int auto_increment
        primary key,
    address_id      int        default 0 null,
    date            datetime             null,
    order_state     tinyint(1) default 0 null,
    confirm_receipt tinyint(1) default 0 null,
    money           float      default 0 null,
    foreign key (address_id)references address_info(address_id)
);

