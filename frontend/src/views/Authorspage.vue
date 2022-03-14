<template>
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
		:rowData="getAuthorsRows"
	></ag-grid-vue>

	<modal-view :showModal="showModal">
		<author-editor
			:author_id="selectedAuthorID"
			@savePressed="handleSaveAuthor"
			@closePressed="handleCloseModal"
		/>
	</modal-view>
</template>

<script lang="ts" setup>
import { ref, reactive, onBeforeMount, computed } from "vue";
import { useStore } from "../store";
import { storeToRefs } from "pinia";
import GridButtonsVue from "../components/GridButtons.vue";
import ActionsButtons from "../components/ActionsButtons.vue"
import ModalView from "../components/ModalView.vue";
import AuthorEditor from "../components/AuthorEditor.vue";
import { AgGridVue } from "ag-grid-vue3";
import {
	GridReadyEvent,
	GridApi,
	GridOptions,
	ColumnApi,
	SelectionChangedEvent,
} from "@ag-grid-community/all-modules";
import { AuthorRow } from "../types/grid";

// Reactive vars from store
const store = useStore();
const { isAuthorsLoaded, getAuthorsRows } = storeToRefs(store);

// Context for action buttons
const actionContext = ref({});
const gridApi = ref<GridApi>();
const columnApi = ref<ColumnApi>();

// Conditional variables
const hasSelected = ref<boolean>(false);
const singleSelected = ref<boolean>(false);

const showModal = ref<boolean>(false);
const selectedAuthorID = ref<string | undefined>();

const handleCloseModal = () => {
	showModal.value = false;
}

const handleSaveAuthor = () => {
	showModal.value = false;
}

const handleAdd = () => {
	selectedAuthorID.value = undefined;
	showModal.value = true;
}

const handleEdit = () => {
	const selected = gridApi.value?.getSelectedRows()
	if (selected?.length !== 1) {
		console.log("Выбрано более одного элемента. Редактирование недоступно.");
		return;
	}

	selectedAuthorID.value = (selected[0] as AuthorRow).id
	showModal.value = true;
}

const handleDelete = () => {
	const selected = gridApi.value?.getSelectedRows()

	if (confirm("Вы уверены, что хотите удалить выбранных авторов?")) {
		selected?.forEach((author: AuthorRow) => {
			store.deleteAuthor(author.id);
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
		noRowsToShow: 'Авторов пока не добавлено',
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
			field: "firstname",
			checkboxSelection: true,
			headerCheckboxSelection: true,
			headerName: "Имя",
			minWidth: 110,
			flex: 1,
			cellClass: ["centered-cell"]
		},
		{
			field: "lastname",
			headerName: "Фамилия",
			minWidth: 110,
			flex: 1,
			cellClass: ["centered-cell"]
		},
		{
			field: "surname",
			headerName: "Отчество",
			minWidth: 110,
			flex: 1,
			cellClass: ["centered-cell"]
		},
		{
			field: "books",
			headerName: "Книги",
			minWidth: 110,
			flex: 1,
			// maxWidth: 115,
			cellClass: ["centered-cell"]
		},
		{
			field: "snippet",
			headerName: "Описание",
			flex: 3,
			minWidth: 380,
			cellClass: ["centered-cell"]
		},
		{
			field: "actions",
			headerName: "Действия",
			cellRenderer: "gridBtn",
			flex: 0.5,
			minWidth: 115,
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
	if (confirm("Вы действительно хотите удалить выбранного автора?")) {
		store.deleteAuthor(id);
	}
};
const editFunc = (id: string) => {
	selectedAuthorID.value = id
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