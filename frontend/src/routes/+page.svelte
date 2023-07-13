<script lang="ts">
	import { initializeApp, type FirebaseOptions } from 'firebase/app';
	import {
		GoogleAuthProvider,
		browserSessionPersistence,
		getAuth,
		onAuthStateChanged,
		setPersistence,
		signInWithPopup,
		type UserCredential,
		type User
	} from 'firebase/auth';
	import { onMount } from 'svelte';

	// TODO: Replace the following with your app's Firebase project configuration
	// See: https://firebase.google.com/docs/web/learn-more#config-object
	const firebaseConfig: FirebaseOptions = {
		apiKey: 'AIzaSyDRgd0yNpM_ziPJmkex9Y-w0jRGs8sp8kw',
		authDomain: 'avenue-f8597.firebaseapp.com',
		projectId: 'avenue-f8597',
		storageBucket: 'avenue-f8597.appspot.com',
		messagingSenderId: '992541868817',
		appId: '1:992541868817:web:b28df303c2cb6ea198c2b0',
		measurementId: 'G-LRK89ZFGY5'
	};

	let token: string = '';
	let user: User;

	// Initialize Firebase
	const app = initializeApp(firebaseConfig);

	// Initialize Firebase Authentication and get a reference to the service
	const auth = getAuth(app);

	async function loginWithGoogle() {
		setPersistence(auth, browserSessionPersistence);
		const userCred = await signInWithPopup(auth, new GoogleAuthProvider());
		token = await userCred.user.getIdToken();
		var header = new Headers();
		header.append('Content-Type', 'application/json');

		var raw = JSON.stringify({
			token: token
		});

		var requestOptions: RequestInit = {
			method: 'POST',
			headers: header,
			body: raw,
			redirect: 'follow'
		};

		const res = await fetch('http://localhost:3000/user/create', requestOptions);
		console.log(await res.json());
	}

	onMount(() => {
		const loginToken = localStorage.getItem('token');
	});

	onAuthStateChanged(auth, async (userState) => {
		if (userState) {
			// User is signed in
			console.log('user is signed in');
			user = userState;
			token = await user.getIdToken();
		} else {
			// User is signed out
			// ...
			console.log('user is signed out');
		}
	});
</script>

<main class="space-y-20 p-20">
	<h1>Create account with Google</h1>
	<button
		class="rounded-md text-bold text-white bg-black px-4 py-2"
		on:click={() => loginWithGoogle()}>Continue with Google</button
	>
	{#if user}
		Logged In with JWT Token: {token}
	{/if}
</main>
