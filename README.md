# Rapberry-pie---Assistant---FrontEnd
Svelte UI for real-time productivity dashboard
<img width="1436" height="775" alt="Screenshot 2025-11-16 at 4 47 39 PM" src="https://github.com/user-attachments/assets/da4e975d-8673-4003-a8a1-cecf6d2a1d95" />

## Setup
### 1. Install dependencies
```bash
  npm install
```
### 2. Start development server
```bash
  npm run dev
```
### 3. Connect to backend
The UI expects the backend WebSocket server at:
```bash
  ws://localhost:9001/ws
# or
  ws://<raspberry-pi-ip>:9001/ws

```
### Testing

- Start the backend
- Run the Svelte dev server
- Interact with notes, Pomodoro, and watch updates sync live

### Build for Production
```bash
  npm run build
```
Outputs go to the /dist folder.
