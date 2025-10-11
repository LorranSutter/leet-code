package main

import "fmt"

type ExamTracker struct {
	Times  []int
	Scores []int
}

func Constructor() ExamTracker {
	et := ExamTracker{
		Times:  nil,
		Scores: nil,
	}

	return et
}

func (this *ExamTracker) Record(time int, score int) {
	// TODO Have to use prefix sum for the scores
	this.Times = append(this.Times, time)
	this.Scores = append(this.Scores, score)
}

func (this *ExamTracker) TotalScore(startTime int, endTime int) int64 {
	startTime_idx := this.Search(startTime)
	endTime_idx := this.Search(endTime)

	total_score := 0
	for i := startTime_idx; i < endTime_idx+1; i++ {
		if startTime <= this.Times[i] && this.Times[i] <= endTime {
			total_score += this.Scores[i]
		}
	}

	return int64(total_score)
}

func (this *ExamTracker) Search(time int) int {
	left, mid := 0, 0
	right := len(this.Times) - 1

	for left <= right {
		mid = left + (right-left)/2

		if this.Times[mid] == time {
			return mid
		}
		if this.Times[mid] < time {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return mid
}

func main() {
	// Your ExamTracker object will be instantiated and called as such:
	obj := Constructor()
	obj.Record(1, 98)
	fmt.Println(obj.TotalScore(1, 1) == 98)
	obj.Record(5, 99)
	fmt.Println(obj.TotalScore(1, 3) == 98)
	fmt.Println(obj.TotalScore(1, 5) == 197)
	fmt.Println(obj.TotalScore(3, 4) == 0)
	fmt.Println(obj.TotalScore(2, 5) == 99)
}
