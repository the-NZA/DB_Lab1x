<template>
	<div class="login">
		<h2 class="login__title">{{ pageTitle }}</h2>

		<form @submit.prevent="onFormSubmit" class="login__form">
			<div class="login__formgroup">
				<label for="username">Имя пользователя</label>
				<input type="text" v-model="username" id="username" required />
			</div>

			<div class="login__formgroup">
				<label for="password">Пароль</label>
				<input type="password" v-model="password" id="password" required />
			</div>

			<div v-if="isSignUp" class="login__formgroup">
				<label for="email">Email(Опционально)</label>
				<input id="email" v-model="email" type="email" />
			</div>

			<button class="login__button">Войти</button>
		</form>
	</div>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useStore } from "../store/index";
import { AUTH } from "../HTTP";
import { User, LoginBody, RegisterBody } from "../types/index"

const router = useRouter()
const route = useRoute()
const store = useStore()

const isSignUp = computed(() => {
	return route.fullPath === "/signup"
})

const pageTitle = computed(() => {
	return isSignUp.value ? "Регистрация" : "Авторизация"
})

const username = ref<string>("")
const password = ref<string>("")
const email = ref<string>("")

const onFormSubmit = async (e: Event) => {
	if (isSignUp.value) {
		try {
			// login
			const res = await AUTH<User, RegisterBody>({
				username: username.value,
				password: password.value,
				email: email.value,
			}, "/auth/signup")

			// save and login new user
			store.setLogin(true, res)

		}
		catch (err) {
			console.error(err);
			// display error
			return
		}
	} else {
		try {
			// login
			const res = await AUTH<User, LoginBody>({
				username: username.value,
				password: password.value,
			}, "/auth/login")

			// save logined user
			store.setLogin(true, res)
		}
		catch (err) {
			console.error(err);
			// display error
			return
		}
	}

	// go to home page
	router.push({
		name: "Home"
	})
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
