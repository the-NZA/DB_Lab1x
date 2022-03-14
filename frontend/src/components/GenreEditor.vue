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
					v-model="currentGenre.title"
					type="text"
					placeholder="Введите название жанра"
					required
					@input="store.setErrorWithMessage(false)"
				/>
			</div>

			<div class="edfields__field">
				<label class="edfields__label" for="edsnippet">Описание</label>
				<textarea
					class="edfields__input edfields__textarea"
					id="edsnippet"
					v-model="currentGenre.snippet"
					placeholder="Введите описание жанра"
					@input="store.setErrorWithMessage(false)"
				></textarea>
			</div>
		</div>
		<div class="editor__footer">
			<button @click="saveGenre">{{ buttonText }}</button>
			<error-message :is-error="store.getIsError" :err-message="store.getErrMessage" />
		</div>
	</div>
</template>

<script lang="ts" setup>
import { ref, onBeforeMount, reactive, computed } from 'vue';
import ErrorMessage from './ErrorMessage.vue';
import { useStore } from "../store"
import { Genre } from '../types';
import { GenreEditorTitle, SaveButtonValue } from '../types/enums';

const props = defineProps({
	genre_id: {
		type: String,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

const store = useStore()
const currentGenre = reactive<Genre>({
	id: "",
	title: "",
	snippet: "",
	deleted: false,
})

const title = ref<GenreEditorTitle>(GenreEditorTitle.Create)
const buttonText = ref<SaveButtonValue>(SaveButtonValue.Save)

onBeforeMount(() => {
	if (props.genre_id) {
		const genre = store.getGenreByID(props.genre_id);
		if (genre) {
			title.value = GenreEditorTitle.Edit
			buttonText.value = SaveButtonValue.Update

			currentGenre.id = genre.id
			currentGenre.deleted = genre.deleted
			currentGenre.title = genre.title
			currentGenre.snippet = genre.snippet
		}
	}
})

const saveGenre = async () => {
	// Check required status for current genre
	if (!currentGenre.title || currentGenre.title.length < 1) {
		store.setErrorWithMessage(true, "Название жанра обязательно для заполнения")
		return
	}

	if (title.value === GenreEditorTitle.Create) {
		// If create new genre
		await store.addGenre(currentGenre)
	} else {
		// Update existing 
		await store.updateGenre(currentGenre)
	}

	emit('savePressed')
}

</script>

<style>
</style>