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
          <el-button type="primary" @click="handleQuery">
            <template #icon><i-ep-search /></template>
            搜索
          </el-button>
          <el-button @click="resetQuery">
            <template #icon><i-ep-refresh /></template>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-card shadow="never" class="table-container">
      <template #header>
        <el-button type="success" @click="openDialog('user-form')">
          <i-ep-plus />
          新增
        </el-button>
        <el-button type="danger" @click="handleDelete()">
          <i-ep-delete />
          删除
        </el-button>
      </template>

      <el-table
        v-loading="loading"
        :data="pageData"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="ID" prop="id" width="80" />
        <el-table-column label="头像" width="80">
          <template #default="scope">
            <img :src="scope.row.avatar" width="40" height="40" />
          </template>
        </el-table-column>
        <el-table-column label="用户名" prop="username" min-width="100" />
        <el-table-column label="昵称" prop="nickname" min-width="100" />
        <el-table-column label="手机" prop="mobile" min-width="100" />
        <el-table-column label="性别" align="center" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.gender == 0">未知</el-tag>
            <el-tag v-if="scope.row.gender == 1">男</el-tag>
            <el-tag v-if="scope.row.gender == 2">女</el-tag>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" min-width="220">
          <template #default="scope">
            <el-button
              type="primary"
              size="small"
              link
              @click="resetPassword(scope.row)"
            >
              <i-ep-refresh-left />
              重置密码
            </el-button>
            <el-button
              type="primary"
              link
              size="small"
              @click="openDialog('user-form', scope.row.id)"
            >
              <i-ep-edit />
              编辑
            </el-button>
            <el-button
              type="primary"
              link
              size="small"
              @click="handleDelete(scope.row.id)"
            >
              <i-ep-delete />
              删除
            </el-button>
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
        v-if="dialog.type === 'user-form'"
        ref="userFormRef"
        :model="formData"
        :rules="userRule"
        label-width="80px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="formData.username"
            :readonly="!!formData.id"
            placeholder="请输入用户名"
          />
        </el-form-item>
        <el-form-item label="用户昵称" prop="nickname">
          <el-input v-model="formData.nickname" placeholder="请输入用户昵称" />
        </el-form-item>
        <el-form-item label="性别" prop="gender">
          <el-select
            v-model="formData.gender"
            placeholder="Select"
            size="large"
          >
            <el-option
              v-for="item in genderOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="formData.mobile" placeholder="请输入手机号码" />
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <SingleImageUpload v-model="formData.avatar" />
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
</template>
<script setup lang="ts">
import UserAPI, { UserQuery, UserPageVO, UserForm } from "@/api/ums/user";

const loading = ref(false);
const total = ref(0);
const queryFormRef = ref(ElForm); // 查询表单
const userFormRef = ref(ElForm); // 用户表单
const pageData = ref<UserPageVO[]>(); // 用户分页数据
const removeIds = ref([]); // 删除用户ID集合 用于批量删除

// 用户表单数据
const formData = reactive<UserForm>({
  id: 0,
});

// 用户表单数据
const queryParams = reactive<UserQuery>({
  pageNum: 1,
  pageSize: 10,
});

// 弹窗对象
const dialog = reactive({
  visible: false,
  type: "user-form",
  width: 800,
  title: "",
});

// 校验规则
const userRule = reactive({
  username: [{ required: true, message: "用户名不能为空", trigger: "blur" }],
  nickname: [{ required: true, message: "用户昵称不能为空", trigger: "blur" }],
  avatar: [{ required: true, message: "头像不能为空", trigger: "blur" }],
  mobile: [
    {
      required: true,
      pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
      message: "请输入正确的手机号码",
      trigger: "blur",
    },
  ],
});

const genderOptions = [
  {
    value: 1,
    label: "男",
  },
  {
    value: 2,
    label: "女",
  },
  {
    value: 0,
    label: "保密",
  },
];

onMounted(() => {
  handleQuery();
});

/**
 * 查询
 */
function handleQuery() {
  // 重置父组件
  loading.value = true;
  UserAPI.getList(queryParams)
    .then((data) => {
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

/** 重置密码 */
function resetPassword(row: { [key: string]: any }) {
  ElMessageBox.prompt(
    "请输入用户「" + row.nickname + "」的新密码",
    "重置密码",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
    }
  )
    .then(({ value }) => {
      if (!value) {
        ElMessage.warning("请输入新密码");
        return false;
      }
      // updateUserPassword(row.id, value).then(() => {
      //   ElMessage.success("密码重置成功，新密码是：" + value);
      // });
    })
    .catch(() => {});
}

function handleDelete(id?: number) {
  const userIds = [id || removeIds.value].join(",");
  if (!userIds) {
    ElMessage.warning("请勾选删除项");
    return;
  }

  ElMessageBox.confirm("确认删除用户?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(function () {
      UserAPI.delUser(userIds).then(() => {
        ElMessage.success("删除成功");
        resetQuery();
      });
    })
    .catch(() => {});
}

/** 行选中 */
function handleSelectionChange(selection: any) {
  removeIds.value = selection.map((item: any) => item.id);
}

async function openDialog(type: string, id?: number) {
  dialog.visible = true;
  dialog.type = type;
  if (dialog.type === "user-form") {
    if (id) {
      dialog.title = "修改用户";
      UserAPI.getUserInfo(id).then((data) => {
        Object.assign(formData, { ...data });
      });
    } else {
      dialog.title = "新增用户";
    }
  }
}

/**
 * 关闭弹窗
 *
 * @param type 弹窗类型  用户表单：user-form | 用户导入：user-import
 */
function closeDialog() {
  dialog.visible = false;
  if (dialog.type === "user-form") {
    userFormRef.value.resetFields();
    userFormRef.value.clearValidate();
    formData.id = 0;
    formData.avatar = "";
  } else {
  }
}

/** 表单提交 */
const handleSubmit = useThrottleFn(() => {
  if (dialog.type === "user-form") {
    userFormRef.value.validate((valid: any) => {
      if (valid) {
        const userId = formData.id;
        console.log(formData);
        loading.value = true;
        if (userId) {
          UserAPI.updateUser(formData)
            .then(() => {
              ElMessage.success("修改用户成功");
              closeDialog();
              resetQuery();
            })
            .finally(() => (loading.value = false));
        } else {
          UserAPI.addUser(formData)
            .then(() => {
              ElMessage.success("新增用户成功");
              closeDialog();
              resetQuery();
            })
            .finally(() => (loading.value = false));
        }
      }
    });
  }
}, 3000);
</script>
