<template>
	<div class="login">
		<h2 class="login__title">{{ pageTitle }}</h2>

		<form @submit.prevent="onFormSubmit" class="login__form">
			<div class="login__formgroup">
				<label for="username">Имя пользователя</label>
				<input type="text" id="username" required />
			</div>

			<div class="login__formgroup">
				<label for="password">Пароль</label>
				<input type="password" id="password" required />
			</div>

			<div v-if="isSignUp" class="login__formgroup">
				<label for="email">Email(Опционально)</label>
				<input id="email" type="email" />
			</div>

			<button class="login__button">Войти</button>
		</form>
	</div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute()

const isSignUp = computed(() => {
	return route.fullPath === "/signup"
})

const pageTitle = computed(() => {
	return isSignUp.value ? "Регистрация" : "Авторизация"
})

const onFormSubmit = (e: Event) => {
	console.log("submited");
}
</script>

<style scoped>
.login {
	border: 2px solid rgb(var(--redsand));
	border-radius: var(--offset-quarter);

	padding: var(--offset);
}

.login__title {
	margin: 0 0 var(--offset) 0;
	text-align: center;
	padding-bottom: var(--offset);
	border-bottom: 2px solid rgb(var(--redsand));
}
.login__form {
	display: flex;
	flex-direction: column;
	min-width: 400px;
}

.login__formgroup {
	display: flex;
	flex-direction: column;
	margin-bottom: var(--offset);
}

.login__formgroup label {
	margin-bottom: var(--offset-quarter);
	font-weight: bold;
	font-size: 1.1rem;
}

.login__formgroup input {
	padding: var(--offset-quarter);
}

.login__button {
	background-color: rgb(var(--greyblue));
	border: none;
	padding: 15px;
	color: rgb(var(--white));
	font-weight: bold;
	font-size: 1.1rem;
	cursor: pointer;
	transition: var(--main-transition);
}

.login__button:hover {
	background-color: rgb(var(--sapphire));
}
</style>
