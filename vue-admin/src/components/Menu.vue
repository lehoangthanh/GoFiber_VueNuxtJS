<template>
	<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		<div class="position-sticky pt-3">
			<ul class="nav flex-column">
				<li class="nav-item">
					<router-link to="/users" class="nav-link" active-class="active">
						Users
<!--						<span data-feather="home"></span>-->
<!--						Dashboard-->
					</router-link>
					<router-link to="/products" class="nav-link" active-class="active">
						Products
					</router-link>
					<router-link to="/orders" class="nav-link" active-class="active">
						Orders
					</router-link>
				</li>
			</ul>
		</div>
	</nav>
</template>

<script>
export default {
	name: "Menu"
}
</script>

<style scoped>
/*
	 * Sidebar
	 */

/*.sidebar {*/
/*	top: 0;*/
/*	!* rtl:raw:*/
/*	right: 0;*/
/*	*!*/
/*	bottom: 0;*/
/*	!* rtl:remove *!*/
/*	left: 0;*/
/*	z-index: 100; !* Behind the navbar *!*/
/*	padding: 48px 0 0; !* Height of navbar *!*/
/*	box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);*/
/*}*/

@media (max-width: 767.98px) {
	.sidebar {
		top: 5rem;
	}
}

.sidebar-sticky {
	position: relative;
	top: 0;
	height: calc(100vh - 48px);
	padding-top: .5rem;
	overflow-x: hidden;
	overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
}

.sidebar .nav-link {
	font-weight: 500;
	color: #333;
}

.sidebar .nav-link .feather {
	margin-right: 4px;
	color: #727272;
}

.sidebar .nav-link.router-link-exact-active.active {
	color: #2470dc;
}

.sidebar .nav-link:hover .feather,
.sidebar .nav-link.active .feather {
	color: inherit;
}

.sidebar-heading {
	font-size: .75rem;
	text-transform: uppercase;
}
</style>