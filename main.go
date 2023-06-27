package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()

	val, err := fetchUserData(ctx, 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", val)
	fmt.Printf("Elapsed %s ms\n", time.Since(start))
}

type Response struct {
	value string
	err   error
}

func fetchUserData(ctx context.Context, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyReallySlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("fetching data from third party took to long")
		case res := <-respch:
			return res.value, res.err
		}
	}
}

func fetchThirdPartyReallySlow() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return `{"name": "Peter"}`, nil
}
