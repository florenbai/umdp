<template>
  <div class="container">
    <Breadcrumb :items="['menu.channel', `${title}`]" />
    <a-form ref="formRef" layout="vertical" :model="formData">
      <a-space direction="vertical" :size="16">
        <a-card class="general-card">
          <template #title> 基础信息 </template>
          <a-row :gutter="80">
            <a-col :span="8">
              <a-form-item
                label="渠道名称"
                field="channelName"
                :rules="[
                  {
                    required: true,
                    message: '请选择渠道名称',
                  },
                ]"
              >
                <channel-select v-model="formData.channelName"></channel-select>
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item
                label="渠道标识"
                field="channelTag"
                :rules="[
                  {
                    required: true,
                    message: '请输入渠道标识',
                  },
                ]"
              >
                <a-input
                  v-model="formData.channelTag"
                  placeholder="请输入渠道标识"
                ></a-input>
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item
                label="渠道状态"
                field="channelStatus"
                :rules="[
                  {
                    required: true,
                    message: '请选择渠道状态',
                  },
                ]"
              >
                <channel-status-select
                  v-model="formData.channelStatus"
                ></channel-status-select>
              </a-form-item>
            </a-col>
          </a-row>
        </a-card>
        <wechat
          v-if="formData.channelName === '企业微信'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></wechat>
        <wechatBot
          v-if="formData.channelName === '企业微信-群机器人'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></wechatBot>
        <dingtalk
          v-else-if="formData.channelName === '钉钉'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></dingtalk>
        <feishu
          v-else-if="formData.channelName === '飞书'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></feishu>
        <email
          v-else-if="formData.channelName === '邮箱'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></email>
        <ymrt
          v-else-if="formData.channelName === '亿美软通电话'"
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></ymrt>
        <aliyun-sms
          v-else-if="
            formData.channelName === '阿里云短信' ||
            formData.channelName === '腾讯云短信'
          "
          :config="formData.channelConfig"
          @change-config="handleChangeConfig"
        ></aliyun-sms>
      </a-space>

      <div class="actions">
        <a-space>
          <a-button @click="gotoChannelList"> 返回 </a-button>
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
    ChannelForm,
    submitChannelForm,
    submitEditChannelForm,
    WechatConfig,
    EmailConfig,
    getChannelById,
    ChannelRecord,
    DingTalkConfig,
  } from '@/api/channel';
  import { Message } from '@arco-design/web-vue';
  import { useRoute } from 'vue-router';
  import { gotoChannelList } from './router';
  import channelSelect from './components/channelSelect.vue';
  import channelStatusSelect from './components/channelStatusSelect.vue';
  import wechat from './components/wechat.vue';
  import email from './components/email.vue';
  import ymrt from './components/ymrt.vue';
  import aliyunSms from './components/aliyunSms.vue';
  import dingtalk from './components/dingtalk.vue';
  import feishu from './components/feishu.vue';
  import wechatBot from './components/wechatBot.vue';

  const formData = ref<ChannelForm>({
    channelName: '企业微信',
    channelTag: '',
    channelConfig: {},
    channelStatus: 1,
  });
  const title = ref<string>('menu.channel.add');
  const id = ref<string>();
  const route = useRoute();
  const formRef = ref<FormInstance>();
  const { loading, setLoading } = useLoading();
  const onSubmitClick = async () => {
    const res = await formRef.value?.validate();
    if (!res) {
      setLoading(true);
      try {
        if (id.value) {
          await submitEditChannelForm(id.value, formData.value);
        } else {
          await submitChannelForm(formData.value);
        }
        gotoChannelList();
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

  onMounted(async () => {
    if (id.value) {
      title.value = 'menu.channel.edit';
      const { data } = await getChannelById(id.value);
      if (data) {
        formData.value = { ...(data as unknown as ChannelRecord) } as any;
      }
    }
  });

  const handleChangeConfig = (
    model: WechatConfig | EmailConfig | DingTalkConfig
  ) => {
    formData.value.channelConfig = {
      ...model,
    };
  };
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
