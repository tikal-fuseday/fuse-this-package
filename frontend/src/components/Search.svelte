<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    let timeout = null;
    export let value = '';

    function onValueChange(e) {
        clearTimeout(timeout);
        const { value } = e.target;
        timeout = setTimeout(() => dispatch('change', { value }), 500);
    }
</script>

<div class="search">
    <input type="text" bind:value={value} on:input={onValueChange} />
    <span class="fa fa-search" />
</div>

<style>
    .search {
        position: relative;
    }

    input {
        width: 500px;
        height: 50px;
        box-shadow: 5px 5px 20px 0px lightgrey;
        border: none;
        border-radius: 10px;
        outline: none;
        margin-bottom: 50px;
        padding: 0 20px 0 40px;
        font-size: 24px;
    }

    .fa {
        position: absolute;
        left: 10px;
        top: 15px;
        font-size: 20px;
        color: rgba(0, 0, 0, 0.2);
    }
</style>
