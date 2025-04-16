<template>
  <a-card class="general-card">
    <a-row :gutter="80">
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
            v-model="configRef.wechatConfig.value.content"
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
  import { WechatTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';

  export interface Props {
    wechatConfig: WechatTemplateConfig;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
</script>
