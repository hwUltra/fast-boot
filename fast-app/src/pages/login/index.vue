<template>
  <view class="login">
    <image class="logo" src="@/static/logo.png" style="width: 100px; height: 100px" />
    <wd-form ref="loginFormRef" :model="loginFormData">
      <wd-cell-group border>
        <wd-input
          v-model="loginFormData.mobile"
          label="用户名"
          label-width="100px"
          prop="mobile"
          clearable
          placeholder="请输入用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <wd-input
          v-model="loginFormData.vcode"
          label-width="100px"
          type="text"
          label="验证码"
          prop="vcode"
          center
          placeholder="请输入验证码"
          :rules="[{ required: true, message: '请填写密码' }]"
          clearable
        >
          <template #suffix>
            <wd-button v-if="!countDown" size="small" custom-class="button" @click="handleSendCode"
              >获取验证码
            </wd-button>
            <wd-count-down
              v-else
              :time="time"
              format="mm:ss"
              @finish="onTimeFinish"
              class="custom-count-down"
            />
          </template>
        </wd-input>
      </wd-cell-group>

      <view class="footer">
        <wd-button size="large" type="primary" block @click="handleLogin">提交</wd-button>
      </view>
    </wd-form>
  </view>
</template>

<script lang="ts" setup>
import { type LoginFormData } from "@/api/auth";
import { useUserStore } from "@/store/modules/user";

const countDown = ref(false);
const time = ref<number>(2 * 60 * 1000);

const loginFormRef = ref();

const handleSendCode = () => {
  countDown.value = true;
  console.log("发送验证码");
};

const onTimeFinish = () => {
  countDown.value = false;
};

const loginFormData = ref<LoginFormData>({
  mobile: "13312312312",
  vcode: "202020",
});

const userStore = useUserStore();

// 登录处理
const handleLogin = () => {
  loginFormRef.value.validate().then(async ({ valid }: { valid: boolean }) => {
    if (valid) {
      try {
        await userStore.login(loginFormData.value); // 等待登录和获取用户信息完成
        await userStore.getInfo(); // 等待用户信息获取完成
        uni.showToast({ title: "登录成功", icon: "success" });
        const pages = getCurrentPages(); // 获取当前的页面栈

        console.log("pages", pages);
        if (pages.length > 1) {
          // 如果页面栈中有多个页面，则可以返回上一页
          uni.navigateBack();
        } else {
          // 如果页面栈中只有一个页面（通常是首页），则可以跳转到指定页面，避免 navigateBack 无法返回的问题
          uni.reLaunch({
            url: "/pages/index/index", // 替换为你想要跳转的页面路径
          });
        }
      } catch (error: any) {
        console.log("登录失败", error.message);
      }
    }
  });
};
</script>

<style lang="scss" scoped>
.login {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: calc(100vh - 50px);
}
.footer {
  padding: 16px;
}

.custom-count-down {
  display: inline-block;
  width: 84px;
  height: 28px;
  color: #282121;
  font-size: 14px;
  text-align: center;
  background-color: #fff;
  border-radius: 2px;
}
</style>
