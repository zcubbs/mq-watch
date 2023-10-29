import React from 'react';

interface MessageStatProps {
  tenant: string;
  totalMessages: number;
}

const MessageStat: React.FC<MessageStatProps> = ({ tenant, totalMessages }) => {
  return (
    <div className="message-stat">
      <span className="tenant">{tenant}</span>
      <span className="total-messages">{totalMessages}</span>
    </div>
  );
};

interface MessageStatsListProps {
  data: MessageStatProps[];
}

const MessageStatsList: React.FC<MessageStatsListProps> = ({ data }) => {
  return (
    <div className="message-stats-list">
      {data.map((stat, index) => (
        <MessageStat key={index} {...stat} />
      ))}
    </div>
  );
};

export default MessageStatsList;
