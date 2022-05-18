package pagerank

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"sync"
)

// CalcPageRankMultiThread calculates the PageRank of the graph using multiple threads
func (pr *PageRank) CalcPageRankMultiThread() {
	// check if edges are empty
	if len(pr.Edges) == 0 {
		return
	}
	// remove self loops
	pr.RemoveSelfLoops()
	// check if there are nodes provides
	if len(pr.Nodes) == 0 {
		return
	}
	// set the amount of workers based on the cpu
	workers := runtime.NumCPU()

	// init nodes by setting the rank to 1/n
	pr.InitializeNodes()
	// iterate unit convergence
	for pr.Iteration = uint(0); pr.Iteration <= pr.MaxIter; pr.Iteration++ {
		fmt.Println("--------------------------------------------")
		fmt.Println("\t Iteration: ", pr.Iteration, "/", pr.MaxIter)
		// init the l1Err error
		pr.l1Err = 0.0

		// int waiting group
		wg := sync.WaitGroup{}
		// init string queue for node ids
		queue := make(chan NodeID, workers)

		// fill queue
		go pr.fillQueue(&wg, queue)

		for i := 0; i < workers; i++ {
			go pr.worker(&wg, queue)
		}
		// wait for all routines to finish
		wg.Wait()

		fmt.Println("\t new err", pr.l1Err)
		// check if the L1 error is smaller than the initial tolerance
		// fmt.Println(l1Err, pr.Tolerance, l1Err < pr.Tolerance)
		if pr.l1Err < pr.Tolerance {
			// fmt.Println("terminate")
			// norm the results by dividing by the sum
			rankSum := pr.SumTotalNodeRank()
			// fmt.Println("terminate")
			for _, node := range pr.Nodes {
				node.Rank = node.Rank / rankSum
			}
			return
		}
	}
	panic("PageRank did not converge in " + strconv.Itoa(int(pr.MaxIter)) + " iterations.")
}

func (pr *PageRank) fillQueue(wg *sync.WaitGroup, queue chan NodeID) {
	wg.Add(1)
	defer wg.Done()
	// iterate over nodes in graph
	for nodeKey := range pr.Nodes {
		// fill queue with all nodes
		queue <- nodeKey
	}
	close(queue)
}

func (pr *PageRank) worker(wg *sync.WaitGroup, queue chan NodeID) {

	wg.Add(1)
	defer wg.Done()

	var process = func(nodeKey NodeID) {
		// calculate the sum of the surrounding nodes of the current node
		rankSumOfConnectedNode := 0.0
		// iterate over the incoming connections
		for _, edge := range pr.Nodes[nodeKey].Incoming {
			cursorNode := edge.From
			// calculate the new rank based on the current rank and the out-degree
			// (out-degree is the number of neighbors of pr node)
			currentRankDividedByOutDegree := cursorNode.Rank / float64(cursorNode.OutDegree())
			// fmt.Println(nodeKey, edge.From.Id, "\t\t", cursorNode.Rank, "/", float64(cursorNode.OutDegree()), "=", currentRankDividedByOutDegree)
			// add the result to the sum
			rankSumOfConnectedNode += currentRankDividedByOutDegree
			// fmt.Println(nodeKey, edge.From.Id, rankSumOfConnectedNode)
		}
		newRank := 1/float64(pr.N) + pr.Alpha*rankSumOfConnectedNode
		//fmt.Println(nodeKey, "new rank", newRank)
		// compute the L1 norm
		absErr := math.Abs(newRank - pr.Nodes[nodeKey].Rank)
		// fmt.Println(nodeKey, "absErr", newRank, "-", pr.Nodes[nodeKey].Rank, "=", absErr)
		pr.mutex.Lock()
		// set the new error
		pr.l1Err += absErr
		// set the new rank value to the current node
		pr.Nodes[nodeKey].Rank = newRank
		pr.mutex.Unlock()
	}

	for {
		select {
		case nodeID, more := <-queue:
			if more {
				process(nodeID)
			} else {
				return
			}
		}
	}

}
