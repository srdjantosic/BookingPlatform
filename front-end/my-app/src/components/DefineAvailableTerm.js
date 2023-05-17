import {useState} from 'react';
import './styles/Login.css';
import MenuItem from '@mui/material/MenuItem';
export default function DefineAvailableTerm() {
  

    const[availabilityStartDate, setAvailabilityStartDate] = useState('')
    const[availabilityEndDate, setAvailabilityEndDate] = useState('')
    const[price, setPrice] = useState('')
    const[unitPrice, setUnitPrice] = useState('')
    var [productId, setProductId] = useState('');
    
    const handleChange = (event) => {
        setProductId(event.target.value);
      };

    const handleClick = (e) =>{
       
      }
      const handleClick1 = (e) =>{
        
     
      }

    return(
      <body>
        <div className="wrapper">
        <form >
          <h1>Define free term</h1>


          <inputlabel id="demo-simple-select-label">Product name</inputlabel>
<select
  labelId="demo-simple-select-label"
  id="demo-simple-select"
  value={productId}
  label="Age"
  onChange={handleChange}
>
  <MenuItem value={"64649928bd67705a903c9f03"}>Lux Apartmani</MenuItem>
  <MenuItem value={"64649928bd67705a903c9f03"}>Apartmani zavicaj</MenuItem>
  
 
</select>
          <fieldset>
                <label>
                    <p> Start date :</p>
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
      <div className="bodyImg"></div>
      <div className="wrapper">@Chocolate Factory Novi Sad since 2000</div>
      </body>
    )
}