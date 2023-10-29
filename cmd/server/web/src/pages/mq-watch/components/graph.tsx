import React, {useEffect, useRef, useState} from 'react';
import * as d3 from 'd3';
import {useFetchData} from "@/pages/mq-watch/api.ts";

type DataPoint = {
  date: string;
  value: number;
};

type GraphData = {
  [tenantId: string]: DataPoint[];
};

interface GraphProps {
  data: GraphData;
}

const Graph: React.FC<GraphProps> = () => {
  const svgRef = useRef(null);
  const [startDate, setStartDate] = useState<string>(() => {
    const lastWeek = new Date();
    lastWeek.setDate(lastWeek.getDate() - 7);
    return lastWeek.toISOString().substr(0, 10);
  });

  const [endDate, setEndDate] = useState<string>(() => new Date().toISOString().substr(0, 10));

  const startDateTime = `${startDate}T00:00:00Z`;
  const endDateTime = `${endDate}T23:59:59Z`;

  const { data, isLoading, isError } = useFetchData(startDateTime, endDateTime);

  useEffect(() => {
    const svg = d3.select(svgRef.current);
    const margin = { top: 20, right: 20, bottom: 30, left: 50 };
    const width = 960 - margin.left - margin.right;
    const height = 500 - margin.top - margin.bottom;

    const x = d3.scaleTime().range([0, width]);
    const y = d3.scaleLinear().range([height, 0]);

    const color = d3.scaleOrdinal(d3.schemeCategory10);

    const line = d3.line<DataPoint>()
      .x(d => x(new Date(d.date)))
      .y(d => y(d.value));

    // Clear existing content
    svg.selectAll('*').remove();

    svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);

    const allDataPoints: DataPoint[] = [];
    Object.values(data).forEach(tenantData => {
      allDataPoints.push(...tenantData);
    });

    allDataPoints.forEach(d => {
      d.date = new Date(d.date);
    });

    x.domain(d3.extent(allDataPoints, d => d.date) as [Date, Date]);
    y.domain([0, d3.max(allDataPoints, d => d.value) as number]);

    svg.append("g")
      .attr("transform", `translate(${margin.left},${height + margin.top})`)
      .call(d3.axisBottom(x));

    svg.append("g")
      .attr("transform", `translate(${margin.left},${margin.top})`)
      .call(d3.axisLeft(y));

    Object.entries(data).forEach(([tenantId, tenantData], index) => {
      svg.append("path")
        .datum(tenantData)
        .attr("fill", "none")
        .attr("stroke", color(index.toString()))
        .attr("stroke-width", 1.5)
        .attr("d", line);
    });

    // Define the tooltip div outside the SVG
    const div = d3.select("body").append("div")
      .attr("class", "tooltip")
      .style("opacity", 0);

// On mouseover for each point, show the tooltip
    const mouseover = function(d: DataPoint) {
      div.transition()
        .duration(200)
        .style("opacity", .9);
      div.html(d.value.toString())
        .style("left", (d3.event.pageX) + "px")
        .style("top", (d3.event.pageY - 28) + "px");
    }

// On mouseout, hide the tooltip
    const mouseout = function(d: DataPoint) {
      div.transition()
        .duration(500)
        .style("opacity", 0);
    }

    Object.entries(data).forEach(([tenantId, tenantData], index) => {
      const tenantLine = svg.append("path")
        .datum(tenantData)
      // ... other attributes

      tenantLine.on("mouseover", mouseover)
        .on("mouseout", mouseout);
    });

  }, [data]);

  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error loading data</div>;

  return <svg ref={svgRef} width={960} height={500}></svg>;
};

export default Graph;
