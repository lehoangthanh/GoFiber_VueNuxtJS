<template>
	<Products :products="products"/>
</template>

<script lang="ts">
import Products from "@/views/Products";
import {onMounted, ref} from "vue";
import {Product} from "@/models/product";
import axios from "axios";
export default {
	name: "ProductsBackend",
	components: {Products},
	setup() {
		const products = ref<Product[]>([])
		onMounted(async () => {
			const {data} = await axios.get('products/backend')
			products.value = data.data;
		});
		return {
			products
		}
	}
}
</script>

<style scoped>

</style>