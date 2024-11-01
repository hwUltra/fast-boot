/*
 Navicat Premium Data Transfer

 Source Server         : p
 Source Server Type    : PostgreSQL
 Source Server Version : 160001 (160001)
 Source Host           : localhost:15432
 Source Catalog        : mall-boot
 Source Schema         : mall-boot

 Target Server Type    : PostgreSQL
 Target Server Version : 160001 (160001)
 File Encoding         : 65001

 Date: 01/11/2024 19:14:49
*/


-- ----------------------------
-- Type structure for sys_log_module
-- ----------------------------
DROP TYPE IF EXISTS "mall-boot"."sys_log_module";
CREATE TYPE "mall-boot"."sys_log_module" AS ENUM (
  'LOGIN',
  'USER',
  'ROLE',
  'DEPT',
  'MENU',
  'DICT',
  'OTHER'
);
ALTER TYPE "mall-boot"."sys_log_module" OWNER TO "kyle";

-- ----------------------------
-- Sequence structure for oms_order_delivery_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_delivery_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_delivery_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for oms_order_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for oms_order_item_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_item_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_item_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for oms_order_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_log_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_log_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for oms_order_pay_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_pay_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_pay_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for oms_order_setting_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."oms_order_setting_id_seq";
CREATE SEQUENCE "mall-boot"."oms_order_setting_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_brand_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_brand_id_seq";
CREATE SEQUENCE "mall-boot"."pms_brand_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_category_attribute_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_category_attribute_id_seq";
CREATE SEQUENCE "mall-boot"."pms_category_attribute_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_category_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_category_id_seq";
CREATE SEQUENCE "mall-boot"."pms_category_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_goods_attribute_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_goods_attribute_id_seq";
CREATE SEQUENCE "mall-boot"."pms_goods_attribute_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_goods_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_goods_id_seq";
CREATE SEQUENCE "mall-boot"."pms_goods_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_shop_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_shop_id_seq";
CREATE SEQUENCE "mall-boot"."pms_shop_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pms_sku_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."pms_sku_id_seq";
CREATE SEQUENCE "mall-boot"."pms_sku_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_dept_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_dept_id_seq";
CREATE SEQUENCE "mall-boot"."sys_dept_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_dict_data_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_dict_data_id_seq";
CREATE SEQUENCE "mall-boot"."sys_dict_data_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_dict_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_dict_id_seq";
CREATE SEQUENCE "mall-boot"."sys_dict_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_log_id_seq";
CREATE SEQUENCE "mall-boot"."sys_log_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_menu_id_seq";
CREATE SEQUENCE "mall-boot"."sys_menu_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_notice_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_notice_id_seq";
CREATE SEQUENCE "mall-boot"."sys_notice_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_role_id_seq";
CREATE SEQUENCE "mall-boot"."sys_role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_user_id_seq";
CREATE SEQUENCE "mall-boot"."sys_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_user_notice_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."sys_user_notice_id_seq";
CREATE SEQUENCE "mall-boot"."sys_user_notice_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_address_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."user_address_id_seq";
CREATE SEQUENCE "mall-boot"."user_address_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_cart_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."user_cart_id_seq";
CREATE SEQUENCE "mall-boot"."user_cart_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."user_id_seq";
CREATE SEQUENCE "mall-boot"."user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_third_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mall-boot"."user_third_id_seq";
CREATE SEQUENCE "mall-boot"."user_third_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for oms_order
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order";
CREATE TABLE "mall-boot"."oms_order" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_id_seq'::regclass),
  "order_sn" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "total_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "total_quantity" int4 NOT NULL DEFAULT 0,
  "source" int2 NOT NULL,
  "status" int2 NOT NULL DEFAULT 101,
  "remark" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "member_id" int8 NOT NULL DEFAULT '0'::bigint,
  "coupon_id" int8 NOT NULL DEFAULT '0'::bigint,
  "coupon_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "freight_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "payment_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "pay_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "payment_time" timestamptz(6),
  "payment_method" int2,
  "out_trade_no" varchar(32) COLLATE "pg_catalog"."default",
  "transaction_id" varchar(32) COLLATE "pg_catalog"."default",
  "out_refund_no" varchar(32) COLLATE "pg_catalog"."default",
  "refund_id" varchar(32) COLLATE "pg_catalog"."default",
  "delivery_time" timestamptz(6),
  "receive_time" timestamptz(6),
  "comment_time" timestamptz(6),
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order"."order_sn" IS '订单号';
COMMENT ON COLUMN "mall-boot"."oms_order"."total_amount" IS '订单总额（分）';
COMMENT ON COLUMN "mall-boot"."oms_order"."total_quantity" IS '商品总数';
COMMENT ON COLUMN "mall-boot"."oms_order"."source" IS '订单来源(1:APP；2:网页)';
COMMENT ON COLUMN "mall-boot"."oms_order"."status" IS '订单状态：
101->待付款；
102->用户取消；
103->系统取消；
201->已付款；
202->申请退款；
203->已退款；
301->待发货；
401->已发货；
501->用户收货；
502->系统收货；
901->已完成；';
COMMENT ON COLUMN "mall-boot"."oms_order"."remark" IS '订单备注';
COMMENT ON COLUMN "mall-boot"."oms_order"."member_id" IS '会员id';
COMMENT ON COLUMN "mall-boot"."oms_order"."coupon_id" IS '使用的优惠券';
COMMENT ON COLUMN "mall-boot"."oms_order"."coupon_amount" IS '优惠券抵扣金额（分）';
COMMENT ON COLUMN "mall-boot"."oms_order"."freight_amount" IS '运费金额（分）';
COMMENT ON COLUMN "mall-boot"."oms_order"."payment_amount" IS '应付总额（分）';
COMMENT ON COLUMN "mall-boot"."oms_order"."pay_amount" IS '应付总额（分）';
COMMENT ON COLUMN "mall-boot"."oms_order"."payment_time" IS '支付时间';
COMMENT ON COLUMN "mall-boot"."oms_order"."payment_method" IS '支付方式(1：微信JSAPI；2：支付宝；3：余额；4：微信APP)';
COMMENT ON COLUMN "mall-boot"."oms_order"."out_trade_no" IS '微信支付等第三方支付平台的商户订单号';
COMMENT ON COLUMN "mall-boot"."oms_order"."transaction_id" IS '微信支付订单号';
COMMENT ON COLUMN "mall-boot"."oms_order"."out_refund_no" IS '商户退款单号';
COMMENT ON COLUMN "mall-boot"."oms_order"."refund_id" IS '微信退款单号';
COMMENT ON COLUMN "mall-boot"."oms_order"."delivery_time" IS '发货时间';
COMMENT ON COLUMN "mall-boot"."oms_order"."receive_time" IS '确认收货时间';
COMMENT ON COLUMN "mall-boot"."oms_order"."comment_time" IS '评价时间';
COMMENT ON COLUMN "mall-boot"."oms_order"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order" IS '订单详情表';

-- ----------------------------
-- Records of oms_order
-- ----------------------------

-- ----------------------------
-- Table structure for oms_order_delivery
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order_delivery";
CREATE TABLE "mall-boot"."oms_order_delivery" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_delivery_id_seq'::regclass),
  "order_id" int8 NOT NULL,
  "delivery_company" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "delivery_sn" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_phone" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_post_code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_province" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_city" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_region" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "receiver_detail_address" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "delivery_status" int2 DEFAULT '0'::smallint,
  "delivery_time" timestamptz(6),
  "receive_time" timestamptz(6),
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."order_id" IS '订单id';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."delivery_company" IS '物流公司(配送方式)';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."delivery_sn" IS '物流单号';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_name" IS '收货人姓名';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_phone" IS '收货人电话';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_post_code" IS '收货人邮编';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_province" IS '省份/直辖市';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_city" IS '城市';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_region" IS '区';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receiver_detail_address" IS '详细地址';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."remark" IS '备注';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."delivery_status" IS '物流状态【0->运输中；1->已收货】';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."delivery_time" IS '发货时间';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."receive_time" IS '确认收货时间';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order_delivery"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order_delivery" IS '订单物流记录表';

-- ----------------------------
-- Records of oms_order_delivery
-- ----------------------------

-- ----------------------------
-- Table structure for oms_order_item
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order_item";
CREATE TABLE "mall-boot"."oms_order_item" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_item_id_seq'::regclass),
  "order_id" int8 NOT NULL,
  "spu_name" varchar(128) COLLATE "pg_catalog"."default",
  "sku_id" int8 NOT NULL DEFAULT '0'::bigint,
  "sku_sn" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sku_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "pic_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "price" int8 NOT NULL DEFAULT '0'::bigint,
  "quantity" int4,
  "total_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order_item"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."order_id" IS '订单ID';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."spu_name" IS '商品名称';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."sku_id" IS '商品ID';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."sku_sn" IS '商品编码';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."sku_name" IS '规格名称';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."pic_url" IS '商品图片';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."price" IS '商品单价(单位：分)';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."quantity" IS '商品数量';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."total_amount" IS '商品总价(单位：分)';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order_item"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order_item" IS '订单商品信息表';

