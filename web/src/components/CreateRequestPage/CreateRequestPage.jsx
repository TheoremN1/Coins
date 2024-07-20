import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '../Header/Header';
import Button from '../Button/Button';
import ToggleButton from '../ToggleButton/ToggleButton';
import ReqForCoinsForm from './ReqForCoinsForm/ReqForCoinsForm'
import ReqForMerchForm from './ReqForMerchForm/ReqForMerchForm'
import './CreateRequestPage.css';


const CreateRequestPage = () => {

    const navigate = useNavigate();
    const [selection, setSelection] = useState(null);
    // const [isReadyToSubmit, setIsReadyToSubmit] = useState(false);
    const [selectedAchievement, setSelectedAchievement] = useState('');
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


    const handleSubmit = async () => {
      // if (!isReadyToSubmit) return;
      
  
      const CoinReq = {
        id_user: 0, //TODO: задать id юзера после реализации сервиса регистрации
        id_achievement: selectedAchievement,
        comment_hr: '1',
        comment_s: comment,
        id_status: 1
      };
      
      alert (JSON.stringify(CoinReq))
      // try {
      //   const response = await fetch('https://simplbot.onrender.com/api/submitCoinRequest', {
      //     method: 'POST',
      //     headers: {
      //       'Content-Type': 'application/json',
      //     },
      //     body: JSON.stringify(CoinReq),
      //   });
  
      //   if (response.ok) {
      //     console.log('Application submitted successfully');
  
      //   } else {
      //     console.error('Failed to submit application');
      //   }
      // } catch (error) {
      //   console.error('Error submitting application:', error);
      // }
  
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

        {selection === 'left' && <ReqForMerchForm />}


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