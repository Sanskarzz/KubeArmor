package state

import (
	"encoding/json"
	"sync"
	"testing"

	tp "github.com/kubearmor/KubeArmor/KubeArmor/types"
)

func TestSnapshotOmitsExitedContainers(t *testing.T) {
	sa := &StateAgent{
		Containers: map[string]tp.Container{
			"live": {ContainerID: "live", Status: "running"},
			"old":  {ContainerID: "old", Status: "waiting"},
		},
		ContainersLock: new(sync.RWMutex),
	}
	events := sa.snapshotContainerEvents()
	if len(events) != 1 {
		t.Fatalf("got %d containers, want only the running container", len(events))
	}
	var container tp.Container
	if err := json.Unmarshal(events[0].Object, &container); err != nil || container.ContainerID != "live" {
		t.Fatalf("wrong snapshot container: %#v, %v", container, err)
	}
}
