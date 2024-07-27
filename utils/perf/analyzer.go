package perf

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Analyze function to measure execution time and memory usage of a code block
func Analyze(description string, codeBlock func()) {
	// Memory stats and Goroutines count before execution
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)
	memBefore := memStatsBefore.Alloc / 1024   // Convert to KB
	goroutinesBefore := runtime.NumGoroutine() // Number of Goroutines before execution

	// Measure execution time
	start := time.Now()
	codeBlock()
	elapsed := time.Since(start) // Calculate elapsed time

	// Memory stats and Goroutines count after execution
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)
	memAfter := memStatsAfter.Alloc / 1024    // Convert to KB
	goroutinesAfter := runtime.NumGoroutine() // Number of Goroutines after execution

	// Memory used and percentage change
	memUsed := memAfter - memBefore                                   // Memory used during execution
	memChangePercent := (float64(memUsed) / float64(memBefore)) * 100 // Percentage change in memory usage

	// Print the results
	printResults(description, elapsed, memBefore, memAfter, memUsed, memChangePercent, goroutinesBefore, goroutinesAfter)
}

// printResults function to display results in a formatted table
func printResults(description string, elapsed time.Duration, memBefore, memAfter, memUsed uint64, memChangePercent float64, goroutinesBefore, goroutinesAfter int) {
	data := [][]string{
		{"Description", description},                                 // Description of the code block being analyzed
		{"Total Time (s)", fmt.Sprintf("%.6f", elapsed.Seconds())},   // Total execution time in seconds
		{"Memory Before (KB)", fmt.Sprintf("%d", memBefore)},         // Memory usage before execution in KB
		{"Memory After (KB)", fmt.Sprintf("%d", memAfter)},           // Memory usage after execution in KB
		{"Memory Used (KB)", fmt.Sprintf("%d", memUsed)},             // Memory used during execution in KB
		{"Memory Change (%)", fmt.Sprintf("%.2f", memChangePercent)}, // Percentage change in memory usage
		{"Goroutines Before", fmt.Sprintf("%d", goroutinesBefore)},   // Number of Goroutines before execution
		{"Goroutines After", fmt.Sprintf("%d", goroutinesAfter)},     // Number of Goroutines after execution
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"}) // Table headers
	table.SetBorder(true)                        // Enable table border
	table.AppendBulk(data)                       // Append rows to the table
	table.Render()                               // Render the table
}
