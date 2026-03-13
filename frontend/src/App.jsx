import { useState, useEffect } from "react"

const API = import.meta.env.VITE_API_URL || "http://localhost:8080/api"

export default function App() {
  const [today, setToday] = useState(null)
  const [bhajan, setBhajan] = useState(null)
  const [loading, setLoading] = useState(true)
  const [userRashi, setUserRashi] = useState(
      localStorage.getItem("rashi") || ""
  )

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const [todayRes, bhajanRes] = await Promise.all([
        fetch(`${API}/today`),
        fetch(`${API}/bhajan${userRashi ? `?rashi=${userRashi}` : ""}`)
      ])
      const todayData = await todayRes.json()
      const bhajanData = await bhajanRes.json()
      setToday(todayData)
      setBhajan(bhajanData)
    } catch (err) {
      console.error("Error:", err)
    } finally {
      setLoading(false)
    }
  }

  if (loading) return (
      <div style={{ textAlign: "center", padding: "40px", color: "#B8541B" }}>
        <div style={{ fontSize: "48px" }}>🙏</div>
        <p>AstroMandir load ho raha hai...</p>
      </div>
  )

  return (
      <div style={{ padding: "16px" }}>

        {/* Header */}
        <div style={{ textAlign: "center", padding: "16px 0" }}>
          <h1 style={{ color: "#B8541B", fontSize: "26px" }}>
            🙏 AstroMandir
          </h1>
          <p style={{ color: "#888", fontSize: "13px" }}>
            Aapka daily spiritual companion
          </p>
        </div>

        {/* Today Card */}
        {today && <TodayCard data={today} />}

        {/* Bhajan Player */}
        {bhajan && <BhajanPlayer data={bhajan} />}

        {/* Rashi Section */}
        {userRashi
            ? <RashifalCard rashi={userRashi} onReset={() => {
              localStorage.removeItem("rashi")
              setUserRashi("")
            }} />
            : <SetRashiCard onSet={(r) => {
              localStorage.setItem("rashi", r)
              setUserRashi(r)
            }} />
        }

        {/* Kundali CTA */}
        <KundaliCTA />

      </div>
  )
}

// ─── Today Card ───────────────────────────────────
function TodayCard({ data }) {
  return (
      <div style={cardStyle}>
        <h3 style={headingStyle}>📅 Aaj Ka Din</h3>
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "8px", marginBottom: "10px" }}>
          <InfoRow icon="🌅" label="Sunrise" value={data.sunrise} />
          <InfoRow icon="🌙" label="Sunset" value={data.sunset} />
          <InfoRow icon="📅" label="Tithi" value={data.tithi} />
          <InfoRow icon="⭐" label="Nakshatra" value={data.nakshatra} />
          <InfoRow icon="🧘" label="Yoga" value={data.yoga} />
          <InfoRow icon="🌕" label="Moonrise" value={data.moonrise} />
        </div>
        {data.vrat && (
            <div style={{
              backgroundColor: "#FFF8E1",
              borderRadius: "8px",
              padding: "10px",
              textAlign: "center"
            }}>
          <span style={{ color: "#F57F17", fontWeight: "bold" }}>
            🙏 Aaj Ka Vrat: {data.vrat}
          </span>
            </div>
        )}
      </div>
  )
}

// ─── Bhajan Player ────────────────────────────────
function BhajanPlayer({ data }) {
  const [playing, setPlaying] = useState(false)

  return (
      <div style={cardStyle}>
        <h3 style={headingStyle}>🎵 Aaj Ka Bhajan</h3>
        <p style={{ marginBottom: "12px", color: "#333", fontSize: "14px" }}>
          {data.title}
        </p>
        {!playing ? (
            <button onClick={() => setPlaying(true)} style={primaryButton}>
              ▶ Bhajan Sunein
            </button>
        ) : (
            <iframe
                width="100%"
                height="200"
                src={`https://www.youtube.com/embed/${data.youtube_id}?autoplay=1`}
                allow="autoplay; encrypted-media"
                allowFullScreen
                style={{ borderRadius: "8px", border: "none" }}
            />
        )}
      </div>
  )
}

