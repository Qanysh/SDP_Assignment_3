# SDP_Assignment_3
# Singleton Pattern in Go

This is a simple Go code example that demonstrates the Singleton pattern. The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance. In this example, we create a Singleton structure with lazy initialization and ensure that there is only one instance of it.

## Singleton Pattern Overview

The Singleton pattern restricts the instantiation of a class to a single instance. It is useful when a single object is needed to coordinate actions across the system. In our example, the `Singleton` structure holds a piece of data, and we ensure that only one instance of this structure exists.

## Code Structure

The code consists of the following components:

- `Singleton` structure: It represents the Singleton instance. In this case, it has a `Data` field, which can be set and accessed.
- `GetInstance` function: This function returns the unique instance of the Singleton. It uses the `sync.Once` package to guarantee that the instance is initialized only once.
- `main` function: It demonstrates the usage of the Singleton pattern. It obtains two instances of the Singleton, sets data in one of them, and then prints the data from both instances. It also checks if both instances are the same object.

## How to Use

To use the Singleton pattern in your Go code, follow these steps:

1. Import the necessary packages, such as `fmt` and `sync`.
2. Create a `Singleton` structure to represent the Singleton instance. You can add data and methods as needed.
3. Implement the `GetInstance` function, which ensures that only one instance of the Singleton is created and returned.
4. In your `main` function or other parts of your code, obtain the Singleton instance using `GetInstance`.
5. Use the Singleton instance to access its data or invoke its methods.

# Command Pattern in Go

## Overview

This Go code example demonstrates the Command pattern, a behavioral design pattern. The Command pattern encapsulates a request as an object, thereby allowing for parameterization of clients with queues, requests, and operations. In this example, we create a simple remote control application that can turn a light on and off using commands.

## Code Structure

The code consists of the following components:

- `Command` interface: It defines the `Execute` method, which all concrete commands must implement.
- `Light` receiver: It represents the device that performs specific actions. In this case, it's a light that can be turned on and off.
- Concrete commands: `TurnOnLightCommand` and `TurnOffLightCommand` are specific commands that tie a receiver (`Light`) to an action.
- `Invoker` (`RemoteControl`): It's the initiator that invokes commands.

## How to Use

To use the Command pattern in your Go code, follow these steps:

1. Define the `Command` interface with the `Execute` method to be implemented by all concrete commands.
2. Create the receiver, in this case, the `Light` structure, and define actions as methods on it.
3. Implement concrete commands, each of which binds a receiver to a specific action.
4. Create an initiator, such as the `RemoteControl` structure, that can call commands.
5. Assign commands to the initiator and execute them.


