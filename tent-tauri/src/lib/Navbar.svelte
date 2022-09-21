<script lang="ts">
	import { page } from '$app/stores';
	import ManageAccount from './ManageAccount.svelte';

	let path: string;

	function getPath(currentPath: string) {
		path = currentPath;
		console.log(path);
	}

	$: getPath($page.url.pathname);

	// your script goes here
	type route = {
		name: string;
		path: string;
	};

	const leftRoutes: route[] = [
		{ name: 'home', path: '/' },
		{ name: 'products', path: '/products' },
		{ name: 'catalog', path: '/catalog' },
		{ name: 'buys', path: '/buys' },
		{ name: 'sales', path: '/sales' }
	];
</script>

<main>
	<nav class="border-gray-200 bg-white px-2 text-white dark:border-gray-700 dark:bg-black md:h-12">
		<div class="flex flex-col md:flex-row md:h-full items-center">
			{#each leftRoutes as route}
				<a
					href={route.path}
					class="text-center  text-sm hover:bg-gray-900 uppercase  border-gray-400
                    md:px-3 md:py-3 w-36 {path === route.path ? 'md:border-b-2' : ''}"
				>
					{route.name}
				</a>
			{/each}

			<div class="md:flex md:grow justify-end">
				<ManageAccount />
			</div>
		</div>
	</nav>
</main>
