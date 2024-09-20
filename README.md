# QueueMaster - Distributed Task Queue System 🚀

**QueueMaster** is a distributed task queue system built with **Go**, **gRPC**, and **RabbitMQ**. It allows clients to submit tasks, which are processed asynchronously by worker nodes. The system is designed to handle long-running or resource-intensive jobs by distributing tasks across multiple workers with **RabbitMQ** as the message broker.

## Features 🌟

- **gRPC API** for submitting tasks and receiving results 🖥️
- **RabbitMQ** for task queueing and reliable message passing 📦
- **Worker Nodes** that process tasks concurrently, enabling horizontal scaling 🔄
- Real-time task status updates via gRPC streaming 🔔
- Configurable task retry mechanism and error handling ⚠️
- Modular architecture for easy extension and maintenance 🛠️
- Real-time metrics for monitoring and alerting 📊

## Tech Stack 🧩

- **Go**: Core language for building the system 🏗️
- **gRPC**: High-performance RPC framework for communication between client, server, and workers 📡
- **RabbitMQ**: Message broker to queue and distribute tasks to worker nodes 🐇
- **Docker**: Containerization for easy deployment and environment setup 🐳
- **Prometheus**: Real-time metrics for monitoring and alerting 📊
- **Grafana**: Dashboard for monitoring and visualization 📊

## Getting Started 🚀

### Prerequisites 🛠️

- **Go** 1.19+
- **Docker** and **Docker Compose**
- **RabbitMQ** (can be run via Docker Compose)
- **Prometheus** and **Grafana** (can be run via Docker Compose)

### Installation 🛠️

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yourusername/QueueMaster.git
    cd QueueMaster
    ```

2. **Set up environment**:
    You can use Docker Compose to spin up RabbitMQ, the gRPC server, and the worker nodes:

    ```bash
    docker-compose up --build
    ```

3. **Run the gRPC server** (if not using Docker):

    ```bash
    go run cmd/server/main.go
    ```

4. **Run a worker**:

    ```bash
    go run cmd/worker/main.go
    ```

5. **Submit a task via the client**:

    ```bash
    go run cmd/client/main.go submit -T "math" -P "1 + 2 * 5"
    ```

### gRPC API 🛠️

The gRPC service is defined in the `taskqueue.proto` file located in the `/api/` directory. After running the `protoc` compiler, the generated Go files can be found in `/api/pb/`.

Example service definition (in `taskqueue.proto`):

```proto
service TaskQueue {
    rpc SubmitTask (TaskRequest) returns (TaskResponse);
    rpc GetTaskStatus (TaskStatusRequest) returns (TaskStatusResponse);
    rpc StreamTaskResults (stream TaskResultRequest) returns (stream TaskResultResponse);
}
```

### Configuration ⚙️

The configuration is managed via a `config.yaml` file located in the `/config/` folder. You can adjust RabbitMQ credentials, server ports, and other settings there.

### Task Flow 🔄

1. The **client** submits a task via gRPC.
2. The **gRPC server** pushes the task into the **RabbitMQ** queue.
3. A **worker node** picks up the task, processes it, and sends the result back to the client via gRPC.

## Contributing 🤝

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

## License 📜

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
