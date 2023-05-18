import React, { useEffect, useState } from 'react'
import './styles/Login.css'

const HostReservations = () => {
    const [requests, setRequests] = useState([]);

    var hostId=localStorage.getItem("userId")

    useEffect(()=>{
        fetch("http://localhost:8080/api/user/reservationRequests/requests/getRequests/"+hostId,{
        }).then(res =>res.json()).then((result)=> {
            setRequests(result);
        }).catch((err)=>{
            console.log(err)
        })
    }, []);

    const handleNo = (e) =>{

        var req_id = localStorage.getItem('requestId')


        fetch("http://localhost:8080/api/user/request/accept/reservationRequest/"+req_id,{
            method:"DELETE",
        }).then(() =>{
            alert("Successful deletion!")
        }).catch((err) => {
            console.log(err)
        });
    }

  return(
        <body>
            <div class="topnav">
                <a  href="/HostHomepage">Home Page</a>
                <a href="/HostApartments">View my apartments</a>
                <a  href="/CreateApartment">Add new apartment</a>
                <a  className="active" href="/HostReservations">Reservation requests</a>
                <a href="/HostUpdate">Profile</a>
            </div>
            <div className='wrapper'>
                <h1>Requests</h1>
            </div>
            <div className='wrapper'>
                <table>
                    {requests.map((val, key) => {
                        return(
                            <tr key={key}>
                                <table>
                                    <tr>
                                        <th>
                                            UserId
                                        </th>
                                        <th>ApartmentId</th>
                                        <th></th>
                                        <th></th>
                                        <th></th>
                                        <th></th>
                                        <th></th>
                                    </tr>
                                    {val.map((val2, key2)=>{
                                        return(
                                            <tr>
                                                <td>{val2.userId}</td>
                                                <td>{val2.apartmentId}</td>
                                                <td>{val2.startDate}</td>
                                                <td>{val2.endDate}</td>
                                                <td>{val2.guestsNumber}</td>
                                                <td>
                                                    <button onClick={(e) => {
                                                        e.preventDefault()
                                                        localStorage.setItem('requestId', val2.id)
                                                        // handleYes()
                                                        var req_id = localStorage.getItem('requestId')
                                                        console.log("BUTTON YES")

                                                        fetch("http://localhost:8080/api/user/request/accept/reservationRequest/"+req_id,{
                                                            method:"POST",
                                                        }).then(() =>{
                                                            alert("Request accepted!")
                                                        }).catch((err) => {
                                                            console.log(err)
                                                        });
                                                    }}
                                                    >Yes</button>
                                                </td>
                                                <td>
                                                    <button  onClick={(e) => {
                                                        e.preventDefault()
                                                        localStorage.setItem('requestId', val2.id)
                                                        // handleYes()
                                                        var req_id = localStorage.getItem('requestId')


                                                        fetch("http://localhost:8080/api/user/delete/request/"+req_id,{
                                                            method:"DELETE",
                                                        }).then(() =>{
                                                            alert("Request rejected!")
                                                        }).catch((err) => {
                                                            console.log(err)
                                                        });
                                                    }}> No</button>
                                                </td>
                                            </tr>
                                        )
                                    })}
                                </table>
                            </tr>
                        )
                    })}
                </table>
            </div>
        </body>
    
  )
}
export default HostReservations