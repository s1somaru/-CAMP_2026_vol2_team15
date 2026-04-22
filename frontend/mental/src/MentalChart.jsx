import {
  ResponsiveContainer,
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid
} from "recharts";

function MentalChart({ logs }) {
  return (
    <div style={{ width: "100%", height: 250 }}>
      <ResponsiveContainer width="100%" height="100%">
        <LineChart data={logs}>
          <CartesianGrid strokeDasharray="3 3" />

          <XAxis 
            dataKey="date"
            tick={{ fontSize: 12 }}
          />

          <YAxis 
            domain={[1, 5]}
            tick={{ fontSize: 12 }}
          />

          <Tooltip />

          <Line
            type="monotone"
            dataKey="mood"
            stroke="#4a3f35"  //ブラウン
            strokeWidth={3}    // ← 少し太くして見やすく
            dot={{ r: 4 }}     // ← 点を大きく（スマホ用）
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}

export default MentalChart;