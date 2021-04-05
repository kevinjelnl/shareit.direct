<script>
    import CopyToClipboard from "svelte-copy-to-clipboard";
    import ShowSecretButton from "./reuseables/ShowSecretButton.svelte";

    export let secValue;
    export let btsDisabled = true;
    let secField;
    let copiedMessage = "";

    const handleCopied = (e) => {
        copiedMessage = "Secret copied!";
    };

    const invalidCopied = () => {
        copiedMessage = "Could not copy the secret!";
    };
</script>

<div class="field">
    <input
        bind:this={secField}
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
            disabled={btsDisabled}
        >
            copy secret
        </button>
    </CopyToClipboard>

    <ShowSecretButton {secField} disableButton={btsDisabled} />
</div>
