import React, { Component } from "react";
import CAppMenu from "../../components/CAppMenu";
import styled from "styled-components";
import { Typography } from "@material-ui/core";

const StyledDiv = styled.div`
    margin: 100px 50px 0px 50px;
`;

const Title = styled(Typography)`
    text-align: center;
`;

const Content = styled(Typography)`
    margin-top: 20px !important;
    text-align: justify;
    text-justify: inter-word;
`;

class PAbout extends Component {

    render() {
        return (
            <React.Fragment>
                <CAppMenu user={this.props.location.state.user} />
                <StyledDiv>
                    <Title variant="h4">
                        Todolistreact
                    </Title>
                    <Content>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit,
                        sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
                        Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
                        nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
                        reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
                        Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia
                        deserunt mollit anim id est laborum.
                    </Content>
                </StyledDiv>
            </React.Fragment>
        );
    };
};

export default PAbout;