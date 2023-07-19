<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from './Button.svelte';
	import SubmitButton from './SubmitButton.svelte';
	import type { Avenue, Link } from '$lib/helper';

	const dispatch = createEventDispatcher();
	export let avenue: Avenue | undefined;
	export let userToken: string | undefined;
	export let link: Link | undefined;
	const close = () => {
		dispatch('close');
	};
	let linkName: string = link?.description!;
	let urlString: string = link?.url!;

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

	const updateLink = async () => {
		var header = new Headers();
		header.append('Authorization', `Bearer ${userToken}`);
		header.append('Content-Type', 'application/json');
		console.log('User Token' + userToken);
		console.log('AvenueID' + avenue?.ID);

		const valid = validation();
		if (valid == false) {
			return;
		}
		let updatedLink: Link = {
			ID: link?.ID,
			url: urlString!,
			description: linkName,
			avenue_id: avenue?.ID
		};
		console.log(updatedLink);

		var requestOptions: RequestInit = {
			method: 'PATCH',
			headers: header,
			body: JSON.stringify(updatedLink),
			redirect: 'follow'
		};

		const res = await fetch(`http://localhost:3000/links/update`, requestOptions);

		if (res.ok) {
			console.log('Update the link succesfully');
			close();
		} else {
			throw new Error('Unable to update link' + res.statusText);
		}
	};
</script>

<div class="fixed top-0 left-0 z-50 w-full h-full bg-black bg-opacity-80" on:close={close}>
	<div class="max-w-lg mx-auto bg-orange-500 rounded-lg my-52 dark:bg-secondary shadow-brutal">
		<form
			class="flex flex-col items-center w-full h-full p-8 grainy gap-4"
			name="updateForm"
			on:submit|preventDefault={updateLink}
		>
			<h3 class="mb-2 text-xl font-medium">Update Link</h3>
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
				<SubmitButton class="w-32 h-10 text-xs" color={'fuchsia'}>Update</SubmitButton>
				<Button class="w-32 h-10 text-xs" color={'red'} on:click={close}>Close</Button>
			</div>
		</form>
	</div>
</div>
