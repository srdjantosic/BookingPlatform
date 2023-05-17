import React, { useEffect, useState } from 'react'
import './styles/Login.css'
import { useNavigate } from 'react-router-dom'

function HostHomepage(){
    const [apartments, setApartments] = useState([]);
    const[availableApartments, setAvailableApartments] = useState([]);
  
    const[location, setLocation] = useState('');
    const[availabilityStartDate, setAvailabilityStartDate] = useState('');
    const[availabilityEndDate, setAvailabilityEndDate] = useState('');
    const[number, setNumber] = useState('');

    useEffect(()=>{

        fetch("http://localhost:8080/api/apartment/",{
        })
        .then(res =>res.json())
        .then((result)=> {
            console.log(result)
            setApartments(result);
        })
    }, []);


    const navigate = useNavigate();
  
    const navigateToAddNew = (e) =>{
        e.preventDefault()
        window.location.href = "/CreateApartment"
    }

    const search = (e) => {

        document.getElementById('toHide').hidden=true;
        document.getElementById('toShow').hidden=false;

        fetch("http://localhost:8080/api/apartment/filter/" + location + "/" + number + "/" + availabilityStartDate + "/" + availabilityEndDate, {}).then(res => res.json()).then((result) => {
            console.log(result)
            setAvailableApartments(result)
        }).catch((err) => {
            console.log(err);
        });
    }

    return(
        <body>
            <div class="topnav">
                <a class="active" href="/HostHomepage">Home Page</a>
                <a  href="/">All apartments</a>
                <a >Contracts</a>
                <a href="/UserUpdate">Profile</a>
             
            </div>
            <div class="wrapper">
                <table>
                    <tr>
                        <th>Location</th>
                        <th>Number of persons</th>
                        <th>Start date</th>
                        <th>End date</th>
                        <th></th>
                    </tr>
                    <tr>
                        <td><input id="location" name="location" onChange={(e)=>setLocation(e.target.value)}/></td>
                        <td><input id="number" name="number" onChange={(e)=>setNumber(e.target.value)}/></td>
                        <td><input id="availabilityStartDate" name="availabilityStartDate" onChange={(e)=>setAvailabilityStartDate(e.target.value)}/></td>
                        <td><input id="availabilityEndDate" name="availabilityEndDate" onChange={(e)=>setAvailabilityEndDate(e.target.value)}/></td>
                        <td><button onClick={search}> Search</button></td>
                    </tr>
                </table>
            </div>
            <div className='wrapper'>
                <table id="toHide">
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
                                        localStorage.setItem('Id', val.id)
                                        }}>View
                                    </button>
                                </td>
                            </tr>
                        )
                    })}
                </table>
                <table id="toShow" hidden='true'>
                    <tr>
                        <th>Name</th>
                        <th>Location</th>
                        <th>Benefits</th>
                        <th>Minimum number of guests</th>
                        <th>Maximum number of guests</th>
                        <th>Total Price</th>
                        <th>Unit Price</th>
                        <th></th>
                    </tr>
                    {availableApartments.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.Apartment.name}</td>
                                <td>{val.Apartment.location}</td>
                                <td>{val.Apartment.benefits}</td>
                                <td>{val.Apartment.minGuestsNumber}</td>
                                <td>{val.Apartment.maxGuestsNumber}</td>
                                <td>{val.TotalPrice}</td>
                                <td>{val.UnitPrice}</td>
                                <td>
                                    <button onClick={(e) => {
                                        e.preventDefault()
                                        localStorage.setItem('Id', val.id)
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

export default HostHomepage