import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import UserLogin from './components/UserLogin';
import UserRegistration from './components/UserRegistration';
import UserUpdate from './components/UserUpdate';
import Homepage from './components/Homepage';
import CreateApartment from './components/CreateApartment';

// const root = ReactDOM.createRoot(document.getElementById('root'));
// root.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>
// );

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
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