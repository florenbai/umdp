<template>
  <a-card class="general-card">
    <template #title> 钉钉配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="应用编号"
          field="channelConfig.appKey"
          :rules="[
            {
              required: true,
              message: '请输入钉钉Client ID',
            },
          ]"
        >
          <a-input
            v-model="dingtalk.config.value.appKey"
            placeholder="请输入钉钉Client ID"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="应用密钥"
          field="channelConfig.appSecret"
          :rules="[
            {
              required: true,
              message: '请输入钉钉Client Secret',
            },
          ]"
        >
          <a-input-password
            v-model="dingtalk.config.value.appSecret"
            placeholder="请输入钉钉Client Secret"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="应用编号"
          field="channelConfig.agentId"
          :rules="[
            {
              required: true,
              message: '请输入钉钉内部应用AgentId',
            },
          ]"
        >
          <a-input
            v-model="dingtalk.config.value.agentId"
            placeholder="请输入钉钉内部应用AgentId"
          ></a-input>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { DingTalkConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: DingTalkConfig;
  }>();
  const dingtalk = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => dingtalk.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
