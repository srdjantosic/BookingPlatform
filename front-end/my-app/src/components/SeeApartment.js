import React, { useEffect, useState } from 'react'
import './styles/Login.css'

function SeeApartment(){
    const [apartment, setApartment] = useState([]);
    const [items, setItems] = useState([]);

    var apartmentId=localStorage.getItem('apartmentId')

    useEffect(()=>{

        fetch("http://localhost:8080/api/apartment/getOne/"+apartmentId,{
        }).then(res =>res.json()).then((result)=> {
            console.log(result)
            setApartment(result);
            setItems(result.pricelist);
        }).catch((err)=>{
            console.log(err)
        })
    }, []);

    const handleClick = (e)=>{
        e.preventDefault()
        window.location.href='/DefineAvailableTerm'
    }

    return(
        <body>
        <div className="topnav">
            <a  href="/HostHomepage">Home Page</a>
            <a className="active" href="/HostApartments">View my apartments</a>
            <a  href="/CreateApartment">Add new apartment</a>
            <a href="/HostReservations">Reservations</a>
            <a href="/HostUpdate">Profile</a>
        </div>
        <div className='wrapper'>
            <h1>Selected apartment</h1>
        </div>
        <div className='wrapper'>
            <table>
                <tr>
                    <th>Name</th>
                    <th>Location</th>
                    <th>Benefits</th>
                    <th>Minimum number of guests</th>
                    <th>Maximum number of guests</th>
                    <th>Pricelist</th>
                </tr>
                <tr>
                    <td>{apartment.name}</td>
                    <td>{apartment.location}</td>
                    <td>{apartment.benefits}</td>
                    <td>{apartment.minGuestsNumber}</td>
                    <td>{apartment.maxGuestsNumber}</td>
                    <td>


                  {
                      items.map((val, key)=>{

                            var priceUnit = val.unitPrice
                            var priceUnitString = "Price for apartment"
                            if (priceUnit == "0"){
                            priceUnitString = "Price per person"
                        }

                            return(
                            <div>
                            <p><strong>Pricelist item:</strong></p>
                            <p>Start date: {val.availabilityStartDate}</p>
                            <p>End date: {val.availabilityEndDate}</p>
                            <p>Price: {val.price}</p>
                            <p>Price unit: {priceUnitString}</p>
                            </div>
                            )

                        })}

                    </td>
                </tr>
            </table>
        </div>
        <div className="wrapper">
            <button onClick={handleClick}>Add new pricelist item</button>
        </div>
        </body>
    )
}

export default SeeApartment