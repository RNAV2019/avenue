export type Avenue = {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	DeletedAt?: Date | null;
	title?: string;
	description?: string;
	user_id?: number;
	links?: null;
};

export type User = {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	DeletedAt?: Date | null;
	FirebaseID?: string;
	avenue?: Avenue;
};
