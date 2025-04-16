<template>
  <a-select placeholder="请选择渠道状态" :options="statusOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import { computed } from 'vue';

  const statusOptions = computed<SelectOptionData[]>(() => [
    {
      label: '禁用',
      value: 0,
    },
    {
      label: '启用',
      value: 1,
    },
  ]);
</script>
