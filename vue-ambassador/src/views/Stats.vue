<template>
	<div class="album py-5 bg-light">
		<div class="container">
			<div class="table-responsive">
				<table class="table table-striped table-hover">
					<thead>
						<tr>
							<th>Link</th>
							<th>Users</th>
							<th>Revenue</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="link in links" :key="link.id">
							<td>{{checkoutURL(link.code)}}</td>
							<td>{{link.count}}</td>
							<td>{{link.revenue}}</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import {onMounted, ref} from "vue";
import {Link} from "@/models/link"
import axios from "axios";
export default {
	name: "Stats",
	setup() {
		const links = ref<Link[]>([])
		onMounted(async () => {
			const { data } = await axios.get('stats');
			links.value = data;
		});
		const checkoutURL = (code: string) => `${process.env.VUE_APP_CHECKOUT_URL}/${code}`
		return {
			links,
			checkoutURL
		}
	}
}
</script>

<style scoped>

</style>