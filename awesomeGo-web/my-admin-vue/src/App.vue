<template>
  <div id="app">
    <Sidebar @menu-change="handleMenuChange" />
    <UserList v-if="activeMenu === 'users'" />
    <div v-if="activeMenu === 'products'" class="content">产品管理内容</div>
    <div v-if="activeMenu === 'orders'" class="content">订单管理内容</div>
  </div>
</template>

<script setup>
import Sidebar from './components/Sidebar.vue';
import UserList from './components/UserList.vue';
import { ref } from 'vue';

const activeMenu = ref('users');

const handleMenuChange = (menu) => {
  activeMenu.value = menu;
};
</script>

<style scoped>
#app {
  font-family: sans-serif;
  margin: 0;
  padding: 0;
  display: flex;
}

.content {
  flex: 1;
  padding: 20px;
}
</style>