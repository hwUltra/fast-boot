import request from '@/utils/request';

const BASE_API = '';

export default {
	login(data) {
		return request({
			url: BASE_API + '/auth/code',
			method: 'post',
			data,
		});
	},
	logout(params) {
		return request({
			url: '/auth/logout',
			method: 'delete',
			params: params,
		});
	},
	decrypt(data) {
		return request({
			url: BASE_API + '/auth/decrypt',
			method: 'post',
			data,
		});
	},
	getPhoneNumber(data) {
		return request({
			url: BASE_API + '/auth/getMobile',
			method: 'post',
			data,
		});
	},
	bind(data) {
		return request({
			url: BASE_API + '/auth/bind',
			method: 'post',
			data,
		});
	},
	me() {
		return request({
			url: BASE_API + '/me',
			method: 'get'
		});
	},
	updateUserInfo(data) {
		return request({
			url: BASE_API + '/user/update',
			method: 'post',
			data,
		});
	},
};