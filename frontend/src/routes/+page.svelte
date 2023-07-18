<script lang="ts">
	import { goto } from '$app/navigation';
	import type { User } from '@supabase/supabase-js';
	export let data;
	let { supabase, session } = data;
	$: ({ supabase } = data);
	let user: User | null = session?.user!;

	async function loginWithGoogle() {
		let { data } = await supabase.auth.signInWithOAuth({
			provider: 'google'
		});
		console.log('Logged in successfully with google');
		console.log(`Logged in with ${data.provider} at url ${data.url}`);

		// var header = new Headers();
		// header.append('Authorization', `Bearer ${await userCred.user.getIdToken()}`);

		// var requestOptions: RequestInit = {
		// 	method: 'POST',
		// 	headers: header,
		// 	redirect: 'follow'
		// };

		// const res = await fetch('http://localhost:3000/user/create', requestOptions);
		// console.log(await res.json());
	}

	async function getDetails() {
		user = (await supabase.auth.getUser()).data.user;
		console.log(user?.email);
		console.log(user?.user_metadata.avatar_url);
	}

	async function logout() {
		await supabase.auth.signOut();
		user = (await supabase.auth.getUser()).data.user;
		// user.set(null);
		// userToken.set(undefined);
		console.log('Signed the user out on button press event');
	}
</script>

<main class="space-y-10 p-10">
	<nav class="flex flex-row gap-3 justify-between items-center">
		<div class="flex flex-row gap-1 text-2xl font-medium items-center">
			<h1>Avenue</h1>
			{#if user}
				<span>- logged in as {user.email}</span>
			{/if}
		</div>
		<div class="flex flex-row gap-4 items-center">
			{#if user}
				<button
					class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
					on:click={() => logout()}>Sign Out</button
				>
				<img src={user?.user_metadata.avatar_url} alt="Profile" class="rounded-full w-9 h-9" />
			{:else}
				<button
					class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
					on:click={() => loginWithGoogle()}>Continue with Google</button
				>
			{/if}
		</div>
	</nav>
	{#if user}
		<button
			class="rounded-md font-bold text-sm text-white bg-black px-4 py-2"
			on:click={() => {
				goto(`/avenues/${user?.id}`);
			}}>Visit my avenue</button
		>
	{/if}
	{user}
</main>
