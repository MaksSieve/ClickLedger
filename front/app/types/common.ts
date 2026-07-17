export interface Dictionary<T> {
    [Key: string]: T;
}

export const colors =[
	'gray',
	'red',
	'green',
	'yellow',
	'blue',
	'indigo',
	'purple',
	'pink',
	'black',
	'cyan',
	'teal',
	'orange',
	'amber'
] as const
export type Color = typeof colors[number];

export type SortDirection = 'asc' | 'desc' | 'none';