<template>
	<div class="space-y-6">
		<nav class="flex items-center gap-2 text-xs text-muted-foreground">
			<NuxtLink to="/links" class="hover:text-foreground">Links</NuxtLink>
			<ChevronRightIcon class="size-3.5" />
			<NuxtLink :to="`/links/${linkId}`" class="hover:text-foreground">{{ link.Name }}</NuxtLink>
			<ChevronRightIcon class="size-3.5" />
			<span class="text-foreground">Clicks</span>
		</nav>

		<header class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
			<div>
				<p class="mb-1 text-sm font-medium text-primary">Activity</p>
				<h1 class="text-2xl font-semibold tracking-tight sm:text-3xl">Click history</h1>
				<p class="mt-1.5 text-sm text-muted-foreground">Inspect recent traffic without exposing sensitive visitor data.</p>
			</div>
			<div class="flex items-center gap-2">
				<button type="button" class="flex h-10 items-center gap-2 rounded-lg border border-border bg-card px-3 text-sm font-medium">
					<CalendarDaysIcon class="size-4 text-muted-foreground" />
					Last 30 days
					<ChevronDownIcon class="size-4 text-muted-foreground" />
				</button>
				<button type="button" class="flex h-10 items-center gap-2 rounded-lg border border-border bg-card px-3 text-sm font-medium">
					<ArrowDownTrayIcon class="size-4 text-muted-foreground" />
					Export
				</button>
			</div>
		</header>

		<section class="grid gap-3 sm:grid-cols-3">
			<div class="rounded-xl border border-border bg-card p-4 shadow-sm sm:col-span-2">
				<div class="flex min-w-0 items-start gap-3">
					<div class="grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
						<LinkIcon class="size-5" />
					</div>
					<div class="min-w-0">
						<div class="flex items-center gap-2">
							<h2 class="truncate font-semibold">{{ link.Name }}</h2>
							<span class="size-2 shrink-0 rounded-full bg-success" />
						</div>
						<NuxtLink :to="`/links/${linkId}`" class="mt-0.5 block truncate text-sm font-medium text-primary">
							clickledger.io/{{ link.Slug }}
						</NuxtLink>
						<p class="mt-1 truncate text-xs text-muted-foreground">{{ link.Target }}</p>
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4 shadow-sm">
				<p class="text-xs font-medium text-muted-foreground">Visible activity</p>
				<div class="mt-2 flex items-end justify-between gap-3">
					<p class="text-2xl font-semibold tracking-tight">{{ filteredClicks.length.toLocaleString() }}</p>
					<p class="mb-1 text-xs font-semibold text-success">{{ humanTrafficRate }}% human</p>
				</div>
				<div class="mt-3 h-1.5 overflow-hidden rounded-full bg-muted">
					<div class="h-full rounded-full bg-success" :style="{ width: `${humanTrafficRate}%` }" />
				</div>
			</div>
		</section>

		<section class="overflow-hidden rounded-xl border border-border bg-card shadow-sm">
			<div class="border-b border-border p-3 sm:p-4">
				<div class="flex flex-col gap-3 xl:flex-row xl:items-center xl:justify-between">
					<label class="relative block w-full xl:max-w-sm">
						<MagnifyingGlassIcon class="absolute left-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground" />
						<input
							v-model="search"
							type="search"
							placeholder="Search location, source, campaign..."
							class="h-10 w-full rounded-lg border border-border bg-background pl-9 pr-3 text-sm outline-none focus:ring-2 focus:ring-ring"
						>
					</label>
					<div class="flex flex-wrap items-center gap-2">
						<label class="sr-only" for="source-filter">Source</label>
						<select id="source-filter" v-model="sourceFilter" class="h-10 rounded-lg border border-border bg-background px-3 text-sm font-medium outline-none focus:ring-2 focus:ring-ring">
							<option value="">All sources</option>
							<option v-for="source in sourceOptions" :key="source" :value="source">{{ source }}</option>
						</select>
						<label class="sr-only" for="device-filter">Device</label>
						<select id="device-filter" v-model="deviceFilter" class="h-10 rounded-lg border border-border bg-background px-3 text-sm font-medium outline-none focus:ring-2 focus:ring-ring">
							<option value="">All devices</option>
							<option v-for="device in deviceOptions" :key="device" :value="device">{{ device }}</option>
						</select>
						<label class="flex h-10 cursor-pointer items-center gap-2 rounded-lg border border-border px-3 text-sm font-medium">
							<input v-model="includeBots" type="checkbox" class="size-4 accent-primary">
							Include bots
						</label>
					</div>
				</div>
			</div>

			<div v-if="loading" class="space-y-3 p-5">
				<div v-for="index in 8" :key="index" class="h-12 animate-pulse rounded-lg bg-muted" />
			</div>

			<div v-else-if="filteredClicks.length" class="overflow-x-auto">
				<table class="w-full min-w-[880px] text-left text-sm">
					<thead class="border-b border-border bg-muted/40 text-[10px] uppercase tracking-[0.12em] text-muted-foreground">
						<tr>
							<th class="px-5 py-3 font-semibold">Clicked at</th>
							<th class="px-4 py-3 font-semibold">Location</th>
							<th class="px-4 py-3 font-semibold">Source</th>
							<th class="px-4 py-3 font-semibold">Device</th>
							<th class="px-4 py-3 font-semibold">Campaign</th>
							<th class="px-5 py-3 text-right font-semibold">Traffic</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						<tr v-for="click in filteredClicks" :key="click.id" class="hover:bg-muted/30">
							<td class="px-5 py-3.5">
								<p class="font-medium">{{ formatTime(click.clickedAt) }}</p>
								<p class="mt-0.5 text-xs text-muted-foreground">{{ formatDate(click.clickedAt) }}</p>
							</td>
							<td class="px-4 py-3.5">
								<div class="flex items-center gap-2.5">
									<div class="grid size-8 place-items-center rounded-full bg-muted text-[10px] font-semibold text-muted-foreground">
										{{ click.countryCode || '—' }}
									</div>
									<div>
										<p class="font-medium">{{ click.city || 'Unknown city' }}</p>
										<p class="text-xs text-muted-foreground">{{ click.country || 'Unknown country' }}</p>
									</div>
								</div>
							</td>
							<td class="px-4 py-3.5">
								<p class="font-medium">{{ click.source }}</p>
								<p class="max-w-40 truncate text-xs text-muted-foreground">{{ referrerLabel(click.referrer) }}</p>
							</td>
							<td class="px-4 py-3.5">
								<p class="font-medium capitalize">{{ click.device || 'Unknown' }}</p>
								<p class="text-xs text-muted-foreground">{{ [click.os, click.browser].filter(Boolean).join(' · ') || 'Unknown client' }}</p>
							</td>
							<td class="px-4 py-3.5">
								<span v-if="click.attribution.campaign" class="rounded-md bg-primary/10 px-2 py-1 text-xs font-medium text-primary">
									{{ click.attribution.campaign }}
								</span>
								<span v-else class="text-xs text-muted-foreground">Unattributed</span>
							</td>
							<td class="px-5 py-3.5 text-right">
								<span
									class="inline-flex items-center gap-1.5 rounded-full px-2 py-1 text-xs font-medium"
									:class="click.isBot ? 'bg-warning/10 text-warning' : 'bg-success/10 text-success'"
								>
									<span class="size-1.5 rounded-full bg-current" />
									{{ click.isBot ? 'Bot' : 'Human' }}
								</span>
							</td>
						</tr>
					</tbody>
				</table>
			</div>

			<div v-else class="flex flex-col items-center px-5 py-16 text-center">
				<div class="grid size-11 place-items-center rounded-full bg-muted text-muted-foreground">
					<CursorArrowRaysIcon class="size-5" />
				</div>
				<h2 class="mt-3 text-sm font-semibold">No clicks found</h2>
				<p class="mt-1 max-w-sm text-xs text-muted-foreground">Try changing the filters or wait for new activity on this link.</p>
				<button type="button" class="mt-4 text-xs font-semibold text-primary" @click="clearFilters">Clear filters</button>
			</div>

			<footer v-if="!loading && filteredClicks.length" class="flex items-center justify-between border-t border-border px-4 py-3 sm:px-5">
				<p class="text-xs text-muted-foreground">Showing {{ filteredClicks.length }} click{{ filteredClicks.length === 1 ? '' : 's' }}</p>
				<button
					type="button"
					class="rounded-lg border border-border px-3 py-2 text-xs font-semibold disabled:cursor-not-allowed disabled:opacity-50"
					:disabled="!nextCursor || loadingMore"
					@click="loadMore"
				>
					{{ loadingMore ? 'Loading…' : 'Load more' }}
				</button>
			</footer>
		</section>
	</div>
