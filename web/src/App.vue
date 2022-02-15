<script>

import ChannelsTable from "./components/ChannelsTable.vue"

export default {
  components: {
    ChannelsTable
  }
}
</script>

<template>
  <div class="custom-container">
    <div class="row">
      <h1>TCP Server Monitor</h1>
      <ChannelsTable />
    </div>
  </div>
</template>

<style>
.custom-container {
  padding: 10px 50px;
}
</style>