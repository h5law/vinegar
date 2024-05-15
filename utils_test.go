package vinegar

import (
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789[](){}~\\/;:"

func TestFormatKeyword(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < min(runtime.NumCPU(), runtime.GOMAXPROCS(0)); i++ {
		wg.Add(1)
		go func(t *testing.T, wg *sync.WaitGroup) {
			defer wg.Done()
			b := make([]byte, len(alphabet))
			for i := range b {
				b[i] = alphabet[rand.Intn(len(alphabet))]
			}
			randString := string(b)
			str := strings.ReplaceAll(randString, " ", "")
			str = strings.ToLower(str)
			for _, c := range str {
				if !strings.Contains(alphabet[:26], string(c)) {
					str = strings.ReplaceAll(str, string(c), "")
				}
			}
			formatStr := formatKeyword(randString)
			require.Equal(t, str, formatStr)
		}(t, wg)
	}
	wg.Wait()
}
