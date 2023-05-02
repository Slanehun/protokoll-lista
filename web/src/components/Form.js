import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Input from "@mui/material/Input";
import * as React from "react";
import {useState} from "react";
import {columns} from "./Table";
import Button from "@mui/material/Button";
import axios from "axios";

const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 600,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
};

const formStlye = {
    margin: '0.3rem'
}

const initFormData = {
    government: "",
    event: "",
    name: "",
    position: "",
    address: "",
    email: "",
    phone: "",
    comment: ""
}

export const Form = React.forwardRef((props, ref) => {
    const [formData, setFormData] = useState(initFormData)

    const handleChange = (ev, field) => {
        setFormData({...formData, [field]: ev.target.value})
    }

    const handleForm = () => {
        axios.post(`http://localhost:8080/add`, formData)
        props.onSentForm()
    }

    return(
        <Box
            component="form"
            sx={style}
            noValidate
            autoComplete="off"
        >
            {columns.map(c => (
                <FormControl sx={formStlye} key={c.field}>
                    <InputLabel htmlFor={c.field}>{c.field.charAt(0).toUpperCase() + c.field.slice(1)}</InputLabel>
                    <Input id={c.field} aria-describedby={c.field} onChange={(ev) => handleChange(ev, c.field)}/>
                </FormControl>
            ))}
            <Button
                variant="contained"
                onClick={handleForm}
            >Add event</Button>
        </Box>
    )
})