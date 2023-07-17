<script lang="ts">
	import { authStore } from '$lib/stores';
	import { initializeApp } from 'firebase/app';
	import { getAuth, onAuthStateChanged } from 'firebase/auth';
	import { onMount } from 'svelte';
	import '../app.css';
	import { firebaseConfig } from '../config/firebase/firebase-config';
	// Initialize Firebase
	const app = initializeApp(firebaseConfig);

	// Initialize Firebase Authentication and get a reference to the service
	const auth = getAuth(app);

	onMount(async () => {
		onAuthStateChanged(auth, async (userState) => {
			authStore.update((curr) => {
				return { ...curr, isLoading: true, currentUser: userState };
			});
		});
	});
</script>

<slot />

<style lang="postcss">
	:global(html) {
		width: 100%;
		height: 100%;
		@apply bg-amber-400;
	}

	:global(body) {
		width: 100%;
		height: 100%;
		@apply bg-grainy;
	}
</style>
