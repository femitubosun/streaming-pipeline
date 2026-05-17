FROM oven/bun:1.2.21-slim
WORKDIR /app

RUN apt-get update -y && apt-get install -y openssl

COPY package.json bun.lock ./
COPY src src/
COPY tsconfig.json ./

RUN bun install

EXPOSE 3000
ENV PORT=3000

CMD ["bun", "run", "start"]
