package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/simelo/rextporter/src/util/file"
	"github.com/simelo/rextporter/test/integration/testrand"
	log "github.com/sirupsen/logrus"
)

const exposedMetricsResponse = `
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.10.4"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.060672e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.060672e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.443364e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 1007
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 169984
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.060672e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 401408
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 2.285568e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 8042
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 2.686976e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 15
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 9049
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 31312
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.05954e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 458752
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 458752
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 5.867768e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1024
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 9
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.2218368e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.54180052506e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 3.95628544e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP skycoin_wallet1_seq1 I am running since
# TYPE skycoin_wallet1_seq1 counter
skycoin_wallet1_seq1 0
# HELP skycoin_wallet1_seq1_up Says if the same name metric(skycoin_wallet1_seq1) was success updated, 1 for ok, 0 for failed.
# TYPE skycoin_wallet1_seq1_up gauge
skycoin_wallet1_seq1_up 1
# HELP skycoin_wallet1_seq2 I am running since
# TYPE skycoin_wallet1_seq2 counter
skycoin_wallet1_seq2 0
# HELP skycoin_wallet1_seq2_up Says if the same name metric(skycoin_wallet1_seq2) was success updated, 1 for ok, 0 for failed.
# TYPE skycoin_wallet1_seq2_up gauge
skycoin_wallet1_seq2_up 1
# HELP skycoin_wallet2_seq1 I am running since
# TYPE skycoin_wallet2_seq1 counter
skycoin_wallet2_seq1 0
# HELP skycoin_wallet2_seq1_up Says if the same name metric(skycoin_wallet2_seq1) was success updated, 1 for ok, 0 for failed.
# TYPE skycoin_wallet2_seq1_up gauge
skycoin_wallet2_seq1_up 1
# HELP skycoin_wallet2_seq2 I am running since
# TYPE skycoin_wallet2_seq2 counter
skycoin_wallet2_seq2 0
# HELP skycoin_wallet2_seq2_up Says if the same name metric(skycoin_wallet2_seq2) was success updated, 1 for ok, 0 for failed.
# TYPE skycoin_wallet2_seq2_up gauge
skycoin_wallet2_seq2_up 0
`

const jsonHealthResponse = `
{
    "blockchain": {
        "head": {
            "seq": 58894,
            "block_hash": "3961bea8c4ab45d658ae42effd4caf36b81709dc52a5708fdd4c8eb1b199a1f6",
            "previous_block_hash": "8eca94e7597b87c8587286b66a6b409f6b4bf288a381a56d7fde3594e319c38a",
            "timestamp": 1537581604,
            "fee": 485194,
            "version": 0,
            "tx_body_hash": "c03c0dd28841d5aa87ce4e692ec8adde923799146ec5504e17ac0c95036362dd",
            "ux_hash": "f7d30ecb49f132283862ad58f691e8747894c9fc241cb3a864fc15bd3e2c83d3"
        },
        "unspents": 38171,
        "unconfirmed": 1,
        "time_since_last_block": "4m46s"
    },
    "version": {
        "version": "0.24.1",
        "commit": "8798b5ee43c7ce43b9b75d57a1a6cd2c1295cd1e",
        "branch": "develop"
    },
    "open_connections": 8,
    "uptime": "6m30.629057248s",
    "csrf_enabled": true,
    "csp_enabled": true,
    "wallet_api_enabled": true,
    "gui_enabled": true,
    "unversioned_api_enabled": false,
    "json_rpc_enabled": false
}
`

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("r.RequestURI\n\n", r.RequestURI)
	if _, err := w.Write([]byte(jsonHealthResponse)); err != nil {
		log.WithError(err).Panicln("unable to write response")
	}
	w.WriteHeader(http.StatusOK)
}

func exposedMetricHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("r.RequestURI\n\n", r.RequestURI)
	if _, err := w.Write([]byte(exposedMetricsResponse)); err != nil {
		log.WithError(err).Panicln("unable to write response")
	}
	w.WriteHeader(http.StatusOK)
}

func writeListenPortInFile(port uint16) (err error) {
	var path string
	path, err = testrand.FilePathToSharePort()
	if err != nil {
		return err
	}
	if !file.ExistFile(path) {
		var file, err = os.Create(path)
		if err != nil {
			log.WithError(err).Errorln("error creating file")
			return err
		}
		defer file.Close()
	}
	var file *os.File
	file, err = os.OpenFile(path, os.O_WRONLY, 0400)
	if err != nil {
		log.WithError(err).Errorln("error opening file")
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", port))
	if err != nil {
		log.WithError(err).Errorln("error writing file")
		return err
	}
	err = file.Sync()
	if err != nil {
		log.WithError(err).Errorln("error flushing file")
		return err
	}
	return err
}

func main() {
	var fakeNodePort = testrand.RandomPort()
	if err := writeListenPortInFile(fakeNodePort); err != nil {
		log.Fatal(err)
	}
	log.WithField("port", fakeNodePort).Infoln("starting fake server")
	log.WithError(http.ListenAndServe(fmt.Sprintf(":%d", fakeNodePort), handler)).Fatalln("server fail")
}
