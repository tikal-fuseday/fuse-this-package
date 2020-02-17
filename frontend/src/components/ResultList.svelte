<script>
  import Card from "./Card.svelte";
  export let items = [];
</script>

<style>
  .results {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-content: center;
  }
</style>

<div class="results">
  {#each items as item}
    <Card {item} />
  {/each}
</div>