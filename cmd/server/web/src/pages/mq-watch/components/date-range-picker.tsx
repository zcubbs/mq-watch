import { CalendarIcon } from "@radix-ui/react-icons";
import { format, subDays } from "date-fns";
import { DateRange } from "react-day-picker";

import { cn } from "@/lib/utils";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {ReactNode, useState} from "react";

// Define a type for the component's props
interface CalendarDateRangePickerProps {
  onDateChange?: (range: DateRange) => void; // Optional prop function
  className?: string;
}

function formatDateRange(date: DateRange): ReactNode {
  if (date.from) {
    if (date.to) {
      return (
        <>
          {format(date.from, "LLL dd, y")} - {format(date.to, "LLL dd, y")}
        </>
      );
    }
    return format(date.from, "LLL dd, y");
  }
  return <span>Pick a date</span>;
}

export function CalendarDateRangePicker({
                                          onDateChange,
                                          className,
                                        }: Readonly<CalendarDateRangePickerProps>) {
  // Set the initial state to be from 7 days ago to today
  const [date, setDate] = useState<DateRange>({
    from: subDays(new Date(), 7),
    to: new Date(),
  });

  // This function is called when a new date range is selected
  const handleDateSelect = (newDateRange: DateRange | undefined) => {
    if (newDateRange) {
      setDate(newDateRange);
      onDateChange?.(newDateRange); // Call the prop function with the new date range if it exists
    }
  };

  return (
    <div className={cn("grid gap-2", className)}>
      <Popover>
        <PopoverTrigger asChild>
          <Button
            id="date"
            variant="outline"
            className={cn(
              "w-[260px] justify-start text-left font-normal",
              !date.from && "text-muted-foreground"
            )}
          >
            <CalendarIcon className="mr-2 h-4 w-4" />
            {formatDateRange(date)}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-auto p-0" align="end">
          <Calendar
            initialFocus
            mode="range"
            defaultMonth={date.from}
            selected={date}
            onSelect={handleDateSelect}
            numberOfMonths={2}
          />
        </PopoverContent>
      </Popover>
    </div>
  );
}
