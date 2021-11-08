<template>
	<div class="album py-5 bg-light">
		<div class="container">
			<div class="table-responsive">
				<table class="table table-striped table-hover">
					<thead>
						<tr>
							<th>#</th>
							<th>Name</th>
							<th>Revenue</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="(val, name, i) in rankings" :key="i">
							<td>{{i+1}}</td>
							<td>{{name}}</td>
							<td>{{val}}</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import {onMounted, ref} from "vue";
import {Ranking} from "@/models/ranking"
import axios from "axios";

export default {
	name: "Stats",
	setup() {
		const rankings = ref<Ranking[]>([])
		onMounted(async () => {
			const { data } = await axios.get('rankings');
			rankings.value = data;
		});
		
		return {
			rankings
		}
	}
}
</script>

<style scoped>

</style>