import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";
import { withSnackbar } from "notistack";
import { Card, CardContent, CardActions } from "@material-ui/core";
import PersonAddIcon from "@material-ui/icons/PersonAddOutlined";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import api from "../../services/api";
import Logo from "../../assets/logo512.png";

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

class PRegister extends Component {
    state = {
        username: "",
        email: "",
        password: "",

        errorUsername: false,
        errorPassword: false,
    }

    _handleRegister = event => {
        event.preventDefault();
        const { username, email, password } = this.state;
        if (!username) {
            this.setState({ errorUsername: true });
            this.props.enqueueSnackbar("Por favor insira um nome de usuário", { variant: "error" });
        }
        if (!password) {
            this.setState({ errorPassword: true });
            this.props.enqueueSnackbar("Por favor insira uma senha", { variant: "error" });
        }

        if (username && password) {
            api.post("/register", { username, email, password })
                .then(res => {
                    if (res.data === "This username already exists.") {
                        this.setState({ errorUsername: true });
                        this.props.enqueueSnackbar("Nome de usuário já existente, por favor tente outro", { variant: "error" });
                    } else {
                        this.props.enqueueSnackbar("Usuário cadastrado com sucesso!", { variant: "success" });
                    }
                })
                .catch(err => {
                    console.log(err);
                    this.props.enqueueSnackbar("Ocorreu um erro ao registrar a conta, por favor tente novamente...", { variant: "error" });
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
                                type="email" variant="outlined" style={styles.inputs}
                                onChange={event => this.setState({ email: event.target.value })}
                                label="Endereço de e-mail"
                                placeholder="seu@email.com" />

                            <StyledInput
                                type="password" variant="outlined" style={styles.inputs}
                                onChange={event => this.setState({ errorPassword: false, password: event.target.value })}
                                label="Senha"
                                required
                                error={this.state.errorPassword} />
                        </CardContent>

                        <CardActions style={styles.cardActionsContainer}>
                            <StyledButton variant="outlined" onClick={this._handleRegister}>
                                <PersonAddIcon style={{ color: "#999", marginRight: "8px" }} />
                                Cadastrar-se
                            </StyledButton>
                            <hr style={styles.cardActions} />
                            <Link style={styles.cardActions} to="/">
                                <ExitToAppIcon />
                                <br />
                                Acessar
                            </Link>
                        </CardActions>
                    </Card>
                </Form>
            </Container>
        )
    }
}

export default withRouter(withSnackbar(PRegister));