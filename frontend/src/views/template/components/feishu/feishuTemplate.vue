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
          <MessageTypeSelect
            v-model="configRef.feishuConfig.value.messageType"
          ></MessageTypeSelect>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <textMessage
    v-if="configRef.feishuConfig.value.messageType === 1"
    :config-key="props.configKey"
    :feishu-config="configRef.feishuConfig.value"
  ></textMessage>
  <a-form-item>
    <a-input-number
      v-model="configRef.feishuConfig.value.id"
      type="hidden"
    ></a-input-number>
  </a-form-item>
</template>

<script lang="ts" setup>
  import { FeiShuTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';
  import MessageTypeSelect from './messageTypeSelect.vue';
  import textMessage from './textMessage.vue';

  export interface Props {
    feishuConfig: FeiShuTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.feishuConfig.value.id = configRef.channel.value;
  configRef.feishuConfig.value.type = '飞书';
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 100%;
    margin-bottom: 150px;
    padding: 20px;
    background-color: rgb(var(--gray-1));
  }
</style>
