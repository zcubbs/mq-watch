import React from 'react';

interface DateTimeInputGroupProps {
  label: string;
  value: string;
  onChange: (value: string) => void;
}

const DateTimeInputGroup: React.FC<DateTimeInputGroupProps> = ({ label, value, onChange }) => {
  return (
    <div className="date-time-input-group">
      <label>{label}</label>
      <input type="datetime-local" value={value} onChange={(e) => onChange(e.target.value)} />
    </div>
  );
};

export default DateTimeInputGroup;
