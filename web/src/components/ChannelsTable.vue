<script>
export default {
  data() {
    return {
      channels: [
        {
          name: 'channel1'
        },
        {
          name: 'channel2'
        },
      ]
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
  <div class="accordion" id="accordionExample">
    <div
      class="accordion-item"
      v-if="channels.length > 0"
      v-for="channel in channels"
      :key="channel.channelName"
    >
      <h2 class="accordion-header" :id="`heading${channel.name}`">
        <button
          class="accordion-button collapsed"
          type="button"
          data-bs-toggle="collapse"
          :data-bs-target="`#collapse${channel.name}`"
          aria-expanded="false"
          :aria-controls="`collapse${channel.name}`"
        >Channel {{ channel.name }}</button>
      </h2>
      <div
        :id="`collapse${channel.name}`"
        class="accordion-collapse collapse"
        :aria-labelledby="`heading${channel.name}`"
        data-bs-parent="#accordionExample"
      >
        <div class="accordion-body">
          <strong>This is the first item's accordion body.</strong> It is shown by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions. You can modify any of this with custom CSS or overriding our default variables. It's also worth noting that just about any HTML can go within the
          <code>.accordion-body</code>, though the transition does limit overflow.
        </div>
      </div>
    </div>
  </div>

  <!-- <div class="wrapper">
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
  </div>-->
</template>

<style>
table,
th,
td {
  border: 1px solid black;
  border-collapse: collapse;
}
</style>
