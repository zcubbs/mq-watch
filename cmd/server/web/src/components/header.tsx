import React from 'react';
import {useTheme} from "@/context/theme-context.tsx";

const Header: React.FC = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <div className={`header ${theme}`}>
      <h1>MQ Watch</h1>
      <button onClick={toggleTheme}>
        Switch to {theme === 'light' ? 'Dark' : 'Light'} Mode
      </button>
    </div>
  );
};

export default Header;
