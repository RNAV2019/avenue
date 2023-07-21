import type { Statistic } from './helper';

export function processDataForLineGraph(data: Statistic[]) {
	type ClickCounts = {
		[date: string]: number;
	};

	const clickCounts: ClickCounts = {};
	data.forEach((statistic) => {
		const clickDate = new Date(statistic.click_timestamp).toUTCString();
		clickCounts[clickDate] = clickCounts[clickDate] + 1 || 1;
	});

	const dates = Object.keys(clickCounts);

	const clickData = dates.map((date) => clickCounts[date]);

	return { dates, clickData };
}
