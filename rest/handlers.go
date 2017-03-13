package rest

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/predixdeveloperACN/vcap-support"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {

	// Always set content type and status code
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	var timestamp = time.Now().UTC().Format(time.RFC850)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	message := []string{

		fmt.Sprintf("Received Ping on ... %s \n\n", timestamp),
		fmt.Sprint("Disruption Detection Microservice - Copyright GE 2015-2016\n"),
		fmt.Sprintf("Alloc: %d  (bytes allocated and not yet freed)\n", mem.Alloc),
		fmt.Sprintf("Total Alloc: %d  (bytes allocated (even if freed))\n", mem.TotalAlloc),
		fmt.Sprintf("Sys: %d  (bytes obtained from system)\n", mem.Sys),
		fmt.Sprintf("Free Memory: %d  (Sys - Alloc)\n", mem.Sys-mem.Alloc),
		fmt.Sprintf("Lookups: %d  (number of pointer lookups)\n", mem.Lookups),
		fmt.Sprintf("Mallocs: %d  (number of mallocs)\n", mem.Mallocs),
		fmt.Sprintf("Frees: %d  (number of frees)\n", mem.Frees),
	}
	fmt.Fprintf(w, strings.Join(message, ""))
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Always set content type and status code
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	vcapAppMap, _ := vcap.LoadApplication()
	guid := vcapAppMap.ID
	uri := response_path
	predixSpace := GetPredixSpace()

	var ts = time.Now().UTC()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	ma := float32(mem.Alloc) / 1024.0 / 1024.0 //  convert to MB
	ta := float32(mem.TotalAlloc) / 1024.0 / 1024.0    //  convert to MB
	sm := float32(mem.Sys) / 1024.0 / 1024.0           //  convert to MB
	fm := float32(mem.Sys-mem.Alloc) / 1024.0 / 1024.0 //  convert to MB

	message := []string{

		fmt.Sprintf("Received Info on ... %s \n", ts.Format(time.RFC850)),
		fmt.Sprintf("Application Name: %v \n", vcapAppMap.Name),
		fmt.Sprintf("Predix Space: %s\n", predixSpace),
		fmt.Sprintf("Instance GUID: %s\n", guid),
		fmt.Sprintf("Instance IP: %s\n", os.Getenv("CF_INSTANCE_ADDR")),
		fmt.Sprintf("Instance #: %s\n", os.Getenv("CF_INSTANCE_INDEX")),
		fmt.Sprintf("Started: %s\n", startTime.Format(time.RFC850)),
		fmt.Sprintf("Uptime:  %s\n", ts.Sub(startTime).String()),
		fmt.Sprint("\n"),
		fmt.Sprint("Memory Stats\n"),
		fmt.Sprintf("  Alloc:       %f  (MB allocated and not yet freed)\n", ma),
		fmt.Sprintf("  Total Alloc: %f  (MB allocated (even if freed))\n", ta),
		fmt.Sprintf("  Sys:         %f  (MB obtained from system)\n", sm),
		fmt.Sprintf("  Free Memory: %f  (MB Sys - Alloc)\n", fm),
		fmt.Sprintf("  Lookups:     %d  (number of pointer lookups)\n", mem.Lookups),
		fmt.Sprintf("  Mallocs:     %d  (number of mallocs)\n", mem.Mallocs),
		fmt.Sprintf("  Frees:       %d  (number of frees)\n", mem.Frees),
		fmt.Sprint("\n"),
		fmt.Sprint("Service State\n"),
		fmt.Sprint("Configuration Variables\n"),
		fmt.Sprintf("  URI:  %s\n", uri),
		fmt.Sprintf("  PORT: %s\n", os.Getenv("PORT")),
	}

	if debug {
		fmt.Printf("Info Command\n%s", strings.Join(message, ""))
	}
	fmt.Fprintf(w, strings.Join(message, ""))

}
