import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { ChakraProvider } from '@chakra-ui/react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Home } from './Home/Home';
import { Customer } from './Customer/Customer';
import { Insurances } from './Insurances/Insurances';
import { Contracts } from './Contracts/Contracts';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <ChakraProvider>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/customers' element={<Customer />} />
          <Route path='/insurances' element={<Insurances />} />
          <Route path='/contracts' element={<Contracts />} />
        </Routes>
      </BrowserRouter>
    </ChakraProvider>
  </React.StrictMode>
);