</template>

<script setup lang="ts">
import {
	ArrowDownTrayIcon,
	CalendarDaysIcon,
	ChevronDownIcon,
	ChevronRightIcon,
	CursorArrowRaysIcon,
	LinkIcon,
	MagnifyingGlassIcon,
} from '@heroicons/vue/24/outline'
import type { Click } from '~/types/api/click'
import type { ClickActivity, ClickActivityResponse } from '~/types/api/dashboard'
import type { Link } from '~/types/api/link'

const route = useRoute()
const linkId = computed(() => String(route.params.id))
const loading = ref(true)
const search = ref('')
const sourceFilter = ref('')
const deviceFilter = ref('')
const includeBots = ref(false)
const nextCursor = ref<string | null>(null)
const loadingMore = ref(false)

const link = ref<Link>({
	ID: Number(route.params.id) || 1,
	Name: 'Summer campaign',
	Slug: 'summer-26',
	Target: 'https://example.com/summer-sale',
	LinkType: 'campaign',
})

const fallbackClicks: ClickActivity[] = [
	activity('1', 2, 'Berlin', 'Germany', 'DE', 'Instagram', 'instagram.com', 'mobile', 'iOS', 'Safari', 'summer_26'),
	activity('2', 8, 'Austin', 'United States', 'US', 'Google', 'google.com', 'desktop', 'Windows', 'Chrome', 'summer_26'),
	activity('3', 14, 'London', 'United Kingdom', 'GB', 'Direct', null, 'mobile', 'Android', 'Chrome', 'newsletter'),
	activity('4', 21, 'Toronto', 'Canada', 'CA', 'LinkedIn', 'linkedin.com', 'desktop', 'macOS', 'Safari', 'summer_26'),
	activity('5', 32, 'Munich', 'Germany', 'DE', 'Instagram', 'instagram.com', 'mobile', 'iOS', 'Instagram', 'summer_26'),
	activity('6', 47, 'New York', 'United States', 'US', 'Google', 'google.com', 'desktop', 'macOS', 'Chrome', null),
	activity('7', 63, null, null, null, 'Direct', null, 'bot', 'Linux', 'Crawler', null, true),
	activity('8', 76, 'Manchester', 'United Kingdom', 'GB', 'LinkedIn', 'linkedin.com', 'mobile', 'Android', 'Chrome', 'newsletter'),
]

