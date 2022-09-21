<script lang="ts">
	export let title = '';
	let isHovered = false;
	let x: number;
	let y: number;

	function mouseOver(event: MouseEvent) {
		isHovered = true;
		x = event.pageX - 10;
		y = event.pageY + 5;
	}
	function mouseMove(event: MouseEvent) {
		x = event.pageX - 10;
		y = event.pageY + 5;
	}
	function mouseLeave() {
		isHovered = false;
	}
</script>

<div
	on:mouseover={mouseOver}
	on:focus={() => {}}
	on:mouseleave={mouseLeave}
	on:mousemove={mouseMove}
>
	<slot />
</div>

{#if isHovered}
	<div style="top: {y}px; left: {x}px;" class="bg-gray-900 text-white rounded-lg absolute">
		{title}
	</div>
{/if}
