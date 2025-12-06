UI scaffold served by `nginx` in container.

- Access after running compose at `http://localhost:8080`
- `index.html` queries `http://localhost:8000/health` for quick integration test

Next.js UI scaffold (Tailwind)

How to run locally (without Docker):

```bash
cd ui
npm install
npm run dev
```

Docker (built by root `docker-compose.yml`) exposes the app on port `3000` inside container; in compose we map it to `8080`.

Shadcn integration notes:
- This scaffold includes Tailwind and Next.js ready to accept `shadcn/ui` components.
- To add shadcn: run `npx shadcn-ui@latest init` inside `ui/` and follow its setup (it will update Tailwind configs and add component files).
- After adding shadcn components, add them under `components/` and import into pages.
