<template>
  <div class="container">
    <Breadcrumb :items="['menu.profession', `${title}`]" />
    <a-form ref="formRef" layout="vertical" :model="formData">
      <a-space direction="vertical" :size="16">
        <a-card class="general-card">
          <template #title> 基础信息 </template>
          <a-row :gutter="80">
            <a-col :span="12">
              <a-form-item
                label="业务名称"
                field="professionName"
                :rules="[
                  {
                    required: true,
                    message: '请输入业务名称',
                  },
                ]"
              >
                <a-input
                  v-model="formData.professionName"
                  placeholder="请输入业务名称"
                ></a-input>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item
                label="验证TOKEN"
                field="token"
                :rules="[
                  {
                    required: true,
                    message: '请输入验证TOKEN',
                  },
                ]"
              >
                <a-input v-model="formData.token" placeholder="请输入验证TOKEN">
                  <template #suffix>
                    <icon-refresh @click="generateRandomToken" /> </template
                ></a-input>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item
                label="开通渠道"
                field="channels"
                :rules="[
                  {
                    required: true,
                    message: '请选择渠道',
                  },
                ]"
              >
                <channel-name-select
                  v-model="formData.channels"
                  :selected="formData.channels"
                ></channel-name-select>
              </a-form-item>
            </a-col>
          </a-row>
        </a-card>
      </a-space>

      <div class="actions">
        <a-space>
          <a-button @click="gotoProfessionList"> 返回 </a-button>
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
    ProfessionForm,
    submitProfessionForm,
    submitEditProfessionForm,
    getProfessionById,
    ProfessionRecord,
  } from '@/api/profession';
  import { Message } from '@arco-design/web-vue';
  import { useRoute } from 'vue-router';
  import { gotoProfessionList } from './router';
  import channelNameSelect from './components/channelNameSelect.vue';

  const formData = ref<ProfessionForm>({
    professionName: '',
    token: '',
    channels: [],
  });
  const title = ref<string>('menu.profession.add');
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
          await submitEditProfessionForm(id.value, formData.value);
        } else {
          await submitProfessionForm(formData.value);
        }
        gotoProfessionList();
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
      title.value = 'menu.profession.edit';
      const { data } = await getProfessionById(id.value);
      if (data) {
        formData.value = { ...(data as unknown as ProfessionRecord) } as any;
      }
    }
  });

  const generateRandomToken = () => {
    const buffer = crypto.randomUUID();
    formData.value.token = buffer;
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
