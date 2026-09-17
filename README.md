# Practical Go Foundations — Ardan Labs

My exercises and notes from the **"Practical Go Foundations"** course by
[Ardan Labs](https://www.ardanlabs.com/) (instructor: Miki Tebeka).
Started 5 September 2026.

Each directory is a standalone `package main` — a self-contained exercise
from one module of the course.

## Progress

| Exercise | Topics | Status |
|---|---|:--:|
| [hw/hw.go](hw/hw.go) | Hello World, `fmt` | ✅ |
| [banner/banner.go](banner/banner.go) | Strings, Unicode, `strings.Repeat` | ✅ |
| [github/github.go](github/github.go) | REST APIs, JSON decoding, struct tags | ✅ |
| [kill_server/kill_server.go](kill_server/kill_server.go) | Files, `defer`, error wrapping, `log/slog` | ✅ |
| [sha1/sha1.go](sha1/sha1.go) | `io.Reader` / `io.Writer`, gzip, hashing | ✅ |
| [cart/cart.go](cart/cart.go) | Slices: `len`, indexing | 🚧 |
| `div/div.go` | Handling panics | ⬜ not started |

## What each exercise covers

**`hw`** — The starting point: `fmt.Println`.

**`banner`** — Centres text inside a fixed-width banner using `strings.Repeat`.
*Note:* it uses `len(text)`, which counts **bytes**, not runes — so the padding
misaligns on non-ASCII input like `"héllo"`. Fixing it with
`utf8.RuneCountInString` is the actual Unicode lesson.

**`github`** — Queries `api.github.com/users/<login>` and returns the user's
name and public repo count. Decodes with `json.NewDecoder` and a struct tag
(`json:"public_repos"`). The parsing lives in `parseResponse(r io.Reader)`
rather than taking the `*http.Response` directly, so it can be tested without
a network call.

**`kill_server`** — Reads a PID from a file and deletes it. Covers the
acquire → check error → `defer` release idiom, `defer` in a closure so
`Close()`'s error isn't discarded, error wrapping with `%w`, and unwrapping
with `errors.Is` / `errors.Unwrap` against `fs.ErrNotExist`.

**`sha1`** — Computes the SHA-1 of a file, transparently decompressing it only
when the name ends in `.gz`. Equivalent to `cat http.log.gz | gunzip | sha1sum`.
The key move is `var r io.Reader = file`: declaring the variable as an
*interface* allows swapping in `*gzip.Reader` later. Declaring it with
`r := file` would infer `*os.File` and fail to compile — and using
`r, err := gzip.NewReader(file)` inside the `if` would shadow it, leaving the
outer `r` pointing at the still-compressed file.

## Running the exercises

Each exercise reads its data files by **relative path**, so you have to run it
from inside its own directory:

```bash
cd sha1
go run .
```

Running `go run ./sha1` from the repo root compiles fine but fails at runtime —
the working directory is wherever you invoked it, so `http.log.gz` isn't found.

Build everything at once from the root:

```bash
go build ./...
go vet ./...
```

### `kill_server` needs a PID file

It **deletes** `server.pid` on success, so recreate it before each run:

```bash
cd kill_server
echo 7 > server.pid
go run .
```

Run it with the file absent to exercise the `fs.ErrNotExist` path.

## Requirements

- Go 1.24 or later (module: `practical_go`)
- No external dependencies — standard library only

## Layout

```
.
├── banner/       strings and Unicode
├── cart/         slices (in progress)
├── github/       REST + JSON
├── hw/           hello world
├── kill_server/  files, defer, errors
├── sha1/         io.Reader / io.Writer
└── Material/     course PDFs (not redistributable)
```

Compiled binaries, `*.pid` files and editor backups are excluded via
[.gitignore](.gitignore).
