<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from './Button.svelte';
	import SubmitButton from './SubmitButton.svelte';
	import type { Avenue, Link } from '$lib/helper';
	import { PUBLIC_API_ENDPOINT_URL } from '$env/static/public';

	const dispatch = createEventDispatcher();
	export let avenue: Avenue | undefined;
	export let userToken: string | undefined;
	const close = () => {
		dispatch('close');
	};
	let linkName: string;
	let urlString: string;
	const urlRegex = /^(https?:\/\/)?([\w.-]+\.[a-zA-Z]{2,})(:\d+)?(\/[^\s]*)?$/;

	function validation() {
		if (linkName == null || linkName == '') {
			return false;
		}
		if (urlString == null || linkName == '') {
			return false;
		}
		if (urlRegex.test(urlString) == false) {
			return false;
		}
		if (typeof avenue == undefined) {
			return false;
		}
		return true;
	}

	const createLink = async () => {
		var header = new Headers();
		header.append('Authorization', `Bearer ${userToken}`);
		header.append('Content-Type', 'application/json');

		const valid = validation();
		if (valid == false) {
			return;
		}
		let link: Link = {
			url: urlString,
			description: linkName
		};

		var requestOptions: RequestInit = {
			method: 'POST',
			headers: header,
			body: JSON.stringify(link),
			redirect: 'follow'
		};

		const res = await fetch(`${PUBLIC_API_ENDPOINT_URL}/links/create`, requestOptions);

		if (res.ok) {
			close();
		} else {
			throw new Error('Unable to create link' + res.statusText);
		}
	};
</script>

<div class="fixed top-0 left-0 z-50 w-full h-full bg-black bg-opacity-80" on:close={close}>
	<div class="max-w-lg mx-auto rounded-lg my-52 bg-cyan-500 dark:bg-secondary shadow-brutal">
		<form
			class="flex flex-col items-center w-full h-full gap-4 p-8 grainy"
			name="createForm"
			on:submit|preventDefault={createLink}
		>
			<h3 class="mb-2 text-xl font-medium">Create Link</h3>
			<input
				type="text"
				id="name"
				name="name"
				placeholder="Link name"
				bind:value={linkName}
				class="w-2/3 p-2 px-3 text-sm border-2 border-black rounded-sm outline-none shadow-brutal placeholder:text-black"
			/>
			<input
				type="text"
				id="url"
				name="url"
				placeholder="URL e.g www.google.com"
				bind:value={urlString}
				class="w-2/3 p-2 px-3 mb-2 text-sm border-2 border-black rounded-sm outline-none shadow-brutal placeholder:text-black"
			/>
			<div class="flex flex-row gap-8">
				<SubmitButton class="w-32 h-10 text-xs" color={'emerald'}>Create</SubmitButton>
				<Button class="w-32 h-10 text-xs" color={'red'} on:click={close}>Close</Button>
			</div>
		</form>
	</div>
</div>
