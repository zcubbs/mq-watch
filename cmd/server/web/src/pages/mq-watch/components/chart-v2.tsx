import React, {RefObject, useEffect, useRef, useState} from 'react';
import {Chart, ChartData, ChartDataset, ChartOptions} from 'chart.js';
import "chart.js/auto";
// @ts-ignore
import 'chartjs-adapter-moment';
import {Line} from 'react-chartjs-2'
import {fetchData} from "@/pages/mq-watch/api.ts";

interface LineChartData extends ChartData {
  labels: string[];
  datasets: ChartDataset<"line", number[]>[];
}

const getRandomColor = (): string => {
  const letters = '0123456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
};

// Function to format date to RFC3339
function formatDateToRFC3339(date: Date) {
  function pad(value: number) {
    return value < 10 ? `0${value}` : value;
  }

  return `${date.getUTCFullYear()}-${pad(date.getUTCMonth() + 1)}-${pad(date.getUTCDate())}T${pad(date.getUTCHours())}:${pad(date.getUTCMinutes())}:${pad(date.getUTCSeconds())}Z`;
}

// Define the options for the Line chart
const options: ChartOptions<'line'> = {
  scales: {
    x: {
      type: 'time',
      time: {
        parser: 'YYYY-MM-DD',
        unit: 'day',
        displayFormats: {
          day: 'DD MMM'
        },
        tooltipFormat: 'DD MMM'
      },
      ticks: {
        autoSkip: true,
        maxTicksLimit: 10  // adjust to your preference
      },
      title: {
        display: true,
        text: 'Date',
      },
    },
    y: {
      title: {
        display: true,
        text: 'Messages',
      },
    },
  },
};


const MessagesLineChart: React.FC = () => {
  const [chartData, setChartData] = useState<LineChartData | null>(null);
  const chartRef = useRef<RefObject<Chart> | null>(null);

  useEffect(() => {
    const endDate = new Date();

    // Calculate start date as 7 days ago at 00:00
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 60);
    startDate.setHours(0, 0, 0, 0);

    // Fetch data from API
    fetchData(formatDateToRFC3339(startDate), formatDateToRFC3339(endDate))
      .then((data) => {
        const dailyData = data.daily_data;

        // Extracting unique tenants
        let tenants: string[] = [];
        Object.values(dailyData).forEach((day: any) => {
          tenants = [...tenants, ...Object.keys(day)];
        });
        tenants = [...new Set(tenants)]; // Removing duplicates

        // Extracting dates
        const dates = Object.keys(dailyData);

        // Creating datasets
        const datasets = tenants.map((tenant) => {
          return {
            label: tenant,
            data: dates.map((date) => dailyData[date][tenant] || 0), // Using 0 if no data for tenant on a date
            fill: false,
            borderColor: getRandomColor(),
          };
        });

        setChartData({
          labels: dates,
          datasets,
        });
      });
  }, []);

  if (!chartData?.labels || chartData.labels.length === 0 || !chartData.datasets) {
    return <div>Loading...</div>;
  }

  const toggleLines = () => {
    if (!chartRef.current) return;

    // Toggle lines directly on the chart instance
    (chartRef.current as any)?.data.datasets.forEach((dataset: ChartDataset<"line", number[]>) => {
      dataset.hidden = !dataset.hidden;
    });
    (chartRef.current as any)?.update(); // Update the chart to reflect changes
  };

  return (
    <div>
      {/*<button onClick={toggleLines}>*/}
      {/*  Toggle Lines*/}
      {/*</button>*/}
      <Line
        height="220"
        width="auto"
        ref={chartRef as any}
        data={chartData}
        options={options}
      />
    </div>
  );
};

export default MessagesLineChart;
