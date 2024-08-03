import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './ReqForMerchForm.css';
const ReqForMerchForm = ({ onMerchSelect, selectedMerch }) => {

    const navigate = useNavigate();
    const [merchItems, setMerchItems] = useState([]);

    useEffect(() => {
      const fetchMerchItems = async () => {

        try {
          const response = await fetch(import.meta.env.VITE_PRODUCTS_SERVICE_URL + '/api/merch');
          const data = await response.json();
          setMerchItems(data);
        } catch (error) {
          console.error('Error fetching merch items:', error);
        }
      };
  
      fetchMerchItems();
    }, []);
  
    return (
      <div className="form-merch">
        <h3>Выберите товар</h3>
        <div className="merch-grid">
          {merchItems.map((item) => (
            <div className= 'merch-item-container'>
                <div
                  key={item.Id}
                  className={`merch-item ${selectedMerch === item.Id ? 'selected' : ''}`}
                  onClick={() => onMerchSelect(item.Id)}
                >
                <img src={'src/assets/Merch.png'} alt={item.Name} className="merch-image" />   
                <p>{item.Price} <span className="coin-icon">💰</span></p>
                </div>

                <div className="merch-info">      
                <p>{item.Name}</p>                             
                </div>

            </div>
          ))}
        </div>
      </div>
    );
  };
  
  export default ReqForMerchForm;