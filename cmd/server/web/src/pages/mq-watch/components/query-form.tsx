import React, {useState} from 'react';
import {fetchData} from "@/pages/mq-watch/api.ts";

const QueryForm: React.FC = () => {
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [data, setData] = useState(null);

  const handleQuery = async () => {
    try {
      const result = await fetchData(startDate, endDate);
      setData(result);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  return (
    <div className="query-form-container">
    </div>
  );
};

export default QueryForm;
