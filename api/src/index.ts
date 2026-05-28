import { credentials } from "@grpc/grpc-js";
import { Hono } from "hono";
import { Empty, ObservabilityServiceClient } from "./proto/observability";

const app = new Hono();

const obsClient = new ObservabilityServiceClient(
  "observability:50051",
  credentials.createInsecure(),
);

let cachedMetrics: any = {
  statusCounts: {},
  totalProcessed: 0,
  lastUpdatedMs: 0,
};
// Poll every 30s
setInterval(() => {
  obsClient.getTransactionMetrics(Empty.create(), (err, response) => {
    if (err) {
      console.error("Failed to fetch metrics:", err);
      return;
    }
    if (response) {
      cachedMetrics = {
        statusCounts: response.statusCounts,
        totalProcessed: response.totalProcessed,
        lastUpdatedMs: response.lastUpdatedMs,
      };
    }
  });
}, 30000);

app.get("/", (c) => {
  return c.text("Hello Hono!");
});

app.get("/stats", (c) => {
  return c.json({
    metrics: cachedMetrics,
    polledAt: new Date().toISOString(),
  });
});

export default app;
