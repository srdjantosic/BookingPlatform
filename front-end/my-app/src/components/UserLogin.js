import { useState } from "react";
import './styles/Login.css';

export default function UserLogin(){
    const[username, setUserName] = useState('');
    const[password, setPassword] = useState('');

   


    const handleSubmit = (e) =>{
      e.preventDefault()
    
      fetch("http://localhost:8080/api/user/" + username+"/"+password,{
      })
      .then(res =>res.json())
      .then((result)=>
      {
        localStorage.setItem('user_userName', username);
        localStorage.setItem('userId',result.data.id)
        localStorage.setItem('role',result.data.role)
        console.log(localStorage.getItem('role'))
        
        console.log(localStorage.getItem('user_userName'));
          if(localStorage.getItem('role')==="HOST"){
            window.location.href='/HostHomepage';
          } else{
        console.log(localStorage.getItem('userId'));
window.location.href='/Homepage';
          }
      }
      )
    };
    
 return(

    <body>
    <div className="wrapper">
      <form onSubmit={handleSubmit}>
        <h1>Booking Platform</h1>
        <fieldset>
              <label>
                  <p>Username</p>
                  <input id="userName" name="userName" onChange={(e)=>setUserName(e.target.value)}/>
              </label>
          </fieldset>
          <fieldset>
              <label>
                  <p>Password</p>
                  <input type="password" id="password" name="password" onChange={(e)=>setPassword(e.target.value)}/>
              </label>
          </fieldset>
          <button type="submit">Login</button>
      </form>
    </div>
    <div className="wrapper">
        Create an account? <a href="/UserRegistration">Sing Up</a>
    </div>
  </body>
 );

}