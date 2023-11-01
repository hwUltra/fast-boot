<template>
  <view class="flex-column h-full w-full">
    <slot />
  </view>
</template>
<script setup>
const { proxy } = getCurrentInstance();
const props = defineProps({
  activeTabName: { type: String, required: true },
});
</script>
<style lang="scss" scoped></style>
