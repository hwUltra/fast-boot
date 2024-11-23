<template>
  <view class="work">
    <template v-for="item in gridList">
      <wd-card :title="item.title">
        <wd-grid clickable :column="4">
          <wd-grid-item
            v-for="(child, index) in item.children"
            :key="index"
            use-slot
            link-type="navigateTo"
            :url="child.url"
          >
            <view class="p-2">
              <image class="w-80rpx h-80rpx rounded-8rpx" :src="child.icon" />
            </view>
            <view class="text">{{ child.title }}</view>
          </wd-grid-item>
        </wd-grid>
      </wd-card>
    </template>
  </view>
</template>

<script lang="ts" setup>
const gridList = reactive([
  {
    title: "系统管理",
    children: [
      {
        icon: "/static/icons/user.png",
        title: "用户管理",
        url: "/subPackages/work/user/index",
      },

      {
        icon: "/static/icons/role.png",
        title: "角色管理",
      },
      {
        icon: "/static/icons/menu.png",
        title: "菜单管理",
      },
      {
        icon: "/static/icons/dept.png",
        title: "部门管理",
      },
      {
        icon: "/static/icons/dict.png",
        title: "字典管理",
      },
      {
        icon: "/static/icons/config.png",
        title: "系统配置",
      },
      {
        icon: "/static/icons/notice.png",
        title: "通知公告",
      },
      {
        icon: "/static/icons/more.png",
        title: "更多模块",
      },
    ],
  },
  {
    title: "系统监控",
    children: [
      {
        icon: "/static/icons/monitor.png",
        title: "实时监控",
      },
      {
        icon: "/static/icons/log.png",
        title: "日志管理",
      },
      {
        icon: "/static/icons/performance.png",
        title: "性能监控",
      },
      {
        icon: "/static/icons/more.png",
        title: "更多模块",
      },
    ],
  },
]);
</script>

<style lang="scss" scoped>
/* stylelint-disable selector-type-no-unknown */
page {
  background: #f8f8f8;
}
/* stylelint-enable selector-type-no-unknown */

.work {
  padding: 40rpx 0;
}
</style>
