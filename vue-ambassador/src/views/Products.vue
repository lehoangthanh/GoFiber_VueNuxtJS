<template>
	<div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
		<div class="col" v-for="product in products" :key="product.id">
			<div class="card shadow-sm">
				<img :src="product.image"/>
				
				<div class="card-body">
					<p class="card-text">{{ product.title }}</p>
					<div class="d-flex justify-content-between align-items-center">
						<small class="text-muted">${{ product.price }}</small>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "Products",
	props: ['products']
}
</script>

<style scoped>

</style>