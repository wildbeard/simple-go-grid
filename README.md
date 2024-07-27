## Grid Generator

This simple program generates a grid given the `width` and `height` args.
Each Node within the grid gets a unique `UUID` and a list of `ConnectedIDs`.
The `ConnectedIDs` is a list of neighboring nodes directly adjacent and not diagnal to the Node.

#### Example Grid
||||
|:---:|:---:|:---:|
| 0,0 | 1,0 | 2,0 |
| 0,1 | 1,1 | 2,1 |
| 0,2 | 1,2 | 2,2 |

In the above grid, the Node at `1,1` will have `ConnectedIDs` of `(1,0), (0,1), (2,1), (1,2)` but not `0,0` or `2,2`.
