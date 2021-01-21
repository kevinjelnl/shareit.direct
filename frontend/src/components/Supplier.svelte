<script>
    const params = new URLSearchParams(window.location.search);
    // https://www.npmjs.com/package/crypto-js
    let token = "";
    let pwfield;
    let pwButtonText = "show";

    if (params.get("token")) {
        token = params.get("token");
    }
    $: tlen = token.length;

    let keyphrase = false;
    let value = "";

    // functions that shares the supplied secret
    async function shareSecret() {
        console.log(value);
        const res = await fetch("http://127.0.0.1:8080/supply", {
            method: "POST",
            body: JSON.stringify({ token, value, keyphrase }),
        });
        const todo = await res.json();
        console.log(todo);
    }
    // toggle the password field
    const togglePW = () => {
        if (pwfield.type === "password") {
            pwfield.type = "text";
            pwbutton = "hide";
        } else {
            pwfield.type = "password";
            pwbutton = "show";
        }
    };
</script>

<div class="column is-full">
    <h2 class="title" style="display:inline">Supply a secret:</h2>
</div>

<div class="column is-full">
    <div class="field">
        <div class="columns">
            <div class="column is-one-fifth">
                <input
                    class="input"
                    bind:value={token}
                    placeholder="Token"
                    style="text-transform: uppercase"
                    autocomplete="off"
                />
            </div>
        </div>
        <!-- enable inputfield on 5 charlen of token-->
        {#if tlen === 5}
            <!-- https://stackoverflow.com/questions/57392773/error-type-attribute-cannot-be-dynamic-if-input-uses-two-way-binding -->
            <input
                bind:this={pwfield}
                class="input"
                type="password"
                bind:value
            />
            <span class="icon is-small is-left"><i class="fas fa-lock" /></span>
            <button class="button" on:click={shareSecret}>share</button>
            <button class="button" on:click={togglePW}>{pwButtonText}</button>
        {/if}
    </div>
</div>
