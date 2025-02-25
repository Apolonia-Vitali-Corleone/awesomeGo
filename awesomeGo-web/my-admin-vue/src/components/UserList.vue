<template>
  <div class="content">
    <h2>用户管理</h2>
    <table>
      <thead>
      <tr>
        <th>ID</th>
        <th>用户名</th>
        <th>邮箱</th>
        <th>操作</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="user in users" :key="user.id">
        <td>{{ user.id }}</td>
        <td>{{ user.username }}</td>
        <td>{{ user.email }}</td>
        <td class="action-buttons">
          <button>修改</button>
          <button>删除</button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import axios from 'axios';

const users = ref([]);

onMounted(async () => {
  try {
    const response = await axios.get('/api/users.json');
    users.value = response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
  }
});
</script>

<style scoped>
.content {
  flex: 1;
  padding: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

.action-buttons button {
  margin-right: 5px;
}
</style>