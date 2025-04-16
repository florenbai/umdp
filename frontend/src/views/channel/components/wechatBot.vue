<template>
  <a-card class="general-card">
    <template #title> 企业微信-群机器人配置 </template>
    <a-row :gutter="80">
      <a-col :span="12"> 无需配置 </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { WechatBotConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: WechatBotConfig;
  }>();
  const wechat = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => wechat.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