const clicks = ref<ClickActivity[]>(fallbackClicks)

onMounted(async () => {
	try {
		const [linkResponse, clicksResponse] = await Promise.all([
			fetch(`http://localhost:8888/api/link/${linkId.value}`),
			fetch(`http://localhost:8888/api/link/${linkId.value}/clicks?limit=25`),
		])

		if (linkResponse.ok) link.value = await linkResponse.json()
		if (clicksResponse.ok) {
			const payload = await clicksResponse.json() as Click[] | ClickActivityResponse
			if (Array.isArray(payload)) {
				clicks.value = payload.map(normalizeLegacyClick)
			} else {
				clicks.value = payload.items
				nextCursor.value = payload.page.nextCursor
			}
		}
	} catch {
		// Stable prototype activity remains visible when the local API is offline.
	} finally {
		loading.value = false
	}
})

const sourceOptions = computed(() => [...new Set(clicks.value.map(click => click.source))].sort())
const deviceOptions = computed(() => [...new Set(clicks.value.map(click => click.device).filter(Boolean))].sort())
const filteredClicks = computed(() => {
	const query = search.value.trim().toLowerCase()
	return clicks.value.filter((click) => {
		if (!includeBots.value && click.isBot) return false
		if (sourceFilter.value && click.source !== sourceFilter.value) return false
		if (deviceFilter.value && click.device !== deviceFilter.value) return false
		if (!query) return true
		return [
			click.city,
			click.country,
			click.source,
			click.referrer,
			click.attribution.campaign,
			click.device,
			click.browser,
		].some(value => value?.toLowerCase().includes(query))
	})
})

