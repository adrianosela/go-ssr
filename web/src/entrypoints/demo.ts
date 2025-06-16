import DemoWidget from '../components/DemoWidget.svelte'

const conf = (window as any).APP_CONFIG || {};
console.log(`Got Config: `, conf);

const aboutEl = document.getElementById('demo-widget');
if (aboutEl) {
  new DemoWidget({ target: aboutEl });
}