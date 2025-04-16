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
            v-model="configRef.dingTalkConfig.value.messageType"
          ></MessageTypeSelect>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <textMessage
    v-if="
      configRef.dingTalkConfig.value.messageType === 1 ||
      configRef.dingTalkConfig.value.messageType === 2
    "
    :config-key="props.configKey"
    :ding-talk-config="configRef.dingTalkConfig.value"
  ></textMessage>
  <a-form-item>
    <a-input-number
      v-model="configRef.dingTalkConfig.value.id"
      type="hidden"
    ></a-input-number>
  </a-form-item>
</template>

<script lang="ts" setup>
  import { DingTalkTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';
  import MessageTypeSelect from './messageTypeSelect.vue';
  import textMessage from './textMessage.vue';

  export interface Props {
    dingTalkConfig: DingTalkTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.dingTalkConfig.value.id = configRef.channel.value;
  configRef.dingTalkConfig.value.type = '钉钉';
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 100%;
    margin-bottom: 150px;
    padding: 20px;
    background-color: rgb(var(--gray-1));
  }
</style>