const humanTrafficRate = computed(() => {
	if (!clicks.value.length) return 0
	return Math.round((clicks.value.filter(click => !click.isBot).length / clicks.value.length) * 100)
})

async function loadMore() {
	if (!nextCursor.value || loadingMore.value) return
	loadingMore.value = true
	try {
		const response = await fetch(`http://localhost:8888/api/link/${linkId.value}/clicks?limit=25&cursor=${encodeURIComponent(nextCursor.value)}`)
		if (!response.ok) return
		const payload = await response.json() as ClickActivityResponse
		const knownIds = new Set(clicks.value.map(click => click.id))
		clicks.value.push(...payload.items.filter(click => !knownIds.has(click.id)))
		nextCursor.value = payload.page.nextCursor
	} finally {
		loadingMore.value = false
	}
}

function activity(
	id: string,
	minutesAgo: number,
	city: string | null,
	country: string | null,
	countryCode: string | null,
	source: string,
	referrer: string | null,
	device: string,
	os: string,
	browser: string,
	campaign: string | null,
	isBot = false,
): ClickActivity {
	return {
		id,
		clickedAt: new Date(Date.UTC(2026, 6, 23, 12, 0) - minutesAgo * 60_000).toISOString(),
		linkId: Number(route.params.id) || 1,
		target: link.value.Target,
		referrer,
		source,
		countryCode,
		country,
		city,
		device,
		os,
		browser,
		isBot,
		attribution: { source: source.toLowerCase(), medium: null, campaign, content: null, term: null },
	}
}

function normalizeLegacyClick(click: Click): ClickActivity {
	const source = click.UtmContent?.utm_source || sourceFromReferrer(click.Referer)
	return {
		id: String(click.ID),
		clickedAt: click.CreatedAt,
		linkId: click.LinkID,
		target: click.Target,
		referrer: click.Referer || null,
		source,
		countryCode: click.CountryCode || null,
		country: click.Country || null,
		city: click.City || null,
		device: click.DeviceType || 'unknown',
		os: click.OS || '',
		browser: click.Browser || '',
		isBot: click.IsBot,
		attribution: {
			source: click.UtmContent?.utm_source || null,
			medium: click.UtmContent?.utm_medium || null,
			campaign: click.UtmContent?.utm_campaign || null,
			content: click.UtmContent?.utm_content || null,
			term: click.UtmContent?.utm_term || null,
		},
	}
}

function sourceFromReferrer(referrer?: string) {
	if (!referrer) return 'Direct'
	try {
		return new URL(referrer).hostname.replace(/^www\./, '')
	} catch {
		return referrer
	}
}

function referrerLabel(referrer: string | null) {
	return referrer ? sourceFromReferrer(referrer) : 'No referrer'
}

function formatTime(value: string) {
	return new Intl.DateTimeFormat('en', { hour: '2-digit', minute: '2-digit' }).format(new Date(value))
}

function formatDate(value: string) {
	return new Intl.DateTimeFormat('en', { day: '2-digit', month: 'short', year: 'numeric' }).format(new Date(value))
}

function clearFilters() {
	search.value = ''
	sourceFilter.value = ''
	deviceFilter.value = ''
	includeBots.value = false
}
</script>
