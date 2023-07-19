import type { Avenue, Statistic } from '$lib/helper';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, data }) => {
	type StatisticJSON = {
		statistics: Statistic[];
		clicks: number;
	};

	const requestOptions: RequestInit = {
		method: 'GET'
	};
	let avenue: Avenue;
	let aggregateClicks: number;
	let statistics: Statistic[];
	const userID = data.session.user.id;
	const resFind = await fetch(`http://localhost:3000/avenue/find/${userID}`, requestOptions);

	if (resFind.ok) {
		avenue = (await resFind.json()) as Avenue;
	} else {
		throw new Error('Avenue not found');
	}

	const statisticsRes = await fetch(
		`http://localhost:3000/dashboard/all/${avenue.ID}`,
		requestOptions
	);

	if (statisticsRes.ok) {
		const statisticJSON = (await statisticsRes.json()) as StatisticJSON;
		statistics = statisticJSON.statistics;
		aggregateClicks = statisticJSON.clicks;
	} else {
		throw new Error('Statistics not found');
	}
	return {
		avenue,
		aggregateClicks,
		statistics
	};
};
