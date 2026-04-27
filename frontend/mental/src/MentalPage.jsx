import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";
import { useState, useEffect } from "react";
import MentalChart from "./MentalChart";

function MentalPage() {
  const [logs, setLogs] = useState(() => {
    const saved = localStorage.getItem("logs");
    return saved ? JSON.parse(saved) : [];
  });

  const [mood, setMood] = useState(3);
  const [event, setEvent] = useState("");
  const [date, setDate] = useState(
    new Date().toISOString().slice(0, 10)
  );

  const [editIndex, setEditIndex] = useState(null);
  const [selectedLog, setSelectedLog] = useState(null); // モーダル用

  useEffect(() => {
    localStorage.setItem("logs", JSON.stringify(logs));
  }, [logs]);

  // 追加・編集
  function addLog() {
    if (!event.trim()) return;

    const newLog = {
      date,
      mood: Number(mood),
      event
    };

    let updated;

    if (editIndex !== null) {
      updated = logs.map((log, i) =>
        i === editIndex ? newLog : log
      );
      setEditIndex(null);
    } else {
      updated = [...logs, newLog];
    }

    setLogs(updated);
    setEvent("");
  }

  // 削除
  function deleteLog(index) {
    const updated = logs.filter((_, i) => i !== index);
    setLogs(updated);
    setSelectedLog(null);
  }

  // カレンダークリック
  function handleDateClick(value) {
    const d = value.toISOString().slice(0, 10);
    setDate(d);

    const found = logs.find((l) => l.date === d);
    if (found) {
      setMood(found.mood);
      setEvent(found.event);
    } else {
      setMood(3);
      setEvent("");
    }
  }

  // 感情アイコン
  function moodIcon(mood) {
    if (mood <= 2) return "😭";
    if (mood === 3) return "😐";
    return "😄";
  }

  return (
    <div className="container">
      <h1>メンタル記録</h1>

      {/* 📅 カレンダー */}
      <Calendar
        onClickDay={handleDateClick}
        value={new Date(date)}
        tileContent={({ date, view }) => {
          if (view === "month") {
            const d = date.toISOString().slice(0, 10);
            const log = logs.find((l) => l.date === d);

            if (log) {
              const color =
                log.mood <= 2
                  ? "#3b82f6"
                  : log.mood === 3
                  ? "#9ca3af"
                  : "#f59e0b";

              return (
                <div
                  style={{
                    width: 6,
                    height: 6,
                    borderRadius: "50%",
                    background: color,
                    margin: "2px auto"
                  }}
                />
              );
            }
          }
        }}
      />

      {/* 📊 グラフ */}
      <div className="chart">
        <MentalChart
          logs={[...logs].sort(
            (a, b) => new Date(a.date) - new Date(b.date)
          )}
        />
      </div>

      {/* 📝 入力 */}
      <input
        type="date"
        value={date}
        onChange={(e) => setDate(e.target.value)}
      />

      <select value={mood} onChange={(e) => setMood(e.target.value)}>
        <option value="1">😭 1</option>
        <option value="2">😢 2</option>
        <option value="3">😐 3</option>
        <option value="4">🙂 4</option>
        <option value="5">😄 5</option>
      </select>

      <input
        placeholder="出来事"
        value={event}
        onChange={(e) => setEvent(e.target.value)}
      />

      <button onClick={addLog}>
        {editIndex !== null ? "更新" : "記録"}
      </button>

      {/* 📜 履歴 */}
      <ul>
        {[...logs].reverse().map((log, i) => {
          const realIndex = logs.length - 1 - i;

          return (
            <li
              key={i}
              className="card"
              onClick={() => setSelectedLog(log)}
              style={{ cursor: "pointer" }}
            >
              <p>📅 {log.date}</p>
              <p>
                {moodIcon(log.mood)} 気分: {log.mood}
              </p>
              <p>📝 {log.event}</p>

              <button
                onClick={(e) => {
                  e.stopPropagation();
                  setEditIndex(realIndex);
                  setMood(log.mood);
                  setEvent(log.event);
                  setDate(log.date);
                }}
              >
                編集
              </button>

              <button
                onClick={(e) => {
                  e.stopPropagation();
                  deleteLog(realIndex);
                }}
              >
                削除
              </button>
            </li>
          );
        })}
      </ul>

      {/* ➕ FAB */}
      <button
        className="fab"
        onClick={() =>
          window.scrollTo({ top: 0, behavior: "smooth" })
        }
      >
        ＋
      </button>

      {/* 🔥 モーダル */}
      {selectedLog && (
        <div className="modal" onClick={() => setSelectedLog(null)}>
          <div
            className="modal-content"
            onClick={(e) => e.stopPropagation()}
          >
            <p>📅 {selectedLog.date}</p>
            <p>
              {moodIcon(selectedLog.mood)} 気分: {selectedLog.mood}
            </p>
            <p>📝 {selectedLog.event}</p>

            <button onClick={() => setSelectedLog(null)}>
              閉じる
            </button>
          </div>
        </div>
      )}
    </div>
  );
}

export default MentalPage;