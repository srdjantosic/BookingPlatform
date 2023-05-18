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
        var role=localStorage.getItem('role')
        const new_user = {password, firstName, lastName, email, username, address, role}

        console.log(userId)
        console.log(new_user)

        fetch("http://localhost:8080/api/user/update/"+ userId,{
        method:"PUT",
        body:JSON.stringify(new_user)
        }).then(() =>{
            alert("Successful update!")
        }).catch((err) => {
            console.log(err)
        });
      }


    const handleDelete = (e) =>{
        e.preventDefault()

        var userId=localStorage.getItem('userId');
        var role=localStorage.getItem('role')


        fetch("http://localhost:8080/api/user/delete/"+ userId + "/" + role,{
            method:"DELETE",
        }).then(() =>{
            alert("Successful delete!")
            window.location.href = "/UserLogin";
        }).catch((err) => {
            console.log(err)
        });
    }

    return(
      <body>
        <div class="topnav">
                <a href="/Homepage">Home Page</a>
                <a  href="/UserReservations">Reservations</a>
                
                <a class="active" href="/UserUpdate">Profile</a>
        </div>
        <div className="wrapper">
            <form>
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
            <button type="submit" onClick={handleDelete}>Delete account</button>
            </form>
        </div>
      </body>
    )
}