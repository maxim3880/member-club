import { Typography } from '@mui/material';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import InnerTable from '../../components/Table';
import AlertButton from '../../components/AlertButton';
import { Component } from 'react';

class Main extends Component {
    constructor(props) {
        super(props);
        this.state = {
            rows: [],
            name: "",
            email: "",
            open: false,
            openSuccess: false,
            alertMessage: ""
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleAddClick = this.handleAddClick.bind(this)
        this.handleClearClick = this.handleClearClick.bind(this)
        this.setOpen = this.setOpen.bind(this)
    }

    componentDidMount() {
        this.updateRows()
    }

    async updateRows() {
        const response = await fetch(`${process.env.REACT_APP_API_URL}/users`, {
            mode: 'cors',
        });
        const respBody = await response.json();
        this.setState({ rows: respBody.data })
    }

    handleChange(event) {
        const { name, value } = event.target;
        this.setState({ [name]: value });
    };

    handleClearClick() {
        this.setState({
            name: "",
            email: "",
            open: false
        })
    }

    async handleAddClick() {
        let reqBody = {
            name: this.state.name,
            email: this.state.email,
        }
        const response = await fetch(`${process.env.REACT_APP_API_URL}/users`, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(reqBody)
        });
        var respBody = await response.json();

        if (!!(respBody?.error)) {
            this.setState({
                open: true,
                alertMessage: respBody.error
            })
            return
        }
        await this.updateRows()
        this.setOpen("openSuccess", true)
    }

    setOpen(key, value) {
        this.setState({ [key]: value })
    }

    render() {
        return (
            <Grid xs={12} container direction="column" justifyContent="space-between" alignItems="center">
                <Grid item xs={12} justifyContent="center" alignItems="center">
                    <Typography variant="h2">Member club UI</Typography>
                </Grid>
                <Grid xs={10} container direction="row" justifyContent="space-between" alignItems="stretch" spacing={2}>
                    <Grid item xs={4}>
                        <Paper square elevation={12}>
                            <Grid container direction="column" alignItems="center">
                                <Grid item xs={10}>
                                    <Typography variant="h4">Add new member</Typography>
                                </Grid>
                                <Grid xs={10} container direction="column" justifyContent="space-around" alignItems="stretch" spacing={4}>
                                    <Grid item>
                                        <Grid container direction="column" justifyContent="space-between" alignItems="stretch" spacing={1}>
                                            <Grid item ><TextField fullWidth name="name" value={this.state.name} onChange={this.handleChange} required id="outlined-required" label="Your name" /></Grid>
                                            <Grid item ><TextField fullWidth name="email" value={this.state.email} onChange={this.handleChange} required id="outlined-required" label="Your email" /></Grid>
                                        </Grid>
                                    </Grid>
                                    <Grid item>
                                        <AlertButton
                                            open={this.state.open}
                                            openSuccess={this.state.openSuccess}
                                            handleAddClick={this.handleAddClick}
                                            handleClearClick={this.handleClearClick}
                                            setOpen={this.setOpen}
                                            alertMessage={this.state.alertMessage}
                                        />
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Grid>
                    <Grid item xs={8}>
                        <Paper elevation={12}>
                            <Grid container direction="row" justifyContent="center" alignItems="center">
                                <Grid container xs={10} justifyContent="center">
                                    <Typography variant="h4">Members</Typography>
                                </Grid>
                                <Grid container xs={10}>
                                    <InnerTable rows={this.state.rows}></InnerTable>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Grid>
                </Grid>
            </Grid>
        )
    }
}
export default Main;