<template>
  <a-select placeholder="请选择消息类型" :options="statusOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import { computed } from 'vue';

  const statusOptions = computed<SelectOptionData[]>(() => [
    {
      label: '文本消息',
      value: 1,
    },
  ]);
</script>
