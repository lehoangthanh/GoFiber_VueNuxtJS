<template>
	<div class="card">
		<h5 class="card-header">Account Information</h5>
		<div class="card-body">
			<form @submit.prevent="updateIFO">
				<div class="mb-3">
					<label for="first_name" class="form-label">First Name</label>
					<input class="form-control" id="first_name" placeholder="First Name" v-model="infoData.first_name">
				</div>
				
				<div class="mb-3">
					<label for="last_name" class="form-label">Last Name</label>
					<input class="form-control" id="last_name" placeholder="Last Name" v-model="infoData.last_name">
				</div>
				
				<div class="mb-3">
					<label for="email" class="form-label">Email address</label>
					<input type="email" class="form-control" id="email" placeholder="name@example.com" v-model="infoData.email">
				</div>
				<div class="mb-3">
					<button type="submit" class="btn btn-primary">Save</button>
				</div>
			</form>
		</div>
	</div>
	<div class="mt-3"></div>
	<div class="card">
		<h5 class="card-header">Change Password</h5>
		<div class="card-body">
			<form @submit.prevent="changePassword">
				<div class="mb-3">
					<label for="password" class="form-label">Password</label>
					<input type="password" class="form-control" id="password" placeholder="First Name" v-model="passwordData.password">
				</div>
				
				<div class="mb-3">
					<label for="password_confirm" class="form-label">Password Confirm</label>
					<input type="password" class="form-control" id="password_confirm" placeholder="First Name" v-model="passwordData.password_confirm">
				</div>
				<div class="mb-3">
					<button type="submit" class="btn btn-primary">Save</button>
				</div>
			</form>
		</div>
	</div>
	
</template>

<script>
import axios from "axios";
import {useStore} from 'vuex';
import {computed, reactive, watch} from "vue";

export default {
	name: "Profile",
	setup() {
		const user = computed(()=> store.state.user)
		const store = useStore()
		const infoData = reactive({
			first_name: user.value.first_name,
			last_name: user.value.last_name,
			email: user.value.email
		})
		
		const passwordData = reactive({
			password: '',
			password_confirm: ''
		})
		
		watch(user, () => {
			infoData.first_name = user.value.first_name,
			infoData.last_name = user.value.last_name,
			infoData.email = user.value.email
		})
		
		const updateIFO = async () => {
			const { data } = await axios.put('users/info', infoData);
			await store.dispatch('setUser', data)
		};

		const changePassword = async () => {
			await axios.put('users/password', passwordData)
			alert('Success!')
		};
		return {
			infoData,
			passwordData,
			updateIFO,
			changePassword,
			accountInformation_key: 'account_information',
			changePassword_key: 'change_password',
		}
	},
}
</script>

<style scoped>

</style>