import React, { useState } from 'react';

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

const QueryButton: React.FC<{ onClick: () => void }> = ({ onClick }) => {
  return <button onClick={onClick}>Query</button>;
};

const QueryForm: React.FC = () => {
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');

  const handleQuery = () => {
    console.log('Fetching data for:', startDate, endDate);
    // Fetch data here
  };

  return (
    <div className="query-form">
      <DateTimeInputGroup label="Start Date" value={startDate} onChange={setStartDate} />
      <DateTimeInputGroup label="End Date" value={endDate} onChange={setEndDate} />
      <QueryButton onClick={handleQuery} />
    </div>
  );
};

export default QueryForm;
