FROM oven/bun:1.2.21-slim
WORKDIR /app

COPY package.json bun.lock tsconfig.json ./
RUN bun install --production

COPY src ./src

ENV PORT=3000
EXPOSE 3000
CMD ["bun", "run", "src/index.ts"]
