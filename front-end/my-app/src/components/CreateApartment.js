import {useState} from 'react';
import './styles/Login.css';

export default function CreateApartment() {
  

    const[name, setName] = useState('')
    const[location, setLocation] = useState('')
    const[benefits, setBenefits] = useState('')
    const[minGuestsNumber, setMinGuestsNumber] = useState('')
    const[maxGuestsNumber, setMaxGuestsNumber] = useState('')
    const[automaticReservation, setAutomaticReservation] = useState('')

    const handleClick = (e) =>{
        var userId=localStorage.getItem('userId');
        var role= localStorage.getItem('role');
        var regexPattern = new RegExp("true");
        var boolValue1 = regexPattern.test(automaticReservation);
        e.preventDefault()
        var number = parseInt(minGuestsNumber, 10 );
        var number1 = parseInt(maxGuestsNumber, 10 );
        const new_apartment = {name, location, benefits, number, number1,boolValue1};
  
        fetch("http://localhost:8080/api/apartment/insert/"+role+"/"+userId,{ 
        method:"POST",
       // headers:{},
        body:JSON.stringify(new_apartment)
      }).then(() =>{
        alert("Successful created!")
        //window.location.href = '/';
        console.log(new_apartment)
      }).catch((err) => {
        console.log(err)
      });
      }
      const handleClick1 = (e) =>{
        
        window.location.href = '/DefineAvailableTerm';
     
      }

    return(
      <body>
        <div class="topnav">
                <a href="/Homepage">Home Page</a>
                <a class="active" href="/">Your orders</a>
                <a >Contracts</a>
                <a href="/UserUpdate">Profile</a>
             
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
                    <p>Automatic reservation</p>
                    <input type='automaticReservation' id="automaticReservation" name="automaticReservation" onChange={(e)=>setAutomaticReservation(e.target.value)}/>
                </label>
            </fieldset>
            
            <button type="submit" onClick={handleClick}>submit</button>
        
        </form>
        
        <button type="submit" onClick={handleClick1}>Define</button>
      </div>
      <div className="bodyImg"></div>
      <div className="wrapper">@Chocolate Factory Novi Sad since 2000</div>
      </body>
    )
}