import {useState} from 'react';
import './styles/Login.css';

export default function DefineAvailableTerm() {

    const[availabilityStartDate, setAvailabilityStartDate] = useState('')
    const[availabilityEndDate, setAvailabilityEndDate] = useState('')
    const[priceString, setPrice] = useState('')
    const[unitPriceString, setUnitPrice] = useState('')

    var apartmentId=localStorage.getItem('apartmentId')
    var role = localStorage.getItem('role')

    const handleClick = (e) =>{
        e.preventDefault()

        var price = parseInt(priceString)
        var unitPrice = parseInt(unitPriceString)

        const new_user = {availabilityStartDate, availabilityEndDate, price, unitPrice}

        console.log(new_user)
        fetch("http://localhost:8080/api/apartment/insertItem/"+apartmentId+"/"+role,{
        method:"POST",
        body:JSON.stringify(new_user)
      }).then(() =>{
            alert("Successful create!")
            window.location.href='/SeeApartment'
      }).catch((err) => {
            console.log(err)
      });
      }

    return(
      <body>
        <div class="topnav">
                <a href="/HostHomepage">Home Page</a>
                <a class="active" href="/">Add new apartment</a>
               <a href="/HostReservations">Reservations</a>
                <a href="/HostUpdate">Profile</a>
             
        </div>
        <div className='wrapper'>
            <h1>Add new item</h1>
        </div>
        <div className="wrapper">
        <form>
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
            <fieldset>
                <label>
                    <p>Price</p>
                    <input id="price" name="price" onChange={(e)=>setPrice(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
            <select  onClick={(e)=>setUnitPrice(e.target.value)}>
        <option value="0">Price per person</option>
        <option value="1">Price per apartment </option>
        </select>
            </fieldset>
            <button type="submit" onClick={handleClick}>Create</button>
        </form>
      </div>
      </body>
    )
}