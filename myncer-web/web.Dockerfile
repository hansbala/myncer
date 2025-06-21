# Build stage
FROM node:20 AS builder

WORKDIR /app
COPY . .

# Avoid interactive prompts
ENV CI=true
ENV VITE_API_BASE_URL=https://myncer-api.hansbala.com/api/v1
ENV VITE_GRPC_BASE_URL=https://myncer-api.hansbala.com
ENV VITE_SPOTIFY_CLIENT_ID=e1e0f351205c451faa7ce8af1f862305
ENV VITE_SPOTIFY_REDIRECT_URI=https://myncer.hansbala.com/datasource/spotify/callback
ENV VITE_YOUTUBE_CLIENT_ID=922318482986-2lk1g2u31oc5ucitk7eu83kvbshin1kh.apps.googleusercontent.com
ENV VITE_YOUTUBE_REDIRECT_URI=https://myncer.hansbala.com/datasource/youtube/callback

RUN corepack enable && pnpm install && pnpm build

# Serve with Nginx
FROM nginx:stable-alpine AS runner

# Copy built assets
COPY --from=builder /app/dist /usr/share/nginx/html

# Optional: custom nginx config for SPA routing
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]

