<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="短信签名"
          :field="`config[${props.configKey}].signName`"
          :rules="[
            {
              required: true,
              message: '请输入短信签名名称',
            },
          ]"
        >
          <a-input
            v-model="configRef.smsConfig.value.signName"
            placeholder="请输入短信签名名称"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="短信模板Code"
          :field="`config[${props.configKey}].templateCode`"
          :rules="[
            {
              required: true,
              message: '请输入短信模板Code',
            },
          ]"
        >
          <a-input
            v-model="configRef.smsConfig.value.templateCode"
            placeholder="请输入短信模板Code"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="上行短信扩展码">
          <a-input
            v-model="configRef.smsConfig.value.smsUpExtendCode"
            placeholder="请输入上行短信扩展码"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="外部流水扩展字段">
          <a-input
            v-model="configRef.smsConfig.value.outId"
            placeholder="请输入外部流水扩展字段"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-form-item>
        <a-input-number
          v-model="configRef.smsConfig.value.id"
          type="hidden"
        ></a-input-number>
      </a-form-item>
    </a-row>
  </a-card>
  <a-card class="general-card">
    <template #title> 短信模板变量 </template>
    <template #extra>
      <a-button type="primary" @click="handleAdd">
        <template #icon>
          <icon-plus />
        </template>
        新增短信模板变量
      </a-button>
    </template>
    <a-row v-for="(value, index) of data" :key="index" :gutter="80">
      <a-col :span="8">
        <a-form-item label="变量名称">
          <a-input
            v-model="value.name"
            placeholder="请输入变量名称"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="参数值">
          <a-input
            v-model="value.value"
            placeholder="请输入参数值"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-button
        type="primary"
        :style="{ marginLeft: '10px', marginTop: '30px' }"
        @click="handleDelete(index)"
      >
        <template #icon> <icon-delete /> </template>
        <template #default>删除</template>
      </a-button>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { AliyunSmsTemplateConfig, AliyunTemplateParam } from '@/api/template';
  import { toRefs, ref, onMounted } from 'vue';

  export interface Props {
    smsConfig: AliyunSmsTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  const data = ref<AliyunTemplateParam[]>([]);
  onMounted(() => {
    if (props.smsConfig.templateParam !== undefined) {
      data.value = props.smsConfig.templateParam;
    }
    configRef.smsConfig.value.id = configRef.channel.value;
    configRef.smsConfig.value.type = '阿里云短信';
  });

  const handleAdd = () => {
    data.value.push({
      name: '',
      value: '',
    });
    configRef.smsConfig.value.templateParam = data.value;
  };

  const handleDelete = (index: number) => {
    data.value.splice(index, 1);
    configRef.smsConfig.value.templateParam = data.value;
  };
</script>
