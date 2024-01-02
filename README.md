# MicroForge Microservice Template

Welcome to MicroForge, a streamlined microservice template designed to kickstart your service-oriented applications with ease and efficiency. Crafted with best practices in mind, this template encapsulates the essential scaffolding required to build, deploy, and maintain microservices at scale. It provides a robust foundation for developers looking to launch services with a Docker-centric workflow, structured logging, and organized configuration management.

---

1. **`cmd`**: Contains the application's entry points. Each subdirectory here represents a standalone application. It's where the main function for each executable of your project resides.

2. **`internal`**: This is where the core logic of the application is stored, divided into several subdirectories:
   - **`api`**: Holds the API/HTTP layer code, including router setup and request handling. This is where you define your REST endpoints and the HTTP server.
   - **`service`**: Contains the business logic of the application. This layer usually calls methods from the `repository` layer and performs business operations.
   - **`repository`**: Includes code for data storage/retrieval, abstracting the data source details. This layer interacts with the database or any other data sources.
   - **`models`**: Defines data structures and types used across the application, like your entities and DTOs (Data Transfer Objects).
   - **`util`**: Utility functions and shared code that can be used across different parts of the application. It's for code that doesn't fit into other layers but is shared across them.

3. **`pkg`**: Reusable libraries and packages that can be used in other projects. This directory is intended for code that can be safely used by external applications.

4. **`configs`**: Configuration files and constants. This includes configuration files, like JSON or YAML, and Go files for constants or configuration structures.

5. **`scripts`**: Utility scripts for tasks like building, deploying, or database migrations. This directory is often used for operational tasks.

6. **`docs`**: Documentation for the project, like API specs and design documents. It's where you put your READMEs, API documentation, and other explanatory materials.

7. **`tests`**: Contains all tests, typically mirroring the structure of the `internal` directory. This is where you place your unit tests, integration tests, and any other testing-related files.

This structure helps in keeping the code organized and maintainable, especially important for larger projects or when working in a team.
