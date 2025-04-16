<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="20">
        <a-form-item
          label="消息类型"
          :field="`config[${props.configKey}].messageType`"
          :rules="[
            {
              required: true,
              message: '请选择消息类型',
            },
          ]"
        >
          <WechatMessageTypeSelect
            v-model="configRef.wechatConfig.value.messageType"
          ></WechatMessageTypeSelect>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <contentMessage
    v-if="
      configRef.wechatConfig.value.messageType === 1 ||
      configRef.wechatConfig.value.messageType === 2
    "
    :config-key="props.configKey"
    :wechat-config="configRef.wechatConfig.value"
  ></contentMessage>
  <buttonTemplateMessage
    v-if="configRef.wechatConfig.value.messageType === 3"
    :config-key="props.configKey"
    :button-template="configRef.wechatConfig.value.buttonTemplate"
    @message-changed="messageChange"
  ></buttonTemplateMessage>
  <textNoticeTemplateMessage
    v-if="configRef.wechatConfig.value.messageType === 4"
    :config-key="props.configKey"
    :text-notice-template="configRef.wechatConfig.value.textNoticeTemplate"
    @message-changed="textNoticeChange"
  >
  </textNoticeTemplateMessage>
  <a-form-item>
    <a-input-number
      v-model="configRef.wechatConfig.value.id"
      type="hidden"
    ></a-input-number>
  </a-form-item>
</template>

<script lang="ts" setup>
  import {
    ButtonTemplateData,
    TextNoticeTemplateData,
    WechatTemplateConfig,
  } from '@/api/template';
  import { toRefs } from 'vue';
  import WechatMessageTypeSelect from './wechatMessageTypeSelect.vue';
  import contentMessage from './wechatContentMessage.vue';
  import buttonTemplateMessage from './buttonTemplateMessage.vue';
  import textNoticeTemplateMessage from './textNoticeTemplateMessage.vue';

  export interface Props {
    wechatConfig: WechatTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.wechatConfig.value.id = configRef.channel.value;
  configRef.wechatConfig.value.type = '企业微信';

  const messageChange = (val: ButtonTemplateData) => {
    configRef.wechatConfig.value.buttonTemplate = val;
  };

  const textNoticeChange = (val: TextNoticeTemplateData) => {
    configRef.wechatConfig.value.textNoticeTemplate = val;
  };
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 100%;
    margin-bottom: 150px;
    padding: 20px;
    background-color: rgb(var(--gray-1));
  }
</style>
