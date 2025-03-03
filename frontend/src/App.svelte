<script lang="ts">
	import Focus from './Focus.svelte';
	import Pomodoro from './Pomodoro.svelte';
	import Settings from './Settings.svelte';

	let currentView = 'pomodoro';
</script>

<div class="flex h-screen w-full">
	<!-- Side Navigation -->
	<div class="flex w-16 flex-col gap-4 border-r border-gray-200 p-4">
		<button
			class="flex items-center justify-center rounded-lg p-2 text-2xl transition-colors {currentView === 'pomodoro' ? 'bg-blue-100 text-blue-600' : 'hover:bg-gray-100'}"
			on:click={() => (currentView = 'pomodoro')}
			title="Pomodoro Timer"
		>
			⏰
		</button>
		<button
			class="flex items-center justify-center rounded-lg p-2 text-2xl transition-colors {currentView === 'focus' ? 'bg-blue-100 text-blue-600' : 'hover:bg-gray-100'}"
			on:click={() => (currentView = 'focus')}
			title="Focus Mode"
		>
			🎯
		</button>
		<button
			class="flex items-center justify-center rounded-lg p-2 text-2xl transition-colors {currentView === 'settings' ? 'bg-blue-100 text-blue-600' : 'hover:bg-gray-100'}"
			on:click={() => (currentView = 'settings')}
			title="Settings"
		>
			⚙️
		</button>
	</div>

	<!-- Main Content -->
	<main class="flex-1">
		{#if currentView === 'pomodoro'}
			<Pomodoro />
		{:else if currentView === 'focus'}
			<Focus />
		{:else if currentView === 'settings'}
			<Settings />
		{/if}
	</main>
</div>
