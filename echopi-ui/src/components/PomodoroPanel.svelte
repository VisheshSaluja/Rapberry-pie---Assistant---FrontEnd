<script>
  export let running = false;
  export let minutes = "25";
  export let seconds = "00";
</script>

<div class="pomodoro-container">
  <h1 class="timer {running ? 'glow-active' : ''}">
    {minutes}:{seconds}
  </h1>

  {#if running}
    <p class="status">FOCUS MODE ACTIVE</p>
  {:else}
    <p class="status">Idle</p>
  {/if}
</div>

<style>
  .pomodoro-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    color: #00ff9c;
  }

  .timer {
    font-size: 62px;
    letter-spacing: 2px;
    margin-bottom: 10px;
    color: #00ff9c;
  }

  /* Glow only when running */
  .glow-active {
    animation: glow 2s infinite;
  }

  .status {
    font-size: 22px;
    opacity: 0.7;
  }

  @keyframes glow {
    0% { text-shadow: 0 0 10px #00ffbb; }
    50% { text-shadow: 0 0 25px #00ffbb; }
    100% { text-shadow: 0 0 10px #00ffbb; }
  }
</style>
