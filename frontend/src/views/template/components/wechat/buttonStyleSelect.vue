<template>
  <a-select placeholder="请选择按钮样式" :options="statusOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import { computed } from 'vue';

  const statusOptions = computed<SelectOptionData[]>(() => [
    {
      label: '蓝背景白字',
      value: 1,
    },
    {
      label: '灰背景蓝字',
      value: 2,
    },
    {
      label: '灰背景红字',
      value: 3,
    },
    {
      label: '灰背景黑字',
      value: 4,
    },
  ]);
</script>
