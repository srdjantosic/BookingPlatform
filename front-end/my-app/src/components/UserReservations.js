import React, { useEffect, useState } from 'react'
import './styles/Login.css'
import { useNavigate } from 'react-router-dom'



const UserReservations = () => {
  const [apartments, setApartments] = useState([]);
  
  const[location, setLocation] = useState('');
  const[availabilityStartDate, setAvailabilityStartDate] = useState('');
  const[availabilityEndDate, setAvailabilityEndDate] = useState('');
  const[number, setNumber] = useState('');







  useEffect(()=>{

    fetch("http://localhost:8080/api/user/getReservations/"+ localStorage.getItem('userId'),{
    })
    .then(res =>res.json())
    .then((result)=>
    {
        setApartments(result);
    }
    )
  }, [])

  const navigate = useNavigate();
  
  const navigateToAddNew = (e) =>{
    e.preventDefault()
    window.location.href = "/CreateApartment"
}


const deleteReservation = (e) =>{
    var Id = localStorage.getItem('Id');
      
      fetch("http://localhost:8080/api/user/deleteReservation/"+Id,{
          method:"DELETE",
          headers : { 
            'Content-Type': 'application/json'
            
           },
      
        }).then(() =>{
         window.location.href = "/UserReservations"
        })
     
  }
  


  return(
   
   
    
        <body>
            <div class="topnav">
                <a class="active" href="/Homepage">Home Page</a>
                <a  href="/UserReservations">Reservations</a>
                <a >Contracts</a>
                <a href="/UserUpdate">Profile</a>
             
            </div>
           
            
            <div className='wrapper'>
                <table>
                    <tr>
                        <th>Apartment id</th>
                        <th>Start date</th>
                        <th>End date</th>
                        <th>Number of guests</th>
                      
                    </tr>
                    {apartments.map((val, key) => {
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

                                       deleteReservation();
                                        
                                    }}>Delete</button>
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