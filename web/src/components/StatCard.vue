<script>
export default {
  props: {
    bgcolor: String,
    title: String,

  }
}
</script>
<template>
  <div :class="`card text-white ${this.bgcolor}`">
    <div class="card-header">{{ this.title}}</div>
    <div class="card-body">
      <slot>Fallback Content</slot>
    </div>
  </div>
</template>

<style>
.card {
  max-height: 60vh;
}

.card-body {
  overflow: auto;
}
</style>