package birdwatcher


// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, val := range birdsPerDay {
		total += val
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	var startIdx int;
	if week == 1 {
		startIdx = 0
	} else {
		startIdx = (week - 1) * 7
	}
	endIdx := (week * 7) -1
	for i:=startIdx; i <= endIdx; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := range birdsPerDay {
		if idx == 0 {
			birdsPerDay[idx] += 1
		} else if idx % 2 == 0 {
			birdsPerDay[idx] += 1
		}

	}
	return birdsPerDay
}
