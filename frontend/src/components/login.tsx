import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';


export default function MyComponent() {
      const [users, setUsers] = React.useState<any[]>([]);
      console.log(users)
      const getUsers = async () => {
            const apiUrl = "http://localhost:8080/GetListTriages";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
              .then((response) => response.json())
              .then((res) => {
                if (res.data.Patient) {
                  setUsers(res.data.Patient.Gender_ID);
                }
              });
          };
      React.useEffect(() => {
            getUsers();
      }, []);

  return (



      <Paper elevation={0}>
            <ul>
                  {users.map( item => (
                  <li key={item.ID}>
                        {item.Patient_Name}
                  </li>
                  ))}
            </ul> 


      </Paper>
    
      
      );
  }
   
