<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="20">
        <a-form-item
          label="未接通呼叫次数"
          :field="`config[${props.configKey}].callCount`"
          :rules="[
            {
              required: true,
              message: '请输入未接通呼叫次数',
            },
          ]"
        >
          <a-input-number
            v-model="configRef.ymrtConfig.value.callCount"
            :style="{ width: '120px' }"
            placeholder="请输入未接通呼叫次数"
            :min="0"
            :max="5"
            mode="button"
          />
        </a-form-item>
      </a-col>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="20">
        <a-form-item label="消息内容">
          <a-textarea
            v-model="configRef.ymrtConfig.value.content"
            placeholder="请输入消息内容"
            show-word-limit
            disabled
          />
          <template #help>
            <span> 内容说明：模板电话渠道内容已在服务端定义，不能被修改</span>
          </template>
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <a-form-item>
    <a-input-number
      v-model="configRef.ymrtConfig.value.id"
      type="hidden"
    ></a-input-number>
  </a-form-item>
</template>

<script lang="ts" setup>
  import { PhoneTemplateConfig } from '@/api/template';
  import { toRefs } from 'vue';

  export interface Props {
    ymrtConfig: PhoneTemplateConfig;
    channel: number;
    configKey: any;
  }
  const props = defineProps<Props>();
  const configRef = toRefs(props);
  configRef.ymrtConfig.value.id = configRef.channel.value;
  configRef.ymrtConfig.value.type = '亿美软通电话';
  configRef.ymrtConfig.value.content =
    '告警通知，{$alertTitle} 出现告警，告警次数 {$alertCount} 次。已持续 {$duration}，请尽快处理。';
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 100%;
    margin-bottom: 150px;
    padding: 20px;
    background-color: rgb(var(--gray-1));
  }
</style>
