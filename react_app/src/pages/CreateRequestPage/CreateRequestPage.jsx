import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '../../components/Header/Header';
import Button from '../../components/Button/Button';
import ToggleButton from '../../components/ToggleButton/ToggleButton';
import ReqForCoinsForm from './ReqForCoinsForm/ReqForCoinsForm'
import ReqForMerchForm from './ReqForMerchForm/ReqForMerchForm'
import './CreateRequestPage.css';

//TODO: сделать так, чтобы кнопка отправки и комментарий появлялись только при загрузке достижений

const CreateRequestPage = () => {

    const navigate = useNavigate();
    const [selection, setSelection] = useState(null);
    // const [isReadyToSubmit, setIsReadyToSubmit] = useState(false);
    const [selectedAchievement, setSelectedAchievement] = useState('');
    const [selectedMerch, setSelectedMerch] = useState('');
    const [comment, setComment] = useState('');

    
    const handleSelect = (selection) => {
      setSelection(selection);

    };
    
    const handleAchievementChange = (achievementId) => {
      setSelectedAchievement(achievementId);
    };
  
    const handleCommentChange = (event) => {
      setComment(event.target.value);
    };
    const handleMerchChange = (merchId) => {
      setSelectedMerch(merchId);
    };



    const submitRequest = async (request, path) => {
      try {
           const response = await fetch(import.meta.env.VITE_REQUESTS_SERVICE_URL + path, {
             method: 'POST',
             headers: {
               'Content-Type': 'application/json',
             },
             body: JSON.stringify(request),
           });
    
           if (response.ok) {
             console.log('Application submitted successfully');
    
           } else {
             console.error('Failed to submit request');
           }
         } catch (error) {
           console.error('Error submitting request:', error);
         }


    }

    const handleSubmit = () => {

      const userId = sessionStorage.getItem('userId');

      if (selection==="right") {

        const CoinReq = {
          userId: userId, 
          userMessage: comment,
          achievementId: selectedAchievement    
        };

        submitRequest(CoinReq, '/coinsrequests')      
        alert (JSON.stringify(CoinReq))

    } else {

        const MerchReq = {
          userId: userId, 
          userMessage: comment,
          merchId: 1
        };

        submitRequest(MerchReq, '/merchrequests')

    }
       
      navigate('/');
  
    };


    return (
      <div className='create-request-page'>
        <Header title="Новая заявка" />
        
        <p>На что вы хотите подать заявку?</p>

        <ToggleButton
        leftLabel="Мерч"
        rightLabel="Коины"
        onSelect={handleSelect}
        />
        <form onSubmit={handleSubmit}>

        {selection === 'right' && 
        <ReqForCoinsForm  
          selectedAchievement={selectedAchievement}
          onAchievementChange={handleAchievementChange} />}

        {selection === 'left' &&
        <ReqForMerchForm 
        onMerchSelect={handleMerchChange}
        selectedMerch={selectedMerch}
        />}


        <p>Прикрепите доказательства наличия достижения</p>
        {(selection) && (
        <textarea
          placeholder="Введите комментарий..."
          value={comment}
          onChange={handleCommentChange}
        />
      )}


      {(selection) && (
        <button
         type="submit"
          className="submit-button"
          // disabled={!selectedAchievement && selection === 'right' || !comment}
          >Отправить
        </button> 
      )}

        </form>
      </div>
    );
  };
  
  export default CreateRequestPage;