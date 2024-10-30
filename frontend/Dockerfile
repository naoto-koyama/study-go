FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json yarn.lock ./
RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build

FROM node:20-alpine

WORKDIR /app

COPY --from=builder /app/package*.json /app/yarn.lock ./
RUN yarn install --frozen-lockfile

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.ts .
COPY --from=builder /app/package.json .

USER node

EXPOSE 3000

CMD ["yarn", "start"]
