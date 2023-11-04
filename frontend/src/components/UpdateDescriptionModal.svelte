<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from './Button.svelte';
	import SubmitButton from './SubmitButton.svelte';
	import type { Avenue, Description } from '$lib/helper';
	import { PUBLIC_API_ENDPOINT_URL } from '$env/static/public';

	const dispatch = createEventDispatcher();
	export let avenue: Avenue | undefined;
	export let userToken: string | undefined;
	const close = () => {
		dispatch('close');
	};
	let updateDescription: string = avenue?.description!;

	function validation() {
		if (updateDescription == null || updateDescription == '') {
			return false;
		}
		if (typeof avenue == undefined) {
			return false;
		}
		return true;
	}

	const updateDescriptionFunction = async () => {
		var header = new Headers();
		header.append('Authorization', `Bearer ${userToken}`);
		header.append('Content-Type', 'application/json');

		const valid = validation();
		if (valid == false) {
			return;
		}
		let description: Description = {
			description: updateDescription
		};

		var requestOptions: RequestInit = {
			method: 'PATCH',
			headers: header,
			body: JSON.stringify(description),
			redirect: 'follow'
		};

		const res = await fetch(`${PUBLIC_API_ENDPOINT_URL}/avenue/update`, requestOptions);

		if (res.ok) {
			close();
		} else {
			throw new Error('Unable to update description' + res.statusText);
		}
	};
</script>

<div class="fixed top-0 left-0 z-50 w-full h-full bg-black bg-opacity-80" on:close={close}>
	<div class="max-w-lg mx-auto rounded-lg my-52 bg-cyan-500 dark:bg-secondary shadow-brutal">
		<form
			class="flex flex-col items-center w-full h-full gap-4 p-8 grainy"
			name="createForm"
			on:submit|preventDefault={updateDescriptionFunction}
		>
			<h3 class="mb-2 text-xl font-medium">Update Description</h3>
			<input
				type="text"
				id="name"
				name="name"
				placeholder="Edit description"
				bind:value={updateDescription}
				class="w-2/3 p-2 px-3 text-sm border-2 border-black rounded-sm outline-none shadow-brutal placeholder:text-black"
			/>
			<div class="flex flex-row gap-8">
				<SubmitButton class="w-32 h-10 text-xs" color={'emerald'}>Create</SubmitButton>
				<Button class="w-32 h-10 text-xs" color={'red'} on:click={close}>Close</Button>
			</div>
		</form>
	</div>
</div>
