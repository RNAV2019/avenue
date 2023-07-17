import type { Avenue, User as DBUSER, UserInfo } from '$lib/helper';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	const requestOptions: RequestInit = {
		method: 'GET'
	};
	let avenue: Avenue | undefined;
	let userInfo: UserInfo | undefined;
	let dbuser: DBUSER | undefined;

	const res = await fetch(`http://localhost:3000/avenue/find/${params.slug}`, requestOptions);

	if (res.ok) {
		dbuser = (await res.json()) as DBUSER;
		avenue = dbuser.avenue;
	} else {
		throw new Error('Avenue not found');
	}

	const profileRes = await fetch(
		`http://localhost:3000/user/info/${dbuser.FirebaseID}`,
		requestOptions
	);
	if (profileRes.ok) {
		userInfo = (await profileRes.json()) as UserInfo;
	} else {
		throw new Error('User not found');
	}

	console.log(params.slug);

	return {
		slug: params.slug,
		avenue: avenue,
		userInfo: userInfo
	};
};
