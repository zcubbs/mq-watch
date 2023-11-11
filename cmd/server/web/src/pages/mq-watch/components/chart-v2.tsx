import React, { FC, useEffect, useState } from 'react';
import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
  TooltipProps,
} from 'recharts';
import { fetchData } from "@/pages/mq-watch/api"; // Assuming this is the correct path
import { formatDateToRFC3339, getRandomPastelColor } from "@/lib/utils"; // Assuming this is the correct path

interface TenantData {
  [tenant: string]: number;
}

interface DailyData {
  [date: string]: TenantData;
}

interface MessagesLineChartProps {
  startDate: Date;
  endDate: Date;
}

interface TenantCounts {
  [tenant: string]: number;
}

interface RechartsLineChartData {
  date: string;
  counts: TenantCounts;
}

interface CustomTooltipProps extends TooltipProps<number, string> {
  payload?: {
    color: string;
    name: string;
    value: number;
  }[];
}

const CustomTooltip: React.FC<CustomTooltipProps> = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
    return (
      <div className="custom-tooltip" style={{
        backgroundColor: '#222',
        border: '1px solid #555',
        padding: '5px',
        borderRadius: '5px',
        color: 'white'
      }}>
        <p className="label">{`Date: ${label}`}</p>
        {payload.map((entry) => (
          <p key={entry.name} style={{ color: entry.color }}>
            {`${entry.name} : ${entry.value}`}
          </p>
        ))}
      </div>
    );
  }

  return null;
};

const MessagesLineChart: FC<MessagesLineChartProps> = ({ startDate, endDate }) => {
  const [chartData, setChartData] = useState<RechartsLineChartData[]>([]);

  useEffect(() => {
    fetchData(formatDateToRFC3339(startDate), formatDateToRFC3339(endDate))
      .then((response: { daily_data: DailyData }) => {
        const newData: RechartsLineChartData[] = Object.entries(response.daily_data).map(([date, tenantsData]) => {
          return { date, counts: tenantsData };
        });

        setChartData(newData);
      });
  }, [startDate, endDate]);

  if (!chartData.length) {
    return <div>Loading...</div>;
  }

  return (
    <ResponsiveContainer width="100%" height={380}>
      <LineChart data={chartData}>
        <CartesianGrid strokeDasharray="3 3" stroke="#444" />
        <XAxis dataKey="date" stroke="#ccc" />
        <YAxis stroke="#ccc" />
        <Tooltip content={<CustomTooltip />} />
        {Object.keys(chartData[0].counts).map(key => (
          <Line type="monotone" dataKey={`counts.${key}`} stroke={getRandomPastelColor()} key={key} dot={false} />
        ))}
      </LineChart>
    </ResponsiveContainer>
  );
};

export default MessagesLineChart;
