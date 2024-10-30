<script lang="ts">
	import { onMount } from 'svelte';
	import { GetTimeLeft } from '../wailsjs/go/main/App';

	let remaining = 'waiting...';

	onMount(async () => {
		remaining = await GetTimeLeft();
	});
</script>

<main class="">
	<div class="flex items-center justify-center">
		{remaining}
	</div>
</main>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
