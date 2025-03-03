<script lang="ts">
  export let text: string;
  export let method: string;
  export let onClick: () => Promise<void>;

  const colorMap: { [key: string]: string } = {
    start: "bg-green-500 hover:bg-green-600 focus:ring-green-500",
    pause: "bg-yellow-500 hover:bg-yellow-600 focus:ring-yellow-500",
    resume: "bg-green-500 hover:bg-green-600 focus:ring-green-500",
    stop: "bg-red-500 hover:bg-red-600 focus:ring-red-500",
  };
</script>

<button
  on:click={onClick}
  class="rounded-lg px-6 py-3 font-semibold text-white transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 {colorMap[
    method
  ]}"
>
  {text}
</button>
