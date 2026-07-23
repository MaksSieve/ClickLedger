<template>
	<div class="space-y-6">
		<nav class="flex items-center gap-2 text-xs text-muted-foreground">
			<NuxtLink to="/links" class="hover:text-foreground">Links</NuxtLink>
			<ChevronRightIcon class="size-3.5" />
			<span class="text-foreground">{{ link.Name }}</span>
		</nav>

		<header class="flex flex-col gap-4 xl:flex-row xl:items-start xl:justify-between">
			<div class="min-w-0">
				<div class="flex items-center gap-2">
					<span class="size-2 rounded-full bg-success" />
					<span class="text-xs font-medium text-success">Active</span>
				</div>
				<h1 class="mt-2 truncate text-2xl font-semibold tracking-tight sm:text-3xl">{{ link.Name }}</h1>
				<div class="mt-2 flex flex-wrap items-center gap-x-3 gap-y-2 text-sm">
					<a :href="shortUrl" target="_blank" class="font-medium text-primary hover:underline">{{ shortUrl }}</a>
					<button type="button" class="text-muted-foreground hover:text-foreground" :aria-label="copied ? 'Copied' : 'Copy short URL'" @click="copyLink">
						<CheckIcon v-if="copied" class="size-4 text-success" />
						<ClipboardDocumentIcon v-else class="size-4" />
					</button>
					<span class="hidden text-border sm:inline">|</span>
					<a :href="link.Target" target="_blank" class="max-w-md truncate text-muted-foreground hover:text-foreground">{{ link.Target }}</a>
				</div>
			</div>
			<div class="flex shrink-0 items-center gap-2">
				<button type="button" class="flex h-10 items-center gap-2 rounded-lg border border-border bg-card px-3 text-sm font-medium">
					<CalendarDaysIcon class="size-4 text-muted-foreground" />
					Last 30 days
					<ChevronDownIcon class="size-4 text-muted-foreground" />
				</button>
				<button type="button" class="grid size-10 place-items-center rounded-lg border border-border bg-card" aria-label="More options">
					<EllipsisHorizontalIcon class="size-5" />
				</button>
			</div>
		</header>

		<section class="grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
			<DashboardStatCard label="Total clicks" value="4,862" change="24.3%">
				<CursorArrowRaysIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Unique clicks" value="3,714" change="17.8%">
				<UserGroupIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Top source" value="Instagram" caption="2,041 clicks">
				<ArrowTrendingUpIcon class="size-5" />
			</DashboardStatCard>
			<DashboardStatCard label="Conversion rate" value="9.7%" change="1.4%">
				<BoltIcon class="size-5" />
			</DashboardStatCard>
		</section>

		<section class="grid gap-4 xl:grid-cols-[minmax(0,1.7fr)_minmax(18rem,0.8fr)]">
			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<div class="mb-6 flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
						<div>
							<h2 class="font-semibold">Click activity</h2>
							<p class="mt-1 text-xs text-muted-foreground">Performance for this link over time</p>
						</div>
						<div class="flex rounded-lg bg-muted p-1 text-xs font-medium">
							<button
								v-for="interval in ['Daily', 'Weekly']"
								:key="interval"
								type="button"
								class="rounded-md px-3 py-1.5"
								:class="chartInterval === interval ? 'bg-card text-foreground shadow-sm' : 'text-muted-foreground'"
								@click="chartInterval = interval"
							>
								{{ interval }}
							</button>
						</div>
					</div>
					<DashboardTrendChart :points="visibleTrend" :labels="trendLabels" aria-label="Link clicks over the last 30 days" />
				</div>
			</UIPanel>

			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<h2 class="font-semibold">Traffic sources</h2>
					<p class="mt-1 text-xs text-muted-foreground">Attributed referrers</p>
					<div class="mt-5">
						<DashboardBreakdownList :items="sourceItems" />
					</div>
				</div>
			</UIPanel>
		</section>

		<section class="grid gap-4 lg:grid-cols-2">
			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<div class="mb-5 flex items-center justify-between">
						<div>
							<h2 class="font-semibold">Audience</h2>
							<p class="mt-1 text-xs text-muted-foreground">Device breakdown</p>
						</div>
						<DevicePhoneMobileIcon class="size-5 text-muted-foreground" />
					</div>
					<DashboardBreakdownList :items="deviceItems" />
				</div>
			</UIPanel>
			<UIPanel titleless borderless>
				<div class="p-4 sm:p-5">
					<div class="mb-5 flex items-center justify-between">
						<div>
							<h2 class="font-semibold">Top locations</h2>
							<p class="mt-1 text-xs text-muted-foreground">Clicks by country</p>
						</div>
						<GlobeAltIcon class="size-5 text-muted-foreground" />
					</div>
					<DashboardBreakdownList :items="countryItems" />
				</div>
			</UIPanel>
		</section>

		<UIPanel titleless borderless class="overflow-hidden">
			<div class="flex items-center justify-between border-b border-border px-4 py-4 sm:px-5">
				<div>
					<h2 class="font-semibold">Recent clicks</h2>
					<p class="mt-1 text-xs text-muted-foreground">Latest activity for this link</p>
				</div>
				<NuxtLink :to="`/links/${route.params.id}/clicks`" class="text-xs font-semibold text-primary">View all</NuxtLink>
			</div>
			<div class="overflow-x-auto">
				<table class="w-full min-w-[680px] text-left text-sm">
					<thead class="border-b border-border bg-muted/40 text-[10px] uppercase tracking-[0.12em] text-muted-foreground">
						<tr>
							<th class="px-5 py-2.5 font-semibold">Time</th>
							<th class="px-4 py-2.5 font-semibold">Location</th>
							<th class="px-4 py-2.5 font-semibold">Source</th>
							<th class="px-4 py-2.5 font-semibold">Device</th>
							<th class="px-5 py-2.5 font-semibold">Campaign</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						<tr v-for="click in recentClicks" :key="`${click.time}-${click.location}`">
							<td class="px-5 py-3 text-muted-foreground">{{ click.time }}</td>
							<td class="px-4 py-3 font-medium">{{ click.location }}</td>
							<td class="px-4 py-3">{{ click.source }}</td>
							<td class="px-4 py-3 text-muted-foreground">{{ click.device }}</td>
							<td class="px-5 py-3"><span class="rounded-md bg-primary/10 px-2 py-1 text-xs font-medium text-primary">{{ click.campaign }}</span></td>
						</tr>
					</tbody>
				</table>
			</div>
		</UIPanel>
	</div>
