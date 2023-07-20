import type { Statistic } from './helper';

export function processDataForLineGraph(data: Statistic[]) {
	type ClickCounts = {
		[date: string]: number;
	};

	// Group the data by click date and count the number of clicks on each date
	const clickCounts: ClickCounts = {};
	data.forEach((statistic) => {
		// const clickDate = new Date(statistic.click_timestamp).toLocaleString(undefined, {
		// 	month: 'numeric',
		// 	day: 'numeric',
		// 	year: 'numeric',
		// 	hour: 'numeric',
		// 	minute: 'numeric',
		// 	hour12: false
		// });
		const clickDate = new Date(statistic.click_timestamp).toUTCString();
		clickCounts[clickDate] = clickCounts[clickDate] + 1 || 1;
	});

	// Extract the unique dates
	const dates = Object.keys(clickCounts);

	// Create an array of click counts corresponding to the dates
	const clickData = dates.map((date) => clickCounts[date]);

	return { dates, clickData };
}
