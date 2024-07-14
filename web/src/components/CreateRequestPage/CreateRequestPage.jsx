import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '../Header/Header';
import Button from '../Button/Button';
import ToggleButton from '../ToggleButton/ToggleButton';
import ReqForCoinsForm from './ReqForCoinsForm/ReqForCoinsForm'
import './CreateRequestPage.css';


const CreateRequestPage = () => {

    const navigate = useNavigate();
  
    const [selection, setSelection] = useState(null);
    const handleSelect = (selection) => {
      setSelection(selection);

    };
    






    return (
      <div className='create-request-page'>
        
        <Header title="Новая заявка" />
        <p>На что вы хотите подать заявку?</p>

        <ToggleButton
        leftLabel="Мерч"
        rightLabel="Simpl-коины"
        onSelect={handleSelect}
        />
        {selection === 'right' && <ReqForCoinsForm />}

        
      </div>
    );
  };
  
  export default CreateRequestPage;