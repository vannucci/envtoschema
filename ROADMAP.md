# envtoschema — Roadmap & Motivation

## Why Did We Want This?

Generic JSON Schema validators like AJV or `jsonschema` are excellent at what they do — but they know nothing about AppConfig. They validate structure. They don't know where your schema lives, whether someone changed it in the AWS console last Tuesday, or whether your dev and prod schemas have quietly drifted apart.

The existing workflow for most teams is: write a `.env` file, have someone manually translate it into a JSON Schema, upload it to AppConfig, and hope nobody edits it through the dashboard without telling anyone. That manual step is error-prone and the drift problem is silent.

This tool exists to close that gap. It automates schema generation from your existing config, gives you a guided UI to add the semantics a tool can't infer (required fields, bounds, descriptions), and then grows into a build-time enforcement layer that is AppConfig-aware — not just JSON Schema-aware.

The validation logic itself isn't the interesting part. The AppConfig integration and the semantic layer on top are.

---

## Prior Art — What Already Exists

These tools each solve part of the problem. None solve all of it.

| Tool | What it does | What it doesn't do |
|---|---|---|
| [AJV](https://ajv.js.org/) (JS/Node) | Validates JSON against a JSON Schema | Knows nothing about AppConfig, doesn't generate schemas |
| [jsonschema](https://python-jsonschema.readthedocs.io/) (Python) | Same as AJV, Python ecosystem | Same limitations |
| [gojsonschema](https://github.com/xeipuuv/gojsonschema) / [santhosh-tekuri/jsonschema](https://github.com/santhosh-tekuri/jsonschema) (Go) | Validates JSON against a schema in Go | Schema generation, AppConfig integration: none |
| [check-jsonschema](https://check-jsonschema.readthedocs.io/) (CLI) | Drop-in CI validator, pip installable | No generation, no AppConfig awareness |
| AWS Console | Edit AppConfig schemas manually | Exactly the problem we're solving |

The pattern is the same across all of them: **bring your own schema, bring your own config, get a pass/fail**. You still have to write the schema by hand, you still have to know to pull the remote version before validating, and none of them will tell you that someone changed the schema in the AWS console last week.

`envtoschema` is the missing first step — generate the schema from what you already have, refine it with guardrails a tool can't infer on its own, and then enforce it at every layer from local dev to CI to production drift detection.

---

## V0 — Core Tool

- [x] CLI reads a flat JSON config file
- [ ] Type inference for `string`, `integer`, `number`, `boolean`
- [ ] Localhost form UI to review and override inferred types
- [ ] Mark fields as required, add descriptions
- [ ] Generate a valid AppConfig JSON Schema file on disk
- [ ] Basic validator: does this JSON config pass the schema?

---

## V1 — Inferencer & Parser Improvements

- [ ] Nested object support (1–2 levels)
- [ ] UUID / format hints (`string` + `"format": "uuid"`)
- [ ] Handle `.env` file format as input, not just JSON
- [ ] Smarter integer bounds suggestions (port ranges, percentages, etc.)

---

## V2 — Build Step & CI Integration

- [ ] `--validate` flag: exits with code 1 on failure, pipeline-friendly
- [ ] `--schema` flag: validate a config against a local schema file
- [ ] `--remote` flag: pull the live schema directly from AWS AppConfig API and validate against that — catches drift from console edits automatically
- [ ] Schema drift warning: detect and report when local schema and remote AppConfig schema have diverged

---

## V3 — Multi-Environment & Semantic Validation

- [ ] Multi-environment validation: validate one config against dev and prod schemas in a single pass
- [ ] Semantic / convention validation — rules JSON Schema can't express:
  - Naming conventions (e.g. flag names must be camelCase)
  - Cross-field logic (e.g. if `enabled: true`, an `owner` field is required)
  - Range enforcement (e.g. percentage values must be 0–100)
  - Deprecation enforcement (no field removed without a `deprecated` key)

---

## V4 — Living Schema & Feature Flags

- [ ] Feature flag support: schema treated as a living document, updatable without full regeneration
- [ ] Human-readable config diff before deploy: show what changed, not just pass/fail
- [ ] Pull remote AppConfig, diff against local, summarize changes in plain English
