package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	var in chan interface{}
	out := make(chan interface{})
	wg := &sync.WaitGroup{}

	for _, curCmd := range cmds {
		wg.Add(1)

		go func(in, out chan interface{}, wg *sync.WaitGroup, curCmd cmd) {
			defer wg.Done()
			defer close(out)
			curCmd(in, out)
		}(in, out, wg, curCmd)

		in = out
		out = make(chan interface{})
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	users := make(map[string]bool)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for elem := range in {
		email, ok := elem.(string)
		if !ok {
			log.Fatal("string conversion error")
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			user := GetUser(email)
			mu.Lock()
			defer mu.Unlock()

			if _, ok = users[user.Email]; !ok {
				users[user.Email] = true
				out <- user
			}
		}()
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	users := make([]User, 0, GetMessagesMaxUsersBatch)
	wg := sync.WaitGroup{}

	for elem := range in {
		user, ok := elem.(User)
		if !ok {
			log.Fatal("User conversion error")
		}

		users = append(users, user)

		if len(users) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			batch := make([]User, len(users))
			copy(batch, users)

			go func(users []User, out chan interface{}, wg *sync.WaitGroup) {
				defer wg.Done()

				messages, err := GetMessages(users...)
				if err != nil {
					log.Fatal(fmt.Errorf("GetMessages error: %v", err))
				}

				for _, message := range messages {
					out <- message
				}
			}(batch, out, &wg)

			users = users[:0]
		}
	}

	if len(users) > 0 {
		wg.Add(1)

		go func(users []User, out chan interface{}, wg *sync.WaitGroup) {
			defer wg.Done()

			messages, err := GetMessages(users...)
			if err != nil {
				log.Fatal(fmt.Errorf("GetMessages error: %v", err))
			}

			for _, message := range messages {
				out <- message
			}
		}(users, out, &wg)
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	var requests = make(chan struct{}, HasSpamMaxAsyncRequests)
	wg := sync.WaitGroup{}

	for elem := range in {
		messageID, ok := elem.(MsgID)
		if !ok {
			log.Fatal("MsgID conversion error")
		}

		requests <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()

			hasFlag, err := HasSpam(messageID)
			if err != nil {
				log.Fatal(fmt.Errorf("HasSpam error: %v", err))
			}

			<-requests
			out <- MsgData{
				ID:      messageID,
				HasSpam: hasFlag,
			}
		}()
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var messages []MsgData

	for elem := range in {
		curMessage, ok := elem.(MsgData)
		if !ok {
			log.Fatal("MsgData conversion error")
		}

		messages = append(messages, curMessage)
	}

	sort.Slice(messages, func(i, j int) bool {
		if messages[i].HasSpam == messages[j].HasSpam {
			return messages[i].ID < messages[j].ID
		}

		return messages[i].HasSpam && !messages[j].HasSpam
	})

	for _, message := range messages {
		out <- fmt.Sprintf("%s %d", strconv.FormatBool(message.HasSpam), uint64(message.ID))
	}
}
