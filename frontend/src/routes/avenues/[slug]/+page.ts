import type { Avenue, UserInfo } from '$lib/helper';
import Cookie from 'js-cookie';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch, data }) => {
	const headers = new Headers();
	const location = Cookie.get('location') || null;
	if (location) {
		headers.append('X-Geo-Location', location);
	}
	const requestOptions: RequestInit = {
		method: 'GET'
	};
	const geoRequestOptions: RequestInit = {
		method: 'GET',
		headers: headers
	};
	let avenue: Avenue;
	let userInfo: UserInfo;
	let isOwner = false;
	const userId: string | null = data.session?.user?.id;

	const res = await fetch(`http://localhost:3000/avenue/find/${params.slug}`, requestOptions);

	if (res.ok) {
		avenue = (await res.json()) as Avenue;
	} else {
		throw new Error('Avenue not found');
	}

	const profileRes = await fetch(
		`http://localhost:3000/avenue/userinfo/${params.slug}`,
		geoRequestOptions
	);
	if (profileRes.ok) {
		userInfo = (await profileRes.json()) as UserInfo;
	} else {
		throw new Error('User not found');
	}

	if (params.slug == userId) {
		isOwner = true;
	} else {
		isOwner = false;
	}

	console.log(params.slug);

	return {
		slug: params.slug,
		avenue: avenue,
		userInfo: userInfo,
		isOwner: isOwner
	};
};
