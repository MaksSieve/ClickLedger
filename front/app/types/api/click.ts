import type { Link } from "./link"

export type LegacyUtmContent = {
	utm_source?: string
	utm_medium?: string
	utm_campaign?: string
	utm_content?: string
	utm_term?: string
}

/** Legacy shape currently serialized directly from the Go Click model. */
export type Click = {
	ID: string | number
	CreatedAt: string
	LinkID: number
	Link: Link
	Target: string

	UserIP: string
	UserAgentRaw?: string
	Referer?: string
	UtmContent?: LegacyUtmContent
	Country?: string
	CountryCode?: string
	City?: string
	Region?: string
	DeviceType?: string
	OS?: string
	Browser?: string
	IsBot: boolean
	RedirectStatus?: string
}
