<template>
	<section class="search sectionMain">
		<div class="sectionHeader">
			<h2 class="sectionHeader__title">Поиск</h2>
			<p
				class="sectionHeader__subtitle"
			>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facere aut vitae sed. Saepe, ab porro quibusdam impedit explicabo beatae vel?</p>
		</div>

		<search-form @search-submitted="handleSearch"></search-form>
	</section>

	<section class="books sectionMain">
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
import SearchForm from "../components/SearchForm.vue"
import { onBeforeMount, reactive, ref } from 'vue'
import { GET } from '../HTTP';
import { Book } from '../types';
import { useRouter } from "vue-router";

const randomBooks = ref<Book[]>([])
const router = useRouter()

onBeforeMount(async () => {
	try {
		const res = await GET<Book[]>("api/book/random3");
		randomBooks.value = res

	} catch (err) {
		console.error(err);
	}
})

const handleSearch = async (searchBy: string, searchQuery: string) => {
	try {
		router.push({
			name: "Books",
			query: {
				searchBy: searchBy,
				searchQuery: searchQuery,
			}
		})

	} catch (err) {
		console.error(err);
	}
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
</style>
