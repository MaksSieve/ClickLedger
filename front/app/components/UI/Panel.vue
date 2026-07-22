<template>
	<div 
		:style="style" 
		class="glass-card flex flex-col"
	>
		<div
			v-if="!minimizedFrame && !titleless"
			:class="{
				'p-4': title && titleStyle === 'default',
				'py-2 px-4': title && titleStyle === 'flat',
				'cursor-pointer': minimizable,
				'border-b': title
			}"
			class="relative z-30 flex border-border py-2 px-2 items-center justify-between min-h-12 w-full backdrop-blur-md sm:rounded-t-xl"
			@click="minimizeClick"
		>
			<div class="flex justify-between items-center w-fit">
				<span
					v-if="title"
					class="w-full"
					:class="{
						'text-lg font-medium text-foreground': titleStyle === 'default',
						'text-sm  uppercase tracking-wider font-semibold text-muted-foreground': titleStyle === 'flat'
					}"
				>
					{{ title }}
				</span>
				<slot v-if="!minimized" name="header" />
			</div>
			<div class="flex gap-2 items-center justify-end w-fit">
				<slot v-if="!minimized" name="actions" />
				<template v-if="minimizable">
					<ChevronUpIcon v-if="!minimized" class="w-6 h-6" />
					<ChevronDownIcon v-else class="w-6 h-6" />
				</template>
			</div>
			<button
				v-if="closable"
				type="button"
				class="text-muted-foreground hover:text-foreground focus:outline-none focus:text-foreground transition ease-in-out duration-150 cursor-pointer"
				aria-label="Close"
				@click="emit('close')"
			>
				<XMarkIcon class="size-6" />
			</button>
		</div>
		<div
			class="relative z-0 flex min-h-0 overflow-y-auto w-full h-full overscroll-none"
			:class="{
				'h-[calc(100% - 3rem)]': !titleless,
			}"
		>	
			<div
				v-if="!minimized"
				class="h-full w-full"
				:class="{
					'px-4 py-5 sm:p-6': !borderless,
				}"
			>
				<slot  />
			</div>
			
			
			<!-- <DivideIcon  v-else-if="loading && !minimized" class="flex justify-center items-center h-full w-full">
				<UiLoadingBars class="m-auto"/>
			</div> -->
			<div
				v-if="loading && !minimized" 
				class="absolute w-full h-full flex bg-background/50 backdrop-blur-sm z-100 rounded-xl"
			>
				<div class="m-auto flex flex-col items-center gap-3 p-6 text-center glass-card">
					<div class="size-14 rounded-full bg-accent/60 border border-border flex">
						<UiLoadingSpinner class="m-auto size-10"/>
					</div>
					<span class="block text-muted-foreground">
						{{ loadingText ?? 'Lade Daten' }}
					</span>
				</div>
			</div>
		</div>
		
	</div>
</template>

<script setup lang="ts">
import { ChevronDownIcon, ChevronUpIcon } from '@heroicons/vue/24/outline'
import { XMarkIcon } from '@heroicons/vue/24/solid'


type Props = {
	title?: string
	borderless?: boolean
	closable?: boolean
	maxHeight?: string
	minimizedFrame?: boolean
	loading?: boolean,
	minimizable?: boolean
	minimizedByDefault?: boolean
	titleStyle?: 'default' | 'flat'
	titleless?: boolean
	loadingText?: string
}

const props = withDefaults(defineProps<Props>(), {
	title: '',
	borderless: false,
	closable: false,
	maxHeight: '',
	loading: false,
	minimizedFrame: false,
	minimizable: false,
	minimizedByDefault: false,
	titleStyle: 'default',
	titleless: false
})

type Emits = {
	// eslint-disable-next-line no-unused-vars
	(e: 'close'): void
}

const emit = defineEmits<Emits>()

const minimized = ref(props.minimizedByDefault)
const style = props.maxHeight ? `max-height: ${props.maxHeight}` : '';

function minimizeClick() {
	if (props.minimizable) {
		minimized.value = !minimized.value;
	}
}

</script>
