import React from 'react';
import "./styles/global";
import Routes from "./routes";
import GlobalStyle from './styles/global';
import { SnackbarProvider } from "notistack";
import { IconButton } from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Close";

const App = () => (
    <React.Fragment>
        <GlobalStyle />
        <Routes />
    </React.Fragment>
);

const notistackRef = React.createRef();

const AppIntegratedNotistack = () => (
    <SnackbarProvider maxSnack={3}
        ref={notistackRef}
        anchorOrigin={{ vertical: "top", horizontal: "right", }}
        action={key => (
            <IconButton
                style={{ color: "white" }}
                onClick={() => notistackRef.current.closeSnackbar(key)}
            >
                <CloseIcon fontSize="small" />
            </IconButton>)}
    >
        <App />
    </SnackbarProvider >
);

export default AppIntegratedNotistack;
