import {useQuery} from "@tanstack/react-query";

interface MessagesData {
  total_messages: number;
  daily_data: Record<string, Record<string, number>>;
}

export const fetchData = async (startDateTime: string, endDateTime: string): Promise<MessagesData> => {
  const queryURL = `http://localhost:8000/api/messages?start_datetime=${startDateTime}&end_datetime=${endDateTime}`;

  return fetch(queryURL)
    .then((response) => {
      if (!response.ok) {
        throw new Error(response.statusText);
      }
      return response.json();
    });
};

export const useFetchData = (startDateTime: string, endDateTime: string) => {
  return useQuery({
    queryKey: ['domains'],
    refetchIntervalInBackground: true,
    refetchInterval: 10000,
    queryFn: () => fetchData(startDateTime, endDateTime)
  });
};
