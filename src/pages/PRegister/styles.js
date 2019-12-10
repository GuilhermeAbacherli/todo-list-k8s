import styled from "styled-components";
import { TextField, Button } from "@material-ui/core";

export const Container = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
`;

export const Form = styled.form`
    width: 400px;
    /* background: #fff; */
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    img {
        width: 100px;
        margin: 10px 0 25px;
    }
    hr {
        margin: 20px 0px;
        border: none;
        border-bottom: 1px solid #cdcdcd;
        width: 100%;
    }
    a {
        margin-bottom: 8px;
        font-size: 16;
        font-weight: bold;
        color: #999;
        text-decoration: none;
    }
    a:hover{
        color: blue;
    }
`;

export const StyledInput = styled(TextField)`
    width: 100%;
`;

export const StyledButton = styled(Button)`
    width: 96%;
    height: 50px;
`;
