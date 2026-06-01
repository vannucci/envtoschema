```
"...because we can only juggle so many balls, you have to make a decision. How many of those balls do you want to be incidental complexity and how many do you want to be?"

Rich Hickey, "Simple Made Easy"
```

# What is this?

This is a CI/CD pipeline ready, AWS AppConfig-specific schema validation tool.

# The Problem

The motivation came from a real migration I've been a part of from an ECS deployment type to an EKS deployment type for our app. As part of this transition we wanted to use AWS AppConfig to manage the configs of our pods instead of our previous method. The Freeform AWS AppConfig allows the users to upload a JSON Schema version 4.x for inline schema validation. The problem was that the validation of an config file happens at time of deployment to pods, not at build time, potentially causing bad schemas to exist in our repositories and only becoming a problem when we deploy it to our cluster. Also while plenty of JSON Schema exist (AJV, gojsonschema) they are too generic for this application and would need to be adapted regardless to be used in a CI/CD pipeline as well as draw their source-of-truth from AppConfig itself.

This software's goal is to fill that particular niche: an AWS AppConfig aware JSON Schema Validator that is designed with operation in a CI/CD pipeline first.

# Installation

## From source

Requires Go 1.21+

```bash
git clone https://github.com/vannucci/envtoschema
cd envtoschema
make build
```

The binary will be in the project root. Move it to your PATH:

```bash
mv envtoschema /usr/local/bin/
```

# Usage

## Generate a schema from a config file

```bash
envtoschema -target config.json -output schema.json
```

This will:

1. Read and validate `config.json`
2. Open a local browser UI to review inferred types
3. Write the generated schema to `schema.json` on form submission

## Flags

| Flag      | Default       | Description                                          |
| --------- | ------------- | ---------------------------------------------------- |
| `-target` | required      | Path to the JSON config file to parse                |
| `-output` | `schema.json` | Output path for the generated schema                 |
| `-mode`   | —             | `1` for generate, `2` for validate (v2, coming soon) |

# Example Input

# Example Input

```json
{
  "API_KEY": "abc123",
  "DB_HOST": "localhost",
  "DEBUG": "true",
  "MAX_RETRIES": "5",
  "PORT": "8080",
  "TIMEOUT_MS": "1500.5"
}
```

# Example Output

# Example Output

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "API_KEY": {
      "type": "string",
      "description": "Third-party API authentication key"
    },
    "DB_HOST": {
      "type": "string"
    },
    "DEBUG": {
      "type": "boolean"
    },
    "MAX_RETRIES": {
      "type": "integer"
    },
    "PORT": {
      "type": "integer"
    },
    "TIMEOUT_MS": {
      "type": "number"
    }
  },
  "required": ["PORT", "DEBUG", "DB_HOST"]
}
```

# Roadmap

## V0 — Core Tool

- [x] CLI reads a flat JSON config file
- [x] Type inference for `string`, `integer`, `float`, `boolean`
- [x] Localhost form UI to review and override inferred types
- [x] Mark fields as required, add descriptions
- [x] Generate a valid AppConfig JSON Schema file on disk
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

# Resources

- [AWS AppConfig — Creating a freeform configuration profile](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-free-form-configurations-creating.html)
- [AWS AppConfig — About validators](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-free-form-configuration-and-profile-create-console.html)
- [AWS AppConfig — Validators run at deployment time, not build time](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html)
- [AWS Blog — Best practices for validating AppConfig data](https://aws.amazon.com/blogs/mt/best-practices-for-validating-aws-appconfig-feature-flags-and-configuration-data)
- [JSON Schema Draft-04 specification](https://json-schema.org/specification-links#draft-4)
