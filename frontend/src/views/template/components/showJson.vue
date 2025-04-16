<template>
  <a-button type="primary" @click="handleClick">查看请求样例</a-button>
  <a-drawer
    :width="700"
    :visible="visible"
    unmount-on-close
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 请求样例 </template>
    <a-descriptions title="请求信息" bordered :column="1">
      <a-descriptions-item label="请求地址">
        {{ axios.defaults.baseURL }}/api/v1/message/send
      </a-descriptions-item>
      <a-descriptions-item label="请求方式"> POST </a-descriptions-item>
      <a-descriptions-item label="请求头参数"> token </a-descriptions-item>
      <a-descriptions-item label="Content-Type">
        application/json
      </a-descriptions-item>
    </a-descriptions>
    <br />
    <a-descriptions title="请求参数" bordered />
    <a-table :columns="columns" :data="data" :pagination="false" />
    <br />
    <div class="details-wrapper">
      <a-typography-title :heading="6" style="margin-top: 0">
        自定义参数样例
      </a-typography-title>
      <a-typography-paragraph style="margin-bottom: 0">
        在模板中，您可以使用 `{$}`
        进行变量定义。然后在调用时，传递parameters.variable即可。
      </a-typography-paragraph>
      <a-typography-paragraph style="margin-bottom: 0">
        例如：在模板中定义了一个 {$title} 变量。
        在调用时，您需传递parameters.variable.title进行变量赋值。
      </a-typography-paragraph>
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import axios from 'axios';

  const visible = ref(false);
  const columns = [
    {
      title: '参数名称',
      dataIndex: 'name',
      ellipsis: true,
      tooltip: true,
      width: 180,
    },
    {
      title: '参数类型',
      dataIndex: 'type',
      width: 120,
    },
    {
      title: '是否必填',
      dataIndex: 'essential',
      ellipsis: true,
      width: 100,
    },
    {
      title: '说明',
      dataIndex: 'desc',
      ellipsis: true,
      tooltip: { position: 'left' },
      width: 400,
    },
  ];

  const data = reactive([
    {
      key: '1',
      name: 'templateId',
      type: '整形',
      essential: '是',
      desc: '模板编号',
    },
    {
      key: '2',
      name: 'channel',
      type: '字符串',
      essential: '是',
      desc: '渠道标识',
    },
    {
      key: '3',
      name: 'parameters',
      type: '对象',
      essential: '是',
      desc: '参数对象',
    },
    {
      key: '4',
      name: 'parameters.receiver',
      type: '字符串数组',
      essential: '是',
      desc: '发送对象',
    },
    {
      key: '5',
      name: 'parameters.cc',
      type: '数组',
      essential: '否',
      desc: '抄送对象,只针对邮箱渠道有效',
    },
    {
      key: '6',
      name: 'parameters.variable',
      type: '对象',
      essential: '否',
      desc: '自定义参数对象,用于在模板中的自定义参数。',
    },
  ]);
  const handleClick = () => {
    visible.value = true;
  };
  const handleOk = () => {
    visible.value = false;
  };
  const handleCancel = () => {
    visible.value = false;
  };
</script>

<style scoped lang="less">
  .details-wrapper {
    width: 730px;
    margin-top: 20px;
    padding: 20px;
    text-align: left;
    background-color: var(--color-fill-2);
  }
</style>
