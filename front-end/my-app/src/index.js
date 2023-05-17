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
        
    </Routes>
  </Router>,
 document.getElementById('root')
);