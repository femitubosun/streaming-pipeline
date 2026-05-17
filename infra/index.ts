import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

const network = new docker.Network("app-network", {
  name: "grpc-pipeline",
});

const apiImage = new docker.Image("api", {
  imageName: "grpc-api",
  build: { context: "../api" },
  skipPush: true,
});

// const producerImage = new docker.Image("producer", {
//   imageName: "grpc-producer",
//   build: { context: "../producer" },
// });

// const processorImage = new docker.Image("processor", {
//   imageName: "grpc-processor",
//   build: { context: "../processor" },
// });

// const processor = new docker.Container("processor", {
//   image: processorImage.imageName,
//   name: "processor",
//   networksAdvanced: [{ name: network.name }],
//   ports: [],
//   envs: ["PORT=50051"],
//   mustRun: true,
// });

// const producer = new docker.Container("producer", {
//   image: producerImage.imageName,
//   name: "producer",
//   networksAdvanced: [{ name: network.name }],
//   ports: [],
//   envs: [
//     "PORT=50051",
//     pulumi.interpolate`PROCESSOR_ADDRESS=${processor.name}:50051`,
//   ],
//   mustRun: true,
// });

const api = new docker.Container("api", {
  image: apiImage.imageName,
  name: "api",
  networksAdvanced: [{ name: network.name }],
  ports: [
    {
      internal: 3000,
      external: 3000,
    },
  ],
  envs: [
    "PORT=3000",
    // pulumi.interpolate`PRODUCER_ADDRESS=${producer.name}:50051`,
    // pulumi.interpolate`PROCESSOR_ADDRESS=${processor.name}:50051`,
  ],
  mustRun: true,
});

export const apiUrl = `http://localhost:3000`;
