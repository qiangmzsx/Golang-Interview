package main

import "sort"

//【面试题】合并重叠区间
//给定一组 区间，合并所有重叠的 区间。
//
//例如：
//
//给定：[1,3],[2,6],[8,10],[15,18]
//返回：[1,6],[8,10],[15,18]

// https://studygolang.com/topics/3853

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]Interval, 0)
	swap := Interval{}
	for k, v := range intervals {
		if k == 0 {
			swap = v
			continue
		}
		if v.Start <= swap.End {
			swap.End = v.End
		} else {
			res = append(res, swap)
			swap = v
		}
	}
	res = append(res, swap)
	return res
}

func main() {

}
