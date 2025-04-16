import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const CHANNEL: AppRouteRecordRaw = {
  path: '/channel',
  name: 'channel',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.channel',
    requiresAuth: true,
    icon: 'icon-mosaic',
    order: 1,
  },
  children: [
    {
      path: 'list',
      name: 'channel/list',
      component: () => import('@/views/channel/index.vue'),
      meta: {
        locale: 'menu.channel.list',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'add',
      name: 'channel/add',
      component: () => import('@/views/channel/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
    {
      path: 'edit/:id',
      name: 'channel/edit',
      component: () => import('@/views/channel/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
  ],
};

export default CHANNEL;