// ─── Set Rashi Card ───────────────────────────────
function SetRashiCard({ onSet }) {
  const rashis = [
    "mesh", "vrishabh", "mithun", "kark",
    "simha", "kanya", "tula", "vrishchik",
    "dhanu", "makar", "kumbh", "meen"
  ]
  const rashiEmoji = {
    mesh: "♈", vrishabh: "♉", mithun: "♊", kark: "♋",
    simha: "♌", kanya: "♍", tula: "♎", vrishchik: "♏",
    dhanu: "♐", makar: "♑", kumbh: "♒", meen: "♓"
  }

  return (
      <div style={cardStyle}>
        <h3 style={headingStyle}>⭐ Apni Rashi Chunein</h3>
        <p style={{ fontSize: "13px", color: "#666", marginBottom: "12px" }}>
          Personalized rashifal aur bhajan ke liye
        </p>
        <div style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", gap: "8px" }}>
          {rashis.map(r => (
              <button key={r} onClick={() => onSet(r)} style={rashiButton}>
                {rashiEmoji[r]} {r.charAt(0).toUpperCase() + r.slice(1)}
              </button>
          ))}
        </div>
      </div>
  )
}

// ─── Rashifal Card ────────────────────────────────
function RashifalCard({ rashi, onReset }) {
  const [data, setData] = useState(null)

  useEffect(() => {
    fetch(`${API}/rashifal/${rashi}`)
        .then(r => r.json())
        .then(setData)
  }, [rashi])

  if (!data) return null

  return (
      <div style={cardStyle}>
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "8px" }}>
          <h3 style={headingStyle}>
            ⭐ Rashifal — {rashi.charAt(0).toUpperCase() + rashi.slice(1)}
          </h3>
          <button onClick={onReset} style={{
            border: "none", background: "none",
            color: "#888", fontSize: "12px", cursor: "pointer"
          }}>
            Change
          </button>
        </div>
        <p style={{ fontSize: "14px", color: "#333", lineHeight: "1.6" }}>
          {data.description}
        </p>
      </div>
  )
}

// ─── Kundali CTA ──────────────────────────────────
function KundaliCTA() {
  const [show, setShow] = useState(false)

  return (
      <>
        <div style={{
          ...cardStyle,
          backgroundColor: "#FFF3E0",
          border: "1px solid #FFB74D",
          textAlign: "center"
        }}>
          <p style={{ fontWeight: "bold", color: "#E65100", marginBottom: "8px" }}>
            📿 Apni Kundali Jaanein — FREE
          </p>
          <p style={{ fontSize: "13px", color: "#666", marginBottom: "12px" }}>
            Rashi, nakshatra, mangal dosha aur yogas paayein
          </p>
          <button onClick={() => setShow(true)} style={primaryButton}>
            Kundali Banao 🙏
          </button>
        </div>
        {show && <KundaliModal onClose={() => setShow(false)} />}
      </>
  )
}

