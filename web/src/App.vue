<script>

export default {
  data() {
    return {
      channels: []
    }
  },
  methods: {
    fetchChannelData() {
      fetch('http://localhost:8001/data')
      .then(res => res.json())
      .then(data => this.channels = data)
      .catch(console.log)
    }
  },
  mounted() {
    this.fetchChannelData()
    setInterval(this.fetchChannelData, 5000)
  }
}
</script>

<template>
  <header>
    <div class="wrapper">
      <table>
        <tr>
          <th>Channel</th>
          <th>Bytes Sent</th>
          <th>Clients</th>
          <th>Files Sent</th>
          <th>Messages Sent</th>
        </tr>
        <tr v-if="channels.length > 0" v-for="channel in channels" :key="channel.channelName">
          <td>{{ channel.name }}</td>
          <td>{{ channel.bytesSent }}</td>
          <td>{{ channel.clients.join() }}</td>
          <td>{{ channel.filesSent.length }}</td>
          <td>{{ channel.messagesSent.length }}</td>
        </tr>
      </table>
    </div>
  </header>

  <main></main>
</template>

<style>
table,
th,
td {
  border: 1px solid black;
  border-collapse: collapse;
}

header {
  height: 30vh;
}
</style>
