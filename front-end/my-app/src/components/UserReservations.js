import React, { useEffect, useState } from 'react'
import './styles/Login.css'
import { useNavigate } from 'react-router-dom'



const UserReservations = () => {

  const[reservations, setReservations] = useState('');
  const[availabilityStartDate, setAvailabilityStartDate] = useState('');
  const[availabilityEndDate, setAvailabilityEndDate] = useState('');
  const[number, setNumber] = useState('');







  useEffect(()=>{

    fetch("http://localhost:8080/api/user/getReservations/"+ localStorage.getItem('userId'),{
    })
    .then(res =>res.json())
    .then((result)=>
    {
        setReservations(result);
    }
    )
  }, [])


  




  return(
   
   
    
        <body>
            <div class="topnav">
                <a  href="/Homepage">Home Page</a>
                <a class="active" href="/UserReservations">Reservations</a>
                <a >Contracts</a>
                <a href="/UserUpdate">Profile</a>
             
            </div>
           
           
            <div className='wrapper'>
                <table>
                    <tr>
                        
                        <th>Apartment</th>
                        <th>Start date</th>
                        <th>End date </th>
                        <th>Number of guests</th>
                        
                    </tr>
                    {reservations.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.apartmentId}</td>
                                <td>{val.startDate}</td>
                                <td>{val.endDate}</td>
                                <td>{val.guestsNumber}</td>
                            
                                <td>
                                    <button onClick={(e) => {
                                        e.preventDefault()

                                        localStorage.setItem('Id', val.id)

                                        //navigate(`/SeeApartmet/`+val.id);
                                        
                                    }}>Cancel</button>
                                </td>
                                
                                
                            </tr>
                        )
                    })}
                </table>

            </div>
            
           
        </body>
    
  )
 
}

export default UserReservations