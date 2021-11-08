<template>
	<div>
		<v-card elevation="4" shaped class="pa-5" :key="accountInformation_key">
			<v-card-title class="red--text">Account Information</v-card-title>
			<v-form @submit.prevent="updateIFO">
				<v-text-field label="First Name" v-model="user.first_name"/>
				<v-text-field label="Last Name" v-model="user.last_name"/>
				<v-text-field label="Email" v-model="user.email"/>
				<v-btn color="primary" type="submit">Save</v-btn>
			</v-form>
		</v-card>
		
		<v-card elevation="5" shaped class="mt-10 pa-5" :key="changePassword_key">
			<v-card-title class="red--text">Change Password</v-card-title>
			<v-form @submit.prevent="changePassword">
				<v-text-field type="password" label="Password" v-model="password"/>
				<v-text-field type="password" label="Password Confirm" v-model="password_confirm"/>
				<v-btn color="primary" type="submit">Save</v-btn>
			</v-form>
		</v-card>
	</div>
</template>

<script>

export default {
	name: "Profile",
	data () {
		return {
			accountInformation_key: 'account_information',
			changePassword_key: 'change_password',
			password: '',
			password_confirm: ''
		}
	},
	methods: {
		async updateIFO() {
			const { data } = await axios.put('users/info', {
				first_name: this.user.first_name,
				last_name: this.user.last_name,
				email: this.user.email
			})
		},
		
		async changePassword() {
			await axios.put('users/password', {
				password: this.password,
				password_confirm: this.password_confirm
			})
			alert('Success!')
		},
	},
	computed: {
		user() {
			return this.$store.state.user
		}
	}
}
</script>

<style scoped>

</style>