import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const TEMPLATE: AppRouteRecordRaw = {
  path: '/template',
  name: 'template',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.template',
    requiresAuth: true,
    icon: 'icon-list',
    order: 3,
  },
  children: [
    {
      path: 'list',
      name: 'template/list',
      component: () => import('@/views/template/index.vue'),
      meta: {
        locale: 'menu.template.list',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'add',
      name: 'template/add',
      component: () => import('@/views/template/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
    {
      path: 'edit/:id',
      name: 'template/edit',
      component: () => import('@/views/template/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
  ],
};

export default TEMPLATE;
