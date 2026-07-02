# Niche URL Shortener + Analytics - Development Plan

**Project**: Niche-focused URL shortener with deep analytics, targeted at **Affiliate & Performance Marketers**.

**Primary Goal**: Build a monetizable MVP while transitioning to production-grade Go backend. Leverage your 9+ years experience, TypeScript knowledge, and Performance QA background.

**Niche Value Prop**: Affordable, focused analytics for campaign tracking, attribution, and ROI (better than generic Bitly for this audience). Features like campaign tagging, UTM handling, source breakdowns, and actionable dashboards.

**Success Metrics (MVP)**: Working shorten/redirect + click tracking, basic user dashboard, deployed, and 10-50 early users or waitlist signups.

---

## Tech Stack (Recommended)

- **Backend**: Go (focus on learning concurrency, clean architecture, performance)
  - Router: Gin or Echo (simple, fast)
  - ORM/Query: GORM or sqlx + raw for analytics queries
  - Auth: JWT middleware
- **Database**: PostgreSQL (rich analytics queries, JSON for metadata) + Redis (caching, rate limiting, sessions)
- **Frontend/Dashboard**: Your TypeScript strengths — Vite + React (or simpler HTMX + Go templates/Templ for faster backend focus)
- **Analytics**: Start with Postgres aggregations + materialized views. Later: background jobs or lightweight ClickHouse if volume grows.
- **Other**: Docker for consistency, Stripe (later for payments), basic logging/monitoring.
- **Deployment**: Docker + Fly.io or Railway (cheap, easy for pet projects). Add CI later.

**Why this stack?** Go excels at high-throughput redirects + concurrent analytics ingestion. Your QA background helps optimize performance and data pipelines.

---

## High-Level Architecture (Conceptual)

- **Layers**: HTTP Handlers → Services (business logic) → Repositories (data access)
- **Core Flows**:
  - Shorten: Validate URL → Generate slug (base62 or custom) → Store with metadata (user, campaign tags, UTM)
  - Redirect: Lookup slug → Log click event (async-friendly: IP, user-agent, referrer, geo via simple lookup, timestamp) → 301 redirect
  - Analytics: Aggregate clicks by various dimensions (time, source, geo, device, campaign)
- **Data Model Ideas** (plan only):
  - urls table: id, slug, original_url, user_id, custom_domain?, campaign_tags (array/JSON), created_at, expires?
  - clicks table: id, url_id, clicked_at, ip, user_agent, referrer, country?, device?, source?
  - users table: id, email, plan, etc.
- **Key Concerns**: Rate limiting per user/IP, spam/abuse detection (simple heuristics first), privacy (GDPR-friendly logging), performance (cache hot slugs in Redis).

**Your QA Strength**: Design for observability (logs, metrics on redirect latency, error rates) and test edge cases early (invalid URLs, collisions, high load simulation).

---

## Development Phases (Incremental — One Piece at a Time)

**Phase 0: Project Setup & Go Foundations (3-7 days)**
- Initialize Go module, clean folder structure (cmd/api, internal/{handler,service,repository,model}, pkg, migrations).
- Set up Postgres + Redis locally (Docker Compose).
- **Learning Focus**: Go modules/packages, error handling, basic HTTP server, database connections, middleware pattern, concurrency basics (goroutines/channels for async tasks).
- Resources: Official Go Tour + "Let's Go" book (web apps section) or similar focused on backend. Practice small standalone exercises before integrating.
- Deliverable: Hello-world API with DB connection + simple health endpoint.
- Save progress: Commit to git, update this plan with notes.

**Phase 1: Core URL Shortening + Redirect Engine (1-2 weeks)**
- Implement shorten endpoint (POST /shorten) and redirect (GET /{slug}).
- Slug generation: Random or custom (handle collisions simply).
- Basic validation & sanitization.
- **Learning Focus**: HTTP routing/middleware, clean service layer, repository pattern, basic error responses, testing (unit for logic).
- **Niche Tie-in**: Support initial metadata like tags or UTM params.
- Deliverable: Functional shorten → redirect loop. Test manually + simple load simulation (your QA skills).
- Milestone: M1 complete.

**Phase 2: Click Tracking + Basic Analytics (2-3 weeks)**
- Log every redirect as a click event (store metadata).
- Build simple analytics endpoints/queries: total clicks, by time (daily/weekly), top referrers/sources, basic geo/device breakdown.
- Dashboard views (even if static JSON first).
- **Learning Focus**: Efficient DB schema & queries (GROUP BY, indexes), time-series thinking, Redis for hot data/caching, background processing (goroutines or simple queue for non-blocking logs).
- **Niche Tie-in**: Campaign tagging support, basic source attribution.
- Deliverable: Click logging + 3-5 key analytics reports. Dashboard prototype (TS frontend or simple HTML).
- Milestone: M2 complete. Users can see "what's working" for their links.

