FROM node:18-alpine AS builder
WORKDIR /app
COPY package.json package-lock.json* ./
RUN apk add --no-cache python3 make g++
RUN npm install
COPY . .
RUN npm run build

FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app/.next .next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json
COPY public ./public
COPY pages ./pages
COPY next.config.js ./next.config.js
EXPOSE 3000
CMD ["npm", "run", "start"]
