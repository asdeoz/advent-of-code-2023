# Day 8

## Formatting

### Instructions

(\w) -> "$1",\n

Remove last `,`, add `[]`

### Input

(\w{3}) = \((\w{3}), (\w{3})\) -> { "id": "$1", "left": "$2", "right": "$3" },