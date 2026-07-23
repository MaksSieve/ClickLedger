<template>
	<div class="space-y-6">
		<header class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
			<div>
				<p class="mb-1 text-sm font-medium text-primary">Workspace</p>
				<h1 class="text-2xl font-semibold tracking-tight sm:text-3xl">Links</h1>
				<p class="mt-1.5 text-sm text-muted-foreground">Manage and compare every tracked destination.</p>
			</div>
			<button type="button" class="flex h-10 w-fit items-center gap-2 rounded-lg bg-primary px-4 text-sm font-semibold text-primary-foreground shadow-sm">
				<PlusIcon class="size-4" />
				Create link
			</button>
		</header>

		<div class="flex flex-col gap-3 rounded-xl border border-border bg-card p-3 sm:flex-row sm:items-center sm:justify-between">
			<label class="relative block w-full sm:max-w-sm">
				<MagnifyingGlassIcon class="absolute left-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground" />
				<input v-model="search" type="search" placeholder="Search links..." class="h-10 w-full rounded-lg border border-border bg-background pl-9 pr-3 text-sm outline-none focus:ring-2 focus:ring-ring">
			</label>
			<button type="button" class="flex h-10 items-center justify-center gap-2 rounded-lg border border-border px-3 text-sm font-medium">
				<AdjustmentsHorizontalIcon class="size-4 text-muted-foreground" />
				All links
			</button>
		</div>

		<div class="overflow-hidden rounded-xl border border-border bg-card shadow-sm">
			<div v-if="loading" class="space-y-3 p-5">
				<div v-for="index in 6" :key="index" class="h-14 animate-pulse rounded-lg bg-muted" />
			</div>
			<div v-else class="overflow-x-auto">
				<table class="w-full min-w-190 text-left">
					<thead class="border-b border-border bg-muted/40 text-[10px] uppercase tracking-[0.12em] text-muted-foreground">
						<tr>
							<th class="px-5 py-3 font-semibold">Link</th>
							<th class="px-4 py-3 font-semibold">Type</th>
							<th class="px-4 py-3 text-right font-semibold">Clicks</th>
							<th class="px-4 py-3 font-semibold">Created</th>
							<th class="w-14 px-5 py-3" />
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						<tr v-for="link in filteredLinks" :key="link.ID" class="group">
							<td class="px-5 py-4">
								<NuxtLink :to="`/links/${link.ID}`" class="font-medium group-hover:text-primary">{{ link.Name }}</NuxtLink>
								<div class="mt-0.5 flex items-center gap-2 text-xs">
									<span class="text-primary">clickledger.io/{{ link.Slug }}</span>
									<span class="max-w-52 truncate text-muted-foreground">→ {{ link.Target }}</span>
								</div>
							</td>
							<td class="px-4 py-4"><span class="rounded-md bg-muted px-2 py-1 text-xs capitalize text-muted-foreground">{{ link.LinkType || 'tracked' }}</span></td>
							<td class="px-4 py-4 text-right text-sm font-semibold">{{ clicksFor(link.ID).toLocaleString() }}</td>
							<td class="px-4 py-4 text-sm text-muted-foreground">{{ dateFor(link.ID) }}</td>
							<td class="px-5 py-4">
								<NuxtLink :to="`/links/${link.ID}`" aria-label="Open link dashboard">
									<ChevronRightIcon class="size-4 text-muted-foreground group-hover:text-primary" />
								</NuxtLink>
							</td>
						</tr>
					</tbody>
				</table>
				<div v-if="filteredLinks.length === 0" class="p-12 text-center text-sm text-muted-foreground">No links match your search.</div>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {
	AdjustmentsHorizontalIcon,
	ChevronRightIcon,
	MagnifyingGlassIcon,
	PlusIcon,
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
const loading = ref(true)
const search = ref('')
const allLinks = computed(() => links.value.length ? links.value : fallbackLinks)
const filteredLinks = computed(() => {
	const query = search.value.trim().toLowerCase()
	if (!query) return allLinks.value
	return allLinks.value.filter(link => [link.Name, link.Slug, link.Target].some(value => value.toLowerCase().includes(query)))
})

onMounted(async () => {
	try {
		const response = await fetch('http://localhost:8888/api/link')
		if (response.ok) links.value = await response.json()
	} catch {
		// Use stable prototype data until the API is available.
	} finally {
		loading.value = false
	}
})

const clicksFor = (id: number) => [4862, 3140, 2258, 1512, 802, 273][(id - 1) % 6] ?? 320
const dateFor = (id: number) => ['Jul 18, 2026', 'Jul 12, 2026', 'Jul 4, 2026', 'Jun 28, 2026', 'Jun 12, 2026', 'May 30, 2026'][(id - 1) % 6] ?? 'Jul 1, 2026'
</script>
