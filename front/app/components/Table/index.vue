<template>
	<div class="flex flex-col overflow-hidden overscroll-none" 
		:class="{
			...rootClass,
			'glass-card-flat-top': props.stackedTop,
			'glass-card-flat-bottom': props.stackedBottom,
			'glass-card': !props.stackedTop && !props.stackedBottom,
		}"
	>
		<div v-if="!withoutTitle" class="flex justify-between w-full items-center px-4 sticky z-20 top-0 bg-card/80 backdrop-blur-md border-b border-border h-fit py-2 shrink-0 gap-4">
			<span class="text-lg font-semibold text-foreground">{{ title }}</span>
			<div class="flex items-center gap-2 w-fit">

				<slot v-if="!folded" name="actions" />
				<button
					v-if="foldable"
					type="button"
					class="text-foreground cursor-pointer"
					@click="folded = !folded"
				>
					<ChevronDownIcon v-if="!folded" class="size-5" />
					<ChevronUpIcon v-else class="size-5" />
				</button>
			</div>
		</div>
		<div v-if="withMenu">
			<slot name="menu" />
		</div>
		<div
			v-if="!folded"
			class="relative flex min-h-0 flex-col w-full h-full"
			:class="{
				'overflow-y-auto overscroll-none': !withoutScroll
			}"
		>
			<div class="min-h-full w-full border-border flex-col">
				<table
					class="min-w-full border-separate text-foreground"
					:class="tableClass"
					style="border-spacing: 0"
				>
					<thead class="">
						<tr class="sticky top-0 z-10">
							<slot name="header"/>
						</tr>
					</thead>
					<template v-if="loading">
						<tr 
							v-for="idx in Array.from(Array(numberOfRows !== 0 ? numberOfRows : 25).keys())" 
							:key="idx + '_string_placeholder'"
						>
							<TableTextCell 
								v-for="n in Array.from(Array(numberOfColumns).keys())"
								:key="n + '_column_placeholder'"
							>
								<div class="w-full p-2 flex flex-col gap-2">
									<div class="w-full h-2 bg-muted rounded-full animate-pulse" />
									<div class="w-2/3 h-1.5 bg-muted animate-pulse" />
								</div>
							</TableTextCell>
						</tr>
					</template>
					<tbody
						v-if="!loading"
						class="rounded-b-lg"
						:class="tableClass"
					>
						<slot name="body"/>
					</tbody>
					<slot
						v-if="!loading"
						name="multibody"
						class="bg-card rounded-b-lg"
					/>
				</table>
			</div>
			<div v-if="empty && !loading" class="absolute w-full h-full flex p-auto bg-background/50 backdrop-blur-sm z-10">
				<div class="m-auto flex flex-col items-center gap-3 p-6 text-center glass-card">
					<div class="size-10 rounded-full bg-accent/60 border border-border" />
					<span class="block text-muted-foreground">
						{{ emptyText }}
					</span>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ChevronDownIcon, ChevronUpIcon } from '@heroicons/vue/24/outline'
import type { Slot } from 'vue'

type Props = {
	title?: string
	stickyHeader?: boolean
	loading?: boolean
	withoutTitle?: boolean
	withMenu?: boolean
	empty?: boolean
	emptyText?: string
	foldable?: boolean
	withoutScroll?: boolean
	showRowCount?: boolean
	rowCount?: number
	rowCountLabel?: string
	withSearch?: boolean
	searchable?: boolean
	search?: string
	searchPlaceholder?: string
	compact?: boolean
	flat?: boolean
	stackedTop?: boolean
	stackedBottom?: boolean
}

type Slots = {
	header: Slot
	body: Slot
	multibody: Slot
	actions: Slot
	menu: Slot
}

type Emits = {
	// eslint-disable-next-line no-unused-vars
	(e: 'update:search' | 'search', value: string): void
}

const slots = defineSlots<Slots>()
const emit = defineEmits<Emits>()
const folded = ref<boolean>(false)
const localSearch = ref<string>('')

const props = withDefaults(defineProps<Props>(), {
	stickyHeader: false,
	loading: false,
	title: undefined,
	emptyText: 'Es ist leider keine Daten',
	rowCount: undefined,
	rowCountLabel: 'Zeilen',
	search: undefined,
	searchPlaceholder: 'Suchen ...'
})

const tableCompact = computed(() => Boolean(props.compact))
const tableFlat = computed(() => Boolean(props.flat))

provide('tableCompact', tableCompact)
provide('tableFlat', tableFlat)

const rootClass = computed(() => ({
	'rounded-none border-0 shadow-none bg-transparent backdrop-blur-none': props.withoutTitle || props.flat,
}))

const tableClass = computed(() => ({
	'bg-card': !props.flat,
	'bg-transparent': props.flat,
}))

const numberOfColumns = computed(() => {
	if (slots.header) {
		return slots.header().length
	} else {
		return 1
	}
})

const numberOfRows = computed(() => {
	if (slots.body) {
		return slots.body()[0]?.children?.length
	} else if (slots.multibody) {
		return slots.multibody()[0]?.children?.length
	} else {
		return 0
	}
})

</script>
