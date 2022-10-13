
import React, { useEffect } from "react";
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import { AppBar, Button, FormControl, IconButton, Paper, Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { UsersTypeInterface } from '../models/userTypesUI';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';

import TextField from '@mui/material/TextField';


const Item = styled(Paper)(({ theme }) => ({
      backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
      ...theme.typography.body2,
      padding: theme.spacing(1),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    }));




export default function SimpleContainer() {
      const [age, setAge] = React.useState('');
      console.log(age)

      const [users, setUsers] = React.useState<UsersTypeInterface[]>([]);
      const getUsers = async () => {
            const apiUrl = "http://localhost:8080/ListUserTypes";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
              .then((response) => response.json())
              .then((res) => {
                console.log(res.data);
                if (res.data) {
                  setUsers(res.data);
                }
              });
          };
      useEffect(() => {
            getUsers();
      }, []);
      
      const handleChange = (event: SelectChangeEvent) => {
            setAge(event.target.value)
          };
  return (
      <Paper elevation={0}>
            <Box>
                  <AppBar position="static">
                        <Toolbar>
                              <IconButton
                                    size="large"
                                    edge="start"
                                    color="inherit"
                                    aria-label="menu"
                                    sx={{ mr: 2 }}
                              >
                              <MenuIcon />
                              </IconButton>
                              <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                                    News
                              </Typography>
                              <Button color="inherit">
                                    Logout
                              </Button>
                        </Toolbar>
                  </AppBar>
            </Box>

            <Container maxWidth="md">
                  <Paper elevation={0}>
                        <Box
                              display={"flex"}
                              sx={{
                                    marginTop : 2,
                                    marginX : 2 ,
                              }}
                        >
                              <h2>
                                    ระบบบันทึกการใช้งานเตียง
                              </h2>
                        </Box>
                        <hr />
                        <Grid container spacing={2} sx ={{padding : 2}}>
                              <Grid item xs={10}>
                                    <p>ชื่อผู้ป่วย</p>
                                          <FormControl fullWidth>
                                                <Select
                                                id="demo-select-small"
                                                value={age}
                                                displayEmpty
                                                inputProps={{ 'aria-label': 'Without label' }}
                                                onChange={handleChange}
                                                >
                                                      <MenuItem value="">
                                                            <em>None</em>
                                                      </MenuItem>
                                                      {users.map( item => (
                                                            <MenuItem value={item.ID} key = {item.ID}>{item.UserType}</MenuItem>
                                                      ))}
                                                </Select>
                                          </FormControl>
                        
                              </Grid>
                              <Grid item xs={2} >
                                          <Button variant="contained" sx={{marginTop : 8}}>
                                                ค้นหา
                                          </Button>
                                   
                              </Grid>
                              <Grid item xs={4}>
                                    <p>เพศ</p>
                                    <TextField
                                          id="outlined-read-only-input"
                                          defaultValue="Hello World"
                                          InputProps={{
                                                readOnly: true,
                                          }}
                                    />
                                  
                              </Grid>
                              <Grid item xs={4}>
                                    <p>ประเภทโรค</p>
                                    <TextField
                                          id="outlined-read-only-input"
                                          defaultValue="Hello World"
                                          InputProps={{
                                                readOnly: true,
                                          }}
                                    />
                                   
                              </Grid>
                              <Grid item xs={4}>
                                    <p>โรค</p>
                                    <TextField
                                          id="outlined-read-only-input"
                                          defaultValue="Hello World"
                                          InputProps={{
                                                readOnly: true,
                                          }}
                                    />
                                    
                              </Grid>
                              <Grid item xs={12}>
                                    <TextField
                                          id="outlined-read-only-input"
                                          defaultValue="Hello World"
                                          InputProps={{
                                                readOnly: true,
                                          }}
                                    />
                              </Grid>
                              <Grid item xs={4}>
                                   
                              </Grid>
                              <Grid item xs={4}>
                                    
                              </Grid>
                              <Grid item xs={4}>
                                    
                              </Grid>
                              <Grid item xs={12}>
                                   
                              </Grid>
                              <Grid item xs={12}>
                                    
                              </Grid>
                        </Grid>
                  </Paper>
            </Container>

            <Container>
                  
            </Container>
            

           
      </Paper>

  );
}


{/* <ul>
        {users.map( item => (
          <li key={item.ID}>
            {item.UserType}
          </li>
        ))}
       </ul> */}