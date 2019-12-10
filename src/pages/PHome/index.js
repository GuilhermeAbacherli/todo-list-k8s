import React, { Component } from "react"
import { withRouter } from "react-router-dom";
import { withSnackbar } from "notistack";
import CAppMenu from "../../components/CAppMenu";
import CTodolist from "../../components/CTodolist";
import api from "../../services/api";

class PHome extends Component {

    state = {
        todolist: [],
        isLoading: true,
    }

    componentDidMount() {
        if (this.props.location.state.firstLogin) this.props.enqueueSnackbar(`Bem-vindo ${this.props.location.state.user.name}!!!`, { variant: "info" });
        this.getTodolist();
    }

    async getTodolist() {
        try {
            this.setState({ isLoading: true });
            const res = await api.get("/todo");
            this.setState({ todolist: res.data });
            this.props.enqueueSnackbar("Lista de TODOs atualizada com sucesso", { variant: "success" });
            this.setState({ isLoading: false });
        } catch (err) {
            this.setState({ isLoading: false });
            this.props.enqueueSnackbar("Falha ao atualizar lista de TODOs", { variant: "error" });
        }
    }

    async addTodo(todo) {
        try {
            const res = await api.post("/todo", todo);
            todo = res.data;
            let { todolist } = this.state;
            todolist.push(todo);
            this.setState({ todolist });
            this.props.enqueueSnackbar(`Sucesso ao adicionar o TODO ${res.data.id}`, { variant: "success" });

        } catch (err) {
            this.props.enqueueSnackbar("Falha ao adicionar", { variant: "error" });
        }
    }

    async updateTodo(newTodo, oldTodo) {
        try {
            const res = await api.patch(`/todo/${oldTodo.id}`, newTodo);
            let { todolist } = this.state;
            todolist[todolist.indexOf(oldTodo)] = newTodo;
            this.setState({ todolist });
            this.props.enqueueSnackbar(`Sucesso ao atualizar o TODO ${res.data.id}`, { variant: "success" });

        } catch (err) {
            this.props.enqueueSnackbar("Falha ao atualizar", { variant: "error" });
        }
    }

    async deleteTodo(oldTodo) {
        try {
            const res = await api.delete(`/todo/${oldTodo.id}`);
            let { todolist } = this.state;
            todolist.splice(todolist.indexOf(oldTodo), 1);
            this.setState({ todolist });
            this.props.enqueueSnackbar(`Sucesso ao excluir o TODO ${res.data.id}`, { variant: "success" });

        } catch (err) {
            this.props.enqueueSnackbar("Falha ao excluir o TODO", { variant: "error" });
        }
    }

    render() {
        return (
            <React.Fragment>
                <CAppMenu user={this.props.location.state.user} />
                <CTodolist
                    todolist={this.state.todolist}
                    isLoading={this.state.isLoading}
                    getTodolist={this.getTodolist.bind(this)}
                    addTodo={this.addTodo.bind(this)}
                    updateTodo={this.updateTodo.bind(this)}
                    deleteTodo={this.deleteTodo.bind(this)}
                />
            </React.Fragment>
        );
    }
}

export default withRouter(withSnackbar(PHome));