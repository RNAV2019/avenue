import { PUBLIC_API_ENDPOINT_URL } from '$env/static/public';
import type { Avenue, Statistic } from '$lib/helper';
import type { PageLoad } from './avenues/[slug]/$types';

export const load: PageLoad = async ({ fetch, data }) => {
	type DashboardJSON = {
		statistics: Statistic[];
		clicks: number;
		avenue: Avenue;
	};

	const requestOptions: RequestInit = {
		method: 'GET'
	};
	let avenue: Avenue | undefined = undefined;
	let aggregateClicks: number | undefined = undefined;
	let statistics: Statistic[] | undefined = undefined;
	const userID: string | undefined = data.session?.user?.id;

	if (userID != undefined && userID != '') {
		const dashboardRes = await fetch(
			`${PUBLIC_API_ENDPOINT_URL}/dashboard/all/${userID}`,
			requestOptions
		);

		if (dashboardRes.ok) {
			const dashboardJSON = (await dashboardRes.json()) as DashboardJSON;
			statistics = dashboardJSON.statistics;
			aggregateClicks = dashboardJSON.clicks;
			avenue = dashboardJSON.avenue;
		}
	}
	return {
		avenue,
		aggregateClicks,
		statistics
	};
};
