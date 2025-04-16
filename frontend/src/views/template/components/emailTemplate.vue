<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="20">
        <a-form-item
          label="邮件标题"
          :field="`config[${props.configKey}].title`"
          :rules="[
            {
              required: true,
              message: '请输入邮件标题',
            },
          ]"
        >
          <a-input
            v-model="configRef.emailConfig.value.title"
            placeholder="请输入邮件标题"
          ></a-input>
        </a-form-item>
        <template #help>
          <span>
            内容说明：支持参数变量定义，您可以使用例如{$a}，定义a变量，支持多个变量定义</span
          >
        </template>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="20">
        <a-form-item
          label="邮件内容"
          :field="`config[${props.configKey}].content`"
          :rules="[
            {
              required: true,
              message: '请选择邮件内容',
            },
          ]"
        >
          <div style="border: 1px solid #ccc">
            <Toolbar
              style="border-bottom: 1px solid #ccc"
              :editor="editorRef"
              :default-config="toolbarConfig"
              mode="default"
            />
            <Editor
              v-model="configRef.emailConfig.value.content"
              style="height: 500px; overflow-y: hidden"
              :default-config="editorConfig"
              mode="default"
              @on-created="handleCreated"
            />
          </div>
          <template #help>
            <span>
              内容说明：支持参数变量定义，您可以使用例如{$a}，定义a变量，支持多个变量定义</span
            >
          </template>
        </a-form-item>
        <a-form-item>
          <a-input-number
            v-model="configRef.emailConfig.value.id"
            type="hidden"
          ></a-input-number>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { EmailTemplateConfig } from '@/api/template';
  import { toRefs, shallowRef, onBeforeUnmount } from 'vue';
  import '@wangeditor/editor/dist/css/style.css';
  import { Editor, Toolbar } from '@wangeditor/editor-for-vue';

  export interface Props {
    emailConfig: EmailTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.emailConfig.value.id = configRef.channel.value;
  configRef.emailConfig.value.type = '邮箱';
  const editorRef = shallowRef();
  const toolbarConfig = {};
  const editorConfig = { placeholder: '请输入内容...' };
  onBeforeUnmount(() => {
    const editor = editorRef.value;
    if (editor == null) return;
    editor.destroy();
  });
  const handleCreated = (editor: any) => {
    editorRef.value = editor; // 记录 editor 实例，重要！
  };
</script>
