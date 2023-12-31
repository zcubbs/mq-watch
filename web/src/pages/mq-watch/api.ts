import {API_URL} from "@/lib/utils.ts";

const formatDateToRFC3339 = (datetime: string | number | Date) => {
  const date = new Date(datetime);
  return date.toISOString();
};

export const fetchData = async (startDate: string, endDate: string) => {
  const formattedStartDate = formatDateToRFC3339(startDate);
  const formattedEndDate = formatDateToRFC3339(endDate);

  const response = await fetch(API_URL + `/api/messages?start_datetime=${formattedStartDate}&end_datetime=${formattedEndDate}`);
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Network response was not ok");
  }
  return await response.json();
};

export const fetchTotalMessagesPerDay = async (startDate: string, endDate: string) => {
  const formattedStartDate = formatDateToRFC3339(startDate);
  const formattedEndDate = formatDateToRFC3339(endDate);

  const response = await fetch(API_URL + `/api/total-messages-per-day?start_date=${formattedStartDate}&end_date=${formattedEndDate}`);
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Network response was not ok");
  }
  return await response.json();
};

export const fetchTopTenantsForDateRange = async (startDate: string, endDate: string) => {
  const formattedStartDate = formatDateToRFC3339(startDate);
  const formattedEndDate = formatDateToRFC3339(endDate);

  const response = await fetch(API_URL + `/api/top-tenants?start_date=${formattedStartDate}&end_date=${formattedEndDate}`);
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Network response was not ok");
  }
  return await response.json();
};

export const fetchMessageStats = async (startDate: string, endDate: string) => {
  const formattedStartDate = formatDateToRFC3339(startDate);
  const formattedEndDate = formatDateToRFC3339(endDate);

  const response = await fetch(API_URL + `/api/message-stats?start_date=${formattedStartDate}&end_date=${formattedEndDate}`);
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Network response was not ok");
  }
  return await response.json();
};
