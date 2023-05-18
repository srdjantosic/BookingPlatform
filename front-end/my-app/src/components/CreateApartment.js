import React, {useState} from 'react';
import './styles/Login.css';

export default function CreateApartment() {

    const[name, setName] = useState('')
    const[location, setLocation] = useState('')
    const[benefits, setBenefits] = useState('')
    const[minNumber, setMinGuestsNumber] = useState('')
    const[maxNumber, setMaxGuestsNumber] = useState('')
    const[autoReservation, setAutomaticReservation] = useState('')

    const handleClick = (e) =>{
        e.preventDefault()

        var userId=localStorage.getItem('userId');
        var role= localStorage.getItem('role');

        var regexPattern = new RegExp("true");
        var automaticReservation = regexPattern.test(autoReservation);
        var minGuestsNumber = parseInt(minNumber);
        var maxGuestsNumber = parseInt(maxNumber );
        const new_apartment = {name, location, benefits, minGuestsNumber, maxGuestsNumber, automaticReservation};
  
        fetch("http://localhost:8080/api/apartment/insert/"+role+"/"+userId,{ 
        method:"POST",
        body:JSON.stringify(new_apartment)
      }).then(() =>{
        alert("Successful created!")
        window.location.href = "/HostHomepage";
        console.log(new_apartment)
      }).catch((err) => {
        console.log(err)
      });
      }

    return(
      <body>
        <div class="topnav">
            <a  href="/HostHomepage">Home Page</a>
            <a  href="/HostApartments">View my apartments</a>
            <a className='active' href="/CreateApartment">Add new apartment</a>
            <a href="/HostReservations">Reservation requests</a>
            <a href="/HostUpdate">Profile</a>
        </div>
        <div className="wrapper">
        <form >
          <h1>Create new apartment</h1>
          <fieldset>
                <label>
                    <p> Name</p>
                    <input id="name" name="name" onChange={(e)=>setName(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Location</p>
                    <input id="location" name="location" onChange={(e)=>setLocation(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Benefits</p>
                    <input id="benefits" name="benefits" onChange={(e)=>setBenefits(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Minimum number of guests</p>
                    <input id="minGuestsNumber" name="minGuestsNumber" onChange={(e)=>setMinGuestsNumber(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Maximum number of guests</p>
                    <input type='maxGuestsNumber' id="maxGuestsNumber" name="maxGuestsNumber" onChange={(e)=>setMaxGuestsNumber(e.target.value)}/>
                </label>
            </fieldset>
            
            
            <fieldset>
            <label>
        Automatic Reservation
      <select  onClick={(e)=>setAutomaticReservation(e.target.value)}>
        <option value="false">False</option>
        <option value="true">True </option>
        
      </select>
       </label>     
       </fieldset>
            <button type="submit" onClick={handleClick}>Create</button>
        
        </form>
        </div>
      </body>
    )
}