<script lang="ts">
	import { colorVariants } from '$lib/helper';
	import type { DispatchOptions } from 'svelte';

	export let color:
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
		| 'rose' = 'red';

	let className: string;
	export { className as class };
</script>

<button
	class={`group relative inline-block text-sm font-medium text-black focus:outline-none focus:ring ${className}`}
	type="button"
	on:click
>
	<span
		class={`absolute inset-0 translate-x-1 translate-y-1 bg-black transition-transform group-hover:translate-y-0 group-hover:translate-x-0 ${className}`}
	/>

	<span class={`relative block border-2 border-current ${colorVariants[color]} ${className}`}>
		<div class="grainy w-full h-full flex items-center justify-center"><slot /></div>
	</span>
</button>
