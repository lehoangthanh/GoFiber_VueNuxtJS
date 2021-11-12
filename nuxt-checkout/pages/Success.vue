<template>
  <main>
    <div class="py-5 text-center">
      <h2>Success</h2>
      <p class="lead">
        Your purchase has been completed!
      </p>
    </div>
  </main>
</template>

<script>
export default {
  name: "Success",
  async mounted() {
    const {data} = await this.$axios.post('orders/confirm', {
      source: this.$route.params.source
    })
    console.log('=====', data)
  }
}
</script>

<style scoped>

</style>
