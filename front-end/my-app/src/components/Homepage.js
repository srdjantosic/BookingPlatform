
import './styles/Login.css';
import { useNavigate } from 'react-router-dom'
export default function Homepage() {
  

    
    const navigate = useNavigate();
    
    const handleClick = (e) =>{
        navigate(`/UserUpdate/`+localStorage.getItem('userId'));
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
          <h1>User update</h1>
         
            
            
            <button type="submit" onClick={handleClick}>Update</button>
        </form>
      </div>
      <div className="bodyImg"></div>
    
      </body>
    )
}