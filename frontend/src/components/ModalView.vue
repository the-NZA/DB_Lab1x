<template>
	<teleport to="body">
		<transition name="slide" @after-enter="makeFade = true" @before-leave="makeFade = false">
			<div v-if="showModal" class="modal-view" :class="{ 'modal-view--fade': makeFade }">
				<div class="modal-view__wrapper">
					<slot></slot>
				</div>
			</div>
		</transition>
	</teleport>
</template>

<script lang="ts" setup>
import { ref } from "vue"

defineProps({
	showModal: {
		type: Boolean,
		required: true,
	},
})

const makeFade = ref<boolean>(false);

</script>

<style>
.modal-view {
	position: absolute;
	top: 0;
	bottom: 0;
	left: 0;
	right: 0;
	/* background-color: rgba(var(--black), 0.15); */
	box-sizing: border-box;

	z-index: 999;
}

.modal-view--fade {
	background-color: rgba(var(--black), 0.15);
}

.modal-view__wrapper {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
}

.slide-enter-active,
.slide-leave-active {
	transition: all 0.15s ease-out;
}

.slide-enter-from,
.slide-leave-to {
	/* opacity: 0; */
	top: 100%;
}
</style>