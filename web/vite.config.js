import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

export default defineConfig({
	plugins: [vue()],
	server: {
		proxy: {
			"/api": {
				target: "http://localhost:8369",
				changeOrigin: true,
				rewrite: (p) => p.replace(/^\/api/, ""),
			},
		},
	},
	build: {
		rollupOptions: {
			input: {
				login: resolve(__dirname, "login.html"),
				register: resolve(__dirname, "register.html"),
				app: resolve(__dirname, "index.html"),
			},
		},
	},
});
