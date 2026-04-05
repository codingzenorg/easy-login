# Recovery Impact Analysis

## Reason For Impact Note

After claim is implemented, the strongest remaining model gap is that the ownership proof still cannot recover identity on a new browser. The product promise is now structurally present but not operationally complete.

## Impacted Areas

- backend application
  a recovery use case is needed
- persistence
  player lookup by recovery-proof derivative must be supported
- device registrations
  recovery should issue a usable device token for the new browser
- browser client
  a recovery entry path is needed when no current device token exists

## Refinement Judgment

Recovery should come next.

Why this slice now:

- the ownership model already exists
- the passphrase has meaning only if it can restore continuity after device loss
- it completes the central portability promise without introducing heavier authentication systems

## Boundaries To Preserve

- do not introduce passphrase rotation yet
- do not implement device revocation policies yet
- do not expand into email or OAuth
- do not redesign the current guest or claim flows unless necessary for recovery correctness

## Expected Follow-Through For Build

- add recovery lookup against persisted claim data
- issue a device token for the recovering browser
- preserve compatibility with the current resume flow and claimed-state persistence
