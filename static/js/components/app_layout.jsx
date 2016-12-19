'use strict';

import React from 'react';
import {Layout, Content, Drawer, Header, HeaderRow, HeaderTabs, Tab} from 'react-mdl';
import Login from './login.jsx';

export default class AppLayout extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            title: window.data.title
        };
    }

    render() {
        return (
            <div style={{height: '300px', position: 'relative'}}>
                <Layout fixedHeader fixedTabs>
                    <Header>
                        <HeaderRow title={this.state.title}>
                            <Login />
                        </HeaderRow>
                        <HeaderTabs ripple activeTab={1} onChange={(tabId) => {}}>
                            <Tab>Tab1</Tab>
                            <Tab>Tab2</Tab>
                            <Tab>Tab3</Tab>
                        </HeaderTabs>
                    </Header>
                    <Drawer title={this.state.title} />
                    <Content>
                        <div className="page-content">
                        </div>
                    </Content>
                </Layout>
            </div>
        );
    }
}
