<template>
	<view>
		<uni-section title="自定义内容" subTitle="使用 left/right/default 插槽自定义内容" type="line" style="margin-bottom: 3px;">
			<view class="box-bg">
				<uni-nav-bar>
					<block slot="left">
						<view class="city">
							<view>
								<text class="uni-nav-bar-text">武汉</text>
							</view>
							<uni-icons type="arrowdown" color="#666" size="18" />
						</view>
					</block>
					<view class="input-view">
						<uni-icons class="input-uni-icon" type="search" size="18" color="#999" />
						<input confirm-type="search" class="nav-bar-input" type="text" placeholder="输入搜索关键词"
							@confirm="confirm" />
					</view>
					<block slot="right">
						<view class="city">
							搜索
						</view>
					</block>
				</uni-nav-bar>
			</view>
		</uni-section>

	</view>
</template>

<script setup>
	const {
		proxy
	} = getCurrentInstance();
	const disabled = ref(true);

	function enter() {
		console.log("xxxx");
		uni.navigateTo({
			url: '/subPackages/home/tt'
		});
	}

	function confirm() {
		console.log('@leftClick');
	}
</script>

<style lang="scss" scoped>

</style>