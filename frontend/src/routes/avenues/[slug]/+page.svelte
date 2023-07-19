<script lang="ts">
	import { color, type Avenue, type Link } from '$lib/helper.js';
	import { PenLine, Trash2 } from 'lucide-svelte';
	import Button from '../../../components/Button.svelte';
	import CreateModal from '../../../components/CreateModal.svelte';
	import UpdateModal from '../../../components/UpdateModal.svelte';
	import { onMount } from 'svelte';
	import Cookie from 'js-cookie';

	export let data;
	const urlRegex = /^(https?:\/\/)/;
	let avenue: Avenue | undefined = data.avenue;
	let links: Link[] = data.avenue?.links!;
	console.log(links);
	let picture: string | undefined = data.userInfo.profile_picture;
	let name: string | undefined = data.userInfo.name;
	let userToken: string | undefined = data.session?.access_token;
	let isOwner: boolean = data.isOwner;

	let showCreateModal = false;
	let showUpdateModal = false;
	let updateModalIndex: number | undefined = 0;

	const close = async () => {
		reloadLinks();
		showCreateModal = false;
		showUpdateModal = false;
	};

	async function reloadLinks() {
		const requestOptions: RequestInit = {
			method: 'GET'
		};
		const res = await fetch(`http://localhost:3000/avenue/find/${data.slug}`, requestOptions);

		if (res.ok) {
			let avenue = (await res.json()) as Avenue;
			links = avenue?.links!;
		} else {
			throw new Error('Links not found');
		}
	}

	onMount(() => {
		navigator.geolocation.getCurrentPosition((position) => {
			const latitude = position.coords.latitude;
			const longitude = position.coords.longitude;
			Cookie.set('location', `${latitude},${longitude}`);
		});
	});

	function getRandomColor(): color {
		return color[Math.floor(Math.random() * color.length)] as color;
	}

	async function deleteLink(id: number | undefined) {
		var header = new Headers();
		header.append('Authorization', `Bearer ${userToken}`);
		var requestOptions: RequestInit = {
			method: 'DELETE',
			headers: header,
			redirect: 'follow'
		};
		const res = await fetch(`http://localhost:3000/links/delete/${id}`, requestOptions);
		reloadLinks();
	}
</script>

<main class="flex flex-col items-center justify-center h-full gap-1">
	<img
		src={picture}
		alt="Profile"
		class="w-20 h-20 bg-black border-2 border-black rounded-full"
		referrerpolicy="no-referrer"
	/>
	<h1 class="mt-3 text-xl font-bold">{name}'s Avenue</h1>
	<h3 class="mb-3">{avenue?.description}</h3>
	<div class="flex flex-col items-center gap-4">
		{#if isOwner}
			<Button color="amber" class="h-12 w-52" on:click={() => (showCreateModal = true)}
				>Create Link</Button
			>
		{/if}

		{#if links}
			{#if isOwner}
				{#each links as link (link.ID)}
					<div class="flex flex-row items-center justify-center gap-1 group">
						<div class="hidden px-8 py-1 group-hover:block">
							<button
								class="p-2 rounded-full"
								on:click={() => {
									updateModalIndex = link.ID;
									showUpdateModal = true;
									console.log(avenue);
									console.log(updateModalIndex);
								}}
							>
								<PenLine size={20} />
							</button>
						</div>
						<a href={urlRegex.test(link.url) ? link.url : 'https://' + link.url}>
							<Button color={getRandomColor()} class="h-12 w-52">{link.description}</Button>
						</a>
						<div class="hidden px-8 py-1 group-hover:block">
							<button class="p-2 rounded-full" on:click={() => deleteLink(link.ID)}>
								<Trash2 size={20} />
							</button>
						</div>
					</div>
				{/each}
			{:else}
				{#each links as link (link.ID)}
					<a href={urlRegex.test(link.url) ? link.url : 'https://' + link.url}>
						<Button color={getRandomColor()} class="h-12 w-52">{link.description}</Button>
					</a>
				{/each}
			{/if}
		{/if}
	</div>
	<p class="mt-4">
		{#if isOwner}
			You are the owner of this avenue
		{:else}
			You are not the owner of this avenue
		{/if}
	</p>
</main>
{#if showCreateModal}
	<CreateModal on:close={close} {avenue} {userToken} />
{/if}
{#if showUpdateModal}
	<UpdateModal
		on:close={close}
		{avenue}
		{userToken}
		link={links.find((value) => value.ID == updateModalIndex)}
	/>
{/if}
