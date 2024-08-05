<template>
  <div class="component-container">
    <div class="component-container__main">
      <el-form-item label="店铺" prop="roleIds">
        <el-select v-model="shopId" placeholder="请选择">
          <el-option
            v-for="item in shopList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-cascader-panel
        ref="categoryRef"
        v-model="goodsInfo.categoryId"
        :options="categoryOptions"
        :props="{ emitPath: false }"
        @change="handleCategoryChange"
      />
      <div style="margin-top: 20px">
        <el-link v-show="pathLabels.length > 0" type="info" :underline="false"
          >您选择的商品分类:</el-link
        >
        <el-link
          v-for="(item, index) in pathLabels"
          :key="index"
          type="danger"
          :underline="false"
          class="ml-[5px]"
          style="margin-left: 5px"
        >
          {{ item }}
          <span
            v-show="index < pathLabels.length - 1"
            style="width: 1em; height: 1em; margin-left: 5px"
            ><i-ep-caret-right
          /></span>
        </el-link>
      </div>
    </div>
    <div class="component-container__footer">
      <el-button type="primary" @click="handleNext"
        >下一步，填写商品信息</el-button
      >
    </div>
  </div>
</template>

<script setup lang="ts">
// API 引用
import { getCategoryOptions } from "@/api/pms/category";
import { getShopOptions } from "@/api/pms/shop";

const emit = defineEmits(["next", "update:modelValue"]);
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => {},
  },
});

const shopList = ref<OptionType[]>();

const shopId = ref(1);

watch(shopId, (newValue, oldValue) => {
  console.log("watch 已触发", newValue, oldValue);
  loadCategoryOptions(newValue);
});

const goodsInfo: any = computed({
  get: () => props.modelValue,
  set: (value) => {
    emit("update:modelValue", value);
  },
});

const state = reactive({
  categoryOptions: [] as OptionType[],
  pathLabels: [],
});

const { categoryOptions, pathLabels } = toRefs(state);

async function loadData() {
  await loadShopOptions();
  await loadCategoryOptions(shopId.value);
}

function loadShopOptions() {
  getShopOptions().then((response) => {
    shopList.value = response.data;
  });
}

function loadCategoryOptions(id: number) {
  goodsInfo.value.shopId = id;
  getCategoryOptions(id).then(({ data }) => {
    state.categoryOptions = data;
    if (goodsInfo.value.id) {
      nextTick(() => {
        handleCategoryChange();
      });
    }
  });
}

const categoryRef = ref(ElCascaderPanel);

function handleCategoryChange() {
  const checkNode = categoryRef.value.getCheckedNodes()[0];
  state.pathLabels = checkNode.pathLabels; // 商品分类选择层级提示
  goodsInfo.value.categoryId = checkNode.value;
}

function handleNext() {
  if (!goodsInfo.value.categoryId) {
    ElMessage.warning("请选择商品分类");
    return false;
  }
  emit("next");
}

onMounted(() => {
  loadData();
});
</script>

<style lang="scss" scoped>
.component-container {
  &__main {
    margin: 20px auto;
  }

  &__footer {
    position: fixed;
    right: 20px;
    bottom: 20px;
  }
}
</style>
