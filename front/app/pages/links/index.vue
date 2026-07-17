<template>
	<div class="p-2">
		<Table title="Links" class="h-full" :loading="loading">
		<template #header>
			<HeaderCell>ID</HeaderCell>
			<HeaderCell>Name</HeaderCell>
			<HeaderCell>Slug</HeaderCell>
			<HeaderCell>Target</HeaderCell>
			<HeaderCell>Actions</HeaderCell>
		</template>
		<template #body>
			<tr v-for="(link) in links" :key="link.ID">
				<TableTextCell>{{ link.ID }}</TableTextCell>
				<TableTextCell>{{ link.Name}}</TableTextCell>
				<TableTextCell>{{ link.Slug }}</TableTextCell>
				<TableTextCell>{{ link.Target }}</TableTextCell>
				<TableTextCell>
					<NuxtLink :to="`/links/${link.ID}/clicks`">
						clicks
					</NuxtLink>
				</TableTextCell>
			</tr>
		</template>
		</Table>
  	</div>
</template>

<script setup lang="ts">
import HeaderCell from '~/components/Table/HeaderCell.vue';
import type { Link } from '~/types/api/link';

const links = ref<Link[]>()

const loading = ref<boolean>()

onMounted(async() =>{
  loading.value = true
  try {
    const linkResponse = await fetch('http://localhost:8888/api/link', {})
    links.value = await linkResponse.json()
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }

})
</script>