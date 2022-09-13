<script lang="ts">
    let name = "a";

    function greet() {
        name += "aa";
    }
</script>

<div class="p-4 grid gap-4 max-w-5xl ">
    <div class="bg-neutral-500 text-red-300">
        <button on:click={greet}> {name} </button>
    </div>

    <h1 class="text-center text-fuchsia-500 mt-3">
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolores magni
        sequi consequatur a totam nobis iure quam reprehenderit quasi maiores
        ipsum sint quas assumenda, dolorum enim dolor blanditiis obcaecati rem.
    </h1>
    <div class="h-16 bg-blue-400" />
    <div class="h-16 bg-blue-400" />
    <div class="h-16 bg-blue-400" />
    <div class="h-16 bg-blue-400" />
    <div class="h-16 bg-blue-400" />
</div>
