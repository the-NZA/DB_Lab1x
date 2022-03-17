<template>
	<section class="search">
		<div class="sectionHeader">
			<h2 class="sectionHeader__title">Поиск</h2>
			<p
				class="sectionHeader__subtitle"
			>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facere aut vitae sed. Saepe, ab porro quibusdam impedit explicabo beatae vel?</p>
		</div>

		<form @submit.prevent="onFormSubmit" class="search__form">
			<div class="search__formgroup">
				<label for="fselect">Поиск по</label>
				<select v-model="searchBy" id="fselect" required>
					<option selected disabled value>Выберите опцию</option>
					<option value="title">названию</option>
					<option value="author">автору</option>
					<option value="genre">жанру</option>
				</select>
			</div>

			<div class="search__formgroup">
				<label for="fsearch">Запрос</label>
				<input
					v-model.trim="searchQuery"
					required
					type="search"
					id="fsearch"
					placeholder="Введите ваш запрос"
				/>
			</div>
			<div class="search__formgroup">
				<input type="submit" value="Искать" />
			</div>
		</form>
	</section>

	<section class="books">
		<div class="sectionHeader">
			<h2 class="sectionHeader__title">Случайные книги</h2>
			<p
				class="sectionHeader__subtitle"
			>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facere aut vitae sed. Saepe, ab porro quibusdam impedit explicabo beatae vel?</p>
		</div>

		<div class="books__cards">
			<book-card v-for="book in randomBooks" :book="book"></book-card>
		</div>
	</section>
</template>

<script lang="ts" setup>
import BookCard from "../components/BookCard.vue"
import { onBeforeMount, reactive, ref } from 'vue';
import { GET } from '../HTTP';
import { Book } from '../types';

const randomBooks = ref<Book[]>([])
const searchBy = ref("")
const searchQuery = ref("")

onBeforeMount(async () => {
	try {
		const res = await GET<Book[]>("api/book/random3");
		randomBooks.value = res

	} catch (err) {
		console.error(err);
	}
})

const onFormSubmit = () => {
	console.log("searched");
}
</script>

<style scoped>
/* h3 {
	margin: 0 0 var(--offset-half) 0;
} */

.sectionHeader {
	margin-bottom: var(--offset);
}

.sectionHeader__title {
	margin: 0 0 var(--offset-half) 0;
	text-transform: uppercase;
	font-weight: 800;
	font-size: 2rem;
}

.sectionHeader__subtitle {
	margin: 0;
	font-size: 1.2rem;
	font-weight: 500;
}

section {
	border: 4px solid rgb(var(--greyblue));
	border-radius: calc(var(--offset-half) / 2);
	padding: var(--offset);
}

.search__form {
	display: grid;
	grid-template-columns: max-content 1fr 160px;
	gap: var(--offset);
}
.search__formgroup {
	display: flex;
	flex-direction: column;
	justify-content: flex-end;
}
.search__formgroup label {
	font-weight: bold;
	margin-bottom: var(--offset-quarter);
}
.search__formgroup select,
.search__formgroup input {
	padding: var(--offset-quarter);
	border: 1px solid rgb(var(--greyblue));
	border-radius: calc(var(--offset-quarter) / 2);
}

.search__formgroup select {
	background-color: rgb(var(--white));
	font-weight: 500;
}

.search__formgroup input[type="submit"] {
	background-color: rgb(var(--greyblue));
	font-weight: bold;
	color: rgb(var(--white));
	cursor: pointer;
	transition: var(--main-transition);
}

.search__formgroup input[type="submit"]:hover {
	background-color: rgb(var(--sapphire));
}

.books__cards {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: var(--offset);
}
</style>
