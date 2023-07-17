import type { User } from 'firebase/auth';
import { writable } from 'svelte/store';

type AuthStoreType = {
	isLoading: boolean;
	currentUser: User | null;
};

export const authStore = writable<AuthStoreType>({
	isLoading: true,
	currentUser: null
});
