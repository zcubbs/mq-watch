import {FC, useEffect, useState} from 'react';
import {fetchMessageStats} from "@/pages/mq-watch/api.ts";
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card.tsx";

interface TotalMessageStatProps {
  startDate?: Date;
  endDate?: Date;
}

const TotalMessageStat: FC<TotalMessageStatProps> = ({ startDate, endDate }) => {
  const [totalMessages, setTotalMessages] = useState<number | null>(null);

  useEffect(() => {
    // Default start and end dates to the provided props or current date
    const start = startDate ? startDate.toISOString() : new Date().toISOString();
    const end = endDate ? endDate.toISOString() : new Date().toISOString();

    // Fetch the message stats from the backend
    fetchMessageStats(start, end)
      .then((data) => {
        // Assuming the backend returns an object like { total_messages: 123 }
        setTotalMessages(data.total_messages);
      })
      .catch((error) => {
        console.error("Error fetching message stats:", error);
        // Handle the error state appropriately, possibly setting an error message in state
      });
  }, [startDate, endDate]);

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          Total Messages
        </CardTitle>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          className="h-4 w-4 text-muted-foreground"
        >
          <path d="M22 12h-4l-3 9L9 3l-3 9H2" />
        </svg>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">
          {totalMessages !== null ? totalMessages.toLocaleString() : "Loading..."}
        </div>
        <p className="text-xs text-muted-foreground"></p>
      </CardContent>
    </Card>
  );
};

export default TotalMessageStat;
