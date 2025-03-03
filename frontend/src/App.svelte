<script lang="ts">
  import Focus from "./Focus.svelte";
  import Pomodoro from "./Pomodoro.svelte";
  import Settings from "./Settings.svelte";
  import NavButton from "./components/NavButton.svelte";

  let currentView = "pomodoro";
</script>

<div class="flex h-screen w-full">
  <!-- Side Navigation -->
  <div class="flex w-12 flex-col gap-2 border-r border-gray-200 py-4">
    <NavButton
      icon="mdi:timer-outline"
      title="Pomodoro Timer"
      view="pomodoro"
      {currentView}
      onClick={() => (currentView = "pomodoro")}
    />
    <NavButton
      icon="mdi:target"
      title="Focus Mode"
      view="focus"
      {currentView}
      onClick={() => (currentView = "focus")}
    />
    <NavButton
      icon="mdi:cog-outline"
      title="Settings"
      view="settings"
      {currentView}
      onClick={() => (currentView = "settings")}
    />
  </div>

  <!-- Main Content -->
  <main class="flex-1">
    {#if currentView === "pomodoro"}
      <Pomodoro />
    {:else if currentView === "focus"}
      <Focus />
    {:else if currentView === "settings"}
      <Settings />
    {/if}
  </main>
</div>
