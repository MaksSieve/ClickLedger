<template>
	<main class="h-svh w-full flex flex-col">
		<LayoutPageHeader class="w-full md:h-20 gap-4"/>
		<div class="w-full h-full px-16 py-4">
			<slot/>	
		</div>
		<LayoutPageFooter class="w-full md:h-20 justify-self-end" />
	</main>
</template>