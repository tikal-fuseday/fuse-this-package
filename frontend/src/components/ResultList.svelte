<script>
  import Card from "./Card.svelte";
  export let items = [];
</script>

<style>
  ul {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-content: center;
  }
  li {
    margin: 10px;
  }
</style>

<ul>
  {#each items as item}
    <li>
      <Card {item} />
    </li>
  {/each}
</ul>
