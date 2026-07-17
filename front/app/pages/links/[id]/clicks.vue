<template>
	<div class="p-2">
		<Table title="Clicks" class="h-full" :loading="loading">
		<template #header>
			<HeaderCell>ID</HeaderCell>
			<HeaderCell>Clicked At</HeaderCell>
			<HeaderCell>Link ID</HeaderCell>
			<HeaderCell>Target</HeaderCell>
		</template>
		<template #body>
			<tr v-for="(click) in clicks" :key="click.ID">
			<TableTextCell>{{ click.ID }}</TableTextCell>
			<TableTextCell>{{ useDateFormat(click.CreatedAt, 'DD MMM YYYY', {locales: 'de-DE'}) }}</TableTextCell>
			<TableTextCell>
				<NuxtLink :to="`/links/${linkId}`">{{ click.LinkID }}</NuxtLink></TableTextCell>
			<TableTextCell>{{ (click as Click).Target }}</TableTextCell>
			</tr>
		</template>
		</Table>
  	</div>
</template>

<script setup lang="ts">
import HeaderCell from '~/components/Table/HeaderCell.vue';
import type { Click } from '~/types/api/click';
import type { Link } from '~/types/api/link';
import { useDateFormat } from '@vueuse/core';

const clicks = ref<Click[]>([])
const link = ref<Link>()

const loading = ref<boolean>()

const linkId = useRoute().params.id

onMounted(async() =>{
  loading.value = true


	
  try {
    const linkResponse = await fetch('http://localhost:8888/api/link/1', {})
    link.value = await linkResponse.json()
    const clicksResponse = await fetch('http://localhost:8888/api/link/1/clicks', {})
    clicks.value = await clicksResponse.json()
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }

})
</script>