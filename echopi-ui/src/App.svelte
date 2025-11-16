<script>
  // NOTE: Assuming these component imports are correct and available in your environment.
  import FloatingOrb from "./components/FloatingOrb.svelte";
  import NotesPanel from "./components/NotesPanel.svelte";
  import PomodoroPanel from "./components/PomodoroPanel.svelte";

  let ui = {
    page: "home",
    notes: "",
    pomodoro: { running: false, minutes: "25", seconds: "00" }
  };

  let hotwordActive = false;

  // =============================
  // 🔌 STRONG WebSocket Handler
  // =============================
  let socket = new WebSocket("ws://localhost:9001/ws");

socket.onmessage = (event) => {
  const raw = event.data;

  // 🔒 Only accept real JSON that includes "type"
  if (!raw || typeof raw !== "string" || !raw.includes('"type"')) {
    // Silently ignore dev-server garbage
    return;
  }

  let msg;
  try {
    msg = JSON.parse(raw);
  } catch {
    return; // silently ignore
  }

  // 🔥 Hotword
  if (msg.type === "hotword") {
    hotwordActive = true;
    setTimeout(() => (hotwordActive = false), 1500);
    ui.page = "home";
    return;
  }

  // 🔥 UI update
  if (msg.type === "ui") {
    if (msg.page) ui.page = msg.page;
    if (msg.notes) ui.notes = msg.notes;
    if (msg.data) ui.pomodoro = msg.data;

    console.log("UI updated:", msg);
  }
};



  // MATRIX BACKGROUND
  let canvas;
  $: if (canvas) startMatrix();

  function startMatrix() {
    // Only run if canvas context is available
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    let w = (canvas.width = window.innerWidth);
    let h = (canvas.height = window.innerHeight);

    // Update width/height on resize
    window.addEventListener('resize', () => {
        w = (canvas.width = window.innerWidth);
        h = (canvas.height = window.innerHeight);
    });

    const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789あいうえお";
    const fontSize = 18;
    const columns = Math.floor(w / fontSize);
    const drops = new Array(columns).fill(1);

    function draw() {
      ctx.fillStyle = "rgba(0,0,0,0.05)";
      ctx.fillRect(0, 0, w, h);

      ctx.fillStyle = "#00ff88";
      ctx.font = fontSize + "px monospace";

      for (let i = 0; i < drops.length; i++) {
        const char = chars[Math.floor(Math.random() * chars.length)];
        ctx.fillText(char, i * fontSize, drops[i] * fontSize);

        if (drops[i] * fontSize > h && Math.random() > 0.975) {
          drops[i] = 0;
        }

        drops[i]++;
      }
    }

    // Clear any previous interval before setting a new one (good practice)
    const interval = setInterval(draw, 33);

    // Cleanup function (if App.svelte ever unmounts, though unlikely here)
    return () => clearInterval(interval);
  }
</script>


<!-- MATRIX BACKGROUND -->
<div class="matrix-bg">
  <canvas bind:this={canvas} class="matrix-canvas"></canvas>
  <div class="scanlines"></div>
  <div class="crt-overlay"></div>
</div>

<!-- CENTER ORB -->
<div class="center-orb">
  <FloatingOrb {hotwordActive} />
</div>

<!-- PANELS -->
{#if ui.page === "notes"}
  <div class="panel notes-panel glitch-in">
    <NotesPanel notes={ui.notes} />
  </div>
{/if}

{#if ui.page === "pomodoro"}
  <div class="panel pomodoro-panel glitch-in">
    <PomodoroPanel data={ui.pomodoro} />
  </div>
{/if}

<style>
  /* Base reset is good for a full-screen app */
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'monospace', 'Inter', sans-serif;
  }
  
  .matrix-bg {
    position: fixed;
    inset: 0;
    background: black;
    overflow: hidden;
    z-index: -3;
  }
  .matrix-canvas {
    width: 100%;
    height: 100%;
    display: block;
    filter: brightness(1.3) contrast(1.4);
  }

  .scanlines {
    pointer-events: none;
    position: fixed;
    inset: 0;
    background: repeating-linear-gradient(
      to bottom,
      rgba(0, 255, 100, 0.04) 0px,
      rgba(0, 255, 100, 0.04) 2px,
      rgba(0, 0, 0, 0.3) 3px
    );
    z-index: -2;
    opacity: 0.25;
  }

  .crt-overlay {
    pointer-events: none;
    position: fixed;
    inset: 0;
    z-index: -1;
    background: radial-gradient(
      circle at center,
      rgba(0, 255, 100, 0.05),
      transparent 60%
    );
    animation: crt-wobble 6s infinite ease-in-out;
  }

  @keyframes crt-wobble {
    0%,100% { transform: scale(1); }
    50% { transform: scale(1.01); }
  }

  .center-orb {
    position: fixed;
    inset: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 5;
  }

  .panel {
    position: fixed;
    bottom: 10%;
    left: 50%;
    transform: translateX(-50%);
    z-index: 10;
    /* Optional: Add a subtle glow/box-shadow for CRT feel */
    box-shadow: 0 0 10px rgba(0, 255, 136, 0.5);
    border: 1px solid rgba(0, 255, 136, 0.5);
    border-radius: 8px;
    padding: 1rem;
    background: rgba(0, 0, 0, 0.85); /* Slightly darker background for better contrast */
  }

  .glitch-in {
    animation: glitchIn 0.45s ease-out;
  }

  @keyframes glitchIn {
    0% {
      opacity: 0;
      transform: translate(-50%, 20px) scale(0.95);
      filter: blur(6px);
    }
    30% {
      opacity: 1;
      transform: translate(-50%, -4px) scale(1.03);
      filter: blur(1px);
    }
    60% {
      transform: translate(-50%, 2px) scale(0.99);
      filter: blur(0.5px);
    }
    100% {
      transform: translate(-50%, 0px) scale(1);
      filter: blur(0px);
    }
  }
</style>