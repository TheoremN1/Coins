import React, { useState, useEffect , useCallback} from 'react';

import  './Header.css';

const Header = ({ title }) => {

 const [balance, setBalance] = useState(null);  
  useEffect(() => { //при обновлении страницы будет вычисляться баланс


     const fetchBalance = async () => {
       try {
         const response = await fetch(import.meta.env.VITE_BACKEND_URL +'/api/balance' + '?id=1');
         if (!response.ok) {
           throw new Error('Network response was not ok');
         }
         const data = await response.json();   
           
         setBalance(data);
       } catch (error) {
         console.error('Balance fetch error: ', error);
       }
     };
     fetchBalance();



    

  }, []);




    return (
      <header className="Header">
        {balance !== null ? <span>{balance}</span> : <span>Загрузка...подожди чутка</span>}
        <h1>{title}</h1>
        
      </header>
    );
  };
  export default Header;