</template>

<script setup lang="ts">
import {
	ArrowTrendingUpIcon,
	BoltIcon,
	CalendarDaysIcon,
	CheckIcon,
	ChevronDownIcon,
	ChevronRightIcon,
	ClipboardDocumentIcon,
	CursorArrowRaysIcon,
	DevicePhoneMobileIcon,
	EllipsisHorizontalIcon,
	GlobeAltIcon,
	UserGroupIcon,
} from '@heroicons/vue/24/outline'
import type { Link } from '~/types/api/link'

const route = useRoute()
const chartInterval = ref('Daily')
const copied = ref(false)
const link = ref<Link>({
	ID: Number(route.params.id) || 1,
	Name: 'Summer campaign',
	Slug: 'summer-26',
	Target: 'https://example.com/summer-sale',
	LinkType: 'campaign',
})

const shortUrl = computed(() => `clickledger.io/${link.value.Slug}`)

onMounted(async () => {
	try {
		const response = await fetch(`http://localhost:8888/api/link/${route.params.id}`)
		if (response.ok) link.value = await response.json()
	} catch {
		// Keep the prototype data visible when the API is offline.
	}
})

const copyLink = async () => {
	await navigator.clipboard.writeText(`https://${shortUrl.value}`)
	copied.value = true
	window.setTimeout(() => { copied.value = false }, 1600)
}

const dailyTrend = [92, 108, 101, 145, 128, 166, 172, 158, 194, 187, 210, 232, 218, 244, 236, 271, 260, 288, 274, 302, 326, 310, 341, 329, 356, 348, 382, 369, 402, 418]
const weeklyTrend = [674, 1006, 1418, 1818, 2036]
const visibleTrend = computed(() => chartInterval.value === 'Daily' ? dailyTrend : weeklyTrend)
const trendLabels = computed(() => chartInterval.value === 'Daily'
	? ['Jun 24', 'Jun 29', 'Jul 4', 'Jul 9', 'Jul 14', 'Jul 19', 'Jul 23']
	: ['Jun 24', 'Jul 1', 'Jul 8', 'Jul 15', 'Jul 22'])

const sourceItems = [
	{ label: 'Instagram', value: 2041, color: '#7c3aed' },
	{ label: 'Direct', value: 1124, color: '#2563eb' },
	{ label: 'Google', value: 817, color: '#0d9488' },
	{ label: 'LinkedIn', value: 536, color: '#d97706' },
	{ label: 'Other', value: 344, color: '#94a3b8' },
]
const deviceItems = [
	{ label: 'Mobile', value: 3112, color: '#7c3aed' },
	{ label: 'Desktop', value: 1409, color: '#2563eb' },
	{ label: 'Tablet', value: 341, color: '#0d9488' },
]
const countryItems = [
	{ label: 'United States', value: 1945, color: '#7c3aed' },
	{ label: 'Germany', value: 973, color: '#2563eb' },
	{ label: 'United Kingdom', value: 827, color: '#0d9488' },
	{ label: 'Canada', value: 486, color: '#d97706' },
]
const recentClicks = [
	{ time: '2 min ago', location: 'Berlin, DE', source: 'Instagram', device: 'iPhone · Safari', campaign: 'summer_26' },
	{ time: '8 min ago', location: 'Austin, US', source: 'Google', device: 'Desktop · Chrome', campaign: 'summer_26' },
	{ time: '14 min ago', location: 'London, UK', source: 'Direct', device: 'Android · Chrome', campaign: 'newsletter' },
	{ time: '21 min ago', location: 'Toronto, CA', source: 'LinkedIn', device: 'Desktop · Edge', campaign: 'summer_26' },
]
</script>
