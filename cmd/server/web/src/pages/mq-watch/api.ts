const formatDateToRFC3339 = (datetime) => {
  const date = new Date(datetime);
  return date.toISOString();
};

export const fetchData = async (startDate: string, endDate: string) => {
  try {
    const formattedStartDate = formatDateToRFC3339(startDate);
    const formattedEndDate = formatDateToRFC3339(endDate);

    const response = await fetch(`http://localhost:8000/api/messages?start_datetime=${formattedStartDate}&end_datetime=${formattedEndDate}`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Network response was not ok");
    }
    const data = await response.json();
    return data;
  } catch (error) {
    throw error;
  }
};
