<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

// 定义响应式数据
const username = ref('');
const password = ref('');
const error = ref('');

// 获取路由实例
const router = useRouter();

// 定义登录方法
const handleLogin = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name: username.value,
        password: password.value
      })
    });

    const data = await response.json();
    if (data.error) {
      error.value = data.error;
      return;
    }

    localStorage.setItem('token', data.token);
    router.push('/tasks');
  } catch (error) {
    console.error('登录请求失败:', error);
    error.value = '网络连接失败，请重试';
  }
};
</script>

<template>
  <div class="container">
    <h2>任务管理系统</h2>
    <div class="login-box">
      <h3>登录</h3>
      <div v-if="error" class="error">{{ error }}</div>

      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>用户名</label>
          <input
              type="text"
              v-model="username"
              placeholder="请输入用户名"
              required
          />
        </div>

        <div class="form-group">
          <label>密码</label>
          <input
              type="password"
              v-model="password"
              placeholder="请输入密码"
              required
          />
        </div>

        <button type="submit" class="btn-primary">登录</button>
      </form>

      <p class="mt-2">
        还没有账号？
        <router-link to="/register">立即注册</router-link>
      </p>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
}

.login-box {
  background: #f0f2f5;
  padding: 20px;
  border-radius: 8px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

button {
  width: 100%;
  padding: 12px;
  background: #1a73e8;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  margin-top: 10px;
}

button:hover {
  background: #1557b0;
}

.error {
  color: red;
  text-align: center;
  margin: 10px 0;
  padding: 8px;
  background: #ffebee;
  border-radius: 4px;
}

.router-link-exact-active {
  color: #1a73e8;
  font-weight: bold;
}
</style>