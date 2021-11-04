<template>
	<v-form @submit.prevent="submit">
		<v-text-field label="Title" v-model="title"/>
		<v-textarea label="Description" v-model="description"/>
		<v-text-field label="Price" type="number" min="0" v-model="price"/>
		<v-text-field label="Image" v-model="image"/>
		<v-btn color="primary" type="submit">Save</v-btn>
	</v-form>
</template>

<script>
import axios from "axios";

export default {
	name: "ProductForm",
	data() {
		return {
			title: '',
			description: '',
			image: '',
			price: ''
		}
	},
	async mounted() {
		if (this.$route.params.id) {
			const {data} = await axios.get(`products/${this.$route.params.id}`)
			this.title = data.title
			this.description = data.description
			this.image = data.image
			this.price = data.price
		}
		
	},
	methods: {
		async submit() {
			const data = {
				title: this.title,
				description: this.description,
				image: this.image,
				price: parseFloat(this.price),
			}
			if (this.$route.params.id) {
				await axios.put(`products/${this.$route.params.id}`, data)
			} else {
				await axios.post('products', data)
			}
			await this.$router.push('/products')
		}
	}
}
</script>

<style scoped>

</style>