-- ----------------------------
-- Records of oms_order_item
-- ----------------------------

-- ----------------------------
-- Table structure for oms_order_log
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order_log";
CREATE TABLE "mall-boot"."oms_order_log" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_log_id_seq'::regclass),
  "order_id" int8 NOT NULL,
  "user" varchar(100) COLLATE "pg_catalog"."default",
  "detail" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "order_status" int4,
  "remark" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order_log"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."order_id" IS '订单id';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."user" IS '操作人[用户；系统；后台管理员]';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."detail" IS '操作详情';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."order_status" IS '操作时订单状态';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."remark" IS '备注';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order_log"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order_log" IS '订单操作历史记录';

-- ----------------------------
-- Records of oms_order_log
-- ----------------------------

-- ----------------------------
-- Table structure for oms_order_pay
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order_pay";
CREATE TABLE "mall-boot"."oms_order_pay" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_pay_id_seq'::regclass),
  "order_id" int8 NOT NULL,
  "pay_sn" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "pay_amount" int8 NOT NULL DEFAULT '0'::bigint,
  "pay_time" timestamptz(6) NOT NULL,
  "pay_type" int2 NOT NULL,
  "pay_status" int2 NOT NULL,
  "confirm_time" timestamptz(6) NOT NULL,
  "callback_content" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "callback_time" timestamptz(6) NOT NULL,
  "pay_subject" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."order_id" IS '订单id';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_sn" IS '支付流水号';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_amount" IS '应付总额(分)';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_time" IS '支付时间';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_type" IS '支付方式【1->支付宝；2->微信；3->银联； 4->货到付款；】';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_status" IS '支付状态';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."confirm_time" IS '确认时间';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."callback_content" IS '回调内容';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."callback_time" IS '回调时间';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."pay_subject" IS '交易内容';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order_pay"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order_pay" IS '支付信息表';

-- ----------------------------
-- Records of oms_order_pay
-- ----------------------------

