import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {TopTenants} from "@/pages/mq-watch/components/top-tenants.tsx";
import {useState} from "react";
import {CalendarDateRangePicker} from "@/pages/mq-watch/components/date-range-picker.tsx";
import {subDays} from "date-fns";
import {DateRange} from "react-day-picker";
import MessagePerDayChart from "@/pages/mq-watch/components/message-per-day-chart.tsx";
import MessagePerTenantChart from "@/pages/mq-watch/components/message-per-tenant-chart.tsx";
import TotalMessageStat from "@/pages/mq-watch/components/total-message-stat.tsx";
import TopTenantStat from "@/pages/mq-watch/components/top-tenant-stat.tsx";

interface DateRangeState {
  from: Date;
  to: Date;
}

function MQWatchPage() {
  const [dateRange, setDateRange] = useState<DateRangeState>({
    from: subDays(new Date(), 7),
    to: new Date(),
  });

  // Handler that updates the date range state
  const handleDateChange = (newDateRange: DateRange) => {
    if (newDateRange.from && newDateRange.to) {
      setDateRange({
        from: newDateRange.from,
        to: newDateRange.to,
      });
    }
  };

  return (
    <div className="flex-col md:flex">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">MQ Watch</h2>
          <div className="flex items-center space-x-2">
            <CalendarDateRangePicker
              onDateChange={handleDateChange} className={undefined}
              />
          </div>
        </div>
      </div>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6 pl-8 pr-8 pb-6">
        <TotalMessageStat startDate={dateRange.from} endDate={dateRange.to} />
        <TopTenantStat startDate={dateRange.from} endDate={dateRange.to} />
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6 pl-8 pr-8 pb-6">
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Total Messages Overview</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <MessagePerDayChart
              startDate={dateRange.from}
              endDate={dateRange.to}
            />
          </CardContent>
        </Card>
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Messages Count Per Tenant</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <MessagePerTenantChart startDate={dateRange.from} endDate={dateRange.to} />
          </CardContent>
        </Card>
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Top Tenants Overview</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <TopTenants
              startDate={dateRange.from}
              endDate={dateRange.to}
            />
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

export default MQWatchPage;
