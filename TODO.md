# TODO

Stages of done (your TDD arc):

## Stage 1 - Parser exists and fails loudly

- Parse("") returns error
- Parse("{}") returns empty []Field, no error
- Parse with nested object returns error
- Parse with flat JSON returns fields with keys, no types yet

## Stage 2 — Inferencer runs

- Each Field has a type
- Table-driven: bool, int, float, string, null
- Edge case: the float64 trap from encoding/json

## Stage 3 — Schema emitter produces valid output

- Given []Field, emits JSON Schema draft-04 string
- Test: unmarshal output and assert structure

## Stage 4 — Server serves the form

- GET / renders fields as form
- POST /generate writes schema file
- This is where you stop unit testing and just run it

## Stage 5 (interview day) — Validator

- validate --schema out.json --config sample.json
- Prints pass/fail
