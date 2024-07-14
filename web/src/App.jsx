
import './App.css';
import { useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import MainPage from './components/MainPage/MainPage.jsx';
import CreateRequestPage from './components/CreateRequestPage/CreateRequestPage.jsx';
function App() {

useEffect (()=> {

}, [])




return (
  
<div className="App">

    <Routes>
      <Route path='/' element={<MainPage />} />
      <Route path="/create" element={<CreateRequestPage />} />
       {/* <Route index element={<div>No page is selected.</div> } /> */}
     </Routes>

</div>
);


}

export default App;