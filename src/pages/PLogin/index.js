import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";
import { withSnackbar } from "notistack";
import { Card, CardContent, CardActions, CircularProgress } from "@material-ui/core";
import PersonAddIcon from "@material-ui/icons/PersonAddOutlined";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import api from "../../services/api";
import { login } from "../../services/auth";
import Logo from "../../assets/logo512.png";
// import { green } from '@material-ui/core/colors';

import { Form, Container, StyledInput, StyledButton } from "./styles";

const styles = {
    inputs: {
        margin: "10px 0px",
    },
    cardActionsContainer: {
        flexDirection: "column",
        marginLeft: "0px",
    },
    cardActions: {
        marginLeft: "0px",
    },
};

class PLogin extends Component {
    state = {
        username: "",
        password: "",

        errorUsername: false,
        errorPassword: false,

        isLoading: false,
    }

    _handleLogin = event => {
        event.preventDefault();
        this.setState({ isLoading: true });
        const { username, password } = this.state;
        if (!username) {
            this.setState({ errorUsername: true, isLoading: false });
            this.props.enqueueSnackbar("Por favor insira um nome de usuário", { variant: "error" });
        }
        if (!password) {
            this.setState({ errorPassword: true, isLoading: false });
            this.props.enqueueSnackbar("Por favor insira uma senha", { variant: "error" });
        }

        if (username && password) {
            api.post("/login", { username, password })
                .then(res => {
                    if (res.data === "User not found.") {
                        this.setState({ errorUsername: true, isLoading: false });
                        this.props.enqueueSnackbar("Usuário não encontrado", { variant: "error" });
                    }
                    if (res.data === "Invalid password.") {
                        this.setState({ errorPassword: true, isLoading: false });
                        this.props.enqueueSnackbar("Senha incorreta", { variant: "error" });
                    }
                    if (res.data.token) {
                        login(res.data.token);
                        this.props.history.push({
                            pathname: "/home",
                            state: {
                                firstLogin: true,
                                user: {
                                    name: username,
                                },
                            },
                        });
                    }
                })
                .catch(err => {
                    this.setState({ isLoading: false });
                    this.props.enqueueSnackbar("Ocorreu um erro ao logar, por favor tente novamente...", { variant: "error" });
                })
        }
    }

    render() {
        return (
            <Container>
                <Form>
                    <Card>
                        <CardContent>
                            <img src={Logo} alt="Logo" />
                            <StyledInput
                                type="text" variant="outlined" style={styles.inputs}
                                onChange={event => this.setState({ errorUsername: false, username: event.target.value })}
                                label="Nome de usuário"
                                required
                                error={this.state.errorUsername} />

                            <StyledInput
                                type="password" variant="outlined" style={styles.inputs}
                                onChange={event => this.setState({ errorPassword: false, password: event.target.value })}
                                label="Senha"
                                required
                                error={this.state.errorPassword} />
                        </CardContent>

                        <CardActions style={styles.cardActionsContainer}>
                            <StyledButton variant="outlined" onClick={this._handleLogin} disabled={this.state.isLoading}>
                                {!this.state.isLoading && <ExitToAppIcon style={{ marginRight: "8px", color: "#999" }} />}
                                {!this.state.isLoading && "Acessar"}
                                {this.state.isLoading && <CircularProgress size={24} />}
                            </StyledButton>
                            <hr style={styles.cardActions} />
                            <Link style={styles.cardActions} to="/register">
                                <PersonAddIcon />
                                <br />
                                Cadastrar-se
                            </Link>
                        </CardActions>
                    </Card>
                </Form>
            </Container>
        )
    }
}

export default withRouter(withSnackbar(PLogin));