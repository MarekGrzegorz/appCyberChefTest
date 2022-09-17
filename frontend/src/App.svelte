<script>
    import { Greet } from "../wailsjs/go/main/App.js";

    let name = "TEST - write something", result;

    $: name
    $: console.log(result)

    function Greeter() {
        Greet(name).then((res) => (result = res));
        name = "res - Golang/console NEXT"    
    }
</script>

<main>
    <form on:submit|preventDefault={Greeter}>
        <label for="">
            <input class="test_input" type="text" bind:value={name} />
        </label>
        
    </form>
</main>

<style>
    .test_input {
        color: coral;
    }
</style>
