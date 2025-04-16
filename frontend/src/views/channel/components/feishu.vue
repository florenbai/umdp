<template>
  <a-card class="general-card">
    <template #title> 飞书配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="应用编号"
          field="channelConfig.appId"
          :rules="[
            {
              required: true,
              message: '请输入飞书App ID',
            },
          ]"
        >
          <a-input
            v-model="feishu.config.value.appId"
            placeholder="请输入飞书App ID"
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
              message: '请输入飞书App Secret',
            },
          ]"
        >
          <a-input-password
            v-model="feishu.config.value.appSecret"
            placeholder="请输入飞书App Secret"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { FeishuConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: FeishuConfig;
  }>();
  const feishu = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => feishu.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
