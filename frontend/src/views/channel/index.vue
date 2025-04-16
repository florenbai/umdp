<template>
  <div class="container">
    <Breadcrumb :items="['menu.channel', 'menu.channel.list']" />
    <a-card class="general-card" title="渠道列表">
      <a-row>
        <a-col :flex="1">
          <a-form
            :model="formModel"
            :label-col-props="{ span: 6 }"
            :wrapper-col-props="{ span: 18 }"
            label-align="left"
          >
            <a-row :gutter="16">
              <a-col :span="8">
                <a-form-item field="channel" label="渠道名称">
                  <a-input
                    v-model="formModel.channelName"
                    placeholder="请输入渠道名称"
                  />
                </a-form-item>
              </a-col>
              <a-col :span="8">
                <a-form-item field="channelTag" label="渠道标识">
                  <a-input
                    v-model="formModel.channelTag"
                    placeholder="请输入渠道标识"
                  />
                </a-form-item>
              </a-col>
              <a-col :span="8">
                <a-form-item field="channelStatus" label="状态">
                  <channelStatusSelect
                    v-model="formModel.channelStatus"
                  ></channelStatusSelect>
                </a-form-item>
              </a-col>
            </a-row>
          </a-form>
        </a-col>
        <a-divider style="height: 84px" direction="vertical" />
        <a-col :flex="'86px'" style="text-align: right">
          <a-space direction="vertical" :size="18">
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              查询
            </a-button>
            <a-button @click="reset">
              <template #icon>
                <icon-refresh />
              </template>
              重置
            </a-button>
          </a-space>
        </a-col>
      </a-row>
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="gotoAddChannel">
              <template #icon>
                <icon-plus />
              </template>
              新增渠道
            </a-button>
          </a-space>
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :pagination="pagination"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
        @page-change="onPageChange"
      >
        <template #index="{ rowIndex }">
          {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
        </template>
        <template #channelStatus="{ record }">
          <span v-if="record.channelStatus === 0" class="circle"></span>
          <span v-else class="circle pass"></span>
          {{ $t(`channelTable.statusType.status.${record.channelStatus}`) }}
        </template>
        <template #operations="{ record }">
          <a-space>
            <a-button
              type="text"
              size="small"
              @click="gotoEditChannel(record.id)"
            >
              编辑
            </a-button>
            <a-popconfirm
              content="是否确认要删除渠道?"
              @ok="onDelete(record.id)"
            >
              <a-button type="text" size="small"> 删除 </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, inject } from 'vue';
  import useLoading from '@/hooks/loading';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import {
    queryChannelList,
    ChannelRecord,
    ChannelParams,
    deleteChannel,
  } from '@/api/channel';
  import { Message } from '@arco-design/web-vue';
  import channelStatusSelect from './components/channelStatusSelect.vue';
  import { gotoAddChannel, gotoEditChannel } from './router';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  const size = ref<SizeProps>('medium');
  type Column = TableColumnData & { checked?: true };
  const reload = inject('reload') as any;
  const generateFormModel = () => {
    return {
      channelName: '',
      channelTag: '',
      channelStatus: '',
    };
  };
  const { loading, setLoading } = useLoading(true);
  const renderData = ref<ChannelRecord[]>([]);
  const formModel = ref(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);

  const basePagination: Pagination = {
    current: 1,
    pageSize: 20,
  };
  const pagination = reactive({
    ...basePagination,
  });

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '编号',
      dataIndex: 'id',
    },
    {
      title: '渠道名称',
      dataIndex: 'channelName',
    },
    {
      title: '渠道标识',
      dataIndex: 'channelTag',
    },
    {
      title: '状态',
      dataIndex: 'channelStatus',
      slotName: 'channelStatus',
    },
    {
      title: '操作',
      dataIndex: 'operations',
      slotName: 'operations',
    },
  ]);

  const fetchData = async (
    params: ChannelParams = { current: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryChannelList(params);
      renderData.value = data.list;
      pagination.current = params.current;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const search = () => {
    fetchData({
      ...basePagination,
      ...formModel.value,
    } as unknown as ChannelParams);
  };
  const onPageChange = (current: number) => {
    fetchData({ ...basePagination, current });
  };

  fetchData();
  const reset = () => {
    formModel.value = generateFormModel();
  };

  const onDelete = async (id: number) => {
    setLoading(true);
    try {
      await deleteChannel(id);
      Message.success({
        content: '操作成功',
        duration: 5 * 1000,
      });
      reload();
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
      cloneColumns.value.forEach((item, index) => {
        item.checked = true;
      });
      showColumns.value = cloneDeep(cloneColumns.value);
    },
    { deep: true, immediate: true }
  );
</script>

<script lang="ts">
  export default {
    name: 'SearchTable',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
  .arco-btn.arco-btn-text {
    padding: 0;
    height: auto;
  }
</style>
