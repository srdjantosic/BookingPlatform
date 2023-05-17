import {useState} from 'react';
import './styles/Login.css';
export default function UserUpdate() {
  
    const[password, setPassword] = useState('');
    const[firstName, setFirstName] = useState('');
    const[lastName, setLastName] = useState('');
    const[email, setEmail] = useState('');
    const[username, setUsername] = useState('');
    const[address, setAddress] = useState('');
   
   
  


    const handleClick = (e) =>{
        e.preventDefault()
        var userId=localStorage.getItem('userId');
        const new_user = {password, firstName, lastName, email, username, address}
  
        fetch("http://localhost:8080/api/user/update/"+ userId,{ 
        method:"PUT",
        //headers:{  'Content-Type': 'application/json'},
        body:JSON.stringify(new_user)
      }).then(() =>{
        alert("Successful registration!")
        //window.location.href = '/';
      }).catch((err) => {
        console.log(err)
      });
      }

  

    return(
      <body>
        <div className="wrapper">
        <form onSubmit={handleClick}>
          <h1>User update</h1>
          <fieldset>
                <label>
                    <p>First Name</p>
                    <input id="firstName" name="firstName" onChange={(e)=>setFirstName(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Last Name</p>
                    <input id="lastName" name="lastName" onChange={(e)=>setLastName(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>UserName</p>
                    <input id="username" name="username" onChange={(e)=>setUsername(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Email</p>
                    <input id="email" name="email" onChange={(e)=>setEmail(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Password</p>
                    <input type='password' id="password" name="password" onChange={(e)=>setPassword(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Address</p>
                    <input type='address' id="address" name="address" onChange={(e)=>setAddress(e.target.value)}/>
                </label>
            </fieldset>
            
            <button type="submit" onClick={handleClick}>Update</button>
        </form>
      </div>
      <div className="bodyImg"></div>
    
      </body>
    )
}