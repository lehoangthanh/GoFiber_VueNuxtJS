<template>
	<Products :products="products" :filters="filters" @set-filter="load"/>
</template>

<script lang="ts">
import Products from "@/views/Products.vue";
import {onMounted, reactive, ref} from "vue";
import axios from "axios";

import {Product} from "@/models/product";
import {Filter} from "@/models/filter";

export default {
	name: "ProductsBackend",
	components: {Products},
	setup() {
		const products = ref<Product[]>([])
		const filters = reactive<Filter>({
			s: '',
			sort: '',
			page: 1
		});
		
		const load = async (f :Filter) => {
			filters.s = f.s
			filters.sort = f.sort
			filters.page = f.page
			const arr = [];
			if (filters.s) {
				arr.push(`s=${filters.s}`)
			}
			
			if (filters.sort) {
				arr.push(`sort=${filters.sort}`)
			}
			
			if (filters.page) {
				arr.push(`page=${filters.page}`)
			}
			
			const {data} = await axios.get(`products/backend?${arr.join('&')}`)
			products.value = [...products.value, ...data.data];
		}
		onMounted( () => load(filters));
		return {
			products,
			filters,
			load
		}
	}
}
</script>

<style scoped>

</style>