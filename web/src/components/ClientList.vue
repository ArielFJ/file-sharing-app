<script>
import StatCard from "./StatCard.vue"
export default {
  components: {
    StatCard
  },
  data() {
    return {
      currentInterval: -1,
      clients: []
    }
  },
  computed: {
    orderedClients() {
      // Order by name
      return this.clients.sort((a, b) => a.localeCompare(b))
    },
  },
  methods: {
    fetchClientsData() {
      fetch('http://localhost:8001/clients')
        .then(res => res.json())
        .then(data => this.clients = data)
        .catch(console.log)
    }
  },
  mounted() {
    this.fetchClientsData()
    clearInterval(this.currentInterval)
    this.currentInterval = setInterval(this.fetchClientsData, 5000)
  }
}
</script >

  <template>
  <StatCard title="Clients" bgcolor="border-primary">
    <ul>
      <li v-for="clientName, index in orderedClients" :key="index">{{ clientName }}</li>
    </ul>
  </StatCard>
</template>

<style>
li {
  list-style: none;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3);
}
</style>