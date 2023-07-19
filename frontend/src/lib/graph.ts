import moment from 'moment';
import type { ClickCounts, Statistic } from './helper';

export function processDataForLineGraph(data: Statistic[]) {
	// Assuming `data` is an array of statistics objects
	const clickCounts: ClickCounts = {};

	data.forEach((statistic: Statistic) => {
		// Group the data by click timestamp and count the number of clicks on each date
		const clickDate = moment.utc(statistic.ClickTimestamp).toISOString();
		console.log(statistic.ClickTimestamp);
		clickCounts[clickDate] = clickCounts[clickDate] + 1 || 1;
		console.log(clickCounts);
	});

	// Extract the dates and click counts for the line graph
	const dates = Object.keys(clickCounts);
	const clickData = Object.values(clickCounts);

	console.log(dates, clickData);
	console.log(data);

	// Return the data for the line graph
	return { dates, clickData };
}
