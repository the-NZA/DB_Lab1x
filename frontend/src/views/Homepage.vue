<template>
	<section class="search">
		<div class="sectionHeader">
			<h2 class="sectionHeader__title">Поиск</h2>
			<p
				class="sectionHeader__subtitle"
			>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facere aut vitae sed. Saepe, ab porro quibusdam impedit explicabo beatae vel?</p>
		</div>

		<form class="search__form">
			<input type="text" />
		</form>
	</section>

	<section class="books">
		<div class="sectionHeader">
			<h2 class="sectionHeader__title">Случайные книги</h2>
			<p
				class="sectionHeader__subtitle"
			>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facere aut vitae sed. Saepe, ab porro quibusdam impedit explicabo beatae vel?</p>
		</div>
		<div class="books__cards">SELECT * FROM books ORDER BY RANDOM() LIMIT 3;</div>
	</section>
</template>

<script lang="ts" setup>
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
</style>
