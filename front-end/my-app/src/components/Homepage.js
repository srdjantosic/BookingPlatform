import React, { useEffect, useState } from 'react'
import './styles/Login.css'
import { useNavigate } from 'react-router-dom'



const Homepage = () => {
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


  const navigate = useNavigate();
  
  const navigateToAddNew = (e) =>{
    e.preventDefault()
    window.location.href = "/CreateApartment"
}

const search = (e) =>{
 
    fetch("http://localhost:8080/api/apartment/filter/"+location+"/"+number+"/"+availabilityStartDate+"/"+availabilityEndDate,{
    
      }).then((result) =>{
        console.log(result)
        //setApartments(result);
        
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
           
            <div class="searcing">
            <label>
                    <p> Location</p>
                    <input id="location" name="location" onChange={(e)=>setLocation(e.target.value)}/>
                </label>
                <label>
                    <p> Number of persons</p>
                    <input id="number" name="number" onChange={(e)=>setNumber(e.target.value)}/>
                </label>
                <label>
                    <p> Start date</p>
                    <input id="availabilityStartDate" name="availabilityStartDate" onChange={(e)=>setAvailabilityStartDate(e.target.value)}/>
                </label>
                <label>
                    <p> End date</p>
                    <input id="availabilityEndDate" name="availabilityEndDate" onChange={(e)=>setAvailabilityEndDate(e.target.value)}/>
                </label>
                <button onClick={search}> Search</button>
            </div>
            <div className='wrapper'>
                <table>
                    <tr>
                        <th>Name</th>
                        <th>Location</th>
                        <th>Benefits</th>
                        <th>Minimum number of guests</th>
                        <th>Maximum number of guests</th>
                        <th>Final price</th>
                        <th>Price per person</th>
                        
                    </tr>
                    {apartments.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.name}</td>
                                <td>{val.location}</td>
                                <td>{val.benefits}</td>
                                <td>{val.minGuestsNumber}</td>
                                <td>{val.maxGuestsNumber}</td>
                                <td>{val.minGuestsNumber}</td>
                                <td>{val.maxGuestsNumber}</td>
                                <td>
                                    <button onClick={(e) => {
                                        e.preventDefault()

                                        localStorage.setItem('Id', val.id)

                                        //navigate(`/SeeApartmet/`+val.id);
                                        
                                    }}>View</button>
                                </td>
                                <td>
                                    <button onClick={(e) => {
                                        

                                        localStorage.setItem('Id', val.id)

                                       navigate(`/ReserveApartment`);
                                        
                                    }}>Reserve</button>
                                </td>
                                
                            </tr>
                        )
                    })}
                </table>

            </div>
            
           
        </body>
    
  )
 
}

export default Homepage