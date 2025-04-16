<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="12">
        <a-form-item
          label="短信应用ID"
          :field="`config[${props.configKey}].appId`"
          :rules="[
            {
              required: true,
              message: '请输入短信应用ID',
            },
          ]"
        >
          <a-input
            v-model="configRef.txsmsConfig.value.appId"
            placeholder="请输入短信应用ID"
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
            v-model="configRef.txsmsConfig.value.templateCode"
            placeholder="请输入短信模板Code"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="短信签名">
          <a-input
            v-model="configRef.txsmsConfig.value.signName"
            placeholder="请输入短信签名名称"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="短信码号扩展号">
          <a-input
            v-model="configRef.txsmsConfig.value.smsUpExtendCode"
            placeholder="请输入短信码号扩展号"
          ></a-input>
        </a-form-item>
      </a-col>
      <a-form-item>
        <a-input-number
          v-model="configRef.txsmsConfig.value.id"
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
        <a-form-item label="参数值">
          <a-input
            v-model="data[index]"
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
  import { TencentSmsTemplateConfig } from '@/api/template';
  import { toRefs, ref, onMounted } from 'vue';

  export interface Props {
    txsmsConfig: TencentSmsTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  const data = ref<string[]>([]);
  onMounted(() => {
    if (props.txsmsConfig.templateParam !== undefined) {
      data.value = props.txsmsConfig.templateParam;
    }
    configRef.txsmsConfig.value.id = configRef.channel.value;
    configRef.txsmsConfig.value.type = '腾讯云短信';
  });

  const handleAdd = () => {
    data.value.push('');
    configRef.txsmsConfig.value.templateParam = data.value;
  };

  const handleDelete = (index: number) => {
    data.value.splice(index, 1);
    configRef.txsmsConfig.value.templateParam = data.value;
  };
</script>
