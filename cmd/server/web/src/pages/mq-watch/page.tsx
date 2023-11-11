import {DatePickerForm} from "@/pages/mq-watch/components/query-form.tsx";
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import MessagesLineChart from "@/pages/mq-watch/components/chart-v2.tsx";
import {MessagePerTenantList} from "@/pages/mq-watch/components/message-per-tenant-list.tsx";
import {TopTenants} from "@/pages/mq-watch/components/top-tenants.tsx";

function MQWatchPage() {
  return (
    <div className="flex-col md:flex">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">MQ Watch</h2>
          <div className="flex items-center space-x-2">
            <DatePickerForm />
          </div>
        </div>
      </div>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6 pl-8 pr-8 pb-6">
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
            <div className="text-2xl font-bold">350 548</div>
            <p className="text-xs text-muted-foreground"></p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Top Tenant
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
              <div className="text-2xl font-bold">Tenant1</div>
              <p className="text-xs text-muted-foreground">
                with 80 000 messages
              </p>
            </CardContent>
        </Card>
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6 pl-8 pr-8 pb-6">
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Total Messages Overview</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <MessagePerTenantList />
          </CardContent>
        </Card>
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Messages Count Per Tenant</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <MessagesLineChart />
          </CardContent>
        </Card>
        <Card className="col-span-2">
          <CardHeader>
            <CardTitle>Top Tenants Overview</CardTitle>
          </CardHeader>
          <CardContent className="pl-2">
            <TopTenants />
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

export default MQWatchPage;
