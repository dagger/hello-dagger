<template>
  <main class="shell">
    <section class="hero">
      <div>
        <p class="eyebrow">Dagger demo monorepo</p>
        <h1>One repo, three CI surfaces, cloud traces on demand.</h1>
        <p class="summary">
          A Vue storefront, Go API, and shared contract package wired through a
          native Dagger module.
        </p>
      </div>

      <img alt="Vue logo" class="logo" src="./assets/logo.svg" />
    </section>

    <section class="grid">
      <article class="panel">
        <span class="label">apps/web</span>
        <h2>Frontend build</h2>
        <p>Vite production build with type checking in a Node container.</p>
        <code>dagger call web-build</code>
      </article>

      <article class="panel">
        <span class="label">services/api</span>
        <h2>Go backend</h2>
        <p>HTTP API exposing health and product endpoints with unit tests.</p>
        <code>dagger call api-test</code>
      </article>

      <article class="panel">
        <span class="label">.dagger</span>
        <h2>Native CI</h2>
        <p>Builds the web dist, runs backend tests, and assembles the API image.</p>
        <code>dagger call ci --cloud</code>
      </article>
    </section>

    <section class="products">
      <div>
        <p class="eyebrow">API preview</p>
        <h2>Featured products</h2>
      </div>

      <ul>
        <li v-for="product in products" :key="product.id">
          <span>{{ product.name }}</span>
          <strong>${{ (product.price / 100).toFixed(2) }}</strong>
        </li>
      </ul>
    </section>
  </main>
</template>

<script setup lang="ts">
const products = [
  { id: 'sku_001', name: 'Trace Hoodie', price: 6400 },
  { id: 'sku_002', name: 'Cloud Runner Mug', price: 1800 },
  { id: 'sku_003', name: 'Pipeline Notebook', price: 1200 },
]
</script>
