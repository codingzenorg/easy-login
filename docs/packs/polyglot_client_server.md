# Pack: polyglot_client_server

## Purpose

Use this pack when one bounded model spans more than one runtime, such as a browser client and a server process.

This is the pack to prefer when the semantics are shared but execution is intentionally split across technologies.

---

## Shape

- language: multiple
- runtime topology: client/server or peer runtimes
- architecture mode: explicit contracts across runtime boundaries

---

## Layout

```text
src/
  client/
  server/
  shared_contracts/

tests/
  client/
  server/
  integration/
  contracts/
```

---

## Rules

- keep semantic truth shared in `docs/semantics/`
- make protocol, message, and event contracts explicit
- let slices declare which runtime targets are in scope
- do not hide distributed boundaries behind a fake single-process abstraction

---

## Good Fit

- browser client plus Go or Python server
- game logic split across client prediction and server authority
- shared event or message contracts across runtimes
