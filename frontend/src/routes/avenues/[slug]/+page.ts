import type { UserInfo } from '$lib/helper';
import Cookie from 'js-cookie';
import { PUBLIC_API_ENDPOINT_URL } from '$env/static/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch, data }) => {
	const headers = new Headers();
	const location = Cookie.get('location') || null;
	if (location) {
		headers.append('X-Geo-Location', location);
	}
	const geoRequestOptions: RequestInit = {
		method: 'GET',
		headers: headers
	};
	let userInfo: UserInfo;
	let isOwner = false;
	const userId: string | null = data.session?.user?.id;

	const profileRes = await fetch(
		`${PUBLIC_API_ENDPOINT_URL}/avenue/userinfo/${params.slug}`,
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

	return {
		slug: params.slug,
		avenue: userInfo.avenue,
		userInfo: userInfo,
		isOwner: isOwner
	};
};
