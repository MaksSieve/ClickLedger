<template>
	<div class="space-y-6">
		<header class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
			<div>
				<p class="mb-1 text-sm font-medium text-primary">Overview</p>
				<h1 class="text-2xl font-semibold tracking-tight sm:text-3xl">Good morning, Maksim</h1>
				<p class="mt-1.5 text-sm text-muted-foreground">Here’s how your links performed in the last 30 days.</p>
			</div>
			<div class="flex items-center gap-2">
				<button type="button" class="flex h-10 items-center gap-2 rounded-lg border border-border bg-card px-3 text-sm font-medium">
					<CalendarDaysIcon class="size-4 text-muted-foreground" />
					Last 30 days
					<ChevronDownIcon class="size-4 text-muted-foreground" />
				</button>
				<NuxtLink to="/links" class="flex h-10 items-center gap-2 rounded-lg bg-primary px-4 text-sm font-semibold text-primary-foreground shadow-sm">
					<PlusIcon class="size-4" />
					New link
				</NuxtLink>
			</div>
		</header>

		<section class="grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
			<DashboardStatCard label="Total clicks" value="12,847" change="18.2%">
				<CursorArrowRaysIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Unique visitors" value="9,621" change="12.5%">
				<UserGroupIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Active links" :value="String(displayLinks.length)" change="2" caption="new this month">
				<LinkIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Average click rate" value="8.4%" change="1.2%">
				<ChartBarIcon class="size-5" />
			</DashboardStatCard>
		</section>

		<section class="grid gap-4 xl:grid-cols-[minmax(0,1.75fr)_minmax(18rem,0.8fr)]">
			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<div class="mb-6 flex items-start justify-between">
						<div>
							<h2 class="font-semibold">Clicks over time</h2>
							<p class="mt-1 text-xs text-muted-foreground">Total and unique clicks across all links</p>
						</div>
						<div class="flex items-center gap-4 text-xs text-muted-foreground">
							<span class="flex items-center gap-1.5"><i class="size-2 rounded-full bg-primary" /> Clicks</span>
						</div>
					</div>
					<DashboardTrendChart :points="trendPoints" :labels="trendLabels" aria-label="Account clicks over the last 30 days" />
				</div>
			</UIPanel>

			<UIPanel titleless borderless>
				<div class="flex h-full flex-col p-4 sm:p-5">
					<div class="mb-5">
						<h2 class="font-semibold">Top sources</h2>
						<p class="mt-1 text-xs text-muted-foreground">Where your traffic comes from</p>
					</div>
					<DashboardBreakdownList :items="sourceItems" />
					<button type="button" class="mt-auto pt-6 text-left text-xs font-semibold text-primary">View full report →</button>
				</div>
			</UIPanel>
		</section>

		<section class="grid gap-4 xl:grid-cols-[minmax(0,1.75fr)_minmax(18rem,0.8fr)]">
			<UIPanel titleless borderless class="overflow-hidden">
				<div class="flex items-center justify-between border-b border-border px-4 py-4 sm:px-5">
					<div>
						<h2 class="font-semibold">Top performing links</h2>
						<p class="mt-1 text-xs text-muted-foreground">Ranked by total clicks</p>
					</div>
					<NuxtLink to="/links" class="text-xs font-semibold text-primary">View all</NuxtLink>
				</div>
				<div class="overflow-x-auto">
					<table class="w-full min-w-[620px] text-left">
						<thead class="border-b border-border bg-muted/40 text-[10px] uppercase tracking-[0.12em] text-muted-foreground">
							<tr>
								<th class="px-5 py-2.5 font-semibold">Link</th>
								<th class="px-4 py-2.5 font-semibold">Campaign</th>
								<th class="px-4 py-2.5 text-right font-semibold">Clicks</th>
								<th class="px-5 py-2.5 text-right font-semibold">Trend</th>
							</tr>
						</thead>
						<tbody class="divide-y divide-border">
							<tr v-for="link in displayLinks.slice(0, 4)" :key="link.ID" class="group">
								<td class="px-5 py-3.5">
									<NuxtLink :to="`/links/${link.ID}`" class="block font-medium group-hover:text-primary">{{ link.Name }}</NuxtLink>
									<span class="text-xs text-muted-foreground">clickledger.io/{{ link.Slug }}</span>
								</td>
								<td class="px-4 py-3.5"><span class="rounded-md bg-muted px-2 py-1 text-xs text-muted-foreground">{{ campaignFor(link.ID) }}</span></td>
								<td class="px-4 py-3.5 text-right text-sm font-semibold">{{ clicksFor(link.ID).toLocaleString() }}</td>
								<td class="px-5 py-3.5 text-right text-xs font-semibold text-success">↗ {{ trendFor(link.ID) }}%</td>
							</tr>
						</tbody>
					</table>
				</div>
			</UIPanel>

			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<h2 class="font-semibold">Audience snapshot</h2>
					<p class="mt-1 text-xs text-muted-foreground">Top device and location</p>
					<div class="mt-5 grid grid-cols-2 gap-3">
						<div class="rounded-lg bg-muted/60 p-3">
							<DevicePhoneMobileIcon class="mb-3 size-5 text-primary" />
							<p class="text-lg font-semibold">64%</p>
							<p class="text-xs text-muted-foreground">Mobile</p>
						</div>
						<div class="rounded-lg bg-muted/60 p-3">
							<GlobeAltIcon class="mb-3 size-5 text-primary" />
							<p class="text-lg font-semibold">US</p>
							<p class="text-xs text-muted-foreground">Top country</p>
						</div>
					</div>
					<div class="mt-4 rounded-lg border border-border p-3">
						<div class="flex items-center justify-between text-xs">
							<span class="font-medium">Human traffic</span>
							<span class="font-semibold text-success">96.8%</span>
						</div>
						<div class="mt-2 h-1.5 overflow-hidden rounded-full bg-muted">
							<div class="h-full w-[96.8%] rounded-full bg-success" />
						</div>
					</div>
				</div>
			</UIPanel>
		</section>
	</div>
