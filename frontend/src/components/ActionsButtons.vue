<template>
	<!-- Buttons to add: 
			1. Add new item – Always
			2. Edit – If selected only 1 item 
	3. Delete – If selected at least 1 item-->

	<div class="actions">
		<transition name="fade">
			<button @click="emit('addPressed')" v-show="canAdd" class="actions__btn actions__add">Создать</button>
		</transition>

		<transition name="fade">
			<button
				@click="emit('editPressed')"
				v-show="canEdit"
				class="actions__btn actions__edit"
			>Редактировать</button>
		</transition>

		<transition name="fade">
			<button
				@click="emit('deletePressed')"
				v-show="canDelete"
				class="actions__btn actions__delete"
			>Удалить</button>
		</transition>
	</div>
</template>

<script lang="ts" setup>
const emit = defineEmits<{
	(e: "addPressed"): void
	(e: "editPressed"): void
	(e: "deletePressed"): void
}>()

const props = defineProps({
	canAdd: {
		type: Boolean,
		default: true,
	},
	canEdit: {
		type: Boolean,
		required: true,
	},
	canDelete: {
		type: Boolean,
		required: true,
	}
})
</script>

<style>
.actions {
	display: flex;
	align-items: center;
}

.actions__btn {
	box-sizing: border-box;
	margin-right: var(--offset-half);
	padding: var(--offset-quarter);

	cursor: pointer;

	border: 2px solid rgb(var(--greyblue));
	border-radius: calc(var(--offset-quarter) / 2);

	background-color: rgb(var(--white));

	font-weight: bold;
	color: rgb(var(--greyblue));

	transition: var(--main-transition);
}

.actions__btn:last-of-type {
	margin: 0;
}

.actions__add:hover {
	border-color: rgb(var(--greyblue));
	background-color: rgb(var(--greyblue));
	color: rgb(var(--white));
}

.actions__edit:hover {
	border-color: rgb(var(--sapphire));
	background-color: rgb(var(--sapphire));
	color: rgb(var(--white));
}

.actions__delete:hover {
	border-color: rgb(var(--redsand));
	background-color: rgb(var(--redsand));
	color: rgb(var(--white));
}
</style>