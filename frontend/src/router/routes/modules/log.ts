import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const LOG: AppRouteRecordRaw = {
  path: '/log',
  name: 'log',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.log',
    requiresAuth: true,
    icon: 'icon-book',
    order: 6,
  },
  children: [
    {
      path: 'list',
      name: 'log/list',
      component: () => import('@/views/log/index.vue'),
      meta: {
        locale: 'menu.log.list',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default LOG;
