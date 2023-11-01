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
        <el-table
          v-loading="loading"
          :data="pageData"
          highlight-current-row
          border
        >
          <el-table-column label="ID" prop="id" width="80" />
          <el-table-column label="名称" prop="name" width="120" />
          <el-table-column label="电话" prop="tel" width="120" />
          <el-table-column label="配送费" prop="distributionFee" width="90" />
          <el-table-column label="状态" align="center" prop="status">
            <template #default="scope">
              <el-tag :type="scope.row.status == 1 ? 'success' : 'info'">{{
                scope.row.status == 1 ? "启用" : "禁用"
              }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center" prop="createdAt" />

          <el-table-column label="操作" fixed="right" width="220">
            <template #default="scope">
              <el-button
                type="primary"
                size="small"
                link
                @click="opencDialog(scope.row.id)"
                ><i-ep-refresh-left />分类管理</el-button
              >
              <el-button
                type="primary"
                link
                size="small"
                @click="openDialog(scope.row)"
                ><i-ep-edit />编辑</el-button
              >
            </template>
          </el-table-column>
        </el-table>
        <pagination
          v-if="total > 0"
          v-model:total="total"
          v-model:page="queryParams.pageNum"
          v-model:limit="queryParams.pageSize"
          @pagination="handleQuery"
        />
      </el-card>

      <el-dialog
        v-model="dialog.visible"
        :title="dialog.title"
        width="1000px"
        append-to-body
        @close="closeDialog"
      >
        <el-form
          ref="shopFormRef"
          :model="shopData"
          :rules="shopRule"
          label-width="80px"
        >
          <el-form-item label="用户名" prop="name">
            <el-input v-model="shopData.name" placeholder="请输入名称" />
          </el-form-item>
          <el-form-item label="电话" prop="tel">
            <el-input v-model="shopData.tel" placeholder="请输入名称" />
          </el-form-item>
          <el-form-item label="公告" prop="notice">
            <el-input v-model="shopData.notice" placeholder="请输入名称" />
          </el-form-item>
          <el-form-item label="配送费" prop="distributionFee">
            <el-input
              v-model.number="shopData.distributionFee"
              placeholder="请输入名称"
            />
          </el-form-item>
          <el-form-item label="状态" prop="status">
            <el-radio-group v-model="shopData.status">
              <el-radio :label="1">正常</el-radio>
              <el-radio :label="0">禁用</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
        <!-- 弹窗底部操作按钮 -->
        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="handleSubmit">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </template>
      </el-dialog>

      <el-dialog
        v-model="cDataDialog.visible"
        :title="cDataDialog.title"
        width="1000px"
        @close="closeCDialog"
      >
        <category-data
          v-model:shopId="selectedId"
          @handle:over="categoryOver"
        />
      </el-dialog>
    </div>
  </div>
</template>
<script setup lang="ts">
import { shopList, shopUpdate } from "@/api/pms/shop";
import { ShopQuery, ShopVO, ShopForm } from "@/api/pms/shop/types";

import CategoryData from "@/views/pms/shop/CategoryData.vue";

const state = reactive({
  loading: false,
  total: 0,
  pageData: [] as ShopVO[],
  queryParams: {
    pageNum: 1,
    pageSize: 10,
  } as ShopQuery,
  shopData: {
    status: 1,
  } as ShopForm,
  shopRule: {
    name: [{ required: true, message: "用户名不能为空", trigger: "blur" }],
    tel: [{ required: true, message: "请输入正确的手机号码", trigger: "blur" }],
  },
  dialog: {
    visible: false,
    title: "「 编辑商户 」",
  } as DialogType,
  cDataDialog: {
    title: "「 分类管理 」",
    visible: false,
  } as DialogType,
});

const {
  loading,
  total,
  pageData,
  queryParams,
  shopData,
  shopRule,
  dialog,
  cDataDialog,
} = toRefs(state);

const queryFormRef = ref(ElForm);
const shopFormRef = ref(ElForm);

onMounted(() => {
  handleQuery();
});

function handleQuery() {
  // 重置父组件
  loading.value = true;
  shopList(state.queryParams)
    .then(({ data }) => {
      state.pageData = data.list;
      state.total = data.total;
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

const categoryOver = (val: number) => {
  console.log("categoryOver----", val);
};
// function handleDelete(id?: number) {
//   console.log(id);
// }

async function openDialog(row: ShopVO) {
  dialog.value.visible = true;
  Object.assign(shopData.value, row);
}

function closeDialog() {
  dialog.value.visible = false;
  selectedId.value = 0;
  shopFormRef.value.resetFields();
  shopFormRef.value.clearValidate();
  //清空
  shopData.value.id = 0;
  shopData.value.name = "";
  shopData.value.tel = "";
  shopData.value.notice = "";
  shopData.value.distributionFee = 0;
  shopData.value.status = 1;
}

const selectedId: Ref<number> = ref(0);
function opencDialog(id: number) {
  selectedId.value = id;
  cDataDialog.value.visible = true;
}

function closeCDialog() {
  cDataDialog.value.visible = false;
  selectedId.value = 0;
}

function handleSubmit() {
  shopFormRef.value.validate((valid: any) => {
    if (valid) {
      shopUpdate(shopData.value)
        .then(() => {
          ElMessage.success("修改成功");
          closeDialog();
          resetQuery();
        })
        .finally(() => (loading.value = false));
    }
  });
}
</script>
