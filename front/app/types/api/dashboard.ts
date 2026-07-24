import type { LinkResource } from './link'

export type AnalyticsGranularity = 'hour' | 'day' | 'week' | 'month'
export type AnalyticsDimension = 'source' | 'referrer' | 'country' | 'device'
export type LinkStatus = 'active' | 'archived' | 'expired'

/** Shared query parameters for dashboard analytics endpoints. */
export type AnalyticsRangeQuery = {
	/** Inclusive ISO-8601 timestamp. */
	from: string
	/** Exclusive ISO-8601 timestamp. */
	to: string
	/** IANA time zone used for grouping, for example `Europe/Berlin`. */
	timezone?: string
	granularity?: AnalyticsGranularity
}

export type AnalyticsRange = Required<AnalyticsRangeQuery>

export type Metric = {
	value: number
	previousValue: number
	/** Percentage change from the immediately preceding equal-length period. */
	changePercent: number | null
}

export type DashboardSummary = {
	totalClicks: Metric
	uniqueClicks: Metric
	activeLinks: Metric
	/** Unique clicks divided by total clicks, expressed as a percentage. */
	uniqueClickRate: Metric
	/** Non-bot clicks divided by total clicks, expressed as a percentage. */
	humanTrafficRate: Metric
}

export type LinkDashboardSummary = Omit<DashboardSummary, 'activeLinks'> & {
	/** Optional until a conversion event model is implemented. */
	conversions?: Metric
	conversionRate?: Metric
}

export type TimeSeriesPoint = {
	/** Start of the bucket as an ISO-8601 timestamp. */
	timestamp: string
	clicks: number
	uniqueClicks: number
	humanClicks: number
	conversions?: number
}

export type BreakdownItem = {
	/** Stable machine value, for example `instagram`, `US`, or `mobile`. */
	key: string
	label: string
	clicks: number
	uniqueClicks: number
	/** Item clicks divided by all clicks in this breakdown, from 0 to 100. */
	sharePercent: number
}

export type LinkPerformance = {
	link: LinkResource
	campaign: string | null
	clicks: number
	uniqueClicks: number
	changePercent: number | null
}

export type UserDashboardResponse = {
	range: AnalyticsRange
	summary: DashboardSummary
	timeSeries: TimeSeriesPoint[]
	topSources: BreakdownItem[]
	topLinks: LinkPerformance[]
	deviceBreakdown: BreakdownItem[]
	countryBreakdown: BreakdownItem[]
}

export type LinkDashboardResponse = {
	range: AnalyticsRange
	link: LinkResource
	status: LinkStatus
	summary: LinkDashboardSummary
	timeSeries: TimeSeriesPoint[]
	sourceBreakdown: BreakdownItem[]
	deviceBreakdown: BreakdownItem[]
	countryBreakdown: BreakdownItem[]
	recentClicks: ClickActivity[]
}

export type UtmAttribution = {
	source: string | null
	medium: string | null
	campaign: string | null
	content: string | null
	term: string | null
}

export type ClickActivity = {
	id: string
	clickedAt: string
	linkId: number
	target: string
	referrer: string | null
	source: string
	countryCode: string | null
	country: string | null
	city: string | null
	device: string
	os: string
	browser: string
	isBot: boolean
	attribution: UtmAttribution
}

export type CursorPage<T> = {
	items: T[]
	page: {
		nextCursor: string | null
		hasMore: boolean
		limit: number
	}
}

export type LinkListItem = LinkResource & {
	status: LinkStatus
	clicks: number
	uniqueClicks: number
}

export type LinkListResponse = CursorPage<LinkListItem>
export type ClickActivityResponse = CursorPage<ClickActivity>

export type CreateLinkRequest = {
	name: string
	target: string
	linkType: string
	customSlug?: string
	campaign?: string
}

export type CreateLinkResponse = LinkResource

export type ApiError = {
	error: {
		code: string
		message: string
		details?: Record<string, string>
		requestId?: string
	}
}
