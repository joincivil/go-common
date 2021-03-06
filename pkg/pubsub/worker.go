package pubsub

import (
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
)

// WorkersConfig configures the governance event pubsub workers
type WorkersConfig struct {
	PubSubProjectID        string
	PubSubTopicName        string
	PubSubSubscriptionName string
	NumWorkers             int
	QuitChan               chan struct{}
	EventHandlers          []EventHandler
}

// NewWorkers configures and returns a new Workers struct
func NewWorkers(config *WorkersConfig) (*Workers, error) {
	if config.EventHandlers == nil || len(config.EventHandlers) == 0 {
		return nil, errors.New("no event handlers configured")
	}
	if config.QuitChan == nil {
		return nil, errors.New("quitChan required")
	}
	if config.PubSubProjectID == "" {
		return nil, errors.New("projectID required")
	}
	if config.PubSubTopicName == "" {
		return nil, errors.New("Topic name required")
	}
	if config.PubSubSubscriptionName == "" {
		return nil, errors.New("Subscription name required")
	}
	if config.NumWorkers == 0 {
		config.NumWorkers = 1
	}
	return &Workers{
		Errors:                 make(chan error),
		pubSubProjectID:        config.PubSubProjectID,
		pubSubTopicName:        config.PubSubTopicName,
		pubSubSubscriptionName: config.PubSubSubscriptionName,
		numWorkers:             config.NumWorkers,
		quitChan:               config.QuitChan,
		eventHandlers:          config.EventHandlers,
		workerStartChan:        make(chan struct{}),
		workerStopChan:         make(chan struct{}),
	}, nil
}

// Workers controls the events workers that handles incoming pubsub events.
// One instance of Workers normally points to one particular queue of events
// by event types.  i.e. one Workers instance for gov events and a separate
// instance for token events.
// Meant to be generic framework where the pubsub queue and the set of
// event handlers are configured before use.
type Workers struct {
	Errors                 chan error
	pubSubProjectID        string
	pubSubTopicName        string
	pubSubSubscriptionName string
	numWorkers             int
	numActiveWorkers       int
	quitChan               chan struct{}
	eventHandlers          []EventHandler
	mut                    sync.Mutex
	workerStartChan        chan struct{}
	workerStopChan         chan struct{}
}

// NumActiveWorkers returns the number of active workers
func (w *Workers) NumActiveWorkers() int {
	w.mut.Lock()
	numActive := w.numActiveWorkers
	w.mut.Unlock()
	return numActive
}

// Start starts up the governance event pubsub workers.
// This is a blocking call.
func (w *Workers) Start() {
	for i := 0; i < w.numWorkers; i++ {
		w.runWorker()
	}

Loop:
	for {
		select {
		case <-w.workerStartChan:
			w.mut.Lock()
			w.numActiveWorkers++
			w.mut.Unlock()
			log.Infof("Worker started, %v active workers, target: %v", w.numActiveWorkers, w.numWorkers)

		case <-w.workerStopChan:
			w.mut.Lock()
			w.numActiveWorkers--
			w.mut.Unlock()
			log.Infof("Worker stopped, %v active workers, target: %v", w.numActiveWorkers, w.numWorkers)

			if w.numActiveWorkers < w.numWorkers {
				log.Infof("Attempting to start worker")
				// Try to restart this worker
				w.runWorker()
			}

		case <-w.quitChan:
			log.Infof("Quitting worker start loop")
			break Loop
		}
	}
}

func (w *Workers) runWorker() {
	go func() {
		time.Sleep(1 * time.Second)
		w.workerStartChan <- struct{}{}
		// Blocks here, unless initial failure
		w.worker()
		w.workerStopChan <- struct{}{}
	}()
}

func (w *Workers) worker() {
	// Initializing pubsub here so each worker has their own subscriber pool
	ps, err := initPubSub(w.pubSubProjectID)
	if err != nil {
		w.Errors <- err
		return
	}

	log.Infof("%v, %v", w.pubSubTopicName, w.pubSubSubscriptionName)
	err = initPubSubSubscribers(ps, w.pubSubTopicName, w.pubSubSubscriptionName)
	if err != nil {
		w.Errors <- err
		return
	}

	log.Info("Start listening for events")
Loop:
	for {
		select {
		case msg, ok := <-ps.SubscribeChan:
			if !ok {
				log.Errorf("Sending on closed channel")
				break Loop
			}

			// Run the event through all the event handlers
		HandlerLoop:
			for _, handler := range w.eventHandlers {
				ran, err := handler.Handle(msg.Data)
				if err != nil {
					log.Errorf("Error handling event on handler %v: err: %v", handler.Name(), err)
					continue
				}
				// event was handled already, so don't need to check other handlers
				if ran {
					break HandlerLoop
				}
			}

			// Ack the pubsub message
			msg.Ack()

		case err := <-ps.SubscribeErrChan:
			log.Errorf("Error in subscriber worker: err: %v", err)
			w.Errors <- errors.WithMessage(err, "error with worker subscriber")

		case <-w.quitChan:
			defer func() {
				if r := recover(); r != nil {
					log.Errorf("Panic captured during worker quit: %v", r)
				}
			}()
			err := ps.StopSubscribers()
			if err != nil {
				log.Errorf("Error stopping subscribers: err: %v", err)
			}
			log.Info("Quitting events worker")
		}
	}
}
