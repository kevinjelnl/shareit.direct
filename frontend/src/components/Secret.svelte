<script>
    import CopyToClipboard from "svelte-copy-to-clipboard";

    export let secValue;
    export let btsDisabled;
    let pwbox;
    let copiedMessage = "";
    let showHideSecret= "show secret";

    const show = () => {
        if (pwbox.type == "password") {
            pwbox.type = "text";
            showHideSecret= "hide secret";
        } else {
            pwbox.type = "password";
            showHideSecret= "show secret";
        }
    };

    const handleCopied = (e) => {
        copiedMessage = "Secret copied!";
    };

    const invalidCopied = () => {
        copiedMessage = "Could not copy the secret!";
    };
</script>

<div class="field">
    <input
        bind:this={pwbox}
        class="input"
        type="password"
        placeholder="Your secret will arrive here"
        value={secValue}
        disabled={btsDisabled}
    />
    <p>{copiedMessage}</p>
    <span class="icon is-small is-left"><i class="fas fa-lock" /> </span>
    <br />
    <CopyToClipboard
        text={secValue}
        on:copy={handleCopied}
        on:fail={invalidCopied}
        let:copy
    >
        <button
            class="button is-success"
            on:click={copy}
            disabled={btsDisabled}> copy secret </button>
    </CopyToClipboard>
    <button class="button is-danger" on:click={show} disabled={btsDisabled}>
       {showHideSecret} 
    </button>
</div>
