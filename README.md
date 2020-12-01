# TaskQueue

## Overview

TaskQueue is a distributed task queue system implemented in Go. It is designed to manage and distribute tasks efficiently across multiple workers, ensuring reliable and scalable task processing.

## Features

- **Distributed Processing**: Distributes tasks among multiple workers for high performance.
- **Reliable Execution**: Ensures tasks are executed even if workers fail.
- **Flexible Scheduling**: Supports various scheduling strategies for task execution.
- **Customizable Workers**: Configurable worker pools to fit different needs.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/khushmanvar/taskqueue.git
    ```

2. Navigate to the project directory:

    ```bash
    cd taskqueue
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

### Starting the Task Queue

To start the task queue, run:

```bash
go run main.go