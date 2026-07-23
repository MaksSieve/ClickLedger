<template>
	<div class="min-h-svh bg-background text-foreground">
		<LayoutPageHeader />
		<main class="mx-auto min-h-[calc(100svh-4.25rem)] max-w-[1440px] px-4 py-6 sm:px-6 lg:pl-72 lg:pr-8 lg:py-8">
			<slot />
		</main>
	</div>
</template>
