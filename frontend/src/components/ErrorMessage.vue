<template>
	<transition name="fade">
		<div v-if="isError" class="error-message">{{ errMessage }}</div>
	</transition>
</template>

<script lang="ts" setup>
const props = defineProps({
	isError: {
		type: Boolean,
		required: true,
	},
	errMessage: {
		type: String,
		required: true
	}
})
</script>

<style>
.error-message {
	padding: var(--offset-quarter);

	box-shadow: var(--second-shadow);
	border-radius: calc(var(--offset-quarter) / 2);
	border: 2px solid rgb(var(--redsand));

	background-color: rgb(var(--redsand));
	color: rgb(var(--white));
	font-weight: bold;
}

.fade-enter-active,
.fade-leave-active {
	transition: all 0.2s ease-in-out;
}

.fade-enter-from,
.fade-leave-to {
	opacity: 0;
}
</style>