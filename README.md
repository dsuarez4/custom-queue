# custom-queue

#### Collector
http request handler which forms WorkRequests to throw onto the WorkQueue

#### Work queue


#### Dispatcher

Pulls off work requests from the queue and distributes to next available worker 

#### Worker Queue

channel of channels


#### Setup
Build 

`go build -o worker-queue .`

Run

`./worker-queue -n 2048`
