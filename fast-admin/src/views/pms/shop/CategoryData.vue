<script setup lang="ts">
import {
  categoryList,
  getOptions,
  categoryAdd,
  categoryUpdate,
  categoryDel,
} from "@/api/pms/category";
import {
  CategoryQuery,
  CategoryVO,
  CategoryForm,
} from "@/api/pms/category/types";

const props = defineProps({
  shopId: {
    type: Number,
    default: () => {
      return 1;
    },
  },
});

const emits = defineEmits(["handle:over"]);
const loading = ref(false);
const queryParams = reactive<CategoryQuery>({
  pageNum: 1,
  pageSize: 10,
});
const pageList = ref<CategoryVO[]>([]);

const queryFormRef = ref(ElForm);
const removeIds = ref([]);

const formRef = ref(ElForm);
const formData = reactive<CategoryForm>({
  id: 0,
  parentId: 0,
  sort: 1,
  status: 1,
});

const categoryOptionsData = ref<OptionType[]>();
// 弹窗对象
const dialog = reactive({
  visible: false,
  width: 800,
  title: "",
});

onMounted(() => {
  handleQuery();
});

function handleQuery() {
  loading.value = true;
  categoryList(queryParams)
    .then(({ data }) => {
      pageList.value = data.list;
    })
    .finally(() => {
      loading.value = false;
    });
  emits("handle:over", props.shopId);
}
function resetQuery() {
  queryFormRef.value.resetFields();
  handleQuery();
}

function handleSelectionChange(selection: any) {
  removeIds.value = selection.map((item: any) => item.id);
}

function handleDelete(id?: number) {
  const ids = [id || removeIds.value].join(",");
  console.log(ids);
  if (!ids) {
    ElMessage.warning("请勾选删除项");
    return;
  }

  ElMessageBox.confirm("确认删除用户?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(function () {
      categoryDel(ids).then(() => {
        ElMessage.success("删除成功");
        resetQuery();
      });
    })
    .catch(function () {});
}

async function loadOptions() {
  getOptions(props.shopId).then((response) => {
    categoryOptionsData.value = [
      {
        value: 0,
        label: "顶级部门",
        children: response.data.list,
      },
    ];
  });
}

const rules = reactive({
  title: [{ required: true, message: "title不能为空", trigger: "blur" }],
  image: [{ required: true, message: "图像不能为空", trigger: "blur" }],
});

async function openDialog(id: number, row: CategoryVO) {
  await loadOptions();
  console.log("openDialog", categoryOptionsData);
  dialog.visible = true;
  if (id > 0) {
    Object.assign(formData, row);
  }
}

function closeDialog() {
  dialog.visible = false;
  formData.id = 0;
  formData.title = "";
  formData.image = "";
  formData.sort = 0;
  formData.status = 1;
  formData.subtitle = "";
  formData.parentId = 0;
}

function handleSubmit() {
  formRef.value.validate((valid: any) => {
    if (valid) {
      loading.value = true;
      formData.shopId = props.shopId;
      console.log("formData", formData.title, formData.subtitle);
      if (formData.id) {
        categoryUpdate(formData)
          .then(() => {
            ElMessage.success("修改成功");
            closeDialog();
            resetQuery();
          })
          .finally(() => (loading.value = false));
      } else {
        categoryAdd(formData)
          .then(() => {
            ElMessage.success("新增成功");
            closeDialog();
            resetQuery();
          })
          .finally(() => (loading.value = false));
      }
    }
  });
}
</script>
<template>
  <div class="app-container">
    <div class="search-container">
      <!-- 搜索表单 -->
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="名称"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery"
            ><i-ep-search />搜索</el-button
          >
          <el-button @click="resetQuery"> <i-ep-refresh />重置</el-button>
        </el-form-item>
      </el-form>

      <!-- 表格 -->
      <el-card shadow="never">
        <template #header>
          <div class="flex justify-between">
            <div>
              <el-button type="success" @click="openDialog(0)"
                ><i-ep-plus />新增</el-button
              >
              <el-button
                type="danger"
                :disabled="removeIds.length === 0"
                @click="handleDelete()"
                ><i-ep-delete />删除</el-button
              >
            </div>
          </div>
        </template>

        <el-table
          v-loading="loading"
          :data="pageList"
          highlight-current-row
          row-key="id"
          default-expand-all
          :tree-props="{
            children: 'children',
            hasChildren: 'hasChildren',
          }"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="50" align="center" />
          <el-table-column label="标题" align="left" width="120" prop="title" />
          <el-table-column
            label="副标题"
            align="left"
            width="200"
            prop="subtitle"
          />
          <el-table-column label="头像" width="100">
            <template #default="scope">
              <img
                :src="scope.row.image"
                style="max-height: 60px; max-width: 60px"
              />
            </template>
          </el-table-column>
          <el-table-column label="状态" align="center" prop="status">
            <template #default="scope">
              <el-tag :type="scope.row.status == 1 ? 'success' : 'info'">{{
                scope.row.status == 1 ? "启用" : "禁用"
              }}</el-tag>
            </template>
          </el-table-column>

          <el-table-column label="操作" fixed="right" width="220">
            <template #default="scope">
              <el-button
                type="primary"
                size="small"
                link
                @click="handleDelete(scope.row.id)"
                ><i-ep-refresh-left />删除</el-button
              >
              <el-button
                type="primary"
                link
                size="small"
                @click="openDialog(scope.row.id, scope.row)"
                ><i-ep-edit />编辑</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-card>

      <!-- 弹窗 -->
      <el-dialog
        v-model="dialog.visible"
        :title="dialog.title"
        :width="dialog.width"
        append-to-body
        @close="closeDialog"
      >
        <!-- 用户新增/编辑表单 -->
        <el-form
          ref="formRef"
          :model="formData"
          :rules="rules"
          label-width="80px"
        >
          <el-form-item label="上级部门" prop="parentId">
            <el-tree-select
              v-model="formData.parentId"
              placeholder="选择上级部门"
              :data="categoryOptionsData"
              filterable
              check-strictly
              :render-after-expand="false"
            />
          </el-form-item>
          <el-form-item label="标题" prop="title">
            <el-input v-model="formData.title" placeholder="请输入用户昵称" />
          </el-form-item>
          <el-form-item label="副标题" prop="subtitle">
            <el-input
              v-model="formData.subtitle"
              placeholder="请输入用户昵称"
            />
          </el-form-item>
          <el-form-item label="图标" prop="image">
            <single-upload v-model="formData.image" />
          </el-form-item>
          <el-form-item label="显示排序" prop="sort">
            <el-input-number
              v-model="formData.sort"
              controls-position="right"
              style="width: 100px"
              :min="0"
            />
          </el-form-item>
          <el-form-item label="状态">
            <el-radio-group v-model="formData.status">
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
    </div>
  </div>
</template>
<style lang="scoped"></style>
