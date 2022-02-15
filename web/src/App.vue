<script>

import ChannelsTable from "./components/ChannelsTable.vue"
import ClientList from "./components/ClientList.vue"

export default {
  components: {
    ChannelsTable,
    ClientList
  }
}
</script>

<template>
  <div class="custom-container">
    <div class="row">
      <h1>TCP Server Monitor</h1>
      <div class="col-lg-4">
        <ClientList />
      </div>
      <div class="col-lg-8">
        <ChannelsTable />
      </div>
    </div>
  </div>
</template>

<style>
.custom-container {
  padding: 10px 50px;
}
</style>