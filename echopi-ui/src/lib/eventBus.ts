import { writable } from 'svelte/store';

export const voiceState = writable<"idle"|"listening"|"thinking"|"speaking">("idle");
export const notesStore = writable<string[]>([]);
export const pomodoroStore = writable({
    running: false,
    minutesLeft: 25,
    mode: "focus" // focus | break
});
