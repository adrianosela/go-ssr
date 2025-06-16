<script lang="ts">
  interface DemoConfig {
    version?: String
  }
  let demoConfig: DemoConfig = { version: "default" };
  console.log("✅ Config prop:", demoConfig);
</script>

<main>
  <div>
    <p>Version: {demoConfig.version}</p>
  </div>
</main>