# Collaborative Mouse Pointer System

A real-time system for sharing mouse cursor positions between users in collaborative web applications.

## Current Status

This project is in early development. Currently implemented:

- ✅ Protobuf definitions for mouse pointer data structures
- ✅ gRPC service definitions for real-time communication
- ✅ Go code generation from protobuf
- ✅ Tests for protobuf message serialization/deserialization

## Requirements

- Go 1.24+
- buf v1.55.1 (will be installed automatically)
- Make (for build automation)

## Quick Start

1. Install buf (if not already installed):
   ```bash
   make install-buf
   ```

2. Generate Go code from protobuf:
   ```bash
   make generate
   ```

3. Run tests:
   ```bash
   make test
   ```

## Project Structure

```
├── proto/
│   └── mousepointer/v1/
│       ├── mouse_pointer.proto    # Protobuf definitions
│       ├── mouse_pointer.pb.go    # Generated Go protobuf code
│       └── mouse_pointer_grpc.pb.go # Generated gRPC code
├── docs/
│   └── REQUIREMENTS.md        # Detailed requirements
├── buf.yaml                   # buf configuration
├── buf.gen.yaml              # buf code generation config
├── go.mod                     # Go module dependencies
├── Makefile                   # Build automation
└── README.md                  # This file
```

## API Overview

The system defines the following core messages:

- `Position`: x,y coordinates for mouse position
- `MouseUpdate`: Complete mouse position update with session/user context
- `JoinSessionRequest/Response`: Session management
- `LeaveSessionRequest/Response`: Session cleanup
- `UserPosition`: Current position of a user in a session

The `MousePointerService` provides these gRPC methods:

- `UpdatePosition`: Send mouse position updates
- `JoinSession`: Join a collaborative session
- `LeaveSession`: Leave a collaborative session  
- `StreamUpdates`: Stream real-time updates from all users in a session

## Next Steps

- [ ] Implement gRPC server with session management
- [ ] Create WebSocket gateway for browser clients
- [ ] Add load balancing and service discovery
- [ ] Implement client-side JavaScript SDK
- [ ] Add monitoring and metrics

## Development

To clean generated files:
```bash
make clean
```

To test protobuf generation specifically:
```bash
make test-proto
```

To run tests with coverage:
```bash
make test-coverage
```

To lint protobuf files:
```bash
buf lint
```

To check for breaking changes:
```bash
buf breaking
``` 
