import React, { useEffect, useState } from 'react';
import './ReqForCoinsForm.css';


//TODO: переписать подстановку данных о заявках под текущий сервер


const ReqForCoinsForm = ({ selectedAchievement, onAchievementChange}) => {
  const [achievements, setAchievements] = useState([]);

  // const [isReadyToSubmit, setIsReadyToSubmit] = useState(false);

  

  useEffect(() => {
    
    // Fetch achievements from the backend
    const fetchAchievements = async () => {
      try {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/achievements');
        const data = await response.json();
        setAchievements(data);
      } catch (error) {
        console.error('Error fetching achievements:', error);
      }
    };

    fetchAchievements();
  }, []);



  // useEffect(() => {
  //   //Действие после прогрузки страницы и заполнения формы
  //   //TODO: определить, где будет располагаться кнопка отправки, на самой форме или на странице создания заявки

  // }, [isReadyToSubmit]);


  return (
    <div>
    <p>Выберите достижение из списка</p>


    <div className='achievements-form'>
      {achievements.map((achievement) => (
        
        <div className='achievement'> 
           
            
          <input
            type="radio"
            id={`achievement-button ${achievement.id_achievement}`}
            value={achievement.id_achievement}
            checked={selectedAchievement === achievement.id_achievement}
            onChange={() => onAchievementChange(achievement.id_achievement)}
          />
         <label 
          key={achievement.id_achievement}
          // className={`achievement-option ${selectedAchievement === achievement.id_achievement ? 'selected' : ''}`} 
          for={`achievement-button ${achievement.id_achievement}`}>
          {achievement.name}
        </label>
        </div>
      ))}
         </div>   
    
    </div>
  );
};

export default ReqForCoinsForm;