<template>
  <div class="app-container">
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="菜单名称"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery"
            ><template #icon><i-ep-search /></template>搜索</el-button
          >
          <el-button @click="resetQuery">
            <template #icon><i-ep-refresh /></template>
            重置</el-button
          >
        </el-form-item>
      </el-form>

      <el-card shadow="never" class="table-container">
        <el-table v-loading="loading" :data="pageData">
          <el-table-column label="ID" prop="id" width="80" />
          <el-table-column label="UID" prop="uid" width="80" />
          <el-table-column label="平台" prop="platform" width="80" />
          <el-table-column label="头像" width="100">
            <template #default="scope">
              <el-popover placement="right" :width="400" trigger="hover">
                <img :src="scope.row.avatar" width="400" height="400" />
                <template #reference>
                  <img
                    :src="scope.row.avatar"
                    style="max-height: 60px; max-width: 60px"
                  />
                </template>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column label="昵称" prop="nickname" min-width="80" />
          <el-table-column label="openid" prop="openid" min-width="120" />
          <el-table-column label="unionid" prop="unionid" min-width="120" />
        </el-table>
        <pagination
          v-if="total > 0"
          v-model:total="total"
          v-model:page="queryParams.pageNum"
          v-model:limit="queryParams.pageSize"
          @pagination="handleQuery"
        />
      </el-card>
    </div>
  </div>
</template>
<script setup lang="ts">
import { getThirdList } from "@/api/ums/user";
import { UserQuery, UserThirdPageVO } from "@/api/ums/user/types";

const loading = ref(false);
const total = ref(0);
const queryFormRef = ref(ElForm); // 查询表单
const pageData = ref<UserThirdPageVO[]>(); // 用户分页数据
// 用户表单数据
// 用户表单数据
const queryParams = reactive<UserQuery>({
  pageNum: 1,
  pageSize: 10,
});

onMounted(() => {
  handleQuery();
});

function handleQuery() {
  // 重置父组件
  loading.value = true;
  getThirdList(queryParams)
    .then(({ data }) => {
      pageData.value = data.list;
      total.value = data.total;
    })
    .finally(() => {
      loading.value = false;
    });
}

/** 重置查询 */
function resetQuery() {
  queryFormRef.value.resetFields();
  handleQuery();
}
</script>
