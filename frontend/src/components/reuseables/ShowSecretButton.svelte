<script>
    export let secField;
    export let disableButton = false;
    let buttonValue = "show secret";
    let classValue = ["is-danger", "is-warning"];
    let classChange = classValue[0];

    export function toggleValue() {
        if (secField.type == "password") {
            secField.type = "text";
            buttonValue = "hide secret";
            classChange = classValue[1];
        } else {
            secField.type = "password";
            buttonValue = "show secret";
            classChange = classValue[0];
        }
    }
</script>

<!-- <p>&#128275;</p> -->
<button
    class="button {classChange}"
    on:click={toggleValue}
    disabled={disableButton}
>
    {buttonValue}
</button>
