<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>

	<div class="container-fluid">
		<main>
			<RouterView />
		</main>
	</div>
</template>

<style>
main{
	height: 100%;
	width: 100%;
}
</style>
