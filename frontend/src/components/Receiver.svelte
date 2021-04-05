<script>
    import Secret from "./Secret.svelte";
    let secDict = {
        Token: "",
        Value: "",
        Keyphrase: false,
    };

    // let socket = new WebSocket("ws://127.0.0.1:8080/ws");
    let socket = new WebSocket("ws://192.168.1.22:8080/ws");
    let btsDisabled = true;
    socket.onopen = () => {
        console.log("connected...");
    };

    socket.onmessage = (resp) => {
        let secVal = resp.data;
        try {
            secDict = JSON.parse(secVal);
            btsDisabled = false;
        } catch (error) {
            secDict.Token = secVal;
        }
    };
</script>

<div class="column is-full">
    <h2 class="title" style="display:inline">Unique token:</h2>
    <h2
        class="title is-family-monospace"
        style="display:inline; color:#fa7d09;"
    >
        {secDict.Token.toUpperCase()}
    </h2>
</div>
<div class="column is-full">
    <Secret bind:secValue={secDict.Value} {btsDisabled} />
</div>
