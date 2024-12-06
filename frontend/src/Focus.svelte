<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { Connect, Disconnect, GetFocusing, FocusNow } from '../wailsjs/go/coach/Coach';
	import { EventsOn } from '../wailsjs/runtime';

	let focusing: boolean | null = null;

	EventsOn('focusing', async (updated: boolean) => {
    console.log("Got focusing", updated)
		focusing = updated;
	});

	async function tryConnect() {
		try {
			await Connect();
			focusing = await GetFocusing();
		} catch (error) {
			console.log("Connection error:", error);
			setTimeout(tryConnect, 5000); // Retry every 5 seconds
		}
	}

	onMount(() => {
		tryConnect();
	});

	onDestroy(() => {
		// try {
		// 	Disconnect();
		// } catch (error) {
		// 	console.log(error);
		// }
	});
</script>

<main class="h-screen w-full">
	<div class="flex h-full flex-col items-center justify-center gap-8">
		<div class="text-center text-4xl font-bold text-gray-300">
			{#if focusing === null}
				Loading...
			{:else}
				{focusing ? 'Focusing' : 'Not Focusing'}
			{/if}
		</div>
		<button
			on:click={() => FocusNow()}
			class="rounded-lg bg-gray-700 px-6 py-3 text-xl text-gray-300 hover:bg-gray-600"
		>
			Focus Now
		</button>
	</div>
</main>
