<template>
	<td :class="cellClass" :style="cellStyle">
		<div v-if="loading" class="flex h-full w-full">
			<LoadingBars class="m-auto size-4" />
		</div>

		<slot v-else />
	</td>
</template>

<script setup lang="ts">
import { computed, inject } from 'vue'
import type { ComputedRef } from 'vue'

type Props = {
	loading?: boolean
}

const props = withDefaults(defineProps<Props>(), {
	loading: false,
})

const tableCompact = inject<ComputedRef<boolean> | undefined>('tableCompact', undefined)
const tableFlat = inject<ComputedRef<boolean> | undefined>('tableFlat', undefined)

const cellClass = computed(() => ({
	'px-4 py-3 h-8 whitespace-nowrap border-b text-sm text-foreground': !tableCompact?.value,
	'px-3 py-2 h-8 whitespace-nowrap border-b text-sm text-foreground': tableCompact?.value,
	'bg-muted/40': props.loading,
}))

const cellStyle = computed(() => ({
	borderColor: 'color-mix(in oklab, var(--color-border) 40%, transparent)',
	backgroundColor: tableFlat?.value && !props.loading ? 'transparent' : undefined,
}))

const loading = computed(() => props.loading)
</script>
