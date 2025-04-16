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
          <botMessageTypeSelect
            v-model="configRef.wechatConfig.value.messageType"
          ></botMessageTypeSelect>
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
  <a-form-item>
    <a-input-number
      v-model="configRef.wechatConfig.value.id"
      type="hidden"
    ></a-input-number>
  </a-form-item>
</template>

<script lang="ts" setup>
  import { WechatBotTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';
  import botMessageTypeSelect from './botMessageTypeSelect.vue';
  import contentMessage from './wechatContentMessage.vue';

  export interface Props {
    wechatConfig: WechatBotTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.wechatConfig.value.id = configRef.channel.value;
  configRef.wechatConfig.value.type = '企业微信-群机器人';
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 100%;
    margin-bottom: 150px;
    padding: 20px;
    background-color: rgb(var(--gray-1));
  }
</style>
