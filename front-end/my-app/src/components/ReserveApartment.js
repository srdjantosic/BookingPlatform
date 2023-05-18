import {useState} from 'react';
import './styles/Login.css';
export default function ReserveApartment() {
    const[startDate, setAvailabilityStartDate] = useState('')
    const[endDate, setAvailabilityEndDate] = useState('')
    const[guestNumber, setGuestNumber] = useState('');
  
    
    

  var guestId=localStorage.getItem('userId')
  var apartmentId= localStorage.getItem('Id');

    const handleClick = (e) =>{
      var guestsNumber= parseInt(guestNumber,10)
      e.preventDefault()
      const new_reservation = {startDate, endDate, guestsNumber, guestId,apartmentId}
console.log(new_reservation);
      fetch("http://localhost:8080/api/user/insertReservation",{ 
      method:"POST",
     // headers:{"Content-Type":"application/json"},
      body:JSON.stringify(new_reservation)
    }).then(() =>{
      alert("Reservation sent!")
      window.location.href = '/UserReservations';
    }).catch((err) => {
      console.log(err)
    });
    }
    

    return(
      <body>
        <div class="topnav">
                <a class="active" href="/Homepage">Home Page</a>
                <a  href="/UserReservations">Reservations</a>
                <a >Contracts</a>
                <a  href="/UserUpdate">Profile</a>
             
            </div>
        <div className="wrapper">
        <form >
          <h1>Reserve appartment</h1>
         
          <fieldset>
            <label>
                    <p>Number of guests</p>
                    <input id="number" name="number" onChange={(e)=>setGuestNumber(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
              <label>
                    <p> Start date </p>
                    <input id="availabilityStartDate" name="availabilityStartDate" onChange={(e)=>setAvailabilityStartDate(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
            <label>
                    <p>End date</p>
                    <input id="availabilityEndDate" name="availabilityEndDate" onChange={(e)=>setAvailabilityEndDate(e.target.value)}/>
                </label>
            </fieldset>
            
            
            <button type="submit" onClick={handleClick}>Submit</button>
            
        </form>
      </div>
      <div className="bodyImg"></div>
    
      </body>
    )
}