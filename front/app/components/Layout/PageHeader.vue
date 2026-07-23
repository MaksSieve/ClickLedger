<template>
	<div class="contents">
		<header class="sticky top-0 z-50 flex h-[4.25rem] items-center justify-between border-b border-border bg-background/90 px-4 backdrop-blur-xl sm:px-6 lg:hidden">
			<NuxtLink to="/" aria-label="ClickLedger dashboard">
				<img src="/logo-full.png" class="h-10 w-auto max-w-44 object-contain" alt="ClickLedger">
			</NuxtLink>
			<button
				type="button"
				class="grid size-10 place-items-center rounded-lg border border-border bg-card"
				:aria-expanded="mobileOpen"
				aria-label="Toggle navigation"
				@click="mobileOpen = !mobileOpen"
			>
				<XMarkIcon v-if="mobileOpen" class="size-5" />
				<Bars3Icon v-else class="size-5" />
			</button>
		</header>

		<aside
			class="fixed inset-y-0 left-0 z-40 flex w-64 flex-col border-r border-border bg-card/95 p-4 backdrop-blur-xl transition-transform lg:translate-x-0"
			:class="mobileOpen ? 'translate-x-0 top-[4.25rem]' : '-translate-x-full'"
		>
			<NuxtLink class="hidden h-16 items-center px-2 lg:flex" to="/">
				<img src="/logo-full.png" class="h-12 w-auto object-contain" alt="ClickLedger">
			</NuxtLink>

			<nav class="mt-4 flex flex-1 flex-col gap-1.5" @click="mobileOpen = false">
				<p class="px-3 pb-2 text-[10px] font-semibold uppercase tracking-[0.18em] text-muted-foreground">Workspace</p>
				<NuxtLink
					v-for="item in navigation"
					:key="item.to"
					:to="item.to"
					class="nuxt-link flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-muted-foreground transition hover:bg-muted hover:text-foreground"
				>
					<component :is="item.icon" class="size-5" />
					{{ item.label }}
				</NuxtLink>
			</nav>

			<div class="rounded-xl border border-border bg-muted/50 p-3">
				<div class="mb-3 flex items-center gap-3">
					<div class="grid size-9 place-items-center rounded-full bg-primary text-xs font-bold text-primary-foreground">MS</div>
					<div class="min-w-0">
						<p class="truncate text-sm font-semibold">Maksim</p>
						<p class="text-xs text-muted-foreground">Free plan</p>
					</div>
					<EllipsisHorizontalIcon class="ml-auto size-5 text-muted-foreground" />
				</div>
				<div class="h-1.5 overflow-hidden rounded-full bg-border">
					<div class="h-full w-3/5 rounded-full bg-primary" />
				</div>
				<p class="mt-2 text-[11px] text-muted-foreground">6 of 10 links used</p>
			</div>
		</aside>

		<button
			v-if="mobileOpen"
			type="button"
			class="fixed inset-0 top-[4.25rem] z-30 bg-foreground/20 backdrop-blur-sm lg:hidden"
			aria-label="Close navigation"
			@click="mobileOpen = false"
		/>
	</div>
</template>

<script setup lang="ts">
import {
	Bars3Icon,
	ChartBarIcon,
	EllipsisHorizontalIcon,
	LinkIcon,
	UserCircleIcon,
	XMarkIcon,
} from '@heroicons/vue/24/outline'

const mobileOpen = ref(false)

const navigation = [
	{ label: 'Overview', to: '/', icon: ChartBarIcon },
	{ label: 'Links', to: '/links', icon: LinkIcon },
	{ label: 'Profile', to: '/profile/1', icon: UserCircleIcon },
]
</script>
