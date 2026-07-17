import type { Link } from "./link"

export type Click = {
	ID: string
	CreatedAt: string
	LinkID: number
	Link: Link
	Target: string
	
	UserIP: string
}