**Phase 3: User Management, Auth & Multi-User (1-2 weeks)**
- User registration/login (JWT), link ownership.
- My Links list + personal analytics.
- Basic rate limiting & abuse protection.
- **Learning Focus**: Auth middleware/JWT, user context, authorization (own links only), sessions or stateless.
- Deliverable: Auth-protected flows. Users can manage their own short links.
- Milestone: M3 complete.

**Phase 4: Niche Polish & Monetization Prep (2+ weeks)**
- Advanced analytics: Campaign grouping/ROI views, UTM builder helper, time-based trends, export (CSV).
- Custom slugs/domains (stretch).
- Simple landing page or waitlist for early access.
- **Learning Focus**: More complex queries/aggregations, API design for frontend, background jobs, basic observability.
- **Monetization Hooks**: Plan tiers in DB (free/pro), usage limits, Stripe integration placeholder.
- Deliverable: Polished niche analytics that solves affiliate pain points. Deployed version.
- Milestone: MVP ready for beta users.

**Phase 5: Deployment, Hardening & Growth (Ongoing)**
- Dockerize, deploy to Fly.io/Railway.
- Add basic monitoring, logging, backups.
- Security: HTTPS, input validation, rate limits, privacy considerations.
- **Learning Focus**: Production concerns (config, secrets, graceful shutdown, performance tuning).
- Later ideas: API for power users, webhooks, simple AI insights, team features.

---

## Recommended Timeline (Part-time / Evenings)

- Weeks 1-2: Phase 0 + 1 (core engine)
- Weeks 3-5: Phase 2 (analytics — core value)
- Weeks 6-7: Phase 3 (users)
- Weeks 8-10: Phase 4 (niche + deploy)
- Ongoing: Iterate based on feedback, add monetization.

Adjust based on your pace. Aim for **working MVP in 6-8 weeks** with consistent effort. Track weekly progress in this plan or a simple TODO.

---

## Feature Prioritization (MVP Focus for Affiliate Niche)

**Must-Have (MVP)**:
- Shorten with custom/random slug + metadata (tags/campaign)
- Reliable redirect + click logging
- Basic dashboard: clicks over time, by source/referrer, simple breakdowns
- User accounts & own links
- Rate limiting & basic spam protection

**Nice-to-Have (Post-MVP)**:
- UTM parameter helper/builder
- Advanced filters & exports
- Custom domains
- QR codes
- Team/collaboration
- Integrations (Google Analytics, affiliate networks)
- AI: "best time to share" or performance predictions

Start narrow — analytics depth > breadth.

---

## Monetization Path (Pet Project Revenue)

- **Freemium**: Free tier (limited links or basic analytics) → Pro ($15-35/mo): unlimited + advanced reports + custom domains.
- **Business tier**: Teams, API access, priority support, white-label later.
- **Early Validation**: Build simple landing page (TS) with waitlist. Offer lifetime deals or founder pricing to first users.
- **Marketing**: Share progress on X/Reddit (r/affiliatemarketing, r/indiehackers). Focus on "better tracking for affiliates without Bitly prices".
- Target: First paid users within 3-6 months post-MVP.

---

## Risks & Mitigations (Leverage Your QA Background)

- **Performance/Scalability**: Profile redirects & queries early. Use indexes, caching, async logging.
- **Data Privacy/GDPR**: Minimal logging, anonymize where possible, clear policies. Big differentiator.
- **Abuse/Spam**: Strong rate limits, moderation queue, domain reputation.
- **Analytics Accuracy**: Handle bots, multiple clicks, geo IP limitations (simple MaxMind or free DB first).
- **Scope Creep**: Strict MVP definition. Add one niche feature at a time.
- **Learning Curve**: Dedicate focused time to Go concepts before big features. Build small prototypes.

---

## Learning Resources (Targeted for Go Backend Transition)

- **Core Go + Web**: "Let's Go" book (web applications focus), official Go by Example, Tour of Go.
- **Advanced**: "100 Go Mistakes" (avoid common pitfalls), concurrency patterns.
- **DB/Analytics**: Postgres docs + GORM guides. Practice query optimization.
- **Practice Strategy**: Before each phase, do 1-2 small standalone Go exercises matching the concepts (e.g., middleware, DB repo). Then integrate.
- **Your Advantage**: Apply Performance QA mindset — measure, profile, test edges from day one.
- Track what you learn in a separate notes file.

---

## Project Organization & Progress Tracking

- Repo structure: Follow standard Go layout + docs/ folder.
- This plan lives at: `DEVELOPMENT_PLAN.md` (update it weekly with status, learnings, decisions).
- Use git branches: main (stable), feature/phase-x.
- Simple kanban: GitHub Projects or linear.app free tier for tasks.
- Weekly ritual: Review this plan, update milestones, note blockers/learnings.

**Next Immediate Step**: 
1. Set up the repo locally following Phase 0.
2. Reply here with any questions on Phase 0 setup or specific learning resources.
3. We can refine features or dive deeper into any phase.

This plan is saved in the project sources. Build incrementally, learn deliberately, and ship value to the niche early. You've got the experience — this will accelerate your backend transition.

Good luck — start with Phase 0 setup when ready!