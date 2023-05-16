import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';

import UserLogin from './components/UserLogin';
import UserRegistration from './components/UserRegistration';
import UserUpdate from './components/UserUpdate';
import Homepage from './components/Homepage';
import CreateApartment from './components/CreateApartment';


ReactDOM.render(
  <Router>
    <Routes>
        <Route path='/' element={<UserLogin/>}/>
        <Route path='/UserRegistration' element = {<UserRegistration/>}/>
        <Route path='/UserUpdate/:id' element = {<UserUpdate/>}/>
        <Route path='/Homepage' element = {<Homepage/>}/>
        <Route path='/CreateApartment' element = {<CreateApartment/>}/>
     
    </Routes>
  </Router>,
 document.getElementById('root')
);