<script>
export let item;
</script>

<style>
  .card-container {
    border: 1px solid grey;
    border-radius: 10px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    width: 300px;
    color: darkblue;
  }
</style>

<section class="card-container">
  <header class="card-header" >{item.name}</header>
  <main >{item.description}</main>
  <footer class="card-footer" >
</section>
