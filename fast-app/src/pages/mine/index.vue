<template>
  <view class="mine">
    <!-- 个人资料 -->
    <view class="header">
      <view class="flex items-center">
        <image
          class="w-100rpx h-100rpx rounded-full"
          :src="isLogin ? userInfo!.avatar : defaultAvatar"
        />
        <view class="ml-20rpx text-32rpx">
          <text v-if="isLogin">{{ userInfo!.nickname }}</text>
          <view v-else @click="goToLoginPage">
            <view>您还未登录，请先</view>
            <view class="cursor-pointer"> 登录 <wd-icon name="arrow-right" /> </view>
          </view>
        </view>
      </view>
      <view>
        <view v-if="isLogin" class="cursor-pointer" @click="goToProfile">
          个人信息 <wd-icon name="arrow-right" />
        </view>
      </view>
    </view>
    <view class="relative -top-[50px] p-[40rpx]">
      <!-- 功能操作 -->
      <wd-grid clickable class="rounded-[10rpx] bg-white">
        <wd-grid-item :to="{ path: '/group' }" icon="setting" text="交流群" />
        <wd-grid-item :to="{ path: '/customer-service' }" icon="setting" text="在线客服" />
        <wd-grid-item :to="{ path: '/feedback' }" icon="setting" text="反馈社区" />
        <wd-grid-item :to="{ path: '/like-us' }" icon="setting" text="点赞我们" />
      </wd-grid>
      <!-- 菜单列表 -->
      <view class="mt50rpx">
        <wd-cell-group border>
          <wd-cell title="编辑资料" icon="setting" @click="goToEditProfile">
            <wd-icon name="arrow-right" />
          </wd-cell>
          <wd-cell title="常见问题" icon="setting" @click="goToFAQ">
            <wd-icon name="arrow-right" />
          </wd-cell>
          <wd-cell title="关于我们" icon="setting" @click="goToAbout">
            <wd-icon name="arrow-right" />
          </wd-cell>
          <wd-cell title="应用设置" icon="setting" clickable @click="goToSettings">
            <wd-icon name="arrow-right" />
          </wd-cell>

          <wd-cell title="退出登录" icon="setting" clickable @click="handleLogout">
            <wd-icon name="arrow-right" />
          </wd-cell>
        </wd-cell-group>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { useUserStore } from "@/store/modules/user";

const userStore = useUserStore();

const userInfo = computed(() => userStore.userInfo);

const isLogin = computed(() => !!userInfo.value);

const defaultAvatar = "/static/images/default-avatar.png";

const goToLoginPage = () => {
  uni.navigateTo({ url: "/pages/login/index" });
};

const goToProfile = () => {
  uni.navigateTo({ url: "/subPackages/mine/profile/index" });
};
const goToEditProfile = () => {
  uni.navigateTo({ url: "/pages/edit-profile/index" });
};
const goToFAQ = () => {
  uni.navigateTo({ url: "/pages/faq/index" });
};
const goToAbout = () => {
  uni.navigateTo({ url: "/pages/about/index" });
};
const goToSettings = () => {
  uni.navigateTo({ url: "/subPackages/mine/setting/index" });
};

const handleLogout = () => {
  userStore.logout().then(() => {
    uni.showToast({ title: "已退出登录", icon: "success" });
  });
};
</script>

<style lang="scss" scoped>
/* stylelint-disable selector-type-no-unknown */
page {
  background: #f8f8f8;
}
/* stylelint-enable selector-type-no-unknown */

.mine {
  height: 100%;
  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 50rpx 20rpx 100rpx 20rpx;
    color: #fff;
    background: linear-gradient(60deg, #165dff, #6aa1ff);
  }
}
</style>
