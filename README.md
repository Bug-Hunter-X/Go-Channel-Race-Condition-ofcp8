# Go Channel Race Condition

This repository demonstrates a subtle race condition in Go that can occur when using channels.  The provided `bug.go` code exhibits this issue, while `bugSolution.go` offers a corrected version.

## Problem

The problem lies in the interaction between a goroutine sending values on a channel and the main goroutine receiving those values. If the sending goroutine closes the channel before the receiving loop has processed all values, some values might be lost.

## Solution

The solution involves ensuring that the receiving goroutine completely processes all values before the sending goroutine closes the channel. This can be done using a wait group or by using a separate signal channel for completion. The `bugSolution.go` code provides one method of solving this.