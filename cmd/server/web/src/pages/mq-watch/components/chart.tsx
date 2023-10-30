// import { Card, LineChart, Title } from "@tremor/react";
// import React, {useMemo} from "react";
//
// export type ChartData = {
//   tenantMessages: ChartDataMessage[];
// }
//
// export type ChartDataMessage = {
//   tenant: string;
//   totalMessages: number;
//   datetime: string;
// }
//
// const customValueFormatter = (value: number) => {
//   return value.toString();
// }
//
// const MessagesChart: React.FC<ChartData> = ({tenantMessages}) => {
//
//   let cData: ChartDataMessage[] = useMemo(() => {
//     return tenantMessages;
//   }, [tenantMessages]);
//
//   let categories: string[] = useMemo(() => {
//     return tenantMessages?.map((message) => message.tenant);
//   }, [tenantMessages]);
//
//   return (
//     <Card>
//       <Title>Messages per Tenant</Title>
//       <LineChart
//         className="mt-6"
//         data={cData}
//         index="datetime"
//         categories={["tenant"]}
//         valueFormatter={customValueFormatter}
//         yAxisWidth={40}
//         connectNulls={true}
//         allowDecimals={false}
//       />
//     </Card>
//   )
// };
//
// export default MessagesChart;
