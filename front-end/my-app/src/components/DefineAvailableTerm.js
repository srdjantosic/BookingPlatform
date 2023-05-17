import {useState} from 'react';
import './styles/Login.css';

export default function DefineAvailableTerm() {

    const[availabilityStartDate, setAvailabilityStartDate] = useState('')
    const[availabilityEndDate, setAvailabilityEndDate] = useState('')
    const[price, setPrice] = useState('')
    const[unitPrice, setUnitPrice] = useState('')
    var [productId, setProductId] = useState('');

    const handleClick = (e) =>{
        e.preventDefault()
        var number = parseInt(price, 10 );
        var number1 = parseInt(unitPrice, 10 );
        const new_user = {availabilityStartDate, availabilityEndDate, number, number1}
        console.log(new_user)
        fetch("http://localhost:8080/api/apartment/insertItem/"+productId+"/"+localStorage.getItem('role'),{ 
        method:"POST",
       // headers:{"Content-Type":"application/json"},
        body:JSON.stringify(new_user)
      }).then(() =>{
        alert("Successful registration!")
      
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
        <div className="wrapper">
        <form >
          <h1>Define free term</h1>

            <fieldset>
            <label>
    Choose apartment
      <select  onClick={(e)=>setProductId(e.target.value)}>
        <option value="64649928bd67705a903c9f03">Lux Apartmani</option>
        <option value="6464a00161417e4661d07d07">Apartmani Zavicaj</option>
      </select>
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
            <fieldset>
                <label>
                    <p>Price</p>
                    <input id="price" name="price" onChange={(e)=>setPrice(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Unit price</p>
                    <input id="unitPrice" name="unitPrice" onChange={(e)=>setUnitPrice(e.target.value)}/>
                </label>
            </fieldset>
           
            
            <button type="submit" onClick={handleClick}>submit</button>

        </form>
      </div>
 
      
      </body>
    )
}