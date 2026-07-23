<template>
	<div class="rounded-xl border border-border bg-card p-4 shadow-sm sm:p-5">
		<div class="flex items-start justify-between gap-3">
			<div>
				<p class="text-xs font-medium text-muted-foreground">{{ label }}</p>
				<p class="mt-2 text-2xl font-semibold tracking-tight sm:text-[1.7rem]">{{ value }}</p>
			</div>
			<div class="grid size-9 place-items-center rounded-lg bg-primary/10 text-primary">
				<slot />
			</div>
		</div>
		<div class="mt-3 flex items-center gap-1.5 text-xs">
			<span
				v-if="change"
				class="font-semibold"
				:class="positive ? 'text-success' : 'text-destructive'"
			>
				{{ positive ? '↗' : '↘' }} {{ change }}
			</span>
			<span class="text-muted-foreground">{{ caption }}</span>
		</div>
	</div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
	label: string
	value: string
	change?: string
	caption?: string
	positive?: boolean
}>(), {
	change: '',
	caption: 'vs previous period',
	positive: true,
})
</script>
