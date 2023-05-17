import {useState} from 'react';
import './styles/Login.css';
export default function ReserveApartment() {
    const[availabilityStartDate, setAvailabilityStartDate] = useState('')
    const[availabilityEndDate, setAvailabilityEndDate] = useState('')
    const[number, setNumber] = useState('');

  


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
          <fieldset>
            <label>
                    <p>Number of guests</p>
                    <input id="number" name="number" onChange={(e)=>setNumber(e.target.value)}/>
                </label>
            </fieldset>
            
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
            
            
            <button type="submit" >Submit</button>
            
        </form>
      </div>
      <div className="bodyImg"></div>
    
      </body>
    )
}