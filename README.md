# Dijkstra's Algorithm in Go

This repository contains an implementation of **Dijkstra's Algorithm** in Go. Dijkstra's Algorithm is a popular algorithm for finding the shortest path between nodes in a graph, particularly in graphs with non-negative weights. This implementation uses a min-heap (priority queue) to optimize the search for the shortest path.

## Features

- Graph represented as an adjacency list.
- Efficient shortest path search using a min-heap priority queue.
- Customizable graph structure for directed, weighted edges.

## Getting Started

### Prerequisites

To run this code, make sure you have [Go installed](https://golang.org/dl/). This implementation was written in Go version 1.18, but it should work with any recent version of Go.

### Expected Output

The program will output the shortest distance from the start node (node 1 in this example) to each reachable node.

```text
Distance from 1 to 1: 0
Distance from 1 to 2: 2
Distance from 1 to 3: 3
Distance from 1 to 4: 8
Distance from 1 to 5: 6
Distance from 1 to 6: 9
