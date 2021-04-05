<script>
	// imports
	import router from "page";
	import Receiver from "./components/Receiver.svelte";
	import Supplier from "./components/Supplier.svelte";
	import About from "./components/About.svelte";
	import fourOfour from "./components/fourOfour.svelte";
	import { onMount } from "svelte";

	let page;
	let pselected = 0;
	const navPages = [
		{ link: "/", name: "Receive" },
		{ link: "/supply", name: "Supply" },
		{ link: "/about", name: "About" },
	];
	function changeComponent(event) {
		// https://svelte.dev/repl/ca6de655755d4ae88c4c6cbf60f85fda?version=3.23.0
		pselected = event.srcElement.id;
	}

	// https://www.youtube.com/watch?v=eohan_OS8N0&feature=youtu.be
	router("/", () => (page = Receiver));
	router("/supply", () => (page = Supplier));
	router("/about", () => (page = About));
	router("/*", () => (page = fourOfour));
	router.start();

	//! move to own component!
	// highlight the correct page in the menu... hardcoded 'pagenumbers'!
	onMount(() => {
		let pageLanded = window.location.pathname;
		if (pageLanded.startsWith("/supply")) {
			pselected = 1;
		} else if (pageLanded == "/about") {
			pselected = 2;
		}
	});
</script>

<div class="container is-max-desktop">
	<nav class="navbar-end" role="navigation" aria-label="main navigation">
		<div class="navbar-brand" on:click={changeComponent}>
			{#each navPages as pdict, i}
				<a
					class={pselected === i
						? "navbar-item is-active"
						: "navbar-item"}
					href={pdict.link}
					id={i}
				>
					{pdict.name}</a
				>
			{/each}
		</div>
	</nav>
</div>

<div
	class="container is-max-desktop p-5"
	style="background-color: white; border-bottom-right-radius:20px;
	box-shadow: 0 1px 15px 0 rgba(0,0,0,0.2),0 1px 20px 0 rgba(0,0,0,0.19);
	height: 80%;"
>
	<div class="columns is-multiline">
		<!-- <svelte:component this={page} {params} /> -->
		<svelte:component this={page} />
	</div>
</div>
<div class="container is-max-desktop">
	<a
		style="color: #fa7d09;"
		href="https://github.com/kevinjelnl/shareit.direct/"
		target="_blank">See the code on github</a
	>
</div>
