import * as React from 'react';
import { useEffect, useState } from "react";
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import {AppBar, Button, FormControl, IconButton, Paper, Snackbar,Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from '@mui/material/TextField';
import { ZoneInterface,BedInterface ,MappingBedInterface} from "../models/UserInterface";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
      props,
      ref
     ) {
      return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
     });

     
function MappingBedCreate() {


//=======================================================================================================================================
//รับค่าที่ได้จากการเลือก combobox ทั้งหมดเป็นตารางที่ ดึงไปใส่ตารางหลัก 


      ///////////////////บันทึกลงตารางหลัก///////////////////
      const [triageID,setTriageID] = useState('');
      const [zoneID,setZoneID] = useState('');
      const [bedID,setBedID] = useState('');
      const [date, setDate] = useState<Date | null>(null);
      const [comments,setComments] = useState('');
     
      // data ที่ได้มาจากการ fethch
      const [MapBeds, setMapbeds] = useState<Partial<MappingBedInterface>>({});
      const [Zones, setZones] = useState<ZoneInterface[]>([]);
      const [Beds, setBeds] = useState<BedInterface[]>([]);
  
      // data ที่ได้มาจากการ fethch ตารางเพื่อน และ search function
      const [triages, setTriages] = useState<any[]>([]); //getTriages
      const [IPD_Name, setIPD_Name] = useState<any[]>([]);
      const [Disease_Name, setDisease_Name] = useState<any[]>([]);

      /*
            ได้ 
            Disease_Type_Name       -->   ตาราง Disease_Type
            Gender_Name             -->   ตาราง Gender_Name
      */     
      const [DiseaseType, setDiseaseType] = useState<any[]>([]);
      const [GenderType, setGenderType] = useState<any[]>([]);

      // Check save
      const [success, setSuccess] = useState(false);
      const [error, setError] = useState(false);

      const [filtertriages, setFiltertriages] = useState<any[]>([]);
      
      const triage_id = filtertriages.map((id)=>{
            //console.log(id.Triage_ID);
            return id.Triage_ID;
      });

      //console.log(triages);
      const [test,settest] = useState([]);
     
      // Object.assign({},triage_id)
      console.log(Object.assign({},triage_id));

      for (let i of triage_id) {
            //console.log(i.ID);
            // var result = triage_id.filter(number => number != i.ID);
            // console.log(result);
            var result = triages.filter(id => id.ID == i);

            

            //settest(result)

      }

      
      
     
      
      
      
      
      

//=======================================================================================================================================
//สร้างฟังก์ชันสำหรับ คอยรับการกระทำ เมื่อคลิ๊ก หรือ เลือก
      const handleInputChange = (
            event: React.ChangeEvent<{ id?: string; value: any }>
      ) => {
            const id = event.target.id as keyof typeof MappingBedCreate;
            const { value } = event.target;
            setMapbeds({ ...MapBeds, [id]: value });
            setComments(value);
      };
      
      const onChangeTriage = (event: SelectChangeEvent) => {
            setTriageID(event.target.value as string);
      };
      const onChangeZone = (event: SelectChangeEvent) => {
            setZoneID(event.target.value as string);
      };

      const onChangeBed = (event: SelectChangeEvent) => {
            setBedID(event.target.value as string);
      };

      


//=======================================================================================================================================
//function Submit
      const handleClose = (
            event?: React.SyntheticEvent | Event,
            reason?: string
      ) => {
            if (reason === "clickaway") {
            return;
            }
            setSuccess(false);
            setError(false);
            };

      function submit() {
            
            let data = {
            Triage_ID: triageID,
            Bed_ID: bedID,
            Admidtime: date,
            MapBed_Comment: comments,
            //User_ID:
            };

            console.log(data);

            const apiUrl = "http://localhost:8080/CreateMapBed";
            const requestOptions = {
                  method: "POST",
                  headers: { "Content-Type": "application/json" },
                  body: JSON.stringify(data),
                  };
                  fetch(apiUrl, requestOptions)
                        .then((response) => response.json())
                        .then((res) => {
                              if (res.data) {
                                    setSuccess(true);
                              } else {
                                    setError(true);
                              }
                        });
            // reset All after Submit
            setTriageID("");
            setBedID("");
            setDate(null);
            setZoneID("");
            setDisease_Name([]);
            setDiseaseType([])
            setIPD_Name([]);
            setGenderType([]);
            setComments("")
            
      }
      
//=======================================================================================================================================
//function Search
      function search() { 
            const apiUrl1 = `http://localhost:8080/GetTriage/${triageID}`;
            const requestOptions1 = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl1, requestOptions1)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              //console.log(res.data);
                              setIPD_Name(res.data.Ipd.IPD_Name);
                              setDisease_Name(res.data.Disease.Disease_Name)
                              setDiseaseType(res.data.Disease.Disease_Type.DiseaseType)
                              setGenderType(res.data.Patient.Gender.Gender_Type)
                        }
                  });
      }
      
     