</template>

<script setup lang="ts">
import {
	CalendarDaysIcon,
	ChartBarIcon,
	ChevronDownIcon,
	CursorArrowRaysIcon,
	DevicePhoneMobileIcon,
	GlobeAltIcon,
	LinkIcon,
	PlusIcon,
	UserGroupIcon,
} from '@heroicons/vue/24/outline'
import type { Link } from '~/types/api/link'

const fallbackLinks: Link[] = [
	{ ID: 1, Name: 'Summer campaign', Slug: 'summer-26', Target: 'https://example.com/summer', LinkType: 'campaign' },
	{ ID: 2, Name: 'Creator toolkit', Slug: 'creator-kit', Target: 'https://example.com/toolkit', LinkType: 'affiliate' },
	{ ID: 3, Name: 'Weekly newsletter', Slug: 'weekly', Target: 'https://example.com/newsletter', LinkType: 'content' },
	{ ID: 4, Name: 'Product launch', Slug: 'launch', Target: 'https://example.com/launch', LinkType: 'campaign' },
	{ ID: 5, Name: 'Bio link', Slug: 'maksim', Target: 'https://example.com', LinkType: 'social' },
	{ ID: 6, Name: 'Partner offer', Slug: 'partner', Target: 'https://example.com/offer', LinkType: 'affiliate' },
]

const links = ref<Link[]>([])
const displayLinks = computed(() => links.value.length ? links.value : fallbackLinks)

onMounted(async () => {
	try {
		const response = await fetch('http://localhost:8888/api/link')
		if (response.ok) links.value = await response.json()
	} catch {
		// The dashboard remains useful as a prototype when the API is offline.
	}
})

const trendPoints = [280, 360, 330, 510, 470, 580, 620, 540, 680, 730, 690, 820, 790, 910, 860, 970, 1020, 940, 1080, 1140, 1090, 1210, 1180, 1320, 1270, 1390, 1350, 1510, 1470, 1620]
const trendLabels = ['Jun 24', 'Jun 29', 'Jul 4', 'Jul 9', 'Jul 14', 'Jul 19', 'Jul 23']
const sourceItems = [
	{ label: 'Instagram', value: 4280, color: '#7c3aed' },
	{ label: 'Direct', value: 3120, color: '#2563eb' },
	{ label: 'Google', value: 2485, color: '#0d9488' },
	{ label: 'LinkedIn', value: 1732, color: '#d97706' },
	{ label: 'Other', value: 1230, color: '#94a3b8' },
]

const clicksFor = (id: number) => [4862, 3140, 2258, 1512, 802, 273][(id - 1) % 6] ?? 320
const trendFor = (id: number) => [24, 18, 11, 9, 6, 3][(id - 1) % 6] ?? 4
const campaignFor = (id: number) => ['Summer 2026', 'Affiliate', 'Organic', 'Launch', 'Social', 'Partner'][(id - 1) % 6] ?? 'General'
</script>
