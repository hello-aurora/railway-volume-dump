FROM node:20-alpine AS builder

WORKDIR /app
COPY . .

RUN yarn install --production --prefer-offline --frozen-lockfile
RUN yarn autoclean --force

FROM node:20-alpine AS runner

WORKDIR /app
ENV NODE_ENV=production

COPY --from=builder /app/package.json ./
COPY --from=builder /app/index.js ./
COPY --from=builder /app/generate-zip.js ./
COPY --from=builder /app/node_modules/ ./node_modules/

ENTRYPOINT ["node", "index.js"]
