<script>
    // https://jsfiddle.net/42kd06xr/
    // https://kodeazy.com/javascript-cryptojs-example/
    //stackoverflow.com/questions/18279141/javascript-string-encryption-and-decryption

    // Encrypt
    var ciphertext = CryptoJS.AES.encrypt("my message", "secret key 123");
    // Decrypt
    let foo = ciphertext.toString();
    console.log(foo);

    var bytes = CryptoJS.AES.decrypt(foo, " key 123");
    var second = CryptoJS.AES.decrypt(foo, "secret key 123");
    console.log(bytes.toString(CryptoJS.enc.Utf8));
    console.log(second.toString(CryptoJS.enc.Utf8));
</script>

<!-- keyphrase -->
<!-- <button>foobar</button> -->
<!-- <p>&#128274;</p> -->
