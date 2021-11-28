func allPathsSourceTarget(graph [][]int) [][]int {
    out := make([][]int, 0)
    path := make([]int, 0)

    dfs(0, graph, path, &out)

    return out
}

func dfs(cur int, graph [][]int, path []int, out *[][]int) {
    path = append(path, cur)

    if cur == len(graph) - 1 {
        *out = append(*out, dup(path))
    }

    for _, dest := range graph[cur] {
        dfs(dest, graph, path, out)
    }
}

func dup(slice []int) []int {
    out := make([]int, len(slice))

    copy(out, slice)

    return out
}
