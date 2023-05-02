import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';

export const columns = [
    { field: 'government', headerName: 'Government', width: 160 },
    { field: 'event', headerName: 'Event', width: 270 },
    {field: 'name', headerName: 'Name', width: 160},
    {field: 'position', headerName: 'Position', width: 160},
    {field: 'address', headerName: 'Address', width: 160},
    {field: 'email', headerName: 'Email', width: 150},
    {field: 'phone', headerName: 'Phone', width: 90},
    {field: 'comment', headerName: 'Comment', width: 90},
];

export default function Table(props) {

    function handleRowClick(index) {
        props.onRowSelect(index);
    }

    return (
        <DataGrid
            rows={props.data}
            columns={columns}
            hideFooter={true}
            onRowClick={(itm) => handleRowClick(itm.id)}
        />
    );
}