package main

type Job interface {
	Do()
}

// 영역을 나누는 방법

type SquareJob struct {
	index Int
	// contest State // 필요한 정보
}

func (j *SquareJob) Do(){
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 *time.second)
	fmt.Printf("%d 작업 완료 - 결과 %d \n", j.index, j.index *j.index)
}

func main() {
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
		
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
		
	}
	wg.Wait()
}