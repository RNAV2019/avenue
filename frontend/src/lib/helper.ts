export type Avenue = {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	DeletedAt?: null;
	description?: string;
	profile_image?: string;
	name?: string;
	user_id?: string;
	links?: Link[];
};

export type Link = {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	DeletedAt?: Date | null;
	url: string;
	description?: string;
	avenue_id?: number;
};
export type Description = {
	description: string;
};

export type UserInfo = {
	profile_picture?: string;
	name?: string;
	avenue?: Avenue;
};

export const colorVariants = {
	slate: 'bg-slate-500',
	gray: 'bg-gray-500',
	zinc: 'bg-zinc-500',
	neutral: 'bg-neutral-500',
	stone: 'bg-stone-500',
	red: 'bg-red-500',
	orange: 'bg-orange-500',
	amber: 'bg-amber-500',
	yellow: 'bg-yellow-500',
	lime: 'bg-lime-500',
	green: 'bg-green-500',
	emerald: 'bg-emerald-500',
	teal: 'bg-teal-500',
	cyan: 'bg-cyan-500',
	sky: 'bg-sky-500',
	blue: 'bg-blue-500',
	indigo: 'bg-indigo-500',
	violet: 'bg-violet-500',
	purple: 'bg-purple-500',
	fuchsia: 'bg-fuchsia-500',
	pink: 'bg-pink-500',
	rose: 'bg-rose-500'
};

export type color =
	| 'slate'
	| 'gray'
	| 'zinc'
	| 'neutral'
	| 'stone'
	| 'red'
	| 'orange'
	| 'amber'
	| 'yellow'
	| 'lime'
	| 'green'
	| 'emerald'
	| 'teal'
	| 'cyan'
	| 'sky'
	| 'blue'
	| 'indigo'
	| 'violet'
	| 'purple'
	| 'fuchsia'
	| 'pink'
	| 'rose';

export const color = [
	'slate',
	'gray',
	'zinc',
	'neutral',
	'stone',
	'red',
	'orange',
	'amber',
	'yellow',
	'lime',
	'green',
	'emerald',
	'teal',
	'cyan',
	'sky',
	'blue',
	'indigo',
	'violet',
	'purple',
	'fuchsia',
	'pink',
	'rose'
];

export type Statistic = {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	DeletedAt?: null;
	geographic_location?: string;
	traffic_source?: string;
	click_timestamp: Date;
	avenue_id?: number;
};

export type ClickCounts = {
	[clickDate: string]: number;
};
