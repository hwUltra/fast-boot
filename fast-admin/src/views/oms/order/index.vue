<template>
  <div class="container">
    <el-button @click="openDialog"> 打开地图弹窗</el-button>
    {{ locationVal.address }} {{ locationVal.latitude }}
    {{ locationVal.longitude }}
    <el-dialog
      v-model="dialog.visible"
      :title="dialog.title"
      width="800px"
      @close="closeDialog"
    >
      <amap :modelValue="locationVal" @update:location="updateLocation" />
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
let locationVal = reactive<locationRef>({
  longitude: 116.407526,
  latitude: 39.90403,
  address: "北京市东城区王府井大街88号乐天银泰百货八层",
});

let changeLocation = <locationRef>{};

const dialog = reactive<DialogOption>({
  visible: false,
  title: "地图",
});

const openDialog = () => {
  dialog.visible = true;
};

const closeDialog = () => {
  dialog.visible = false;
};

const updateLocation = (val: locationRef) => {
  console.log("updateLocation", val);
  changeLocation = val;
};

const handleSubmit = () => {
  locationVal = changeLocation;
  closeDialog();
};

interface locationRef {
  address?: string;
  latitude?: number;
  longitude?: number;
  zone?: string[];
}
</script>
<style lang="scoped"></style>
