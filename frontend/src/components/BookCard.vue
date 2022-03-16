<template>
	<div class="bookCard">
		<h3 class="bookCard__title">{{ book.title }}</h3>
		<p class="bookCard__snippet">{{ book.snippet }}</p>
		<button class="bookCard__button">Смотреть</button>
	</div>
</template>

<script setup lang="ts">
import { Book } from '../types';

const props = defineProps<{
	book: Book,
}>()

// const props = defineProps({
// 	book: {
// 		type: {} as Book,
// 		require: true,
// 	},
// })
</script>

<style scoped>
.bookCard {
	border: 2px solid rgb(var(--sapphire));
	border-radius: var(--offset-quarter);
	padding: var(--offset-half);
}

.bookCard__title {
	font-size: 1.6rem;
	margin: 0 0 var(--offset) 0;
}

.bookCard__snippet {
	margin: 0 0 var(--offset) 0;
}

.bookCard__button {
	display: block;
	width: 100%;
	padding: var(--offset-quarter);

	border: none;
	border-radius: calc(var(--offset-quarter) / 2);
	background-color: rgb(var(--greyblue));
	color: rgb(var(--white));
	font-weight: bold;
	cursor: pointer;
	transition: var(--main-transition);
}

.bookCard__button:hover {
	background-color: rgb(var(--sapphire));
}
</style>