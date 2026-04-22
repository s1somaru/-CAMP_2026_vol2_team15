import { useState } from "react";
import MentalChart from "./MentalChart";

function MentalPage() {
  const [logs, setLogs] = useState([]);
  const [mood, setMood] = useState(3);
  const [event, setEvent] = useState("");

  function addLog() {
    const newLog = {
      date: new Date().toISOString().slice(0, 10),
      mood: Number(mood),
      event
    };
    console.log("logs:", logs);
    console.log("newLog:", newLog);

  

    setLogs([...logs, newLog]);
    setEvent("");
  }

  return (
    <div className="container">
  <h1>メンタル記録</h1>

  <select value={mood} onChange={(e) => setMood(e.target.value)}>
    <option value="1">1（最悪）</option>
    <option value="2">2</option>
    <option value="3">3（普通）</option>
    <option value="4">4</option>
    <option value="5">5（最高）</option>
  </select>

  <input
    placeholder="出来事（面接など）"
    value={event}
    onChange={(e) => setEvent(e.target.value)}
  />

  <button onClick={addLog}>記録</button>

  <ul>
    {logs.map((log, i) => (
      <li key={i}>
        {log.date} / 気分:{log.mood} / {log.event}
      </li>
    ))}
  </ul>

  
  <div style={{ width: "80%", margin: "20px auto" }}>
    <MentalChart logs={logs} />
  </div>
</div>
  );
}

export default MentalPage;