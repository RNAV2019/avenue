<script lang="ts">
	import type { Avenue } from '$lib/types';
	import { userToken } from '$lib/store';
	import { initializeApp } from 'firebase/app';
	import { firebaseConfig } from '../config/firebase/firebase-config';
	import {
		GoogleAuthProvider,
		browserSessionPersistence,
		getAuth,
		onAuthStateChanged,
		setPersistence,
		signInWithPopup,
		signOut,
		type User
	} from 'firebase/auth';
	import { goto } from '$app/navigation';

	let user: User | null;
	let avenue: Avenue | null;

	// Initialize Firebase
	const app = initializeApp(firebaseConfig);

	// Initialize Firebase Authentication and get a reference to the service
	const auth = getAuth(app);

	async function loginWithGoogle() {
		setPersistence(auth, browserSessionPersistence);
		const userCred = await signInWithPopup(auth, new GoogleAuthProvider());
		userToken.set(await userCred.user.getIdToken());
		var header = new Headers();
		header.append('Authorization', `Bearer ${$userToken}`);

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
		user = null;
		userToken.set(null);
		avenue = null;
		console.log('Signed the user out on button press event');
	}

	onAuthStateChanged(auth, async (userState) => {
		if (userState) {
			// User is signed in
			user = userState;
			userToken.set(await userState.getIdToken());
		} else {
			// User is signed out
		}
	});
</script>

<main class="space-y-10 p-10">
	<nav class="flex flex-row gap-3 justify-between items-center">
		<div class="flex flex-row gap-1 text-lg">
			<h1>Avenue</h1>
			{#if user}
				<span>- logged in as {user.displayName}</span>
			{/if}
		</div>
		<div class="flex flex-row gap-4 items-center">
			{#if user}
				<button
					class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
					on:click={() => logout()}>Sign Out</button
				>
				<img src={user.photoURL} alt="Profile" class="rounded-full w-9 h-9" />
			{:else}
				<button
					class="rounded-md text-bold text-white bg-black px-4 py-2"
					on:click={() => loginWithGoogle()}>Continue with Google</button
				>
			{/if}
		</div>
	</nav>
	{#if user}
		<button
			class="rounded-md text-bold text-white bg-black px-4 py-2"
			on:click={() => {
				goto(`/avenues/${user?.uid}`);
			}}>Visit my avenue</button
		>
	{/if}
	{$userToken}
</main>
