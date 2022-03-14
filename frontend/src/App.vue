<script setup lang="ts">
import { ref, watch, onBeforeMount, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import { useStore } from "./store/index";
import Loader from "./components/Loader.vue";

const route = useRoute()
const store = useStore();

const isNotHomePage = computed(() => {
	return route.fullPath !== "/";
})

const isLoaded = ref(false)

// onBeforeMount(async () => {
// 	isLoaded.value = false;

// 	// Load all data from API
// 	await Promise.all([
// 		store.loadAuthors(),
// 		store.loadGenres(),
// 		store.loadBooks(),
// 		store.loadBooksAuthors()
// 	]);

// 	console.log("initial load");

// 	isLoaded.value = true;
// });

// Reload data when page changed
watch(() => route.params, async () => {
	isLoaded.value = false;

	// Reset error status and message
	store.setErrorWithMessage(false)

	// Load new all data from API
	await Promise.all([
		store.loadAuthors(),
		store.loadGenres(),
		store.loadBooks(),
		store.loadBooksAuthors()
	]);

	// Disable loader
	setTimeout(() => {
		isLoaded.value = true;
	}, 1000)
})
</script>

<template>
	<loader v-show="!isLoaded" />

	<header class="app-header">
		<div class="header wrapper">
			<div class="header__intro">
				<h1 class="header__title">Лабораторная работа №1</h1>
				<h3 class="header__subtitle">
					По дисциплине "Программирование и администрирование в
					среде"
				</h3>
			</div>

			<nav class="header__nav">
				<router-link to="/">Главная</router-link>
				<router-link to="/books">Книги</router-link>
				<router-link to="/genres">Жанры</router-link>
				<router-link to="/authors">Авторы</router-link>
			</nav>
		</div>
	</header>

	<main class="app-main wrapper" :class="{ 'grid-main': isNotHomePage }">
		<router-view></router-view>
	</main>

	<footer class="app-footer">
		<div class="footer wrapper">
			<div class="footer__credits">
				<p>
					Выполнил студент группы 15.11Д-МО12/19б факультета
					ИМИСиЦЭ Козлов Роман
				</p>
			</div>
			<div class="footer__link">
				<a
					href="https://github.com/the-NZA/DB_Lab1"
					target="_blank"
					title="Откроется в новой вкладке"
				>Репозиторий</a>
			</div>
		</div>
	</footer>
</template>

<style>
#app {
	min-height: 100vh;
	display: grid;
	grid-template-rows: max-content 1fr max-content;
	gap: var(--offset);
}

.app-header {
	background-color: rgb(var(--leanen));
}

.app-main {
	background-color: rgb(var(--white));
	width: 100%;
	padding: 0 var(--offset);
}

.app-footer {
	background-color: rgb(var(--greyblue));
}

.app-footer--margin-top {
	margin-top: var(--offset);
}

.grid-main {
	display: grid;
	grid-template-rows: max-content 1fr;
	gap: var(--offset);
}
</style>
