import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '../Header/Header';
import Button from '../Button/Button';
import '../MainPage/MainPage.css'
// import  './MainPage.css';

const MainPage = () => {

  const navigate = useNavigate();
  

  
  return (
    <div>
      
      <Header title="Мои заявки" />
      <Button text="Создать заявку" onClick={() => navigate('/create')}/>
      
      <p className='empty'>Тут пусто...пока</p>
      
    </div>
  );
};

export default MainPage;