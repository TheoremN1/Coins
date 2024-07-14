import React, { useState } from 'react';
import './ToggleButton.css';

const ToggleButton = ({ leftLabel, rightLabel, onSelect }) => {
  const [selected, setSelected] = useState(null);

  const handleClick = (selection) => {
    setSelected(selection);
    onSelect(selection);
  };

  return (
    <div className="toggle-button">
      <button
        className={`toggle-button-left ${selected === 'left' ? 'selected' : ''}`}
        onClick={() => handleClick('left')}
      >
        {leftLabel}
      </button>
      <button
        className={`toggle-button-right ${selected === 'right' ? 'selected' : ''}`}
        onClick={() => handleClick('right')}
      >
        {rightLabel}
      </button>
    </div>
  );
};

export default ToggleButton;