import React from 'react';
import {Card, Metric, Text} from "@tremor/react";

interface MessageStatProps {
  tenant: string;
  totalMessages: number;
}

interface MessageStatsListProps {
  data: MessageStatProps[];
}

const MessageStatsList: React.FC<MessageStatsListProps> = ({ data }) => {
  return (
    <>
      {data?.map((stat) => (
        <Card key={stat.tenant} className="max-w-xs mx-auto">
          <Text>{stat.tenant}</Text>
          <Metric>{stat.totalMessages}</Metric>
        </Card>
      ))}
    </>
  );
};

export default MessageStatsList;
