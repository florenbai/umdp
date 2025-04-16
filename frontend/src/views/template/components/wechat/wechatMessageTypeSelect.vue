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
    {
      label: 'MarkDown消息',
      value: 2,
    },
    {
      label: '模板卡片消息-按钮交互',
      value: 3,
    },
    {
      label: '模板卡片消息-文本通知',
      value: 4,
    },
  ]);
</script>
