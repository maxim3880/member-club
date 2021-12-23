import { Table, TableCell, TableContainer, TableHead, TableRow, TableBody } from "@mui/material";

function InnerTable(props) {
    return (
        <TableContainer >
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell align="left" >Name</TableCell>
                        <TableCell align="center">Email</TableCell>
                        <TableCell align="right">Registration date</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {props.rows?.map((row) => {
                        debugger
                    return (
                            <TableRow key={row?.email}>
                                <TableCell align="left">{row?.name}</TableCell>
                                <TableCell align="center">{row?.email}</TableCell>
                                <TableCell align="right">{row?.registration_date}</TableCell>
                        </TableRow>
                        )
                })}
                </TableBody>
            </Table>
        </TableContainer>
    )
}

export default InnerTable;