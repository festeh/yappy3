<script lang="ts">
	import Focus from './Focus.svelte';
	import Pomodoro from './Pomodoro.svelte';

	let currentView = 'pomodoro';
</script>

<div class="flex h-screen w-full">
	<!-- Side Navigation -->
	<div class="flex w-16 flex-col gap-4 border-r border-gray-200 p-4">
		<button
			class="flex items-center justify-center rounded-lg p-2 text-2xl transition-colors"
			on:click={() => (currentView = 'pomodoro')}
			title="Pomodoro Timer"
		>
			⏰
		</button>
		<button
			class="flex items-center justify-center rounded-lg p-2 text-2xl transition-colors"
			on:click={() => (currentView = 'focus')}
			title="Focus Mode"
		>
			🎯
		</button>
	</div>

	<!-- Main Content -->
	<main class="flex-1">
		{#if currentView === 'pomodoro'}
			<Pomodoro />
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
