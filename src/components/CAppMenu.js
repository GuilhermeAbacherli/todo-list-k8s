import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import styled from "styled-components";
import { AppBar, Toolbar, SwipeableDrawer, Button, IconButton, Typography, List, Grid } from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import PersonIcon from "@material-ui/icons/PersonRounded";
import LogoutIcon from "@material-ui/icons/PowerSettingsNew";
import TodoIcon from "@material-ui/icons/Description";
import InfoIcon from "@material-ui/icons/Info";
import Logo from "../assets/logo192.png";
import { logout } from "../services/auth";

const Menu = styled.div`
`;

const StyledButton = styled(Button)`
    width: 100%;
`;

const StyledMenuList = styled(List)`
    width: 100%;
    text-align: center;
`;

const StyledMenuLogo = styled.img`
    width: 25%;
    margin: 10px 0 25px;
`;

const styles = {
    icon: {
        marginRight: "8px",
    },
    iconButton: {
        marginRight: "8px",
        color: "#999",
    },
}

class CAppMenu extends Component {

    state = {
        user: null,
        menu: false,
    }

    _handleLogout = event => {
        event.preventDefault();
        logout();
        this.props.history.push("/");
    }

    render() {
        return (
            <Menu>
                <AppBar position="fixed">
                    <Toolbar>
                        <IconButton style={{ color: "white" }} edge="start" onClick={() => { this.setState({ menu: true }) }}>
                            <MenuIcon />
                        </IconButton>
                        <Typography variant="h6" style={{ flexGrow: 1 }}>
                            Todolistreact
                        </Typography>
                        <Grid container direction="row" alignItems="center" style={{ flex: 0.2 }}>
                            <Grid item>
                                <PersonIcon style={styles.icon} />
                            </Grid>
                            <Grid item>
                                {this.props.user.name}
                            </Grid>
                        </Grid>
                        <IconButton style={{ color: "white" }} edge="end" onClick={this._handleLogout}>
                            <LogoutIcon />
                        </IconButton>
                    </Toolbar>
                </AppBar>
                <Toolbar />
                <SwipeableDrawer
                    open={this.state.menu}
                    onClose={() => this.setState({ menu: false })}
                    onOpen={() => this.setState({ menu: true })}
                >
                    <StyledMenuList style={styles.listMenu}>
                        <StyledMenuLogo src={Logo} alt="Logo" style={styles.listLogo} />
                        <StyledButton onClick={() => this.props.history.push({
                            pathname: "/home",
                            state: {
                                user: this.props.location.state.user,
                                firstLogin: false,
                            },
                        })}>
                            <TodoIcon style={styles.iconButton} />
                            Todos
                        </StyledButton>
                        <StyledButton onClick={() => this.props.history.push({
                            pathname: "/about",
                            state: {
                                user: this.props.location.state.user,
                            },
                        })}>
                            <InfoIcon style={styles.iconButton} />
                            Sobre
                        </StyledButton>
                    </StyledMenuList>
                </SwipeableDrawer>
            </Menu>
        );
    }
}

export default withRouter(CAppMenu);