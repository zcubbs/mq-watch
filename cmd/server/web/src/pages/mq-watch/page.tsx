import MessagesLineChart from "@/pages/mq-watch/components/chart-v2.tsx";
import {DatePickerWithRange} from "@/components/date-range-picker.tsx";
import {DatePickerForm} from "@/pages/mq-watch/components/query-form.tsx";

function MQWatchPage() {
  return (
    <>
      {/*<QueryForm />*/}
      <DatePickerForm />
      <MessagesLineChart />
      {/*<MessagesChart tenantMessages={chartDataMessages} />*/}
      {/*<Graph data={fetchData} />*/}
      {/*<MessageStatsList data={statData}/>*/}
    </>
  );
}

export default MQWatchPage;
