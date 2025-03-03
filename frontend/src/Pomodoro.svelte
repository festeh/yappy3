<script lang="ts">
	import { onMount } from 'svelte';
	import * as Pomo from '../wailsjs/go/pomodoro/Pomo';
	import { EventsOn } from '../wailsjs/runtime/runtime';
	import Button from './Button.svelte';

	let remaining = 'waiting...';
	let buttons: Array<{ text: string; method: string }> = [];

	const updateButtons = async () => {
		buttons = await Pomo.GetButtons();
	};

	onMount(async () => {
		remaining = await Pomo.GetTimeLeft();
		await updateButtons();
	});

	EventsOn('tick', async (newTimer) => {
		remaining = newTimer;
		await updateButtons();
	});

	const methodMap = {
		"start": Pomo.Start,
		"stop": Pomo.Stop,
		"pause": Pomo.Pause,
    "resume": Pomo.Resume
	};

	async function handleClick(method: string) {
		await methodMap[method]();
		await updateButtons();
	}
</script>

<div class="flex h-full flex-col items-center justify-center gap-8">
	<div class="text-center text-6xl font-bold text-gray-300">
		{remaining}
	</div>
	<div class="flex gap-4">
		{#each buttons as button}
			<Button text={button.text} method={button.method} onClick={() => handleClick(button.method)} />
		{/each}
	</div>
</div>
