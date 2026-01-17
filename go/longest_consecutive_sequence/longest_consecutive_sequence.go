package kata

// ::KATA START::
func longestConsecutive(nums []int) int {
  if len(nums) == 0 {
    return 0
  }

  numSet := make(map[int]bool)
  for _, num := range nums{
    numSet[num]=true
  }

  longest:=0
  for num := range numSet {
    if _, exists := numSet[num-1]; !exists {
      curr := num
      streak := 1

      for numSet[curr+1] {
        curr++
        streak++
      }
      longest=max(longest, streak)
    }
  }

  return longest

  // seq := 1
  // longest := 0
  // for _, num := range nums{
  //   next := num + 1
  //   _, exists := numSet[next]
  //   for exists {
  //     seq++
  //     next+=1
  //     _, exists = numSet[next]
  //   }
  //   longest=max(seq,longest)
  //   seq=1
  // }
  // return longest
}

// ::KATA END::
// func longestConsecutive(nums []int) int {
// 	slices.Sort(nums)
//
//   seq := 1 
//   prev := math.MinInt
//   longest := 0
// 	for _, num := range nums {
//     if prev+1==num {
//       seq++
//     } else {
//       seq=1
//     }
//
//     prev=num
//     longest=max(seq, longest)
// 	}
//
// 	return longest
// }