// ─── Kundali Modal ────────────────────────────────
function KundaliModal({ onClose }) {
  const [form, setForm] = useState({
    name: "", dob: "", tob: "",
    latitude: "26.8467", longitude: "80.9462"
  })
  const [result, setResult] = useState(null)
  const [loading, setLoading] = useState(false)

  const handleSubmit = async () => {
    if (!form.name || !form.dob || !form.tob) return
    setLoading(true)
    try {
      const res = await fetch(`${API}/kundali`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form)
      })
      const data = await res.json()
      setResult(data)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  return (
      <div style={{
        position: "fixed", inset: 0,
        backgroundColor: "rgba(0,0,0,0.6)",
        display: "flex", alignItems: "center",
        justifyContent: "center", zIndex: 1000, padding: "16px"
      }}>
        <div style={{
          backgroundColor: "white", borderRadius: "16px",
          padding: "20px", width: "100%", maxWidth: "400px",
          maxHeight: "85vh", overflowY: "auto"
        }}>
          <div style={{ display: "flex", justifyContent: "space-between", marginBottom: "16px" }}>
            <h3 style={{ color: "#B8541B" }}>📿 Kundali Banao</h3>
            <button onClick={onClose} style={{ border: "none", background: "none", fontSize: "20px", cursor: "pointer" }}>✕</button>
          </div>

          {!result ? (
              <>
                {[
                  { key: "name", label: "Aapka Naam", type: "text", placeholder: "Ramesh Kumar" },
                  { key: "dob", label: "Janm Tithi", type: "date" },
                  { key: "tob", label: "Janm Samay", type: "time" },
                ].map(field => (
                    <div key={field.key} style={{ marginBottom: "12px" }}>
                      <label style={{ display: "block", marginBottom: "4px", fontSize: "13px", color: "#666" }}>
                        {field.label}
                      </label>
                      <input
                          type={field.type}
                          placeholder={field.placeholder}
                          value={form[field.key]}
                          onChange={e => setForm({ ...form, [field.key]: e.target.value })}
                          style={inputStyle}
                      />
                    </div>
                ))}
                <button
                    onClick={handleSubmit}
                    disabled={loading}
                    style={primaryButton}>
                  {loading ? "Kundali ban rahi hai... 🙏" : "Kundali Dekho — FREE"}
                </button>
              </>
          ) : (
              <div>
                <div style={{ textAlign: "center", marginBottom: "16px" }}>
                  <div style={{ fontSize: "48px" }}>🙏</div>
                  <h3 style={{ color: "#B8541B" }}>Kundali Tayaar!</h3>
                </div>
                <ResultRow label="Naam" value={result.name} />
                <ResultRow label="Nakshatra" value={result.nakshatra} />
                <ResultRow label="Nakshatra Lord" value={result.nakshatra_lord} />
                <ResultRow label="Chandra Rasi" value={result.chandra_rasi} />
                <ResultRow label="Soorya Rasi" value={result.soorya_rasi} />
                <ResultRow label="Shubh Rang" value={result.color} />
                <ResultRow label="Birth Stone" value={result.birth_stone} />
                <ResultRow label="Best Direction" value={result.best_direction} />
                <ResultRow
                    label="Mangal Dosha"
                    value={result.mangal_dosha ? "Haan ⚠️" : "Nahi ✅"}
                />
                <div style={{
                  backgroundColor: "#FFF8E1", borderRadius: "8px",
                  padding: "10px", marginTop: "12px"
                }}>
                  <p style={{ fontWeight: "bold", marginBottom: "6px", color: "#B8541B" }}>
                    🧘 Yogas
                  </p>
                  {result.yogas?.map((y, i) => (
                      <p key={i} style={{ fontSize: "13px", color: "#555", marginBottom: "4px" }}>
                        • {y.description}
                      </p>
                  ))}
                </div>
                <button onClick={onClose} style={{ ...primaryButton, marginTop: "16px" }}>
                  ✅ Done
                </button>
              </div>
          )}
        </div>
      </div>
  )
}

// ─── Helper Components ────────────────────────────
function InfoRow({ icon, label, value }) {
  return (
      <div style={{ fontSize: "13px" }}>
        <span style={{ color: "#888" }}>{icon} {label}: </span>
        <span style={{ fontWeight: "bold", color: "#333" }}>{value || "—"}</span>
      </div>
  )
}

function ResultRow({ label, value }) {
  return (
      <div style={{
        display: "flex", justifyContent: "space-between",
        padding: "8px 0", borderBottom: "1px solid #f0f0f0",
        fontSize: "14px"
      }}>
        <span style={{ color: "#888" }}>{label}</span>
        <span style={{ fontWeight: "bold", color: "#333" }}>{value}</span>
      </div>
  )
}

// ─── Shared Styles ────────────────────────────────
const cardStyle = {
  backgroundColor: "white",
  borderRadius: "12px",
  padding: "16px",
  marginBottom: "12px",
  boxShadow: "0 2px 8px rgba(0,0,0,0.08)"
}

const headingStyle = {
  color: "#B8541B",
  marginBottom: "12px",
  fontSize: "16px"
}

const primaryButton = {
  width: "100%",
  backgroundColor: "#E65100",
  color: "white",
  border: "none",
  padding: "14px",
  borderRadius: "8px",
  fontSize: "15px",
  cursor: "pointer"
}

const rashiButton = {
  padding: "8px 4px",
  backgroundColor: "#FFF8F0",
  border: "1px solid #FFB74D",
  borderRadius: "6px",
  cursor: "pointer",
  fontSize: "12px",
  color: "#E65100"
}

const inputStyle = {
  width: "100%",
  padding: "10px",
  border: "1px solid #ddd",
  borderRadius: "8px",
  fontSize: "14px",
  boxSizing: "border-box"
}