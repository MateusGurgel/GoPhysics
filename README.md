-----

# GoPhysics

GoPhysics is a simple 2D physics engine demonstration built with Go and the Ebiten game library. This project showcases basic rigid body physics, collision detection for circles, and a modular architecture for managing game objects and their interactions.

-----

## Features

  - **2D Physics Simulation**: Basic rigid body physics, including applying forces and updating positions.
  - **Collision Detection**: Implements circle-to-circle collision detection.
  - **Event-driven Collisions**: Objects can react to collisions through an observer pattern.
  - **Game Loop Management**: Separate `Update` and `FixedUpdate` cycles for consistent physics simulation.
  - **Ebiten Integration**: Utilizes the Ebiten game engine for rendering and game loop management.
  - **Modular Design**: Clear separation of concerns for physics, collision, and game object logic.

-----

## Architecture Overview

The project is structured to separate the core physics engine from the game-specific logic, promoting reusability and maintainability.

### Engine Core

The `engine/` directory contains the fundamental building blocks of the physics engine:

  - **`rigidbody/`**:
      - `RigidBody`: Manages an object's physical properties like `mass`, `velocity`, `airDrag`, and `gravity`. It provides methods to apply forces and update the object's position based on physics calculations.
      - `PhysicsObject` interface: Defines the contract for any object that can be manipulated by a `RigidBody`, requiring `GetPosition()` and `SetPosition()` methods.
  - **`colliders/`**:
      - `CircleCollider`: Represents a circular collision boundary. It holds its `Radius`, a reference to the `CollidableObject` it's attached to, and a list of `Observers` to notify upon collision.
      - `ColliderManager`: A singleton (`CM`) responsible for tracking all active colliders in the scene and performing the actual collision checks in a loop.
      - `CollidableObject` interface: Defines the contract for objects that can have colliders, similar to `PhysicsObject` but specifically for collision detection.
      - `Observer` interface: Defines the `OnTrigger()` method, allowing objects to react to collision events without direct coupling.
  - **`game_manager/`**:
      - `GameManager`: The central orchestrator of the game loop. It implements Ebiten's `Game` interface (`Update`, `Draw`, `Layout`). It manages a list of `UpdatableObjects` and ensures `FixedUpdate` (for physics and collisions) runs at a consistent rate.
      - `UpdatableObject` interface: Defines methods for objects that need to be updated and drawn in the game loop.
  - **`time_manager/`**:
      - `TimeManager`: Manages `DeltaTime` (time elapsed since last frame) and `Accumulator` for fixed-time step physics updates.
  - **`utils/`**:
      - `Vector`: A utility struct for 2D vectors, providing methods for common vector operations like addition, division, multiplication, and Euclidean distance calculations.

### Game Objects

Game objects, like the `Ball` in `assets/ball/`, are composed using components from the `engine`.

  - A `Ball` has a `Position`, a `RigidBody`, and `Colliders`.
  - It implements the `UpdatableObject`, `PhysicsObject`, and `Observer` interfaces, allowing it to be part of the game loop, receive physics updates, and react to collisions.
  - The `NewBall` constructor demonstrates how to initialize an object with its physics and collision components.

### Collision System

The collision system follows an observer pattern:

1.  `CircleCollider` instances are registered with the global `ColliderManager (CM)`.
2.  In each `FixedUpdate`, `CM.CheckCollisions()` iterates through all colliders, performing circle-to-circle collision checks.
3.  If a collision is detected, `TriggerCollision()` is called on the collider, which then notifies all registered `Observers` by calling their `OnTrigger()` method.
4.  Game objects (like `Ball`) that act as observers can then implement `OnTrigger()` to define custom behavior upon collision (e.g., applying a force in the `Ball` example).

-----

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

  - **Go 1.23 or newer**: Ensure you have Go installed on your system. You can download it from [golang.org/dl](https://golang.org/dl/).

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/<your-username>/GoPhysics.git
    cd GoPhysics
    ```

    (Replace `<your-username>` with the actual repository owner's username if this project is hosted publicly.)

2.  **Download dependencies:**
    Go modules automatically handle dependency management. Just navigate into the project directory. The `go.mod` file specifies the required modules and their versions.

    ```bash
    go mod download
    ```

    This command will download all the necessary Ebiten and other indirect dependencies.

3.  **Prepare assets:**
    The `assets/ball/ball.go` file attempts to load `resources/Ball.png`. Make sure this file exists in the correct relative path from where you run the application. For demonstration purposes, you might need to create a `resources` directory in the root of your project and place an image named `Ball.png` inside it. A simple placeholder image will suffice to get the example running.

    ```
    GoPhysics/
    ├── assets/
    ├── engine/
    ├── resources/
    │   └── Ball.png  <-- Place your image here
    ├── go.mod
    ├── go.sum
    └── main.go
    ```

### Running the Application

Once you've installed dependencies and placed the `Ball.png` asset, you can run the application:

```bash
go run main.go
```

A new window titled "Hello, World\!" will appear, displaying two circles (balls) moving and interacting according to the defined physics and collision rules. When they collide, they will apply a force on each other.

-----

## Project Structure

The project is organized into logical directories:

```
GoPhysics/
├── assets/                     # Game-specific assets (e.g., game objects, sprites)
│   └── ball/                   # Example: Ball object implementation
├── docs/                       # Project documentation, including Architecture Decision Records (ADRs)
│   ├── ADR - Collider Manager.md
│   └── ADR - Colliders.md
├── engine/                     # Core physics and game engine components
│   ├── colliders/              # Collision detection system (CircleCollider, ColliderManager, interfaces)
│   ├── game_manager/           # Game loop and object management (GameManager, UpdatableObject)
│   ├── rigidbody/              # Rigid body physics simulation (RigidBody, PhysicsObject)
│   ├── time_manager/           # Time tracking for physics updates
│   └── utils/                  # Common utility functions (Vector)
├── go.mod                      # Go module definition and direct dependencies
├── go.sum                      # Go module checksums for reproducible builds
├── main.go                     # Main application entry point
└── resources/                  # Folder for game assets like images
    └── Ball.png                # Example image asset for the ball (create this file)
```

-----

## Design Decisions (ADRs)

The `docs/` directory contains Architecture Decision Records (ADRs) that explain key design choices made during the development of this project. These documents provide context, explore alternatives, and detail the chosen solutions and their consequences.

  - **ADR - Collider Manager**: Discusses the decision to use an observer pattern for notifying objects about collisions, acknowledging the potential for increased complexity in object instantiation.
  - **ADR - Colliders**: Explains the approach of using interfaces for collider creation to hide internal attributes and prevent cyclical dependencies between objects and their colliders. It also highlights the consequence of restricted direct access to collider dimensions from objects.

Reviewing these ADRs can provide deeper insight into the design philosophy behind GoPhysics.

-----

## Contributing

Contributions to GoPhysics are welcome\! If you have suggestions, bug reports, or want to add new features, please feel free to:

1.  Fork the repository.
2.  Create a new branch for your changes.
3.  Make your modifications.
4.  Ensure your code adheres to Go best practices and is well-documented.
5.  Submit a pull request.

-----

## License

This project is open-source and available under the [MIT License](https://opensource.org/licenses/MIT).

-----
