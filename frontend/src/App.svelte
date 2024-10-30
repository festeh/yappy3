<script lang="ts">
	import { onMount } from 'svelte';
	import {
		GetTimeLeft,
		StartPomodoro,
		StopPomodoro,
		PausePomodoro,
		ResumePomodoro,
		GetPomodoroButtons
	} from '../wailsjs/go/main/App';
	import { EventsOn } from '../wailsjs/runtime/runtime';
	import Button from './Button.svelte';
	import Focus from './Focus.svelte';

	let remaining = 'waiting...';
	let buttons: Array<{ text: string; method: string }> = [];
	let currentView = 'pomodoro';

	const updateButtons = async () => {
		buttons = await GetPomodoroButtons();
	};

	onMount(async () => {
		remaining = await GetTimeLeft();
		await updateButtons();
	});

	EventsOn('tick', async (newTimer) => {
		remaining = newTimer;
		await updateButtons();
	});

	const methodMap = {
		StartPomodoro,
		StopPomodoro,
		PausePomodoro,
		ResumePomodoro
	};

	async function handleClick(method: string) {
		await methodMap[method]();
		await updateButtons();
	}
</script>

<div class="flex h-screen w-full">
	<!-- Side Navigation -->
	<div class="flex w-16 flex-col gap-4 border-r border-gray-200 bg-white p-4">
		<button
			class="rounded-lg p-2 text-2xl transition-colors hover:bg-gray-100 {currentView === 'pomodoro' ? 'bg-gray-100' : ''}"
			on:click={() => (currentView = 'pomodoro')}
			title="Pomodoro Timer"
		>
			⏰
		</button>
		<button
			class="rounded-lg p-2 text-2xl transition-colors hover:bg-gray-100 {currentView === 'focus' ? 'bg-gray-100' : ''}"
			on:click={() => (currentView = 'focus')}
			title="Focus Mode"
		>
			🎯
		</button>
	</div>

	<!-- Main Content -->
	<main class="flex-1">
		{#if currentView === 'pomodoro'}
			<div class="flex h-full flex-col items-center justify-center gap-8">
				<div class="text-center text-6xl font-bold text-gray-300">
					{remaining}
				</div>
				<div class="flex gap-4">
					{#each buttons as button}
						<Button
							text={button.text}
							method={button.method}
							onClick={() => handleClick(button.method)}
						/>
					{/each}
				</div>
			</div>
		{:else if currentView === 'focus'}
			<Focus />
		{/if}
	</main>
</div>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
