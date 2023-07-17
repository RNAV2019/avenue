<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from './Button.svelte';
	import SubmitButton from './SubmitButton.svelte';
	import { auth } from '$lib/firebase';
	import type { Avenue, Link } from '$lib/helper';

	let userToken: string | undefined;
	auth.onIdTokenChanged(async (token) => (userToken = await token?.getIdToken()));

	const dispatch = createEventDispatcher();
	export let show: boolean;
	export let avenue: Avenue | undefined;
	const close = () => {
		show = false;
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
		console.log('User Token' + userToken);
		console.log('AvenueID' + avenue?.ID);

		const valid = validation();
		if (valid == false) {
			return;
		}
		let link: Link = {
			url: urlString,
			description: linkName,
			avenue_id: avenue?.ID
		};
		console.log(link);

		var requestOptions: RequestInit = {
			method: 'POST',
			headers: header,
			body: JSON.stringify(link),
			redirect: 'follow'
		};

		const res = await fetch(`http://localhost:3000/links/create`, requestOptions);

		if (res.ok) {
			console.log('Created the link succesfully');
			close();
		} else {
			throw new Error('Unable to create link' + res.statusText);
		}
	};
</script>

{#if show}
	<div class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-80 z-50" on:close={close}>
		<div class="mx-auto my-52 max-w-lg bg-cyan-500 dark:bg-secondary rounded-lg shadow-brutal">
			<form
				class="grainy w-full h-full p-8 flex flex-col gap-4 items-center"
				name="createForm"
				on:submit|preventDefault={createLink}
			>
				<h3 class="text-xl font-medium mb-2">Create Link</h3>
				<input
					type="text"
					id="name"
					name="name"
					placeholder="Link name"
					bind:value={linkName}
					class="p-2 w-2/3 px-3 outline-none rounded-sm border-2 border-black shadow-brutal text-sm placeholder:text-black"
				/>
				<input
					type="text"
					id="url"
					name="url"
					placeholder="URL e.g www.google.com"
					bind:value={urlString}
					class="p-2 w-2/3 px-3 outline-none rounded-sm border-2 border-black shadow-brutal text-sm placeholder:text-black mb-2"
				/>
				<div class="flex flex-row gap-8">
					<SubmitButton class="w-32 h-10 text-xs" color={'emerald'}>Create</SubmitButton>
					<Button class="w-32 h-10 text-xs" color={'red'} on:click={close}>Close</Button>
				</div>
			</form>
		</div>
	</div>
{/if}
