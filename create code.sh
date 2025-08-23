#!/bin/bash

# Create directories for file a
mkdir -p a/go/models
# Create file a/go/models/models.go
cat > a/go/models/models.go << EOF
package models

type A struct {
    Name string
}
EOF

# Create directories for file b
mkdir -p b/go/models
# Create file b/go/models/models.go
cat > b/go/models/models.go << EOF
package models

type B struct {
    Name string
}
EOF

echo "Files and directories created successfully."