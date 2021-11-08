<template>
	<Products :products="products"/>
</template>

<script lang="ts">
import {onMounted, ref} from "vue";
import Products from "@/views/Products";
import {Product} from "@/models/product"
import axios from "axios";

export default {
	name: "ProductsFrontend",
	components: {Products},
	setup() {
		const products = ref<Product[]>([])
		onMounted(async () => {
			const {data} = await axios.get('products/frontend')
			products.value = data;
		});
		return {
			products
		}
	}
}
</script>

<style scoped>

</style>