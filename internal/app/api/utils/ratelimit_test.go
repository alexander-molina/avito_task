package utils_test

import (
	"testing"
	"time"

	"github.com/alexander-molina/avito_task/internal/app/api/utils"
)

func TestRateLimit_SimpleAddressBlock(t *testing.T) {
	limiter := utils.GetLimiter()
	for i := 0; i < 200; i++ {
		p := limiter.AllowRequests("127.0.0.1/24")
		if i < 100 && !p {
			t.Errorf("Requests blocked for iteration %d", i)
		}

		if i >= 100 && p {
			t.Errorf("Requests not blocked for iteration %d", i)
		}
	}
}

func TestRateLimit_MultipleAddressesBlock(t *testing.T) {
	limiter := utils.GetLimiter()
	for i := 0; i < 200; i++ {
		p := limiter.AllowRequests("127.0.0.2/24")
		if i < 100 && !p {
			t.Errorf("Requests blocked for iteration %d. Address: %s", i, "127.0.0.2/24")
		}

		if i >= 100 && p {
			t.Errorf("Requests not blocked for iteration %d. Address: %s", i, "127.0.0.2/24")
		}

	}

	for i := 0; i < 50; i++ {
		p := limiter.AllowRequests("8.8.8.8/24")
		if p == false {
			t.Errorf("Requests blocked for iteration %d. Address: %s", i, "127.0.0.1/24")
		}
	}

	for i := 0; i < 200; i++ {
		p := limiter.AllowRequests("127.0.0.5/24")
		if i < 100 && !p {
			t.Errorf("Requests blocked for iteration %d. Address: %s", i, "127.0.0.5/24")
		}

		if i >= 100 && p {
			t.Errorf("Requests not blocked for iteration %d. Address: %s", i, "127.0.0.5/24")
		}

	}
}

func TestRateLimit_BlockTimeExpires(t *testing.T) {
	time.Sleep(2 * time.Minute)
	limiter := utils.GetLimiter()
	for i := 0; i < 200; i++ {
		p := limiter.AllowRequests("127.0.0.1/24")
		if i < 100 && !p {
			t.Errorf("Requests blocked for iteration %d. Address: %s", i, "127.0.0.1/24")
		}

		if i >= 100 && p {
			t.Errorf("Requests not blocked for iteration %d. Address: %s", i, "127.0.0.1/24")
		}
	}
}

func TestRateLimit_ResetTrackers(t *testing.T) {
	limiter := utils.GetLimiter()
	for i := 0; i < 200; i++ {
		limiter.AllowRequests("127.0.0.1/24")
	}

	addresses := []string{"127.0.0.3/24", "127.0.0.1/24"}
	limiter.ResetTrackers(addresses)

	for i := 0; i < 200; i++ {
		p := limiter.AllowRequests("127.0.0.1/24")
		if i < 100 && !p {
			t.Errorf("Requests blocked for iteration %d. Address: %s", i, "127.0.0.1/24")
		}

		if i >= 100 && p {
			t.Errorf("Requests not blocked for iteration %d. Address: %s", i, "127.0.0.1/24")
		}
	}
}
