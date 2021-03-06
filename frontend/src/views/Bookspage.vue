<template>
	<template v-if="store.isAdmin">
		<actions-buttons
			:canEdit="singleSelected"
			:canDelete="hasSelected"
			@add-pressed="handleAdd"
			@edit-pressed="handleEdit"
			@delete-pressed="handleDelete"
		></actions-buttons>

		<ag-grid-vue
			style="width: 100%; height: 100%;"
			class="ag-theme-alpine"
			:gridOptions="gridOptions"
			:context="actionContext"
			:rowData="getBooksRows"
		></ag-grid-vue>

		<modal-view :showModal="showModal">
			<book-editor
				:book_id="selectedBookID"
				@savePressed="handleSaveBook"
				@closePressed="handleCloseModal"
			/>
		</modal-view>
	</template>

	<template v-else>
		<section class="search sectionMain">
			<search-form @searchSubmitted="handleSearch" :initialBy="searchBy" :initialQuery="searchQuery"></search-form>
		</section>

		<section class="books sectionMain">
			<div class="books__cards">
				<template v-if="searchedBooks">
					<book-card v-for="book in searchedBooks" :book="book"></book-card>
				</template>

				<p
					style="margin: 0; font-weight: bold; text-align: center;grid-column-start: 2;"
					v-else
				>Совпадений не найдено</p>
			</div>
		</section>
	</template>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, computed } from "vue";
import { useStore } from "../store";
import { useRoute, useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import GridButtonsVue from "../components/GridButtons.vue";
import BookCard from "../components/BookCard.vue";
import ActionsButtons from "../components/ActionsButtons.vue"
import SearchForm from "../components/SearchForm.vue"
import ModalView from "../components/ModalView.vue";
import BookEditor from "../components/BookEditor.vue";
import { AgGridVue } from "ag-grid-vue3";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";
import { BookRow } from "../types/grid";
import { Book } from "../types";
import { GET } from "../HTTP";

const store = useStore();
const { isBooksLoaded, getBooksRows } = storeToRefs(store);

const router = useRouter()
const route = useRoute()
const searchBy = route.query.searchBy?.toString()
const searchQuery = route.query.searchQuery?.toString()
const searchedBooks = ref<Book[]>([])
onBeforeMount(async () => {

	if (searchBy && searchQuery) {
		handleSearch(searchBy, searchQuery)
		router.replace({
			path: route.path,
		})
		return
	}


	try {
		const res = await GET<Book[]>("api/book/all");
		searchedBooks.value = res

	} catch (err) {
		console.error(err);
	}
})

const handleSearch = async (searchBy: string, searchQuery: string) => {
	try {
		const res = await GET<Book[]>(`api/book/search?${searchBy}=${searchQuery}`);
		searchedBooks.value = res

	} catch (err) {
		console.error(err);
	}
}

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const showModal = ref<boolean>(false);
const selectedBookID = ref<string | undefined>();

const handleCloseModal = () => {
	showModal.value = false;
}

const handleSaveBook = () => {
	showModal.value = false;
}

const handleAdd = () => {
	selectedBookID.value = undefined;
	showModal.value = true;
}

const handleEdit = () => {
	const selected = gridApi.value?.getSelectedRows()
	if (selected?.length !== 1) {
		console.log("Выбрано более одного элемента. Редактирование недоступно.");
		return;
	}

	selectedBookID.value = (selected[0] as BookRow).id
	showModal.value = true;
}

const handleDelete = () => {
	const selected = gridApi.value?.getSelectedRows()

	if (confirm("Вы уверены, что хотите удалить выбранные книги?")) {
		selected?.forEach((book: BookRow) => {
			store.deleteBook(book.id);
		})
	}
}

const onSelectionChanged = (e: SelectionChangedEvent) => {
	const cnt = e.api.getSelectedRows().length
	switch (cnt) {
		case 0:
			hasSelected.value = false;
			singleSelected.value = false;
			break;
		case 1:
			hasSelected.value = true;
			singleSelected.value = true;
			break;
		default:
			hasSelected.value = true;
			singleSelected.value = false;
			break;

	}
}

// Save grid and column API on grid ready event
const onGridReady = (params: GridReadyEvent) => {
	gridApi.value = params.api;
	columnApi.value = params.columnApi;
};

const gridOptions = ref<GridOptions>({
	onSelectionChanged: onSelectionChanged,
	onGridReady: onGridReady,
	rowSelection: "multiple",
	suppressCellSelection: true,
	suppressRowClickSelection: true,
	localeText: {
		noRowsToShow: 'Книг пока не добавлено',
	},
	defaultColDef: {
		wrapText: true,
		autoHeight: true,
	},
	columnDefs: [
		{
			field: "id",
			hide: true,
		},
		{
			field: "title",
			checkboxSelection: true,
			headerCheckboxSelection: true,
			headerName: "Название",
			flex: 1,
			minWidth: 250,
			cellClass: ["centered-cell"]
		},
		{
			field: "snippet",
			headerName: "Описание",
			flex: 3,
			minWidth: 450,
			cellClass: ["centered-cell"]
		},
		{
			field: "authors",
			headerName: "Авторы",
			minWidth: 110,
			flex: 1,
			// maxWidth: 115,
			cellClass: ["centered-cell"]
		},
		{
			field: "genre",
			headerName: "Жанр",
			minWidth: 120,
			maxWidth: 140,
			cellClass: ["centered-cell"]
		},
		{
			field: "pages_cnt",
			headerName: "Страниц",
			width: 100,
			// cellClass: ["centered"],
			cellClass: ["centered", "centered-cell"]
		},
		{
			field: "pub_year",
			headerName: "Год публикации",
			maxWidth: 155,
			cellClass: ["centered", "centered-cell"]
		},
		{
			field: "actions",
			headerName: "Действия",
			cellRenderer: "gridBtn",
			minWidth: 115,
			maxWidth: 120,
			flex: 0.5,
			cellClass: ["centered-cell"]
		}
	],
	frameworkComponents: {
		gridBtn: GridButtonsVue
	},
});

// Actions funcs which passed to ag-grid
// and called after button pressed
const deleteFunc = (id: string) => {
	if (confirm("Вы действительно хотите удалить выбранную книгу?")) {
		store.deleteBook(id);
	}
};
const editFunc = (id: string) => {
	selectedBookID.value = id
	showModal.value = true
};

// Before Mount Event handler
onBeforeMount(() => {
	// Setup actions context
	actionContext.value = {
		deleteFunc: deleteFunc,
		editFunc: editFunc,
	};
});
</script>

<style>
</style>