# Worker Pool Pattern

Job Queue: A buffered queue that holds the incoming tasks or data.

Worker Goroutines: A fixed number of background goroutines running in a loop, waiting for or processing tasks from the channel.
