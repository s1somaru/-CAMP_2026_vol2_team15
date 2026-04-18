import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip
} from "recharts";

function MentalChart({ logs }) {
  return (
    <LineChart width={400} height={200} data={logs}>
      <XAxis dataKey="date" />
      <YAxis domain={[1, 5]} />
      <Tooltip />
      <Line type="monotone" dataKey="mood" />
    </LineChart>
  );
}

export default MentalChart;