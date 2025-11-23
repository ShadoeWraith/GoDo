// Tells TypeScript how to handle files ending in .vue
declare module "*.vue" {
  import type { DefineComponent } from "vue";

  // This declares that any import of a '*.vue' file results in a Vue component.
  // This resolves the error: "implicitly has an 'any' type."
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

declare module "*.scss";
