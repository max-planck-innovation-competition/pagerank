# PageRank

This is a single-threaded implementation of the Personalized PageRank algorithm in golang.
It uses the Gauss - Seidel method and was inspired by the blog post of Giacomo Mariani -
[How to make PageRank faster (with lots of math and a hint of Python)](https://dev.to/karjudev/how-to-make-pagerank-faster-with-lots-of-math-and-a-hint-of-python-2e92)

## Install
```
go get github.com/max-planck-innovation-competition/pagerank
```

## Example

```go
g := NewGraph()
g.AddEdge("0", "1").
    AddEdge("0", "2").
    AddEdge("1", "0").
    AddEdge("1", "2").
    AddEdge("2", "0").
    AddEdge("2", "3").
    AddEdge("3", "0")
pr := NewPageRank(g)
pr.CalcPageRank()

fmt.Pritnln(pr)
```

### Output
```
Pagerank 
Nodes: 4
Iterations: 80
----------------------------------
0			0.370999
1			0.195175
2			0.278124
3			0.155703
----------------------------------
```
### Order Results
you can also order the result with the help of quicksort
```go
g := NewGraph()
g.AddEdge("0", "1").
    AddEdge("0", "2").
    AddEdge("1", "0").
    AddEdge("1", "2").
    AddEdge("2", "0").
    AddEdge("2", "3").
    AddEdge("3", "0")
pr := NewPageRank(g)
pr.CalcPageRank()
pr.OrderResults() // order results

fmt.Println("Max to Min")
for _, k := range pr.GetMaxToMinOrder() {
    fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
}

fmt.Println("Min to Max")
for _, k := range pr.GetMinToMaxOrder() {
    fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
}
```

# License

Apache 2.0