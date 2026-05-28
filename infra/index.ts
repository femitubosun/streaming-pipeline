import * as docker from "@pulumi/docker";

const network = new docker.Network("app-network", {
  name: "grpc-pipeline",
});

// const apiImage = new docker.Image("api", {
//   imageName: "grpc-api",
//   build: { context: "../api" },
//   skipPush: true,
// });

// const producerImage = new docker.Image("producer", {
//   imageName: "grpc-producer",
//   build: { context: "../producer" },
//   skipPush: true,
// });

// const processorImage = new docker.Image("processor", {
//   imageName: "grpc-processor",
//   build: { context: "../processor" },
//   skipPush: true,
// });

const redpanda = new docker.Container("redpanda", {
  image: "docker.redpanda.com/redpandadata/redpanda:v26.1.8",
  name: "redpanda",
  networksAdvanced: [{ name: network.name }],
  ports: [
    {
      internal: 9092,
      external: 9092,
    },
    // {
    //   internal: 9644,
    //   external: 9644,
    // },
  ],
  command: [
    "redpanda",
    "start",
    "--mode",
    "dev-container",
    "--smp",
    "4",
    "--memory",
    "4G",
    "--overprovisioned",
    "--kafka-addr",
    "internal://0.0.0.0:9092",
    "--advertise-kafka-addr",
    "internal://redpanda:9092",
  ],
  memory: 8192,
  mustRun: true,
  healthcheck: {
    tests: ["CMD", "curl", "-f", "http://localhost:9644/v1/status/ready"],
    interval: "5s",
    timeout: "3s",
    retries: 5,
    startPeriod: "10s",
  },
});

// const producer = new docker.Container(
//   "producer",
//   {
//     image: producerImage.imageName,
//     name: "producer",
//     networksAdvanced: [{ name: network.name }],
//     ports: [],
//     envs: ["KAFKA_BROKERS=redpanda:9092"],
//     mustRun: true,
//   },
//   {
//     dependsOn: [redpanda],
//   },
// );

// const processor = new docker.Container(
//   "processor",
//   {
//     image: processorImage.imageName,
//     name: "processor",
//     networksAdvanced: [{ name: network.name }],
//     ports: [],
//     mustRun: true,
//     envs: ["KAFKA_BROKERS=redpanda:9092"],
//   },
//   {
//     dependsOn: [redpanda],
//   },
// );

// const api = new docker.Container("api", {
//   image: apiImage.imageName,
//   name: "api",
//   networksAdvanced: [{ name: network.name }],
//   ports: [
//     {
//       internal: 3000,
//       external: 3000,
//     },
//   ],
//   envs: ["PORT=3000"],
//   mustRun: true,
// });

// export const apiUrl = `http://localhost:3000`;
