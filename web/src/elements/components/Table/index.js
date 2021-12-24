import { Table, TableCell, TableContainer, TableHead, TableRow, TableBody } from "@mui/material";

function InnerTable(props) {
    return (
        <TableContainer >
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell align="left" ><b>Number</b></TableCell>
                        <TableCell align="left" ><b>Name</b></TableCell>
                        <TableCell align="center" ><b>Email</b></TableCell>
                        <TableCell align="right" ><b>Registration date</b></TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {props.rows?.map((row, index) => {
                        var date = new Date(row.registration_date)
                        return (
                            <TableRow key={row?.email}>
                                <TableCell align="left">{index + 1}</TableCell>
                                <TableCell align="left">{row?.name}</TableCell>
                                <TableCell align="center">{row?.email}</TableCell>
                                <TableCell align="right">{ date.toLocaleDateString()}</TableCell>
                            </TableRow>
                        )
                    })}
                </TableBody>
            </Table>
        </TableContainer>
    )
}

export default InnerTable;