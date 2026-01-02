# Todo TUI

A terminal-based todo list manager with animations, themes, and notifications.

![Preview](assets/preview.gif)

## Features

- 10 color themes (Catppuccin, Nord, Gruvbox, Dracula, and more)
- 30 unique completion animations
- Timer notifications with desktop alerts
- Inline editing for quick task management
- Sorting modes to organize tasks
- Automatic save functionality

## Quick Start

### Download Pre-built Binary (Recommended)

1. Go to [Releases](https://github.com/nirabyte/todo/releases)
2. Download the binary for your system.
3. Extract the folder and run:
   - **Windows**: Double-click `todo.exe` or run `.\todo.exe` in terminal
   - **macOS/Linux**: Run `./todo-app-name` in terminal (you may need to `chmod +x todo-app-name` first)

### Install with Go

If you have Go installed:

```bash
go install github.com/nirabyte/todo@latest
```

Then run:

```bash
todo
```

If Go’s bin directory isn’t in your PATH, add it by running:

```
export PATH=$PATH:$HOME/go/bin
```

Then reload your shell:

```
source ~/.bashrc   # or ~/.zshrc if you use Zsh
```

### Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/nirabyte/todo.git
   cd todo
   ```

2. Build for your platform:

   ```bash
   # Using Make (if available)
   make build

   # Or manually
   go build -o build/todo ./cmd/todo
   ```

3. Run it:

   ```bash
   # Windows
   .\build\todo.exe

   # macOS/Linux
   ./build/todo
   ```

## Requirements

- A terminal that supports colors (most modern terminals work)
- A Nerd Font installed (for icons and symbols to display correctly)
  - Download from [nerdfonts.com](https://www.nerdfonts.com/)
  - Popular choices: FiraCode, JetBrains Mono, Hack, Meslo
- Desktop notifications enabled (optional, for timer alerts)
- Go 1.25.5 or later (only if building from source)

## Usage

### Starting the App

Run:

```bash
todo
```

On first launch, you'll see helpful hints to get started.

### Navigation

| Key             | Action    |
| --------------- | --------- |
| `↑` or `k`      | Move up   |
| `↓` or `j`      | Move down |
| `q` or `Ctrl+C` | Quit      |

### Managing Tasks

| Key     | Action                     |
| ------- | -------------------------- |
| `n`     | New task                   |
| `e`     | Edit selected task         |
| `d`     | Delete selected task       |
| `Space` | Toggle complete/uncomplete |
| `Enter` | Confirm (when editing)     |
| `Esc`   | Cancel (when editing)      |

![Edit Task](assets/edit.gif)

### Setting Timers

Press `@` on any task to set a reminder timer.

**Examples:**

- `10m` = 10 minutes
- `1h30m` = 1 hour 30 minutes
- `45s` = 45 seconds
- `2h15m30s` = 2 hours 15 minutes 30 seconds

When the timer expires, you'll get a desktop notification. The countdown displays next to the task.

![Timer Notification](assets/timer.gif)

### Customization

| Key | Action                      |
| --- | --------------------------- |
| `t` | Cycle through themes        |
| `s` | Cycle through sorting modes |

## Themes

Choose from 10 color themes:

1. **Catppuccin** - Soothing pastel colors
2. **Nord** - Cool arctic blues
3. **Gruvbox** - Retro warm tones
4. **Dracula** - Dark purple theme
5. **Tokyo Night** - Clean dark theme
6. **Rose Pine** - Natural earthy colors
7. **Everforest** - Comfortable green tones
8. **One Dark** - Popular dark theme
9. **Solarized** - Easy on the eyes
10. **Kanagawa** - Inspired by Japanese art

Press `t` to cycle through themes. Your choice is saved automatically.

![Theme Selection](assets/theme.gif)

## Sorting Modes

Organize your tasks with three sorting options:

- **Off** - Keep tasks in the order you created them
- **Todo First** - Incomplete tasks at the top
- **Done First** - Completed tasks at the top

Press `s` to cycle through modes. Your preference is saved.

![Sorting Modes](assets/sort.gif)

## Completion Animations

When you complete a task, one of 30 unique animations plays. These include:

- Sparkle effects
- Rainbow transitions
- Typewriter effects
- Matrix-style animations
- And 26 more unique effects

Each animation is randomly selected for a fresh experience.

![Completion Animations](assets/animation.gif)

## Data Storage

Your tasks are saved automatically in a file called `todos.json` in the same directory where you run the app.

**What's saved:**

- All your tasks (title, completion status, due dates)
- Your selected theme
- Your sorting preference

You can backup this file, edit it manually, or move it to another computer.

## Development

### Project Structure

```
todo/
├── assets/          # Images/icons
├── cmd/todo/
│   └── main.go      # Entry point
├── internal/
│   ├── app/         # App setup
│   ├── config/      # Config handling
│   ├── models/      # Data & logic
│   └── ui/          # UI code
├── go.mod           # Go module dependencies
├── go.sum           # Dependency checksums
└── README.md
```

### Building for All Platforms

```bash
# Build all platforms (requires Make)
make build-all

# Or prepare release binaries with simple names
make release
```

This creates binaries for:

- Windows (amd64, 386, arm64)
- macOS (amd64, arm64)
- Linux (amd64, 386, arm64, arm)

## Roadmap

- [x] Core app: tasks, themes, animations, timers, inline edit, sorting
- [x] JSON-based persistence (`todos.json`) with automatic save
- [x] 10 color themes (Catppuccin, Nord, Gruvbox, Tokyo Night,etc)
- [x] 30 unique completion animations
- [x] Timer reminders with desktop notifications
- [x] Inline task editing
- [x] Sorting modes (Off / Todo First / Done First)
- [ ] Add search/filter option
- [ ] Tasks with description,  status indicator and active filename in header
- [ ] Add CLI arguments for dates/days, todolist, file selection, and import/export commands
- [ ] Extend timers to days/weeks/months and add absolute due dates with visual indicators
- [ ] Add markdown support
- [ ] JSON/Markdown import/export feature
- [ ] More UI/UX improvements 

## Built With

- **Go**
- **Bubble Tea** - TUI framework
- **Lip Gloss** - Styling library
- **Beeep** - Desktop notifications

## Contributing

Please see the full contribution guidelines in `CONTRIBUTING.md`.

[CONTRIBUTING.md](CONTRIBUTING.md)

## License

This project is licensed under the MIT License — see the `LICENSE` file for details.

[LICENSE](LICENSE)

## Contact

If you have questions, suggestions, or just want to chat, feel free to reach out to me on Discord:

[Message me on Discord](https://discord.com/users/863252422913687614)
