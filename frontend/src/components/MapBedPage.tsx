import React, { useEffect, useState } from "react";
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import { AppBar, Button, FormControl, IconButton, Paper, Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { TriageInterface} from '../models/userTypesUI';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';

import TextField from '@mui/material/TextField';
import { color } from "@mui/system";
import { ZoneInterface } from "../models/zoneUI";
import { BedInterface } from "../models/bedUI";






export default function SimpleContainer() {
      const [zoneID,setZoneID] = useState('');
      const [bedID,setBedID] = useState('');



      const [users, setUsers] = useState<any[]>([]);
      const [gendersID, setGendersID] = useState<any[]>([]);
      const [genderTypes, setGenderTypes] = useState<any[]>([]);
     
      console.log(gendersID)

      //function fethch data จาก backend Triages
      const getTriages = async () => {
            const apiUrl = "http://localhost:8080/GetListTriages";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
              .then((response) => response.json())
              .then((res) => {
                if (res.data.Patient) {
                  setGendersID(res.data.Patient.Gender_ID);
                }
              });
          };
      //function fethch data จาก backend Genders
      const getGender = async () => {
            const apiUrl = "http://localhost:8080/GetListTriages";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
              .then((response) => response.json())
              .then((res) => {
                if (res.data.Patient) {
                  setGendersID(res.data.Patient.Gender_ID);
                }
              });
          };

      

      const [zones, setZones] = useState<ZoneInterface[]>([]);
      console.log(zones)
      const getZone = async () => {
            const apiUrl = "http://localhost:8080/GetListZones";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              setZones(res.data);
                        }
                  });
      };

      const [beds, setBeds] = React.useState<BedInterface[]>([]);
      const getBed = async () => {
            const apiUrl = `http://localhost:8080/Bed/${zoneID}`;
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              setBeds(res.data);
                        }
                  });
      };
     

      useEffect(() => {
            getTriages();
      }, []);

      useEffect(() => {
            getZone();
      }, []);
     
      useEffect(() => {
            getBed();
      }, [zoneID]);
      


      const handleChange = (event: SelectChangeEvent) => {
           
          };

      const onChangeZone = (event: SelectChangeEvent) => {
            setZoneID(event.target.value as string);
          };
      const onChangeBed = (event: SelectChangeEvent) => {
      setBedID(event.target.value as string);
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
                                       
     
                                          
                              
                                    </Grid>
                                    <Grid item xs={2} >
                                          <Button 
                                                fullWidth 
                                                variant="contained" sx={{marginTop : 8
                                          }}>
                                                ค้นหา
                                          </Button>
                                    
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>เพศ</p>
                                          <TextField
                                                fullWidth
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
                                                fullWidth
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
                                                fullWidth
                                                id="outlined-read-only-input"
                                                defaultValue="Hello World"
                                                InputProps={{
                                                      readOnly: true,
                                                      
                                                }}
                                          />
                                    
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>แผนก</p>
                                          <TextField
                                                fullWidth
                                                id="outlined-read-only-input"
                                                defaultValue="Hello World"
                                                InputProps={{
                                                      readOnly: true,
                                                
                                                }}
                                          />
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>โซน</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="demo-select-small"
                                                      value={zoneID}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={onChangeZone}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {zones.map( zone => (
                                                                  <MenuItem value={zone.ID} key = {zone.ID}>{zone.Zone_Name}</MenuItem>
                                                            ))}
                                                      </Select>
                                                </FormControl>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>เตียง</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="demo-select-small"
                                                      value={bedID}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={onChangeBed}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {beds.map( bed => (
                                                                  <MenuItem value={bed.ID} key = {bed.ID}>{bed.Bed_Name}</MenuItem>
                                                            ))}
                                                      </Select>
                                          </FormControl>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>วันที่เข้ารับการรักษา</p>
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>หมายเหตุ</p>
                                          <TextField
                                                id="outlined-multiline-static"
                                                fullWidth
                                                multiline
                                                rows={4}
                                                defaultValue=""
                                          />
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>ผู้บันทึก</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="demo-select-small"
                                                      value={''}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={handleChange}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {zones.map( zone => (
                                                                  <MenuItem value={zone.ID} key = {zone.ID}>{zone.Zone_Name}</MenuItem>
                                                            ))}
                                                          
                                                      </Select>
                                          </FormControl>
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <Button 
                                                fullWidth
                                                size="large"
                                                color="success"
                                                variant="contained" 
                                                sx={{marginTop : 2}}
                                          >
                                                บันทึก
                                          </Button>
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                              
                              </Grid>
                        </Paper>
                  </Container>

                  

            
            </Paper>

  );
}


      // <ul>
      //   {users.map( item => (
      //     <li key={item.ID}>
      //       {item.UserType}
      //     </li>
      //   ))}
      //  </ul> 