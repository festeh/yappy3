<script lang="ts">
	import { onMount } from 'svelte';
	import { GetTimeLeft, StartPomodoro, StopPomodoro } from '../wailsjs/go/main/App';

	let remaining = 'waiting...';

	onMount(async () => {
		remaining = await GetTimeLeft();
	});

	async function handleStart() {
		await StartPomodoro();
	}

	async function handleStop() {
		await StopPomodoro();
	}
</script>

<main class="h-screen w-full">
	<div class="flex h-full flex-col items-center justify-center gap-8">
		<div class="text-center text-6xl font-bold text-gray-800">
			{remaining}
		</div>
		<div class="flex gap-4">
			<button
				on:click={handleStart}
				class="rounded-lg bg-green-500 px-6 py-3 font-semibold text-white transition-colors hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
			>
				Start
			</button>
			<button
				on:click={handleStop}
				class="rounded-lg bg-red-500 px-6 py-3 font-semibold text-white transition-colors hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
			>
				Stop
			</button>
		</div>
	</div>
</main>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
