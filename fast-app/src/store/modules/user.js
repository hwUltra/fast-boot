import {
	defineStore
} from 'pinia';
import api from '@/api';
import config from '@/config.js';
import store from '@/store';

export const useUserStore = defineStore('user', () => {
	let openId = ref(''); //openID
	let token = ref(''); //token
	let userObj = ref({}); //用户信息

	let userGeoObj = ref({}); // 用户授权后存储的地理位置
	let isTest = ref(config.baseUrl.includes('127.0.0.1'));

	// 登录
	async function login() {
		if (token.value.length > 0) {
			return;
		}
		uni.login({
			provider: 'weixin',
			success: async (res) => {
				console.log("uni login res", res);
				// 请求后台获取openid注册登录成功后返回基本用户信息
				let apiRes = await api.user.login({
					code: res.code,
				});
				openId.value = apiRes.data.openID;
				token.value = apiRes.data.token;
			},
		});
	}

	//
	async function getUserProfile() {

		uni.getUserProfile({
			desc: "获取你的昵称、头像、地区及性别",
			success: async (res) => {
				console.log("getUserProfile - uni", res)
				let apiRes = await api.user.decrypt({
					openId: openId.value,
					iv: res.iv,
					encryptedData: res.encryptedData,
				});
				console.log("getUserProfile - apiRes", apiRes)
			},
			fail() {

			}
		})
	}

	async function bind(iv, encryptedData) {
		let apiRes = await api.user.bind({
			openId: openId.value,
			iv: iv,
			encryptedData: encryptedData,
		});
		openId.value = apiRes.data.openID;
		token.value = apiRes.data.token;
	}

	async function me() {
		let apiRes = await api.user.me();
		userObj.value = apiRes.data;
	}

	//
	async function updateUserInfo(changeUserObj) {
		let apiRes = await api.user.updateUserInfo(changeUserObj);
		if (apiRes.code == 200) {
			userObj.value = changeUserObj;
		}

	}

	// watch监听器
	watch(token, (newValue, oldValue) => {
		if (newValue) {
			if (isTest.value) {
				// #ifdef H5
				userGeoObj.value = {
					longitude: 104.0752,
					latitude: 30.55262
				}; // 天府三街定位 TODO 本地测试使用
				// #endif
			}
			me();
			// 初始化系统设置数据
			// store.system.useSystemStore().init();
		}
	});


	// 退出登录
	function logout() {
		userObj.value = {};
		token.value = '';
		// openId.value = '';
	}

	return {
		token,
		openId,
		isTest,
		userObj,
		userGeoObj,
		login,
		logout,
		bind,
		me,
		updateUserInfo,
	}
}, {
	unistorage: true // 开启后对 state 的数据读写都将持久化
});