# CV (HTML/CSS)

Editable version of Ethan Donovan’s CV, matching the Canva PDF layout.

## Quick start

Open `index.html` in a browser (double-click or use a local server). Print to PDF via **File → Print** (enable “Background graphics” for the navy header).

## What to edit

| Change | Where |
|--------|--------|
| Name, title, summary, jobs, education | `index.html` — search for the section comments |
| Photo | Replace `assets/profile.png` or change the `src` on `.cv-header__photo` |
| Colors, fonts, page width | `styles.css` — `:root` variables at the top (`--cv-width`, default 860px) |
| Section icons | `index.html` — inline `<svg class="cv-icon">` in each section header |
| Infrastructure projects | `index.html` — section below summary (`#infra-heading`) |
| Freelancing (Project X) | `index.html` — section below infrastructure (`#freelance-heading`) |
| LinkedIn / GitHub URLs | `index.html` — `.cv-contact` links |

## Add a new job

Copy a `<article class="cv-job">...</article>` block inside `.cv-timeline` and update role, dates, org, and bullets.

## Add a new section

Copy any `<section class="cv-section">` block (icon SVG + title + `<hr>` + content).

## Files

- `index.html` — all text content
- `styles.css` — layout and theme
- `assets/profile.png` — headshot (extracted from your PDF)
