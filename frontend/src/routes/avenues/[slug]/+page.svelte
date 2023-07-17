<script lang="ts">
	import type { Avenue, Link, User } from '$lib/helper.js';
	import { onMount } from 'svelte';
	import Button from '../../../components/Button.svelte';
	import CreateModal from '../../../components/CreateModal.svelte';
	import { onAuthStateChanged } from 'firebase/auth';
	import { authStore } from '$lib/stores';
	import { auth } from '$lib/firebase';

	export let data;
	const urlRegex = /^(https?:\/\/)/;
	let avenue: Avenue | undefined = data.avenue;
	let links: Link[] = data.avenue?.links!;
	let picture: string | undefined = data.userInfo.profile_picture;
	let name: string | undefined = data.userInfo.name;
	let isOwner: boolean = false;

	onMount(() => {
		onAuthStateChanged(auth, (userState) => {
			if (data.slug == userState?.uid) {
				isOwner = true;
			}
		});
	});

	let showModal = false;

	const close = async () => {
		const requestOptions: RequestInit = {
			method: 'GET'
		};
		const res = await fetch(`http://localhost:3000/avenue/find/${data.slug}`, requestOptions);

		if (res.ok) {
			let user = (await res.json()) as User;
			avenue = user.avenue;
			links = avenue?.links!;
		} else {
			throw new Error('Links not found');
		}
		showModal = false;
	};
</script>

<main class="flex flex-col gap-1 items-center justify-center h-full">
	<img src={picture} alt="Profile" class="h-20 w-20 rounded-full" referrerpolicy="no-referrer" />
	<h1 class="font-bold text-xl mt-3">{name}'s Avenue</h1>
	<h3 class="mb-3">{avenue?.description}</h3>
	<div class="flex flex-col gap-4">
		{#if isOwner}
			<Button color="amber" class="w-52 h-12" on:click={() => (showModal = true)}
				>Create Link</Button
			>
		{/if}

		{#each links as link (link.ID)}
			<a href={urlRegex.test(link.url) ? link.url : 'https://' + link.url}
				><Button color="red" class="w-52 h-12">{link.description}</Button></a
			>
		{/each}
	</div>
	{#if isOwner}
		You are the owner of this avenue
	{:else}
		You are not the owner of this avenue
	{/if}
</main>
<CreateModal show={showModal} on:close={close} {avenue} />
