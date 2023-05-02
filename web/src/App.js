import './App.css';
import Table from "./components/Table";
import {useEffect, useState} from "react";
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Modal from '@mui/material/Modal';
import axios from 'axios';
import {Form} from "./components/Form";


function App() {
    const [data, setData] = useState([])
    const [selectedID, setSelectedID] = useState(null)
    const [open, setOpen] = useState(false);

    const fetchData = async() => {
        const fData= await axios.get("http://localhost:8080/list")
        setData(fData.data);
    };

    useEffect( () => {
        fetchData();
    }, [selectedID]);

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);
    const handleAdd = () => {
        fetchData()
        setOpen(false)
    };

    const handleRowSelect = (index) => {
        setSelectedID(index);
    }

    const handleDelete = () => {
        axios.delete(`http://localhost:8080/event/${selectedID}/delete`)
        setSelectedID(null)
    }

  return (
    <div className="App">
        <div className="basic-table">
            <Stack direction="row" justifyContent="flex-end" spacing={2} mb={2}>
                <Button
                    variant="contained"
                    onClick={handleOpen}
                >Add</Button>
                <Button
                    variant="contained"
                    disabled={!selectedID ? true : false}
                    onClick={handleDelete}
                >Delete</Button>
            </Stack>
            <Table onRowSelect={handleRowSelect} data={data} />
            <Modal
                open={open}
                onClose={handleClose}
            >
                <Form
                    onSentForm={handleAdd}
                />
            </Modal>
        </div>
    </div>
  );
}

export default App;
