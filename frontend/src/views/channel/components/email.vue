<template>
  <a-card class="general-card">
    <template #title> 邮箱配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="账号名称"
          field="channelConfig.username"
          :rules="[
            {
              required: true,
              message: '请输入邮箱账号名称',
            },
          ]"
        >
          <a-input
            v-model="email.config.value.username"
            placeholder="账号名称"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="账号密码"
          field="channelConfig.password"
          :rules="[
            {
              required: true,
              message: '请输入邮箱账号密码',
            },
          ]"
        >
          <a-input-password
            v-model="email.config.value.password"
            placeholder="请输入账号密码"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="STMP服务地址"
          field="channelConfig.smtpServer"
          :rules="[
            {
              required: true,
              message: '请输入邮箱STMP服务地址',
            },
          ]"
        >
          <a-input
            v-model="email.config.value.smtpServer"
            placeholder="请输入STMP服务地址"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="STMP服务端口"
          field="channelConfig.smtpPort"
          :rules="[
            {
              required: true,
              message: '请输入邮箱STMP服务端口',
            },
          ]"
        >
          <a-input
            v-model="email.config.value.smtpPort"
            placeholder="请输入STMP服务端口"
          ></a-input>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { EmailConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: EmailConfig;
  }>();
  const email = toRefs(props);
  const emits = defineEmits(['changeConfig']);

  watch(
    () => email.config,
    (value) => {
      emits('changeConfig', value);
    }
  );
</script>
