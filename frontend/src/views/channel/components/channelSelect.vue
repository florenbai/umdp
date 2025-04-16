<template>
  <a-select placeholder="请选择渠道" :options="statusOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import { computed } from 'vue';

  const statusOptions = computed<SelectOptionData[]>(() => [
    {
      label: '邮箱',
      value: '邮箱',
    },
    {
      label: '企业微信',
      value: '企业微信',
    },
    {
      label: '企业微信-群机器人',
      value: '企业微信-群机器人',
    },
    {
      label: '钉钉',
      value: '钉钉',
    },
    {
      label: '飞书',
      value: '飞书',
    },
    {
      label: '阿里云短信',
      value: '阿里云短信',
    },
    {
      label: '腾讯云短信',
      value: '腾讯云短信',
    },
    {
      label: '亿美软通电话',
      value: '亿美软通电话',
    },
  ]);
</script>
