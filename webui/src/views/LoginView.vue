<script>
export default {
  data() {
    return {
      username: '',
      errormsg: null,
      loading: false,
    }
  },
  methods: {
    async login() {
      this.loading = true;
      this.errormsg = null;
      try {
        const response = await this.$axios.post(__API_URL__ + '/login', { name: this.username });
        localStorage.setItem('token', response.data.identifier);

        console.log('Login avvenuto con successo!', response.data);
        this.$router.push('/home');

      } catch (error) {
        console.error('Errore di login', error);
        this.errormsg = error.response?.data?.message || '⚠ Errore durante il login ⚠';
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>

<template>
    <div class="container">
      <div class="login">
        <h1>Login</h1>
        <input v-model="username" type="text" placeholder="Username" class="login-input" />
        <button :disabled="loading" @click="login()" type="button" class="login-button">
        <span v-if="loading">Loading...</span>
        <span v-else>Login</span>
      </button>
      <p v-if="errormsg" class="error">{{ errormsg }}</p>
      </div>
    </div>
  </template>
  
  <style>
  body, html {
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
    background-color: rgb(210, 180, 140);
  }
  </style>

  <style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: rgb(210, 180, 140);
}

.login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: rgba(255, 250, 240, 0.8);
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  text-align: center;
  min-width: 300px;
  min-height: 300px;
}

h1 {
  margin-bottom: 2rem;
  font-size: 2rem;
  color: rgb(60, 50, 40);
}
  
.login-button {
  background-color: rgb(217, 128, 91);
  color: white;
  border: none;
  padding: 0.75rem 2rem;
  border-radius: 25px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.login-button:hover {
    background-color: rgb(180, 90, 58);
}

.login-input {
  width: 100%;
  padding: 0.75rem 1rem;
  margin-bottom: 1.5rem;
  border: 1px solid #ccc;
  border-radius: 25px;
  background: rgb(255, 250, 240);
  font-size: 1rem;
  color: rgb(60, 50, 40);
  outline: none;
  transition: border-color 0.3s ease;
}

.login-input:focus {
  border-color: rgb(217, 128, 91);
}

.error {
  color: rgb(180, 90, 58);
  font-weight: bold;
  padding: 0.75rem 1rem;
  border-radius: 25px;
  margin-top: 1rem;
  font-size: 1rem;
  text-align: center;
  width: 100%;
  box-sizing: border-box;
  animation: shake 0.5s ease;
}

@keyframes shake {
  0% { transform: translateX(0); }
  25% { transform: translateX(-5px); }
  50% { transform: translateX(5px); }
  75% { transform: translateX(-5px); }
  100% { transform: translateX(0); }
}

</style>
  