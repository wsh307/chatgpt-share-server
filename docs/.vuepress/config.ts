import { defaultTheme, defineUserConfig } from "vuepress";

export default defineUserConfig({
  lang: "zh-CN",
  title: "ChatgptShareServer",
  description: "基于官网UI的共享账号方案",
  theme: defaultTheme({
    repo: "xyhelper/chatgpt-share-server",
    docsBranch: "master",
    docsDir: "docs",
  })
});
