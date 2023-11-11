import MessagesLineChart from "@/pages/mq-watch/components/chart-v2.tsx";
import {DatePickerForm} from "@/pages/mq-watch/components/query-form.tsx";

function MQWatchPage() {
  return (
    <>
      <DatePickerForm />
      <MessagesLineChart />
    </>
  );
}

export default MQWatchPage;
