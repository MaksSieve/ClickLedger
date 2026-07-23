<template>
	<div class="space-y-4">
		<div v-for="item in items" :key="item.label">
			<div class="mb-1.5 flex items-center justify-between gap-4 text-sm">
				<div class="flex min-w-0 items-center gap-2.5">
					<span class="size-2 shrink-0 rounded-full" :style="{ backgroundColor: item.color ?? 'var(--color-primary)' }" />
					<span class="truncate font-medium">{{ item.label }}</span>
				</div>
				<div class="flex shrink-0 items-baseline gap-2">
					<span class="font-semibold">{{ item.value.toLocaleString() }}</span>
					<span class="w-9 text-right text-xs text-muted-foreground">{{ percentage(item.value) }}%</span>
				</div>
			</div>
			<div class="ml-[18px] h-1.5 overflow-hidden rounded-full bg-muted">
				<div
					class="h-full rounded-full"
					:style="{ width: `${percentage(item.value)}%`, backgroundColor: item.color ?? 'var(--color-primary)' }"
				/>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
type BreakdownItem = {
	label: string
	value: number
	color?: string
}

const props = defineProps<{
	items: BreakdownItem[]
}>()

const total = computed(() => props.items.reduce((sum, item) => sum + item.value, 0))
const percentage = (value: number) => total.value ? Math.round((value / total.value) * 100) : 0
</script>
