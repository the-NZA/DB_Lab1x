<template>
	<div class="editor">
		<div class="editor__header">
			<h2>{{ title }}</h2>
			<button @click="emit('closePressed')"></button>
		</div>
		<div class="editor__body edfields">
			<div class="edfields__field">
				<label class="edfields__label" for="edtitle">Название</label>
				<input
					class="edfields__input"
					id="edtitle"
					v-model="currentBook.title"
					type="text"
					placeholder="Введите название книги"
					@input="store.setErrorWithMessage(false)"
					required
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edgenre">Жанр</label>
				<multiselect
					v-model="selectedGenre"
					:options="store.getGenres"
					track-by="id"
					label="title"
					:allow-empty="false"
					:searcheble="true"
					deselectLabel
					selectLabel
					selectedLabel="Выбранный"
					placeholder="Выберите жанр"
					@input="store.setErrorWithMessage(false)"
					id="edgenre"
					required
				>
					<template v-slot:noOptions>Список жанров пуст</template>
					<template v-slot:noResult>Совпадений не найдено</template>
				</multiselect>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edauthors">Авторы</label>
				<multiselect
					v-model="selectedAuthors"
					:options="store.getAuthors"
					track-by="id"
					:custom-label="customAuthorsSelectLabel"
					:multiple="true"
					:allow-empty="true"
					:searcheble="true"
					:close-on-select="false"
					deselectLabel
					selectLabel
					selectedLabel="Выбранный"
					placeholder="Выберите авторов"
					@input="store.setErrorWithMessage(false)"
					id="edauthors"
					required
				>
					<template v-slot:noOptions>Список авторов пуст</template>
					<template v-slot:noResult>Совпадений не найдено</template>
				</multiselect>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsnippet">Описание</label>
				<textarea
					class="edfields__input edfields__textarea"
					id="edsnippet"
					v-model="currentBook.snippet"
					placeholder="Введите описание автора"
					@input="store.setErrorWithMessage(false)"
				></textarea>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edpubyear">Год публикации</label>
				<input
					type="number"
					step="1"
					class="edfields__input"
					id="edpubyear"
					v-model="currentBook.pub_year"
					placeholder="Введите дату публикации"
					@input="store.setErrorWithMessage(false)"
					required
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edpagescnt">Количество страниц</label>
				<input
					type="number"
					step="1"
					class="edfields__input"
					id="edpagescnt"
					v-model="currentBook.pages_cnt"
					placeholder="Введите количество страниц"
					@input="store.setErrorWithMessage(false)"
				/>
			</div>
		</div>
		<div class="editor__footer">
			<button @click="saveBook">{{ buttonText }}</button>
			<error-message :is-error="store.getIsError" :err-message="store.getErrMessage" />
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive, computed, onMounted } from 'vue';
import Multiselect from 'vue-multiselect'
import ErrorMessage from './ErrorMessage.vue';
import { useStore } from "../store"
import { Author, Book, Genre } from '../types';
import { BookEditorTitle, SaveButtonValue } from '../types/enums';

const props = defineProps({
	book_id: {
		type: String,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

const store = useStore()

const genres = ref<Genre[]>([])
const selectedGenre = ref<Genre>()
const selectedAuthors = ref<Author[]>([])
const currentBook = reactive<Book>({
	id: "",
	title: "",
	snippet: "",
	genre_id: "",
	pages_cnt: 0,
	pub_year: 0,
	deleted: false,
})

const title = ref<BookEditorTitle>(BookEditorTitle.Create)
const buttonText = ref<SaveButtonValue>(SaveButtonValue.Save)

onBeforeMount(() => {
	// genres.value = store.getGenres

	if (props.book_id) {
		const book = store.getBookByID(props.book_id);
		if (book) {
			title.value = BookEditorTitle.Edit
			buttonText.value = SaveButtonValue.Update

			currentBook.id = book.id
			currentBook.title = book.title
			currentBook.snippet = book.snippet
			currentBook.pages_cnt = book.pages_cnt
			currentBook.pub_year = book.pub_year
			currentBook.genre_id = book.genre_id

			// Get initial value for selectedGenre
			store.getGenres.forEach(genre => {
				if (genre.id === currentBook.genre_id) {
					selectedGenre.value = genre
					return
				}
			})

			selectedAuthors.value = store.getAuthorsByBookID(book.id)
		}
	}
})

const saveBook = async () => {
	if (store.getGenres.length < 1) {
		store.setErrorWithMessage(true, "Еще не создано ни одного жанра. Начните с этого")
		return
	}

	if (store.getAuthors.length < 1) {
		store.setErrorWithMessage(true, "Еще не создано ни одного автора. Начните с этого")
		return
	}

	if (
		!currentBook.title ||
		!currentBook.pub_year ||
		!selectedGenre.value ||
		selectedAuthors.value.length < 1
	) {
		store.setErrorWithMessage(true, "Название, год публикации, жанр и авторы обязательны для заполнения")
		return
	}

	currentBook.genre_id = selectedGenre.value!.id

	const authorsIDs: string[] = []
	selectedAuthors.value.forEach(author => {
		authorsIDs.push(author.id)
	})

	if (title.value === BookEditorTitle.Create) {
		// If create new genre
		await store.addBook({
			book: currentBook,
			authors_ids: authorsIDs,
		})

	} else {
		// Update existing 
		await store.updateBook({
			book: currentBook,
			authors_ids: authorsIDs,
		})
	}

	emit('savePressed')
}

const customAuthorsSelectLabel = (author: Author): string => {
	return author.surname.length === 0
		? `${author.lastname} ${author.firstname}`
		: `${author.lastname} ${author.firstname} ${author.surname}`
}

</script>

<style>
</style>