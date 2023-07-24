<script lang="ts">
	import { goto } from '$app/navigation';
	import { processDataForLineGraph } from '$lib/graph';
	import type { Statistic } from '$lib/helper';
	import type { User } from '@supabase/supabase-js';
	import {
		Chart,
		LineElement,
		LinearScale,
		PointElement,
		TimeScale,
		Tooltip,
		controllers,
		type ChartItem
	} from 'chart.js';
	import 'chartjs-adapter-moment';
	import { Link, MousePointerClick } from 'lucide-svelte';
	import moment from 'moment';
	import { onMount } from 'svelte';
	import Button from '../components/Button.svelte';
	import { redirect } from '@sveltejs/kit';
	import { PUBLIC_API_ENDPOINT_URL } from '$env/static/public';
	export let data;
	let { supabase, session, aggregateClicks, avenue, statistics } = data;
	$: ({ supabase } = data);
	let user: User | null = session?.user!;
	let token: string | undefined = session?.access_token;

	async function loginWithGoogle() {
		let { data } = await supabase.auth.signInWithOAuth({
			provider: 'google',
			options: {
				redirectTo: 'https://avenue-alpha.vercel.app/auth/callback'
			}
		});
	}

	// Function to create the line graph using Chart.js
	function createLineGraph(data: Statistic[]) {
		const { dates, clickData } = processDataForLineGraph(data);

		// Convert date strings to Moment.js objects and format them
		const formattedLabels = dates.map((dateString) => {
			const dateMoment = moment(dateString);
			return dateMoment.format('M/D/YYYY, h:mm A');
		});

		const ctx = document.getElementById('lineGraph') as ChartItem;
		new Chart(ctx, {
			type: 'line',
			data: {
				labels: dates,
				datasets: [
					{
						label: 'Number of Clicks',
						data: clickData,
						borderColor: 'rgb(12 74 110)',
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
							unit: 'day',
							displayFormats: {
								day: 'MMM D'
							}
						},
						ticks: {
							color: 'rgba(0, 0, 0, 0.8)' // Change the color of the x-axis
						},
						grid: {
							color: 'rgba(0, 0, 0, 0.15)' // Change the color of the x-axis grid lines
						}
					},
					y: {
						beginAtZero: true,
						title: {
							display: true,
							text: 'Number of Clicks',
							color: 'black'
						},
						ticks: {
							color: 'rgba(0, 0, 0, 0.8)' // Change the color of the y-axis
						},
						grid: {
							color: 'rgba(0, 0, 0, 0.15)' // Change the color of the y-axis grid lines
						}
					}
				},
				plugins: {
					title: {
						display: true,
						text: 'Clicks Over Time'
					},
					tooltip: {
						intersect: false, // Allow multiple tooltips to be shown when overlapping points
						mode: 'index', // Show the tooltip for all data points with the same X-axis value
						callbacks: {
							label: function (tooltipItem) {
								// Customize the tooltip label to show the date and the number of clicks
								const clickCount = tooltipItem.parsed.y;
								return ` ${clickCount} Clicks`;
							}
						}
					}
				}
			}
		});
	}

	onMount(async () => {
		if (token) {
			var header = new Headers();
			header.append('Authorization', `Bearer ${token}`);
			var requestOptions: RequestInit = {
				method: 'POST',
				headers: header,
				redirect: 'follow'
			};
			const res = await fetch(`${PUBLIC_API_ENDPOINT_URL}/avenue/create`, requestOptions);
		}

		Chart.register(
			TimeScale,
			PointElement,
			LinearScale,
			controllers.LineController,
			LineElement,
			Tooltip
		);
		if (statistics != undefined) {
			createLineGraph(statistics);
		}
	});

	async function logout() {
		await supabase.auth.signOut();
		user = (await supabase.auth.getUser()).data.user;
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
			class="flex flex-col w-full gap-4 p-6 border-2 border-black bg-rose-400 shadow-brutal grainy"
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
						<p class="text-3xl font-bold text-gray-900">{avenue?.links?.length}</p>

						<p class="text-sm font-medium">Number of links</p>
					</div>
				</article>
			</section>
			<section>
				<article class="p-6 border-2 border-black bg-sky-500 shadow-brutal grainy">
					<canvas id="lineGraph" width="400" height="200" />
				</article>
			</section>
		</div>
	{/if}
</main>
