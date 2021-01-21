<!-- https://gist.github.com/adhithyan15/4350689 -->
<script>
    let timeValue;
    let seconds = 60;
    let minutes = 5;

    function timer() {
        function tick() {
            // var counter = document.getElementById("counter");
            var current_minutes = minutes - 1;
            seconds--;
            timeValue =
                current_minutes.toString() +
                ":" +
                (seconds < 10 ? "0" : "") +
                String(seconds);
            if (seconds > 0) {
                setTimeout(tick, 1000);
            } else {
                if (mins > 1) {
                    countdown(mins - 1);
                }
            }
        }
        tick();
    }
    $: timer();
</script>

<p>{timeValue}</p>
