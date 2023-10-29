import QueryForm from "@/pages/mq-watch/components/query-form.tsx";
import Graph from "@/pages/mq-watch/components/graph.tsx";
import {useFetchData} from "@/pages/mq-watch/api.ts";

function QueryPage() {
  const { data, error, isLoading } = useFetchData();

  if (isLoading) {
    return <span>Loading...</span>;
  }

  if (error) {
    return <span>Error fetching data.</span>;
  }

  return (
    <>
      <QueryForm />
      <Graph data={data} />
    </>
  );
}

export default QueryPage;
