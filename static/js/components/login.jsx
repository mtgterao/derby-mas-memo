'use strict';

import React from 'react';
import {
    Button,
    Textfield,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions
} from 'react-mdl';

export default class Login extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};
        this.handleOpenDialog = this.handleOpenDialog.bind(this);
        this.handleCloseDialog = this.handleCloseDialog.bind(this);
    }

    handleOpenDialog() {
        this.setState({
            openDialog: true
        });
    }

    handleCloseDialog() {
        this.setState({
            openDialog: false
        });
    }

    handleLoginButton() {

    }

    render() {
        return (
            <div>
                <Button colored accent onClick={this.handleOpenDialog} raised ripple>ログイン</Button>
                <Dialog open={this.state.openDialog}>
                    <DialogTitle>ログイン</DialogTitle>
                    <DialogContent>
                        <p>メールアドレスとパスワードを入力して、ログインしてください</p>
                        <Textfield
                            label="Email"
                            style={{width: '200px'}}
                            name="mail_address"
                            floatingLabel
                        />
                        <Textfield
                            label="パスワード"
                            style={{width: '200px'}}
                            name="password"
                            floatingLabel
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button raised ripple onClick={this.handleLoginBuddon}>ログイン</Button>
                        <Button raised ripple onClick={this.handleCloseDialog}>閉じる</Button>
                    </DialogActions>
                </Dialog>
            </div>
        );
    }
}
