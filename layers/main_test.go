package param

import (
	"testing"
	"github.com/ncostamagna/go-memory/layers/param"
	"github.com/ncostamagna/go-memory/layers/result"
	"runtime"
	"fmt"
	"runtime/pprof"
	"log"
	"os"
	"github.com/ncostamagna/go-memory/layers/resultV"
)

const oper = 100000000
func TestMain(m *testing.M) {
	fmt.Println("entra")
	/*go func() {
        for {
            memory()
            time.Sleep(time.Millisecond * 200)
        }
    }()*/

	exitCode := m.Run()
    fmt.Println("Benchmarks completed.")
    os.Exit(exitCode)
}

// we can use prof usint http
func BenchmarkParam(b *testing.B) {
	f, err := os.Create("heap_profile_param.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	repo := param.NewRepository()
	service := param.NewService(repo)
	for i := 0; i < b.N; i++ {
		
		service.Run()
	}

	if err := pprof.WriteHeapProfile(f); err != nil {
        log.Fatal("could not write memory profile: ", err)
    }
}


func BenchmarkResultV(b *testing.B) {
	f, err := os.Create("heap_profile_v.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	repo := resultV.NewRepository()
	service := resultV.NewService(repo)

	for i := 0; i < b.N; i++ {
		
		service.Run()
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
        log.Fatal("could not write memory profile: ", err)
    }
}

func BenchmarkResult(b *testing.B) {
	f, err := os.Create("heap_profile_pointer.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	repo := result.NewRepository()
	service := result.NewService(repo)

	for i := 0; i < b.N; i++ {
		service.Run()
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
        log.Fatal("could not write memory profile: ", err)
    }
}


func memory() {
	var t runtime.MemStats
            runtime.ReadMemStats(&t)
			fmt.Printf("Heap Alloc = %v mb - ", bToMb(t.HeapAlloc))
			fmt.Printf("Sys Memory = %v mb - ", bToMb(t.Sys))
			fmt.Printf("Total Alloc = %v mb - ", bToMb(t.TotalAlloc))
            fmt.Printf("Alloc = %v mb - ", bToMb(t.Alloc))
			fmt.Printf("Stack = %v mb - ", bToMb(t.StackSys))
			fmt.Printf("Frees = %v mb\n", bToMb(t.Frees))
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}