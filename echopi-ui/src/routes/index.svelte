<script>
  import MatrixBackground from "../components/MatrixBackground.svelte";
  import FloatingOrb from "../components/FloatingOrb.svelte";
  import NotesPanel from "../components/NotesPanel.svelte";
  import PomodoroPanel from "../components/PomodoroPanel.svelte";
  
  import { notesStore, pomodoroStore, voiceState } from "../lib/eventBus";

  // TODO: Wire up WebSocket here later
</script>

<MatrixBackground />

<FloatingOrb />

{#if $notesStore.length > 0}
  <NotesPanel />
{/if}

{#if $pomodoroStore.running}
  <PomodoroPanel />
{/if}
