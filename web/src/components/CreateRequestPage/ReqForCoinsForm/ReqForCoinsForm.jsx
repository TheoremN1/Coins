import React, { useEffect, useState } from 'react';
import './ReqForCoinsForm.css';

import { useNavigate } from 'react-router-dom';
const ReqForCoinsForm = () => {
  const [achievements, setAchievements] = useState([]);
  const [selectedAchievement, setSelectedAchievement] = useState('');
  const [comment, setComment] = useState('');
  const [isReadyToSubmit, setIsReadyToSubmit] = useState(false);
  const navigate = useNavigate();
  

  useEffect(() => {
    
    // Fetch achievements from the backend
    const fetchAchievements = async () => {
      try {
        const response = await fetch('https://simplbot.onrender.com/api/getAchievements');
        const data = await response.json();
        setAchievements(data);
      } catch (error) {
        console.error('Error fetching achievements:', error);
      }
    };

    fetchAchievements();
  }, []);


  useEffect(() => {
    // Update the main button state based on form completeness


    if (selectedAchievement && comment) {

    } else {

    }
  }, [selectedAchievement, comment]);




  const handleAchievementChange = (event) => {
    setSelectedAchievement(event.target.value);
  };

  const handleCommentChange = (event) => {
    setComment(event.target.value);
  };


  const handleSubmit = async (event) => {
    if (!isReadyToSubmit) return;

    const CoinReq = {
      id_user: tg.initDataUnsafe?.user.id,
      id_achievement: selectedAchievement,
      comment_hr: '1',
      comment_s: comment,
      id_status: 1
    };

    try {
      const response = await fetch('https://simplbot.onrender.com/api/submitCoinRequest', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(CoinReq),
      });

      if (response.ok) {
        console.log('Application submitted successfully');

      } else {
        console.error('Failed to submit application');
      }
    } catch (error) {
      console.error('Error submitting application:', error);
    }

    navigate('/');
  

  };

  useEffect(() => {
    //Действие после прогрузки страницы и заполнения формы
    //TODO: определить, где будет располагаться кнопка отправки, на самой форме или на странице создания заявки

  }, [isReadyToSubmit]);


  return (
    <div>
    <p>Выберите достижение из списка</p>
    <form className="achievements-form" onSubmit={handleSubmit}>    
    
      {achievements.map((achievement) => (
        <label 
        key={achievement.id_achievement}
        className={`achievement-option ${selectedAchievement === achievement.id_achievement ? 'selected' : ''}`} for="achievement">
        
          <input
            type="radio"
            name="achievement"
            value={achievement.id_achievement}
            checked={selectedAchievement === achievement.id_achievement}
            onChange={handleAchievementChange}
          />
         
          {achievement.name}
        </label>
      ))}
           
    </form>
    <p>Прикрепите доказательства наличия достижения</p>
      <textarea
        placeholder="Нажми, чтобы написать комментарий..."
        value={comment}
        onChange={handleCommentChange}
      />
     <button type="submit" className="submit-button">Отправить</button> 
    </div>
  );
};

export default ReqForCoinsForm;