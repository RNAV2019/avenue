<script lang="ts">
	import { goto } from '$app/navigation';
	import type { User } from '@supabase/supabase-js';
	import { Link, MousePointerClick } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import Button from '../components/Button.svelte';
	import { processDataForLineGraph } from '$lib/graph';
	import type { Statistic } from '$lib/helper';
	import {
		type ChartItem,
		Chart,
		TimeScale,
		LinearScale,
		controllers,
		PointElement,
		LineElement
	} from 'chart.js';
	import 'chartjs-adapter-moment';
	export let data;
	let { supabase, session, aggregateClicks, avenue, statistics } = data;
	$: ({ supabase } = data);
	let user: User | null = session?.user!;
	let token: string | undefined = session?.access_token;

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

	// Function to create the line graph using Chart.js
	function createLineGraph(data: Statistic[]) {
		const { dates, clickData } = processDataForLineGraph(data);

		const ctx = document.getElementById('lineGraph') as ChartItem;
		new Chart(ctx, {
			type: 'line',
			data: {
				labels: dates,
				datasets: [
					{
						label: 'Number of Clicks',
						data: clickData,
						borderColor: 'rgba(75, 192, 192, 1)',
						borderWidth: 2,
						fill: false
					}
				]
			},
			options: {
				responsive: true,
				scales: {
					x: {
						type: 'time',
						time: {
							unit: 'minute'
							// displayFormats: {
							// 	day: 'MMM D'
							// }
						}
					},
					y: {
						beginAtZero: true,
						title: {
							display: true,
							text: 'Number of Clicks'
						}
					}
				},
				plugins: {
					title: {
						display: true,
						text: 'Clicks Over Time'
					}
				}
			}
		});
	}

	onMount(async () => {
		console.log(user);
		console.log(token);
		console.log('This is the correct stats');
		console.log(statistics);

		console.log(aggregateClicks);
		if (token) {
			var header = new Headers();
			header.append('Authorization', `Bearer ${token}`);
			var requestOptions: RequestInit = {
				method: 'POST',
				headers: header,
				redirect: 'follow'
			};
			const res = await fetch('http://localhost:3000/avenue/create', requestOptions);
			if (res.ok) {
				console.log(await res.json());
			} else {
				console.log(res.statusText);
			}
		}
		Chart.register(TimeScale, PointElement, LinearScale, controllers.LineController, LineElement);
		createLineGraph(statistics);
	});

	async function logout() {
		await supabase.auth.signOut();
		user = (await supabase.auth.getUser()).data.user;
		// user.set(null);
		// userToken.set(undefined);
		console.log('Signed the user out on button press event');
	}
</script>

<main class="p-10 space-y-10">
	<nav class="flex flex-row items-center justify-between gap-3">
		<div class="flex flex-row items-center gap-1 text-xl font-bold">
			<h1>Avenue</h1>
			{#if user}
				<span>- logged in as {user.email}</span>
			{/if}
		</div>
		<div class="flex flex-row items-center justify-center gap-8">
			{#if user}
				<Button class="h-12 w-44" color="green" on:click={() => logout()}>Sign Out</Button>
				<img
					src={user?.user_metadata.avatar_url}
					alt="Profile"
					class="w-12 h-12 border-2 border-black rounded-full"
					referrerpolicy="no-referrer"
				/>
			{:else}
				<Button class="h-12 w-44" on:click={() => loginWithGoogle()}>Continue with Google</Button>
			{/if}
		</div>
	</nav>
	{#if user}
		<div
			class="flex flex-col w-full gap-4 p-6 bg-red-300 border-2 border-black shadow-brutal grainy"
		>
			<div class="flex flex-row items-center justify-between">
				<h1 class="ml-2 text-xl font-medium">Dashboard</h1>
				<Button
					class="h-12 w-44"
					color="indigo"
					on:click={() => {
						goto(`/avenues/${user?.id}`);
					}}
					>Visit my avenue
				</Button>
			</div>
			<section class="flex flex-row justify-between gap-3">
				<article
					class="flex items-center flex-1 gap-4 p-6 bg-orange-500 border-2 border-black shadow-brutal grainy sm:justify-between"
				>
					<span
						class="p-3 text-black border-2 border-black rounded-full bg-amber-500 grainy sm:order-last"
					>
						<MousePointerClick />
					</span>

					<div>
						<p class="text-3xl font-bold text-gray-900">{aggregateClicks}</p>

						<p class="text-sm font-medium">Total Avenue Clicks</p>
					</div>
				</article>
				<article
					class="flex items-center flex-1 gap-4 p-6 bg-orange-500 border-2 border-black shadow-brutal grainy sm:justify-between"
				>
					<span
						class="p-3 text-black border-2 border-black rounded-full bg-amber-500 grainy sm:order-last"
					>
						<Link />
					</span>

					<div>
						<p class="text-3xl font-bold text-gray-900">{avenue.links?.length}</p>

						<p class="text-sm font-medium">Number of links</p>
					</div>
				</article>
			</section>
			<section>
				<canvas id="lineGraph" width="400" height="200" />
			</section>
		</div>
	{/if}
</main>
