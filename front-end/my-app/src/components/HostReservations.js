import React, { useEffect, useState } from 'react'
import './styles/Login.css'
import { useNavigate } from 'react-router-dom'



const HostReservations = () => {
  const [apartments, setApartments] = useState([]);
  
  const[location, setLocation] = useState('');
  const[availabilityStartDate, setAvailabilityStartDate] = useState('');
  const[availabilityEndDate, setAvailabilityEndDate] = useState('');
  const[number, setNumber] = useState('');







  useEffect(()=>{

    fetch("http://localhost:8080/api/apartment/",{
    })
    .then(res =>res.json())
    .then((result)=>
    {
        setApartments(result);
    }
    )
  }, [])


  




  return(
   
   
    
        <body>
            <div class="topnav">
                <a  href="/HostHomepage">Home Page</a>
                <a  href="/CreateApartment">Add new apartment</a>
                <a class="active" href="/HostReservations">Reservations</a>
                <a href="/HostUpdate">Profile</a>
             
            </div>
           
           
            <div className='wrapper'>
                <table>
                    <tr>
                        <th>User</th>
                        <th>Apartment</th>
                        <th>Start date</th>
                        <th>End date </th>
                        <th>Number of guests</th>
                        
                    </tr>
                    {apartments.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.name}</td>
                                <td>{val.location}</td>
                                <td>{val.benefits}</td>
                                <td>{val.minGuestsNumber}</td>
                                <td>{val.maxGuestsNumber}</td>
                            
                                <td>
                                    <button onClick={(e) => {
                                        e.preventDefault()

                                        localStorage.setItem('Id', val.id)

                                        //navigate(`/SeeApartmet/`+val.id);
                                        
                                    }}>Accept</button>
                                </td>
                                <td>
                                    <button onClick={(e) => {
                                        

                                        localStorage.setItem('Id', val.id)

                                      //  navigate(`/UpdateOrder/`+val.id);
                                        
                                    }}>Deny</button>
                                </td>
                                
                            </tr>
                        )
                    })}
                </table>

            </div>
            
           
        </body>
    
  )
 
}

export default HostReservations