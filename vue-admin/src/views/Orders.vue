<template>
	<v-card elevation="2">
		<template>
			<v-expansion-panels>
				<v-expansion-panel
					v-for="(order,i) in orders"
					:key="`order-${order.id}`"
				>
					<v-expansion-panel-header>
						{{ order.name }} ${{ order.total }}
					</v-expansion-panel-header>
					<v-expansion-panel-content>
						<v-simple-table>
							<template v-slot:default>
								<thead>
								<tr>
									<th class="text-left">#</th>
									<th class="text-left">Product Title</th>
									<th class="text-left">Price</th>
									<th class="text-left">Quanity</th>
								</tr>
								</thead>
								<tbody>
								<tr v-for="item in order.order_items" :key="item.id">
									<td>{{ item.id }}</td>
									<td>{{ item.product_title }}</td>
									<td>{{ item.price }}</td>
									<td>{{ item.quantity }}</td>
								</tr>
								</tbody>
							</template>
						</v-simple-table>
					</v-expansion-panel-content>
				</v-expansion-panel>
			</v-expansion-panels>
		</template>
	</v-card>
</template>

<script>
import axios from "axios";
import Vue from "vue";

export default Vue.extend({
	name: "Orders",
	data() {
		return {
			orders: [],
			page: 0,
			lastPage: 0,
		}
	},
	async mounted() {
		const { data } = await axios.get('orders')
		this.orders = data
	}
})
</script>

<style scoped>

</style>