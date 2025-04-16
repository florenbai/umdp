<template>
  <a-modal
    v-model:visible="visible"
    fullscreen
    :loading="loading"
    @before-ok="handleOk"
    @cancel="handleCancel"
  >
    <template #title> 调用调试 </template>
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-space direction="vertical" :size="16">
        <a-card class="general-card">
          <template #title> 基础参数 </template>
          <a-row :gutter="80">
            <a-col :span="12">
              <a-form-item
                field="templateId"
                label="模板编号"
                :rules="[
                  {
                    required: true,
                    message: '请输入模板编号',
                  },
                ]"
              >
                <a-input-number v-model="form.templateId" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item
                field="channel"
                label="渠道名称"
                :rules="[
                  {
                    required: true,
                    message: '请选择渠道',
                  },
                ]"
              >
                <template-channel-select
                  v-model="form.channel"
                  :profession="props.profession"
                ></template-channel-select>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="80">
            <a-col :span="8">
              <a-form-item
                v-for="(receiver, index) of form.parameters.receiver"
                :key="index"
                :field="`parameters.receiver[${index}]`"
                :label="`接收人-${index + 1}`"
                :rules="[
                  {
                    required: true,
                    message: '请输入接收人',
                  },
                ]"
              >
                <a-input
                  v-model="form.parameters.receiver[index]"
                  placeholder="请输入接收人"
                />
                <a-button
                  style="margin-left: 20px"
                  type="primary"
                  @click="receiverAdd"
                  ><template #icon> <icon-plus /> </template>
                  <template #default>新增接收人</template></a-button
                >
                <a-button
                  status="danger"
                  :style="{ marginLeft: '10px' }"
                  @click="receiverDelete(index)"
                  ><template #icon> <icon-delete /> </template>
                  <template #default>删除接收人</template></a-button
                >
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="80">
            <a-col :span="8">
              <a-form-item
                v-for="(cc, index) of form.parameters.cc"
                :key="index"
                :label="`抄送人-${index + 1}`"
              >
                <a-input
                  v-model="form.parameters.cc[index]"
                  placeholder="请输入抄送人，抄送人只对邮箱类型的渠道有效"
                />
                <a-button
                  style="margin-left: 20px"
                  type="primary"
                  @click="ccAdd"
                  ><template #icon> <icon-plus /> </template>
                  <template #default>新增抄送人</template></a-button
                >
                <a-button
                  status="danger"
                  :style="{ marginLeft: '10px' }"
                  @click="ccDelete(index)"
                  ><template #icon> <icon-delete /> </template>
                  <template #default>删除抄送人</template></a-button
                >
              </a-form-item>
            </a-col>
          </a-row>
        </a-card>
        <a-card class="general-card">
          <template #title> 模板变量 </template>
          <template #extra>
            <a-button type="primary" @click="handleParamAdd">
              <template #icon>
                <icon-plus />
              </template>
              新增模板变量
            </a-button>
          </template>
          <a-row v-for="(value, index) of dataParam" :key="index" :gutter="80">
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
              status="danger"
              :style="{ marginLeft: '10px', marginTop: '30px' }"
              @click="handleParamDelete(index)"
            >
              <template #icon> <icon-delete /> </template>
              <template #default>删除</template>
            </a-button>
          </a-row>
        </a-card>
        <a-card class="general-card">
          <template #title> 请求头信息 </template>
          <template #extra>
            <showJson></showJson>
          </template>
          <a-descriptions title="" :column="1">
            <a-descriptions-item label="token">
              <a-tag>{{ props.token }}</a-tag>
            </a-descriptions-item>
          </a-descriptions>
        </a-card>
      </a-space>
    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import useLoading from '@/hooks/loading';
  import {
    AliyunTemplateParam,
    TemplateTestParam,
    submitTestTemplate,
  } from '@/api/template';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import { Message } from '@arco-design/web-vue';
  import templateChannelSelect from './templateChannelSelect.vue';
  import showJson from './showJson.vue';

  export interface Props {
    id: number;
    profession: number;
    show: boolean;
    token: string;
  }

  const { loading, setLoading } = useLoading();
  const formRef = ref<FormInstance>();
  const props = defineProps<Props>();
  const emits = defineEmits(['showChange']);
  const dataParam = ref<AliyunTemplateParam[]>([]);

  const form = reactive<TemplateTestParam>({
    templateId: props.id,
    channel: '',
    parameters: {
      receiver: [''],
      cc: [''],
      variable: {},
    },
  });

  const visible = ref(false);
  visible.value = props.show;

  const summitTest = async () => {
    if (dataParam.value.length > 0) {
      dataParam.value.forEach((element) => {
        form.parameters.variable[element.name] = element.value;
      });
    }
    const res = await formRef.value?.validate();
    if (!res) {
      setLoading(true);
      try {
        await submitTestTemplate(form);
        Message.success({
          content: '测试完成',
          duration: 5 * 1000,
        });
      } catch (err) {
        // you can report use errorHandler or other
      } finally {
        setLoading(false);
      }
    }
  };

  const handleOk = (done: any) => {
    done(false);
    summitTest();
  };

  const handleCancel = () => {
    visible.value = false;
    emits('showChange');
  };

  const receiverAdd = () => {
    form.parameters.receiver.push('');
  };
  const receiverDelete = (index: number) => {
    if (form.parameters.receiver.length < 2) {
      Message.error('至少包含1个接收人');
      return;
    }
    form.parameters.receiver.splice(index, 1);
  };

  const ccAdd = () => {
    form.parameters.cc.push('');
  };
  const ccDelete = (index: number) => {
    form.parameters.cc.splice(index, 1);
  };

  const handleParamAdd = () => {
    dataParam.value.push({
      name: '',
      value: '',
    });
  };

  const handleParamDelete = (index: number) => {
    dataParam.value.splice(index, 1);
  };
</script>
