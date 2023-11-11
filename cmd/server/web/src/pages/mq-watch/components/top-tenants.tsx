import {useState, useEffect, FC} from 'react';
import { Avatar, AvatarFallback } from "@/components/ui/avatar.tsx";
import {fetchTopTenantsForDateRange} from "@/pages/mq-watch/api.ts";

interface Tenant {
  tenant: string;
  messageCount: number;
}

interface TopTenantsProps {
  startDate?: Date;
  endDate?: Date;
}

export const TopTenants: FC<TopTenantsProps> = ({ startDate, endDate }) => {
  const [topTenants, setTopTenants] = useState<Tenant[]>([]);

  useEffect(() => {
    const defaultStartDate = startDate ?? new Date();
    const defaultEndDate = endDate ?? new Date();

    // Use the ISO strings directly for the API call
    fetchTopTenantsForDateRange(defaultStartDate.toISOString(), defaultEndDate.toISOString())
      .then((data: Tenant[]) => {
        setTopTenants(data);
      })
      .catch((error) => {
        console.error("Error fetching top tenants:", error);
      });
  }, [startDate, endDate]); // Include dependencies

  return (
    <div className="space-y-8">
      {topTenants.slice(0, 6).map((tenant, index) => (
        <div key={tenant.tenant} className="flex items-center">
          <Avatar className="h-9 w-9">
            <AvatarFallback>{index + 1}</AvatarFallback>
          </Avatar>
          <div className="ml-4 space-y-1">
            <p className="text-sm font-medium leading-none">{tenant.tenant}</p>
          </div>
          <div className="ml-auto font-medium">{tenant.messageCount.toLocaleString()}</div>
        </div>
      ))}
    </div>
  );
};
