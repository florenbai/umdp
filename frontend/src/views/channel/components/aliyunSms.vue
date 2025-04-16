<template>
  <a-card class="general-card">
    <template #title> 短信配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="访问令牌"
          field="channelConfig.accessKeyId"
          :rules="[
            {
              required: true,
              message: '请输入阿里云访问令牌',
            },
          ]"
        >
          <a-input
            v-model="aliyunSms.config.value.accessKeyId"
            placeholder="请输入accessKeyId"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="访问密钥"
          field="channelConfig.accessKeySecret"
          :rules="[
            {
              required: true,
              message: '请输入阿里云访问密钥',
            },
          ]"
        >
          <a-input-password
            v-model="aliyunSms.config.value.accessKeySecret"
            placeholder="请输入accessKeySecret"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="地域ID"
          field="channelConfig.regionId"
          :rules="[
            {
              required: true,
              message: '请输入阿里云地域ID',
            },
          ]"
        >
          <a-input
            v-model="aliyunSms.config.value.regionId"
            placeholder="请输入地域ID，例如：cn-hangzhou"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="公网接入地址"
          field="channelConfig.endpoint"
          :rules="[
            {
              required: true,
              message: '请输入阿里云公网接入地址',
            },
          ]"
        >
          <a-input
            v-model="aliyunSms.config.value.endpoint"
            placeholder="请输入endpoint"
          ></a-input>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { AliyunSmsConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: AliyunSmsConfig;
  }>();
  const aliyunSms = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => aliyunSms.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
