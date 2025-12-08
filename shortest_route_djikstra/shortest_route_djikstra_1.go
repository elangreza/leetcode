package main

import (
	"fmt"
	"os"
	"strings"
)

const INF int64 = 1 << 60

type Edge struct {
	to int
	w  int64
}

func main() {
	var n, m int
	// read node count and edge count
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read node count:", err)
		return
	}
	if _, err := fmt.Scan(&m); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read edge count:", err)
		return
	}

	// temporary storage and mapping from node name to index
	nodeIndex := map[string]int{}
	labels := []string{}
	addNode := func(name string) int {
		if i, ok := nodeIndex[name]; ok {
			return i
		}
		i := len(labels)
		nodeIndex[name] = i
		labels = append(labels, name)
		return i
	}

	edges := make([][3]string, 0, m)
	for i := 0; i < m; i++ {
		var u, v string
		var w int64
		if _, err := fmt.Scan(&u, &v, &w); err != nil {
			fmt.Fprintln(os.Stderr, "failed to read edge:", err)
			return
		}
		edges = append(edges, [3]string{u, v, fmt.Sprintf("%d", w)})
		addNode(u)
		addNode(v)
	}

	var srcName, dstName string
	if _, err := fmt.Scan(&srcName, &dstName); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read query (SRC DST):", err)
		return
	}
	addNode(srcName)
	addNode(dstName)

	// build adjacency list
	N := len(labels)
	adj := make([][]Edge, N)
	for _, e := range edges {
		u := e[0]
		v := e[1]
		w := int64(0)
		fmt.Sscan(e[2], &w)
		ui := addNode(u)
		vi := addNode(v)
		// undirected: add both directions
		adj[ui] = append(adj[ui], Edge{to: vi, w: w})
		adj[vi] = append(adj[vi], Edge{to: ui, w: w})
	}

	// check source and destination exists
	src, okS := nodeIndex[srcName]
	dst, okD := nodeIndex[dstName]
	if !okS || !okD {
		fmt.Println("unreachable")
		return
	}

	// Dijkstra (naive O(V^2) version)
	dist := make([]int64, N)
	prev := make([]int, N)
	vis := make([]bool, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[src] = 0

	for {
		u := -1
		best := INF
		for i := 0; i < N; i++ {
			if !vis[i] && dist[i] < best {
				best = dist[i]
				u = i
			}
		}
		if u == -1 || best == INF {
			break
		}
		vis[u] = true
		// early stop if we fixed destination
		if u == dst {
			break
		}
		for _, e := range adj[u] {
			if dist[u]+e.w < dist[e.to] {
				dist[e.to] = dist[u] + e.w
				prev[e.to] = u
			}
		}
	}

	if dist[dst] == INF {
		fmt.Println("unreachable")
		return
	}

	// reconstruct path
	pathNodes := []string{}
	for cur := dst; cur != -1; cur = prev[cur] {
		pathNodes = append([]string{labels[cur]}, pathNodes...)
	}

	// print result like: A -> C -> D = 2
	fmt.Printf("%s = %d\n", strings.Join(pathNodes, " -> "), dist[dst])
}
