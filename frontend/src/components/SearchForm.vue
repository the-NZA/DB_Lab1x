<template>
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
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const emit = defineEmits<{
	(e: "searchSubmitted", searchBy: string, searchQuery: string): void
}>()

const searchBy = ref("")
const searchQuery = ref("")

const onFormSubmit = () => {
	emit('searchSubmitted', searchBy.value, searchQuery.value)
}
</script>

<style scoped>
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
</style>