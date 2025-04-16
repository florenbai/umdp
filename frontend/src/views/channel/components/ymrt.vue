<template>
  <a-card class="general-card">
    <template #title> 亿美软通电话配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="应用编号"
          field="channelConfig.appId"
          :rules="[
            {
              required: true,
              message: '请输入AppID',
            },
          ]"
        >
          <a-input
            v-model="ymrt.config.value.appId"
            placeholder="请输入App ID"
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
              message: '请输入App Secret',
            },
          ]"
        >
          <a-input-password
            v-model="ymrt.config.value.appSecret"
            placeholder="请输入App Secret"
            allow-clear
          />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="电话模板"
          field="channelConfig.templateId"
          :rules="[
            {
              required: true,
              message: '请输入电话模板编号',
            },
          ]"
        >
          <a-input
            v-model="ymrt.config.value.templateId"
            placeholder="请输入电话模板编号"
            allow-clear
          />
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { YmrtConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: YmrtConfig;
  }>();
  const ymrt = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => ymrt.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
