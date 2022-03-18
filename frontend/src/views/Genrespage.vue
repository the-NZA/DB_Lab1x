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
			:rowData="getGenresRows"
		></ag-grid-vue>

		<modal-view :showModal="showModal">
			<genre-editor
				:genre_id="selectedGenreID"
				@savePressed="handleSaveGenre"
				@closePressed="handleCloseModal"
			/>
		</modal-view>
	</template>

	<template v-else>
		<section class="search sectionMain">
			<search-form @searchSubmitted="handleSearch"></search-form>
		</section>

		<section class="genres sectionMain">
			<div class="genres__cards">
				<div class="genres__card" v-for="genre in getGenres">
					<h3 class="genres__cardtitle">
						<a href="#">{{ genre.title }}</a>
					</h3>
				</div>
			</div>
		</section>
	</template>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, onMounted, computed, watch } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "../store";
import { storeToRefs } from "pinia";
import SearchForm from "../components/SearchForm.vue";
import GridButtonsVue from "../components/GridButtons.vue";
import ActionsButtons from "../components/ActionsButtons.vue"
import ModalView from "../components/ModalView.vue";
import GenreEditor from "../components/GenreEditor.vue";
import { AgGridVue } from "ag-grid-vue3";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";
import { GenreRow } from "../types/grid";

const store = useStore();
const { isGenresLoaded, getGenresRows, getGenres } = storeToRefs(store);

const router = useRouter()

const handleSearch = async (searchBy: string, searchQuery: string) => {
	try {
		router.push({
			name: "Books",
			query: {
				searchBy: searchBy,
				searchQuery: searchQuery,
			}
		})

	} catch (err) {
		console.error(err);
	}
}

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables for buttons
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const showModal = ref<boolean>(false);
const selectedGenreID = ref<string | undefined>();

const handleCloseModal = () => {
	showModal.value = false;
}

const handleSaveGenre = () => {
	showModal.value = false;
}

const handleAdd = () => {
	selectedGenreID.value = undefined;
	showModal.value = true;
}

const handleEdit = () => {
	const selected = gridApi.value?.getSelectedRows()
	if (selected?.length !== 1) {
		console.log("Выбрано более одного элемента. Редактирование недоступно.");
		return;
	}

	selectedGenreID.value = (selected[0] as GenreRow).id
	showModal.value = true;
}

const handleDelete = () => {
	const selected = gridApi.value?.getSelectedRows()

	if (confirm("Вы уверены, что хотите удалить выбранные жанры?")) {
		selected?.forEach((genre: GenreRow) => {
			store.deleteGenre(genre.id);
		})
	}
}

// Handle selection changed event
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

// Current grid options
const gridOptions = ref<GridOptions>({
	onSelectionChanged: onSelectionChanged,
	onGridReady: onGridReady,
	rowSelection: "multiple",
	suppressCellSelection: true,
	suppressRowClickSelection: true,
	localeText: {
		noRowsToShow: 'Жанров пока не добавлено',
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
			flex: 2,
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
			field: "actions",
			headerName: "Действия",
			cellRenderer: "gridBtn",
			minWidth: 115,
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
	if (confirm("Вы действительно хотите удалить выбранный жанр?")) {
		store.deleteGenre(id);
	}
};
const editFunc = (id: string) => {
	selectedGenreID.value = id
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

<style scoped>
.genres__cards {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: var(--offset);
}

.genres__card {
	background-color: rgb(var(--leanen));
	box-shadow: var(--main-shadow);
}

.genres__cardtitle {
	margin: 0;
	padding: var(--offset);
}

.genres__cardtitle a {
	text-decoration: none;
	color: rgb(var(--sapphire));
	transition: var(--main-transition);
}

.genres__cardtitle a:hover {
	text-decoration: underline;
	color: rgb(var(--redsand));
}
</style>