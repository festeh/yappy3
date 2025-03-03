<script lang="ts">
  export let icon: string;
  export let title: string;
  export let view: string;
  export let currentView: string;
  export let onClick: () => void;
</script>

<button
  class="flex items-center justify-center rounded-lg text-4xl p-1 text-center transition-colors {currentView === view
    ? 'bg-blue-100 text-blue-600'
    : 'hover:bg-gray-100'}"
  on:click={onClick}
  {title}
>
  {icon}
</button>
