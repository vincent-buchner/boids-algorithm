# Boids Algorithm Simulation

This project is a Go implementation of the Boids algorithm, a computational model that simulates the flocking behavior of birds or other swarming entities. I wanted to implement this algorithm after observing hockey players maneuvering around the ice rink with the puck, exhibiting flocking-like movements.

## Overview

The Boids algorithm was introduced by Craig Reynolds in 1986. It simulates the collective behavior of a group of agents, known as "boids," by applying three simple rules:

1. **Cohesion**: Steer towards the average position of nearby boids.
2. **Separation**: Steer away from nearby boids to avoid collisions.
3. **Alignment**: Steer towards the average heading of nearby boids.

By following these rules, the boids exhibit emergent flocking behavior, forming realistic and natural-looking movements resembling those of birds, fish, or other swarming entities.

## Implementation

This implementation uses the [Ebiten](https://github.com/hajimehoshi/ebiten) game library for rendering the boids on the screen. The project consists of the following files:

- `main.go`: The entry point of the application, responsible for setting up the game loop and initializing the flock of boids.
- `boid.go`: Contains the `Boid` struct and its associated methods for implementing the Boids algorithm rules, updating positions, and rendering.

## Usage

To run the application, make sure you have Go installed on your system. Then, navigate to the project directory and execute the following command:

```bash
go run .
```

This will compile and run the application, displaying a window with the simulated flock of boids.

## Configuration

The application allows you to configure various parameters, such as the screen dimensions, maximum speed, and the number of boids in the flock. These values are defined as constants in `main.go`.

Feel free to experiment with different values to observe the impact on the flocking behavior.

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bug fixes, please open an issue or submit a pull request.