-- ----------------------------
-- Table structure for oms_order_setting
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."oms_order_setting";
CREATE TABLE "mall-boot"."oms_order_setting" (
  "id" int8 NOT NULL DEFAULT nextval('oms_order_setting_id_seq'::regclass),
  "flash_order_overtime" int4 NOT NULL,
  "normal_order_overtime" int4 NOT NULL,
  "confirm_overtime" int4 NOT NULL,
  "finish_overtime" int4 NOT NULL,
  "comment_overtime" int4 NOT NULL,
  "member_level" int2 NOT NULL,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."flash_order_overtime" IS '秒杀订单超时关闭时间(分)';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."normal_order_overtime" IS '正常订单超时时间(分)';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."confirm_overtime" IS '发货后自动确认收货时间（天）';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."finish_overtime" IS '自动完成交易时间，不能申请退货（天）';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."comment_overtime" IS '订单完成后自动好评时间（天）';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."member_level" IS '会员等级【0-不限会员等级，全部通用；其他-对应的其他会员等级】';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."oms_order_setting"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."oms_order_setting" IS '订单配置信息';

-- ----------------------------
-- Records of oms_order_setting
-- ----------------------------

-- ----------------------------
-- Table structure for pms_brand
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_brand";
CREATE TABLE "mall-boot"."pms_brand" (
  "id" int8 NOT NULL DEFAULT nextval('pms_brand_id_seq'::regclass),
  "shop_id" int8 NOT NULL DEFAULT '0'::bigint,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "logo" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "sort" int4 NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_brand"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."pms_brand"."shop_id" IS 'shopID';
COMMENT ON COLUMN "mall-boot"."pms_brand"."name" IS '品牌名称';
COMMENT ON COLUMN "mall-boot"."pms_brand"."logo" IS 'LOGO图片';
COMMENT ON COLUMN "mall-boot"."pms_brand"."sort" IS '排序';
COMMENT ON COLUMN "mall-boot"."pms_brand"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."pms_brand"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."pms_brand" IS '商品品牌表';

-- ----------------------------
-- Records of pms_brand
-- ----------------------------
INSERT INTO "mall-boot"."pms_brand" VALUES (1, 1, '有来', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, '2021-07-11 19:56:58+08', '2021-07-11 20:02:54+08', 0);
INSERT INTO "mall-boot"."pms_brand" VALUES (10, 1, '小米', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 2, '2022-03-05 16:12:16+08', '2022-03-05 16:12:16+08', 0);
INSERT INTO "mall-boot"."pms_brand" VALUES (11, 1, '华硕', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 3, '2022-03-05 16:12:16+08', '2022-03-05 16:12:16+08', 0);
INSERT INTO "mall-boot"."pms_brand" VALUES (20, 1, '华为', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, '2022-05-06 23:08:33+08', '2022-05-06 23:08:33+08', 0);
INSERT INTO "mall-boot"."pms_brand" VALUES (33, 1, '惠普', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, '2022-07-07 00:12:16+08', '2022-07-07 00:12:16+08', 1699431676);
INSERT INTO "mall-boot"."pms_brand" VALUES (34, 1, '测试', 'https://img.imguo.com/files/2023/11/08/15a7e93b-1287-4538-b450-61976b3379d1.png', 2, '2023-11-08 16:24:11+08', '2023-11-08 16:38:33+08', 0);
INSERT INTO "mall-boot"."pms_brand" VALUES (35, 2, '康师傅', 'https://img.imguo.com/files/2023/11/20/f11e1d79-3a70-430e-b1d8-565c69744bf7.png', 1, '2023-11-20 13:04:48+08', '2023-11-20 13:06:35+08', 0);

-- ----------------------------
-- Table structure for pms_category
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_category";
CREATE TABLE "mall-boot"."pms_category" (
  "id" int8 NOT NULL DEFAULT nextval('pms_category_id_seq'::regclass),
  "shop_id" int8 NOT NULL DEFAULT '0'::bigint,
  "parent_id" int8 NOT NULL DEFAULT '0'::bigint,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "icon" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "sort" int4 NOT NULL DEFAULT 1,
  "visible" bool NOT NULL DEFAULT true,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_category"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."pms_category"."shop_id" IS 'shop_id';
COMMENT ON COLUMN "mall-boot"."pms_category"."parent_id" IS '父级ID';
COMMENT ON COLUMN "mall-boot"."pms_category"."name" IS '商品分类名称';
COMMENT ON COLUMN "mall-boot"."pms_category"."icon" IS '图标地址';
COMMENT ON COLUMN "mall-boot"."pms_category"."sort" IS '排序';
COMMENT ON COLUMN "mall-boot"."pms_category"."visible" IS '显示状态:( 0:隐藏 1:显示)';
COMMENT ON COLUMN "mall-boot"."pms_category"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."pms_category"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."pms_category" IS '商品分类表';

-- ----------------------------
-- Records of pms_category
-- ----------------------------
INSERT INTO "mall-boot"."pms_category" VALUES (3, 1, 0, '手机配件', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 2, 't', '2022-02-25 11:22:44+08', '2023-10-27 16:45:02+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (4, 1, 3, '智能手机', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, 't', '2022-02-25 11:22:44+08', '2023-10-27 16:45:03+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (5, 1, 4, '5g手机', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, 't', '2022-02-25 11:22:44+08', '2023-10-27 16:45:04+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (6, 1, 0, '电脑办公', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 1, 't', '2022-02-25 11:22:44+08', '2023-10-27 16:45:05+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (97, 1, 6, '笔记本电脑', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 100, 't', '2022-07-08 00:10:27+08', '2023-10-27 16:45:05+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (99, 1, 97, '轻薄本', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 100, 't', '2022-07-08 00:14:03+08', '2023-10-27 16:45:06+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (100, 1, 97, '全能本', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 100, 't', '2022-07-08 00:14:10+08', '2023-10-27 16:45:07+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (101, 1, 97, '游戏本', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', 100, 'f', '2022-07-08 00:14:18+08', '2023-11-10 09:54:00+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (102, 2, 0, '零食', 'https://img.imguo.com/files/2023/11/20/6b975f80-950a-4bdf-856d-026b4e07c473.png', 1, 't', '2023-11-20 13:05:24+08', '2023-11-20 13:05:24+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (103, 2, 102, '饼干', 'https://img.imguo.com/files/2023/11/20/76c2f013-3b41-4abe-bd67-625cce86904b.png', 1, 't', '2023-11-20 13:06:18+08', '2023-11-20 13:06:18+08', 0);
INSERT INTO "mall-boot"."pms_category" VALUES (104, 2, 103, '康师傅', 'https://img.imguo.com/files/2023/12/27/bd8c2647-62bc-4cc3-9eb8-2231b5756224.png', 1, 't', '2023-12-27 11:52:53+08', '2023-12-27 11:52:53+08', 0);

-- ----------------------------
-- Table structure for pms_category_attribute
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_category_attribute";
CREATE TABLE "mall-boot"."pms_category_attribute" (
  "id" int8 NOT NULL DEFAULT nextval('pms_category_attribute_id_seq'::regclass),
  "category_id" int8 NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."category_id" IS '分类ID';
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."name" IS '属性名称';
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."type" IS '类型(1:规格;2:属性;)';
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."pms_category_attribute"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."pms_category_attribute" IS '商品属性表';

-- ----------------------------
-- Records of pms_category_attribute
-- ----------------------------
INSERT INTO "mall-boot"."pms_category_attribute" VALUES (34, 5, '颜色', 1, '2021-07-11 17:57:06+08', '2022-07-01 00:08:19+08');
INSERT INTO "mall-boot"."pms_category_attribute" VALUES (35, 5, '规格', 1, '2021-07-11 18:00:06+08', '2022-07-01 00:08:19+08');
INSERT INTO "mall-boot"."pms_category_attribute" VALUES (36, 5, '上市时间', 2, '2021-07-11 18:00:08+08', '2022-06-01 17:41:05+08');

-- ----------------------------
-- Table structure for pms_goods
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_goods";
CREATE TABLE "mall-boot"."pms_goods" (
  "id" int8 NOT NULL DEFAULT nextval('pms_goods_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "shop_id" int4 NOT NULL DEFAULT 1,
  "category_id" int8 NOT NULL,
  "brand_id" int8 NOT NULL,
  "origin_price" int8 NOT NULL,
  "price" int8 NOT NULL,
  "sales" int4 NOT NULL DEFAULT 0,
  "pic_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "album" varchar(5120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "unit" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "description" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "detail" text COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_goods"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."pms_goods"."name" IS '商品名称';
COMMENT ON COLUMN "mall-boot"."pms_goods"."shop_id" IS 'shopID';
COMMENT ON COLUMN "mall-boot"."pms_goods"."category_id" IS '商品类型ID';
COMMENT ON COLUMN "mall-boot"."pms_goods"."brand_id" IS '商品品牌ID';
COMMENT ON COLUMN "mall-boot"."pms_goods"."origin_price" IS '原价【起】';
COMMENT ON COLUMN "mall-boot"."pms_goods"."price" IS '现价【起】';
COMMENT ON COLUMN "mall-boot"."pms_goods"."sales" IS '销量';
COMMENT ON COLUMN "mall-boot"."pms_goods"."pic_url" IS '商品主图';
COMMENT ON COLUMN "mall-boot"."pms_goods"."album" IS '商品图册';
COMMENT ON COLUMN "mall-boot"."pms_goods"."unit" IS '单位';
COMMENT ON COLUMN "mall-boot"."pms_goods"."description" IS '商品简介';
COMMENT ON COLUMN "mall-boot"."pms_goods"."detail" IS '商品详情';
COMMENT ON COLUMN "mall-boot"."pms_goods"."status" IS '商品状态(0:下架 1:上架)';
COMMENT ON COLUMN "mall-boot"."pms_goods"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."pms_goods"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."pms_goods" IS '商品表';

-- ----------------------------
-- Records of pms_goods
-- ----------------------------
INSERT INTO "mall-boot"."pms_goods" VALUES (1, '测试', 1, 5, 1, 100, 200, 0, 'https://img.imguo.com/files/2023/11/09/96b96503-3e5a-4df8-a402-94308c028106.png', 'https://img.imguo.com/files/2023/11/09/692e12e1-2718-47e1-a42a-487f83d5c926.png', '个', '测试', '<p>彻思叔叔</p>', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);

-- ----------------------------
-- Table structure for pms_goods_attribute
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_goods_attribute";
CREATE TABLE "mall-boot"."pms_goods_attribute" (
  "id" int8 NOT NULL DEFAULT nextval('pms_goods_attribute_id_seq'::regclass),
  "goods_id" int8 NOT NULL,
  "attribute_id" int8 NOT NULL DEFAULT '0'::bigint,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "value" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "pic_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."goods_id" IS '产品ID';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."attribute_id" IS '属性ID';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."name" IS '属性名称';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."value" IS '属性值';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."type" IS '类型(1:规格;2:属性;)';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."pic_url" IS '规格图片';
COMMENT ON COLUMN "mall-boot"."pms_goods_attribute"."status" IS '1';
COMMENT ON TABLE "mall-boot"."pms_goods_attribute" IS '商品属性/规格表';

-- ----------------------------
-- Records of pms_goods_attribute
-- ----------------------------
INSERT INTO "mall-boot"."pms_goods_attribute" VALUES (1, 1, 0, '品牌', '芒果', 2, '', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);
INSERT INTO "mall-boot"."pms_goods_attribute" VALUES (2, 1, 0, '颜色', '红', 1, 'https://img.imguo.com/files/2023/11/11/7fea8284-5d15-4621-a2c5-21c9c10e727e.png', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);
INSERT INTO "mall-boot"."pms_goods_attribute" VALUES (3, 1, 0, '颜色', '黄', 1, '', 0, '2023-11-09 16:41:46+08', '2023-11-11 23:39:56+08', 0);
INSERT INTO "mall-boot"."pms_goods_attribute" VALUES (4, 1, 0, '尺寸', 'x', 1, '', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);
INSERT INTO "mall-boot"."pms_goods_attribute" VALUES (5, 1, 0, '尺寸', 'm', 1, '', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);

-- ----------------------------
-- Table structure for pms_shop
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_shop";
CREATE TABLE "mall-boot"."pms_shop" (
  "id" int8 NOT NULL DEFAULT nextval('pms_shop_id_seq'::regclass),
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "tel" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "notice" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "lng" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "lat" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "address" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "distribution_fee" numeric(10,2) NOT NULL DEFAULT 0.00,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_shop"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."pms_shop"."name" IS '店铺名称';
COMMENT ON COLUMN "mall-boot"."pms_shop"."tel" IS '店铺电话';
COMMENT ON COLUMN "mall-boot"."pms_shop"."notice" IS '公告';
COMMENT ON COLUMN "mall-boot"."pms_shop"."distribution_fee" IS '配送费';
COMMENT ON TABLE "mall-boot"."pms_shop" IS '店铺';

-- ----------------------------
-- Records of pms_shop
-- ----------------------------
INSERT INTO "mall-boot"."pms_shop" VALUES (1, '有个商城', '18600008888', '黑店不黑', '', '', NULL, 1.00, 1, '2023-10-25 20:38:36+08', '2023-10-26 11:37:25+08', 0);
INSERT INTO "mall-boot"."pms_shop" VALUES (2, '有个商城2', '18600008888', '黑店不黑', '', '', NULL, 1.00, 1, '2023-10-25 20:38:36+08', '2023-10-26 11:37:25+08', 0);

-- ----------------------------
-- Table structure for pms_sku
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."pms_sku";
CREATE TABLE "mall-boot"."pms_sku" (
  "id" int8 NOT NULL DEFAULT nextval('pms_sku_id_seq'::regclass),
  "goods_id" int8 NOT NULL,
  "sku_sn" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "spec_ids" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "price" int8 NOT NULL,
  "stock" int4 NOT NULL DEFAULT 0,
  "locked_stock" int4 NOT NULL DEFAULT 0,
  "pic_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."pms_sku"."goods_id" IS 'goods ID';
COMMENT ON COLUMN "mall-boot"."pms_sku"."sku_sn" IS '商品编码';
COMMENT ON COLUMN "mall-boot"."pms_sku"."name" IS '商品名称';
COMMENT ON COLUMN "mall-boot"."pms_sku"."spec_ids" IS '商品规格值，以英文逗号(,)分割';
COMMENT ON COLUMN "mall-boot"."pms_sku"."price" IS '商品价格(单位：分)';
COMMENT ON COLUMN "mall-boot"."pms_sku"."stock" IS '库存数量';
COMMENT ON COLUMN "mall-boot"."pms_sku"."locked_stock" IS '库存锁定数量';
COMMENT ON COLUMN "mall-boot"."pms_sku"."pic_url" IS '商品图片地址';
COMMENT ON COLUMN "mall-boot"."pms_sku"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."pms_sku"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."pms_sku" IS '商品库存表';

-- ----------------------------
-- Records of pms_sku
-- ----------------------------
INSERT INTO "mall-boot"."pms_sku" VALUES (1, 1, 'sn1001', '红 x', '2_4', 1100, 99, 0, 'https://img.imguo.com/files/2023/11/11/7fea8284-5d15-4621-a2c5-21c9c10e727e.png', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);
INSERT INTO "mall-boot"."pms_sku" VALUES (2, 1, 'sn1002', '红 m', '2_5', 1150, 999, 0, 'https://img.imguo.com/files/2023/11/11/7fea8284-5d15-4621-a2c5-21c9c10e727e.png', 1, '2023-11-09 16:41:46+08', '2023-12-25 22:30:46+08', 0);
INSERT INTO "mall-boot"."pms_sku" VALUES (3, 1, 'sn103', '黄 x', '3_4', 1200, 111, 0, '', 0, '2023-11-09 16:41:46+08', '2023-11-11 23:39:56+08', 0);
INSERT INTO "mall-boot"."pms_sku" VALUES (4, 1, 'sn104', '黄 m', '3_5', 1200, 333, 0, '', 0, '2023-11-09 16:41:46+08', '2023-11-11 23:39:56+08', 0);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_dept";
CREATE TABLE "mall-boot"."sys_dept" (
  "id" int8 NOT NULL DEFAULT nextval('sys_dept_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "parent_id" int8 NOT NULL DEFAULT '0'::bigint,
  "tree_path" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sort" int4 NOT NULL DEFAULT 0,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "create_by" int8 NOT NULL DEFAULT '0'::bigint,
  "update_by" int8 NOT NULL DEFAULT '0'::bigint,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."sys_dept"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."sys_dept"."name" IS '部门名称';
COMMENT ON COLUMN "mall-boot"."sys_dept"."code" IS '部门名称';
COMMENT ON COLUMN "mall-boot"."sys_dept"."parent_id" IS '父节点id';
COMMENT ON COLUMN "mall-boot"."sys_dept"."tree_path" IS '父节点id路径';
COMMENT ON COLUMN "mall-boot"."sys_dept"."sort" IS '显示顺序';
COMMENT ON COLUMN "mall-boot"."sys_dept"."status" IS '状态(1:正常;0:禁用)';
COMMENT ON COLUMN "mall-boot"."sys_dept"."create_by" IS '创建人ID';
COMMENT ON COLUMN "mall-boot"."sys_dept"."update_by" IS '修改人ID';
COMMENT ON COLUMN "mall-boot"."sys_dept"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_dept"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_dept"."deleted_at" IS '逻辑删除标识(1:已删除;0:未删除)';
COMMENT ON TABLE "mall-boot"."sys_dept" IS '部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO "mall-boot"."sys_dept" VALUES (1, '普鲁顿科技', '', 0, '0', 1, 1, 1, 1, '2022-04-19 12:46:37+08', '2022-04-19 12:46:37+08', 0);
INSERT INTO "mall-boot"."sys_dept" VALUES (2, '研发部门', '', 1, '0,1', 1, 1, 2, 2, '2022-04-19 12:46:37+08', '2022-04-19 12:46:37+08', 0);
INSERT INTO "mall-boot"."sys_dept" VALUES (3, '测试部门', 'test', 1, '0,1', 1, 1, 2, 2, '2022-04-19 12:46:37+08', '2024-10-26 21:17:17+08', 0);
INSERT INTO "mall-boot"."sys_dept" VALUES (4, '销售部门', '', 1, '0,1', 1, 1, 1, 1, '2022-04-19 12:46:37+08', '2022-04-19 12:46:37+08', 0);
INSERT INTO "mall-boot"."sys_dept" VALUES (172, '总经办', '', 1, '0,1', 1, 1, 2, 2, '2023-10-23 11:30:05+08', '2023-10-23 11:30:05+08', 0);
INSERT INTO "mall-boot"."sys_dept" VALUES (173, '事业一部', '', 1, '0,1', 1, 1, 2, 2, '2023-10-23 11:31:57+08', '2023-10-23 11:31:57+08', 0);

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_dict";
CREATE TABLE "mall-boot"."sys_dict" (
  "id" int8 NOT NULL DEFAULT nextval('sys_dict_id_seq'::regclass),
  "code" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "status" int2 NOT NULL DEFAULT '1'::smallint,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."sys_dict"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."sys_dict"."code" IS '字典类型编码';
COMMENT ON COLUMN "mall-boot"."sys_dict"."name" IS '字典项名称';
COMMENT ON COLUMN "mall-boot"."sys_dict"."status" IS '状态(1:正常;0:禁用)';
COMMENT ON COLUMN "mall-boot"."sys_dict"."remark" IS '备注';
COMMENT ON COLUMN "mall-boot"."sys_dict"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_dict"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_dict"."deleted_at" IS '是否默认(1:是;0:否)';
COMMENT ON TABLE "mall-boot"."sys_dict" IS '字典数据表';

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO "mall-boot"."sys_dict" VALUES (1, 'gender', '性别', 1, '性别', '2024-10-22 11:02:58+08', '2024-10-27 12:11:10+08', 0);
INSERT INTO "mall-boot"."sys_dict" VALUES (2, 'notice_type', '通知类型', 1, '', '2024-10-22 11:02:58+08', '2024-10-22 11:02:58+08', 0);
INSERT INTO "mall-boot"."sys_dict" VALUES (3, 'notice_level', '通知级别', 1, '', '2024-10-22 11:02:58+08', '2024-10-22 11:02:58+08', 0);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_dict_data";
CREATE TABLE "mall-boot"."sys_dict_data" (
  "id" int8 NOT NULL DEFAULT nextval('sys_dict_data_id_seq'::regclass),
  "dict_id" int4 NOT NULL,
  "value" varchar(50) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "label" varchar(50) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "tag" varchar(50) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'success'::character varying,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sort" int4 NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6),
  "status" int2 DEFAULT 1
)
;
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."id" IS '主键 ';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."value" IS '类型名称';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."label" IS '类型名称';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."tag" IS '类型名称';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."remark" IS '备注';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_dict_data"."updated_at" IS '创建时间';
COMMENT ON TABLE "mall-boot"."sys_dict_data" IS '字典类型表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO "mall-boot"."sys_dict_data" VALUES (1, 1, '1', '男', 'primary', '', 1, '2024-10-22 11:02:58+08', '2024-10-27 13:37:07+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (2, 1, '2', '女', 'danger', '', 2, '2024-10-22 11:02:58+08', '2024-10-27 13:37:10+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (3, 1, '0', '保密', 'info', '', 3, '2024-10-22 11:02:58+08', '2024-10-27 13:37:12+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (4, 2, '1', '系统升级', 'success', '', 1, '2024-10-22 11:02:58+08', '2024-10-22 14:43:54+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (5, 2, '2', '系统维护', 'primary', '', 2, '2024-10-22 11:02:58+08', '2024-10-22 14:43:54+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (6, 2, '3', '安全警告', 'danger', '', 3, '2024-10-22 11:02:58+08', '2024-10-22 14:43:55+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (7, 2, '4', '假期通知', 'success', '', 4, '2024-10-22 11:02:58+08', '2024-10-22 14:43:56+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (8, 2, '5', '公司新闻', 'primary', '', 5, '2024-10-22 11:02:58+08', '2024-10-22 14:43:57+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (9, 2, '99', '其他', 'info', '', 99, '2024-10-22 11:02:58+08', '2024-10-22 14:43:58+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (10, 3, 'L', '低', 'info', '', 1, '2024-10-22 11:02:58+08', '2024-10-22 14:43:59+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (11, 3, 'M', '中', 'warning', '', 2, '2024-10-22 11:02:58+08', '2024-10-22 14:44:00+08', 1);
INSERT INTO "mall-boot"."sys_dict_data" VALUES (12, 3, 'H', '高', 'danger', '', 3, '2024-10-22 11:02:58+08', '2024-10-22 14:44:00+08', 1);

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_log";
CREATE TABLE "mall-boot"."sys_log" (
  "id" int8 NOT NULL DEFAULT nextval('sys_log_id_seq'::regclass),
  "module" "mall-boot"."sys_log_module" NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "request_uri" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ip" varchar(45) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "province" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "city" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "execution_time" int8 NOT NULL DEFAULT '0'::bigint,
  "browser" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "browser_version" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "os" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_by" int8 NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."sys_log"."id" IS '主键';
COMMENT ON COLUMN "mall-boot"."sys_log"."module" IS '日志模块';
COMMENT ON COLUMN "mall-boot"."sys_log"."content" IS '日志内容';
COMMENT ON COLUMN "mall-boot"."sys_log"."request_uri" IS '请求路径';
COMMENT ON COLUMN "mall-boot"."sys_log"."ip" IS 'IP地址';
COMMENT ON COLUMN "mall-boot"."sys_log"."province" IS '省份';
COMMENT ON COLUMN "mall-boot"."sys_log"."city" IS '城市';
COMMENT ON COLUMN "mall-boot"."sys_log"."execution_time" IS '执行时间(ms)';
COMMENT ON COLUMN "mall-boot"."sys_log"."browser" IS '浏览器';
COMMENT ON COLUMN "mall-boot"."sys_log"."browser_version" IS '浏览器版本';
COMMENT ON COLUMN "mall-boot"."sys_log"."os" IS '终端系统';
COMMENT ON COLUMN "mall-boot"."sys_log"."create_by" IS '创建人ID';
COMMENT ON COLUMN "mall-boot"."sys_log"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_log"."deleted_at" IS '逻辑删除标识(1-已删除 0-未删除)';
COMMENT ON TABLE "mall-boot"."sys_log" IS '系统日志表';

-- ----------------------------
-- Records of sys_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_menu";
CREATE TABLE "mall-boot"."sys_menu" (
  "id" int8 NOT NULL DEFAULT nextval('sys_menu_id_seq'::regclass),
  "parent_id" int8 NOT NULL,
  "tree_path" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "route_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "route_path" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "component" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "perm" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "always_show" bool NOT NULL DEFAULT false,
  "keep_alive" bool NOT NULL DEFAULT false,
  "visible" bool NOT NULL DEFAULT true,
  "sort" int4 NOT NULL DEFAULT 0,
  "icon" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "redirect" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint,
  "params" json,
  "type" int2 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "mall-boot"."sys_menu"."parent_id" IS '父菜单ID';
COMMENT ON COLUMN "mall-boot"."sys_menu"."tree_path" IS '父节点ID路径';
COMMENT ON COLUMN "mall-boot"."sys_menu"."name" IS '菜单名称';
COMMENT ON COLUMN "mall-boot"."sys_menu"."route_name" IS '路由路径(浏览器地址栏路径)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."route_path" IS '路由路径(浏览器地址栏路径)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."component" IS '组件路径(vue页面完整路径，省略.vue后缀)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."perm" IS '权限标识';
COMMENT ON COLUMN "mall-boot"."sys_menu"."always_show" IS '显示状态(1-显示;0-隐藏)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."keep_alive" IS '显示状态(1-显示;0-隐藏)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."visible" IS '显示状态(1-显示;0-隐藏)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."sort" IS '排序';
COMMENT ON COLUMN "mall-boot"."sys_menu"."icon" IS '菜单图标';
COMMENT ON COLUMN "mall-boot"."sys_menu"."redirect" IS '跳转路径';
COMMENT ON COLUMN "mall-boot"."sys_menu"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_menu"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_menu"."params" IS '组件路径(vue页面完整路径，省略.vue后缀)';
COMMENT ON COLUMN "mall-boot"."sys_menu"."type" IS '菜单类型(1:菜单；2:目录；3:外链；4:按钮)';
COMMENT ON TABLE "mall-boot"."sys_menu" IS '菜单管理';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO "mall-boot"."sys_menu" VALUES (2, 1, '0,1', '用户管理', '', 'user', 'system/user/index', '', 'f', 'f', 't', 1, 'user', '', '2021-08-28 09:12:21+08', '2021-08-28 09:12:21+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (4, 1, '0,1', '菜单管理', '', 'menu', 'system/menu/index', '', 'f', 'f', 't', 3, 'menu', '', '2021-08-28 09:12:21+08', '2021-08-28 09:12:21+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (5, 1, '0,1', '部门管理', '', 'dept', 'system/dept/index', '', 'f', 'f', 't', 4, 'tree', '', '2021-08-28 09:12:21+08', '2021-08-28 09:12:21+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (6, 1, '0,1', '字典管理', '', 'dict', 'system/dict/index', '', 'f', 'f', 't', 5, 'dict', '', '2021-08-28 09:12:21+08', '2021-08-28 09:12:21+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (104, 7, '0,7', '会员列表', '', 'user', 'ums/user/index', '', 'f', 'f', 't', 1, 'monitor', '', '2023-10-23 14:15:32+08', '2023-10-23 14:15:48+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (105, 7, '0,7', '第三方列表', '', 'third', 'ums/third/index', '', 'f', 'f', 't', 1, 'monitor', '', '2023-10-23 14:15:32+08', '2023-10-23 14:15:48+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (106, 8, '0,8', '店铺管理', '', 'shop', 'pms/shop/index', '', 'f', 'f', 't', 1, 'message', '', '2023-10-23 14:15:32+08', '2023-11-08 15:16:03+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (107, 8, '0,8', '商品管理', '', 'goods', 'pms/goods/index', '', 'f', 'f', 't', 1, 'monitor', '', '2023-10-23 14:15:32+08', '2023-10-23 14:15:48+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (108, 9, '0,9', '订单列表', '', 'order', 'oms/order/index', '', 'f', 'f', 't', 1, 'peoples', '', '2023-10-23 14:15:32+08', '2023-10-23 14:15:32+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (109, 9, '0,9', '订单设置', '', 'setting', 'oms/setting/index', '', 'f', 'f', 't', 1, 'peoples', '', '2023-10-23 14:15:32+08', '2023-11-10 21:22:16+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (110, 8, '0,8', '添加商品', '', 'goods/detail', 'pms/goods/detail', '', 'f', 'f', 't', 1, 'monitor', '', '2023-10-23 14:15:32+08', '2023-10-23 14:15:48+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (124, 1, '0,1', '字典数据', 'DictData', 'dict-data', 'system/dict/data', '', 't', 'f', 'f', 6, '', '', '2024-10-22 11:02:59+08', '2024-10-27 13:20:25+08', 0, 'null', 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (73, 4, '0,1,4', '菜单新增', '', '', '', 'sys:menu:add', 'f', 'f', 't', 1, '', '', '2023-05-20 23:41:35+08', '2023-05-20 23:41:35+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (74, 4, '0,1,4', '菜单编辑', '', '', '', 'sys:menu:edit', 'f', 'f', 't', 3, '', '', '2023-05-20 23:41:58+08', '2023-05-20 23:41:58+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (1, 0, '0', '系统管理', '', '/system', 'Layout', '', 'f', 'f', 't', 99, 'system', '/system/user', '2021-08-28 09:12:21+08', '2023-10-29 00:21:37+08', 0, NULL, 2);
INSERT INTO "mall-boot"."sys_menu" VALUES (3, 1, '0,1', '角色管理', '', 'role', 'system/role/index', '', 'f', 'f', 't', 2, 'role', '', '2021-08-28 09:12:21+08', '2021-08-28 09:12:21+08', 0, NULL, 1);
INSERT INTO "mall-boot"."sys_menu" VALUES (7, 0, '0', '会员管理', '', '/ums', 'Layout', '', 'f', 'f', 't', 30, 'homepage', '/ums/user', '2023-10-23 14:13:17+08', '2023-10-29 00:22:41+08', 0, NULL, 2);
INSERT INTO "mall-boot"."sys_menu" VALUES (8, 0, '0', '商城管理', '', '/pms', 'Layout', '', 'f', 'f', 't', 1, 'tree', '/pms/user', '2023-10-23 14:13:17+08', '2023-11-08 15:16:23+08', 0, NULL, 2);
INSERT INTO "mall-boot"."sys_menu" VALUES (9, 0, '0', '订单管理', '', '/oms', 'Layout', '', 'f', 'f', 't', 20, 'publish', '/oms/order', '2023-10-23 14:13:17+08', '2023-11-08 15:39:00+08', 0, NULL, 2);
INSERT INTO "mall-boot"."sys_menu" VALUES (31, 2, '0,1,2', '用户新增', '', '', '', 'sys:user:add', 'f', 'f', 't', 1, '', '', '2022-10-23 11:04:08+08', '2022-10-23 11:04:11+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (32, 2, '0,1,2', '用户编辑', '', '', '', 'sys:user:edit', 'f', 'f', 't', 2, '', '', '2022-10-23 11:04:08+08', '2022-10-23 11:04:11+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (33, 2, '0,1,2', '用户删除', '', '', '', 'sys:user:delete', 'f', 'f', 't', 3, '', '', '2022-10-23 11:04:08+08', '2022-10-23 11:04:11+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (70, 3, '0,1,3', '角色新增', '', '', '', 'sys:role:add', 'f', 'f', 't', 1, '', '', '2023-05-20 23:39:09+08', '2023-05-20 23:39:09+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (71, 3, '0,1,3', '角色编辑', '', '', '', 'sys:role:edit', 'f', 'f', 't', 2, '', '', '2023-05-20 23:40:31+08', '2023-05-20 23:40:31+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (72, 3, '0,1,3', '角色删除', '', '', '', 'sys:role:delete', 'f', 'f', 't', 3, '', '', '2023-05-20 23:41:08+08', '2023-05-20 23:41:08+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (75, 4, '0,1,4', '菜单删除', '', '', '', 'sys:menu:delete', 'f', 'f', 't', 3, '', '', '2023-05-20 23:44:18+08', '2023-05-20 23:44:18+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (76, 5, '0,1,5', '部门新增', '', '', '', 'sys:dept:add', 'f', 'f', 't', 1, '', '', '2023-05-20 23:45:00+08', '2023-05-20 23:45:00+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (77, 5, '0,1,5', '部门编辑', '', '', '', 'sys:dept:edit', 'f', 'f', 't', 2, '', '', '2023-05-20 23:46:16+08', '2023-05-20 23:46:16+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (78, 5, '0,1,5', '部门删除', '', '', '', 'sys:dept:delete', 'f', 'f', 't', 3, '', '', '2023-05-20 23:46:36+08', '2023-05-20 23:46:36+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (79, 6, '0,1,6', '字典类型新增', '', '', '', 'sys:dict_type:add', 'f', 'f', 't', 1, '', '', '2023-05-21 00:16:06+08', '2023-05-21 00:16:06+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (81, 6, '0,1,6', '字典类型编辑', '', '', '', 'sys:dict_type:edit', 'f', 'f', 't', 2, '', '', '2023-05-21 00:27:37+08', '2023-05-21 00:27:37+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (84, 6, '0,1,6', '字典类型删除', '', '', '', 'sys:dict_type:delete', 'f', 'f', 't', 3, '', '', '2023-05-21 00:29:39+08', '2023-05-21 00:29:39+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (85, 6, '0,1,6', '字典数据新增', '', '', '', 'sys:dict:add', 'f', 'f', 't', 4, '', '', '2023-05-21 00:46:56+08', '2023-05-21 00:47:06+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (86, 6, '0,1,6', '字典数据编辑', '', '', '', 'sys:dict:edit', 'f', 'f', 't', 5, '', '', '2023-05-21 00:47:36+08', '2023-05-21 00:47:36+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (87, 6, '0,1,6', '字典数据删除', '', '', '', 'sys:dict:delete', 'f', 'f', 't', 6, '', '', '2023-05-21 00:48:10+08', '2023-05-21 00:48:20+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (88, 2, '0,1,2', '重置密码', '', '', '', 'sys:user:reset_pwd', 'f', 'f', 't', 4, '', '', '2023-05-21 00:49:18+08', '2023-05-21 00:49:18+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (111, 9, '0,9', 'ws列表', 'websocket', 'https://www.google.com', 'Layout', '', 'f', 'f', 't', 1, 'peoples', 'https://www.baidu.com', '2023-10-23 14:15:32+08', '2024-10-26 20:57:12+08', 0, '[{"key": "xx", "value": "11"}, {"key": "yy", "value": "22"}]', 3);
INSERT INTO "mall-boot"."sys_menu" VALUES (125, 124, '0,1,124', '字典数据新增', '', '', '', 'sys:dict-data:add', 't', 'f', 't', 4, '', '', '2024-10-22 11:02:59+08', '2024-10-22 11:02:59+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (126, 124, '0,1,124', '字典数据编辑', '', '', '', 'sys:dict-data:edit', 't', 'f', 't', 5, '', '', '2024-10-22 11:02:59+08', '2024-10-22 11:02:59+08', 0, NULL, 4);
INSERT INTO "mall-boot"."sys_menu" VALUES (127, 124, '0,1,124', '字典数据删除', '', '', '', 'sys:dict-data:delete', 't', 'f', 't', 6, '', '', '2024-10-22 11:02:59+08', '2024-10-22 11:02:59+08', 0, NULL, 4);

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_notice";
CREATE TABLE "mall-boot"."sys_notice" (
  "id" int8 NOT NULL DEFAULT nextval('sys_notice_id_seq'::regclass),
  "title" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "level" varchar(5) COLLATE "pg_catalog"."default" NOT NULL,
  "target_type" int2 NOT NULL,
  "target_user_ids" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "publisher_id" int8 NOT NULL,
  "publish_status" int2 NOT NULL DEFAULT '0'::smallint,
  "publish_time" timestamptz(6) NOT NULL,
  "revoke_time" timestamptz(6) NOT NULL,
  "create_by" int8 NOT NULL,
  "update_by" int8 NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."sys_notice"."title" IS '通知标题';
COMMENT ON COLUMN "mall-boot"."sys_notice"."content" IS '通知内容';
COMMENT ON COLUMN "mall-boot"."sys_notice"."type" IS '通知类型（关联字典编码：notice_type）';
COMMENT ON COLUMN "mall-boot"."sys_notice"."level" IS '通知等级（字典code：notice_level）';
COMMENT ON COLUMN "mall-boot"."sys_notice"."target_type" IS '目标类型（1: 全体, 2: 指定）';
COMMENT ON COLUMN "mall-boot"."sys_notice"."target_user_ids" IS '目标人ID集合（多个使用英文逗号,分割）';
COMMENT ON COLUMN "mall-boot"."sys_notice"."publisher_id" IS '发布人ID';
COMMENT ON COLUMN "mall-boot"."sys_notice"."publish_status" IS '发布状态（0: 未发布, 1: 已发布, -1: 已撤回）';
COMMENT ON COLUMN "mall-boot"."sys_notice"."publish_time" IS '发布时间';
COMMENT ON COLUMN "mall-boot"."sys_notice"."revoke_time" IS '撤回时间';
COMMENT ON COLUMN "mall-boot"."sys_notice"."create_by" IS '创建人ID';
COMMENT ON COLUMN "mall-boot"."sys_notice"."update_by" IS '更新人ID';
COMMENT ON COLUMN "mall-boot"."sys_notice"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_notice"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_notice"."deleted_at" IS '是否删除（0: 未删除, 1: 已删除）';
COMMENT ON TABLE "mall-boot"."sys_notice" IS '通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_role";
CREATE TABLE "mall-boot"."sys_role" (
  "id" int8 NOT NULL DEFAULT nextval('sys_role_id_seq'::regclass),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sort" int4 NOT NULL,
  "data_scope" int2 NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint,
  "status" int2 DEFAULT 1
)
;
COMMENT ON COLUMN "mall-boot"."sys_role"."name" IS '角色名称';
COMMENT ON COLUMN "mall-boot"."sys_role"."code" IS '角色编码';
COMMENT ON COLUMN "mall-boot"."sys_role"."sort" IS '显示顺序';
COMMENT ON COLUMN "mall-boot"."sys_role"."data_scope" IS '数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)';
COMMENT ON COLUMN "mall-boot"."sys_role"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_role"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_role"."deleted_at" IS '逻辑删除标识';
COMMENT ON TABLE "mall-boot"."sys_role" IS '角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO "mall-boot"."sys_role" VALUES (1, '超级管理员', 'ROOT', 1, 0, '2018-12-23 16:00:00+08', '2021-05-21 14:56:51+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (2, '系统管理员', 'ADMIN', 2, 1, NULL, '2021-03-25 12:39:54+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (3, '访问游客', 'GUEST', 3, 0, '2019-05-05 16:00:00+08', '2023-11-08 11:14:27+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (4, '系统管理员1', 'ADMIN1', 2, 1, NULL, '2021-03-25 12:39:54+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (5, '系统管理员2', 'ADMIN2', 2, 1, NULL, '2021-03-25 12:39:54+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (6, '系统管理员3', 'ADMIN3', 2, 1, NULL, '2021-03-25 12:39:54+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (7, 'tt001', 'tt33', 1, 0, '2023-11-27 21:29:11+08', '2024-10-24 15:17:13+08', 0, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (8, 'test', 'test222', 1, 3, '2023-11-27 21:30:21+08', '2023-11-27 21:53:02+08', 1701420926, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (9, 'tr', 'tr', 1, 1, '2023-11-27 21:53:17+08', '2023-11-27 21:53:17+08', 1701420926, 1);
INSERT INTO "mall-boot"."sys_role" VALUES (10, '测试', 'test', 1, 0, '2024-10-24 15:01:53+08', '2024-10-24 15:01:53+08', 0, 1);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_role_menu";
CREATE TABLE "mall-boot"."sys_role_menu" (
  "role_id" int8 NOT NULL,
  "menu_id" int8 NOT NULL
)
;
COMMENT ON COLUMN "mall-boot"."sys_role_menu"."role_id" IS '角色ID';
COMMENT ON COLUMN "mall-boot"."sys_role_menu"."menu_id" IS '菜单ID';
COMMENT ON TABLE "mall-boot"."sys_role_menu" IS '角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 1);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 2);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 31);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 32);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 33);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 88);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 3);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 70);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 71);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (5, 72);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 1);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 4);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 73);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 74);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 75);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 5);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 76);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 77);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (4, 78);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 1);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 2);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 31);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 32);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 33);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 88);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 3);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 70);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 71);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 72);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 4);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 73);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 74);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 75);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 5);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 76);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 77);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 78);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 6);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 79);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 81);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 84);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 85);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 86);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 87);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 7);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 104);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 105);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 8);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 106);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (2, 107);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 8);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 106);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 107);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 110);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 7);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 104);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 1);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 2);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 31);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 32);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 33);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (10, 88);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 8);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 106);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 107);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 110);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 9);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 108);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 109);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 111);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 7);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 104);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 105);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 1);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 2);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 31);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 32);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 33);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 88);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 3);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 70);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 71);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 72);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 4);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 73);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 74);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 75);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 5);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 76);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 77);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 78);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 6);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 79);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 81);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 84);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 85);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 86);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 87);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 124);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 125);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 126);
INSERT INTO "mall-boot"."sys_role_menu" VALUES (1, 127);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_user";
CREATE TABLE "mall-boot"."sys_user" (
  "id" int4 NOT NULL DEFAULT nextval('sys_user_id_seq'::regclass),
  "username" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "nickname" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "dept_id" int4 NOT NULL,
  "avatar" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "mobile" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint,
  "status" int2 DEFAULT 1,
  "gender" int2 DEFAULT 1
)
;
COMMENT ON COLUMN "mall-boot"."sys_user"."username" IS '用户名';
COMMENT ON COLUMN "mall-boot"."sys_user"."nickname" IS '昵称';
COMMENT ON COLUMN "mall-boot"."sys_user"."password" IS '密码';
COMMENT ON COLUMN "mall-boot"."sys_user"."dept_id" IS '部门ID';
COMMENT ON COLUMN "mall-boot"."sys_user"."avatar" IS '用户头像';
COMMENT ON COLUMN "mall-boot"."sys_user"."mobile" IS '联系方式';
COMMENT ON COLUMN "mall-boot"."sys_user"."email" IS '用户邮箱';
COMMENT ON COLUMN "mall-boot"."sys_user"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_user"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_user"."deleted_at" IS '逻辑删除标识(0:未删除;1:已删除)';
COMMENT ON TABLE "mall-boot"."sys_user" IS '用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO "mall-boot"."sys_user" VALUES (1, 'kyle', '总监', '2be7b79b4067d1d315c0ef99b4d34f61cbba08ba7882cbb0576d8193bac332d8', 1, 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', '17621590365', 'kyle@163.com', '2019-10-10 13:41:22+08', '2024-10-24 10:43:18+08', 0, 1, 1);
INSERT INTO "mall-boot"."sys_user" VALUES (2, 'admin', '系统管理员', '2be7b79b4067d1d315c0ef99b4d34f61cbba08ba7882cbb0576d8193bac332d8', 1, 'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif', '17621210366', 'kyle@163.com', '2019-10-10 13:41:22+08', '2024-10-24 10:17:28+08', 0, 1, 1);
INSERT INTO "mall-boot"."sys_user" VALUES (3, 'test', '测试小用户', '2be7b79b4067d1d315c0ef99b4d34f61cbba08ba7882cbb0576d8193bac332d8', 3, 'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif', '17621210366', 'kyle@163.com', '2021-06-05 01:31:29+08', '2024-10-24 10:11:05+08', 0, 1, 1);
INSERT INTO "mall-boot"."sys_user" VALUES (4, 'xx', 'xxx', '48e449d9b50c6246bf0c28a61eecf3f9ee2094f42966bd13e1201bfa5e92f083', 2, '', '18627111095', 'sk@qq.com', '2024-10-24 11:46:51+08', '2024-10-24 11:46:51+08', 0, 1, 1);

-- ----------------------------
-- Table structure for sys_user_notice
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_user_notice";
CREATE TABLE "mall-boot"."sys_user_notice" (
  "id" int8 NOT NULL DEFAULT nextval('sys_user_notice_id_seq'::regclass),
  "notice_id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "is_read" int8 NOT NULL DEFAULT '0'::bigint,
  "read_time" timestamptz(6),
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6),
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."notice_id" IS '公共通知id';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."user_id" IS '用户id';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."is_read" IS '读取状态（0: 未读, 1: 已读）';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."read_time" IS '阅读时间';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."updated_at" IS '更新时间';
COMMENT ON COLUMN "mall-boot"."sys_user_notice"."deleted_at" IS '逻辑删除(0: 未删除, 1: 已删除)';
COMMENT ON TABLE "mall-boot"."sys_user_notice" IS '用户通知公告表';

-- ----------------------------
-- Records of sys_user_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."sys_user_role";
CREATE TABLE "mall-boot"."sys_user_role" (
  "user_id" int8 NOT NULL,
  "role_id" int8 NOT NULL
)
;
COMMENT ON COLUMN "mall-boot"."sys_user_role"."user_id" IS '用户ID';
COMMENT ON COLUMN "mall-boot"."sys_user_role"."role_id" IS '角色ID';
COMMENT ON TABLE "mall-boot"."sys_user_role" IS '用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO "mall-boot"."sys_user_role" VALUES (1, 1);
INSERT INTO "mall-boot"."sys_user_role" VALUES (1, 2);
INSERT INTO "mall-boot"."sys_user_role" VALUES (2, 2);
INSERT INTO "mall-boot"."sys_user_role" VALUES (3, 1);
INSERT INTO "mall-boot"."sys_user_role" VALUES (3, 2);
INSERT INTO "mall-boot"."sys_user_role" VALUES (4, 1);
INSERT INTO "mall-boot"."sys_user_role" VALUES (4, 2);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."user";
CREATE TABLE "mall-boot"."user" (
  "id" int8 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
  "mobile" char(11) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "username" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "nickname" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(512) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "id_card" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "gender" int2 NOT NULL DEFAULT '0'::smallint,
  "signature" varchar(10240) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "birthday" date,
  "tags" varchar(5000) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "source" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'WXAPP'::character varying,
  "source_uid" int4 NOT NULL DEFAULT 0,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."user"."mobile" IS '电话号码';
COMMENT ON COLUMN "mall-boot"."user"."username" IS '姓名';
COMMENT ON COLUMN "mall-boot"."user"."nickname" IS '昵称';
COMMENT ON COLUMN "mall-boot"."user"."avatar" IS '头像';
COMMENT ON COLUMN "mall-boot"."user"."id_card" IS '身份证号码';
COMMENT ON COLUMN "mall-boot"."user"."gender" IS '性别 0 未知 1男 2女';
COMMENT ON COLUMN "mall-boot"."user"."signature" IS '签名';
COMMENT ON COLUMN "mall-boot"."user"."birthday" IS '生日';
COMMENT ON COLUMN "mall-boot"."user"."tags" IS 'tags';
COMMENT ON COLUMN "mall-boot"."user"."source" IS '来源，APP H5';
COMMENT ON COLUMN "mall-boot"."user"."source_uid" IS '邀请uid';

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."user_address";
CREATE TABLE "mall-boot"."user_address" (
  "id" int8 NOT NULL DEFAULT nextval('user_address_id_seq'::regclass),
  "shop_id" int4 NOT NULL DEFAULT 0,
  "uid" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "mobile" char(11) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "area" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "info" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_default" bool NOT NULL DEFAULT false,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "mall-boot"."user_address"."id" IS 'ID';
COMMENT ON COLUMN "mall-boot"."user_address"."shop_id" IS '店铺id';
COMMENT ON COLUMN "mall-boot"."user_address"."uid" IS '用户ID';
COMMENT ON COLUMN "mall-boot"."user_address"."name" IS '联系人';
COMMENT ON COLUMN "mall-boot"."user_address"."mobile" IS '手机号';
COMMENT ON COLUMN "mall-boot"."user_address"."area" IS '学校地址';
COMMENT ON COLUMN "mall-boot"."user_address"."info" IS '详细地址';
COMMENT ON COLUMN "mall-boot"."user_address"."is_default" IS '是否默认';
COMMENT ON TABLE "mall-boot"."user_address" IS '学校-用户收货地址表';

-- ----------------------------
-- Records of user_address
-- ----------------------------
INSERT INTO "mall-boot"."user_address" VALUES (1, 1, 1, '黄先生', '18627111095', '武汉', '3501', 't', '2023-11-10 18:29:23+08', '2023-11-10 18:29:25+08', 0);

-- ----------------------------
-- Table structure for user_cart
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."user_cart";
CREATE TABLE "mall-boot"."user_cart" (
  "id" int8 NOT NULL DEFAULT nextval('user_cart_id_seq'::regclass),
  "shop_id" int4 NOT NULL,
  "uid" int4 NOT NULL,
  "goods_id" int4 NOT NULL,
  "goods_name" varchar(127) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "checked" bool NOT NULL DEFAULT true,
  "pic_url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "num" int4 NOT NULL,
  "sku_id" int4 NOT NULL,
  "spec" json NOT NULL,
  "spec_str" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" int4 NOT NULL,
  "updated_at" int4 NOT NULL
)
;
COMMENT ON COLUMN "mall-boot"."user_cart"."id" IS 'ID';
COMMENT ON COLUMN "mall-boot"."user_cart"."shop_id" IS '商户id';
COMMENT ON COLUMN "mall-boot"."user_cart"."uid" IS '用户id';
COMMENT ON COLUMN "mall-boot"."user_cart"."goods_id" IS '商品id';
COMMENT ON COLUMN "mall-boot"."user_cart"."goods_name" IS '商品名称';
COMMENT ON COLUMN "mall-boot"."user_cart"."checked" IS '购物车中商品是否选择状态';
COMMENT ON COLUMN "mall-boot"."user_cart"."pic_url" IS '商品图片或者商品货品图片';
COMMENT ON COLUMN "mall-boot"."user_cart"."num" IS '数量';
COMMENT ON COLUMN "mall-boot"."user_cart"."sku_id" IS 'sku id';
COMMENT ON COLUMN "mall-boot"."user_cart"."spec" IS 'spec';
COMMENT ON COLUMN "mall-boot"."user_cart"."spec_str" IS 'spec 描述';
COMMENT ON COLUMN "mall-boot"."user_cart"."created_at" IS '创建时间';
COMMENT ON COLUMN "mall-boot"."user_cart"."updated_at" IS '更新时间';
COMMENT ON TABLE "mall-boot"."user_cart" IS '用户购物表';

-- ----------------------------
-- Records of user_cart
-- ----------------------------

-- ----------------------------
-- Table structure for user_third
-- ----------------------------
DROP TABLE IF EXISTS "mall-boot"."user_third";
CREATE TABLE "mall-boot"."user_third" (
  "id" int8 NOT NULL DEFAULT nextval('user_third_id_seq'::regclass),
  "uid" int4 NOT NULL,
  "platform" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'wxapp'::character varying,
  "openid" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "unionid" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "nickname" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(512) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" int8 NOT NULL DEFAULT '0'::bigint
)
;
COMMENT ON COLUMN "mall-boot"."user_third"."id" IS 'id';
COMMENT ON COLUMN "mall-boot"."user_third"."uid" IS '会员ID';
COMMENT ON COLUMN "mall-boot"."user_third"."platform" IS '第三方应用';
COMMENT ON COLUMN "mall-boot"."user_third"."openid" IS '第三方唯一ID';
COMMENT ON COLUMN "mall-boot"."user_third"."unionid" IS '小程序unionid';
COMMENT ON COLUMN "mall-boot"."user_third"."nickname" IS '第三方会员昵称';
COMMENT ON COLUMN "mall-boot"."user_third"."avatar" IS '第三方会员头像';
COMMENT ON TABLE "mall-boot"."user_third" IS '第三方登录表';

-- ----------------------------
-- Records of user_third
-- ----------------------------
INSERT INTO "mall-boot"."user_third" VALUES (4, 7, 'wxapp', 'odbtT5eCbsuLGOdUDW7EcY5ZT-Bg', 'odbtT5eCbsuLGOdUDW7EcY5ZT-Bg', 'kyle', 'https://img.imguo.com/files/2023/10/13/7d51f441-9554-4008-9120-304254e00ce5.jpeg', '2023-10-12 22:39:38+08', '2023-10-24 17:43:03+08', 0);

-- ----------------------------
-- Function structure for on_update_current_timestamp_sys_dict_data
-- ----------------------------
DROP FUNCTION IF EXISTS "mall-boot"."on_update_current_timestamp_sys_dict_data"();
CREATE OR REPLACE FUNCTION "mall-boot"."on_update_current_timestamp_sys_dict_data"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_delivery_id_seq"
OWNED BY "mall-boot"."oms_order_delivery"."id";
SELECT setval('"mall-boot"."oms_order_delivery_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_id_seq"
OWNED BY "mall-boot"."oms_order"."id";
SELECT setval('"mall-boot"."oms_order_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_item_id_seq"
OWNED BY "mall-boot"."oms_order_item"."id";
SELECT setval('"mall-boot"."oms_order_item_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_log_id_seq"
OWNED BY "mall-boot"."oms_order_log"."id";
SELECT setval('"mall-boot"."oms_order_log_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_pay_id_seq"
OWNED BY "mall-boot"."oms_order_pay"."id";
SELECT setval('"mall-boot"."oms_order_pay_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."oms_order_setting_id_seq"
OWNED BY "mall-boot"."oms_order_setting"."id";
SELECT setval('"mall-boot"."oms_order_setting_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_brand_id_seq"
OWNED BY "mall-boot"."pms_brand"."id";
SELECT setval('"mall-boot"."pms_brand_id_seq"', 35, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_category_attribute_id_seq"
OWNED BY "mall-boot"."pms_category_attribute"."id";
SELECT setval('"mall-boot"."pms_category_attribute_id_seq"', 36, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_category_id_seq"
OWNED BY "mall-boot"."pms_category"."id";
SELECT setval('"mall-boot"."pms_category_id_seq"', 104, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_goods_attribute_id_seq"
OWNED BY "mall-boot"."pms_goods_attribute"."id";
SELECT setval('"mall-boot"."pms_goods_attribute_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_goods_id_seq"
OWNED BY "mall-boot"."pms_goods"."id";
SELECT setval('"mall-boot"."pms_goods_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_shop_id_seq"
OWNED BY "mall-boot"."pms_shop"."id";
SELECT setval('"mall-boot"."pms_shop_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."pms_sku_id_seq"
OWNED BY "mall-boot"."pms_sku"."id";
SELECT setval('"mall-boot"."pms_sku_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_dept_id_seq"
OWNED BY "mall-boot"."sys_dept"."id";
SELECT setval('"mall-boot"."sys_dept_id_seq"', 173, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_dict_data_id_seq"
OWNED BY "mall-boot"."sys_dict_data"."id";
SELECT setval('"mall-boot"."sys_dict_data_id_seq"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_dict_id_seq"
OWNED BY "mall-boot"."sys_dict"."id";
SELECT setval('"mall-boot"."sys_dict_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_log_id_seq"
OWNED BY "mall-boot"."sys_log"."id";
SELECT setval('"mall-boot"."sys_log_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_menu_id_seq"
OWNED BY "mall-boot"."sys_menu"."id";
SELECT setval('"mall-boot"."sys_menu_id_seq"', 127, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_notice_id_seq"
OWNED BY "mall-boot"."sys_notice"."id";
SELECT setval('"mall-boot"."sys_notice_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_role_id_seq"
OWNED BY "mall-boot"."sys_role"."id";
SELECT setval('"mall-boot"."sys_role_id_seq"', 10, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_user_id_seq"
OWNED BY "mall-boot"."sys_user"."id";
SELECT setval('"mall-boot"."sys_user_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."sys_user_notice_id_seq"
OWNED BY "mall-boot"."sys_user_notice"."id";
SELECT setval('"mall-boot"."sys_user_notice_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."user_address_id_seq"
OWNED BY "mall-boot"."user_address"."id";
SELECT setval('"mall-boot"."user_address_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."user_cart_id_seq"
OWNED BY "mall-boot"."user_cart"."id";
SELECT setval('"mall-boot"."user_cart_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."user_id_seq"
OWNED BY "mall-boot"."user"."id";
SELECT setval('"mall-boot"."user_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mall-boot"."user_third_id_seq"
OWNED BY "mall-boot"."user_third"."id";
SELECT setval('"mall-boot"."user_third_id_seq"', 4, true);

-- ----------------------------
-- Triggers structure for table sys_dict_data
-- ----------------------------
CREATE TRIGGER "on_update_current_timestamp" BEFORE UPDATE ON "mall-boot"."sys_dict_data"
FOR EACH ROW
EXECUTE PROCEDURE "mall-boot"."on_update_current_timestamp_sys_dict_data"();

-- ----------------------------
-- Indexes structure for table sys_menu
-- ----------------------------
CREATE UNIQUE INDEX "sys_menu_id_idx" ON "mall-boot"."sys_menu" USING btree (
  "id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table sys_user
-- ----------------------------
ALTER TABLE "mall-boot"."sys_user" ADD CONSTRAINT "sys_user_pkey" PRIMARY KEY ("id");
