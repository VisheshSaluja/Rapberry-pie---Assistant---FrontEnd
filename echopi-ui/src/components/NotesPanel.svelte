<script>
  export let notes = "";
</script>

<div class="notes-container">
  <h2 class="title">NOTES</h2>
  {#if notes.trim() === ""}
    <p class="empty">No notes yet.</p>
  {:else}
    <pre class="notes">{notes}</pre>
  {/if}
</div>

<style>
  .notes-container {
    padding: 40px;
    color: #00ff9c;
    font-family: monospace;
  }

  .title {
    font-size: 28px;
    margin-bottom: 15px;
    opacity: 0.9;
  }

  .empty {
    opacity: 0.5;
  }

  .notes {
    white-space: pre-wrap;
    background: rgba(0, 255, 150, 0.05);
    border-left: 2px solid #00ff9c;
    padding: 20px;
    border-radius: 6px;
    font-size: 18px;
  }
</style>
