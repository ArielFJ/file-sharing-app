<script>
import StatCard from "./StatCard.vue"

export default {
  components: {
    StatCard
  },
  data() {
    return {
      currentInterval: -1,
      channels: []
    }
  },
  computed: {
    bytesByChannel() {
      return this.channels.map(channel => channel.bytesSent)
    },
    orderedChannels() {
      if (this.totalBytesSent) {
        // Order from greater to smaller
        return this.channels.sort((chanA, chanB) => chanB.bytesSent - chanA.bytesSent)
      }
      else {
        // Order by name
        return this.channels.sort((a, b) => a.name.localeCompare(b.name))
      }
    },
    maxBytesSent() {
      return Math.max(...this.bytesByChannel)
    },
    totalBytesSent() {
      return this.bytesByChannel.reduce((prev, curr) => prev + curr)
    }
  },
  methods: {
    fetchChannelData() {
      fetch('http://localhost:8001/data')
        .then(res => res.json())
        .then(data => this.channels = data)
        .catch(console.log)
    },
    getBytesSentAsPecent(channel) {
      const percent = parseInt(100 * (channel.bytesSent / this.totalBytesSent))
      return Object.is(percent, NaN) ? 0 : percent
    }
  },
  mounted() {
    this.fetchChannelData()
    clearInterval(this.currentInterval)
    this.currentInterval = setInterval(this.fetchChannelData, 5000)
  }
}
</script>

<template>
  <div class="accordion" id="accordionExample">
    <div
      class="accordion-item"
      v-if="channels.length > 0"
      v-for="channel in orderedChannels"
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
        >
          <div class="w-100 row">
            <div class="col-3">Channel {{ channel.name }}</div>
            <div class="col-6 offset-2">
              {{ getBytesSentAsPecent(channel) }}% of total bytes sent to this channel
              <div class="progress">
                <div
                  class="progress-bar bg-warning"
                  role="progressbar"
                  :style="`width: ${getBytesSentAsPecent(channel)}%;`"
                  aria-valuenow="75"
                  aria-valuemin="0"
                  aria-valuemax="100"
                ></div>
              </div>
            </div>
          </div>
        </button>
      </h2>
      <div
        :id="`collapse${channel.name}`"
        class="accordion-collapse collapse"
        :aria-labelledby="`heading${channel.name}`"
        data-bs-parent="#accordionExample"
      >
        <div class="accordion-body">
          <div class="row">
            <div class="col-lg-4">
              <StatCard title="Files Sent" bgcolor="bg-info">
                <ul class="p-0">
                  <li v-for="fileName, index in channel.filesSent" :key="index">{{ fileName }}</li>
                </ul>
              </StatCard>
            </div>
            <div class="col-lg-5">
              <StatCard title="Messages Sent" bgcolor="bg-primary">
                <ul class="p-0">
                  <li v-for="msg, index in channel.messagesSent" :key="index">{{ msg }}</li>
                </ul>
              </StatCard>
            </div>
            <div class="col-lg-3">
              <StatCard title="Additional Data" bgcolor="border-success">
                <p>Name: {{ channel.name }}</p>
                <p>Bytes sent: {{ channel.bytesSent }}</p>
              </StatCard>
              <br />
              <StatCard title="Clients" bgcolor="bg-secondary">
                <ul class="p-0">
                  <li v-for="client, index in channel.clients" :key="index">{{ client }}</li>
                </ul>
              </StatCard>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
table,
th,
td {
  border: 1px solid black;
  border-collapse: collapse;
}

li {
  list-style: none;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3);
}

.accordion-collapse {
  max-height: 75vh;
}
</style>
