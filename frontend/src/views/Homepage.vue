<template>
	<h2>Описание задачи</h2>
	<p>
		Данное приложение реализует работу с базой данных, которая хранит
		информацию о книгах, жанрах и авторах.
	</p>
	<p>
		Доступно создание, редактирование и удаление записей для каждой из
		сущностей.
	</p>

	<div class="group">
		<h3>Книги</h3>
		<p>
			Каждая книги содержит информацию о её названии, описании,
			жанре, количестве страниц, годе публикации.
		</p>
	</div>

	<div class="group">
		<h3>Жанры</h3>
		<p>
			Каждая запись о жанре содержит информацию о названии и
			описании.
		</p>
	</div>

	<div class="group">
		<h3>Авторы</h3>
		<p>
			Для каждого автора доступна информация о его имени, фамилии,
			отчестве и краткой характеристики.
		</p>
	</div>

	<p>
		Исходные тексты доступны в репозитории на
		<a href="https://github.com/the-NZA/DB_Lab1">Github</a>
	</p>
</template>

<script lang="ts" setup>
</script>

<style scoped>
h2 {
	margin: 0 0 var(--offset) 0;
	text-transform: uppercase;
	font-weight: 800;
	font-size: 2rem;
}

h3 {
	margin: 0 0 var(--offset-half) 0;
}

p {
	margin: 0 0 var(--offset-half) 0;
	font-size: 1.2rem;
	font-weight: 500;
}

.group {
	border: 3px solid rgb(var(--redsand));
	border-radius: var(--offset-quarter);
	margin-bottom: var(--offset);
	padding: var(--offset);
}

.group:first-of-type {
	margin-top: var(--offset);
}

.group p {
	margin-bottom: var(--offset-half);
	line-height: 1.5;
}

.group p:last-of-type {
	margin: 0;
}
</style>
