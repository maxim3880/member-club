import * as React from 'react';
import Alert from '@mui/material/Alert';
import IconButton from '@mui/material/IconButton';
import Collapse from '@mui/material/Collapse';
import Button from '@mui/material/Button';
import Snackbar from '@mui/material/Snackbar';
import CloseIcon from '@mui/icons-material/Close';
import Grid from '@mui/material/Grid';


function AlertButton(props) {
    return (
        <Grid container direction="row" justifyContent="space-between" alignItems="stretch" spacing={1}>
            
            <Grid item xs={6}>
                <Button fullWidth onClick={props.handleClearClick} variant="outlined">Clear</Button>
            </Grid>
            <Grid item xs={6}>
                <Button fullWidth disabled={props.open} onClick={props.handleAddClick} variant="contained">Add</Button>
            </Grid>
            <Grid item xs={12}>
                <Snackbar open={props.open} autoHideDuration={6000} onClose={() => { props.setOpen("open", false) }}>
                    <Collapse in={props.open}>
                        <Alert
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => {
                                        props.setOpen("open", false);
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                            severity="error"
                        >
                            {props.alertMessage}
                        </Alert>
                    </Collapse>
                </Snackbar>
            </Grid>
            <Grid item xs={12}>
                <Snackbar open={props.openSuccess} autoHideDuration={6000} onClose={() => { props.setOpen("openSuccess", false) }}>
                    <Collapse in={props.openSuccess}>
                        <Alert
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => {
                                        props.setOpen("openSuccess", false);
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                        >
                            Data was successfully added
                        </Alert>
                    </Collapse>
                </Snackbar>
            </Grid>
        </Grid>
    )
};

export default AlertButton;