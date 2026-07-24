/**
 * Legacy shape currently serialized directly from the Go model.
 * New endpoints should use LinkResource instead.
 */
export type Link = {
	ID: number
	Target: string
	Slug: string
	Name: string
	LinkType: string
}

/** Canonical public link DTO for new API endpoints. */
export type LinkResource = {
	id: number
	target: string
	slug: string
	name: string
	linkType: string
	createdAt: string
}
