<template>
  <div class="sidebar">
    <h2>管理菜单</h2>
    <ul>
      <li><a href="#users" @click="setActive('users')">用户管理</a></li>
      <li><a href="#products" @click="setActive('products')">产品管理</a></li>
      <li><a href="#orders" @click="setActive('orders')">订单管理</a></li>
    </ul>
  </div>
</template>

<script setup>
import {ref} from 'vue';

const activeMenu = ref('users');
const emit = defineEmits(['menu-change']);

const setActive = (menu) => {
  activeMenu.value = menu;
  emit('menu-change', menu);
};
</script>

<style scoped>
.sidebar {
  width: 200px;
  background-color: #f0f0f0;
  padding: 20px;
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar li {
  margin-bottom: 10px;
}

.sidebar a {
  text-decoration: none;
  color: #333;
}
</style>