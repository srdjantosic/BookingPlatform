import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';

import UserLogin from './components/UserLogin';
import UserRegistration from './components/UserRegistration';
import UserUpdate from './components/UserUpdate';
import Homepage from './components/Homepage';
import CreateApartment from './components/CreateApartment';
import DefineAvailableTerm from './components/DefineAvailableTerm';
import HostHomepage from './components/HostHomepage';
import HostUpdate from './components/HostUpdate';
import HostReservations from './components/HostReservations';
import SeeApartment from './components/SeeApartment';
import UserReservations from './components/UserReservations';
import ReserveApartment from './components/ReserveApartment';
import SearchedApartments from './components/SearchedApartments';


ReactDOM.render(
  <Router>
    <Routes>
        <Route path='/' element={<UserLogin/>}/>
        <Route path='/UserRegistration' element = {<UserRegistration/>}/>
        <Route path='/UserUpdate' element = {<UserUpdate/>}/>
        <Route path='/Homepage' element = {<Homepage/>}/>
        <Route path='/CreateApartment' element = {<CreateApartment/>}/>
        <Route path='/DefineAvailableTerm' element = {<DefineAvailableTerm/>}/>
        <Route path='/HostHomepage' element = {<HostHomepage/>}/>
        <Route path='/HostUpdate' element = {<HostUpdate/>}/>
        <Route path='/HostReservations' element = {<HostReservations/>}/>
        <Route path='/SeeApartment' element = {<SeeApartment/>}/>
        <Route path='/UserReservations' element={<UserReservations/>}/>
        <Route path='/SearchedApartments' element={<SearchedApartments/>}/>

        
    </Routes>
  </Router>,
 document.getElementById('root')
);