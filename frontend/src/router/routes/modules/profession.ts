import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const PROFESSION: AppRouteRecordRaw = {
  path: '/profession',
  name: 'profession',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.profession',
    requiresAuth: true,
    icon: 'icon-apps',
    order: 2,
  },
  children: [
    {
      path: 'list',
      name: 'profession/list',
      component: () => import('@/views/profession/index.vue'),
      meta: {
        locale: 'menu.profession.list',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'add',
      name: 'profession/add',
      component: () => import('@/views/profession/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
    {
      path: 'edit/:id',
      name: 'profession/edit',
      component: () => import('@/views/profession/form.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
  ],
};

export default PROFESSION;
