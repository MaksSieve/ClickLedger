<template>
	<div class="w-full">
		<div class="h-56 w-full sm:h-64">
			<svg viewBox="0 0 720 240" class="size-full overflow-visible" preserveAspectRatio="none" role="img" :aria-label="ariaLabel">
				<defs>
					<linearGradient :id="gradientId" x1="0" x2="0" y1="0" y2="1">
						<stop offset="0%" stop-color="var(--color-primary)" stop-opacity=".28" />
						<stop offset="100%" stop-color="var(--color-primary)" stop-opacity="0" />
					</linearGradient>
				</defs>
				<line v-for="y in [18, 73, 128, 183, 238]" :key="y" x1="0" x2="720" :y1="y" :y2="y" stroke="var(--color-border)" stroke-width="1" />
				<path :d="areaPath" :fill="`url(#${gradientId})`" />
				<polyline :points="linePoints" fill="none" stroke="var(--color-primary)" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" vector-effect="non-scaling-stroke" />
				<circle :cx="lastPoint.x" :cy="lastPoint.y" r="5" fill="var(--color-card)" stroke="var(--color-primary)" stroke-width="3" vector-effect="non-scaling-stroke" />
			</svg>
		</div>
		<div class="mt-3 flex justify-between text-[10px] font-medium text-muted-foreground sm:text-xs">
			<span v-for="label in displayLabels" :key="label">{{ label }}</span>
		</div>
	</div>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
	points: number[]
	labels: string[]
	ariaLabel?: string
}>(), {
	ariaLabel: 'Clicks trend',
})

const gradientId = useId().replace(/:/g, '')

const coordinates = computed(() => {
	const max = Math.max(...props.points, 1)
	const min = Math.min(...props.points, 0)
	const range = Math.max(max - min, 1)
	return props.points.map((value, index) => ({
		x: props.points.length === 1 ? 0 : (index / (props.points.length - 1)) * 720,
		y: 218 - ((value - min) / range) * 185,
	}))
})

const linePoints = computed(() => coordinates.value.map(point => `${point.x},${point.y}`).join(' '))
const areaPath = computed(() => `M 0 240 L ${coordinates.value.map(point => `${point.x} ${point.y}`).join(' L ')} L 720 240 Z`)
const lastPoint = computed(() => coordinates.value.at(-1) ?? { x: 0, y: 218 })

const displayLabels = computed(() => {
	if (props.labels.length <= 6) return props.labels
	const step = Math.ceil(props.labels.length / 6)
	return props.labels.filter((_, index) => index % step === 0 || index === props.labels.length - 1)
})
</script>
