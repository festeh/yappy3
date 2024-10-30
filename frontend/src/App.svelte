<script lang="ts">
	import { onMount } from 'svelte';
	import { GetTimeLeft, StartPomodoro, StopPomodoro } from '../wailsjs/go/main/App';

	let remaining = 'waiting...';

	onMount(async () => {
		remaining = await GetTimeLeft();
	});
</script>

<main class="h-screen w-full">
	<div class="flex h-full items-center justify-center text-center text-4xl font-bold">
		{remaining}
	</div>
</main>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
