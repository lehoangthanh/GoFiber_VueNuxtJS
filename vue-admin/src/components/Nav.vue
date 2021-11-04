<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3" href="#">Company name</a>
		
		<input class="form-control form-control-dark w-100" type="text" placeholder="Search" aria-label="Search">
		<div class="navbar-nav">
			<div class="nav-item text-nowrap">
				<router-link to="/profile" class="nav-link px-3" href="#">
					{{user.first_name}} {{user.last_name}}
				</router-link>
				<a class="nav-link px-3" href="#" @click="logout">Sign out</a>
			</div>
		</div>
	</header>
</template>

<script>
import axios from "axios";

export default {
	name: "Nav",
	methods: {
		async logout() {
			await axios.post('logout')
			await this.$router.push('/login')
		}
	},
	computed: {
		user() {
			return this.$store.state.user
		}
	}
}
</script>

<style scoped>
/*
 * Navbar
 */

.navbar-brand {
	padding-top: .75rem;
	padding-bottom: .75rem;
	font-size: 1rem;
	background-color: rgba(0, 0, 0, .25);
	box-shadow: inset -1px 0 0 rgba(0, 0, 0, .25);
}

.navbar .navbar-toggler {
	top: .25rem;
	right: 1rem;
}

.navbar .form-control {
	padding: .75rem 1rem;
	border-width: 0;
	border-radius: 0;
}
</style>