package main

import (
	"net/http"
	"strings"

	"container/heap"
	"github.com/crwgregory/golang-api-skeleton/handlers"
	"github.com/crwgregory/golang-api-skeleton/components"
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/config"
	"time"
)

type Worker struct {
	requests chan handlers.Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in heap
}

type Pool []*Worker

type Router struct {
	routes     []handlers.HandlerRoute
	request chan handlers.Request // incoming request
	pool Pool
	done chan *Worker
	logChan chan components.Log
}

func BuildRouter(routes handlers.HandlerRoutes) *Router {
	router := new(Router)
	for _, route := range routes {
		if route.Path[0] != '/' {
			panic("Path has to start with a '/'.")
		}
		router.routes = append(router.routes, route)
	}
	// instantiate the worker pool
	var pool Pool
	for i := 0; i < config.WorkerPoolSize; i++ {
		worker := new(Worker)
		worker.requests = make(chan handlers.Request)
		pool = append(pool, worker)
	}
	heap.Init(&pool)
	router.pool = pool

	router.request = make(chan handlers.Request)
	router.done = make(chan *Worker)

	router.logChan = make(chan components.Log, config.LogChanSize)
	go components.LogRequest(router.logChan)

	go router.balance()
	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	route := r.findMatchingRoute(req)

	if route == nil {
		components.WriteApiResponse(w, components.ApiResponse{
			StatusCode:http.StatusNotFound,
			Message:"Route not found",
		})
	} else {
		response := make(chan components.ApiResponse)
		// put the request into the Routers request channel with the created response channel for this request
		r.request <- handlers.Request{
			Request: req,
			Response: response,
			Route: *route,
			When: time.Now(),
		}
		res := <-response // wait for response
		components.WriteApiResponse(w, res)
	}
}

func (r *Router) dispatch(req handlers.Request) {
	// get the a worker
	w := heap.Pop(&r.pool).(*Worker)
	// go do the work
	go w.work(r)
	// ...send it the task
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&r.pool, w)
}

// Job is complete; update heap
func (r *Router) completed(w *Worker) {
	// One fewer in the queue.
	if w.pending > 0 {
		w.pending--
	}
	// Remove it from heap.
	heap.Remove(&r.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&r.pool, w)
}

func (w *Worker) work(r *Router) {
	for {
		req := <- w.requests
		req.Response <- req.Route.Handler.Handle(req, r.logChan)      // call Handle and send response
		r.done <- w                        		                             		// we've finished this request
	}
}

func (r *Router) balance() {
	for {
		select {
		case req := <- r.request:
			r.dispatch(req) // request came in on routers request channel, dispatch
		case w := <- r.done:
			r.completed(w) // response came in on routers done channel, dispatch
		}
	}
}

func (r *Router) findMatchingRoute(req *http.Request) *handlers.HandlerRoute {
	// find the matching route/handler
	reqPath := strings.Split(req.URL.Path, "/")[1]
	for _, route := range r.routes {
		routePath := route.Path
		if routePath == "/"+reqPath {
			return &route
		}
	}
	return nil
}

func (p Pool) Len() int { return len(p) }
func (p Pool) Less(i, j int) bool { return p[i].pending < p[j].pending }
func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Pool) Push(x interface{}) {
	worker := x.(*Worker)
	worker.index = p.Len()
	*p = append(*p, worker)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	x.index = -1
	*p = old[0 : n-1]
	return x
}

func (r *Router) describePool() {
	fmt.Println("\npool describe:")
	for i := 0; i < r.pool.Len(); i++ {
		worker := r.pool[i]
		fmt.Println("index: ", worker.index, " pending: ", worker.pending)
	}
}
