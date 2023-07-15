<script lang="ts">
	import type { Avenue, User } from '$lib/types.js';
	import { onMount } from 'svelte';

	export let data;
	let avenue: Avenue | undefined = data.avenue;
</script>

<h1>Hello {data.slug}</h1>
<h3>This is the data {avenue?.title}</h3>
