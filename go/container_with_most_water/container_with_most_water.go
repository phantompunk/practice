package kata

// ::KATA START::
func maxArea(height []int) int {
  left, right := 0, len(height)-1    

  maxWater:=0
  for left<right{
    h:=min(height[left],height[right])
    w:=right-left
    maxWater=max(maxWater,h*w)
    if height[left] < height[right] {
      left++
    } else {
      right--
    }
  }
  return maxWater
}
// ::KATA END::
