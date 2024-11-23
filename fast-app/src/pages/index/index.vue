<template>
  <view class="flex-center flex-col">
    <image class="logo" src="/static/logo.png" />
    <view>
      <view class="text-blue font-bold text-lg">{{ title }}</view>
    </view>

    <view class="mt-5 flex-center">
      <view class="text-green font-bold text-lg">
        {{ userStore.userInfo?.nickname ? userStore.userInfo.nickname : "未登录" }}</view
      >
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/store/modules/user";

const title = ref("fast-app");

const userStore = useUserStore();
</script>

<style>
.logo {
  width: 200rpx;
  height: 200rpx;
  margin-top: 200rpx;
  margin-right: auto;
  margin-bottom: 50rpx;
  margin-left: auto;
}
</style>