//=======================================================================================================================================
//function fethch data จาก backend
      const getMappigBed = async () => {
            const apiUrl = "http://localhost:8080/GetListMapBeds";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              setFiltertriages(res.data)
                              
                        }
                  });
      };

      const getTriages = async () => {
            const apiUrl = "http://localhost:8080/GetListTriages";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              //setGendersID(res.data.Patient.Gender_ID);
                              setTriages(res.data)
                        }
                  });
      };

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

      //========function useEffect ========
      useEffect(() => {
            getTriages();
            getZone();
            getBed();
            getMappigBed();
      }, [zoneID]);

//=======================================================================================================================================
//Uer inter face
//=======================================================================================================================================
      
      return(
            <Paper elevation={0}>
                  <Snackbar
                        open={success}
                        autoHideDuration={6000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                  >
                        <Alert onClose={handleClose} severity="success">
                              บันทึกข้อมูลสำเร็จ
                        </Alert>
                  </Snackbar>
                  <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="error">
                              บันทึกข้อมูลไม่สำเร็จ
                        </Alert>
                  </Snackbar>
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
                                          <FormControl fullWidth >
                                                <Select
                                                id="Patient_Name"
                                                value={triageID}
                                                displayEmpty
                                                inputProps={{ 'aria-label': 'Without label' }}
                                                onChange={onChangeTriage}
                                                >
                                                      <MenuItem value="">
                                                            กรุณาเลือกผู้ป่วย
                                                      </MenuItem>
                                                      {triages.map( triage => (
                                                            <MenuItem value={triage.ID} key = {triage.ID}>
                                                                  {triage.Patient.Patient_Name}
                                                            </MenuItem>
                                                      ))}
                                                </Select>
                                          </FormControl>  
                                                            
                                    </Grid>
                                    <Grid item xs={2} >
                                          
                                          <Button
                                                onClick={search}
                                                variant="contained"
                                                color="primary"
                                                sx={{marginTop : 8}}
                                          >
                                                ค้นหา
                                          </Button>
                                    
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>เพศ</p>
                                          <TextField
                                                fullWidth
                                                id="GenderType"
                                                value={GenderType}
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
                                                value={DiseaseType}
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
                                                value={Disease_Name}
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
                                                value={IPD_Name}
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
                                                            กรุณาเลือกโซน
                                                      </MenuItem>
                                                      {Zones.map( zone => (
                                                            <MenuItem value={zone.ID} key = {zone.ID}>{zone.Zone_Name}</MenuItem>
                                                      ))}
                                                </Select>
                                          </FormControl>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>เตียง</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="beds"
                                                      value={bedID}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={onChangeBed}
                                                      >
                                                            <MenuItem value="">
                                                                  กรุณาเลือกเตียง
                                                            </MenuItem>
                                                            {Beds.map( bed => (
                                                                  <MenuItem value={bed.ID} key = {bed.ID}>
                                                                        {bed.Bed_Name}
                                                                  </MenuItem>
                                                            ))}
                                                      </Select>
                                          </FormControl>
                                          
                                    </Grid>
                                    <Grid item xs={4}>
                                          <p>วันที่เข้ารับการรักษา</p>
                                          <FormControl fullWidth variant="outlined">
                                          <LocalizationProvider dateAdapter={AdapterDateFns}>
                                                <DatePicker
                                                      value={date}
                                                      onChange={(newValue) => {setDate(newValue);}}
                                                      renderInput={(params) => <TextField {...params} />}
                                                />
                                                </LocalizationProvider>
                                          </FormControl>
                                          <FormControl fullWidth variant="outlined">
                                                
                                          </FormControl>
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>หมายเหตุ</p>
                                          <FormControl fullWidth variant="outlined">
                                                <TextField
                                                id="MapBed_Comment"
                                                variant="outlined"
                                                type="string"
                                                size="medium"
                                                value={comments}
                                                defaultValue=""
                                                onChange={handleInputChange}
                                                />
                                          </FormControl>

                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>ผู้บันทึก</p>
                                          <FormControl fullWidth>
                                                <Select
                                                id="demo-select-small"
                                                value={zoneID}
                                                displayEmpty
                                                inputProps={{ 'aria-label': 'Without label' }}
                                                onChange={onChangeZone}
                                                >
                                                      <MenuItem value="">
                                                            -
                                                      </MenuItem>
                                                      
                                                </Select>
                                          </FormControl>
                                         
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <Button
                                                fullWidth
                                                onClick={submit}
                                                variant="contained"
                                                color="primary"
                                          >
                                                Submit
                                          </Button>
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                                    <Grid item xs={12}>
                                   
                                    </Grid>
                             
                              </Grid>
                        </Paper>
                  </Container>
            </Paper>
      );
}
export default MappingBedCreate;