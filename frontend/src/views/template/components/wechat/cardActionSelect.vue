<template>
  <a-select placeholder="请选择跳转事件类型" :options="cardActionOptions">
  </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import { computed } from 'vue';

  const cardActionOptions = computed<SelectOptionData[]>(() => [
    {
      label: '无',
      value: 0,
    },
    {
      label: '跳转url',
      value: 1,
    },
    {
      label: '打开小程序',
      value: 2,
    },
  ]);
</script>
