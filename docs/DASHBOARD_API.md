# Dashboard API contracts

These endpoints provide the data used by the Phase 2 user and link dashboards.
New responses use camelCase JSON. Dates are ISO-8601 UTC timestamps and percentages
are numbers from `0` to `100`.

The canonical frontend types live in
`front/app/types/api/dashboard.ts`.

## Required endpoints

| Method | Endpoint | Used by |
| --- | --- | --- |
| `GET` | `/api/dashboard` | User overview: summary, trend, sources, audience, and top links |
| `GET` | `/api/link` | Searchable, paginated links directory |
| `POST` | `/api/link/shorten` | Create-link action |
| `GET` | `/api/link/{id}/dashboard` | Complete analytics view for one link |
| `GET` | `/api/link/{id}/clicks` | Paginated click activity and the full clicks view |

The aggregate dashboard endpoints intentionally return all cards and reports for
their view in one response. Every report shares the same time range, so splitting
them into many requests would add latency and consistency problems without adding
useful flexibility.

## Shared analytics query

Both dashboard endpoints accept:

| Parameter | Required | Format | Meaning |
| --- | --- | --- | --- |
| `from` | yes | ISO-8601 timestamp | Inclusive range start |
| `to` | yes | ISO-8601 timestamp | Exclusive range end |
| `timezone` | no | IANA time zone | Bucket boundaries; defaults to `UTC` |
| `granularity` | no | `hour`, `day`, `week`, `month` | Defaults according to range length |

Example:

```http
GET /api/dashboard?from=2026-06-24T00:00:00Z&to=2026-07-24T00:00:00Z&timezone=Europe%2FBerlin&granularity=day
```

The server calculates `previousValue` and `changePercent` against the immediately
preceding range of the same duration. `changePercent` is `null` when the previous
value is zero.

## `GET /api/dashboard`

Returns the complete user-level dashboard.

```json
{
  "range": {
    "from": "2026-06-24T00:00:00Z",
    "to": "2026-07-24T00:00:00Z",
    "timezone": "Europe/Berlin",
    "granularity": "day"
  },
  "summary": {
    "totalClicks": { "value": 12847, "previousValue": 10869, "changePercent": 18.2 },
    "uniqueClicks": { "value": 9621, "previousValue": 8552, "changePercent": 12.5 },
    "activeLinks": { "value": 6, "previousValue": 4, "changePercent": 50 },
    "uniqueClickRate": { "value": 74.9, "previousValue": 72.4, "changePercent": 3.45 },
    "humanTrafficRate": { "value": 96.8, "previousValue": 95.7, "changePercent": 1.15 }
  },
  "timeSeries": [
    {
      "timestamp": "2026-06-24T00:00:00Z",
      "clicks": 280,
      "uniqueClicks": 221,
      "humanClicks": 270
    }
  ],
  "topSources": [
    {
      "key": "instagram",
      "label": "Instagram",
      "clicks": 4280,
      "uniqueClicks": 3156,
      "sharePercent": 33.3
    }
  ],
  "topLinks": [
    {
      "link": {
        "id": 1,
        "name": "Summer campaign",
        "slug": "summer-26",
        "target": "https://example.com/summer",
        "linkType": "campaign",
        "createdAt": "2026-06-20T14:32:00Z"
      },
      "campaign": "Summer 2026",
      "clicks": 4862,
      "uniqueClicks": 3714,
      "changePercent": 24.3
    }
  ],
  "deviceBreakdown": [],
  "countryBreakdown": []
}
```

`topSources`, `topLinks`, `deviceBreakdown`, and `countryBreakdown` should contain
at most five items. Missing metadata is grouped under a stable `unknown` key.

## `GET /api/link`

Extends the existing link collection with search, status, analytics totals, and
cursor pagination.

| Parameter | Required | Meaning |
| --- | --- | --- |
| `search` | no | Case-insensitive name, slug, or target match |
| `status` | no | `active`, `archived`, or `expired` |
| `sort` | no | `createdAt`, `clicks`, or `name`; default `createdAt` |
| `order` | no | `asc` or `desc`; default `desc` |
| `cursor` | no | Opaque cursor returned by the previous page |
| `limit` | no | `1`–`100`; default `25` |
| `from`, `to`, `timezone` | no | Range used for click totals |

Response type: `LinkListResponse`.

## `POST /api/link/shorten`

Request type: `CreateLinkRequest`. Successful response is a
`CreateLinkResponse` with HTTP status `201 Created`.

```json
{
  "name": "Summer campaign",
  "target": "https://example.com/summer",
  "linkType": "campaign",
  "customSlug": "summer-26",
  "campaign": "Summer 2026"
}
```

`customSlug` and `campaign` are optional. The server returns `409 Conflict` when
the requested slug is already in use.

## `GET /api/link/{id}/dashboard`

Accepts the shared analytics query and returns `LinkDashboardResponse`. It has
the same time-series and breakdown shapes as the user endpoint, plus:

- link identity and status;
- link-level summary metrics;
- source, device, and country breakdowns;
- the five most recent click activities.

`conversions` and `conversionRate` are omitted until conversion events are
implemented. They are optional in the frontend contract for that reason.

## `GET /api/link/{id}/clicks`

| Parameter | Required | Meaning |
| --- | --- | --- |
| `cursor` | no | Opaque cursor returned by the previous page |
| `limit` | no | `1`–`100`; default `25` |
| `from`, `to` | no | Click timestamp filter |
| `source` | no | Exact normalized source |
| `country` | no | ISO 3166-1 alpha-2 country code |
| `device` | no | Normalized device type |
| `includeBots` | no | Boolean; default `false` |

Response type: `ClickActivityResponse`. Results are ordered by `clickedAt`
descending and then by `id` descending, which keeps cursor pagination stable.
Raw IP addresses and complete user-agent strings must not be returned to the
dashboard.

## Error response

All failures use the `ApiError` contract:

```json
{
  "error": {
    "code": "invalid_time_range",
    "message": "`from` must be earlier than `to`.",
    "details": { "from": "Must be earlier than `to`." },
    "requestId": "req_01J..."
  }
}
```

Expected status codes are `400` for invalid input, `404` for a missing link,
`409` for a duplicate custom slug, and `500` for an unexpected server error.

## Existing endpoint migration

- Keep `GET /api/link/{id}` for lightweight link metadata.
- Replace the current referrer-only `GET /api/link/{id}/analytics` with
  `GET /api/link/{id}/dashboard`.
- Change `GET /api/link/{id}/clicks` from an unbounded array to
  `ClickActivityResponse`.
- Migrate legacy Go model JSON (`ID`, `Name`, `Slug`, and similar fields) to
  explicit DTOs over time. The dashboard contracts otherwise use camelCase.
