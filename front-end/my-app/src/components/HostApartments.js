import React, { useEffect, useState } from 'react'
import './styles/Login.css'

function HostApartments(){
    const [apartments, setApartments] = useState([]);

    var hostId = localStorage.getItem('userId')

    useEffect(()=>{

        fetch("http://localhost:8080/api/apartment/getByHostId/"+hostId,{
        }).then(res =>res.json()).then((result)=> {
                console.log(result)
                setApartments(result);
        }).catch((err)=>{
            console.log(err)
        })
    }, []);

    return(
        <body>
        <div className="topnav">
            <a  href="/HostHomepage">Home Page</a>
            <a className="active" href="/HostApartments">View my apartments</a>
            <a  href="/CreateApartment">Add new apartment</a>
            <a href="/HostReservations">Reservations Requests</a>
            <a href="/HostUpdate">Profile</a>
        </div>
        <div className='wrapper'>
            <h1>Apartments</h1>
        </div>
        <div className='wrapper'>
            <table>
                <tr>
                    <th>Name</th>
                    <th>Location</th>
                    <th>Benefits</th>
                    <th>Minimum number of guests</th>
                    <th>Maximum number of guests</th>
                    <th></th>
                </tr>
                {apartments.map((val, key) => {
                    return(
                        <tr key={key} >
                            <td>{val.name}</td>
                            <td>{val.location}</td>
                            <td>{val.benefits}</td>
                            <td>{val.minGuestsNumber}</td>
                            <td>{val.maxGuestsNumber}</td>
                            <td>
                                <button onClick={(e) => {
                                    e.preventDefault()
                                    localStorage.setItem('apartmentId', val.id)
                                    window.location.href="/SeeApartment"
                                }}>View
                                </button>
                            </td>
                        </tr>
                    )
                })}
            </table>
        </div>
        </body>
    )
}

export default HostApartments