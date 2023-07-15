import type { Avenue, User } from '$lib/types';

export const load = async ({ params, fetch }) => {
	const requestOptions: RequestInit = {
		method: 'GET'
	};
	let avenue: Avenue | undefined;

	const res = await fetch(`http://localhost:3000/avenue/find/${params.slug}`, requestOptions);

	if (res.ok) {
		avenue = ((await res.json()) as User).avenue;
	} else {
		throw new Error('Avenue not found');
	}

	return {
		slug: params.slug,
		avenue: avenue
	};
};
