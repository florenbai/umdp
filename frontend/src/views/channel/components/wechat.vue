<template>
  <a-card class="general-card">
    <template #title> 企业微信配置 </template>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="企业编号"
          field="channelConfig.corpid"
          :rules="[
            {
              required: true,
              message: '请输入企业微信corpid',
            },
          ]"
        >
          <a-input
            v-model="wechat.config.value.corpid"
            placeholder="请输入corpid"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="企业密钥"
          field="channelConfig.corpsecret"
          :rules="[
            {
              required: true,
              message: '请输入企业微信corpsecret',
            },
          ]"
        >
          <a-input-password
            v-model="wechat.config.value.corpsecret"
            placeholder="请输入corpsecret"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="应用编号"
          field="channelConfig.agentid"
          :rules="[
            {
              required: true,
              message: '请输入企业微信agentid',
            },
          ]"
        >
          <a-input
            v-model="wechat.config.value.agentid"
            placeholder="请输入agentid"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="应用密钥"
          field="channelConfig.secret"
          :rules="[
            {
              required: true,
              message: '请输入企业微信secret',
            },
          ]"
        >
          <a-input-password
            v-model="wechat.config.value.secret"
            placeholder="请输入secret"
            allow-clear
          ></a-input-password>
        </a-form-item>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item label="接收消息Token">
          <a-input
            v-model="wechat.config.value.token"
            placeholder="请输入接收消息Token"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="接收消息EncodingAESKey">
          <a-input
            v-model="wechat.config.value.encodingAESKey"
            placeholder="请输入接收消息EncodingAESKey"
          ></a-input>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { WechatConfig } from '@/api/channel';
  import { toRefs, watch } from 'vue';

  const props = defineProps<{
    config: WechatConfig;
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
