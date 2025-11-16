<script>
  export let active = false;
</script>

<div class="orb-container">
  <div class="orb {active ? 'orb-active' : ''}"></div>
</div>

<style>
  .orb-container {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

  .orb {
    width: 110px;
    height: 110px;
    border-radius: 50%;
    background: radial-gradient(circle, #00ffcc, #007755);
    box-shadow: 0 0 20px #00ffcc;
    animation: float 5s ease-in-out infinite;
  }

  .orb-active {
    box-shadow: 0 0 35px #00ffcc, 0 0 70px #00ffcc;
    animation: pulse 1.5s infinite;
  }

  @keyframes float {
    0% { transform: translateY(-10px); }
    50% { transform: translateY(10px); }
    100% { transform: translateY(-10px); }
  }

  @keyframes pulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.07); }
  }
</style>
