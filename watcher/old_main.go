package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/hashicorp/nomad/api"
// )

// // make tyoe for job scheduling rules
// type JobStatusMap map[string]string
// type JobMeta map[string]string
// type JobSchedulingRules map[string]JobMeta

// func oldMain() {
// 	fmt.Println("Hello, world.")
// 	client, err := createNomadClient()
// 	if err != nil {
// 		fmt.Printf("nomad-watcher: error creating client: %v", err)
// 	}

// 	self, err := client.Agent().Self()
// 	if err != nil {
// 		fmt.Printf("nomad-watcher: error retrieving nodes: %v", err)
// 	}

// 	fmt.Println("Agent ID is: ", self.Member.Name)

// 	// TODO: Change so that this works
// 	// Pretty sure I can just interpolate this into an env var!!
// 	// ${node.unique.name} Or id?
// 	nodeName := "mike-VMW1R39HWM"
// 	nodeFilter := "Name == \"" + nodeName + "\""
// 	nodeQueryOpts := &api.QueryOptions{
// 		Filter: nodeFilter,
// 	}
// 	nodes, _, err := client.Nodes().List(nodeQueryOpts)
// 	if err != nil {
// 		fmt.Printf("nomad-watcher: error retrieving nodes: %v", err)
// 	}

// 	nodeId := ""

// 	for _, node := range nodes {
// 		fmt.Println("Node Id:", node.ID)
// 		nodeId = node.ID
// 	}

// 	jobSchedulingRules := make(JobSchedulingRules)
// 	updatedScheduleRules(client, nodeId, jobSchedulingRules)

// 	ticker := time.NewTicker(1 * time.Second)

// 	go func() {
// 		for range ticker.C {
// 			handleSchedule(client, jobSchedulingRules)
// 		}
// 	}()

// 	for {
// 		fmt.Println("Watcher alive")
// 		time.Sleep(10 * time.Second)
// 	}
// }

// func handleSchedule(client *api.Client, schedulesMap JobSchedulingRules) {
// 	// Get and parse the stop time
// 	// If the job is running and it is past the stop time - send a kill command

// 	// TODO: IF THE JOB NEEDS TO START AND ITS ALLOC WAS GCED, WE WONT GET IT!
// 	// So then we would need to have some sort of schedule-aware

// 	// What if... this is really a pre-start task + a sidecar?
// 	// Prestart boots the main when the time passes (still reads job meta?)
// 	// No need to even use the Nomad client? (can I get meta in env?)
// 	// TODO: this should not be a prestart?!?!
// 	// Sidecar locks the lock, kills the task, turns it off in Nomad

// 	// Retries is infinite since this'll keep going?

// 	for key, value := range schedulesMap {
// 		fmt.Println("Job ID:", key, "Schedule:", value["schedule"])
// 	}
// }

// func updatedScheduleRules(client *api.Client, nodeId string, schedulesMap JobSchedulingRules) {
// 	jobNameMap := make(JobStatusMap)

// 	// TODO: This doesn't work... why?
// 	allocFilter := "DesiredStatus == \"no\""
// 	allocQueryOpts := &api.QueryOptions{
// 		Filter: allocFilter,
// 	}
// 	allocs, _, err := client.Nodes().Allocations(nodeId, allocQueryOpts)
// 	if err != nil {
// 		fmt.Printf("nomad-watcher: error retrieving allocations: %v", err)
// 	}
// 	for _, alloc := range allocs {
// 		// fmt.Println("Alloc Name:", alloc.Name)
// 		// fmt.Println("Job for Alloc:", alloc.JobID)
// 		jobNameMap[alloc.JobID] = alloc.DesiredStatus
// 	}

// 	for key, value := range jobNameMap {
// 		fmt.Println("Job ID:", key, "Alloc Desired Status:", value)
// 	}

// 	for key, _ := range jobNameMap {
// 		// fmt.Println("nomad-watcher: Getting meta for job: ", key)

// 		job, _, err := client.Jobs().Info(key, nil)
// 		if err != nil {
// 			// fmt.Println("nomad-watcher: error retrieving job ", key, err)
// 			continue
// 		}

// 		fmt.Println("Meta:", job.Meta)

// 		if job.Meta["schedule"] != "" {
// 			// fmt.Println("Job ID:", key, "Schedule:", job.Meta["schedule"])
// 			schedulesMap[key] = job.Meta
// 		}
// 	}
// }

// func createNomadClient() (*api.Client, error) {
// 	clientConfig := api.DefaultConfig()
// 	client, err := api.NewClient(clientConfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }