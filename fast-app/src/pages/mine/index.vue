<template>
	<view>
		<uni-badge text="1"></uni-badge>
		<uni-badge text="2" type="success" ></uni-badge>
		<uni-badge text="3" type="primary" :inverted="true"></uni-badge>
	</view>
</template>


<script setup>
	import {BlankImg,BlankName} from "@/static/resources.js";
	
	const {
		proxy
	} = getCurrentInstance();
	
	let {
		token,
		userObj
	} = toRefs(proxy.$store.user.useUserStore());
	
	async function bind(e) {
		console.log(e)
		if (e.detail.iv) {
			await proxy.$store.user.useUserStore().bind(e.detail.iv, e.detail.encryptedData);
		}
	}
	
	async function localLogin() {
	  // await proxy.$store.user.useUserStore().localLogin();
	}

	function logout() {
		proxy.$store.user.useUserStore().logout();
	}

	function tips() {
		// proxy.submitFail('暂未开发...');
		uni.navigateTo({
		    url: '/subPackages/mine/address'
		});
	}
</script>

<style lang="scss" scoped>
	page {
		background-color: #ededed;
	}

	.user-box {
		background-color: #fff;
	}

	.u-cell-group {
		margin-top: 20rpx;
		background-color: #fff;
	}
</style>