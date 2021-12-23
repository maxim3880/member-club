import { Typography } from '@mui/material';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import InnerTable from '../../components/Table';

const rows = [
    {id: "1", email: "maxim3880@gmail.com", name: "Max Navrotskiy", registration_date: "2021-01-02"},
    {id: "1", email: "maxim3880@gmail.com", name: "Max Navrotskiy", registration_date: "2021-01-02"},
    {id: "1", email: "maxim3880@gmail.com", name: "Max Navrotskiy", registration_date: "2021-01-02"},
]

function Main() {
    return (
        <Grid xs={12} container direction="column" justifyContent="space-between" alignItems="center">
            <Grid item xs={12} justifyContent="center" alignItems="center">
                <Typography variant="h2">Member club UI</Typography>
            </Grid>
            <Grid xs={10} container direction="row" justifyContent="space-between" alignItems="stretch" spacing={2}>
                <Grid item xs={4}>
                    <Paper square>
                        <Grid container direction="column"  alignItems="center">
                            <Grid item xs={10}>
                                <Typography variant="h4">Member info fields</Typography>
                            </Grid>
                            <Grid xs={10} container direction="column" justifyContent="space-around" alignItems="stretch" spacing={2}>
                                <Grid item>
                                    <Grid container direction="column" justifyContent="space-between" alignItems="stretch" spacing={1}>
                                        <Grid item ><TextField fullWidth required id="outlined-required" label="Your name" /></Grid>
                                        <Grid item ><TextField fullWidth required id="outlined-required" label="Your email"/></Grid>
                                    </Grid>
                                </Grid>
                                <Grid item>
                                    <Grid container direction="row" justifyContent="space-between" alignItems="stretch" spacing={1}>
                                        <Grid item xs={6}><Button fullWidth variant="outlined">Clear</Button></Grid>
                                        <Grid item xs={6}><Button fullWidth variant="contained">Add</Button></Grid>
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
                <Grid item xs={8}>
                    <Paper>
                        <Grid container direction="row" justifyContent="center"  alignItems="center">
                            <Grid container xs={10} justifyContent="center">
                                <Typography variant="h4">Members</Typography>
                            </Grid>
                            <Grid container xs={10}>
                                <InnerTable rows={rows}></InnerTable>
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
            </Grid>
        </Grid>
    )
}
export default Main;