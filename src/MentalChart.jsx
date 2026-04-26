import {
  ResponsiveContainer,
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  ReferenceLine
} from "recharts";

function MentalChart({ logs }) {
  // 🔥 平均値を計算
  const average =
    logs.length > 0
      ? logs.reduce((sum, log) => sum + log.mood, 0) / logs.length
      : 0;

  return (
    <div style={{ width: "100%", height: 300 }}>
      
      {/* 🔥 平均値表示 */}
      <p style={{ textAlign: "center", marginBottom: "8px" }}>
        平均気分: {average.toFixed(2)}
      </p>

      <ResponsiveContainer width="100%" height="100%">
        <LineChart data={logs}>
          <CartesianGrid strokeDasharray="3 3" />

          <XAxis dataKey="date" tick={{ fontSize: 12 }} />

          <YAxis domain={[1, 5]} tick={{ fontSize: 12 }} />

          {/* 🔥 Tooltip */}
          <Tooltip
            content={({ active, payload }) => {
              if (active && payload && payload.length) {
                const data = payload[0].payload;
                return (
                  <div
                    style={{
                      background: "white",
                      padding: "8px",
                      border: "1px solid #ccc",
                      borderRadius: "8px"
                    }}
                  >
                    <p>{data.date}</p>
                    <p>気分: {data.mood}</p>
                    <p style={{ fontSize: "12px", color: "#4a3f35" }}>
                      {data.event}
                    </p>
                  </div>
                );
              }
              return null;
            }}
          />

          {/* 🔥 平均線 */}
          <ReferenceLine
            y={average}
            stroke="#ef4444"
            strokeDasharray="4 4"
            label="平均"
          />

          {/* 🔥 メイン線 */}
          <Line
            type="monotone"
            dataKey="mood"
            stroke="#726250"
            strokeWidth={3}
            dot={({ cx, cy, payload }) => {
              let color;

              if (payload.mood <= 2) color = "#3b82f6"; // 青
              else if (payload.mood === 3) color = "#9ca3af"; // グレー
              else color = "#f59e0b"; // オレンジ

              return (
                <circle
                  cx={cx}
                  cy={cy}
                  r={5}
                  fill={color}
                  stroke="#4a3f35"
                  strokeWidth={2}
                />
              );
            }}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}

export default MentalChart;