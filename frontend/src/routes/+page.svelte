<script lang="ts">
	import { goto } from '$app/navigation';
	import { auth } from '$lib/firebase';
	import type { Avenue } from '$lib/helper';
	import { authStore } from '$lib/stores';
	import {
		GoogleAuthProvider,
		browserSessionPersistence,
		setPersistence,
		signInWithPopup,
		signOut
	} from 'firebase/auth';

	let avenue: Avenue | null;

	async function loginWithGoogle() {
		setPersistence(auth, browserSessionPersistence);
		const userCred = await signInWithPopup(auth, new GoogleAuthProvider());
		var header = new Headers();
		header.append('Authorization', `Bearer ${await userCred.user.getIdToken()}`);

		var requestOptions: RequestInit = {
			method: 'POST',
			headers: header,
			redirect: 'follow'
		};

		const res = await fetch('http://localhost:3000/user/create', requestOptions);
		console.log(await res.json());
	}

	async function logout() {
		await signOut(auth);
		// user.set(null);
		// userToken.set(undefined);
		avenue = null;
		console.log('Signed the user out on button press event');
	}
</script>

<main class="space-y-10 p-10">
	<nav class="flex flex-row gap-3 justify-between items-center">
		<div class="flex flex-row gap-1 text-2xl font-medium items-center">
			<h1>Avenue</h1>
			{#if $authStore.currentUser}
				<span>- logged in as {$authStore.currentUser.displayName}</span>
			{/if}
		</div>
		<div class="flex flex-row gap-4 items-center">
			{#if $authStore.currentUser}
				<button
					class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
					on:click={() => logout()}>Sign Out</button
				>
				<img src={$authStore.currentUser.photoURL} alt="Profile" class="rounded-full w-9 h-9" />
			{:else}
				<button
					class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
					on:click={() => loginWithGoogle()}>Continue with Google</button
				>
			{/if}
		</div>
	</nav>
	{#if $authStore.currentUser}
		<button
			class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
			on:click={() => {
				goto(`/avenues/${$authStore.currentUser?.uid}`);
			}}>Visit my avenue</button
		>
	{/if}
	{$authStore.currentUser}
</main>
