<script>
    import Keyphrase from "./Keyphrase.svelte";
    import ShowSecretButton from "./reuseables/ShowSecretButton.svelte";

    const params = new URLSearchParams(window.location.search);
    // https://www.npmjs.com/package/crypto-js
    let token = "";
    let secField;
    let success;
    let shareDisabled = false;
    let sharedMessage = "";

    if (params.get("token")) {
        token = params.get("token");
    }
    $: tlen = token.length;

    let keyphrase = false;
    let value = "";

    // functions that shares the supplied secret
    async function shareSecret() {
        console.log(value);
        const res = await fetch("http://192.168.1.22:8080/supply", {
            // const res = await fetch("http://127.0.0.1:8080/supply", {
            method: "POST",
            body: JSON.stringify({ token, value, keyphrase }),
        });
        success = await res.json();
        console.log(success);
        if (!success) {
            alert("invalid token supplied");
            return;
        } else {
            console.log("Token supplied");
            sharedMessage = "Secret supplied!";
        }
    }

    function foo() {
        alert("foobar");
    }
</script>

<!-- icons -->
<!-- https://www.compart.com/en/unicode/U+1F513 -->

<div class="column is-full">
    <h2 class="title" style="display:inline">Supply a secret:</h2>
</div>
<div class="column is-full">
    <p>1. Insert the token</p>
    <p>2. Insert the secret</p>
    <p>3. opt. Use a passphrase for <b>extra security</b>!</p>
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
        {#if tlen === 5}
            <div class="control has-icons-left has-icons-right">
                <input
                    bind:this={secField}
                    class="input"
                    type="password"
                    bind:value
                />
                <span class="icon is-small is-left">
                    <p class="lock-icon" style="cursor: pointer" on:click={foo}>
                        &#128274;
                    </p>
                </span>
            </div>
            <p>{sharedMessage}</p>
            <Keyphrase />
            <br />
            <button
                class="button is-success"
                on:click={shareSecret}
                disabled={shareDisabled}>share</button
            >
            <ShowSecretButton buttonValue="show secret" {secField} />
        {/if}
    </div>
</div>
