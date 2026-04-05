import m from "mithril"
import "./style.css"

const storageKey = "easy-login.device-token"
const apiBaseUrl =
  import.meta.env.VITE_API_BASE_URL ||
  `http://localhost:${import.meta.env.VITE_API_PORT || "8080"}`

const state = {
  apiBaseUrl,
  displayName: "",
  recoveryPassphrase: "",
  deviceToken: window.localStorage.getItem(storageKey) || "",
  identity: null,
  error: "",
  loading: false,
}

async function createGuestIdentity() {
  state.loading = true
  state.error = ""

  try {
    const response = await fetch(`${state.apiBaseUrl}/identities/guest`, {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({display_name: state.displayName}),
    })

    const body = await response.json()
    if (!response.ok) {
      throw new Error(body.error || "guest creation failed")
    }

    state.identity = body
    state.deviceToken = body.device_token
    window.localStorage.setItem(storageKey, body.device_token)
  } catch (error) {
    state.error = error.message
  } finally {
    state.loading = false
    m.redraw()
  }
}

async function resumeIdentity() {
  if (!state.deviceToken) {
    return
  }

  state.loading = true
  state.error = ""

  try {
    const response = await fetch(`${state.apiBaseUrl}/identities/resume`, {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({device_token: state.deviceToken}),
    })

    const body = await response.json()
    if (!response.ok) {
      window.localStorage.removeItem(storageKey)
      state.deviceToken = ""
      throw new Error(body.error || "resume failed")
    }

    state.identity = body
  } catch (error) {
    state.error = error.message
  } finally {
    state.loading = false
    m.redraw()
  }
}

async function claimIdentity() {
  if (!state.deviceToken) {
    return
  }

  state.loading = true
  state.error = ""

  try {
    const response = await fetch(`${state.apiBaseUrl}/identities/claim`, {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
        device_token: state.deviceToken,
        recovery_passphrase: state.recoveryPassphrase,
      }),
    })

    const body = await response.json()
    if (!response.ok) {
      throw new Error(body.error || "claim failed")
    }

    state.identity = body
  } catch (error) {
    state.error = error.message
  } finally {
    state.loading = false
    m.redraw()
  }
}

async function recoverIdentity() {
  state.loading = true
  state.error = ""

  try {
    const response = await fetch(`${state.apiBaseUrl}/identities/recover`, {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
        recovery_passphrase: state.recoveryPassphrase,
      }),
    })

    const body = await response.json()
    if (!response.ok) {
      throw new Error(body.error || "recovery failed")
    }

    state.identity = body
    state.deviceToken = body.device_token
    window.localStorage.setItem(storageKey, body.device_token)
  } catch (error) {
    state.error = error.message
  } finally {
    state.loading = false
    m.redraw()
  }
}

const App = {
  oninit() {
    void resumeIdentity()
  },
  view() {
    return m("main.shell", [
      m("section.card", [
        m("p.eyebrow", "easy-login"),
        m("h1", "Guest identity with continuity"),
        m(
          "p.copy",
          "Create a guest identity from a display name and resume it later on the same browser with the stored device token.",
        ),
        m("label.label", {for: "display_name"}, "Display name"),
        m("input.input", {
          id: "display_name",
          value: state.displayName,
          oninput: (event) => {
            state.displayName = event.target.value
          },
          placeholder: "henrique",
        }),
        m(
          "button.button",
          {
            disabled: state.loading,
            onclick: () => void createGuestIdentity(),
          },
          state.loading ? "Working..." : "Create guest identity",
        ),
        state.deviceToken
          ? m("button.button.secondary", {disabled: state.loading, onclick: () => void resumeIdentity()}, "Resume from device token")
          : null,
        !state.deviceToken
          ? [
              m("label.label", {for: "recovery_passphrase"}, "Recover with passphrase"),
              m("input.input", {
                id: "recovery_passphrase",
                value: state.recoveryPassphrase,
                oninput: (event) => {
                  state.recoveryPassphrase = event.target.value
                },
                placeholder: "moon-river-42",
              }),
              m(
                "button.button.secondary",
                {
                  disabled: state.loading,
                  onclick: () => void recoverIdentity(),
                },
                "Recover identity",
              ),
            ]
          : null,
        state.identity && state.identity.claim_status === "guest"
          ? [
              m("label.label", {for: "recovery_passphrase"}, "Recovery passphrase"),
              m("input.input", {
                id: "recovery_passphrase",
                value: state.recoveryPassphrase,
                oninput: (event) => {
                  state.recoveryPassphrase = event.target.value
                },
                placeholder: "moon-river-42",
              }),
              m(
                "button.button.secondary",
                {
                  disabled: state.loading,
                  onclick: () => void claimIdentity(),
                },
                "Claim identity",
              ),
            ]
          : null,
        state.error ? m("p.error", state.error) : null,
      ]),
      state.identity
        ? m("section.card.identity", [
            m("p.eyebrow", "Current identity"),
            m("dl.identity-grid", [
              m("div", [m("dt", "Player ID"), m("dd", state.identity.player_id)]),
              m("div", [m("dt", "Display name"), m("dd", state.identity.display_name)]),
              m("div", [m("dt", "Claim status"), m("dd", state.identity.claim_status)]),
              state.deviceToken ? m("div", [m("dt", "Device token"), m("dd.code", state.deviceToken)]) : null,
            ]),
          ])
        : null,
    ])
  },
}

m.mount(document.getElementById("app"), App)
