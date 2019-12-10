import React, { forwardRef } from "react";
import { withRouter } from "react-router-dom";
import MaterialTable from "material-table";

import {
    AddBox, ArrowDownward, Check, ChevronLeft, ChevronRight,
    Clear, DeleteOutline, Edit, FilterList, FirstPage,
    LastPage, Remove, SaveAlt, Search, ViewColumn, Refresh
} from '@material-ui/icons';

const tableIcons = {
    Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
    Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
    Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
    DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
    Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
    Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
    FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
    LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
    NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
    ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
    SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
    ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
    ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />),
};

const refreshIcon = () => (<Refresh />);


const CTodolist = (props) => {
    return (
        <div style={{ margin: "25px 25px 25px 25px" }}>
            <MaterialTable
                title="Todolist"
                icons={tableIcons}
                data={props.todolist}
                isLoading={props.isLoading}
                columns={[
                    { title: "Título", field: "title", emptyValue: "Sem título" },
                    { title: "Descrição", field: "description", emptyValue: "Sem descrição" },
                    { title: "Concluído", field: "done", type: "boolean" },
                ]}
                options={{
                    selection: true,
                    rowStyle: rowData => ({ backgroundColor: rowData.tableData.checked ? '#f1f1f1' : '' }),
                }}
                actions={[
                    {
                        icon: refreshIcon,
                        tooltip: 'Atualizar dados',
                        isFreeAction: true,
                        onClick: props.getTodolist,
                    }
                ]}
                editable={{
                    onRowAdd: newData =>
                        new Promise(async resolve => {
                            await props.addTodo(newData);
                            resolve();
                        }),
                    onRowUpdate: (newData, oldData) =>
                        new Promise(async resolve => {
                            if (oldData) {
                                await props.updateTodo(newData, oldData);
                            }
                            resolve();
                        }),
                    onRowDelete: oldData =>
                        new Promise(async resolve => {
                            await props.deleteTodo(oldData);
                            resolve();
                        }),
                }}
                localization={{
                    pagination: {
                        labelDisplayedRows: "{from}-{to} de {count}",
                        labelRowsSelect: "linhas",
                        firstTooltip: "Primeira página",
                        lastTooltip: "Última página",
                        previousTooltip: "Página anterior",
                        nextTooltip: "Próxima página",
                    },
                    toolbar: {
                        nRowsSelected: "{0} linha(s) selecionada(s)",
                        searchPlaceholder: "Buscar",
                        searchTooltip: "Buscar",
                    },
                    header: {
                        actions: "Ações",
                    },
                    body: {
                        emptyDataSourceMessage: "Não há dados para exibir",
                        addTooltip: "Adicionar",
                        deleteTooltip: "Excluir",
                        editTooltip: "Editar",
                        editRow: {
                            deleteText: "Tem certeza que deseja excluir esse TODO?",
                            cancelTooltip: "Cancelar",
                            saveTooltip: "Salvar",
                        },
                        filterRow: {
                            filterTooltip: "Filtro",
                        },
                    },
                }}
            />
        </div>
    )
};

export default withRouter(CTodolist);