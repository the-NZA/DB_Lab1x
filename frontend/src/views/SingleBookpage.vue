<template>
	<div class="book">
		<h2 class="book__title">{{ book?.title }}</h2>

		<p class="book__snippet">{{ book?.snippet }}</p>

		<div class="book__info">
			<div class="book__pages_cnt">
				<span>Страниц:</span>
				{{ book?.pages_cnt }}
			</div>

			<div class="book__pub_year">
				<span>Год публикации:</span>
				{{ book?.pub_year }}
			</div>

			<div class="book__genres">
				<span>Жанр:</span>
				{{ bookGenre?.title }}
			</div>
		</div>

		<div class="book__download">
			<h4 class="book__downheader">Ссылки для скачивания</h4>

			<ol class="book__downlist">
				<li class="book__downitem">
					<a href="#" download>download</a>
				</li>
				<li class="book__downitem">
					<a href="#" download>download</a>
				</li>
				<li class="book__downitem">
					<a href="#" download>download</a>
				</li>
			</ol>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount } from "vue";
import { useStore } from "../store";
import { useRoute, useRouter } from "vue-router";
import { Book, Genre } from "../types"
import { GET } from "../HTTP"

const store = useStore()
const route = useRoute()
const book = ref<Book>()
const bookGenre = ref<Genre>()

onBeforeMount(async () => {
	try {
		const res = await GET<Book>(`api/book/${route.params.bookId}`);
		book.value = res

		const resGenre = await GET<Genre>(`api/genre/${book.value.genre_id}`);
		bookGenre.value = resGenre
	} catch (err) {
		console.error(err);

	}
})

</script>

<style scoped>
.book {
	border: 2px solid rgb(var(--greyblue));
	border-radius: calc(var(--offset-quarter) / 2);
	padding: calc(var(--offset) * 2);

	display: grid;
	grid-template-rows: repeat(4, min-content);
	gap: calc(var(--offset) * 2);
}

.book__title {
	margin: 0;
	text-align: center;
	text-transform: uppercase;
	font-weight: 900;
	font-size: 2rem;
}

.book__snippet {
	margin: 0;
	font-weight: 500;
}

.book__info {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	justify-items: center;

	border-top: 1px solid #000;
	padding-top: calc(var(--offset) * 2);
}

.book__info span {
	font-weight: bold;
}

.book__download {
	border-top: 1px solid #000;
	padding-top: calc(var(--offset) * 2);

	display: flex;
	flex-direction: column;
	align-items: center;
}

.book__downheader {
	text-transform: uppercase;
	margin: 0 0 var(--offset-half) 0;
}

.book__downlist {
	padding: 0;
	margin: 0;
	list-style: decimal;
}

.book__downitem:not(:last-of-type) {
	margin-bottom: var(--offset-half);
}

.book__downitem::marker {
	font-weight: bold;
}
</style>