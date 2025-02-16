func main() {
    c := make(chan int)
    done := make(chan bool)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
        done <- true
    }()
    for i := range c {
        fmt.Println(i)
    }
    <-done
} 