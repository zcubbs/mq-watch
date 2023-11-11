import { useState, useEffect, FC } from 'react';
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis, Tooltip } from 'recharts';
import { fetchTotalMessagesPerDay } from '../api';

interface MessagePerTenantListProps {
  startDate?: Date;
  endDate?: Date;
}

interface ChartData {
  name: string;
  total: number;
}

export const MessagePerDayChart: FC<MessagePerTenantListProps> = ({ startDate, endDate }) => {
  const [data, setData] = useState<ChartData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const defaultStartDate = startDate ?? new Date();
    const defaultEndDate = endDate ?? new Date();

    const formattedStartDate = defaultStartDate.toISOString();
    const formattedEndDate = defaultEndDate.toISOString();

    fetchTotalMessagesPerDay(formattedStartDate, formattedEndDate)
      .then((fetchedData: ChartData[]) => {
        setData(fetchedData);
        setLoading(false);
      })
      .catch((err: Error) => {
        setError(err.message);
        setLoading(false);
      });
  }, [startDate, endDate]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <ResponsiveContainer width="100%" height={380}>
      <BarChart data={data}>
        <XAxis
          dataKey="name"
          stroke="#367588"
          fontSize={12}
          tickLine={false}
          axisLine={false}
        />
        <YAxis
          stroke="#367588"
          fontSize={12}
          tickLine={false}
          axisLine={false}
          tickFormatter={(value: number) => `${value}`}
        />
        <Tooltip
          cursor={{ fill: 'transparent' }}
          formatter={(value: number) => [`Total: ${value}`]}
          labelFormatter={(name: string) => `Date: ${name}`}
          contentStyle={{
            backgroundColor: '#333',
            borderColor: '#777',
            borderRadius: '4px',
            color: '#fff'
          }}
          itemStyle={{
            color: '#fff'
          }}
        />
        <Bar dataKey="total" fill="#008080" radius={[4, 4, 0, 0]} />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default MessagePerDayChart;
