import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './ReqForMerchForm.css';
const ReqForMerchForm = ({ onMerchSelect, selectedMerch }) => {

    const navigate = useNavigate();
    const [merchItems, setMerchItems] = useState([]);

    useEffect(() => {
      const fetchMerchItems = async () => {
        try {
          const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/merchItems');
          const data = await response.json();
          setMerchItems(data);
        } catch (error) {
          console.error('Error fetching merch items:', error);

          const data = {}



        }
      };
  
      fetchMerchItems();
    }, []);
  
    return (
      <div className="form-merch">
        <h3>Выберите товар</h3>
        <div className="merch-grid">
          {merchItems.map((item) => (
            <div
              key={item.id}
              className={`merch-item ${selectedMerch === item.id ? 'selected' : ''}`}
              onClick={() => onMerchSelect(item.id)}
            >
              <img src={'src/assets/Merch.png'} alt={item.name} className="merch-image" />
              <div className="merch-info">
                <p>{item.name}</p>
                <p>{item.price} <span className="coin-icon">💰</span></p>
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };
  
  export default ReqForMerchForm;