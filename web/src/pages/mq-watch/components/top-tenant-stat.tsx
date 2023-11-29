import { useState, useEffect, FC } from 'react';
import { Card, Metric, Text } from "@tremor/react";
import { fetchTopTenantsForDateRange } from "@/pages/mq-watch/api.ts";

interface TopTenant {
  tenant: string;
  messageCount: number;
}

interface TopTenantProps {
  startDate?: Date;
  endDate?: Date;
}

export const TopTenantStat: FC<TopTenantProps> = ({ startDate, endDate }) => {
  const [topTenant, setTopTenant] = useState<TopTenant | null>(null);

  useEffect(() => {
    const defaultStartDate = startDate ?? new Date();
    const defaultEndDate = endDate ?? new Date();

    fetchTopTenantsForDateRange(defaultStartDate.toISOString(), defaultEndDate.toISOString())
      .then((data: TopTenant[]) => {
        // Assuming the top tenant is the first in the array
        if (data.length > 0) {
          setTopTenant(data[0]);
        }
      })
      .catch((error) => {
        console.error("Error fetching top tenant:", error);
      });
  }, [startDate, endDate]); // Include dependencies

  return (
    <Card className="max-w-xs mx-auto">
      <Text>Top Tenant</Text>
      {topTenant ? (
        <>
          <Text>{topTenant.tenant}</Text>
          <Metric>{topTenant.messageCount.toLocaleString()}</Metric>
        </>
      ) : (
        <Text>Loading...</Text>
      )}
    </Card>
  );
};

export default TopTenantStat;
