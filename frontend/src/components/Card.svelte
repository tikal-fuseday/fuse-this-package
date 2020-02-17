<script>
  export let item = {};
  let status = 'warning';

  function kFormatter(num) {
    return Math.abs(num) > 999 ? Math.sign(num)*((Math.abs(num)/1000).toFixed(1)) + 'k' : Math.sign(num)*Math.abs(num)
  }

  function gitlink(str) {
    return str.replace(/(api\.)|(users\/)|(repos\/)/g, '');
  }
</script>

<style>
  .card-container {
    width: 300px;
    height: 170px;
    box-shadow: 5px 5px 20px 0px lightgrey;
    border-radius: 10px;
    padding: 20px;
    margin: 10px;
    display: flex;
    flex-direction: column;
    position: relative;
  }

  .card-header {
    width: calc(100% - 50px);
  }

  .card-header a:first-child {
    font-size: 20px;
  }

  .card-footer {
    display: flex;
    justify-content: space-between;
  }

  .card-footer div {
    color: rgb(0 0 0 / 0.4);
  }

  .card-footer div:last-child {
    width: 50px;
    text-align: right;
  }

  main {
    margin-top: 10px;
    flex: 1;
  }

  .description {
    -webkit-box-orient: vertical;
    display: -webkit-box;
    -webkit-line-clamp: 5;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: normal;
  }

  .score {
    position: absolute;
    font-size: 50px;
    right: 10px;
    top: -4px;
  }

  .score span {
    position: absolute;
    top: 0;
    right: 0;
    font-size: 14px;
    height: 80%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    color: black;
  }

  .good .score {
    color: rgb(0 255 0 / 0.2);
  }

  .ok .score {
    color: rgb(255 255 0 / 0.2);
  }

  .bad .score {
    color: rgb(255 0 0 / 0.2);
  }
</style>

<section
  class="card-container"
  class:good="{item.our_score >= 8}"
  class:ok="{item.our_score >= 5 && item.our_score < 8}"
  class:bad="{(item.our_score || 0) < 5}"
>
  <div class="score fa fa-bookmark">
    <span>{(item.our_score || 0).toFixed(1)}</span>
  </div>
  <header class="card-header">
    <a href={gitlink(item.url)} target="_blank">{item.name}</a> / 
    <a href={gitlink(item.owner.url)} target="_blank">{item.owner.login}</a>
  </header>
  <main><div class="description">{item.description}</div></main>
  <footer class="card-footer">
    <div>
      <span class="fa fa-star"></span>
      {kFormatter(item.stargazers_count)}
    </div>
    <div>
      <span class="fa fa-code-fork"></span>
      {kFormatter(item.forks_count)}
    </div>
    <div>
      <span class="fa fa-eye"></span>
      {kFormatter(item.watchers_count)}
    </div>
    <div>
      <span class="fa fa-bug"></span>
      {kFormatter(item.open_issues_count)}
    </div>
    <div>
      {#if item.trend > 1}
      <span class="fa fa-line-chart"></span>
      {/if}
    </div>
  </footer>
</section>
