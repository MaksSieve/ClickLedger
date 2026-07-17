<template>
	<th
		:class="[
			cellClass,
			{
				'cursor-pointer': sortable,
			},
		]"
		@click="sort"
	>
		<div class="flex items-center justify-between w-full gap-2">
			<div class="w-full">
				<template v-if="text">
					{{ text }}
				</template>
				<template v-else>
					<slot />
				</template>
			</div>
			<component :is="sortingComponent" v-if="sortable" class="size-3.5 shrink-0" />
		</div>
	</th>
</template>

<script setup lang="ts">
import { ChevronDownIcon, ChevronUpDownIcon, ChevronUpIcon } from '@heroicons/vue/24/outline';
import { computed, inject } from 'vue';
import type { ComputedRef } from 'vue';
import type { SortDirection } from '~/types/common';

type Props = {
	sortable?: boolean;
	text?: string
	// eslint-disable-next-line vue/require-default-prop
	sortDirection?: SortDirection
};

const props = withDefaults(defineProps<Props>(),{
	sortable: false,
	initialSortDirection: 'none' as SortDirection,
	text: undefined
});

const tableCompact = inject<ComputedRef<boolean> | undefined>('tableCompact', undefined)
const tableFlat = inject<ComputedRef<boolean> | undefined>('tableFlat', undefined)

const cellClass = computed(() => ({
	'bg-card/80 backdrop-blur-md shadow-sm px-4 py-2.5 text-left text-[10px] font-semibold uppercase tracking-[0.12em] text-muted-foreground border-b border-border': !tableCompact?.value,
	'bg-card/80 backdrop-blur-md shadow-sm px-3 py-2 text-left text-[10px] font-semibold uppercase tracking-[0.12em] text-muted-foreground border-b border-border': tableCompact?.value,
	'bg-transparent': tableFlat?.value,
}))

const emit = defineEmits<{
		// eslint-disable-next-line no-unused-vars
		(e: 'sort', direction: SortDirection): void,
		// eslint-disable-next-line no-unused-vars
		(e: 'clearSorting'): void;
	}>();

const sortingComponent = computed(() => {
	switch(props.sortDirection) {
		case 'asc':
			return ChevronUpIcon
		case 'desc':
			return ChevronDownIcon
		case 'none':
		default:
			return ChevronUpDownIcon
	}
})

const sort = () => {
	if (props.sortable) {
		if (props.sortDirection === 'none') {
			emit('sort', 'desc');
		} else if (props.sortDirection === 'desc') {
			emit('sort', 'asc');
		} else {
			emit('sort', 'none');
		}
	}
};
</script>
