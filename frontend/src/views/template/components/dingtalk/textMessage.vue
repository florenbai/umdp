<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col v-if="configRef.dingTalkConfig.value.messageType === 2" :span="20">
        <a-form-item
          label="消息标题"
          :field="`config[${props.configKey}].title`"
          :rules="[
            {
              required: true,
              message: '请输入消息标题',
            },
          ]"
        >
          <a-input
            v-model="configRef.dingTalkConfig.value.title"
            placeholder="请输入消息内容"
            allow-clear
          />
        </a-form-item>
      </a-col>
      <a-col :span="20">
        <a-form-item
          label="消息内容"
          :field="`config[${props.configKey}].content`"
          :rules="[
            {
              required: true,
              message: '请选择消息内容',
            },
          ]"
        >
          <a-textarea
            v-model="configRef.dingTalkConfig.value.content"
            placeholder="请输入消息内容"
            allow-clear
            show-word-limit
          />
          <template #help>
            <span>
              内容说明：支持参数变量定义，您可以使用例如{$a}，定义a变量，支持多个变量定义</span
            >
          </template>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { DingTalkTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';

  export interface Props {
    dingTalkConfig: DingTalkTemplateConfig;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
</script>
