package main

import "fmt"

// single channel relay
func single_channel_author(write chan<- any) {
	for {
		write <- "message from author"
	}
}
func single_channel_reader(read <-chan any) {
	for {
		select {
		case msg := <-read:
			fmt.Println("reader received message: ", msg)
		default:
			fmt.Println("no message for reader")
		}
	}
}
func single_channel_broker(read <-chan any, write chan<- any) {
	for {
		select {
		case msg := <-read:
			fmt.Println("relaying message '", msg, "' from author to reader")
			write <- msg
		default:
			fmt.Println("no message to relay to reader")
		}
	}
}
func single_channel_main() {
	write := make(chan any) // author can write to this channel
	read := make(chan any)  // reader can read from this channel
	// broker reads from author's write channel and relays to reader's read channel
	go single_channel_broker(write, read)
	go single_channel_reader(read)
	go single_channel_author(write)
}

// multi channel relay
// have three channel to read from and three channels to write to
func multi_channel_author1(write chan<- any) {
	for i := 0; i < 10; i++ {
		write <- fmt.Sprintf("message from author1: %d", i)
	}
}
func multi_channel_author2(write chan<- any) {
	for i := 0; i < 10; i++ {
		write <- fmt.Sprintf("message from author2: %d", i)
	}
}
func multi_channel_author3(write chan<- any) {
	for i := 0; i < 10; i++ {
		write <- fmt.Sprintf("message from author3: %d", i)
	}
}

func multi_channel_reader1(read <-chan any) {
	for {
		select {
		case msg := <-read:
			fmt.Println("reader1 received message: ", msg)
		default:
			fmt.Println("no message for reader1")
		}
	}
}
func multi_channel_reader2(read <-chan any) {
	for {
		select {
		case msg := <-read:
			fmt.Println("reader2 received message: ", msg)
		default:
			fmt.Println("no message for reader2")
		}
	}
}
func multi_channel_reader3(read <-chan any) {
	for {
		select {
		case msg := <-read:
			fmt.Println("reader3 received message: ", msg)
		default:
			fmt.Println("no message for reader3")
		}
	}
}
func multi_channel_broker(read1, read2, read3 <-chan any, write1, write2, write3 chan<- any) {
	for {
		select {
		case msg1 := <-read1:
			fmt.Println("relaying message '", msg1, "' from author1 to all readers")
			write1 <- msg1
			write2 <- msg1
			write3 <- msg1
		case msg2 := <-read2:
			fmt.Println("relaying message '", msg2, "' from author2 to all readers")
			write1 <- msg2
			write2 <- msg2
			write3 <- msg2
		case msg3 := <-read3:
			fmt.Println("relaying message '", msg3, "' from author3 to all readers")
			write1 <- msg3
			write2 <- msg3
			write3 <- msg3
		default:
			fmt.Println("no message to relay")
		}
	}
}

func multi_channel_main() {
	write1 := make(chan any)
	write2 := make(chan any)
	write3 := make(chan any)

	read1 := make(chan any)
	read2 := make(chan any)
	read3 := make(chan any)

	go multi_channel_broker(write1, write2, write3, read1, read2, read3)

	go multi_channel_reader1(read1)
	go multi_channel_reader2(read2)
	go multi_channel_reader3(read3)

	go multi_channel_author1(write1)
	go multi_channel_author2(write2)
	go multi_channel_author3(write3)
}

// relay from unknown number of channel to unknown number of channels
// have unknown number of channels to read from, and unknown number of channels to write to
func u_multi_channel_reader(readMap map[int]chan any) {
	for {
		for key, value := range readMap {
			select {
			case msg := <-value:
				fmt.Println("reader ", key, "received message: ", msg)
			default:
				fmt.Println("no message for reader ", key)
			}
		}
	}
}
func u_multi_channel_author(writeMap map[int]chan any) {
	for {
		for key, value := range writeMap {
			value <- fmt.Sprintf("message from author %d", key)
		}
	}
}
func u_multi_channel_broker(readMap map[int]chan any, writeMap map[int]chan any) {
	for {
		for readKey, readValue := range readMap {
			select {
			case msg := <-readValue:
				for writeKey, writeValue := range writeMap {
					fmt.Println("relaying message '", msg, "' from author ", readKey, "to reader ", writeKey)
					writeValue <- msg
				}
			default:
				fmt.Println("no message to relay")
			}
		}
	}
}
func u_multi_channel_fanout(read <-chan any, writeMap map[int]chan any) {}
func u_multi_channel_main() {
	readMap := make(map[int]chan any)
	writeMap := make(map[int]chan any)
	go u_multi_channel_broker(writeMap, readMap)
	go u_multi_channel_reader(readMap)
	go u_multi_channel_author(writeMap)
	for i := 0; i < 10; i++ {
		readMap[i] = make(chan any)
		writeMap[i] = make(chan any)
	}
}
