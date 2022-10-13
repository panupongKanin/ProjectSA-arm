import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import {FormControl, Grid, InputLabel, MenuItem, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField} from '@mui/material';
import { styled } from '@mui/material/styles';
import { useState,useEffect} from 'react';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { id } from 'date-fns/locale';







export default function MyComponent() {
 
  const [users, setUsers] = useState<any[]>([]);
  console.log(users)

  const getUsers = async () => {
  const apiUrl = "http://localhost:8080/ListUserTypes";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => setUsers(res));
  };

  useEffect(() => {
    getUsers();
  }, []);

  
    return (
    <ul>
     
    </ul>
    );
}
