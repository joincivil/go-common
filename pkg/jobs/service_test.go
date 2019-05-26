package jobs_test

import (
	"testing"
	"time"

	"github.com/joincivil/go-common/pkg/jobs"
)

func TestInMemoryJobService(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	testFunc := func(updates chan<- string) {
		time.Sleep(3 * time.Second)
		updates <- "step1"
		time.Sleep(4 * time.Second)
		updates <- "step2"
	}

	id := "id1"

	_, err := jobsService.StartJob(id, testFunc)
	if err != nil {
		t.Fatalf("Should have started the job: err: %v", err)
	}

	_, err = jobsService.GetJob(id)
	if err != nil {
		t.Errorf("Should have retrieved the job: err: %v", err)
	}
}
func TestInMemoryJobServiceNoJob(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	id := "id1"
	_, err := jobsService.GetJob(id)
	if err == nil {
		t.Errorf("Should not have retrieved the job: err: %v", err)
	}
}
func TestInMemoryJobServiceSameJob(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	testFunc := func(updates chan<- string) {
		time.Sleep(3 * time.Second)
		updates <- "step1"
		time.Sleep(4 * time.Second)
		updates <- "step2"
	}

	id := "id1"

	_, err := jobsService.StartJob(id, testFunc)
	if err != nil {
		t.Fatalf("Should have started the job: err: %v", err)
	}

	_, err = jobsService.StartJob(id, testFunc)
	if err == nil {
		t.Fatalf("Should have failed to start a job with same ID: err: %v", err)
	}
}

func TestInMemoryJobServiceSubscribe(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	testFunc := func(updates chan<- string) {
		time.Sleep(2 * time.Second)
		updates <- "step1"
		time.Sleep(2 * time.Second)
		updates <- "step2"
	}

	id := "id1"

	_, err := jobsService.StartJob(id, testFunc)
	if err != nil {
		t.Fatalf("Should have started the job: err: %v", err)
	}

	sub, err := jobsService.StartSubscription(id)
	if err != nil {
		t.Fatalf("Should have gotten a subscription: err: %v", err)
	}

	updateCount := 0
Loop:
	for {
		select {
		case <-sub.Updates:
			updateCount++
			if updateCount == 2 {
				break Loop
			}
		case <-time.After(10 * time.Second):
			t.Errorf("Should have received all the updates")
			break Loop
		}
	}
}

func TestInMemoryJobServiceSubscribeNoJob(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	testFunc := func(updates chan<- string) {
		time.Sleep(2 * time.Second)
		updates <- "step1"
		time.Sleep(2 * time.Second)
		updates <- "step2"
	}

	id := "id1"

	_, err := jobsService.StartJob(id, testFunc)
	if err != nil {
		t.Fatalf("Should have started the job: err: %v", err)
	}

	badID := "id2"
	_, err = jobsService.StartSubscription(badID)
	if err == nil {
		t.Fatalf("Should have failed to get subscription: err: %v", err)
	}
}

func TestInMemoryJobServiceSubscribeStopSub(t *testing.T) {
	jobsService := jobs.NewInMemoryJobService()

	testFunc := func(updates chan<- string) {
		time.Sleep(2 * time.Second)
		updates <- "step1"
		time.Sleep(2 * time.Second)
		updates <- "step2"
	}

	id := "id1"

	_, err := jobsService.StartJob(id, testFunc)
	if err != nil {
		t.Fatalf("Should have started the job: err: %v", err)
	}

	sub, err := jobsService.StartSubscription(id)
	if err != nil {
		t.Fatalf("Should not have failed to get subscription: err: %v", err)
	}

	err = jobsService.StopSubscription(sub)
	if err != nil {
		t.Fatalf("Should not have failed to stop subscription: err: %v", err)
	}

}
