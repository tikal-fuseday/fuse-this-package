<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    let timeout = null;
    export let value = 'hello';

    function onValueChange(val) {
        clearTimeout(timeout);
        timeout = setTimeout(() => dispatch('change', val), 500);
    }
</script>

<style>
    .search input {
        display: flex;
    }
</style>

<div class="search">
    <input type="text" bind:value={value} on:change={onValueChange} />
</div>
