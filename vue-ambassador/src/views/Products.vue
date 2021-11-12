<template>
	<div class="col-md-12" v-if="link">
		<div class="alert alert-info" role="alert">
			{{link}}
		</div>
	</div>
	<div class="col-md-12" v-if="error">
		<div class="alert alert-danger" role="alert">
			{{error}}
		</div>
	</div>
	<div class="col-md-12 mb-4 input-group">
		<input class="form-control" placeholder="Search Products...." @keyup="search($event.target.value)">
		<div class="input-group-append" v-if="selected.length > 0">
			<button class="btn btn-info" @click="generate">Generate Link</button>
		</div>
		<div class="input-group-append">
			<select class="form-select" @change="sort($event.target.value)">
				<option></option>
				<option value="asc">Price Ascending</option>
				<option value="desc">Price Descending</option>
			</select>
		</div>
	</div>
	<div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
		<div class="col" v-for="product in products" :key="product.id" @click="select(product.id)">
			<div :class="selected.some(s => s === product.id) ? 'card shadow-sm selected' : 'card shadow-sm'">
				<img :src="product.image"/>
				<div class="card-body">
					<p class="card-text">{{ product.title }}</p>
					<p class="card-subtitle mb-2 text-muted">{{ product.description }}</p>
					<div class="d-flex justify-content-between align-items-center">
						<small class="text-muted">${{ product.price }}</small>
					</div>
				</div>
			</div>
		</div>
	</div>
	
	<div class="d-flex justify-content-center mt-4">
		<button class="btn btn-primary" @click="loadMore">Load More</button>
	</div>
</template>

<script lang="ts">
import {ref, SetupContext} from "vue";
import axios from "axios";

export default {
	name: "Products",
	props: ['products', 'filters'],
	emits: ['set-filters'],
	setup(props: any, context: SetupContext) {
		const selected = ref<number[]>([])
		const link = ref('')
		const error = ref('')
		
		const search = (s: string) => {
			context.emit('set-filter', {
				...props.filters,
				s
			});
		}
		
		const sort = (sort: string) => {
			context.emit('set-filter', {
				...props.filters,
				sort
			});
		}
		
		const loadMore = () => {
			context.emit('set-filter', {
				...props.filters,
				page: props.filters.page + 1
			});
		}
		
		const select = (id :number) => {
			if (selected.value.some(s => s === id)) {
				selected.value = selected.value.filter(s => s !== id)
				return;
			}
			selected.value = [...selected.value, id];
		}
		
		const generate = async () => {
			try {
				const {data} = await axios.post('links', {
					products: selected.value
				});
				link.value = `Link generated: ${process.env.VUE_APP_CHECKOUT_URL}/${data.code}`
			}catch (e) {
				error.value = 'You should be logged in to generate a link!';
			} finally {
				setTimeout(()=>{
					error.value = ''
				}, 5000)
			}
		}

		return {
			search,
			sort,
			loadMore,
			select,
			selected,
			generate,
			link,
			error
		}
	}
}
</script>

<style scoped>
	.card {
		cursor: pointer;
	}
	
	.card.selected {
		border: 4px solid darkcyan;
	}
</style>