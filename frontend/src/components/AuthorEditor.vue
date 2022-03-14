<template>
	<div class="editor">
		<div class="editor__header">
			<h2>{{ title }}</h2>
			<button @click="emit('closePressed')"></button>
		</div>
		<div class="editor__body edfields">
			<div class="edfields__field">
				<label class="edfields__label" for="edfirstname">Имя</label>
				<input
					class="edfields__input"
					id="edfirstname"
					v-model="currentAuthor.firstname"
					type="text"
					placeholder="Введите имя автора"
					@input="store.setErrorWithMessage(false)"
					required
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edlastname">Фамилия</label>
				<input
					class="edfields__input"
					id="edlastname"
					v-model="currentAuthor.lastname"
					type="text"
					placeholder="Введите фамилию автора"
					@input="store.setErrorWithMessage(false)"
					required
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsurname">Отчество</label>
				<input
					class="edfields__input"
					id="edsurname"
					v-model="currentAuthor.surname"
					type="text"
					placeholder="Введите отчество автора"
					@input="store.setErrorWithMessage(false)"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edbooks">Книги</label>
				<multiselect
					v-model="selectedBooks"
					:options="store.getBooks"
					track-by="id"
					label="title"
					:multiple="true"
					:allow-empty="true"
					:searcheble="true"
					:close-on-select="false"
					deselectLabel
					selectLabel
					selectedLabel="Выбранный"
					placeholder="Выберите книги"
					@input="store.setErrorWithMessage(false)"
					id="edbooks"
				>
					<template v-slot:noOptions>Список книг пуст</template>
					<template v-slot:noResult>Совпадений не найдено</template>
				</multiselect>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsnippet">Описание</label>
				<textarea
					class="edfields__input edfields__textarea"
					id="edsnippet"
					v-model="currentAuthor.snippet"
					placeholder="Введите описание автора"
					@input="store.setErrorWithMessage(false)"
				></textarea>
			</div>
		</div>
		<div class="editor__footer">
			<button @click="saveAuthor">{{ buttonText }}</button>
			<error-message :is-error="store.getIsError" :err-message="store.getErrMessage" />
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive } from 'vue';
import Multiselect from 'vue-multiselect'
import ErrorMessage from './ErrorMessage.vue';
import { useStore } from "../store"
import { Author, Book } from '../types';
import { AuthorEditorTitle, SaveButtonValue } from '../types/enums';

const props = defineProps({
	author_id: {
		type: String,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

const store = useStore()

const selectedBooks = ref<Book[]>([])
const currentAuthor = reactive<Author>({
	id: "",
	firstname: "",
	lastname: "",
	surname: "",
	snippet: "",
	deleted: false,
})

const title = ref<AuthorEditorTitle>(AuthorEditorTitle.Create)
const buttonText = ref<SaveButtonValue>(SaveButtonValue.Save)

onBeforeMount(() => {
	if (props.author_id) {
		const author = store.getAuthorByID(props.author_id);
		if (author) {
			title.value = AuthorEditorTitle.Edit
			buttonText.value = SaveButtonValue.Update

			currentAuthor.id = author.id
			currentAuthor.firstname = author.firstname
			currentAuthor.lastname = author.lastname
			currentAuthor.surname = author.surname
			currentAuthor.snippet = author.snippet

			selectedBooks.value = store.getBooksByAuthorID(author.id)
		}

	}
})

const saveAuthor = async () => {
	// Check required status for current author
	if (
		!currentAuthor.lastname ||
		!currentAuthor.firstname ||
		currentAuthor.lastname.length < 1 ||
		currentAuthor.firstname.length < 1
	) {
		store.setErrorWithMessage(true, "Имя и фамилия обязательны для заполнения")
		return
	}

	const booksIDs: string[] = []
	selectedBooks.value.forEach(book => {
		booksIDs.push(book.id)
	})

	if (title.value === AuthorEditorTitle.Create) {
		// If create new genre
		await store.addAuthor({
			author: currentAuthor,
			books_ids: booksIDs,
		})
	} else {
		// Update existing 
		await store.updateAuthor({
			author: currentAuthor,
			books_ids: booksIDs,
		})
	}

	emit('savePressed')
}

</script>

<style>
</style>