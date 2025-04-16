<template>
  <div class="container">
    <Breadcrumb :items="['menu.template', `${title}`]" />
    <a-form ref="formRef" layout="vertical" :model="formData">
      <a-space direction="vertical" :size="16">
        <a-card class="general-card">
          <template #title> 基础信息 </template>
          <a-row :gutter="80">
            <a-col :span="12">
              <a-form-item
                label="模板名称"
                field="templateName"
                :rules="[
                  {
                    required: true,
                    message: '请输入模板名称',
                  },
                ]"
              >
                <a-input
                  v-model="formData.templateName"
                  placeholder="请输入模板名称"
                ></a-input>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item
                label="所属业务"
                field="professionId"
                :rules="[
                  {
                    required: true,
                    message: '请选择所属业务',
                  },
                ]"
              >
                <professionNameSelect
                  v-model="formData.professionId"
                  @change="getChannels()"
                ></professionNameSelect>
              </a-form-item>
            </a-col>
          </a-row>
        </a-card>
        <a-card class="general-card">
          <template #title> 模板配置 </template>
          <a-tabs>
            <a-tab-pane
              v-for="(item, key) in tabData"
              :key="item.id"
              :title="item.channelName + '-' + item.channelTag"
            >
              <EmailTemplate
                v-if="item.channelName === '邮箱'"
                :email-config="emailTemplateConfig"
                :config-key="key"
                :channel="item.id"
              ></EmailTemplate>
              <WechatTemplate
                v-else-if="item.channelName === '企业微信'"
                :wechat-config="wechatTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></WechatTemplate>
              <botTemplate
                v-else-if="item.channelName === '企业微信-群机器人'"
                :wechat-config="wechatBotTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></botTemplate>
              <DingTalkTemplate
                v-else-if="item.channelName === '钉钉'"
                :ding-talk-config="dingtalkTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></DingTalkTemplate>
              <FeiShuTemplate
                v-else-if="item.channelName === '飞书'"
                :feishu-config="feishuTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></FeiShuTemplate>
              <AliyunSmsTemplate
                v-else-if="item.channelName === '阿里云短信'"
                :sms-config="aliyunSmsTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></AliyunSmsTemplate>
              <TencentSmsTemplate
                v-else-if="item.channelName === '腾讯云短信'"
                :txsms-config="tencentSmsTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></TencentSmsTemplate>
              <YmrtPhoneTemplate
                v-else-if="item.channelName === '亿美软通电话'"
                :ymrt-config="ymrtTemplateConfig"
                :channel="item.id"
                :config-key="key"
              ></YmrtPhoneTemplate>
            </a-tab-pane>
          </a-tabs>
        </a-card>
      </a-space>
      <div class="actions">
        <a-space>
          <a-button @click="gotoTemplateList"> 返回 </a-button>
          <a-button type="primary" :loading="loading" @click="onSubmitClick">
            提交
          </a-button>
        </a-space>
      </div>
    </a-form>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref, watch } from 'vue';
  import { FormInstance } from '@arco-design/web-vue/es/form';
  import useLoading from '@/hooks/loading';
  import {
    TemplateForm,
    submitTemplateForm,
    submitEditTemplateForm,
    getTemplateById,
    TemplateRecord,
    EmailTemplateConfig,
    WechatTemplateConfig,
    AliyunSmsTemplateConfig,
    TencentSmsTemplateConfig,
    DingTalkTemplateConfig,
    FeiShuTemplateConfig,
    PhoneTemplateConfig,
    WechatBotTemplateConfig,
  } from '@/api/template';
  import {
    getProfessionChannels,
    ProfessionChannelRecord,
  } from '@/api/profession';
  import { Message } from '@arco-design/web-vue';
  import { useRoute } from 'vue-router';
  import { gotoTemplateList } from './router';
  import professionNameSelect from './components/professionNameSelect.vue';
  import EmailTemplate from './components/emailTemplate.vue';
  import WechatTemplate from './components/wechat/wechatTemplate.vue';
  import AliyunSmsTemplate from './components/sms/aliyunSmsTemplate.vue';
  import TencentSmsTemplate from './components/sms/tencentSmsTemplate.vue';
  import DingTalkTemplate from './components/dingtalk/dingtalkTemplate.vue';
  import FeiShuTemplate from './components/feishu/feishuTemplate.vue';
  import YmrtPhoneTemplate from './components/ymrt/phoneTemplate.vue';
  import botTemplate from './components/wechat/botTemplate.vue';

  const title = ref<string>('menu.template.add');
  const formData = ref<TemplateForm>({
    professionId: null,
    templateName: '',
    config: [],
  });
  const id = ref<string>();
  const editChannels = ref<string[]>([]);
  const tabData = ref<ProfessionChannelRecord[]>();
  const route = useRoute();
  const formRef = ref<FormInstance>();
  const emailTemplateConfig = ref<EmailTemplateConfig>({
    id: 0,
    type: '邮箱',
    title: '',
    content: '',
  });
  const wechatTemplateConfig = ref<WechatTemplateConfig>({
    id: 0,
    type: '企业微信',
    messageType: 1,
  });
  const wechatBotTemplateConfig = ref<WechatBotTemplateConfig>({
    id: 0,
    type: '企业微信-群机器人',
    messageType: 1,
    content: '',
  });
  const aliyunSmsTemplateConfig = ref<AliyunSmsTemplateConfig>({
    id: 0,
    type: '阿里云短信',
    signName: '',
    templateCode: '',
    outId: '',
  });
  const tencentSmsTemplateConfig = ref<TencentSmsTemplateConfig>({
    id: 0,
    type: '腾讯云短信',
    appId: '',
    signName: '',
    templateCode: '',
  });
  const dingtalkTemplateConfig = ref<DingTalkTemplateConfig>({
    id: 0,
    type: '钉钉',
    messageType: 1,
    content: '',
  });
  const feishuTemplateConfig = ref<FeiShuTemplateConfig>({
    id: 0,
    type: '飞书',
    messageType: 1,
    content: '',
  });
  const ymrtTemplateConfig = ref<PhoneTemplateConfig>({
    id: 0,
    type: '亿美软通电话',
    callCount: 0,
    content: '',
  });

  const { loading, setLoading } = useLoading();
  const onSubmitClick = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      setLoading(true);
      try {
        if (id.value) {
          await submitEditTemplateForm(id.value, formData.value);
        } else {
          await submitTemplateForm(formData.value);
        }
        gotoTemplateList();
        Message.success({
          content: '操作成功',
          duration: 5 * 1000,
        });
      } catch (err) {
        // you can report use errorHandler or other
      } finally {
        setLoading(false);
      }
    }
    setTimeout(() => {
      setLoading(false);
    }, 1000);
  };

  watch(
    route,
    () => {
      if (route.params?.id) {
        id.value = String(route.params?.id);
      } else {
        id.value = undefined;
      }
    },
    {
      deep: true,
      immediate: true,
    }
  );

  const getChannels = async () => {
    if (formData.value.professionId !== null) {
      const { data } = await getProfessionChannels(formData.value.professionId);
      if (data) {
        tabData.value = {
          ...(data as unknown as ProfessionChannelRecord[]),
        };
        if (tabData.value !== undefined) {
          if (!id.value) {
            data.forEach((item) => {
              if (item.channelName === '企业微信') {
                formData.value.config.push(wechatTemplateConfig.value);
              } else if (item.channelName === '企业微信-群机器人') {
                formData.value.config.push(wechatBotTemplateConfig.value);
              } else if (item.channelName === '邮箱') {
                formData.value.config.push(emailTemplateConfig.value);
              } else if (item.channelName === '阿里云短信') {
                formData.value.config.push(aliyunSmsTemplateConfig.value);
              } else if (item.channelName === '腾讯云短信') {
                formData.value.config.push(tencentSmsTemplateConfig.value);
              } else if (item.channelName === '钉钉') {
                formData.value.config.push(dingtalkTemplateConfig.value);
              } else if (item.channelName === '飞书') {
                formData.value.config.push(feishuTemplateConfig.value);
              } else if (item.channelName === '亿美软通电话') {
                formData.value.config.push(ymrtTemplateConfig.value);
              }
            });
          } else {
            data.forEach((item) => {
              if (
                item.channelName === '企业微信' &&
                !editChannels.value.includes('企业微信')
              ) {
                formData.value.config.push(wechatTemplateConfig.value);
              } else if (
                item.channelName === '企业微信-群机器人' &&
                !editChannels.value.includes('企业微信-群机器人')
              ) {
                formData.value.config.push(wechatBotTemplateConfig.value);
              } else if (
                item.channelName === '邮箱' &&
                !editChannels.value.includes('邮箱')
              ) {
                formData.value.config.push(emailTemplateConfig.value);
              } else if (
                item.channelName === '阿里云短信' &&
                !editChannels.value.includes('阿里云短信')
              ) {
                formData.value.config.push(aliyunSmsTemplateConfig.value);
              } else if (
                item.channelName === '腾讯云短信' &&
                !editChannels.value.includes('腾讯云短信')
              ) {
                formData.value.config.push(tencentSmsTemplateConfig.value);
              } else if (
                item.channelName === '钉钉' &&
                !editChannels.value.includes('钉钉')
              ) {
                formData.value.config.push(dingtalkTemplateConfig.value);
              } else if (
                item.channelName === '飞书' &&
                !editChannels.value.includes('飞书')
              ) {
                formData.value.config.push(feishuTemplateConfig.value);
              } else if (
                item.channelName === '亿美软通电话' &&
                !editChannels.value.includes('亿美软通电话')
              ) {
                formData.value.config.push(ymrtTemplateConfig.value);
              }
            });
          }
        }
      }
    }
  };

  onMounted(async () => {
    if (id.value) {
      title.value = 'menu.template.edit';
      const { data } = await getTemplateById(id.value);
      if (data) {
        formData.value = { ...(data as unknown as TemplateRecord) } as any;
        formData.value.config.forEach((item: any) => {
          editChannels.value.push(item.type);
          if (item.type === '企业微信') {
            wechatTemplateConfig.value = item as WechatTemplateConfig;
          } else if (item.type === '企业微信-群机器人') {
            wechatBotTemplateConfig.value = item as WechatBotTemplateConfig;
          } else if (item.type === '邮箱') {
            emailTemplateConfig.value = item as EmailTemplateConfig;
          } else if (item.type === '阿里云短信') {
            aliyunSmsTemplateConfig.value = item as AliyunSmsTemplateConfig;
          } else if (item.type === '腾讯云短信') {
            tencentSmsTemplateConfig.value = item as TencentSmsTemplateConfig;
          } else if (item.type === '钉钉') {
            dingtalkTemplateConfig.value = item as DingTalkTemplateConfig;
          } else if (item.type === '飞书') {
            feishuTemplateConfig.value = item as FeiShuTemplateConfig;
          } else if (item.type === '亿美软通电话') {
            ymrtTemplateConfig.value = item as PhoneTemplateConfig;
          }
        });
        getChannels();
      }
    }
  });
</script>

<script lang="ts">
  export default {
    name: 'Group',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 40px 20px;
    overflow: hidden;
  }

  .actions {
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;
    height: 60px;
    padding: 14px 20px 14px 0;
    background: var(--color-bg-2);
    text-align: right;
  }
</style>
