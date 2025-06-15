# Build stage
FROM node:20 AS builder

WORKDIR /app
COPY . .

# Avoid interactive prompts
ENV CI=true

RUN corepack enable && pnpm install && pnpm build

# Serve with Nginx
FROM nginx:stable-alpine AS runner

# Copy built assets
COPY --from=builder /app/dist /usr/share/nginx/html

# Optional: custom nginx config for SPA routing
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

